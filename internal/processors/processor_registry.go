package processors

import (
	"fmt"
	"sync"

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

// createRegistryKey creates a composite key from sourceType and listType
func createRegistryKey(sourceType, listType string) string {
	return fmt.Sprintf("%s:%s", sourceType, listType)
}

// ProcessorRegistry is a thread-safe registry for processors.
type ProcessorRegistry struct {
	// CustomProcessorRegistry is a map of source types and list types to processors
	// It allows the application to dynamically register and retrieve processors
	// for different source and list type combinations.
	customProcessorRegistry map[string]Processor
	registryMutex           sync.RWMutex
}

// NewProcessorRegistry provides a thread-safe registry for processors.
func NewProcessorRegistry() *ProcessorRegistry {
	return &ProcessorRegistry{
		customProcessorRegistry: make(map[string]Processor),
	}
}

// RegisterProcessor registers a processor for a source type and list type.
// If a processor for the given source type and list type already exists, it will be replaced.
//
// Parameters:
//   - sourceType: The source type identifier for this processor
//   - listType: The list type identifier for this processor
//   - processor: The processor implementation to register
func (pr *ProcessorRegistry) RegisterProcessor(sourceType, listType string, processor Processor) {
	pr.registryMutex.Lock()
	defer pr.registryMutex.Unlock()
	key := createRegistryKey(sourceType, listType)
	pr.customProcessorRegistry[key] = processor
}

// GetProcessor retrieves a processor for a source type and list type.
//
// Parameters:
//   - sourceType: The source type to look up
//   - listType: The list type to look up
//
// Returns:
//   - The registered processor for the source type and list type
//   - A boolean indicating whether a processor was found
func (pr *ProcessorRegistry) GetProcessor(sourceType, listType string) (Processor, bool) {
	pr.registryMutex.RLock()
	defer pr.registryMutex.RUnlock()
	key := createRegistryKey(sourceType, listType)
	processor, ok := pr.customProcessorRegistry[key]
	if !ok {
		return nil, false
	}
	return processor, ok
}

// ListProcessors returns a copy of all registered processors.
// The returned map is a clone of the registry, so modifications to it
// won't affect the actual registry.
//
// Returns:
//   - A map of composite keys (sourceType:listType) to processors
func (pr *ProcessorRegistry) ListProcessors() map[string]Processor {
	pr.registryMutex.RLock()
	defer pr.registryMutex.RUnlock()

	// Create a manual copy for compatibility with Go versions before 1.21
	result := make(map[string]Processor, len(pr.customProcessorRegistry))
	for k, v := range pr.customProcessorRegistry {
		result[k] = v
	}
	return result
}

var Processors = NewProcessorRegistry()
