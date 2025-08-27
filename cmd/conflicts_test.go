package cmd

import (
	"os"
	"path/filepath"
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
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
	constants.OutputDir = filepath.Join(testDataDir, constants.OutputDir)

	return func() {
		constants.OutputDir = oldOutputDir
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

	processed := []c.ProcessedFile{
		{Name: "bl-src", ListType: constants.ListTypeBlocklist, Filepath: bl, Valid: true},
		{Name: "al-src", ListType: constants.ListTypeAllowlist, Filepath: al, Valid: true},
	}

	path, err := GenerateConflictReport(logger, processed)
	assert.NoError(t, err)
	assert.Equal(t, "", path, "expected no report path when no conflicts")
}

func TestGenerateConflictReport_WithConflicts(t *testing.T) {
	t.Parallel()
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

	processed := []c.ProcessedFile{
		{Name: "bl-src-1", ListType: constants.ListTypeBlocklist, Filepath: bl, Valid: true},
		{Name: "al-src-1", ListType: constants.ListTypeAllowlist, Filepath: al, Valid: true},
	}

	path, err := GenerateConflictReport(logger, processed)
	assert.NoError(t, err)
	if assert.NotEmpty(t, path) {
		_, err := os.Stat(path)
		assert.NoError(t, err)
		content, err := os.ReadFile(path)
		assert.NoError(t, err)
		assert.Contains(t, string(content), "common.example.com")
	}
}
