package cmd

import (
	"os"
	"path/filepath"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	cfg "github.com/phani-kb/dns-toolkit/internal/config"
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
			groupResults := processGroupConsolidation(
				Logger,
				group,
				processedFiles,
				genericSourceTypes,
			)
			for identifier, summaries := range groupResults {
				consolidatedSummariesByGroup[identifier] = append(
					consolidatedSummariesByGroup[identifier],
					summaries...,
				)
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

// getFilesForGroup filters processed files by group
func getFilesForGroup(processedFiles []c.ProcessedFile, group string) []c.ProcessedFile {
	var groupFiles []c.ProcessedFile
	for _, file := range processedFiles {
		for _, fileGroup := range file.Groups {
			if fileGroup == group && file.Valid {
				groupFiles = append(groupFiles, file)
				break
			}
		}
	}
	return groupFiles
}

// processGroupConsolidation processes consolidation for a specific group
func processGroupConsolidation(
	logger *multilog.Logger,
	group string,
	processedFiles []c.ProcessedFile,
	genericSourceTypes []string,
) map[string][]c.ConsolidatedSummary {
	config := ProcessingConfig{
		Identifier:         group,
		IdentifierField:    "Group",
		ProcessedFiles:     processedFiles,
		GenericSourceTypes: genericSourceTypes,
		GetFilesFunc:       getFilesForGroup,
		ConsolidateFunc:    consolidateByGroup,
	}

	return processIdentifierConsolidation(logger, config)
}

// consolidateByGroup consolidates files for a specific size group
func consolidateByGroup(
	logger *multilog.Logger,
	genericSourceType, listType, group string,
	entriesToIgnore u.StringSet,
	processedFiles []c.ProcessedFile,
) (u.StringSet, c.ConsolidatedSummary) {
	params := ConsolidationParams{
		GenericSourceType: genericSourceType,
		ListType:          listType,
		Identifier:        group,
		OutputDir:         constants.ConsolidatedGroupsDir,
		IdentifierField:   "Group",
	}

	return consolidateGeneric(logger, params, entriesToIgnore, processedFiles)
}
