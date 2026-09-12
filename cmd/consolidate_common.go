package cmd

import (
	"context"

	"github.com/phani-kb/dns-toolkit/internal/db"
	"github.com/phani-kb/multilog"
)

func runScopedConsolidationCommand(consolidationType, scopeLabel, startMsg, doneMsg string) {
	Logger.Infof("%s", startMsg)
	ctx := context.Background()

	database, consolidatedRepo, dbErr := openConsolidatedRepo(ctx, Logger, consolidationType, forceConsolidate)
	if dbErr != nil {
		Logger.Errorf("Failed to open consolidated repo: %v", dbErr)
		return
	}
	if database == nil {
		return
	}
	defer database.CloseLogError(Logger)

	entriesRepo := db.NewEntriesRepo(database)
	genericSourceTypes, typesErr := entriesRepo.GetGenericSourceTypes(ctx)
	if typesErr != nil {
		Logger.Errorf("Failed to get source types: %v", typesErr)
		return
	}
	if len(genericSourceTypes) == 0 {
		Logger.Infof("No entries found in database")
		return
	}

	runScopedConsolidation(ctx, Logger, consolidatedRepo, consolidationType, scopeLabel, genericSourceTypes)
	Logger.Infof("%s", doneMsg)
}

func runScopedConsolidation(
	ctx context.Context,
	logger *multilog.Logger,
	repo *db.ConsolidatedRepo,
	consolidationType, label string,
	genericSourceTypes []string,
) {
	typesToProcess, typeFingerprints := filterTypesForConsolidation(
		ctx, logger, repo, consolidationType, genericSourceTypes, forceConsolidate,
	)

	if len(typesToProcess) == 0 {
		logger.Infof("All %s source types unchanged, nothing to consolidate", consolidationType)
		saveConsolidationFingerprints(logger, repo, consolidationType, typesToProcess, typeFingerprints)
		return
	}

	logger.Infof(
		"Processing %d changed source types for %s: %v",
		len(typesToProcess),
		consolidationType,
		typesToProcess,
	)

	if err := repo.DropConsolidatedIndexes(ctx); err != nil {
		logger.Warnf("Failed to drop consolidated indexes: %v", err)
	}

	for _, gst := range typesToProcess {
		if allowRes, err := repo.ConsolidateScopedAllowlistAll(ctx, consolidationType, gst, true); err != nil {
			logger.Errorf("Failed allowlist consolidation for %s: %v", gst, err)
		} else {
			for _, r := range allowRes {
				logger.Infof("%s allowlist [%s %s]: %d entries",
					gst, label, r.ScopeValue, r.FinalCount)
			}
		}

		blockRes, err := repo.ConsolidateScopedBlocklistAll(ctx, consolidationType, gst, true)
		if err != nil {
			logger.Errorf("Failed blocklist consolidation for %s: %v", gst, err)
			continue
		}
		for _, r := range blockRes {
			logger.Infof("%s blocklist [%s %s]: %d sources, %d final",
				gst, label, r.ScopeValue, r.SourceCount, r.FinalCount)
		}

		if includeInvalid {
			if _, err := repo.ConsolidateScopedBlocklistAll(ctx, consolidationType, gst, false); err != nil {
				logger.Errorf("Failed invalid blocklist consolidation for %s: %v", gst, err)
			}
		}
	}

	if err := repo.CreateConsolidatedIndexes(ctx); err != nil {
		logger.Warnf("Failed to recreate consolidated indexes: %v", err)
	}

	saveConsolidationFingerprints(logger, repo, consolidationType, typesToProcess, typeFingerprints)
}
