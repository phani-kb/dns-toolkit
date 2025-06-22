package consolidators

import (
	c "github.com/phani-kb/dns-toolkit/internal/common"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

type ConsolidateFunc func(logger *multilog.Logger, processedFiles []c.ProcessedFile) (u.StringSet, []c.FileInfo)

type GenericConsolidator struct {
	BaseConsolidator
	consolidateFunc ConsolidateFunc
}

func (gc *GenericConsolidator) Consolidate(
	logger *multilog.Logger,
	processedFiles []c.ProcessedFile,
) (u.StringSet, []c.FileInfo) {
	return gc.consolidateFunc(logger, processedFiles)
}

func (gc *GenericConsolidator) IsValid(processedFile c.ProcessedFile) bool {
	return gc.BaseConsolidator.IsValid(processedFile)
}

func NewGenericConsolidator(sourceType, listType string, consolidateFunc ConsolidateFunc) *GenericConsolidator {
	return &GenericConsolidator{
		BaseConsolidator: NewBaseConsolidator(sourceType, listType),
		consolidateFunc:  consolidateFunc,
	}
}

func RegisterGenericConsolidator(
	sourceType string,
	listType string,
	consolidateFunc ConsolidateFunc,
) {
	factory := func(st string, lt string) Consolidator {
		return NewGenericConsolidator(st, lt, consolidateFunc)
	}

	Consolidators.RegisterConsolidator(
		sourceType,
		listType,
		factory(sourceType, listType),
	)
}

func RegisterGenericConsolidatorTypes(
	sourceType string,
	listTypes []string,
	consolidateFunc ConsolidateFunc,
) {
	for _, listType := range listTypes {
		RegisterGenericConsolidator(
			sourceType,
			listType,
			consolidateFunc,
		)
	}
}
