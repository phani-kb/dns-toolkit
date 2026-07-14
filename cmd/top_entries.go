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

		if minSources < 1 {
			Logger.Errorf("--min-sources must be at least 1, got %d", minSources)
			return
		}
		if maxEntries < 1 {
			Logger.Errorf("--max-entries must be at least 1, got %d", maxEntries)
			return
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

		listTypes := []string{constants.ListTypeBlocklist, constants.ListTypeAllowlist}

		Logger.Infof("Processing top entries: min_sources=%d, max_entries=%d", minSources, maxEntries)

		totalPersisted := 0
		for _, gst := range genericSourceTypes {
			for _, listType := range listTypes {
				entries, qErr := topRepo.GetTopEntries(ctx, gst, listType, minSources, maxEntries)
				if qErr != nil {
					Logger.Errorf("Failed to get top entries for %s/%s: %v", gst, listType, qErr)
					continue
				}

				if len(entries) == 0 {
					Logger.Debugf("No top entries for %s/%s with min_sources=%d", gst, listType, minSources)
					continue
				}

				if pErr := topRepo.PersistTopEntries(ctx, gst, listType, minSources, entries); pErr != nil {
					Logger.Errorf("Failed to persist top entries for %s/%s: %v", gst, listType, pErr)
					continue
				}

				totalPersisted += len(entries)
				Logger.Infof("%s %s: %d entries appearing in %d+ sources",
					gst, listType, len(entries), minSources)
			}
		}

		Logger.Infof("Top entries complete: %d entries persisted to database", totalPersisted)
	},
}

func init() {
	topEntriesCmd.Flags().IntVarP(&minSources, "min-sources", "m", 3, "Minimum sources (default 3)")
	topEntriesCmd.Flags().IntVarP(&maxEntries, "max-entries", "x", int(^uint(0)>>1), "Max entries to output")
}
