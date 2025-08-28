package cmd

import (
	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/multilog"
)

// ConsolidationManager handles general consolidation conflict resolution
type ConsolidationManager struct {
	logger *multilog.Logger
}

// NewConsolidationManager creates a new manager
func NewConsolidationManager(logger *multilog.Logger) *ConsolidationManager {
	return &ConsolidationManager{
		logger: logger,
	}
}

// GenerateConflictReport generates conflict report for general consolidation
func (cm *ConsolidationManager) GenerateConflictReport(processedFiles []c.ProcessedFile) error {
	cm.logger.Infof("Building resolution sets for conflict report...")

	// Build resolution sets for conflict analysis
	allowByType, blockByType, conflicts, manualAllowToBlock, manualBlockToAllow, detailsMap := BuildResolutionSets(
		cm.logger,
		processedFiles,
	)

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
