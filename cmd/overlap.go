package cmd

import (
	"context"

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

		totalPairs := 0
		for _, gst := range genericSourceTypes {
			count := computeOverlapForGenericType(ctx, Logger, overlapRepo, gst)
			totalPairs += count
		}

		Logger.Infof("Overlap analysis complete: %d pairs computed", totalPairs)
	},
}

// computeOverlapForGenericType computes overlap for all source pairs of a given generic type,
func computeOverlapForGenericType(
	ctx context.Context,
	logger *multilog.Logger,
	overlapRepo *db.OverlapRepo,
	genericSourceType string,
) int {
	sources, err := overlapRepo.GetSourceEntryCountsAllListTypes(ctx, genericSourceType)
	if err != nil {
		logger.Errorf("Failed to get source counts for %s: %v", genericSourceType, err)
		return 0
	}

	if len(sources) < 2 {
		return 0
	}

	logger.Infof("Computing overlap for %s: %d sources", genericSourceType, len(sources))

	sourceByKey := make(map[db.SourceListKey]db.SourceEntryCount, len(sources))
	for _, s := range sources {
		sourceByKey[db.SourceListKey{SourceID: s.SourceID, ListType: s.ListType}] = s
	}

	pairs, pErr := overlapRepo.ComputeAllPairOverlapsAcrossListTypes(ctx, genericSourceType)
	if pErr != nil {
		logger.Errorf("Failed to compute overlaps for %s: %v", genericSourceType, pErr)
		return 0
	}

	if len(pairs) == 0 {
		return 0
	}

	var results []db.OverlapResultRow
	for _, p := range pairs {
		srcKey := db.SourceListKey{SourceID: p.SourceID, ListType: p.SourceListType}
		tgtKey := db.SourceListKey{SourceID: p.TargetID, ListType: p.TargetListType}
		src := sourceByKey[srcKey]
		tgt := sourceByKey[tgtKey]

		srcPercent := float64(0)
		if src.EntryCount > 0 {
			srcPercent = float64(p.OverlapCount) * 100.0 / float64(src.EntryCount)
		}
		results = append(results, db.OverlapResultRow{
			SourceName:        src.SourceName,
			TargetName:        tgt.SourceName,
			GenericSourceType: genericSourceType,
			SourceListType:    p.SourceListType,
			TargetListType:    p.TargetListType,
			OverlapCount:      p.OverlapCount,
			SourceCount:       src.EntryCount,
			TargetCount:       tgt.EntryCount,
			OverlapPercent:    srcPercent,
		})

		tgtPercent := float64(0)
		if tgt.EntryCount > 0 {
			tgtPercent = float64(p.OverlapCount) * 100.0 / float64(tgt.EntryCount)
		}
		results = append(results, db.OverlapResultRow{
			SourceName:        tgt.SourceName,
			TargetName:        src.SourceName,
			GenericSourceType: genericSourceType,
			SourceListType:    p.TargetListType,
			TargetListType:    p.SourceListType,
			OverlapCount:      p.OverlapCount,
			SourceCount:       tgt.EntryCount,
			TargetCount:       src.EntryCount,
			OverlapPercent:    tgtPercent,
		})
	}

	pairCount := len(pairs)
	if len(results) > 0 {
		if err := overlapRepo.PersistOverlapResults(ctx, results); err != nil {
			logger.Errorf("Failed to persist overlap results for %s: %v", genericSourceType, err)
		} else {
			logger.Infof("%s: %d pairs with overlap persisted", genericSourceType, pairCount)
		}
	}

	return pairCount
}
