package consolidators

import (
	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

var commonSourceTypes = []string{
	constants.SourceTypeIpv4,
	constants.SourceTypeIpv6,
	constants.SourceTypeCidrIpv4,
	constants.SourceTypeDomain,
}

var commonSupportedListTypes = []string{
	constants.ListTypeAllowlist,
	constants.ListTypeBlocklist,
}

type CommonConsolidator struct {
	BaseConsolidator
}

func NewCommonConsolidator(sourceType, listType string) *CommonConsolidator {
	return &CommonConsolidator{
		BaseConsolidator: NewBaseConsolidator(sourceType, listType),
	}
}

func (c *CommonConsolidator) Consolidate(
	logger *multilog.Logger,
	processedFiles []c.ProcessedFile,
) (u.StringSet, []c.FileInfo) {
	return c.BaseConsolidator.Consolidate(logger, processedFiles)
}

func (c *CommonConsolidator) FilterEntries(
	logger *multilog.Logger,
	entrySet u.StringSet,
	filterSet u.StringSet,
) (u.StringSet, u.StringSet) {
	return c.BaseConsolidator.FilterEntries(logger, entrySet, filterSet)
}

func (c *CommonConsolidator) SaveEntries(
	logger *multilog.Logger,
	entrySet u.StringSet,
	filePath string,
) error {
	return c.BaseConsolidator.SaveEntries(logger, entrySet, filePath)
}

func (c *CommonConsolidator) IsValid(processedFile c.ProcessedFile) bool {
	return c.BaseConsolidator.IsValid(processedFile)
}

func init() {
	for _, st := range commonSourceTypes {
		RegisterConsolidatorTypes(st, commonSupportedListTypes, func(st, lt string) Consolidator {
			return NewCommonConsolidator(st, lt)
		})
	}
}
