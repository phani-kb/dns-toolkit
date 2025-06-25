package cmd

import (
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

// Test command argument handling and flag parsing
func TestCommandArguments(t *testing.T) {
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

	if searchCmd.Run != nil {
		searchCmd.Run(searchCmd, []string{"test.com"})
		searchCmd.Run(searchCmd, []string{"192.168.1.1"})
	}

	if consolidateCmd.HasSubCommands() {
		subcommands := consolidateCmd.Commands()
		for _, subcmd := range subcommands {
			if subcmd.Run != nil {
				subcmd.Run(subcmd, []string{})
			}
		}
	}

	// Test generate command with subcommands
	if generateCmd.HasSubCommands() {
		subcommands := generateCmd.Commands()
		for _, subcmd := range subcommands {
			if subcmd.Run != nil {
				subcmd.Run(subcmd, []string{})
			}
		}
	}
}

// Test command help functionality
func TestCommandHelp(t *testing.T) {
	commands := []*cobra.Command{
		versionCmd,
		downloadCmd,
		processCmd,
		consolidateCmd,
		sourceTypesSummaryCmd,
		validateSourcesCmd,
		overlapCmd,
		topEntriesCmd,
		searchCmd,
		archiveCmd,
		generateCmd,
	}

	for _, cmd := range commands {
		usage := cmd.UsageString()
		assert.NotEmpty(t, usage, "Command %s should have usage string", cmd.Use)

		assert.NotEmpty(t, cmd.Short, "Command %s should have short description", cmd.Use)
	}
}

// Test flag handling
func TestCommandFlags(t *testing.T) {
	rootFlags := rootCmd.Flags()
	assert.NotNil(t, rootFlags)

	commands := []*cobra.Command{
		downloadCmd,
		processCmd,
		consolidateCmd,
		overlapCmd,
		topEntriesCmd,
		searchCmd,
		archiveCmd,
		generateCmd,
	}

	for _, cmd := range commands {
		flags := cmd.Flags()
		assert.NotNil(t, flags, "Command %s should have flags", cmd.Use)
	}
}
