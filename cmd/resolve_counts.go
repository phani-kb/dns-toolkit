package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/db"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

// ConflictDetail captures entries that need manual review
type ConflictDetail struct {
	Entry             string   `json:"entry"`
	GenericSourceType string   `json:"generic_source_type"`
	BlockSources      []string `json:"block_sources"`
	AllowSources      []string `json:"allow_sources"`
	BlockCount        int      `json:"block_count"`
	AllowCount        int      `json:"allow_count"`
}

// OverrideRecord represents a single override decision
type OverrideRecord struct {
	Entry      string   `json:"entry"`
	Decision   string   `json:"decision"`
	Reason     string   `json:"reason"`
	BlockSrcs  []string `json:"block_sources"`
	AllowSrcs  []string `json:"allow_sources"`
	BlockCount int      `json:"block_count"`
	AllowCount int      `json:"allow_count"`
}

// ResolutionResult contains the results of conflict resolution
type ResolutionResult struct {
	ManualOverride struct {
		AllowToBlock map[string]struct{}
		BlockToAllow map[string]struct{}
	}
	AllowByType map[string]u.StringSet
	BlockByType map[string]u.StringSet
	DetailsMap  map[string]ConflictDetail
	Conflicts   []ConflictDetail
}

// Decision constants
const (
	DecisionBlock    = "block"
	DecisionAllow    = "allow"
	DecisionConflict = "conflict"

	ReasonCounts            = "counts"
	ReasonManualForcedBlock = "manual_forced_block"
	ReasonManualForcedAllow = "manual_forced_allow"
	ReasonEqualCounts       = "equal_counts"
)

// BuildResolutionSets performs count-based resolution using DB queries.
// It only loads conflicting entries.
func BuildResolutionSets(
	logger *multilog.Logger,
	database *db.DB,
) (
	map[string]u.StringSet,
	map[string]u.StringSet,
	[]ConflictDetail,
	map[string]struct{},
	map[string]struct{},
	map[string]ConflictDetail,
	error,
) {
	result := &ResolutionResult{
		AllowByType: make(map[string]u.StringSet),
		BlockByType: make(map[string]u.StringSet),
		DetailsMap:  make(map[string]ConflictDetail),
	}

	entriesRepo := db.NewEntriesRepo(database)

	// query all entries to determine which have sources on both sides
	allRows, err := entriesRepo.GetAllEntryCounts(context.Background())
	if err != nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("querying all entry counts: %w", err)
	}

	logger.Infof("Total unique entries in database: %d", len(allRows))

	allowWins := 0
	blockWins := 0
	equalCounts := 0
	conflicts := make([]ConflictDetail, 0)

	for _, row := range allRows {
		// get the actual source names for this entry
		blockSources, allowSources, err2 := entriesRepo.GetSourcesForEntry(
			context.Background(),
			row.Entry,
			row.GenericSourceType,
		)
		if err2 != nil {
			logger.Warnf("Failed to get sources for entry %s: %v", row.Entry, err2)
			continue
		}

		detail := ConflictDetail{
			Entry:             row.Entry,
			GenericSourceType: row.GenericSourceType,
			BlockSources:      blockSources,
			AllowSources:      allowSources,
			BlockCount:        row.BlockCount,
			AllowCount:        row.AllowCount,
		}

		result.DetailsMap[row.Entry] = detail

		// apply threshold logic
		minAllow := 1
		minBlock := 1
		if AppConfig != nil && AppConfig.DNSToolkit.Override.Enabled {
			for _, t := range AppConfig.DNSToolkit.Override.Thresholds {
				if strings.EqualFold(t.Name, "allowlist") && t.MinSources > 0 {
					minAllow = t.MinSources
				}
				if strings.EqualFold(t.Name, "blocklist") && t.MinSources > 0 {
					minBlock = t.MinSources
				}
			}
		}

		switch {
		case row.BlockCount > row.AllowCount:
			if row.AllowCount == 0 {
				addToBlockSetsDB(result, row.Entry, row.GenericSourceType)
				blockWins++
			} else {
				if row.BlockCount >= minBlock {
					addToBlockSetsDB(result, row.Entry, row.GenericSourceType)
					blockWins++
				} else {
					conflicts = append(conflicts, detail)
				}
			}
		case row.AllowCount > row.BlockCount:
			if row.BlockCount == 0 {
				addToAllowSetsDB(result, row.Entry, row.GenericSourceType)
				allowWins++
			} else {
				if row.AllowCount >= minAllow {
					addToAllowSetsDB(result, row.Entry, row.GenericSourceType)
					allowWins++
				} else {
					conflicts = append(conflicts, detail)
				}
			}
		default:
			if row.BlockCount > 0 { // equal non-zero counts = conflict
				conflicts = append(conflicts, detail)
				equalCounts++
			}
		}
	}

	result.Conflicts = conflicts
	logger.Infof(
		"resolveByCounts (from DB): allowWins=%d, blockWins=%d, equalCounts=%d, conflicts=%d",
		allowWins,
		blockWins,
		equalCounts,
		len(conflicts),
	)

	// apply manual overrides
	applyManualOverrides(logger, result)

	// filter out manually overridden entries from conflicts
	result.Conflicts = filterConflictsAfterOverrides(result)
	logger.Infof("Total conflicts after manual overrides: %d", len(result.Conflicts))

	return result.AllowByType,
		result.BlockByType,
		result.Conflicts,
		result.ManualOverride.AllowToBlock,
		result.ManualOverride.BlockToAllow,
		result.DetailsMap,
		nil
}

// filterConflictsAfterOverrides removes manually overridden entries from conflicts list
func filterConflictsAfterOverrides(result *ResolutionResult) []ConflictDetail {
	filteredConflicts := make([]ConflictDetail, 0)

	for _, conflict := range result.Conflicts {
		// skip if this entry was manually overridden
		_, isAllowToBlock := result.ManualOverride.AllowToBlock[conflict.Entry]
		_, isBlockToAllow := result.ManualOverride.BlockToAllow[conflict.Entry]

		if !isAllowToBlock && !isBlockToAllow {
			filteredConflicts = append(filteredConflicts, conflict)
		}
	}

	return filteredConflicts
}

// applyManualOverrides applies manual overrides.
func applyManualOverrides(logger *multilog.Logger, result *ResolutionResult) {
	manualAllowToBlockByType, manualBlockToAllowByType := readCustomOverrides(logger)

	result.ManualOverride.AllowToBlock = make(map[string]struct{})
	result.ManualOverride.BlockToAllow = make(map[string]struct{})

	for gst, entries := range manualAllowToBlockByType {
		for entry := range entries {
			// only apply if entry exists in DetailsMap
			if detail, ok := result.DetailsMap[entry]; ok && detail.GenericSourceType == gst {
				moveToBlockSet(result, entry, gst)
				result.ManualOverride.AllowToBlock[entry] = struct{}{}
			}
		}
	}

	// Apply forced allows (takes precedence)
	for gst, entries := range manualBlockToAllowByType {
		for entry := range entries {
			if detail, ok := result.DetailsMap[entry]; ok && detail.GenericSourceType == gst {
				moveToAllowSet(result, entry, gst)
				result.ManualOverride.BlockToAllow[entry] = struct{}{}
				// Remove from forced block if it exists
				delete(result.ManualOverride.AllowToBlock, entry)
			}
		}
	}
}

// writeOverrideSummary writes the override summary JSON
func writeOverrideSummary(logger *multilog.Logger, result *ResolutionResult) (string, error) {
	overrides := buildOverrideRecords(logger, result)

	overridesPath := filepath.Join(
		constants.SummaryDir,
		constants.SummaryTypesOutputSummaryFileMap[constants.SummaryTypeOverrides],
	)

	data, err := json.MarshalIndent(overrides, "", "  ")
	if err != nil {
		return overridesPath, fmt.Errorf("failed to marshal overrides: %w", err)
	}

	if err := os.WriteFile(overridesPath, data, 0o644); err != nil {
		return overridesPath, fmt.Errorf("failed to write overrides json: %w", err)
	}

	return overridesPath, nil
}

// buildOverrideRecords creates override records from resolution results
func buildOverrideRecords(logger *multilog.Logger, result *ResolutionResult) []OverrideRecord {
	overrides := make([]OverrideRecord, 0)

	totalAllowEntries := 0
	for typeName, set := range result.AllowByType {
		if set != nil {
			totalAllowEntries += set.Size()
			logger.Debugf("AllowByType[%s]: %d entries", typeName, set.Size())
		} else {
			logger.Debugf("AllowByType[%s]: nil", typeName)
		}
	}
	logger.Infof("buildOverrideRecords: Total AllowByType entries across all types: %d", totalAllowEntries)

	overrides = append(overrides, getAutomaticDecisions(logger, result)...)

	overrides = append(overrides, getManualOverrideRecords(logger, result)...)

	overrides = append(overrides, getConflictRecords(result.Conflicts)...)

	return overrides
}

func ensureStringSet(setMap map[string]u.StringSet, key string) {
	if setMap[key] == nil {
		setMap[key] = u.NewStringSet([]string{})
	}
}

func moveToBlockSet(result *ResolutionResult, entry, gst string) {
	// Remove from allow sets
	if result.AllowByType[gst] != nil {
		result.AllowByType[gst].Remove(entry)
	}
	ensureStringSet(result.BlockByType, gst)
	result.BlockByType[gst].Add(entry)
}

func moveToAllowSet(result *ResolutionResult, entry, gst string) {
	// Remove from block sets
	if result.BlockByType[gst] != nil {
		result.BlockByType[gst].Remove(entry)
	}

	ensureStringSet(result.AllowByType, gst)
	result.AllowByType[gst].Add(entry)
}

// addToBlockSetsDB adds an entry to block sets.
func addToBlockSetsDB(result *ResolutionResult, entry, gst string) {
	ensureStringSet(result.BlockByType, gst)
	result.BlockByType[gst].Add(entry)
}

// addToAllowSetsDB adds an entry to allow sets.
func addToAllowSetsDB(result *ResolutionResult, entry, gst string) {
	ensureStringSet(result.AllowByType, gst)
	result.AllowByType[gst].Add(entry)
}

func getAutomaticDecisions(logger *multilog.Logger, result *ResolutionResult) []OverrideRecord {
	records := make([]OverrideRecord, 0)

	allowCountGreater := 0
	allowDecisionSet := 0
	allowDecisionNotSet := 0

	for entry, detail := range result.DetailsMap {
		if isManualOverride(entry, result) {
			continue
		}

		if len(detail.BlockSources) == 0 || len(detail.AllowSources) == 0 {
			continue
		}

		var decision string
		switch {
		case detail.BlockCount > detail.AllowCount:
			for _, set := range result.BlockByType {
				if set != nil && set.Contains(entry) {
					decision = DecisionBlock
					break
				}
			}
		case detail.AllowCount > detail.BlockCount:
			allowCountGreater++
			found := false
			for _, set := range result.AllowByType {
				if set != nil && set.Contains(entry) {
					decision = DecisionAllow
					found = true
					break
				}
			}
			if found {
				allowDecisionSet++
			} else {
				allowDecisionNotSet++
				logger.Debugf("Entry with allowCount>blockCount NOT in AllowByType: %s (allow=%d, block=%d)",
					entry, detail.AllowCount, detail.BlockCount)
			}
		default:
			continue
		}

		if decision == "" {
			// conflict
			continue
		}

		records = append(records, OverrideRecord{
			Entry:      entry,
			Decision:   decision,
			Reason:     ReasonCounts,
			BlockCount: detail.BlockCount,
			AllowCount: detail.AllowCount,
			BlockSrcs:  detail.BlockSources,
			AllowSrcs:  detail.AllowSources,
		})
	}

	return records
}

func getManualOverrideRecords(logger *multilog.Logger, result *ResolutionResult) []OverrideRecord {
	records := make([]OverrideRecord, 0)

	// Forced blocks (only if not also forced allow)
	for entry := range result.ManualOverride.AllowToBlock {
		if _, hasConflict := result.ManualOverride.BlockToAllow[entry]; hasConflict {
			logger.Warnf("Skipping conflicting manual override for entry: %s", entry)
			continue // Skip conflicting manual overrides
		}

		if detail, ok := result.DetailsMap[entry]; ok && hasSourcesOnBothSides(detail) {
			records = append(records, OverrideRecord{
				Entry:      entry,
				Decision:   DecisionBlock,
				Reason:     ReasonManualForcedBlock,
				BlockCount: detail.BlockCount,
				AllowCount: detail.AllowCount,
				BlockSrcs:  detail.BlockSources,
				AllowSrcs:  detail.AllowSources,
			})
		}
	}

	for entry := range result.ManualOverride.BlockToAllow {
		if detail, ok := result.DetailsMap[entry]; ok && hasSourcesOnBothSides(detail) {
			records = append(records, OverrideRecord{
				Entry:      entry,
				Decision:   DecisionAllow,
				Reason:     ReasonManualForcedAllow,
				BlockCount: detail.BlockCount,
				AllowCount: detail.AllowCount,
				BlockSrcs:  detail.BlockSources,
				AllowSrcs:  detail.AllowSources,
			})
		}
	}

	return records
}

func getConflictRecords(conflicts []ConflictDetail) []OverrideRecord {
	records := make([]OverrideRecord, 0, len(conflicts))

	for _, conflict := range conflicts {
		if hasSourcesOnBothSides(conflict) {
			records = append(records, OverrideRecord{
				Entry:      conflict.Entry,
				Decision:   DecisionConflict,
				Reason:     ReasonEqualCounts,
				BlockCount: conflict.BlockCount,
				AllowCount: conflict.AllowCount,
				BlockSrcs:  conflict.BlockSources,
				AllowSrcs:  conflict.AllowSources,
			})
		}
	}

	return records
}

func isManualOverride(entry string, result *ResolutionResult) bool {
	_, isAllowToBlock := result.ManualOverride.AllowToBlock[entry]
	_, isBlockToAllow := result.ManualOverride.BlockToAllow[entry]
	return isAllowToBlock || isBlockToAllow
}

func hasSourcesOnBothSides(detail ConflictDetail) bool {
	return len(detail.BlockSources) > 0 && len(detail.AllowSources) > 0
}

// readCustomOverrides reads manual override files
func readCustomOverrides(logger *multilog.Logger) (map[string]map[string]struct{}, map[string]map[string]struct{}) {
	manualAllowToBlockByType := make(map[string]map[string]struct{})
	manualBlockToAllowByType := make(map[string]map[string]struct{})

	for gst, overrideMap := range constants.CustomOverrideFilesMap {
		if overrideMap == nil {
			continue
		}

		if path, ok := overrideMap[constants.ForcedAllow]; ok {
			if entries := readManualEntries(logger, "manual allowlist for "+gst, path); len(entries) > 0 {
				manualBlockToAllowByType[gst] = entries
			}
		}

		if path, ok := overrideMap[constants.ForcedBlock]; ok {
			if entries := readManualEntries(logger, "manual blocklist for "+gst, path); len(entries) > 0 {
				manualAllowToBlockByType[gst] = entries
			}
		}
	}

	return manualAllowToBlockByType, manualBlockToAllowByType
}

// readManualEntries reads entries from a manual override file
func readManualEntries(logger *multilog.Logger, label, path string) map[string]struct{} {
	entries, duplicates, err := u.ReadEntriesFromFile(logger, path)
	if err != nil {
		logger.Debugf("Failed to read %s %s: %v", label, path, err)
		return nil
	}

	if duplicates > 0 {
		logger.Infof("Found %d duplicate entries in %s %s", duplicates, label, path)
	}

	result := make(map[string]struct{}, len(entries))
	for _, entry := range entries {
		entry = strings.TrimSpace(entry)
		result[entry] = struct{}{}
	}

	return result
}
