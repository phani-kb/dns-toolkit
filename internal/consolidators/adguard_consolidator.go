package consolidators

import (
	"strings"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

const sourceType = "adguard"

const adguardExceptionPrefix = "@@"

var adguardExceptionSuffixesToStrip = []string{
	"$important",
}

type AdguardConsolidator struct {
	BaseConsolidator
}

func NewAdguardConsolidator(sourceType, listType string) *AdguardConsolidator {
	return &AdguardConsolidator{
		BaseConsolidator: NewBaseConsolidator(sourceType, listType),
	}
}

func (c *AdguardConsolidator) Consolidate(
	logger *multilog.Logger,
	processedFiles []c.ProcessedFile,
) (u.StringSet, []c.FileInfo) {
	return c.BaseConsolidator.Consolidate(logger, processedFiles)
}

func (c *AdguardConsolidator) FilterEntries(
	logger *multilog.Logger,
	entrySet u.StringSet,
	filterSet u.StringSet,
) (u.StringSet, u.StringSet) {
	newFilterSet := u.NewStringSet(nil)
	for entry := range filterSet {
		if strings.HasPrefix(entry, adguardExceptionPrefix) {
			newEntry := entry[len(adguardExceptionPrefix):]
			for _, suffix := range adguardExceptionSuffixesToStrip {
				newEntry = strings.TrimSuffix(newEntry, suffix)
			}
			newFilterSet.Add(newEntry)
		} else {
			newFilterSet.Add(entry)
		}
	}
	filterSet = newFilterSet
	return c.BaseConsolidator.FilterEntries(logger, entrySet, filterSet)
}

func (c *AdguardConsolidator) SaveEntries(
	logger *multilog.Logger,
	entrySet u.StringSet,
	filePath string,
) error {
	return c.BaseConsolidator.SaveEntries(logger, entrySet, filePath)
}

func (c *AdguardConsolidator) IsValid(processedFile c.ProcessedFile) bool {
	return c.BaseConsolidator.IsValid(processedFile)
}

func init() {
	RegisterConsolidatorTypes(sourceType, commonSupportedListTypes, func(_, lt string) Consolidator {
		return NewAdguardConsolidator(sourceType, lt)
	})
}
