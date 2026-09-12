package cmd

import (
	"github.com/spf13/cobra"
)

var consolidateGroupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "Generate sized consolidated lists (mini, lite, normal, big)",
	Run: func(cmd *cobra.Command, args []string) {
		runScopedConsolidationCommand(
			"group",
			"Group",
			"Generating sized consolidated lists...",
			"Groups consolidation complete",
		)
	},
}
