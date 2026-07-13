package cmd

import (
	"fmt"

	"github.com/phani-kb/dns-toolkit/internal/db"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

// GetResolutionSets builds resolution sets from the database.
func GetResolutionSets(logger *multilog.Logger, database *db.DB) (
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

	allowByType, blockByType, conflicts, manualAllowToBlock, manualBlockToAllow, detailsMap, err := BuildResolutionSets(
		logger,
		database,
	)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	return allowByType, blockByType, conflicts, manualAllowToBlock, manualBlockToAllow, detailsMap, nil
}
