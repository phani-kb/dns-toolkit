package cmd

import (
	"os"
	"path/filepath"
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/utils"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildResolutionSets_SimpleCounts(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)

	cleanup, testDataDir := setupTestEnvironmentForCmdTests(t)
	defer cleanup()

	bl := filepath.Join(testDataDir, "bl_test.txt")
	if err := os.WriteFile(bl, []byte("bad.example.com\n"), 0o644); err != nil {
		t.Fatalf("failed to write blocklist: %v", err)
	}

	al := filepath.Join(testDataDir, "al_test.txt")
	if err := os.WriteFile(al, []byte("good.example.com\n"), 0o644); err != nil {
		t.Fatalf("failed to write allowlist: %v", err)
	}

	processed := []c.ProcessedFile{
		{
			Name:              "bl-src",
			GenericSourceType: constants.SourceTypeDomain,
			ListType:          constants.ListTypeBlocklist,
			Filepath:          bl,
			Valid:             true,
		},
		{
			Name:              "al-src",
			GenericSourceType: constants.SourceTypeDomain,
			ListType:          constants.ListTypeAllowlist,
			Filepath:          al,
			Valid:             true,
		},
	}

	allowByType, blockByType, conflicts, _, _, _ := BuildResolutionSets(logger, processed)

	assert.NotNil(t, allowByType)
	assert.NotNil(t, blockByType)
	assert.NotNil(t, conflicts)

	assert.Contains(t, blockByType[constants.SourceTypeDomain].ToSlice(), "bad.example.com")
	assert.NotContains(t, allowByType[constants.SourceTypeDomain].ToSlice(), "bad.example.com")
}

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

func TestHelperFunctions(t *testing.T) {
	valid := isValidProcessedFile(c.ProcessedFile{Filepath: "/path/file.txt", Valid: true})
	invalid := isValidProcessedFile(c.ProcessedFile{Filepath: "", Valid: false})
	assert.True(t, valid)
	assert.False(t, invalid)

	sources := map[string]struct{}{"source_c": {}, "source_a": {}, "source_b": {}}
	result := getSourcesList(sources)
	assert.Equal(t, []string{"source_a", "source_b", "source_c"}, result)

	detail := ConflictDetail{BlockSources: []string{"src1"}, AllowSources: []string{"src2"}}
	assert.True(t, hasSourcesOnBothSides(detail))

	emptyDetail := ConflictDetail{BlockSources: []string{}, AllowSources: []string{}}
	assert.False(t, hasSourcesOnBothSides(emptyDetail))
}

func TestBuildSourceMaps_Integration(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	cleanup, testDataDir := setupTestEnvironmentForCmdTests(t)
	defer cleanup()

	blockFile := filepath.Join(testDataDir, "block_test.txt")
	allowFile := filepath.Join(testDataDir, "allow_test.txt")

	err := os.WriteFile(blockFile, []byte("bad.example.com\nevil.com\n"), 0o644)
	require.NoError(t, err)

	err = os.WriteFile(allowFile, []byte("good.example.com\nsafe.com\n"), 0o644)
	require.NoError(t, err)

	processedFiles := []c.ProcessedFile{
		{
			Filepath:          blockFile,
			Name:              "TestBlocklist",
			ListType:          constants.ListTypeBlocklist,
			GenericSourceType: constants.SourceTypeDomain,
			Valid:             true,
		},
		{
			Filepath:          allowFile,
			Name:              "TestAllowlist",
			ListType:          constants.ListTypeAllowlist,
			GenericSourceType: constants.SourceTypeDomain,
			Valid:             true,
		},
	}

	maps := buildSourceMaps(logger, processedFiles)

	assert.Contains(t, maps.BlockMap, "bad.example.com")
	assert.Contains(t, maps.BlockMap["bad.example.com"], "TestBlocklist")

	assert.Contains(t, maps.AllowMap, "good.example.com")
	assert.Contains(t, maps.AllowMap["good.example.com"], "TestAllowlist")

	assert.Contains(t, maps.EntryTypes["bad.example.com"], constants.SourceTypeDomain)
}

func TestResolveByCounts_Integration(t *testing.T) {
	maps := &SourceMaps{
		BlockMap: map[string]map[string]struct{}{
			"conflict.com": {"block1": {}, "block2": {}},
			"block.com":    {"block1": {}},
			"equal.com":    {"block1": {}},
		},
		AllowMap: map[string]map[string]struct{}{
			"conflict.com": {"allow1": {}},
			"allow.com":    {"allow1": {}, "allow2": {}},
			"equal.com":    {"allow1": {}},
		},
		EntryTypes: map[string]map[string]struct{}{
			"conflict.com": {constants.SourceTypeDomain: {}},
			"block.com":    {constants.SourceTypeDomain: {}},
			"allow.com":    {constants.SourceTypeDomain: {}},
			"equal.com":    {constants.SourceTypeDomain: {}},
		},
	}

	result := &ResolutionResult{
		AllowByType: make(map[string]u.StringSet),
		BlockByType: make(map[string]u.StringSet),
		DetailsMap:  make(map[string]ConflictDetail),
	}

	conflicts := resolveByCounts(maps, result)

	assert.Contains(t, result.BlockByType[constants.SourceTypeDomain], "conflict.com")
	assert.NotContains(t, result.AllowByType[constants.SourceTypeDomain], "conflict.com")

	assert.Contains(t, result.AllowByType[constants.SourceTypeDomain], "allow.com")

	assert.Contains(t, result.BlockByType[constants.SourceTypeDomain], "block.com")

	assert.Len(t, conflicts, 1)
	assert.Equal(t, "equal.com", conflicts[0].Entry)
	assert.Equal(t, 1, conflicts[0].BlockCount)
	assert.Equal(t, 1, conflicts[0].AllowCount)
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

func TestUncoveredFunctions(t *testing.T) {
	result := &ResolutionResult{
		AllowByType: make(map[string]u.StringSet),
		BlockByType: make(map[string]u.StringSet),
	}
	result.BlockByType[constants.SourceTypeDomain] = u.NewStringSet([]string{"test.com"})

	moveToAllowSet(result, "test.com", constants.SourceTypeDomain)
	assert.Contains(t, result.AllowByType[constants.SourceTypeDomain], "test.com")
	assert.NotContains(t, result.BlockByType[constants.SourceTypeDomain], "test.com")

	setsByType := map[string]u.StringSet{
		"type1": u.NewStringSet([]string{"entry1", "entry2"}),
		"type2": u.NewStringSet([]string{"entry3"}),
	}
	count := countSetEntries(setsByType)
	assert.Equal(t, 3, count)
}

func TestWriteResolvedLists(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	cleanup, _ := setupTestEnvironmentForCmdTests(t)
	defer cleanup()

	result := &ResolutionResult{
		AllowByType: map[string]u.StringSet{
			constants.SourceTypeDomain: u.NewStringSet([]string{"allow.com"}),
		},
		BlockByType: map[string]u.StringSet{
			constants.SourceTypeDomain: u.NewStringSet([]string{"block.com"}),
		},
	}

	oldEmitResolvedLists := emitResolvedLists
	emitResolvedLists = false
	allowPath, blockPath, err := writeResolvedLists(logger, result)
	assert.NoError(t, err)
	assert.Empty(t, allowPath)
	assert.Empty(t, blockPath)

	emitResolvedLists = true
	allowPath, blockPath, err = writeResolvedLists(logger, result)
	assert.NoError(t, err)
	assert.NotEmpty(t, allowPath)
	assert.NotEmpty(t, blockPath)

	assert.FileExists(t, allowPath)
	assert.FileExists(t, blockPath)

	emitResolvedLists = oldEmitResolvedLists
}

func TestResolveConflictsByCounts_Integration(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	cleanup, testDataDir := setupTestEnvironmentForCmdTests(t)
	defer cleanup()

	summaryDir := filepath.Join(testDataDir, "summary")
	err := os.MkdirAll(summaryDir, 0o755)
	require.NoError(t, err)

	blockFile := filepath.Join(testDataDir, "block_resolve_test.txt")
	allowFile := filepath.Join(testDataDir, "allow_resolve_test.txt")

	err = os.WriteFile(blockFile, []byte("conflict.example.com\nblock-only.com\n"), 0o644)
	require.NoError(t, err)

	err = os.WriteFile(allowFile, []byte("conflict.example.com\nallow-only.com\n"), 0o644)
	require.NoError(t, err)

	processedFiles := []c.ProcessedFile{
		{
			Filepath:          blockFile,
			Name:              "TestBlockSource",
			ListType:          constants.ListTypeBlocklist,
			GenericSourceType: constants.SourceTypeDomain,
			Valid:             true,
		},
		{
			Filepath:          allowFile,
			Name:              "TestAllowSource",
			ListType:          constants.ListTypeAllowlist,
			GenericSourceType: constants.SourceTypeDomain,
			Valid:             true,
		},
	}

	allowPath, blockPath, overridePath, err := ResolveConflictsByCounts(logger, processedFiles)
	assert.NoError(t, err)
	assert.NotEmpty(t, overridePath) // Should create override summary

	if !emitResolvedLists {
		assert.Empty(t, allowPath)
		assert.Empty(t, blockPath)
	}
}

func TestPerTypeOverride_AppliesOnlyToMatchingType(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)

	cleanup, testDataDir := setupTestEnvironmentForCmdTests(t)
	defer cleanup()

	customDir := filepath.Join(testDataDir, "data", "custom")
	if err := os.MkdirAll(customDir, 0o755); err != nil {
		t.Fatalf("failed to create custom dir: %v", err)
	}

	domainFile := filepath.Join(testDataDir, "domain_pf.txt")
	if err := os.WriteFile(domainFile, []byte("bad.example.com\n"), 0o644); err != nil {
		t.Fatalf("failed to write domain processed file: %v", err)
	}

	domainForcedBlock := filepath.Join(customDir, "domain_forced_block.txt")
	if err := os.WriteFile(domainForcedBlock, []byte("bad.example.com\n"), 0o644); err != nil {
		t.Fatalf("failed to write domain override file: %v", err)
	}

	oldMap := constants.CustomOverrideFilesMap
	defer func() { constants.CustomOverrideFilesMap = oldMap }()
	constants.CustomOverrideFilesMap = map[string]map[string]string{
		constants.SourceTypeDomain: {
			constants.ForcedBlock: domainForcedBlock,
			constants.ForcedAllow: "",
		},
	}

	processed := []c.ProcessedFile{
		{
			Name:              "domain-src",
			GenericSourceType: constants.SourceTypeDomain,
			ListType:          constants.ListTypeBlocklist,
			Filepath:          domainFile,
			Valid:             true,
		},
	}

	allowByType, blockByType, _, manualAllowToBlock, _, _ := BuildResolutionSets(logger, processed)
	_, applied := manualAllowToBlock["bad.example.com"]
	assert.True(t, applied, "expected manualAllowToBlock to contain bad.example.com")
	assert.Contains(t, blockByType[constants.SourceTypeDomain].ToSlice(), "bad.example.com")
	assert.NotContains(t, allowByType[constants.SourceTypeDomain].ToSlice(), "bad.example.com")
}

func TestPerTypeOverride_DoesNotApplyToOtherType(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)

	cleanup, testDataDir := setupTestEnvironmentForCmdTests(t)
	defer cleanup()

	customDir := filepath.Join(testDataDir, "data", "custom")
	if err := os.MkdirAll(customDir, 0o755); err != nil {
		t.Fatalf("failed to create custom dir: %v", err)
	}

	domainFile := filepath.Join(testDataDir, "domain_pf2.txt")
	if err := os.WriteFile(domainFile, []byte("bad.example.com\n"), 0o644); err != nil {
		t.Fatalf("failed to write domain processed file: %v", err)
	}

	ipv4ForcedBlock := filepath.Join(customDir, "ipv4_forced_block.txt")
	if err := os.WriteFile(ipv4ForcedBlock, []byte("bad.example.com\n"), 0o644); err != nil {
		t.Fatalf("failed to write ipv4 override file: %v", err)
	}

	oldMap := constants.CustomOverrideFilesMap
	defer func() { constants.CustomOverrideFilesMap = oldMap }()
	constants.CustomOverrideFilesMap = map[string]map[string]string{
		constants.SourceTypeIpv4: {
			constants.ForcedBlock: ipv4ForcedBlock,
			constants.ForcedAllow: "",
		},
	}

	processed := []c.ProcessedFile{
		{
			Name:              "domain-src",
			GenericSourceType: constants.SourceTypeDomain,
			ListType:          constants.ListTypeBlocklist,
			Filepath:          domainFile,
			Valid:             true,
		},
	}

	_, _, _, manualAllowToBlock, _, _ := BuildResolutionSets(logger, processed)
	_, applied := manualAllowToBlock["bad.example.com"]
	assert.False(t, applied, "expected manualAllowToBlock NOT to contain bad.example.com for ipv4 override")
}
