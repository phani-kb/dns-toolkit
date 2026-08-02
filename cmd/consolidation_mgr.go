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

// GenerateConflictReport generates conflict report by using resolution sets
func (cm *ConsolidationManager) GenerateConflictReport() error {
	cm.logger.Infof("Building resolution sets for conflict report...")

	result, err := GetResolutionSets(cm.logger, cm.database)
	if err != nil {
		return err
	}

	return cm.GenerateConflictReportFromResult(result)
}

// GenerateConflictReportFromResult generates conflict report by using resolution sets
func (cm *ConsolidationManager) GenerateConflictReportFromResult(result *ResolutionResult) error {
	if result == nil {
		cm.logger.Warnf("No resolution result provided, skipping conflict report")
		return nil
	}

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
