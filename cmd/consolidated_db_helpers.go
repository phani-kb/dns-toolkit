package cmd

import (
	"context"
	"fmt"

	"github.com/phani-kb/dns-toolkit/internal/db"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

func persistConsolidatedEntries(
	ctx context.Context,
	logger *multilog.Logger,
	repo *db.ConsolidatedRepo,
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

func filterTypesForConsolidation(
	ctx context.Context,
	logger *multilog.Logger,
	repo *db.ConsolidatedRepo,
	consolidationType string,
	genericSourceTypes []string,
	force bool,
) (typesToProcess []string, typeFingerprints map[string]string) {
	typesToProcess = make([]string, 0, len(genericSourceTypes))
	typeFingerprints = make(map[string]string)
	for _, gst := range genericSourceTypes {
		fp, fpErr := repo.ComputeTypeFingerprint(consolidationType, gst)
		if fpErr == nil && fp != "" {
			typeFingerprints[gst] = fp
			if !force {
				stored := repo.GetStoredTypeFingerprint(consolidationType, gst)
				if stored == fp && repo.HasConsolidatedDataForType(consolidationType, gst) {
					logger.Infof("Skipping %s %s consolidation (unchanged)", consolidationType, gst)
					continue
				}
			}
		}
		if clearErr := repo.ClearConsolidatedRowsForType(ctx, consolidationType, gst); clearErr != nil {
			logger.Errorf("Failed to clear %s/%s: %v", consolidationType, gst, clearErr)
			continue
		}
		typesToProcess = append(typesToProcess, gst)
	}
	return typesToProcess, typeFingerprints
}

// saveConsolidationFingerprints persists per-type and global fingerprints so the
// next run can skip unchanged types/consolidations.
func saveConsolidationFingerprints(
	logger *multilog.Logger,
	repo *db.ConsolidatedRepo,
	consolidationType string,
	typesToProcess []string,
	typeFingerprints map[string]string,
) {
	for _, gst := range typesToProcess {
		if fp, ok := typeFingerprints[gst]; ok && fp != "" {
			if setErr := repo.SetStoredTypeFingerprint(consolidationType, gst, fp); setErr != nil {
				logger.Warnf("Failed to save type fingerprint for %s/%s: %v", consolidationType, gst, setErr)
			}
		}
	}

	// global
	if fp, err := repo.ComputeConsolidationFingerprint(consolidationType); err == nil {
		if setErr := repo.SetStoredFingerprint(consolidationType, fp); setErr != nil {
			logger.Warnf("Failed to save consolidation fingerprint for %s: %v", consolidationType, setErr)
		}
	}
}

func openConsolidatedRepo(
	ctx context.Context,
	logger *multilog.Logger,
	consolidationType string,
	force bool,
) (*db.DB, *db.ConsolidatedRepo, error) {
	dbPath := getDBPath()
	database, err := db.Open(ctx, logger, dbPath, false)
	if err != nil {
		return nil, nil, fmt.Errorf("opening database: %w", err)
	}

	repo := db.NewConsolidatedRepo(database)

	if !force {
		fingerprint, fpErr := repo.ComputeConsolidationFingerprint(consolidationType)
		if fpErr == nil && fingerprint != "" {
			stored := repo.GetStoredFingerprint(consolidationType)
			if stored == fingerprint && repo.HasConsolidatedData(consolidationType) {
				logger.Infof("Skipping %s consolidation (unchanged since last run)", consolidationType)
				database.CloseLogError(logger)
				return nil, nil, nil
			}
		}
	}

	return database, repo, nil
}
