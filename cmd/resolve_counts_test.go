package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/utils"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestEnvironmentForCmdTests(t *testing.T) (func(), string) {
	projectRoot, err := utils.FindProjectRoot("")
	if err != nil {
		t.Fatalf("failed to find project root: %v", err)
	}
	testDataDir := filepath.Join(projectRoot, "testdata")
	oldSummaryDir := constants.SummaryDir
	oldOutputDir := constants.OutputDir
	constants.SummaryDir = filepath.Join(testDataDir, "summary")
	constants.OutputDir = filepath.Join(testDataDir, "output")

	return func() {
		constants.SummaryDir = oldSummaryDir
		constants.OutputDir = oldOutputDir
	}, testDataDir
}

func TestResolutionResult_Struct(t *testing.T) {
	result := &ResolutionResult{
		AllowByType: make(map[string]u.StringSet),
		BlockByType: make(map[string]u.StringSet),
		Conflicts:   []ConflictDetail{},
		DetailsMap:  make(map[string]ConflictDetail),
	}
	result.ManualOverride.AllowToBlock = make(map[string]struct{})
	result.ManualOverride.BlockToAllow = make(map[string]struct{})

	assert.NotNil(t, result.AllowByType)
	assert.NotNil(t, result.BlockByType)
	assert.NotNil(t, result.Conflicts)
	assert.NotNil(t, result.DetailsMap)
	assert.NotNil(t, result.ManualOverride.AllowToBlock)
	assert.NotNil(t, result.ManualOverride.BlockToAllow)
}

func TestConflictProcessing(t *testing.T) {
	conflicts := []ConflictDetail{
		{
			Entry:        "conflict1.com",
			BlockSources: []string{"block1"},
			AllowSources: []string{"allow1"},
			BlockCount:   1,
			AllowCount:   1,
		},
		{Entry: "no-sources.com", BlockSources: []string{}, AllowSources: []string{}, BlockCount: 0, AllowCount: 0},
	}

	records := getConflictRecords(conflicts)
	assert.Len(t, records, 1) // only entries with sources on both sides
	assert.Equal(t, "conflict1.com", records[0].Entry)
	assert.Equal(t, DecisionConflict, records[0].Decision)

	logger, _ := multilog.NewTestLogger(t)
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "manual.txt")

	content := "entry1.com\nentry2.com\n  entry3.com  \n\n"
	err := os.WriteFile(testFile, []byte(content), 0o644)
	require.NoError(t, err)

	result := readManualEntries(logger, "test manual", testFile)
	assert.Len(t, result, 3)
	assert.Contains(t, result, "entry1.com")
	assert.Contains(t, result, "entry3.com") // trimmed

	result = readManualEntries(logger, "nonexistent", filepath.Join(tempDir, "nonexistent.txt"))
	assert.Nil(t, result)
}

func TestMoveToBlockSet(t *testing.T) {
	result := &ResolutionResult{
		AllowByType: make(map[string]u.StringSet),
		BlockByType: make(map[string]u.StringSet),
	}
	result.AllowByType[constants.SourceTypeDomain] = u.NewStringSet([]string{"test.com"})

	moveToBlockSet(result, "test.com", constants.SourceTypeDomain)

	assert.True(t, result.BlockByType[constants.SourceTypeDomain].Contains("test.com"))
	assert.False(t, result.AllowByType[constants.SourceTypeDomain].Contains("test.com"))
}

func TestIsManualOverride(t *testing.T) {
	result := &ResolutionResult{}
	result.ManualOverride.AllowToBlock = map[string]struct{}{"blocked.com": {}}
	result.ManualOverride.BlockToAllow = map[string]struct{}{"allowed.com": {}}

	assert.True(t, isManualOverride("blocked.com", result))
	assert.True(t, isManualOverride("allowed.com", result))
	assert.False(t, isManualOverride("neutral.com", result))
}

func TestGetAutomaticDecisions(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	cleanup, _ := setupTestEnvironmentForCmdTests(t)
	defer cleanup()

	result := &ResolutionResult{
		AllowByType: make(map[string]u.StringSet),
		BlockByType: make(map[string]u.StringSet),
		DetailsMap:  make(map[string]ConflictDetail),
	}
	result.ManualOverride.AllowToBlock = make(map[string]struct{})
	result.ManualOverride.BlockToAllow = make(map[string]struct{})

	// Entry resolved as block (more block sources)
	result.BlockByType[constants.SourceTypeDomain] = u.NewStringSet([]string{"block-wins.com"})
	result.DetailsMap["block-wins.com"] = ConflictDetail{
		Entry:        "block-wins.com",
		BlockSources: []string{"src1", "src2"},
		AllowSources: []string{"src3"},
		BlockCount:   2,
		AllowCount:   1,
	}

	// Entry resolved as allow (more allow sources)
	result.AllowByType[constants.SourceTypeDomain] = u.NewStringSet([]string{"allow-wins.com"})
	result.DetailsMap["allow-wins.com"] = ConflictDetail{
		Entry:        "allow-wins.com",
		BlockSources: []string{"src1"},
		AllowSources: []string{"src2", "src3"},
		BlockCount:   1,
		AllowCount:   2,
	}

	records := getAutomaticDecisions(logger, result)

	assert.Len(t, records, 2)
	decisionMap := make(map[string]string)
	for _, r := range records {
		decisionMap[r.Entry] = r.Decision
	}
	assert.Equal(t, DecisionBlock, decisionMap["block-wins.com"])
	assert.Equal(t, DecisionAllow, decisionMap["allow-wins.com"])
}

func TestGetManualOverrideRecords(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)

	result := &ResolutionResult{
		AllowByType: make(map[string]u.StringSet),
		BlockByType: make(map[string]u.StringSet),
		DetailsMap:  make(map[string]ConflictDetail),
	}
	result.ManualOverride.AllowToBlock = map[string]struct{}{"forced-block.com": {}}
	result.ManualOverride.BlockToAllow = map[string]struct{}{"forced-allow.com": {}}

	// Add details with sources on both sides
	result.DetailsMap["forced-block.com"] = ConflictDetail{
		Entry:        "forced-block.com",
		BlockSources: []string{"src1"},
		AllowSources: []string{"src2"},
		BlockCount:   1,
		AllowCount:   1,
	}
	result.DetailsMap["forced-allow.com"] = ConflictDetail{
		Entry:        "forced-allow.com",
		BlockSources: []string{"src3"},
		AllowSources: []string{"src4"},
		BlockCount:   1,
		AllowCount:   1,
	}

	records := getManualOverrideRecords(logger, result)

	assert.Len(t, records, 2)
	decisionMap := make(map[string]string)
	for _, r := range records {
		decisionMap[r.Entry] = r.Decision
	}
	assert.Equal(t, DecisionBlock, decisionMap["forced-block.com"])
	assert.Equal(t, DecisionAllow, decisionMap["forced-allow.com"])
}

func TestBuildOverrideRecords(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)

	result := &ResolutionResult{
		AllowByType: make(map[string]u.StringSet),
		BlockByType: make(map[string]u.StringSet),
		DetailsMap:  make(map[string]ConflictDetail),
		Conflicts:   []ConflictDetail{},
	}
	result.ManualOverride.AllowToBlock = make(map[string]struct{})
	result.ManualOverride.BlockToAllow = make(map[string]struct{})

	// Add a conflict
	result.Conflicts = append(result.Conflicts, ConflictDetail{
		Entry:        "conflict.com",
		BlockSources: []string{"src1"},
		AllowSources: []string{"src2"},
		BlockCount:   1,
		AllowCount:   1,
	})

	records := buildOverrideRecords(logger, result)

	// Should have conflict record
	assert.Len(t, records, 1)
	assert.Equal(t, "conflict.com", records[0].Entry)
	assert.Equal(t, DecisionConflict, records[0].Decision)
}

func TestWriteOverrideSummary(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	cleanup, _ := setupTestEnvironmentForCmdTests(t)
	defer cleanup()

	result := &ResolutionResult{
		AllowByType: make(map[string]u.StringSet),
		BlockByType: make(map[string]u.StringSet),
		DetailsMap:  make(map[string]ConflictDetail),
		Conflicts:   []ConflictDetail{},
	}
	result.ManualOverride.AllowToBlock = make(map[string]struct{})
	result.ManualOverride.BlockToAllow = make(map[string]struct{})

	path, err := writeOverrideSummary(logger, result)
	assert.NoError(t, err)
	assert.NotEmpty(t, path)
	assert.FileExists(t, path)
}

func TestApplyManualOverrides(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	cleanup, _ := setupTestEnvironmentForCmdTests(t)
	defer cleanup()

	result := &ResolutionResult{
		AllowByType: make(map[string]u.StringSet),
		BlockByType: make(map[string]u.StringSet),
		DetailsMap:  make(map[string]ConflictDetail),
		Conflicts:   []ConflictDetail{},
	}

	// Add an entry in the allow set with details
	result.AllowByType[constants.SourceTypeDomain] = u.NewStringSet([]string{"test.com"})
	result.DetailsMap["test.com"] = ConflictDetail{
		Entry:             "test.com",
		GenericSourceType: constants.SourceTypeDomain,
		BlockSources:      []string{"src1"},
		AllowSources:      []string{"src2"},
	}

	// Apply manual overrides - this will initialize ManualOverride maps
	applyManualOverrides(logger, result)

	// Verify ManualOverride maps are initialized
	assert.NotNil(t, result.ManualOverride.AllowToBlock)
	assert.NotNil(t, result.ManualOverride.BlockToAllow)
}

func TestApplyManualOverrides_EmptyResult(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	cleanup, _ := setupTestEnvironmentForCmdTests(t)
	defer cleanup()

	result := &ResolutionResult{
		AllowByType: make(map[string]u.StringSet),
		BlockByType: make(map[string]u.StringSet),
		DetailsMap:  make(map[string]ConflictDetail),
	}

	// Should not panic with empty result
	applyManualOverrides(logger, result)

	assert.NotNil(t, result.ManualOverride.AllowToBlock)
	assert.NotNil(t, result.ManualOverride.BlockToAllow)
}
