package overlap

import (
	"github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/multilog"
)

// FileOverlapService defines the interface for overlap-related operations
type FileOverlapService interface {
	// FindOverlap processes a collection of files of the same generic source type to identify overlaps
	FindOverlap(
		logger *multilog.Logger,
		genericSourceType string,
		files []common.ProcessedFile,
	) common.OverlapSourceType

	// WriteCompactOverlapSummaries generates and writes overlap summaries from processed files
	WriteCompactOverlapSummaries(
		logger *multilog.Logger,
		processedFiles []common.ProcessedFile,
		genericSourceTypes []string,
		maxWorkers int,
	) ([]common.OverlapSummary, error)

	// GetOverlapFilename generates a filename for an overlap file based on source and target file information
	GetOverlapFilename(name1, listType1, name2, listType2, sourceType string) string

	// SaveOverlap saves overlap entries to a file
	SaveOverlap(logger *multilog.Logger, overlap []string, filename string) (string, error)
}

// IsOverlapSourceTypeValid checks if an overlap source type contains valid data
func IsOverlapSourceTypeValid(summary common.OverlapSourceType) bool {
	return len(summary.Pairs) > 0
}

// IsOverlapDetailedSummaryValid checks if an overlap detailed summary contains valid data
func IsOverlapDetailedSummaryValid(summary common.OverlapDetailedSummary) bool {
	return len(summary.SourceTypes) > 0
}
