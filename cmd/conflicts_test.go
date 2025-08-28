package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
)

func setupTestEnvironment(t *testing.T) (func(), string) {
	projectRoot, err := utils.FindProjectRoot("")
	if err != nil {
		t.Fatalf("failed to find project root: %v", err)
	}

	testDataDir := filepath.Join(projectRoot, "testdata")
	oldOutputDir := constants.OutputDir
	oldSummaryDir := constants.SummaryDir
	constants.OutputDir = filepath.Join(testDataDir, "output")
	constants.SummaryDir = filepath.Join(testDataDir, "data")

	return func() {
		constants.OutputDir = oldOutputDir
		constants.SummaryDir = oldSummaryDir
	}, testDataDir
}

func TestGenerateConflictReport_NoConflicts(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)

	cleanup, testDataDir := setupTestEnvironment(t)
	defer cleanup()

	bl := filepath.Join(testDataDir, "bl.txt")
	if err := os.WriteFile(bl, []byte("bad.example.com\n"), 0o644); err != nil {
		t.Fatalf("failed to write blocklist: %v", err)
	}

	al := filepath.Join(testDataDir, "al.txt")
	if err := os.WriteFile(al, []byte("good.example.com\n"), 0o644); err != nil {
		t.Fatalf("failed to write allowlist: %v", err)
	}

	path, err := GenerateConflictReport(logger, "non_existent_summary.json")
	assert.NoError(t, err)
	assert.Equal(t, "", path, "expected no report path when no conflicts")
}

func TestGenerateConflictReport_WithConflicts(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)

	cleanup, testDataDir := setupTestEnvironment(t)
	defer cleanup()

	bl := filepath.Join(testDataDir, "bl2.txt")
	if err := os.WriteFile(bl, []byte("a.example.com\ncommon.example.com\n"), 0o644); err != nil {
		t.Fatalf("failed to write blocklist: %v", err)
	}

	al := filepath.Join(testDataDir, "al2.txt")
	if err := os.WriteFile(al, []byte("common.example.com\n"), 0o644); err != nil {
		t.Fatalf("failed to write allowlist: %v", err)
	}

	// processed := []c.ProcessedFile{
	// 	{Name: "bl-src-1", ListType: constants.ListTypeBlocklist, Filepath: bl, Valid: true},
	// 	{Name: "al-src-1", ListType: constants.ListTypeAllowlist, Filepath: al, Valid: true},
	// }

	dataDir := filepath.Join(testDataDir, "data")
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		t.Fatalf("failed to create data dir: %v", err)
	}

	overrides := []struct {
		Entry      string   `json:"entry"`
		Decision   string   `json:"decision"`
		BlockSrcs  []string `json:"block_sources"`
		AllowSrcs  []string `json:"allow_sources"`
		BlockCount int      `json:"block_count"`
		AllowCount int      `json:"allow_count"`
	}{
		{
			Entry:      "common.example.com",
			Decision:   "conflict",
			BlockSrcs:  []string{"bl-src-1"},
			AllowSrcs:  []string{"al-src-1"},
			BlockCount: 1,
			AllowCount: 1,
		},
	}
	confPath := filepath.Join(dataDir, constants.SummaryTypesOutputSummaryFileMap[constants.SummaryTypeOverrides])
	bb, err := json.MarshalIndent(overrides, "", "  ")
	if err != nil {
		t.Fatalf("failed to marshal overrides json: %v", err)
	}
	if err := os.WriteFile(confPath, bb, 0o644); err != nil {
		t.Fatalf("failed to write overrides json: %v", err)
	}

	path, err := GenerateConflictReport(logger, confPath)
	assert.NoError(t, err)
	if assert.NotEmpty(t, path) {
		_, err := os.Stat(path)
		assert.NoError(t, err)
		content, err := os.ReadFile(path)
		assert.NoError(t, err)
		assert.Contains(t, string(content), "common.example.com")
	}
}
