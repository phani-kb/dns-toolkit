package cmd

import (
	"context"
	"fmt"
	"sync"

	"github.com/phani-kb/dns-toolkit/internal/db"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

func persistConsolidatedEntries(
	ctx context.Context,
	logger *multilog.Logger,
	repo *db.ConsolidatedRepo,
	mu *sync.Mutex,
	entries u.StringSet,
	genericSourceType, listType, consolidationType string,
	groupName, category string,
	valid bool,
) error {
	if entries == nil || entries.Size() == 0 {
		return nil
	}

	rows := make([]db.ConsolidatedEntryRow, 0, entries.Size())
	for entry := range entries {
		rows = append(rows, db.ConsolidatedEntryRow{
			Entry:             entry,
			GenericSourceType: genericSourceType,
			ListType:          listType,
			ConsolidationType: consolidationType,
			GroupName:         groupName,
			Category:          category,
			Valid:             valid,
			SourceCount:       1,
		})
	}

	if mu != nil {
		mu.Lock()
		defer mu.Unlock()
	}

	count, err := repo.BulkInsertEntries(ctx, rows)
	if err != nil {
		return fmt.Errorf("persisting %s %s consolidated entries: %w", genericSourceType, listType, err)
	}
	logger.Infof(
		"Persisted %d %s %s entries to database (consolidation_type=%s)",
		count,
		genericSourceType,
		listType,
		consolidationType,
	)
	return nil
}

func openConsolidatedRepo(
	ctx context.Context,
	logger *multilog.Logger,
	consolidationType string,
) (*db.DB, *db.ConsolidatedRepo, error) {
	dbPath := getDBPath()
	database, err := db.Open(ctx, logger, dbPath, false)
	if err != nil {
		return nil, nil, fmt.Errorf("opening database: %w", err)
	}

	repo := db.NewConsolidatedRepo(database)
	if clearErr := repo.ClearConsolidated(consolidationType); clearErr != nil {
		database.CloseLogError(logger)
		return nil, nil, fmt.Errorf("clearing %s consolidated entries: %w", consolidationType, clearErr)
	}
	logger.Infof("Cleared existing %s consolidated entries", consolidationType)

	return database, repo, nil
}
