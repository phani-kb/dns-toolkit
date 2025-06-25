package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

// Test the Execute function
func TestExecute(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"dns-toolkit", "version"}

	Execute()
}

func TestValidateAndSetDirs(t *testing.T) {
	oldConfig := AppConfig
	defer func() { AppConfig = oldConfig }()

	// Test with nil config
	AppConfig = nil
	validateAndSetDirs()

	AppConfig = &config.AppConfig{
		DNSToolkit: config.DNSToolkitConfig{
			Folders: config.FoldersConfig{
				Download:               "test/download",
				Processed:              "test/processed",
				Consolidated:           "test/consolidated",
				Summary:                "test/summary",
				Overlap:                "test/overlap",
				Top:                    "test/top",
				ConsolidatedGroups:     "test/consolidated_groups",
				ConsolidatedCategories: "test/consolidated_categories",
				Archive:                "test/archive",
				Output:                 "test/output",
				Backup:                 "test/backup",
			},
		},
	}

	os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
	validateAndSetDirs()
	assert.True(t, filepath.IsAbs(constants.DownloadDir))
	assert.True(t, filepath.Base(filepath.Dir(constants.DownloadDir)) == "test" && filepath.Base(constants.DownloadDir) == "download")

	os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	validateAndSetDirs()

	assert.True(t, filepath.IsAbs(constants.DownloadDir))
	assert.Contains(t, constants.DownloadDir, "test/download")

	AppConfig.DNSToolkit.Folders = config.FoldersConfig{}
	validateAndSetDirs()
}

func TestInitForTesting(t *testing.T) {
	oldConfig := AppConfig
	defer func() { AppConfig = oldConfig }()

	AppConfig = &config.AppConfig{
		DNSToolkit: config.DNSToolkitConfig{
			Folders: config.FoldersConfig{
				Summary: "testdata/summary",
			},
		},
	}

	os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	InitForTesting()

	assert.True(t, filepath.IsAbs(constants.SummaryDir))
}

// Test the root command's PersistentPreRun and PersistentPostRun
func TestRootCmdHooks(t *testing.T) {
	cmd := &cobra.Command{Use: "help"}
	var args []string

	rootCmd.PersistentPreRun(cmd, args)

	rootCmd.PersistentPostRun(cmd, args)

	cmd = &cobra.Command{Use: "test-cmd"}

	if rootCmd.PersistentPreRun != nil {
		rootCmd.PersistentPreRun(cmd, args)
	}

	// PersistentPostRun should handle non-ignored commands
	if rootCmd.PersistentPostRun != nil {
		rootCmd.PersistentPostRun(cmd, args)
	}
}

// Test the ignoreCmds map functionality
func TestIgnoreCmds(t *testing.T) {
	assert.True(t, ignoreCmds["help"])
	assert.True(t, ignoreCmds["sts"])
	assert.True(t, ignoreCmds["version"])
	assert.False(t, ignoreCmds["download"])
}

// Test command structure
func TestRootCommandStructure(t *testing.T) {
	assert.Equal(t, "dns-toolkit", rootCmd.Use)
	assert.Equal(t, "DNS Toolkit", rootCmd.Short)
	assert.True(t, rootCmd.CompletionOptions.DisableDefaultCmd)
}

// Test all validateAndSetDirs switch cases
func TestValidateAndSetDirsAllCases(t *testing.T) {
	oldConfig := AppConfig
	defer func() { AppConfig = oldConfig }()

	AppConfig = &config.AppConfig{
		DNSToolkit: config.DNSToolkitConfig{
			Folders: config.FoldersConfig{
				Download:               "custom/download",
				Processed:              "custom/processed",
				Consolidated:           "custom/consolidated",
				Summary:                "custom/summary",
				Overlap:                "custom/overlap",
				Top:                    "custom/top",
				ConsolidatedGroups:     "custom/consolidated_groups",
				ConsolidatedCategories: "custom/consolidated_categories",
				Archive:                "custom/archive",
				Output:                 "custom/output",
				Backup:                 "custom/backup",
			},
		},
	}

	// Test in normal mode (not test mode)
	os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
	validateAndSetDirs()

	// Verify all constants are set as absolute paths with correct suffixes
	assert.True(t, filepath.IsAbs(constants.DownloadDir))
	assert.True(t, filepath.Base(filepath.Dir(constants.DownloadDir)) == "custom" && filepath.Base(constants.DownloadDir) == "download")
	assert.True(t, filepath.IsAbs(constants.ProcessedDir))
	assert.True(t, filepath.Base(filepath.Dir(constants.ProcessedDir)) == "custom" && filepath.Base(constants.ProcessedDir) == "processed")
	assert.True(t, filepath.IsAbs(constants.ConsolidatedDir))
	assert.True(t, filepath.Base(filepath.Dir(constants.ConsolidatedDir)) == "custom" && filepath.Base(constants.ConsolidatedDir) == "consolidated")
	assert.True(t, filepath.IsAbs(constants.SummaryDir))
	assert.True(t, filepath.Base(filepath.Dir(constants.SummaryDir)) == "custom" && filepath.Base(constants.SummaryDir) == "summary")
	assert.True(t, filepath.IsAbs(constants.OverlapDir))
	assert.True(t, filepath.Base(filepath.Dir(constants.OverlapDir)) == "custom" && filepath.Base(constants.OverlapDir) == "overlap")
	assert.True(t, filepath.IsAbs(constants.TopDir))
	assert.True(t, filepath.Base(filepath.Dir(constants.TopDir)) == "custom" && filepath.Base(constants.TopDir) == "top")
	assert.True(t, filepath.IsAbs(constants.ConsolidatedGroupsDir))
	assert.True(t, filepath.Base(filepath.Dir(constants.ConsolidatedGroupsDir)) == "custom" && filepath.Base(constants.ConsolidatedGroupsDir) == "consolidated_groups")
	assert.True(t, filepath.IsAbs(constants.ConsolidatedCategoriesDir))
	assert.True(t, filepath.Base(filepath.Dir(constants.ConsolidatedCategoriesDir)) == "custom" && filepath.Base(constants.ConsolidatedCategoriesDir) == "consolidated_categories")
	assert.True(t, filepath.IsAbs(constants.ArchiveDir))
	assert.True(t, filepath.Base(filepath.Dir(constants.ArchiveDir)) == "custom" && filepath.Base(constants.ArchiveDir) == "archive")
	assert.True(t, filepath.IsAbs(constants.OutputDir))
	assert.True(t, filepath.Base(filepath.Dir(constants.OutputDir)) == "custom" && filepath.Base(constants.OutputDir) == "output")
	assert.True(t, filepath.IsAbs(constants.BackupDir))
	assert.True(t, filepath.Base(filepath.Dir(constants.BackupDir)) == "custom" && filepath.Base(constants.BackupDir) == "backup")
}
