package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	cfg "github.com/phani-kb/dns-toolkit/internal/config"
	con "github.com/phani-kb/dns-toolkit/internal/consolidators"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/spf13/cobra"
)

var consolidateCategoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "Generate category-based consolidated lists (advertising, malware, privacy, etc)",
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Infof("Generating category-based consolidated lists...")

		if err := u.EnsureDirectoryExists(Logger, constants.ConsolidatedCategoriesDir); err != nil {
			Logger.Errorf("Failed to create consolidated categories directory: %v", err)
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

		// Get unique categories from all processed files
		categories := getUniqueCategories(processedFiles)
		Logger.Infof("Found %d unique categories: %v", len(categories), categories)

		// Maps to store consolidated summaries by category
		consolidatedSummariesByCategory := make(map[string][]c.ConsolidatedSummary)
		for _, category := range categories {
			consolidatedSummariesByCategory[category] = []c.ConsolidatedSummary{}
		}

		// Process each category and create consolidated lists
		for _, category := range categories {
			Logger.Infof("Processing category: %s", category)
			// Filter processed files by category
			var categoryFiles []c.ProcessedFile
			for _, file := range processedFiles {
				for _, fileCategory := range file.Categories {
					if fileCategory == category && file.Valid {
						categoryFiles = append(categoryFiles, file)
						break
					}
				}
			}
			if len(categoryFiles) == 0 {
				Logger.Infof("No files found for category: %s", category)
				continue
			}

			// Create a map to allowlisted entries by source type
			allowlistEntriesByType := make(map[string]u.StringSet)
			// First process allowlists for each category and source type
			for _, gst := range genericSourceTypes {
				var allowlistFiles []c.ProcessedFile
				for _, file := range categoryFiles {
					if file.GenericSourceType == gst &&
						file.ListType == constants.ListTypeAllowlist {
						allowlistFiles = append(allowlistFiles, file)
					}
				}
				if len(allowlistFiles) > 0 {
					entries, allowlistSummary := consolidateByCategory(
						Logger,
						gst,
						constants.ListTypeAllowlist,
						category,
						u.NewStringSet([]string{}),
						allowlistFiles,
					)
					allowlistEntriesByType[gst] = entries
					allowlistSummary.Category = category
					consolidatedSummariesByCategory[category] = append(
						consolidatedSummariesByCategory[category],
						allowlistSummary,
					)
				} else {
					allowlistEntriesByType[gst] = u.NewStringSet([]string{})
				}
			}

			// Then process blocklists using the allowlists from above
			for _, gst := range genericSourceTypes {
				var blocklistFiles []c.ProcessedFile
				for _, file := range categoryFiles {
					if file.GenericSourceType == gst &&
						file.ListType == constants.ListTypeBlocklist {
						blocklistFiles = append(blocklistFiles, file)
					}
				}
				if len(blocklistFiles) > 0 {
					allowlistEntries := allowlistEntriesByType[gst]
					_, blocklistSummary := consolidateByCategory(
						Logger,
						gst,
						constants.ListTypeBlocklist,
						category,
						allowlistEntries,
						blocklistFiles,
					)
					blocklistSummary.Category = category
					consolidatedSummariesByCategory[category] = append(
						consolidatedSummariesByCategory[category],
						blocklistSummary,
					)
				}
			}
		}

		// Create consolidated categories summaries
		var consolidatedCategoriesSummaries []c.ConsolidatedCategoriesSummary
		timestamp := u.GetTimestamp()
		for _, category := range categories {
			summaries := consolidatedSummariesByCategory[category]
			if len(summaries) > 0 {
				consolidatedCategoriesSummary := c.ConsolidatedCategoriesSummary{
					Category:                  category,
					ConsolidatedSummaries:     summaries,
					LastConsolidatedTimestamp: timestamp,
				}
				consolidatedCategoriesSummaries = append(
					consolidatedCategoriesSummaries,
					consolidatedCategoriesSummary,
				)
			}
		}

		// Save all consolidated summaries to the regular consolidated summary file if not skipped
		if !skipConsolidatedSummary {
			var allConsolidatedSummaries []c.ConsolidatedSummary
			for _, summariesByCategory := range consolidatedSummariesByCategory {
				allConsolidatedSummaries = append(allConsolidatedSummaries, summariesByCategory...)
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

		// Save grouped consolidated categories summaries to the new summary file
		if len(consolidatedCategoriesSummaries) > 0 {
			categoriesSummaryFile := filepath.Join(
				constants.SummaryDir,
				constants.DefaultSummaryFiles["consolidated_categories"],
			)
			summariesCount, err := u.SaveSummaries(
				Logger,
				consolidatedCategoriesSummaries,
				categoriesSummaryFile,
				c.ConsolidatedCategoriesSummaryLessFunc,
			)
			if err != nil {
				Logger.Errorf(
					"Error saving consolidated categories summaries to %s: %v",
					categoriesSummaryFile,
					err,
				)
			} else {
				if summariesCount > 0 {
					Logger.Infof(
						"Successfully saved %d category summaries to %s",
						len(consolidatedCategoriesSummaries),
						categoriesSummaryFile,
					)
				}
			}
		}
	},
}

// getUniqueCategories returns a slice of unique categories from all processed files
func getUniqueCategories(processedFiles []c.ProcessedFile) []string {
	categoriesSet := make(map[string]struct{})
	for _, file := range processedFiles {
		for _, category := range file.Categories {
			if category != "" {
				categoriesSet[category] = struct{}{}
			}
		}
	}

	// Convert the map to a slice
	categories := make([]string, 0, len(categoriesSet))
	for category := range categoriesSet {
		categories = append(categories, category)
	}

	sort.Strings(categories)

	return categories
}

// consolidateByCategory consolidates files for a specific category
func consolidateByCategory(
	logger *multilog.Logger,
	genericSourceType, listType, category string,
	entriesToIgnore u.StringSet,
	processedFiles []c.ProcessedFile,
) (u.StringSet, c.ConsolidatedSummary) {
	logger.Debugf(
		"Starting consolidation for %s %s in category %s",
		listType,
		genericSourceType,
		category,
	)
	if len(processedFiles) == 0 {
		logger.Debugf(
			"No processed files found for %s %s in category %s",
			listType,
			genericSourceType,
			category,
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
			"Consolidated %s %s %d entry(s) for category %s",
			listType,
			genericSourceType,
			len(consolidatedEntries),
			category,
		)
	}

	allEntries, ignoredEntries := consolidator.FilterEntries(
		logger,
		consolidatedEntries,
		entriesToIgnore,
	)

	// Create a summary with category information
	consolidatedSummary := c.ConsolidatedSummary{
		Type:                      genericSourceType,
		FilesCount:                len(fileInfos),
		Files:                     consolidatedFileStrings,
		Valid:                     true,
		Count:                     len(allEntries),
		IgnoredEntriesCount:       len(ignoredEntries),
		ListType:                  listType,
		Category:                  category,
		LastConsolidatedTimestamp: u.GetTimestamp(),
	}

	if len(ignoredEntries) > 0 {
		logger.Infof(
			"Ignored %s %s %d entry(s) for category %s",
			listType,
			genericSourceType,
			len(ignoredEntries),
			category,
		)
		filenamePrefix := fmt.Sprintf("%s_%s_%s", category, genericSourceType, listType)
		ignoredFilePath := filepath.Join(
			constants.ConsolidatedCategoriesDir,
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
			"No entry(s) to consolidate for %s %s in category %s",
			listType,
			genericSourceType,
			category,
		)
		return u.NewStringSet([]string{}), c.ConsolidatedSummary{}
	}

	// Use category-specific filename
	filenamePrefix := fmt.Sprintf("%s_%s_%s", category, genericSourceType, listType)
	consolidatedFilePath := filepath.Join(
		constants.ConsolidatedCategoriesDir,
		filenamePrefix+".txt",
	)
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
		"Finished consolidation for %s %s in category %s",
		listType,
		genericSourceType,
		category,
	)
	return allEntries, consolidatedSummary
}

func init() {
	consolidateCmd.AddCommand(consolidateCategoriesCmd)
}
