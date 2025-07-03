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

	if len(consolidatedEntries) > 0 {
		logger.Infof(
			"Consolidated %s %s %d entry(s) for %s %s",
			params.ListType,
			params.GenericSourceType,
			len(consolidatedEntries),
			params.IdentifierField,
			params.Identifier,
		)
	}

	allEntries, ignoredEntries := consolidator.FilterEntries(
		logger,
		consolidatedEntries,
		entriesToIgnore,
	)

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
		logger.Infof(
			"Ignored %s %s %d entry(s) for %s %s",
			params.ListType,
			params.GenericSourceType,
			len(ignoredEntries),
			params.IdentifierField,
			params.Identifier,
		)
		filenamePrefix := fmt.Sprintf("%s_%s_%s", params.Identifier, params.GenericSourceType, params.ListType)
		ignoredFilePath := filepath.Join(
			params.OutputDir,
			filenamePrefix+"_ignored.txt",
		)
		err := consolidator.SaveEntries(logger, ignoredEntries, ignoredFilePath)
		if err != nil {
			logger.Errorf("Error writing ignored entry(s) to file %s: %v", ignoredFilePath, err)
		} else {
			consolidatedSummary.IgnoredFilepath = ignoredFilePath
		}
	}

	// Use identifier-specific filename
	filenamePrefix := fmt.Sprintf("%s_%s_%s", params.Identifier, params.GenericSourceType, params.ListType)
	consolidatedFilePath := filepath.Join(params.OutputDir, filenamePrefix+".txt")
	consolidatedSummary.Filepath = consolidatedFilePath

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
	ConsolidateFunc    func(*multilog.Logger, string, string, string, u.StringSet, []c.ProcessedFile) (u.StringSet, c.ConsolidatedSummary)
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

	var logMessage string
	switch config.IdentifierField {
	case "Group":
		logMessage = fmt.Sprintf("Processing size group: %s", config.Identifier)
	case "Category":
		logMessage = fmt.Sprintf("Processing category: %s", config.Identifier)
	default:
		logMessage = fmt.Sprintf("Processing %s: %s", config.IdentifierField, config.Identifier)
	}
	logger.Infof(logMessage)

	// Filter processed files by identifier
	identifierFiles := config.GetFilesFunc(config.ProcessedFiles, config.Identifier)
	if len(identifierFiles) == 0 {
		logger.Infof("No files found for %s: %s", strings.ToLower(config.IdentifierField), config.Identifier)
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
			allowlistEntries := allowlistEntriesByType[gst]
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
