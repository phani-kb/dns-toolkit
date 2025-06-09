package processors

import (
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/multilog"
)

// ProcessFunc is a function type for the specific processing logic
type ProcessFunc func(logger *multilog.Logger, content string) ([]string, []string)

// GenericProcessor is a customizable processor that uses a function for processing
type GenericProcessor struct {
	BaseProcessor
	processFunc ProcessFunc
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
