package cmd

import (
	"github.com/spf13/cobra"
)

var consolidateCategoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "Generate category-based consolidated lists",
	Run: func(cmd *cobra.Command, args []string) {
		runScopedConsolidationCommand(
			"category",
			"Category",
			"Generating category-based consolidated lists...",
			"Categories consolidation complete",
		)
	},
}
