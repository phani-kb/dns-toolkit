package cmd

import (
	"context"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/db"
	"github.com/phani-kb/multilog"
	"github.com/spf13/cobra"
)

var overlapCmd = &cobra.Command{
	Use:   "overlap",
	Short: "Find overlap between sources",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		database := openDB(ctx)
		defer database.CloseLogError(Logger)

		overlapRepo := db.NewOverlapRepo(database)
		entriesRepo := db.NewEntriesRepo(database)

		if err := overlapRepo.ClearOverlapResults(); err != nil {
			Logger.Errorf("Failed to clear overlap results: %v", err)
			return
		}

		genericSourceTypes, err := entriesRepo.GetGenericSourceTypes(ctx)
		if err != nil {
			Logger.Errorf("Failed to get source types: %v", err)
			return
		}

		if len(genericSourceTypes) == 0 {
			Logger.Infof("No entries found")
			return
		}

		listTypes := []string{constants.ListTypeBlocklist, constants.ListTypeAllowlist}
		totalPairs := 0

		for _, gst := range genericSourceTypes {
			for _, listType := range listTypes {
				count := computeOverlapForType(ctx, Logger, overlapRepo, gst, listType)
				totalPairs += count
			}
		}

		Logger.Infof("Overlap analysis complete: %d pairs computed", totalPairs)
	},
}

// computeOverlapForType computes overlap for all pairs of sources of a given type/list.
func computeOverlapForType(
	ctx context.Context,
	logger *multilog.Logger,
	overlapRepo *db.OverlapRepo,
	genericSourceType string,
	listType string,
) int {
	sources, err := overlapRepo.GetSourceEntryCounts(ctx, genericSourceType, listType)
	if err != nil {
		logger.Errorf("Failed to get source counts for %s/%s: %v", genericSourceType, listType, err)
		return 0
	}

	if len(sources) < 2 {
		return 0
	}

	logger.Infof("Computing overlap for %s %s: %d sources", genericSourceType, listType, len(sources))

	var results []db.OverlapResultRow
	pairCount := 0

	for i := range sources {
		for j := i + 1; j < len(sources); j++ {
			src := sources[i]
			tgt := sources[j]

			overlapCount, cErr := overlapRepo.ComputePairOverlap(
				ctx, src.SourceID, tgt.SourceID, genericSourceType,
			)
			if cErr != nil {
				logger.Errorf("Failed to compute overlap %s↔%s: %v", src.SourceName, tgt.SourceName, cErr)
				continue
			}

			if overlapCount == 0 {
				continue
			}

			// source and target
			srcPercent := float64(0)
			if src.EntryCount > 0 {
				srcPercent = float64(overlapCount) * 100.0 / float64(src.EntryCount)
			}
			results = append(results, db.OverlapResultRow{
				SourceName:        src.SourceName,
				TargetName:        tgt.SourceName,
				GenericSourceType: genericSourceType,
				SourceListType:    listType,
				TargetListType:    listType,
				OverlapCount:      overlapCount,
				SourceCount:       src.EntryCount,
				TargetCount:       tgt.EntryCount,
				OverlapPercent:    srcPercent,
			})

			// target and source
			tgtPercent := float64(0)
			if tgt.EntryCount > 0 {
				tgtPercent = float64(overlapCount) * 100.0 / float64(tgt.EntryCount)
			}
			results = append(results, db.OverlapResultRow{
				SourceName:        tgt.SourceName,
				TargetName:        src.SourceName,
				GenericSourceType: genericSourceType,
				SourceListType:    listType,
				TargetListType:    listType,
				OverlapCount:      overlapCount,
				SourceCount:       tgt.EntryCount,
				TargetCount:       src.EntryCount,
				OverlapPercent:    tgtPercent,
			})

			pairCount++
		}
	}

	if len(results) > 0 {
		if err := overlapRepo.PersistOverlapResults(ctx, results); err != nil {
			logger.Errorf("Failed to persist overlap results for %s/%s: %v", genericSourceType, listType, err)
		} else {
			logger.Infof("%s %s: %d pairs with overlap persisted", genericSourceType, listType, pairCount)
		}
	}

	return pairCount
}
