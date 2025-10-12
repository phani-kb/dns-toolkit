package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	con "github.com/phani-kb/dns-toolkit/internal/consolidators"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

// getFileStrings converts FileInfo slice to string slice
func getFileStrings(fileInfos []c.FileInfo) []string {
	fileStrings := make([]string, 0, len(fileInfos))
	for _, fileInfo := range fileInfos {
		fileStrings = append(fileStrings, fileInfo.GetString())
	}
	return fileStrings
}

// ConsolidationParams holds parameters for consolidation functions
type ConsolidationParams struct {
	GenericSourceType string
	ListType          string
	Identifier        string // group or category
	OutputDir         string
	IdentifierField   string // "Group" or "Category"
}

// consolidateGeneric is a generic consolidation function that can be used by both groups and categories
func consolidateGeneric(
	logger *multilog.Logger,
	params ConsolidationParams,
	entriesToIgnore u.StringSet,
	processedFiles []c.ProcessedFile,
) (u.StringSet, c.ConsolidatedSummary) {
	logger.Debugf(
		"Starting consolidation for %s %s in %s %s",
		params.ListType,
		params.GenericSourceType,
		params.IdentifierField,
		params.Identifier,
	)

	if len(processedFiles) == 0 {
		logger.Debugf(
			"No processed files found for %s %s in %s %s",
			params.ListType,
			params.GenericSourceType,
			params.IdentifierField,
			params.Identifier,
		)
		return u.NewStringSet([]string{}), c.ConsolidatedSummary{}
	}

	consolidator, exists := con.Consolidators.GetConsolidator(params.GenericSourceType, params.ListType)
	if !exists {
		logger.Warnf(
			"No consolidator found for generic source type: %s, list type: %s",
			params.GenericSourceType,
			params.ListType,
		)
		return u.NewStringSet([]string{}), c.ConsolidatedSummary{}
	}

	consolidatedEntries, fileInfos := consolidator.Consolidate(logger, processedFiles)
	consolidatedFileStrings := getFileStrings(fileInfos)

	allEntries, ignoredEntries := consolidator.FilterEntries(
		logger,
		consolidatedEntries,
		entriesToIgnore,
	)

	originalCount := calculateOriginalCount(fileInfos)

	if len(consolidatedEntries) > 0 {
		identifierStr := fmt.Sprintf("%s %s", params.IdentifierField, params.Identifier)
		if params.IdentifierField == "" {
			identifierStr = "" // For regular consolidation without identifier
		} else {
			identifierStr = " [" + identifierStr + "]"
		}

		if len(ignoredEntries) > 0 {
			logger.Infof(
				"%s %s%s: %d sources, %d total → %d final (%d filtered)",
				params.GenericSourceType,
				params.ListType,
				identifierStr,
				len(fileInfos),
				originalCount,
				len(allEntries),
				len(ignoredEntries),
			)
		} else {
			logger.Infof(
				"%s %s%s: %d sources, %d total → %d final",
				params.GenericSourceType,
				params.ListType,
				identifierStr,
				len(fileInfos),
				originalCount,
				len(allEntries),
			)
		}
	}

	if len(allEntries) <= 0 {
		logger.Infof(
			"No entry(s) to consolidate for %s %s in %s %s",
			params.ListType,
			params.GenericSourceType,
			params.IdentifierField,
			params.Identifier,
		)
		return u.NewStringSet([]string{}), c.ConsolidatedSummary{}
	}

	// Create a summary with identifier information
	consolidatedSummary := c.ConsolidatedSummary{
		Type:                      params.GenericSourceType,
		FilesCount:                len(fileInfos),
		Files:                     consolidatedFileStrings,
		Valid:                     true,
		Count:                     len(allEntries),
		OriginalCount:             calculateOriginalCount(fileInfos),
		IgnoredEntriesCount:       len(ignoredEntries),
		ListType:                  params.ListType,
		LastConsolidatedTimestamp: u.GetTimestamp(),
	}

	// Set the appropriate field based on identifier type
	switch params.IdentifierField {
	case "Group":
		consolidatedSummary.Group = params.Identifier
	case "Category":
		consolidatedSummary.Category = params.Identifier
	}

	if len(ignoredEntries) > 0 {
		filenamePrefix := fmt.Sprintf("%s_%s_%s", params.Identifier, params.GenericSourceType, params.ListType)
		ignoredFilePath := filepath.Join(
			params.OutputDir,
			filenamePrefix+"_ignored.txt",
		)

		// annotate ignored entries with a reason
		reason := "filtered by provided filter set"
		switch params.ListType {
		case constants.ListTypeBlocklist:
			reason = "filtered by consolidated allowlist"
		case constants.ListTypeAllowlist:
			reason = "filtered by local blocklist"
		}

		annotated := make([]string, 0, len(ignoredEntries))
		for entry := range ignoredEntries {
			annotated = append(annotated, fmt.Sprintf("%s # ignored: %s", entry, reason))
		}

		if err := u.WriteEntriesToFile(logger, ignoredFilePath, annotated); err != nil {
			logger.Errorf("Error writing ignored entry(s) to file %s: %v", ignoredFilePath, err)
		} else {
			consolidatedSummary.IgnoredFilepath = ignoredFilePath
		}
	}

	// Use identifier-specific filename
	filenamePrefix := fmt.Sprintf("%s_%s_%s", params.Identifier, params.GenericSourceType, params.ListType)
	consolidatedFilePath := filepath.Join(params.OutputDir, filenamePrefix+".txt")
	consolidatedSummary.Filepath = consolidatedFilePath

	resolvedBefore := 0
	if allEntries != nil {
		resolvedBefore = allEntries.Size()
	}

	resolvedAfter := resolvedBefore

	logger.Debugf(
		"Writing %s for %s_%s: original=%d final=%d",
		params.ListType,
		params.Identifier,
		params.GenericSourceType,
		calculateOriginalCount(fileInfos),
		resolvedAfter,
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

	logger.Debugf(
		"Finished consolidation for %s %s in %s %s",
		params.ListType,
		params.GenericSourceType,
		params.IdentifierField,
		params.Identifier,
	)
	return allEntries, consolidatedSummary
}

// ProcessingConfig holds configuration for processing consolidation
type ProcessingConfig struct {
	GetFilesFunc       func([]c.ProcessedFile, string) []c.ProcessedFile
	ConsolidateFunc    func(*multilog.Logger, string, string, string, u.StringSet, []c.ProcessedFile) (u.StringSet, c.ConsolidatedSummary) // nolint:lll
	AllowFilterByType  map[string]u.StringSet
	Identifier         string
	IdentifierField    string
	ProcessedFiles     []c.ProcessedFile
	GenericSourceTypes []string
}

// processIdentifierConsolidation is a generic function for processing consolidation by identifier (group or category)
func processIdentifierConsolidation(
	logger *multilog.Logger,
	config ProcessingConfig,
) map[string][]c.ConsolidatedSummary {
	consolidatedSummariesByIdentifier := make(map[string][]c.ConsolidatedSummary)
	consolidatedSummariesByIdentifier[config.Identifier] = []c.ConsolidatedSummary{}

	// filter processed files by identifier
	identifierFiles := config.GetFilesFunc(config.ProcessedFiles, config.Identifier)
	if len(identifierFiles) == 0 {
		var identifierType string
		switch config.IdentifierField {
		case "Group":
			identifierType = "size group"
		case "Category":
			identifierType = "category"
		default:
			identifierType = strings.ToLower(config.IdentifierField)
		}
		logger.Infof("No files found for %s: %s", identifierType, config.Identifier)
		return consolidatedSummariesByIdentifier
	}

	// Create a map of allowlisted entries by source type
	allowlistEntriesByType := make(map[string]u.StringSet)

	// First process allowlists for each identifier and source type
	for _, gst := range config.GenericSourceTypes {
		var allowlistFiles []c.ProcessedFile
		for _, file := range identifierFiles {
			if file.GenericSourceType == gst &&
				file.ListType == constants.ListTypeAllowlist {
				allowlistFiles = append(allowlistFiles, file)
			}
		}

		if len(allowlistFiles) > 0 {
			entries, allowlistSummary := config.ConsolidateFunc(
				logger,
				gst,
				constants.ListTypeAllowlist,
				config.Identifier,
				u.NewStringSet([]string{}),
				allowlistFiles,
			)
			allowlistEntriesByType[gst] = entries

			// Set the appropriate field based on identifier type
			switch config.IdentifierField {
			case "Group":
				allowlistSummary.Group = config.Identifier
			case "Category":
				allowlistSummary.Category = config.Identifier
			}

			consolidatedSummariesByIdentifier[config.Identifier] = append(
				consolidatedSummariesByIdentifier[config.Identifier],
				allowlistSummary,
			)
		} else {
			allowlistEntriesByType[gst] = u.NewStringSet([]string{})
		}
	}

	// Then process blocklists using the allowlists from above
	for _, gst := range config.GenericSourceTypes {
		var blocklistFiles []c.ProcessedFile
		for _, file := range identifierFiles {
			if file.GenericSourceType == gst &&
				file.ListType == constants.ListTypeBlocklist {
				blocklistFiles = append(blocklistFiles, file)
			}
		}

		if len(blocklistFiles) > 0 {
			var allowlistEntries u.StringSet
			if config.AllowFilterByType != nil {
				if aset, ok := config.AllowFilterByType[gst]; ok && aset != nil && aset.Size() > 0 {
					allowlistEntries = aset
				} else {
					allowlistEntries = u.NewStringSet([]string{})
				}
			} else {
				allowlistEntries = u.NewStringSet([]string{})
			}

			_, blocklistSummary := config.ConsolidateFunc(
				logger,
				gst,
				constants.ListTypeBlocklist,
				config.Identifier,
				allowlistEntries,
				blocklistFiles,
			)

			// Set the appropriate field based on identifier type
			switch config.IdentifierField {
			case "Group":
				blocklistSummary.Group = config.Identifier
			case "Category":
				blocklistSummary.Category = config.Identifier
			}

			consolidatedSummariesByIdentifier[config.Identifier] = append(
				consolidatedSummariesByIdentifier[config.Identifier],
				blocklistSummary,
			)
		}
	}

	return consolidatedSummariesByIdentifier
}

// processConsolidationWithTransform is a generic function that handles both identifier-based consolidation
// and transformation to source-type-based results
func processConsolidationWithTransform(
	logger *multilog.Logger,
	config ProcessingConfig,
) map[string][]c.ConsolidatedSummary {
	// Get result keyed by identifier
	resultByIdentifier := processIdentifierConsolidation(logger, config)

	// Transform to result keyed by source type
	resultBySourceType := make(map[string][]c.ConsolidatedSummary)
	for _, summaries := range resultByIdentifier {
		for _, summary := range summaries {
			// Skip empty summaries
			if summary.Type == "" {
				continue
			}
			sourceType := summary.Type
			if resultBySourceType[sourceType] == nil {
				resultBySourceType[sourceType] = []c.ConsolidatedSummary{}
			}
			resultBySourceType[sourceType] = append(resultBySourceType[sourceType], summary)
		}
	}

	return resultBySourceType
}
