package cmd

import (
	"github.com/phani-kb/dns-toolkit/internal/db"
	"github.com/phani-kb/multilog"
)

// ConsolidationManager handles general consolidation conflict resolution
type ConsolidationManager struct {
	logger   *multilog.Logger
	database *db.DB
}

// NewConsolidationManager creates a new manager
func NewConsolidationManager(logger *multilog.Logger, database *db.DB) *ConsolidationManager {
	return &ConsolidationManager{
		logger:   logger,
		database: database,
	}
}

// GenerateConflictReport generates conflict report for general consolidation
func (cm *ConsolidationManager) GenerateConflictReport() error {
	cm.logger.Infof("Building resolution sets for conflict report...")

	// build resolution sets for conflict analysis using DB
	allowByType, blockByType, conflicts, manualAllowToBlock, manualBlockToAllow, detailsMap, err := GetResolutionSets(
		cm.logger,
		cm.database,
	)
	if err != nil {
		return err
	}

	result := &ResolutionResult{
		AllowByType: allowByType,
		BlockByType: blockByType,
		Conflicts:   conflicts,
		DetailsMap:  detailsMap,
	}
	result.ManualOverride.AllowToBlock = manualAllowToBlock
	result.ManualOverride.BlockToAllow = manualBlockToAllow

	overridesPath, err := writeOverrideSummary(cm.logger, result)
	if err != nil {
		return err
	}

	reportPath, err := GenerateConflictReport(cm.logger, overridesPath)
	if err != nil {
		return err
	}

	if reportPath != "" {
		cm.logger.Infof("Generated conflicts report: %s", reportPath)
	}

	return nil
}
