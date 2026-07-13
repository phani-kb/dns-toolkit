package cmd

import (
	"fmt"

	"github.com/phani-kb/dns-toolkit/internal/db"
	"github.com/phani-kb/multilog"
)

// GetResolutionSets builds resolution sets and returns a ResolutionResult.
func GetResolutionSets(logger *multilog.Logger, database *db.DB) (*ResolutionResult, error) {
	if database == nil {
		return nil, fmt.Errorf("database is required for resolution")
	}

	allowByType, blockByType, conflicts, manualAllowToBlock, manualBlockToAllow, detailsMap, err := BuildResolutionSets(
		logger,
		database,
	)
	if err != nil {
		return nil, err
	}

	result := &ResolutionResult{
		AllowByType: allowByType,
		BlockByType: blockByType,
		Conflicts:   conflicts,
		DetailsMap:  detailsMap,
	}
	result.ManualOverride.AllowToBlock = manualAllowToBlock
	result.ManualOverride.BlockToAllow = manualBlockToAllow

	return result, nil
}
