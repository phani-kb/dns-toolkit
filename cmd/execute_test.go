package cmd

import (
	"testing"

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

	// Test archiveCmd
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
