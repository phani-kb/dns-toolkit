package cmd

import (
	"os"
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/stretchr/testify/require"

	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/stretchr/testify/assert"
)

// Test execution of commands with proper setup
func TestExecuteCommands(t *testing.T) {
	oldConfig := AppConfig
	defer func() { AppConfig = oldConfig }()

	AppConfig = &config.AppConfig{
		Application: config.ApplicationConfig{
			Name:    "DNS-Toolkit-Test",
			Version: "test-version",
		},
		DNSToolkit: config.DNSToolkitConfig{
			Folders: config.FoldersConfig{
				Download:     "testdata/download",
				Processed:    "testdata/processed",
				Consolidated: "testdata/consolidated",
				Summary:      "testdata/summary",
				Overlap:      "testdata/overlap",
				Top:          "testdata/top",
				Archive:      "testdata/archive",
				Output:       "testdata/output",
				Backup:       "testdata/backup",
			},
		},
	}

	err := os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	require.NoError(t, err)

	wd, err := os.Getwd()
	require.NoError(t, err)

	projectRoot, err := utils.FindProjectRoot(wd)
	assert.NoError(t, err)

	err = os.Setenv("DNS_TOOLKIT_TEST_CONFIG_PATH", projectRoot+"/testdata/config.yml")
	assert.NoError(t, err)

	InitForTesting()

	if versionCmd.Run != nil {
		versionCmd.Run(versionCmd, []string{})
	}

	assert.NotNil(t, downloadCmd.Run)

	if downloadCmd.Run != nil {
		downloadCmd.Run(downloadCmd, []string{})
	}

	if processCmd.Run != nil {
		processCmd.Run(processCmd, []string{})
	}

	if sourceTypesSummaryCmd.Run != nil {
		sourceTypesSummaryCmd.Run(sourceTypesSummaryCmd, []string{})
	}

	if validateSourcesCmd.Run != nil {
		validateSourcesCmd.Run(validateSourcesCmd, []string{})
	}

	if consolidateCmd.Run != nil {
		consolidateCmd.Run(consolidateCmd, []string{})
	}

	if overlapCmd.Run != nil {
		overlapCmd.Run(overlapCmd, []string{})
	}

	if topEntriesCmd.Run != nil {
		topEntriesCmd.Run(topEntriesCmd, []string{})
	}

	if archiveCmd.Run != nil {
		archiveCmd.Run(archiveCmd, []string{})
	}

	if generateCmd.Run != nil {
		generateCmd.Run(generateCmd, []string{})
	}

	if searchCmd.Run != nil {
		searchCmd.Run(searchCmd, []string{"example.com"})
	}
}
