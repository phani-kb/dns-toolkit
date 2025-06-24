package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	cfg "github.com/phani-kb/dns-toolkit/internal/config"
	con "github.com/phani-kb/dns-toolkit/internal/consolidators"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/spf13/cobra"
)

var consolidateGroupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "Generate different sized consolidated lists (mini, lite, normal, big)",
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Infof("Generating sized consolidated lists...")

		if err := u.EnsureDirectoryExists(Logger, constants.ConsolidatedGroupsDir); err != nil {
			Logger.Errorf("Failed to create consolidated groups directory: %v", err)
			os.Exit(1)
		}
		if err := u.EnsureDirectoryExists(Logger, constants.SummaryDir); err != nil {
			Logger.Errorf("Failed to create summary directory: %v", err)
			os.Exit(1)
		}

		processedSummaries, genericSourceTypes, processedFiles := cfg.GetProcessedSummaries(
			Logger,
			SourcesConfigs,
			*AppConfig,
		)

		if len(processedSummaries) == 0 {
			Logger.Errorf("No processed summaries found")
			return
		}

		// Maps to store consolidated summaries by group
		consolidatedSummariesByGroup := make(map[string][]c.ConsolidatedSummary)
		for _, group := range constants.SizeGroups {
			consolidatedSummariesByGroup[group] = []c.ConsolidatedSummary{}
		}

		// Process each size group and create consolidated lists
		for _, group := range constants.SizeGroups {
			Logger.Infof("Processing size group: %s", group)

			// Filter processed files by group
			var groupFiles []c.ProcessedFile
			for _, file := range processedFiles {
				for _, fileGroup := range file.Groups {
					if fileGroup == group && file.Valid {
						groupFiles = append(groupFiles, file)
						break
					}
				}
			}

			if len(groupFiles) == 0 {
				Logger.Infof("No files found for group: %s", group)
				continue
			}

			// Create a map of allowlisted entries by source type
			allowlistEntriesByType := make(map[string]u.StringSet)

			// First process allowlists for each group and source type
			for _, gst := range genericSourceTypes {
				var allowlistFiles []c.ProcessedFile
				for _, file := range groupFiles {
					if file.GenericSourceType == gst &&
						file.ListType == constants.ListTypeAllowlist {
						allowlistFiles = append(allowlistFiles, file)
					}
				}

				if len(allowlistFiles) > 0 {
					entries, allowlistSummary := consolidateByGroup(
						Logger,
						gst,
						constants.ListTypeAllowlist,
						group,
						u.NewStringSet([]string{}),
						allowlistFiles,
					)
					allowlistEntriesByType[gst] = entries
					allowlistSummary.Group = group
					consolidatedSummariesByGroup[group] = append(
						consolidatedSummariesByGroup[group],
						allowlistSummary,
					)
				} else {
					allowlistEntriesByType[gst] = u.NewStringSet([]string{})
				}
			}

			// Then process blocklists using the allowlists from above
			for _, gst := range genericSourceTypes {
				var blocklistFiles []c.ProcessedFile
				for _, file := range groupFiles {
					if file.GenericSourceType == gst &&
						file.ListType == constants.ListTypeBlocklist {
						blocklistFiles = append(blocklistFiles, file)
					}
				}

				if len(blocklistFiles) > 0 {
					allowlistEntries := allowlistEntriesByType[gst]
					_, blocklistSummary := consolidateByGroup(
						Logger,
						gst,
						constants.ListTypeBlocklist,
						group,
						allowlistEntries,
						blocklistFiles,
					)
					blocklistSummary.Group = group
					consolidatedSummariesByGroup[group] = append(
						consolidatedSummariesByGroup[group],
						blocklistSummary,
					)
				}
			}
		}

		// Create consolidated groups summaries
		var consolidatedGroupsSummaries []c.ConsolidatedGroupsSummary
		timestamp := u.GetTimestamp()

		for _, group := range constants.SizeGroups {
			summaries := consolidatedSummariesByGroup[group]
			if len(summaries) > 0 {
				consolidatedGroupsSummary := c.ConsolidatedGroupsSummary{
					Group:                     group,
					ConsolidatedSummaries:     summaries,
					LastConsolidatedTimestamp: timestamp,
				}
				consolidatedGroupsSummaries = append(
					consolidatedGroupsSummaries,
					consolidatedGroupsSummary,
				)
			}
		}

		// Save all consolidated summaries to the regular consolidated summary file if not skipped
		if !skipConsolidatedSummary {
			var allConsolidatedSummaries []c.ConsolidatedSummary
			for _, summariesByGroup := range consolidatedSummariesByGroup {
				allConsolidatedSummaries = append(allConsolidatedSummaries, summariesByGroup...)
			}

			if len(allConsolidatedSummaries) > 0 {
				summaryFile := filepath.Join(
					constants.SummaryDir,
					constants.DefaultSummaryFiles["consolidated"],
				)
				summariesCount, err := u.SaveSummaries(
					Logger,
					allConsolidatedSummaries,
					summaryFile,
					c.ConsolidatedSummaryLessFunc,
				)
				if err != nil {
					Logger.Errorf("Error saving consolidated summaries to %s: %v", summaryFile, err)
				} else {
					if summariesCount > 0 {
						Logger.Infof("Saved consolidated summaries to %s", summaryFile)
					}
				}
			}
		} else {
			Logger.Infof("Skipping regular consolidated summary file as requested")
		}

		// Save grouped consolidated groups summaries to the new summary file
		if len(consolidatedGroupsSummaries) > 0 {
			GroupsSummaryFile := filepath.Join(
				constants.SummaryDir,
				constants.DefaultSummaryFiles["consolidated_groups"],
			)
			summariesCount, err := u.SaveSummaries(
				Logger,
				consolidatedGroupsSummaries,
				GroupsSummaryFile,
				c.ConsolidatedGroupsSummaryLessFunc,
			)
			if err != nil {
				Logger.Errorf(
					"Error saving consolidated groups summaries to %s: %v",
					GroupsSummaryFile,
					err,
				)
			} else {
				if summariesCount > 0 {
					Logger.Infof(
						"Successfully saved %d size group summaries to %s",
						len(consolidatedGroupsSummaries),
						GroupsSummaryFile,
					)
				}
			}
		}
	},
}

// consolidateByGroup consolidates files for a specific size group
func consolidateByGroup(
	logger *multilog.Logger,
	genericSourceType, listType, group string,
	entriesToIgnore u.StringSet,
	processedFiles []c.ProcessedFile,
) (u.StringSet, c.ConsolidatedSummary) {
	logger.Debugf(
		"Starting consolidation for %s %s in group %s",
		listType,
		genericSourceType,
		group,
	)

	if len(processedFiles) == 0 {
		logger.Debugf(
			"No processed files found for %s %s in group %s",
			listType,
			genericSourceType,
			group,
		)
		return u.NewStringSet([]string{}), c.ConsolidatedSummary{}
	}

	consolidator, exists := con.Consolidators.GetConsolidator(genericSourceType, listType)
	if !exists {
		logger.Warnf(
			"No consolidator found for generic source type: %s, list type: %s",
			genericSourceType,
			listType,
		)
		return u.NewStringSet([]string{}), c.ConsolidatedSummary{}
	}

	consolidatedEntries, fileInfos := consolidator.Consolidate(logger, processedFiles)
	consolidatedFileStrings := getFileStrings(fileInfos)

	if len(consolidatedEntries) > 0 {
		logger.Infof(
			"Consolidated %s %s %d entry(s) for group %s",
			listType,
			genericSourceType,
			len(consolidatedEntries),
			group,
		)
	}

	allEntries, ignoredEntries := consolidator.FilterEntries(
		logger,
		consolidatedEntries,
		entriesToIgnore,
	)

	// Create a summary with group information
	consolidatedSummary := c.ConsolidatedSummary{
		Type:                      genericSourceType,
		FilesCount:                len(fileInfos),
		Files:                     consolidatedFileStrings,
		Valid:                     true,
		Count:                     len(allEntries),
		IgnoredEntriesCount:       len(ignoredEntries),
		ListType:                  listType,
		Group:                     group,
		LastConsolidatedTimestamp: u.GetTimestamp(),
	}

	if len(ignoredEntries) > 0 {
		logger.Infof(
			"Ignored %s %s %d entry(s) for group %s",
			listType,
			genericSourceType,
			len(ignoredEntries),
			group,
		)
		filenamePrefix := fmt.Sprintf("%s_%s_%s", group, genericSourceType, listType)
		ignoredFilePath := filepath.Join(
			constants.ConsolidatedGroupsDir,
			filenamePrefix+"_ignored.txt",
		)
		err := consolidator.SaveEntries(logger, ignoredEntries, ignoredFilePath)
		if err != nil {
			logger.Errorf("Error writing ignored entry(s) to file %s: %v", ignoredFilePath, err)
		} else {
			consolidatedSummary.IgnoredFilepath = ignoredFilePath
		}
	}

	if len(allEntries) <= 0 {
		logger.Infof(
			"No entry(s) to consolidate for %s %s in group %s",
			listType,
			genericSourceType,
			group,
		)
		return u.NewStringSet([]string{}), c.ConsolidatedSummary{}
	}

	// Use group-specific filename
	filenamePrefix := fmt.Sprintf("%s_%s_%s", group, genericSourceType, listType)
	consolidatedFilePath := filepath.Join(constants.ConsolidatedGroupsDir, filenamePrefix+".txt")
	consolidatedSummary.Filepath = consolidatedFilePath

	err := consolidator.SaveEntries(logger, allEntries, consolidatedSummary.Filepath)
	if err != nil {
		logger.Errorf("Error writing entry(s) to file %s: %v", consolidatedSummary.Filepath, err)
	} else {
		if calculateChecksum || AppConfig.DNSToolkit.FilesChecksum.Enabled {
			consolidatedSummary.Checksum = u.CalculateChecksum(
				logger,
				consolidatedSummary.Filepath,
				AppConfig.DNSToolkit.FilesChecksum.Algorithm,
			)
		}
	}

	logger.Debugf(
		"Finished consolidation for %s %s in group %s",
		listType,
		genericSourceType,
		group,
	)
	return allEntries, consolidatedSummary
}
