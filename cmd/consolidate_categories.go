package cmd

import (
	"context"
	"sync"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/db"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/spf13/cobra"
)

var consolidateCategoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "Generate category-based consolidated lists (ads, malware, privacy, etc)",
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Infof("Generating category-based consolidated lists...")
		ctx := context.Background()

		processedSummaries, genericSourceTypes, processedFiles, loadErr := loadProcessedInputsForConsolidation(
			ctx,
			Logger,
			"categories",
		)
		if loadErr != nil {
			Logger.Errorf("Failed to load processed summaries from database: %v", loadErr)
			return
		}
		if len(processedSummaries) == 0 {
			Logger.Errorf("No processed summaries found")
			return
		}

		database, consolidatedRepo, dbErr := openConsolidatedRepo(ctx, Logger, "category")
		if dbErr != nil {
			Logger.Errorf("Failed to open consolidated repo: %v", dbErr)
			return
		}
		defer database.CloseLogError(Logger)

		var persistMu sync.Mutex

		// Get unique categories from all processed files
		categories := getUniqueCategories(processedFiles)
		Logger.Infof("Found %d unique categories: %v", len(categories), categories)

		// Process each category and create consolidated lists
		for _, category := range categories {
			processCategoryConsolidation(
				Logger,
				category,
				processedFiles,
				genericSourceTypes,
				consolidatedRepo,
				ctx,
				&persistMu,
			)
		}

		Logger.Infof("Categories consolidation complete")
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

	u.SortCaseInsensitiveStrings(categories)

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
	consolidatedRepo *db.ConsolidatedRepo,
	ctx context.Context,
	persistMu *sync.Mutex,
) map[string][]c.ConsolidatedSummary {
	config := ProcessingConfig{
		Identifier:         category,
		IdentifierField:    "Category",
		ProcessedFiles:     processedFiles,
		GenericSourceTypes: genericSourceTypes,
		GetFilesFunc:       getFilesForCategory,
		ConsolidateFunc:    consolidateByCategory,
		AllowFilterByType:  nil, // no cross-category filtering
		ConsolidatedRepo:   consolidatedRepo,
		DBCtx:              ctx,
		PersistMu:          persistMu,
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
