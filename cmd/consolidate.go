package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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
		allowlistEntriesByType := make(map[string]u.StringSet)

		// first phase: process allowlists which will populate allowlistEntriesByType
		processAllowlists(genericSourceTypes, processedFiles, allowlistEntriesByType, &allConsolidatedSummaries)

		// build resolution sets (counts-based) to optionally use resolved allows for filtering
		allowByType, _, _, _, _, _ := GetCachedResolutionSets(Logger, processedFiles)

		// prefer resolved allow sets when present
		allowFilterByType := make(map[string]u.StringSet)
		for _, gst := range genericSourceTypes {
			if aset, ok := allowByType[gst]; ok && aset != nil && aset.Size() > 0 {
				allowFilterByType[gst] = aset
				if existing, ok := allowlistEntriesByType[gst]; ok && existing != nil {
					Logger.Infof(
						"Using resolved allow set for filtering %s: resolved=%d consolidated=%d",
						gst,
						aset.Size(),
						existing.Size(),
					)
				} else {
					Logger.Infof("Using resolved allow set for filtering %s: resolved=%d consolidated=0", gst, aset.Size())
				}
			} else if existing, ok := allowlistEntriesByType[gst]; ok && existing != nil {
				allowFilterByType[gst] = existing
			} else {
				allowFilterByType[gst] = u.NewStringSet([]string{})
			}
		}

		// apply resolved allow sets to consolidated outputs
		if applyResolvedToConsolidated {
			Logger.Infof("Applying resolved allow sets to consolidated outputs (opt-in)")
			for _, gst := range genericSourceTypes {
				if aset, ok := allowByType[gst]; ok && aset != nil && aset.Size() > 0 {
					consolidatedPath := filepath.Join(
						constants.ConsolidatedDir,
						gst+"_"+constants.ListTypeAllowlist+".txt",
					)
					consolidator, exists := con.Consolidators.GetConsolidator(gst, constants.ListTypeAllowlist)
					if exists {
						if err := consolidator.SaveEntries(Logger, aset, consolidatedPath); err != nil {
							Logger.Errorf("Failed to write resolved allow entries to %s: %v", consolidatedPath, err)
						} else {
							Logger.Infof("Overwrote consolidated allowlist: %s (entries=%d)", consolidatedPath, aset.Size())

							// update any in-memory consolidated summary toreflect the new entry count.
							for i := range allConsolidatedSummaries {
								cs := &allConsolidatedSummaries[i]
								if cs.Type == gst && cs.ListType == constants.ListTypeAllowlist && cs.Valid {
									// preserve original count before overwriting
									cs.OriginalCount = cs.Count
									cs.Count = aset.Size()
									cs.Filepath = consolidatedPath
								}
							}
						}
					} else {
						entries := aset.ToSliceSorted()
						if err := u.WriteEntriesToFile(Logger, consolidatedPath, entries); err != nil {
							Logger.Errorf("Failed to write resolved allow entries to %s: %v", consolidatedPath, err)
						} else {
							Logger.Infof("Wrote resolved allowlist to %s (entries=%d)", consolidatedPath, aset.Size())

							// update in-memory consolidated summary as above
							for i := range allConsolidatedSummaries {
								cs := &allConsolidatedSummaries[i]
								if cs.Type == gst && cs.ListType == constants.ListTypeAllowlist && cs.Valid {
									cs.OriginalCount = cs.Count
									cs.Count = aset.Size()
									cs.Filepath = consolidatedPath
								}
							}
						}
					}
				}
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
	allowlistEntriesByType map[string]u.StringSet,
	allConsolidatedSummaries *[]c.ConsolidatedSummary,
) {
	Logger.Infof("Processing allowlists...")
	for _, genericSourceType := range genericSourceTypes {
		// Get local blocklist entries for this source type to use as filter
		localBlocklistEntries := getLocalBlocklistEntries(genericSourceType, processedFiles)

		entries, allowlistSummary := consolidateFilesBasedOnSTLT(
			Logger,
			genericSourceType,
			constants.ListTypeAllowlist,
			true,
			localBlocklistEntries, // Use local blocklist as filter
			processedFiles,
		)
		allowlistEntriesByType[genericSourceType] = entries
		appendSummary(allConsolidatedSummaries, allowlistSummary, IsConsolidatedSummaryValid)

		Logger.Debugf("Valid Allowlisted entry(s) count for %s: %d", genericSourceType, len(entries))
		if len(localBlocklistEntries) > 0 {
			Logger.Debugf("Local blocklist entries filtered for %s: %d", genericSourceType, len(localBlocklistEntries))
		}

		if includeInvalid {
			_, invalidAllowlistSummary := consolidateFilesBasedOnSTLT(
				Logger,
				genericSourceType,
				constants.ListTypeAllowlist,
				false,
				localBlocklistEntries,
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

func getLocalBlocklistEntries(genericSourceType string, processedFiles []c.ProcessedFile) u.StringSet {
	localBlocklistFiles := make([]c.ProcessedFile, 0)
	for _, file := range processedFiles {
		if file.GenericSourceType == genericSourceType &&
			file.ListType == constants.ListTypeBlocklist &&
			file.Valid &&
			strings.HasPrefix(strings.ToLower(file.Name), "local") {
			localBlocklistFiles = append(localBlocklistFiles, file)
		}
	}

	if len(localBlocklistFiles) == 0 {
		return u.NewStringSet([]string{})
	}

	Logger.Debugf("Found %d local blocklist files for %s", len(localBlocklistFiles), genericSourceType)

	consolidator, exists := con.Consolidators.GetConsolidator(genericSourceType, constants.ListTypeBlocklist)
	if !exists {
		Logger.Warnf("No consolidator found for generic source type: %s, list type: %s",
			genericSourceType, constants.ListTypeBlocklist)
		return u.NewStringSet([]string{})
	}

	// Consolidate local blocklist entries
	consolidatedEntries, _ := consolidator.Consolidate(Logger, localBlocklistFiles)
	Logger.Debugf("Consolidated %d local blocklist entries for %s", len(consolidatedEntries), genericSourceType)

	return consolidatedEntries
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
			reason = "filtered by consolidated allowlist"
		case constants.ListTypeAllowlist:
			reason = "filtered by local blocklist"
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
