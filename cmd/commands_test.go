package cmd

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

// Test command structure and properties
func TestCommandStructures(t *testing.T) {
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
		assert.NotEmpty(t, cmd.Use, "Command should have Use field")
		assert.NotEmpty(t, cmd.Short, "Command should have Short description")
	}

	assert.Equal(t, "version", versionCmd.Use)
	assert.Equal(t, "download", downloadCmd.Use)
	assert.Equal(t, "process", processCmd.Use)
	assert.Equal(t, "consolidate", consolidateCmd.Use)
	assert.Equal(t, "sts", sourceTypesSummaryCmd.Use)
	assert.Equal(t, "validate-sources", validateSourcesCmd.Use)
	assert.Equal(t, "overlap", overlapCmd.Use)
	assert.Equal(t, "top", topEntriesCmd.Use)
	assert.Contains(t, searchCmd.Use, "search")
	assert.Equal(t, "archive", archiveCmd.Use)
	assert.Equal(t, "generate", generateCmd.Use)
}

// Test that all commands have been added to root
func TestRootCommandHasSubcommands(t *testing.T) {
	subcommands := rootCmd.Commands()

	assert.Greater(t, len(subcommands), 5, "Root command should have multiple subcommands")

	commandNames := make(map[string]bool)
	for _, cmd := range subcommands {
		commandNames[cmd.Use] = true
	}

	assert.True(t, commandNames["version"], "Should have version command")
	assert.True(t, commandNames["download"], "Should have download command")
}
