package cmd

import (
	"os"
	"path/filepath"
	"sort"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	cfg "github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/spf13/cobra"
)

var consolidateCategoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "Generate category-based consolidated lists (ads, malware, privacy, etc)",
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

		processedSummaries, genericSourceTypes, processedFiles := cfg.GetProcessedSummariesForConsolidation(
			Logger,
			SourcesConfigs,
			*AppConfig,
			"categories",
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
			categoryResults := processCategoryConsolidation(
				Logger,
				category,
				processedFiles,
				genericSourceTypes,
			)
			for identifier, summaries := range categoryResults {
				consolidatedSummariesByCategory[identifier] = append(
					consolidatedSummariesByCategory[identifier],
					summaries...,
				)
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
					constants.DefaultSummaryFiles["consolidated_categories"],
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

// getFilesForCategory filters processed files by category
func getFilesForCategory(processedFiles []c.ProcessedFile, category string) []c.ProcessedFile {
	var categoryFiles []c.ProcessedFile
	for _, file := range processedFiles {
		for _, fileCategory := range file.Categories {
			if fileCategory == category && file.Valid {
				categoryFiles = append(categoryFiles, file)
				break
			}
		}
	}
	return categoryFiles
}

// processCategoryConsolidation processes consolidation for a specific category
func processCategoryConsolidation(
	logger *multilog.Logger,
	category string,
	processedFiles []c.ProcessedFile,
	genericSourceTypes []string,
) map[string][]c.ConsolidatedSummary {
	config := ProcessingConfig{
		Identifier:         category,
		IdentifierField:    "Category",
		ProcessedFiles:     processedFiles,
		GenericSourceTypes: genericSourceTypes,
		GetFilesFunc:       getFilesForCategory,
		ConsolidateFunc:    consolidateByCategory,
	}

	return processConsolidationWithTransform(logger, config)
}

// consolidateByCategory consolidates files for a specific category
func consolidateByCategory(
	logger *multilog.Logger,
	genericSourceType, listType, category string,
	entriesToIgnore u.StringSet,
	processedFiles []c.ProcessedFile,
) (u.StringSet, c.ConsolidatedSummary) {
	params := ConsolidationParams{
		GenericSourceType: genericSourceType,
		ListType:          listType,
		Identifier:        category,
		OutputDir:         constants.ConsolidatedCategoriesDir,
		IdentifierField:   "Category",
	}

	return consolidateGeneric(logger, params, entriesToIgnore, processedFiles)
}
