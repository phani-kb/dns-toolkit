package cmd

import (
	"context"
	"runtime"
	"sync"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	con "github.com/phani-kb/dns-toolkit/internal/consolidators"
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
		ctx := context.Background()

		processedSummaries, genericSourceTypes, processedFiles, loadErr := loadProcessedInputsForConsolidation(
			ctx,
			Logger,
			"general",
		)
		if loadErr != nil {
			Logger.Errorf("Failed to load processed summaries from database: %v", loadErr)
			return
		}
		if len(processedSummaries) == 0 {
			Logger.Errorf("No processed summaries found")
			return
		}

		database, consolidatedRepo, dbErr := openConsolidatedRepo(ctx, Logger, "general")
		if dbErr != nil {
			Logger.Errorf("Failed to open consolidated repo: %v", dbErr)
			return
		}
		defer database.CloseLogError(Logger)

		var allConsolidatedSummaries []c.ConsolidatedSummary
		var mu sync.Mutex
		var persistMu sync.Mutex

		allowByType, blockByType, _, _, _, _, resErr := GetCachedResolutionSets(Logger, database, processedFiles)
		if resErr != nil {
			Logger.Errorf("Failed to build resolution sets: %v", resErr)
			return
		}

		allowlistEntriesByType := make(map[string]u.StringSet)
		processAllowlists(
			ctx,
			consolidatedRepo,
			genericSourceTypes,
			processedFiles,
			allowByType,
			blockByType,
			allowlistEntriesByType,
			&allConsolidatedSummaries,
		)

		// use resolved allow sets for filtering blocklists
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

		// Second phase: Process all blocklists in parallel, now that we have all allowlist entries
		Logger.Infof("Processing blocklists...")
		blocklistTypes := make([]string, len(genericSourceTypes))
		copy(
			blocklistTypes,
			genericSourceTypes,
		) // Make a copy to prevent loop variable capture issues

		maxWorkers := runtime.GOMAXPROCS(0)
		if AppConfig != nil && AppConfig.DNSToolkit.MaxWorkers > 0 {
			maxWorkers = AppConfig.DNSToolkit.MaxWorkers
		}
		maxWorkers = max(maxWorkers, 1)
		Logger.Infof("Using worker pool with %d worker(s) for consolidation", maxWorkers)
		workerPool := c.NewDTWorkerPool(maxWorkers)

		for i := range blocklistTypes {
			genericSourceType := blocklistTypes[i] // Local variable for this iteration
			workerPool.Submit(func() {
				gst := genericSourceType
				Logger.Debugf("Processing blocklist for generic source type: %s", gst)

				allowlistEntries := allowFilterByType[gst]
				Logger.Debugf("Filtering %s blocklist with %d resolved allowlist entries", gst, allowlistEntries.Size())

				blockEntries, blocklistSummary := consolidateFilesBasedOnSTLT(
					Logger,
					gst,
					constants.ListTypeBlocklist,
					true,
					allowlistEntries,
					processedFiles,
				)
				mu.Lock()
				appendSummary(
					&allConsolidatedSummaries,
					blocklistSummary,
					IsConsolidatedSummaryValid,
				)
				mu.Unlock()

				if blockEntries.Size() > 0 {
					if err := persistConsolidatedEntries(
						ctx, Logger, consolidatedRepo, &persistMu,
						blockEntries, gst,
						constants.ListTypeBlocklist, "general", "", "", true,
					); err != nil {
						Logger.Errorf("Failed to persist blocklist entries for %s: %v", gst, err)
					}
				}

				if includeInvalid {
					invalidEntries, invalidBlocklistSummary := consolidateFilesBasedOnSTLT(
						Logger,
						gst,
						constants.ListTypeBlocklist,
						false,
						allowlistEntries,
						processedFiles,
					)
					mu.Lock()
					appendSummary(
						&allConsolidatedSummaries,
						invalidBlocklistSummary,
						IsConsolidatedSummaryValid,
					)
					mu.Unlock()

					if invalidEntries.Size() > 0 {
						if err := persistConsolidatedEntries(
							ctx, Logger, consolidatedRepo, &persistMu,
							invalidEntries, gst,
							constants.ListTypeBlocklist, "general", "", "", false,
						); err != nil {
							Logger.Errorf("Failed to persist invalid blocklist entries for %s: %v", gst, err)
						}
					}
				}
			})
		}

		Logger.Debugf("Waiting for all blocklists to finish processing...")
		workerPool.Wait()

		if generateConflictsReport {
			manager := NewConsolidationManager(Logger, database)
			if err := manager.GenerateConflictReport(processedFiles); err != nil {
				Logger.Errorf("Failed to generate conflicts report: %v", err)
			}
		}

		Logger.Infof("Consolidation complete")
	},
}

func processAllowlists(
	ctx context.Context,
	consolidatedRepo *db.ConsolidatedRepo,
	genericSourceTypes []string,
	processedFiles []c.ProcessedFile,
	_ map[string]u.StringSet,
	resolvedBlockByType map[string]u.StringSet,
	allowlistEntriesByType map[string]u.StringSet,
	allConsolidatedSummaries *[]c.ConsolidatedSummary,
) {
	Logger.Infof("Processing allowlists...")
	for _, genericSourceType := range genericSourceTypes {
		// consolidate all allowlist source files (no filtering)
		entries, allowlistSummary := consolidateFilesBasedOnSTLT(
			Logger,
			genericSourceType,
			constants.ListTypeAllowlist,
			true,
			u.NewStringSet([]string{}), // no filtering during initial consolidation
			processedFiles,
		)

		var resolvedBlocklist u.StringSet
		if bset, ok := resolvedBlockByType[genericSourceType]; ok && bset != nil {
			resolvedBlocklist = bset
			Logger.Infof("Resolved blocklist for %s: %d entries", genericSourceType, bset.Size())
		} else {
			resolvedBlocklist = u.NewStringSet([]string{})
		}

		mustConsiderSet := u.NewStringSet([]string{})
		for _, pf := range processedFiles {
			if pf.GenericSourceType == genericSourceType &&
				pf.ListType == constants.ListTypeAllowlist &&
				pf.MustConsider {
				fileEntries, _, err := u.ReadEntriesFromFile(Logger, pf.Filepath)
				if err != nil {
					Logger.Warnf(
						"Unable to read must-consider source file %s: %v",
						pf.Filepath,
						err,
					)
					continue
				}
				mustConsiderSet.AddAll(fileEntries, true)
			}
		}

		finalAllowlist := u.NewStringSet([]string{})
		removedByResolution := 0

		for entry := range entries {
			mustConsider, _ := entries.Get(entry)
			inResolvedBlock := resolvedBlocklist.Contains(entry)
			isMustConsider := mustConsiderSet.Contains(entry)

			if isMustConsider {
				// always include must-consider entries
				finalAllowlist.AddWithConsider(entry, true)
			} else if !inResolvedBlock {
				// include if not in resolved blocklist
				finalAllowlist.AddWithConsider(entry, mustConsider)
			} else {
				// resolved as a blocklist entry, exclude it
				removedByResolution++
			}
		}

		// add must-consider entries
		for entry := range mustConsiderSet {
			if !finalAllowlist.Contains(entry) {
				finalAllowlist.AddWithConsider(entry, true)
			}
		}

		Logger.Infof(
			"Final allowlist for %s: consolidated=%d removed_by_resolution=%d must_consider=%d final=%d",
			genericSourceType,
			entries.Size(),
			removedByResolution,
			mustConsiderSet.Size(),
			finalAllowlist.Size(),
		)

		// update summary with final counts
		allowlistSummary.Count = finalAllowlist.Size()
		allowlistSummary.IgnoredEntriesCount = removedByResolution

		// persist allowlist entries to database
		if finalAllowlist.Size() > 0 {
			if err := persistConsolidatedEntries(
				ctx, Logger, consolidatedRepo, nil,
				finalAllowlist, genericSourceType,
				constants.ListTypeAllowlist, "general", "", "", true,
			); err != nil {
				Logger.Errorf("Failed to persist allowlist entries for %s: %v", genericSourceType, err)
			}
		} else {
			Logger.Debugf("Skipping persist for empty allowlist: %s", genericSourceType)
		}

		allowlistEntriesByType[genericSourceType] = finalAllowlist
		appendSummary(allConsolidatedSummaries, allowlistSummary, IsConsolidatedSummaryValid)

		if includeInvalid {
			invalidEntries, invalidAllowlistSummary := consolidateFilesBasedOnSTLT(
				Logger,
				genericSourceType,
				constants.ListTypeAllowlist,
				false,
				u.NewStringSet([]string{}),
				processedFiles,
			)
			appendSummary(
				allConsolidatedSummaries,
				invalidAllowlistSummary,
				IsConsolidatedSummaryValid,
			)
			if invalidEntries.Size() > 0 {
				if err := persistConsolidatedEntries(
					ctx, Logger, consolidatedRepo, nil,
					invalidEntries, genericSourceType,
					constants.ListTypeAllowlist, "general", "", "", false,
				); err != nil {
					Logger.Errorf("Failed to persist invalid allowlist entries for %s: %v", genericSourceType, err)
				}
			}
		}
	}
	Logger.Debugf("Finished processing allowlists")
}

// calculateOriginalCount sums the Count from all file infos
func calculateOriginalCount(fileInfos []c.FileInfo) int {
	total := 0
	for _, fi := range fileInfos {
		total += fi.Count
	}
	return total
}

func consolidateFilesBasedOnSTLT(
	logger *multilog.Logger,
	genericSourceType, listType string,
	valid bool,
	entriesToIgnore u.StringSet,
	allProcessedFiles []c.ProcessedFile,
) (u.StringSet, c.ConsolidatedSummary) {
	logger.Debugf("Starting consolidation for %s %s", listType, genericSourceType)

	processedFiles := make([]c.ProcessedFile, 0)
	for _, processedFile := range allProcessedFiles {
		if valid == processedFile.Valid {
			processedFiles = append(processedFiles, processedFile)
		}
	}
	logger.Debugf("Processed files count: %d", len(processedFiles))
	consolidator, exists := con.Consolidators.GetConsolidator(genericSourceType, listType)
	if !exists {
		Logger.Warnf(
			"No consolidator found for generic source type: %s, list type: %s",
			genericSourceType,
			listType,
		)
		return u.NewStringSet([]string{}), c.ConsolidatedSummary{}
	}
	consolidatedEntries, fileInfos := consolidator.Consolidate(Logger, processedFiles)
	consolidatedFileStrings := getFileStrings(fileInfos)
	if len(consolidatedEntries) > 0 {
		logger.Infof(
			"Consolidated %s %s %d entry(s)",
			listType,
			genericSourceType,
			len(consolidatedEntries),
		)
	}
	allEntries, ignoredEntries := consolidator.FilterEntries(
		logger,
		consolidatedEntries,
		entriesToIgnore,
	)

	consolidatedSummary := c.ConsolidatedSummary{
		Type:                      genericSourceType,
		FilesCount:                len(fileInfos),
		Files:                     consolidatedFileStrings,
		Valid:                     valid,
		Count:                     len(allEntries),
		OriginalCount:             calculateOriginalCount(fileInfos),
		IgnoredEntriesCount:       len(ignoredEntries),
		ListType:                  listType,
		LastConsolidatedTimestamp: u.GetTimestamp(),
	}

	if len(ignoredEntries) > 0 {
		logger.Infof("Ignored %s %s %d entry(s)", listType, genericSourceType, len(ignoredEntries))
	}

	if len(allEntries) <= 0 {
		logger.Infof("No entry(s) to consolidate for %s %s", listType, genericSourceType)
		return u.NewStringSet([]string{}), c.ConsolidatedSummary{}
	}

	logger.Debugf("Finished consolidation for %s %s", listType, genericSourceType)
	return allEntries, consolidatedSummary
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
