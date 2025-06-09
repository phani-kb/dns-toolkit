package processors

import "github.com/phani-kb/dns-toolkit/internal/constants"

// BaseProcessor provides common functionality for all processors
type BaseProcessor struct {
	sourceType string
	listType   string
}

// NewBaseProcessor creates a new base processor with the given source type and list type
func NewBaseProcessor(sourceType, listType string) BaseProcessor {
	return BaseProcessor{
		sourceType: sourceType,
		listType:   listType,
	}
}

// GetSourceType returns the source type identifier for this processor
func (bp *BaseProcessor) GetSourceType() string {
	return bp.sourceType
}

// GetListType returns the list type identifier for this processor
func (bp *BaseProcessor) GetListType() string {
	return bp.listType
}

// RegisterProcessor registers a processor for both blocklist and allowlist types
func RegisterProcessor(sourceType string, processorCreator func(string, string) Processor) {
	// Default register blocklist processor
	Processors.RegisterProcessor(
		sourceType,
		constants.ListTypeBlocklist,
		processorCreator(sourceType, constants.ListTypeBlocklist),
	)
}

func RegisterProcessorTypes(
	sourceType string,
	listTypes []string,
	processorCreator func(string, string) Processor,
) {
	for _, listType := range listTypes {
		Processors.RegisterProcessor(
			sourceType,
			listType,
			processorCreator(sourceType, listType),
		)
	}
}

// ProcessorFactory is a generic factory function type for creating processors
type ProcessorFactory func(sourceType, listType string) Processor
