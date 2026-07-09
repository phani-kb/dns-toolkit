package cmd

import (
	"fmt"
	"strings"
	"sync"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/db"
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

// GetCachedResolutionSets returns cached resolution sets, or builds them from the database if not cached.
func GetCachedResolutionSets(logger *multilog.Logger, database *db.DB, processedFiles []c.ProcessedFile) (
	map[string]u.StringSet,
	map[string]u.StringSet,
	[]ConflictDetail,
	map[string]struct{},
	map[string]struct{},
	map[string]ConflictDetail,
	error,
) {
	if database == nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf(
			"database is required for resolution",
		)
	}

	key := getProcessedFilesKey(processedFiles)

	resolutionCacheMu.Lock()
	defer resolutionCacheMu.Unlock()

	logger.Debugf("Cache check: files=%d, keyMatch=%v, cacheExists=%v",
		len(processedFiles), key == resolutionCacheKey, resolutionCachedAllow != nil)

	if key != "" && key == resolutionCacheKey && resolutionCachedAllow != nil {
		logger.Debugf("Counts FROM cache: allow=%d, block=%d, conflicts=%d, manualAllow=%d, manualBlock=%d",
			len(resolutionCachedAllow),
			len(resolutionCachedBlock),
			len(resolutionCachedConflicts),
			len(resolutionCachedManualAllow),
			len(resolutionCachedManualBlock),
		)
		return resolutionCachedAllow,
			resolutionCachedBlock,
			resolutionCachedConflicts,
			resolutionCachedManualAllow,
			resolutionCachedManualBlock,
			resolutionCachedDetails,
			nil
	}

	logger.Debugf("Cache miss - rebuilding resolution sets from database")

	allowByType, blockByType, conflicts, manualAllowToBlock, manualBlockToAllow, detailsMap, err := BuildResolutionSetsFromDB( // nolint: lll
		logger,
		database,
	)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	resolutionCacheKey = key
	resolutionCachedAllow = allowByType
	resolutionCachedBlock = blockByType
	resolutionCachedConflicts = conflicts
	resolutionCachedManualAllow = manualAllowToBlock
	resolutionCachedManualBlock = manualBlockToAllow
	resolutionCachedDetails = detailsMap

	logger.Infof("Counts TO cache: allow=%d, block=%d, conflicts=%d, manualAllow=%d, manualBlock=%d",
		len(resolutionCachedAllow),
		len(resolutionCachedBlock),
		len(resolutionCachedConflicts),
		len(resolutionCachedManualAllow),
		len(resolutionCachedManualBlock),
	)

	return allowByType, blockByType, conflicts, manualAllowToBlock, manualBlockToAllow, detailsMap, nil
}
