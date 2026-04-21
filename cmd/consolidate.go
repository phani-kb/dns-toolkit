package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	cfg "github.com/phani-kb/dns-toolkit/internal/config"
	con "github.com/phani-kb/dns-toolkit/internal/consolidators"
	"github.com/phani-kb/dns-toolkit/internal/constants"
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
		if err := u.EnsureDirectoryExists(Logger, constants.ConsolidatedDir); err != nil {
			Logger.Errorf("Failed to create consolidated directory: %v", err)
			os.Exit(1)
		}
		if err := u.EnsureDirectoryExists(Logger, constants.SummaryDir); err != nil {
			Logger.Errorf("Failed to create summary directory: %v", err)
			os.Exit(1)
		}

		processedSummaries, genericSourceTypes, processedFiles := cfg.GetProcessedSummariesForConsolidation(
			Logger,
			SourcesConfigs,
			*AppConfig,
			"general",
		)
		if len(processedSummaries) == 0 {
			Logger.Errorf("No processed summaries found")
			return
		}

		var allConsolidatedSummaries []c.ConsolidatedSummary
		var mu sync.Mutex

		allowByType, blockByType, _, _, _, _ := GetCachedResolutionSets(Logger, processedFiles)

		allowlistEntriesByType := make(map[string]u.StringSet)
		processAllowlists(
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

				_, blocklistSummary := consolidateFilesBasedOnSTLT(
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

				if includeInvalid {
					_, invalidBlocklistSummary := consolidateFilesBasedOnSTLT(
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
				}
			})
		}

		Logger.Debugf("Waiting for all blocklists to finish processing...")
		workerPool.Wait()

		summaryFile := filepath.Join(
			constants.SummaryDir,
			constants.DefaultSummaryFiles["consolidated"],
		)
		summariesCount, err := u.SaveSummaries(
			Logger,
			allConsolidatedSummaries,
			summaryFile,
			c.ConsolidatedSummaryLessFunc,
		)
		if err != nil {
			Logger.Errorf("Error saving consolidated summaries to %s: %v", summaryFile, err)
		} else {
			if summariesCount > 0 {
				Logger.Infof("Saved consolidated summaries to %s", summaryFile)
			}

			if generateConflictsReport {
				manager := NewConsolidationManager(Logger)
				if err := manager.GenerateConflictReport(processedFiles); err != nil {
					Logger.Errorf("Failed to generate conflicts report: %v", err)
				}
			}
		}
	},
}

func processAllowlists(
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

		if finalAllowlist.Size() > 0 {
			consolidator, exists := con.Consolidators.GetConsolidator(genericSourceType, constants.ListTypeAllowlist)
			if exists {
				consolidatedPath := filepath.Join(
					constants.ConsolidatedDir,
					allowlistSummary.GetFilename(),
				)
				if err := consolidator.SaveEntries(Logger, finalAllowlist, consolidatedPath); err != nil {
					Logger.Errorf("Failed to write allowlist to %s: %v", consolidatedPath, err)
				} else {
					Logger.Infof("Wrote allowlist to %s (entries=%d)", consolidatedPath, finalAllowlist.Size())
					allowlistSummary.Filepath = consolidatedPath

					if calculateChecksum || AppConfig.DNSToolkit.FilesChecksum.Enabled {
						allowlistSummary.Checksum = u.CalculateChecksum(
							Logger,
							consolidatedPath,
							AppConfig.DNSToolkit.FilesChecksum.Algorithm,
						)
					}
				}
			}
		} else {
			Logger.Debugf("Skipping write for empty allowlist: %s", genericSourceType)
		}

		allowlistEntriesByType[genericSourceType] = finalAllowlist
		appendSummary(allConsolidatedSummaries, allowlistSummary, IsConsolidatedSummaryValid)

		if includeInvalid {
			_, invalidAllowlistSummary := consolidateFilesBasedOnSTLT(
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
		ignoredFilePath := filepath.Join(
			constants.ConsolidatedDir,
			consolidatedSummary.GetIgnoredFilename(),
		)

		// annotate ignored entries with a reason
		reason := "filtered by provided filter set"
		switch listType {
		case constants.ListTypeBlocklist:
			reason = "filtered by resolved allowlist (conflict resolution)"
		case constants.ListTypeAllowlist:
			reason = "filtered by resolved blocklist (conflict resolution)"
		}

		annotated := make([]string, 0, len(ignoredEntries))
		for entry := range ignoredEntries {
			annotated = append(annotated, fmt.Sprintf("%s # ignored: %s", entry, reason))
		}

		if err := u.WriteEntriesToFile(Logger, ignoredFilePath, annotated); err != nil {
			logger.Errorf("Error writing ignored entry(s) to file %s: %v", ignoredFilePath, err)
		} else {
			consolidatedSummary.IgnoredFilepath = ignoredFilePath
		}
	}

	if len(allEntries) <= 0 {
		logger.Infof("No entry(s) to consolidate for %s %s", listType, genericSourceType)
		return u.NewStringSet([]string{}), c.ConsolidatedSummary{}
	}

	consolidatedSummary.Filepath = filepath.Join(
		constants.ConsolidatedDir,
		consolidatedSummary.GetFilename(),
	)
	err := consolidator.SaveEntries(logger, allEntries, consolidatedSummary.Filepath)
	if err != nil {
		logger.Errorf("Error writing entry(s) to file %s: %v", consolidatedSummary.Filepath, err)
	} else {
		if calculateChecksum || AppConfig.DNSToolkit.FilesChecksum.Enabled {
			consolidatedSummary.Checksum = u.CalculateChecksum(
				logger,
				consolidatedSummary.Filepath,
				AppConfig.DNSToolkit.FilesChecksum.Algorithm,
			)
		}
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
