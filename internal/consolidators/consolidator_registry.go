package consolidators

import (
	"fmt"
	"sync"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

type Consolidator interface {
	Consolidate(logger *multilog.Logger, processedFiles []c.ProcessedFile) (u.StringSet, []c.FileInfo)
	FilterEntries(logger *multilog.Logger, entrySet u.StringSet, filterSet u.StringSet) (u.StringSet, u.StringSet)
	SaveEntries(logger *multilog.Logger, entrySet u.StringSet, filePath string) error
	IsValid(processedFile c.ProcessedFile) bool
	GetSourceType() string
	GetListType() string
}

func createRegistryKey(sourceType, listType string) string {
	return fmt.Sprintf("%s:%s", sourceType, listType)
}

type ConsolidatorRegistry struct {
	customConsolidatorRegistry map[string]Consolidator
	registryMutex              sync.RWMutex
}

func NewConsolidatorRegistry() *ConsolidatorRegistry {
	return &ConsolidatorRegistry{
		customConsolidatorRegistry: make(map[string]Consolidator),
	}
}

func (cr *ConsolidatorRegistry) RegisterConsolidator(sourceType, listType string, processor Consolidator) {
	cr.registryMutex.Lock()
	defer cr.registryMutex.Unlock()
	key := createRegistryKey(sourceType, listType)
	cr.customConsolidatorRegistry[key] = processor
}

func (cr *ConsolidatorRegistry) GetConsolidator(sourceType, listType string) (Consolidator, bool) {
	key := createRegistryKey(sourceType, listType)
	cr.registryMutex.RLock()
	processor, ok := cr.customConsolidatorRegistry[key]
	cr.registryMutex.RUnlock()
	return processor, ok
}

func (cr *ConsolidatorRegistry) ListConsolidators() map[string]Consolidator {
	cr.registryMutex.RLock()
	defer cr.registryMutex.RUnlock()

	result := make(map[string]Consolidator, len(cr.customConsolidatorRegistry))
	for k, v := range cr.customConsolidatorRegistry {
		result[k] = v
	}
	return result
}

var Consolidators = NewConsolidatorRegistry()
