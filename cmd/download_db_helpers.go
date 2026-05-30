package cmd

import (
	"context"
	"fmt"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	cfg "github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/db"
	"github.com/phani-kb/multilog"
)

const runtimeSourceFile = "runtime-config"

func syncSourcesToDB(
	ctx context.Context,
	logger *multilog.Logger,
	repo *db.SourcesRepo,
	sourcesConfigs []cfg.SourcesConfig,
) error {
	for _, sourcesConfig := range sourcesConfigs {
		if _, _, err := repo.ImportSourcesFromConfig(ctx, logger, sourcesConfig, runtimeSourceFile); err != nil {
			return fmt.Errorf("syncing source definitions to database: %w", err)
		}
	}

	return nil
}

func loadDownloadSummaries(ctx context.Context, logger *multilog.Logger) ([]c.DownloadSummary, error) {
	dbPath := getDBPath()
	database, openErr := db.Open(ctx, logger, dbPath, false)
	if openErr != nil {
		return nil, fmt.Errorf("opening database for download summaries: %w", openErr)
	}
	defer database.CloseLogError(logger)

	downloadsRepo := db.NewDownloadsRepo(database)
	summaries, listErr := downloadsRepo.ListDownloadSummaries(constants.DownloadDir)
	if listErr != nil {
		return nil, fmt.Errorf("loading download summaries from database: %w", listErr)
	}

	return summaries, nil
}

func loadPreviousDownloadSummary(
	_ *multilog.Logger,
	downloadsRepo *db.DownloadsRepo,
	sourceName string,
	targetFilePath string,
) (*c.DownloadSummary, error) {
	if downloadsRepo == nil {
		return nil, nil
	}

	summaries, err := downloadsRepo.GetDownloadSummaryBySourceName(sourceName, constants.DownloadDir)
	if err != nil {
		return nil, fmt.Errorf("loading previous download summary for %s from database: %w", sourceName, err)
	}
	if len(summaries) == 0 {
		return nil, nil
	}

	if targetFilePath != "" {
		for i := range summaries {
			if summaries[i].Filepath == targetFilePath {
				return &summaries[i], nil
			}
		}
	}

	return &summaries[0], nil
}
