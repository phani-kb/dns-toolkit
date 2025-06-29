package processors

import (
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/multilog"
)

// ProcessFunc is a function type for the specific processing logic
type ProcessFunc func(logger *multilog.Logger, content string) ([]string, []string)

// GenericProcessor is a customizable processor that uses a function for processing
type GenericProcessor struct {
	processFunc ProcessFunc
	BaseProcessor
}

// Process delegates to the custom process function
func (gp *GenericProcessor) Process(logger *multilog.Logger, content string) ([]string, []string) {
	return gp.processFunc(logger, content)
}

// NewGenericProcessor creates a new processor with custom processing logic
func NewGenericProcessor(sourceType, listType string, processFunc ProcessFunc) *GenericProcessor {
	return &GenericProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
		processFunc:   processFunc,
	}
}

// RegisterGenericProcessor registers a processor with a custom processing function
func RegisterGenericProcessor(
	sourceType string,
	processFunc ProcessFunc,
	blocklistOnly bool,
	allowlistOnly bool,
) {
	factory := func(st string, lt string) Processor {
		return NewGenericProcessor(st, lt, processFunc)
	}

	if !allowlistOnly {
		Processors.RegisterProcessor(
			sourceType,
			constants.ListTypeBlocklist,
			factory(sourceType, constants.ListTypeBlocklist),
		)
	}

	if !blocklistOnly {
		Processors.RegisterProcessor(
			sourceType,
			constants.ListTypeAllowlist,
			factory(sourceType, constants.ListTypeAllowlist),
		)
	}
}

// SpecialProcessorRegistry provides a way to register processors with special conditions different handling
// for blocklists vs. allowlists
type SpecialProcessorRegistry struct {
	processorFactories map[string]map[string]ProcessorFactory
}

// NewSpecialProcessorRegistry creates a new registry for special processors
func NewSpecialProcessorRegistry() *SpecialProcessorRegistry {
	return &SpecialProcessorRegistry{
		processorFactories: make(map[string]map[string]ProcessorFactory),
	}
}

// Register adds a processor factory for a specific source and list type
func (r *SpecialProcessorRegistry) Register(sourceType, listType string, factory ProcessorFactory) {
	if _, exists := r.processorFactories[sourceType]; !exists {
		r.processorFactories[sourceType] = make(map[string]ProcessorFactory)
	}
	r.processorFactories[sourceType][listType] = factory
}

// GetProcessor returns a processor for the requested source and list types
func (r *SpecialProcessorRegistry) GetProcessor(sourceType, listType string) (Processor, bool) {
	if factories, exists := r.processorFactories[sourceType]; exists {
		if factory, exists := factories[listType]; exists {
			return factory(sourceType, listType), true
		}
	}
	return nil, false
}

// SpecialProcessors Global instance of the special processor registry
var SpecialProcessors = NewSpecialProcessorRegistry()
