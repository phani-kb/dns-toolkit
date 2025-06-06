package processors

import (
	"github.com/phani-kb/multilog"
)

// Processor defines the interface that all processor types should implement.
// Each processor must be able to parse content and extract valid and invalid entries,
// as well as report its source type.
type Processor interface {
	// Process parses the content and returns valid and invalid entries
	Process(logger *multilog.Logger, content string) ([]string, []string)
	// GetSourceType returns the source type this processor handles
	GetSourceType() string
	// GetListType returns the list type this processor handles
	GetListType() string
}
