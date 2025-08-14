package cmd

import (
	"os"
	"path/filepath"
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
	ignoreAllowlist         bool
	includeInvalid          bool
	calculateChecksum       bool
	skipConsolidatedSummary bool
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

		_, genericSourceTypes, processedFiles := cfg.GetProcessedSummariesForConsolidation(
			Logger,
			SourcesConfigs,
			*AppConfig,
			"general",
		)
		var allConsolidatedSummaries []c.ConsolidatedSummary
		allowlistEntriesByType := make(map[string]u.StringSet)
		var mu sync.Mutex

		// First phase: Process all allowlists synchronously
		if !ignoreAllowlist {
			processAllowlists(
				genericSourceTypes,
				processedFiles,
				allowlistEntriesByType,
				&allConsolidatedSummaries,
			)
		}

		// Second phase: Process all blocklists in parallel, now that we have all allowlist entries
		Logger.Infof("Processing blocklists...")
		var wg sync.WaitGroup
		blocklistTypes := make([]string, len(genericSourceTypes))
		copy(
			blocklistTypes,
			genericSourceTypes,
		) // Make a copy to prevent loop variable capture issues

		for i := range blocklistTypes {
			genericSourceType := blocklistTypes[i] // Local variable for this iteration
			wg.Add(1)
			go func(gst string) {
				defer wg.Done()
				Logger.Debugf("Processing blocklist for generic source type: %s", gst)

				allowlistEntries := allowlistEntriesByType[gst]

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
			}(genericSourceType)
		}

		Logger.Debugf("Waiting for all blocklists to finish processing...")
		wg.Wait()

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
			strings.HasPrefix(file.Name, "Local") {
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
		err := consolidator.SaveEntries(logger, ignoredEntries, ignoredFilePath)
		if err != nil {
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

func getFileStrings(fileInfos []c.FileInfo) []string {
	fileStrings := make([]string, 0)
	for _, fileInfo := range fileInfos {
		fileStrings = append(fileStrings, fileInfo.GetString())
	}
	return fileStrings
}

func init() {
	consolidateCmd.PersistentFlags().
		BoolVar(&ignoreAllowlist, "ignore-allowlist", false, "Ignore allowlist during consolidation where applicable")
	consolidateCmd.PersistentFlags().
		BoolVar(&includeInvalid, "include-invalid", false, "Include invalid entry(s) during consolidation")
	consolidateCmd.PersistentFlags().
		BoolVar(&calculateChecksum, "calculate-checksum", false, "Calculate checksum on the consolidated files")
	consolidateCategoriesCmd.PersistentFlags().
		BoolVar(&skipConsolidatedSummary, "skip-consolidated-summary", false, "Skip creating the consolidated summary file")
	// nolint:lll
	consolidateGroupsCmd.PersistentFlags().
		BoolVar(&skipConsolidatedSummary, "skip-consolidated-summary", false, "Skip creating the regular consolidated summary file")
	consolidateCmd.AddCommand(consolidateAllCmd)
	consolidateCmd.AddCommand(consolidateGroupsCmd)
	consolidateCmd.AddCommand(consolidateCategoriesCmd)
}
