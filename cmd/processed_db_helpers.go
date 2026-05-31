package cmd

import (
	"context"
	"fmt"
	"sort"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	cfg "github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/db"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

func loadProcessedSummariesFromDB(ctx context.Context, logger *multilog.Logger) ([]c.ProcessedSummary, error) {
	dbPath := getDBPath()
	database, openErr := db.Open(ctx, logger, dbPath, false)
	if openErr != nil {
		return nil, fmt.Errorf("opening database for processed summaries: %w", openErr)
	}
	defer database.CloseLogError(logger)

	processedRepo := db.NewProcessedRepo(database)
	summaries, listErr := processedRepo.ListProcessedSummaries(constants.ProcessedDir)
	if listErr != nil {
		return nil, fmt.Errorf("loading processed summaries from database: %w", listErr)
	}

	return summaries, nil
}

func loadProcessedInputsForConsolidation(
	ctx context.Context,
	logger *multilog.Logger,
	consolidationType string,
) ([]c.ProcessedSummary, []string, []c.ProcessedFile, error) {
	summaries, err := loadProcessedSummariesFromDB(ctx, logger)
	if err != nil {
		return nil, nil, nil, err
	}

	enabledSummaries := filterEnabledProcessedSummaries(logger, summaries, consolidationType)
	sort.Slice(enabledSummaries, func(i, j int) bool {
		return u.CaseInsensitiveLess(enabledSummaries[i].Name, enabledSummaries[j].Name)
	})

	genericSourceTypes := extractGenericSourceTypesFromSummaries(enabledSummaries)
	processedFiles := cfg.GetAllProcessedFiles(enabledSummaries)

	logger.Infof(
		"Processed summaries count: %d, generic source types count: %d, files count: %d, consolidation type: %s",
		len(enabledSummaries),
		len(genericSourceTypes),
		len(processedFiles),
		consolidationType,
	)

	return enabledSummaries, genericSourceTypes, processedFiles, nil
}

func filterEnabledProcessedSummaries(
	logger *multilog.Logger,
	summaries []c.ProcessedSummary,
	consolidationType string,
) []c.ProcessedSummary {
	enabledSummaries := make([]c.ProcessedSummary, 0, len(summaries))
	for _, summary := range summaries {
		if cfg.IsEnabledSourceForConsolidation(summary.Name, SourcesConfigs, *AppConfig, consolidationType) {
			enabledSummaries = append(enabledSummaries, summary)
		} else {
			logger.Debugf("Skipping summary %s: not enabled for consolidation", summary.Name)
		}
	}
	return enabledSummaries
}

func extractGenericSourceTypesFromSummaries(summaries []c.ProcessedSummary) []string {
	sourceTypeMap := make(map[string]struct{})
	for _, summary := range summaries {
		for _, processedFile := range summary.ValidFiles {
			sourceTypeMap[processedFile.GenericSourceType] = struct{}{}
		}
		for _, processedFile := range summary.InvalidFiles {
			sourceTypeMap[processedFile.GenericSourceType] = struct{}{}
		}
	}

	genericSourceTypes := make([]string, 0, len(sourceTypeMap))
	for sourceType := range sourceTypeMap {
		genericSourceTypes = append(genericSourceTypes, sourceType)
	}
	u.SortCaseInsensitiveStrings(genericSourceTypes)
	return genericSourceTypes
}
