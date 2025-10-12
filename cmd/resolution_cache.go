package cmd

import (
	"strings"
	"sync"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

var (
	resolutionCacheMu  sync.Mutex
	resolutionCacheKey string

	resolutionCachedAllow       map[string]u.StringSet
	resolutionCachedBlock       map[string]u.StringSet
	resolutionCachedConflicts   []ConflictDetail
	resolutionCachedManualAllow map[string]struct{}
	resolutionCachedManualBlock map[string]struct{}
	resolutionCachedDetails     map[string]ConflictDetail
)

func getProcessedFilesKey(processedFiles []c.ProcessedFile) string {
	parts := make([]string, 0, len(processedFiles))
	for _, pf := range processedFiles {
		parts = append(parts, pf.Filepath+":"+pf.Checksum)
	}
	return strings.Join(parts, "|")
}

// GetCachedResolutionSets returns cached resolution sets
func GetCachedResolutionSets(logger *multilog.Logger, processedFiles []c.ProcessedFile) (
	map[string]u.StringSet,
	map[string]u.StringSet,
	[]ConflictDetail,
	map[string]struct{},
	map[string]struct{},
	map[string]ConflictDetail,
) {
	key := getProcessedFilesKey(processedFiles)

	resolutionCacheMu.Lock()
	defer resolutionCacheMu.Unlock()

	if key != "" && key == resolutionCacheKey && resolutionCachedAllow != nil {
		return resolutionCachedAllow,
			resolutionCachedBlock,
			resolutionCachedConflicts,
			resolutionCachedManualAllow,
			resolutionCachedManualBlock,
			resolutionCachedDetails
	}

	allowByType, blockByType, conflicts, manualAllowToBlock, manualBlockToAllow, detailsMap := BuildResolutionSets(
		logger,
		processedFiles,
	)

	resolutionCacheKey = key
	resolutionCachedAllow = allowByType
	resolutionCachedBlock = blockByType
	resolutionCachedConflicts = conflicts
	resolutionCachedManualAllow = manualAllowToBlock
	resolutionCachedManualBlock = manualBlockToAllow
	resolutionCachedDetails = detailsMap

	return allowByType, blockByType, conflicts, manualAllowToBlock, manualBlockToAllow, detailsMap
}
