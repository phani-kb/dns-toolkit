package cmd

import (
	"context"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/db"
	"github.com/spf13/cobra"
)

var (
	minSources int
	maxEntries int
)

var topEntriesCmd = &cobra.Command{
	Use:   "top",
	Short: "Find top entry(s) in each generic source type",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		if maxEntries < 1 {
			Logger.Errorf("--max-entries must be at least 1, got %d", maxEntries)
			return
		}

		minSourcesRange := []int{minSources}
		if minSources == 0 {
			minSourcesRange = constants.DefaultMinSourcesRange
		}

		database := openDB(ctx)
		defer database.CloseLogError(Logger)

		topRepo := db.NewTopEntriesRepo(database)
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

		listTypes := []string{constants.ListTypeBlocklist, constants.ListTypeAllowlist} // TODO: get list from db

		Logger.Infof("Processing top entries: min_sources_range=%v, max_entries=%d", minSourcesRange, maxEntries)

		totalPersisted := 0
		for _, gst := range genericSourceTypes {
			for _, listType := range listTypes {
				for _, minSrc := range minSourcesRange {
					entries, qErr := topRepo.GetTopEntries(ctx, gst, listType, minSrc, maxEntries)
					if qErr != nil {
						Logger.Errorf("Failed to get top entries for %s/%s (min=%d): %v", gst, listType, minSrc, qErr)
						continue
					}

					if len(entries) == 0 {
						Logger.Debugf("No top entries for %s/%s with min_sources=%d", gst, listType, minSrc)
						continue
					}

					if pErr := topRepo.PersistTopEntries(ctx, gst, listType, minSrc, entries); pErr != nil {
						Logger.Errorf(
							"Failed to persist top entries for %s/%s (min=%d): %v",
							gst,
							listType,
							minSrc,
							pErr,
						)
						continue
					}

					totalPersisted += len(entries)
					Logger.Infof("%s %s: %d entries appearing in %d+ sources",
						gst, listType, len(entries), minSrc)
				}
			}
		}

		Logger.Infof("Top entries complete: %d entries persisted to database", totalPersisted)
	},
}

func init() {
	topEntriesCmd.Flags().IntVarP(&minSources, "min-sources", "m", 0, "Minimum sources (0 = default range 3-12)")
	topEntriesCmd.Flags().IntVarP(&maxEntries, "max-entries", "x", int(^uint(0)>>1), "Max entries to output")
}
