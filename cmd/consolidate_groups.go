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

var consolidateGroupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "Generate different sized consolidated lists (mini, lite, normal, big)",
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Infof("Generating sized consolidated lists...")
		ctx := context.Background()

		processedSummaries, genericSourceTypes, processedFiles, loadErr := loadProcessedInputsForConsolidation(
			ctx,
			Logger,
			"groups",
		)
		if loadErr != nil {
			Logger.Errorf("Failed to load processed summaries from database: %v", loadErr)
			return
		}

		if len(processedSummaries) == 0 {
			Logger.Errorf("No processed summaries found")
			return
		}

		database, consolidatedRepo, dbErr := openConsolidatedRepo(ctx, Logger, "group")
		if dbErr != nil {
			Logger.Errorf("Failed to open consolidated repo: %v", dbErr)
			return
		}
		defer database.CloseLogError(Logger)

		var persistMu sync.Mutex

		// Process each size group and create consolidated lists
		for _, group := range constants.SizeGroups {
			processGroupConsolidationWithAllow(
				Logger,
				group,
				processedFiles,
				genericSourceTypes,
				consolidatedRepo,
				ctx,
				&persistMu,
			)
		}

		Logger.Infof("Groups consolidation complete")
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

func processGroupConsolidationWithAllow(
	logger *multilog.Logger,
	group string,
	processedFiles []c.ProcessedFile,
	genericSourceTypes []string,
	consolidatedRepo *db.ConsolidatedRepo,
	ctx context.Context,
	persistMu *sync.Mutex,
) map[string][]c.ConsolidatedSummary {
	allowByType, _, _, _, _, _ := GetCachedResolutionSets(logger, processedFiles)
	config := ProcessingConfig{
		Identifier:         group,
		IdentifierField:    "Group",
		ProcessedFiles:     processedFiles,
		GenericSourceTypes: genericSourceTypes,
		GetFilesFunc:       getFilesForGroup,
		ConsolidateFunc:    consolidateByGroup,
		AllowFilterByType:  allowByType,
		ConsolidatedRepo:   consolidatedRepo,
		DBCtx:              ctx,
		PersistMu:          persistMu,
	}

	return processConsolidationWithTransform(logger, config)
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
