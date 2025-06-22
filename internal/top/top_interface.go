package top

import (
	"github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

// EntriesService defines the interface for top entries operations
type EntriesService interface {
	// FindTopEntries processes files to find the top entries that appear in at least minSources
	// sources for a specific generic source type and list type
	FindTopEntries(
		logger *multilog.Logger,
		genericSourceType string,
		listType string,
		allProcessedFiles []common.ProcessedFile,
		minSources int,
		maxEntries int,
		stringPool *utils.DTEntryPool,
	) (common.TopSummary, error)

	// ProcessTopEntries processes all processed files to generate top entries for different
	// combinations of generic source types, list types and minimum sources thresholds
	ProcessTopEntries(
		logger *multilog.Logger,
		genericSourceTypes []string,
		processedFiles []common.ProcessedFile,
		minSourcesValues []int,
		maxEntries int,
		maxWorkers int,
	) ([]common.TopSummary, error)

	// FilterTopSummaries removes empty or invalid top summaries
	FilterTopSummaries(topSummaries []common.TopSummary) []common.TopSummary

	// GetTopNEntries returns the top N entries that appear in at least minSources sources
	GetTopNEntries(
		entrySources map[string]map[uint16]struct{},
		minSources int,
		maxEntries int,
	) []common.EntryCountPair

	// SaveTopEntries saves the top entries to a file and returns the filepath
	SaveTopEntries(
		logger *multilog.Logger,
		genericSourceType string,
		listType string,
		minSources int,
		entries []common.EntryCountPair,
	) (string, error)

	// SaveTopSummaries saves the top summaries to the summary file
	SaveTopSummaries(
		logger *multilog.Logger,
		topSummaries []common.TopSummary,
	) (int, error)
}
