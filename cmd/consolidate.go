package cmd

import (
	"context"
	"runtime"
	"sync"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/db"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/spf13/cobra"
)

var (
	ignoreAllowlist             bool
	includeInvalid              bool
	calculateChecksum           bool
	skipConsolidatedSummary     bool
	generateConflictsReport     bool
	emitResolvedLists           bool
	applyResolvedToConsolidated bool
)

var consolidateCmd = &cobra.Command{
	Use:   "consolidate",
	Short: "Consolidate processed files",
	Run: func(cmd *cobra.Command, args []string) {
		consolidateAllCmd.Run(cmd, args)
	},
}

var consolidateAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Consolidate all processed files",
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Infof("Consolidating all processed entries (DB-based)...")
		ctx := context.Background()

		database, consolidatedRepo, dbErr := openConsolidatedRepo(ctx, Logger, "general")
		if dbErr != nil {
			Logger.Errorf("Failed to open consolidated repo: %v", dbErr)
			return
		}
		defer database.CloseLogError(Logger)

		entriesRepo := db.NewEntriesRepo(database)

		// get generic source types
		genericSourceTypes, typesErr := entriesRepo.GetGenericSourceTypesFromDB(ctx)
		if typesErr != nil {
			Logger.Errorf("Failed to get source types from database: %v", typesErr)
			return
		}
		if len(genericSourceTypes) == 0 {
			Logger.Errorf("No source types found in database")
			return
		}

		Logger.Infof("Found %d generic source types: %v", len(genericSourceTypes), genericSourceTypes)

		// build resolution sets for conflict resolution
		allowByType, blockByType, _, _, _, _, resErr := GetResolutionSets(Logger, database)
		if resErr != nil {
			Logger.Warnf("Failed to build resolution sets (continuing without): %v", resErr)
			allowByType = make(map[string]u.StringSet)
			blockByType = make(map[string]u.StringSet)
		}

		var allConsolidatedSummaries []c.ConsolidatedSummary
		var mu sync.Mutex
		var persistMu sync.Mutex

		// Phase 1: Process allowlists
		Logger.Infof("Processing allowlists...")
		allowlistEntriesByType := make(map[string]u.StringSet)
		processAllowlistsFromDB(
			ctx,
			Logger,
			entriesRepo,
			consolidatedRepo,
			genericSourceTypes,
			blockByType,
			allowlistEntriesByType,
			&allConsolidatedSummaries,
			&persistMu,
		)

		// build allowlist filter for blocklists
		allowFilterByType := make(map[string]u.StringSet)
		for _, gst := range genericSourceTypes {
			if aset, ok := allowByType[gst]; ok && aset != nil && aset.Size() > 0 {
				allowFilterByType[gst] = aset
			} else if existing, ok := allowlistEntriesByType[gst]; ok && existing != nil {
				allowFilterByType[gst] = existing
			} else {
				allowFilterByType[gst] = u.NewStringSet([]string{})
			}
		}

		// phase 2: Process blocklists in parallel
		Logger.Infof("Processing blocklists...")
		maxWorkers := runtime.GOMAXPROCS(0)
		if AppConfig != nil && AppConfig.DNSToolkit.MaxWorkers > 0 {
			maxWorkers = AppConfig.DNSToolkit.MaxWorkers
		}
		maxWorkers = max(maxWorkers, 1)
		Logger.Infof("Using worker pool with %d worker(s) for consolidation", maxWorkers)
		workerPool := c.NewDTWorkerPool(maxWorkers)

		for i := range genericSourceTypes {
			gst := genericSourceTypes[i]
			workerPool.Submit(func() {
				Logger.Debugf("Processing blocklist for generic source type: %s", gst)

				allowlistEntries := allowFilterByType[gst]
				Logger.Debugf("Filtering %s blocklist with %d resolved allowlist entries", gst, allowlistEntries.Size())

				// get blocklist entries from DB
				blockEntries, err := entriesRepo.GetEntriesForGeneralConsolidation(
					ctx,
					gst,
					constants.ListTypeBlocklist,
				)
				if err != nil {
					Logger.Errorf("Failed to get blocklist entries for %s: %v", gst, err)
					return
				}

				if len(blockEntries) == 0 {
					return
				}

				// build entry set and source count
				entrySet := u.NewStringSet([]string{})
				sourceNames := make(map[string]struct{})
				for _, e := range blockEntries {
					entrySet.AddWithConsider(e.Entry, e.MustConsider)
					sourceNames[e.SourceName] = struct{}{}
				}

				originalCount := entrySet.Size()

				// filter with allowlist entries
				filteredSet, ignoredSet := filterEntriesWithAllowlist(entrySet, allowlistEntries)

				if filteredSet.Size() > 0 {
					if ignoredSet.Size() > 0 {
						Logger.Infof(
							"%s blocklist: %d sources, %d total → %d final (%d filtered)",
							gst, len(sourceNames), originalCount, filteredSet.Size(), ignoredSet.Size(),
						)
					} else {
						Logger.Infof(
							"%s blocklist: %d sources, %d total → %d final",
							gst, len(sourceNames), originalCount, filteredSet.Size(),
						)
					}

					summary := c.ConsolidatedSummary{
						Type:                      gst,
						FilesCount:                len(sourceNames),
						Valid:                     true,
						Count:                     filteredSet.Size(),
						OriginalCount:             originalCount,
						IgnoredEntriesCount:       ignoredSet.Size(),
						ListType:                  constants.ListTypeBlocklist,
						LastConsolidatedTimestamp: u.GetTimestamp(),
					}

					mu.Lock()
					appendSummary(&allConsolidatedSummaries, summary, IsConsolidatedSummaryValid)
					mu.Unlock()

					if err := persistConsolidatedEntries(
						ctx, Logger, consolidatedRepo, &persistMu,
						filteredSet, gst,
						constants.ListTypeBlocklist, "general", "", "", true,
					); err != nil {
						Logger.Errorf("Failed to persist blocklist entries for %s: %v", gst, err)
					}
				}

				// handle invalid entries if requested
				if includeInvalid {
					invalidEntries, invErr := entriesRepo.GetInvalidEntriesForGeneralConsolidation(
						ctx,
						gst,
						constants.ListTypeBlocklist,
					)
					if invErr != nil {
						Logger.Errorf("Failed to get invalid blocklist entries for %s: %v", gst, invErr)
						return
					}

					if len(invalidEntries) > 0 {
						invalidSet := u.NewStringSet([]string{})
						for _, e := range invalidEntries {
							invalidSet.AddWithConsider(e.Entry, e.MustConsider)
						}

						invalidFiltered, _ := filterEntriesWithAllowlist(invalidSet, allowlistEntries)
						if invalidFiltered.Size() > 0 {
							invalidSummary := c.ConsolidatedSummary{
								Type:                      gst,
								Valid:                     false,
								Count:                     invalidFiltered.Size(),
								OriginalCount:             invalidSet.Size(),
								ListType:                  constants.ListTypeBlocklist,
								LastConsolidatedTimestamp: u.GetTimestamp(),
							}

							mu.Lock()
							appendSummary(&allConsolidatedSummaries, invalidSummary, IsConsolidatedSummaryValid)
							mu.Unlock()

							if err := persistConsolidatedEntries(
								ctx, Logger, consolidatedRepo, &persistMu,
								invalidFiltered, gst,
								constants.ListTypeBlocklist, "general", "", "", false,
							); err != nil {
								Logger.Errorf("Failed to persist invalid blocklist entries for %s: %v", gst, err)
							}
						}
					}
				}
			})
		}

		Logger.Debugf("Waiting for all blocklists to finish processing...")
		workerPool.Wait()

		if generateConflictsReport {
			manager := NewConsolidationManager(Logger, database)
			if err := manager.GenerateConflictReport(); err != nil {
				Logger.Errorf("Failed to generate conflicts report: %v", err)
			}
		}

		Logger.Infof("Consolidation complete")
	},
}

// processAllowlistsFromDB processes allowlists directly from DB entries
func processAllowlistsFromDB(
	ctx context.Context,
	logger *multilog.Logger,
	entriesRepo *db.EntriesRepo,
	consolidatedRepo *db.ConsolidatedRepo,
	genericSourceTypes []string,
	resolvedBlockByType map[string]u.StringSet,
	allowlistEntriesByType map[string]u.StringSet,
	allConsolidatedSummaries *[]c.ConsolidatedSummary,
	persistMu *sync.Mutex,
) {
	for _, gst := range genericSourceTypes {
		// get all allowlist entries
		allowEntries, err := entriesRepo.GetEntriesForGeneralConsolidation(ctx, gst, constants.ListTypeAllowlist)
		if err != nil {
			logger.Errorf("Failed to get allowlist entries for %s: %v", gst, err)
			allowlistEntriesByType[gst] = u.NewStringSet([]string{})
			continue
		}

		if len(allowEntries) == 0 {
			allowlistEntriesByType[gst] = u.NewStringSet([]string{})
			continue
		}

		// build entry set, tracking must_consider and source names
		entrySet := u.NewStringSet([]string{})
		mustConsiderSet := u.NewStringSet([]string{})
		sourceNames := make(map[string]struct{})
		for _, e := range allowEntries {
			entrySet.AddWithConsider(e.Entry, e.MustConsider)
			if e.MustConsider {
				mustConsiderSet.AddWithConsider(e.Entry, true)
			}
			sourceNames[e.SourceName] = struct{}{}
		}

		// get resolved blocklist for filtering
		var resolvedBlocklist u.StringSet
		if bset, ok := resolvedBlockByType[gst]; ok && bset != nil {
			resolvedBlocklist = bset
			logger.Infof("Resolved blocklist for %s: %d entries", gst, bset.Size())
		} else {
			resolvedBlocklist = u.NewStringSet([]string{})
		}

		// filter: remove entries that are in resolved blocklist (unless must_consider)
		finalAllowlist := u.NewStringSet([]string{})
		removedByResolution := 0

		for entry := range entrySet {
			mustConsider, _ := entrySet.Get(entry)
			inResolvedBlock := resolvedBlocklist.Contains(entry)
			isMustConsider := mustConsiderSet.Contains(entry)

			if isMustConsider {
				finalAllowlist.AddWithConsider(entry, true)
			} else if !inResolvedBlock {
				finalAllowlist.AddWithConsider(entry, mustConsider)
			} else {
				removedByResolution++
			}
		}

		// ensure must-consider entries are included
		for entry := range mustConsiderSet {
			if !finalAllowlist.Contains(entry) {
				finalAllowlist.AddWithConsider(entry, true)
			}
		}

		logger.Infof(
			"Final allowlist for %s: consolidated=%d removed_by_resolution=%d must_consider=%d final=%d",
			gst, entrySet.Size(), removedByResolution, mustConsiderSet.Size(), finalAllowlist.Size(),
		)

		allowlistSummary := c.ConsolidatedSummary{
			Type:                      gst,
			FilesCount:                len(sourceNames),
			Valid:                     true,
			Count:                     finalAllowlist.Size(),
			OriginalCount:             entrySet.Size(),
			IgnoredEntriesCount:       removedByResolution,
			ListType:                  constants.ListTypeAllowlist,
			LastConsolidatedTimestamp: u.GetTimestamp(),
		}

		// persist allowlist entries
		if finalAllowlist.Size() > 0 {
			if err := persistConsolidatedEntries(
				ctx, logger, consolidatedRepo, persistMu,
				finalAllowlist, gst,
				constants.ListTypeAllowlist, "general", "", "", true,
			); err != nil {
				logger.Errorf("Failed to persist allowlist entries for %s: %v", gst, err)
			}
		}

		allowlistEntriesByType[gst] = finalAllowlist
		appendSummary(allConsolidatedSummaries, allowlistSummary, IsConsolidatedSummaryValid)

		// handle invalid allowlist entries
		if includeInvalid {
			invalidEntries, invErr := entriesRepo.GetInvalidEntriesForGeneralConsolidation(
				ctx,
				gst,
				constants.ListTypeAllowlist,
			)
			if invErr != nil {
				logger.Errorf("Failed to get invalid allowlist entries for %s: %v", gst, invErr)
				continue
			}

			if len(invalidEntries) > 0 {
				invalidSet := u.NewStringSet([]string{})
				for _, e := range invalidEntries {
					invalidSet.AddWithConsider(e.Entry, e.MustConsider)
				}

				invalidSummary := c.ConsolidatedSummary{
					Type:                      gst,
					Valid:                     false,
					Count:                     invalidSet.Size(),
					OriginalCount:             invalidSet.Size(),
					ListType:                  constants.ListTypeAllowlist,
					LastConsolidatedTimestamp: u.GetTimestamp(),
				}
				appendSummary(allConsolidatedSummaries, invalidSummary, IsConsolidatedSummaryValid)

				if err := persistConsolidatedEntries(
					ctx, logger, consolidatedRepo, persistMu,
					invalidSet, gst,
					constants.ListTypeAllowlist, "general", "", "", false,
				); err != nil {
					logger.Errorf("Failed to persist invalid allowlist entries for %s: %v", gst, err)
				}
			}
		}
	}
	logger.Debugf("Finished processing allowlists")
}

func IsConsolidatedSummaryValid(summary c.ConsolidatedSummary) bool {
	return summary.Count > 0
}

func init() {
	consolidateCmd.PersistentFlags().
		BoolVar(&ignoreAllowlist, "ignore-allowlist", false, "Ignore allowlist during consolidation where applicable")
	consolidateCmd.PersistentFlags().
		BoolVar(&includeInvalid, "include-invalid", false, "Include invalid entry(s) during consolidation")
	consolidateCmd.PersistentFlags().
		BoolVar(&calculateChecksum, "calculate-checksum", false, "Calculate checksum on the consolidated files")
	// nolint:lll
	consolidateCmd.PersistentFlags().
		BoolVar(&generateConflictsReport, "gen-conflicts", false, "Generate a conflict report, allowlist vs. blocklist")
	consolidateCmd.PersistentFlags().
		BoolVar(&emitResolvedLists, "emit-resolved-lists", false, "Emit allowlist and blocklist when resolving conflicts")
	// nolint:lll
	consolidateCmd.PersistentFlags().
		BoolVar(&applyResolvedToConsolidated, "apply-resolved-to-consolidated", true, "Apply resolved allow sets to consolidated output files (opt-in)")
	consolidateCategoriesCmd.PersistentFlags().
		BoolVar(&skipConsolidatedSummary, "skip-consolidated-summary", false, "Skip creating the consolidated summary file")
	// nolint:lll
	consolidateGroupsCmd.PersistentFlags().
		BoolVar(&skipConsolidatedSummary, "skip-consolidated-summary", false, "Skip creating the regular consolidated summary file")
	consolidateCmd.AddCommand(consolidateAllCmd)
	consolidateCmd.AddCommand(consolidateGroupsCmd)
	consolidateCmd.AddCommand(consolidateCategoriesCmd)
}
