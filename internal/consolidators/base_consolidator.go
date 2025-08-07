package consolidators

import (
	c "github.com/phani-kb/dns-toolkit/internal/common"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

type BaseConsolidator struct {
	sourceType string
	listType   string
}

func NewBaseConsolidator(sourceType, listType string) BaseConsolidator {
	return BaseConsolidator{
		sourceType: sourceType,
		listType:   listType,
	}
}

func (bc *BaseConsolidator) GetSourceType() string {
	return bc.sourceType
}

func (bc *BaseConsolidator) GetListType() string {
	return bc.listType
}

func (bc *BaseConsolidator) IsValid(processedFile c.ProcessedFile) bool {
	return processedFile.GenericSourceType == bc.sourceType && processedFile.ListType == bc.listType
}

func (bc *BaseConsolidator) Consolidate(
	logger *multilog.Logger,
	processedFiles []c.ProcessedFile,
) (u.StringSet, []c.FileInfo) {
	consolidatedSet := u.NewStringSet([]string{})
	var fileInfos []c.FileInfo

	for _, processedFile := range processedFiles {
		if !bc.IsValid(processedFile) {
			logger.Debugf("Skipping invalid processed file: %s", processedFile.Filepath)
			continue
		}
		logger.Debugf("Reading entry(s) from file: %s", processedFile.Filepath)
		entries, duplicates, err := u.ReadEntriesFromFile(logger, processedFile.Filepath)
		if err != nil {
			logger.Errorf("Error reading entry(s) from file %s: %v", processedFile.Filepath, err)
			continue
		}
		if len(entries) == 0 {
			logger.Infof("No entry(s) found in file: %s", processedFile.Filepath)
			continue
		}
		if len(entries) != processedFile.NumberOfEntries {
			logger.Warnf(
				"Entry count mismatch for file %s: expected %d, got %d, duplicates: %d",
				processedFile.Filepath,
				processedFile.NumberOfEntries,
				len(entries),
				duplicates,
			)
			continue
		}

		logger.Debugf("Adding %d entry(s) from file: %s", len(entries), processedFile.Filepath)

		consolidatedSet.AddAll(entries, processedFile.MustConsider)

		fileInfos = append(fileInfos, c.FileInfo{
			Name:         processedFile.Name,
			SourceType:   processedFile.ActualSourceType,
			Filepath:     processedFile.Filepath,
			MustConsider: processedFile.MustConsider,
			Count:        len(entries),
		})
	}

	return consolidatedSet, fileInfos
}

func (bc *BaseConsolidator) FilterEntries(
	logger *multilog.Logger,
	entrySet u.StringSet,
	filterSet u.StringSet,
) (u.StringSet, u.StringSet) {
	filteredSet := u.NewStringSet([]string{})
	ignoredSet := u.NewStringSet([]string{})

	if len(entrySet) == 0 {
		logger.Debugf("No entry to filter")
		return filteredSet, ignoredSet
	}

	logger.Infof(
		"Filtering %d entry(s) with %d ignored entry(s), source type: %s, list type: %s",
		len(entrySet),
		len(filterSet),
		bc.sourceType,
		bc.listType,
	)

	for entry := range entrySet {
		mustConsider, _ := entrySet.Get(entry)
		mustFilter, entryExistsInFilterSet := filterSet.Get(entry)
		if mustFilter {
			if mustConsider {
				logger.Warnf("Entry %s is marked as must consider, but it is also marked as must filter", entry)
				filteredSet.AddWithConsider(entry, mustFilter)
			} else {
				ignoredSet.Add(entry)
			}
		} else {
			if !entryExistsInFilterSet {
				filteredSet.AddWithConsider(entry, mustConsider)
			} else {
				if mustConsider {
					filteredSet.AddWithConsider(entry, mustFilter)
				} else {
					ignoredSet.Add(entry)
				}
			}
		}
	}

	logger.Infof("Entry(s) %d filtered, %d ignored", len(filteredSet), len(ignoredSet))

	return filteredSet, ignoredSet
}

func (bc *BaseConsolidator) SaveEntries(
	logger *multilog.Logger,
	entrySet u.StringSet,
	filePath string,
) error {
	logger.Debugf("Saving %d entry(s) to file: %s", len(entrySet), filePath)
	err := u.WriteEntriesToFile(logger, filePath, entrySet.ToSlice())
	return err
}

func RegisterConsolidator(sourceType string, listType string, consolidatorCreator func(string, string) Consolidator) {
	Consolidators.RegisterConsolidator(
		sourceType,
		listType,
		consolidatorCreator(sourceType, listType),
	)
}

func RegisterConsolidatorTypes(
	sourceType string,
	listTypes []string,
	consolidatorCreator func(string, string) Consolidator,
) {
	for _, listType := range listTypes {
		Consolidators.RegisterConsolidator(
			sourceType,
			listType,
			consolidatorCreator(sourceType, listType),
		)
	}
}
