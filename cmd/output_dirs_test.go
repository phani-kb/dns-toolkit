package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/stretchr/testify/assert"
)

// TestOutputDirectoriesSetCorrectly tests that computed output directories are set correctly
func TestOutputDirectoriesSetCorrectly(t *testing.T) {
	oldConfig := AppConfig
	defer func() { AppConfig = oldConfig }()

	AppConfig = &config.AppConfig{
		DNSToolkit: config.DNSToolkitConfig{
			Folders: config.FoldersConfig{
				Output: "testdata/output",
				Backup: "testdata/backup",
			},
		},
	}

	err := os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	assert.NoError(t, err)
	defer func() {
		err := os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
		if err != nil {
			t.Logf("Failed to unset DNS_TOOLKIT_TEST_MODE: %v", err)
		}
	}()

	InitForTesting()

	assert.True(t, filepath.IsAbs(constants.OutputDir), "OutputDir should be absolute in test mode")
	assert.True(
		t,
		strings.HasSuffix(constants.OutputDir, "testdata/output"),
		"OutputDir should end with testdata/output",
	)

	assert.True(
		t,
		filepath.IsAbs(constants.OutputSummariesDir),
		"OutputSummariesDir should be absolute in test mode, got: %s",
		constants.OutputSummariesDir,
	)
	assert.True(
		t,
		strings.HasSuffix(constants.OutputSummariesDir, "testdata/output/summaries"),
		"OutputSummariesDir should end with testdata/output/summaries, got: %s",
		constants.OutputSummariesDir,
	)

	assert.True(
		t,
		filepath.IsAbs(constants.OutputGroupsDir),
		"OutputGroupsDir should be absolute in test mode, got: %s",
		constants.OutputGroupsDir,
	)
	assert.True(
		t,
		strings.HasSuffix(constants.OutputGroupsDir, "testdata/output/groups"),
		"OutputGroupsDir should end with testdata/output/groups, got: %s",
		constants.OutputGroupsDir,
	)

	assert.True(
		t,
		filepath.IsAbs(constants.OutputCategoriesDir),
		"OutputCategoriesDir should be absolute in test mode, got: %s",
		constants.OutputCategoriesDir,
	)
	assert.True(
		t,
		strings.HasSuffix(constants.OutputCategoriesDir, "testdata/output/categories"),
		"OutputCategoriesDir should end with testdata/output/categories, got: %s",
		constants.OutputCategoriesDir,
	)

	assert.True(t, filepath.IsAbs(constants.BackupDir), "BackupDir should be absolute in test mode")
	assert.True(
		t,
		strings.HasSuffix(constants.BackupDir, "testdata/backup"),
		"BackupDir should end with testdata/backup",
	)
}
