package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

// ConflictDetail captures entries that need manual review
type ConflictDetail struct {
	Entry        string   `json:"entry"`
	BlockSources []string `json:"block_sources"`
	AllowSources []string `json:"allow_sources"`
	BlockCount   int      `json:"block_count"`
	AllowCount   int      `json:"allow_count"`
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

// ResolveConflictsByCounts builds allowlist and blocklist and JSON summary.
// Uses count-based resolution: higher count wins, equal counts create conflicts.
func ResolveConflictsByCounts(
	logger *multilog.Logger,
	processedFiles []c.ProcessedFile,
) (string, string, string, error) {
	logger.Infof("Resolving conflicts and producing final sets...")

	allowByType, blockByType, conflicts, manualAllowToBlock, manualBlockToAllow, detailsMap := GetCachedResolutionSets(
		logger,
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

	allowPath, blockPath, err := writeResolvedLists(logger, result)
	if err != nil {
		return "", "", "", err
	}

	overridesPath, err := writeOverrideSummary(logger, result)
	if err != nil {
		return allowPath, blockPath, "", err
	}

	logger.Infof(
		"Consolidated allowlist (%d), blocklist (%d), overrides (%d)",
		countSetEntries(result.AllowByType),
		countSetEntries(result.BlockByType),
		len(result.DetailsMap),
	)

	return allowPath, blockPath, overridesPath, nil
}

// BuildResolutionSets processes files and returns complete resolution results
func BuildResolutionSets(
	logger *multilog.Logger,
	processedFiles []c.ProcessedFile,
) (
	map[string]u.StringSet,
	map[string]u.StringSet,
	[]ConflictDetail,
	map[string]struct{},
	map[string]struct{},
	map[string]ConflictDetail,
) {
	sourceMaps := buildSourceMaps(logger, processedFiles)

	// Resolve by counts
	result := &ResolutionResult{
		AllowByType: make(map[string]u.StringSet),
		BlockByType: make(map[string]u.StringSet),
		DetailsMap:  make(map[string]ConflictDetail),
	}

	result.Conflicts = resolveByCounts(logger, sourceMaps, result)
	logger.Infof("Total conflicts before manual overrides: %d", len(result.Conflicts))

	applyManualOverrides(logger, sourceMaps, result)

	// Filter out manually overridden entries from conflicts
	result.Conflicts = filterConflictsAfterOverrides(result)
	logger.Infof("Total conflicts after manual overrides: %d", len(result.Conflicts))

	fillDetailsForResolution(sourceMaps, result)

	return result.AllowByType,
		result.BlockByType,
		result.Conflicts,
		result.ManualOverride.AllowToBlock,
		result.ManualOverride.BlockToAllow,
		result.DetailsMap
}

// SourceMaps contains all source mapping data
type SourceMaps struct {
	BlockMap   map[string]map[string]struct{}
	AllowMap   map[string]map[string]struct{}
	EntryTypes map[string]map[string]struct{}
}

func buildSourceMaps(logger *multilog.Logger, processedFiles []c.ProcessedFile) *SourceMaps {
	maps := &SourceMaps{
		BlockMap:   make(map[string]map[string]struct{}),
		AllowMap:   make(map[string]map[string]struct{}),
		EntryTypes: make(map[string]map[string]struct{}),
	}

	for _, pf := range processedFiles {
		if !isValidProcessedFile(pf) {
			continue
		}

		entries, err := readFileEntries(logger, pf.Filepath)
		if err != nil {
			logger.Warnf("Skipping file %s: %v", pf.Filepath, err)
			continue
		}

		processFileEntries(maps, entries, pf)
	}

	return maps
}

// resolveByCounts performs count-based resolution with single pass
func resolveByCounts(logger *multilog.Logger, maps *SourceMaps, result *ResolutionResult) []ConflictDetail {
	conflicts := make([]ConflictDetail, 0)

	allEntries := getAllUniqueEntries(maps.BlockMap, maps.AllowMap)

	allowWins := 0
	blockWins := 0
	equalCounts := 0
	allowOnlyAdded := 0

	for entry := range allEntries {
		blockSources := getSourcesList(maps.BlockMap[entry])
		allowSources := getSourcesList(maps.AllowMap[entry])
		blockCount := len(blockSources)
		allowCount := len(allowSources)

		detail := ConflictDetail{
			Entry:        entry,
			BlockSources: blockSources,
			AllowSources: allowSources,
			BlockCount:   blockCount,
			AllowCount:   allowCount,
		}

		result.DetailsMap[entry] = detail

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
		case blockCount > allowCount:
			if allowCount == 0 {
				addToBlockSets(result, entry, maps.EntryTypes[entry])
				blockWins++
			} else {
				if blockCount >= minBlock {
					addToBlockSets(result, entry, maps.EntryTypes[entry])
					blockWins++
				} else {
					conflicts = append(conflicts, detail)
				}
			}
		case allowCount > blockCount:
			if blockCount == 0 {
				addToAllowSets(result, entry, maps.EntryTypes[entry])
				allowWins++
			} else {
				if allowCount >= minAllow {
					addToAllowSets(result, entry, maps.EntryTypes[entry])
					allowWins++
				} else {
					conflicts = append(conflicts, detail)
				}
			}
		default:
			if blockCount > 0 { // equal non-zero counts = conflict
				conflicts = append(conflicts, detail)
				equalCounts++
			}
		}
	}

	logger.Infof(
		"resolveByCounts: allowWins=%d (both sides), blockWins=%d, allowOnlyAdded=%d, equalCounts=%d, conflicts=%d",
		allowWins,
		blockWins,
		allowOnlyAdded,
		equalCounts,
		len(conflicts),
	)

	return conflicts
}

// filterConflictsAfterOverrides removes manually overridden entries from conflicts list
func filterConflictsAfterOverrides(result *ResolutionResult) []ConflictDetail {
	filteredConflicts := make([]ConflictDetail, 0)

	for _, conflict := range result.Conflicts {
		// Skip if this entry was manually overridden
		_, isAllowToBlock := result.ManualOverride.AllowToBlock[conflict.Entry]
		_, isBlockToAllow := result.ManualOverride.BlockToAllow[conflict.Entry]

		if !isAllowToBlock && !isBlockToAllow {
			filteredConflicts = append(filteredConflicts, conflict)
		}
	}

	return filteredConflicts
}

// applyManualOverrides applies manual overrides
func applyManualOverrides(logger *multilog.Logger, maps *SourceMaps, result *ResolutionResult) {
	manualAllowToBlockByType, manualBlockToAllowByType := readCustomOverrides(logger)

	result.ManualOverride.AllowToBlock = make(map[string]struct{})
	result.ManualOverride.BlockToAllow = make(map[string]struct{})

	for gst, entries := range manualAllowToBlockByType {
		for entry := range entries {
			if hasEntryType(maps.EntryTypes[entry], gst) {
				moveToBlockSet(result, entry, gst)
				result.ManualOverride.AllowToBlock[entry] = struct{}{}
			}
		}
	}

	// Apply forced allows (takes precedence)
	for gst, entries := range manualBlockToAllowByType {
		for entry := range entries {
			if hasEntryType(maps.EntryTypes[entry], gst) {
				moveToAllowSet(result, entry, gst)
				result.ManualOverride.BlockToAllow[entry] = struct{}{}
				// Remove from forced block if it exists
				delete(result.ManualOverride.AllowToBlock, entry)
			}
		}
	}
}

// writeResolvedLists writes the final allow and block lists
func writeResolvedLists(logger *multilog.Logger, result *ResolutionResult) (string, string, error) {
	if !emitResolvedLists {
		return "", "", nil
	}

	outDir := constants.OutputDir
	if err := u.EnsureDirectoryExists(logger, outDir); err != nil {
		return "", "", fmt.Errorf("failed to ensure output dir: %w", err)
	}

	allowEntries := collectAllEntries(result.AllowByType)
	allowPath := filepath.Join(outDir, "consolidated_allowlist.txt")
	resolvedBeforeAllow := len(allowEntries)
	logger.Infof(
		"Writing resolved allowlist %s: original=%d resolved_before=%d must_consider_merged=%d resolved_after=%d",
		allowPath,
		-1,
		resolvedBeforeAllow,
		0,
		resolvedBeforeAllow,
	)
	if err := u.WriteEntriesToFile(logger, allowPath, allowEntries); err != nil {
		return "", "", fmt.Errorf("failed to write allowlist: %w", err)
	}

	blockEntries := collectAllEntries(result.BlockByType)
	blockPath := filepath.Join(outDir, "consolidated_blocklist.txt")
	resolvedBeforeBlock := len(blockEntries)
	logger.Infof(
		"Writing resolved blocklist %s: original=%d resolved_before=%d must_consider_merged=%d resolved_after=%d",
		blockPath,
		-1,
		resolvedBeforeBlock,
		0,
		resolvedBeforeBlock,
	)
	if err := u.WriteEntriesToFile(logger, blockPath, blockEntries); err != nil {
		return allowPath, "", fmt.Errorf("failed to write blocklist: %w", err)
	}

	return allowPath, blockPath, nil
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

func readFileEntries(logger *multilog.Logger, filepath string) ([]string, error) {
	entries, _, err := u.ReadEntriesFromFile(logger, filepath)
	return entries, err
}

func processFileEntries(maps *SourceMaps, entries []string, pf c.ProcessedFile) {
	for _, entry := range entries {
		switch pf.ListType {
		case constants.ListTypeBlocklist:
			addToSourceMap(maps.BlockMap, entry, pf.Name)
		case constants.ListTypeAllowlist:
			addToSourceMap(maps.AllowMap, entry, pf.Name)
		}

		addToSourceMap(maps.EntryTypes, entry, pf.GenericSourceType)
	}
}

func addToSourceMap(sourceMap map[string]map[string]struct{}, entry, source string) {
	if sourceMap[entry] == nil {
		sourceMap[entry] = make(map[string]struct{})
	}
	sourceMap[entry][source] = struct{}{}
}

func getAllUniqueEntries(blockMap, allowMap map[string]map[string]struct{}) map[string]struct{} {
	entries := make(map[string]struct{})
	for entry := range blockMap {
		entries[entry] = struct{}{}
	}
	for entry := range allowMap {
		entries[entry] = struct{}{}
	}
	return entries
}

func getSourcesList(sources map[string]struct{}) []string {
	if sources == nil {
		return nil
	}
	list := make([]string, 0, len(sources))
	for source := range sources {
		list = append(list, source)
	}
	sort.Strings(list)
	return list
}

func addToBlockSets(result *ResolutionResult, entry string, types map[string]struct{}) {
	entryTypes := getEntryTypesOrDefault(types)
	for entryType := range entryTypes {
		ensureStringSet(result.BlockByType, entryType)
		result.BlockByType[entryType].Add(entry)
	}
}

func addToAllowSets(result *ResolutionResult, entry string, types map[string]struct{}) {
	entryTypes := getEntryTypesOrDefault(types)
	for entryType := range entryTypes {
		ensureStringSet(result.AllowByType, entryType)
		result.AllowByType[entryType].Add(entry)
	}
}

func getEntryTypesOrDefault(types map[string]struct{}) map[string]struct{} {
	if len(types) > 0 {
		return types
	}
	return map[string]struct{}{constants.SourceTypeDomain: {}}
}

func ensureStringSet(setMap map[string]u.StringSet, key string) {
	if setMap[key] == nil {
		setMap[key] = u.NewStringSet([]string{})
	}
}

func hasEntryType(types map[string]struct{}, targetType string) bool {
	if types == nil {
		return false
	}
	_, exists := types[targetType]
	return exists
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

func countSetEntries(setsByType map[string]u.StringSet) int {
	total := 0
	for _, set := range setsByType {
		if set != nil {
			total += set.Size()
		}
	}
	return total
}

func isValidProcessedFile(pf c.ProcessedFile) bool {
	return pf.Filepath != "" && pf.Valid
}

func collectAllEntries(setsByType map[string]u.StringSet) []string {
	seen := make(map[string]struct{})
	for _, set := range setsByType {
		for entry := range set {
			seen[entry] = struct{}{}
		}
	}

	entries := make([]string, 0, len(seen))
	for entry := range seen {
		entries = append(entries, entry)
	}
	sort.Strings(entries)
	return entries
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

func fillDetailsForResolution(maps *SourceMaps, result *ResolutionResult) {
	resolvedEntries := make(map[string]struct{})

	for _, set := range result.AllowByType {
		for entry := range set {
			resolvedEntries[entry] = struct{}{}
		}
	}

	for _, set := range result.BlockByType {
		for entry := range set {
			resolvedEntries[entry] = struct{}{}
		}
	}

	for entry := range resolvedEntries {
		if _, exists := result.DetailsMap[entry]; !exists {
			blockSources := getSourcesList(maps.BlockMap[entry])
			allowSources := getSourcesList(maps.AllowMap[entry])

			result.DetailsMap[entry] = ConflictDetail{
				Entry:        entry,
				BlockSources: blockSources,
				AllowSources: allowSources,
				BlockCount:   len(blockSources),
				AllowCount:   len(allowSources),
			}
		}
	}
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
