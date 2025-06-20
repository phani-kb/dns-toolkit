package cmd

import (
	"os"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/spf13/cobra"
)

var consolidateCategoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "Generate category-based consolidated lists (advertising, malware, privacy, etc)",
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Infof("Generating category-based consolidated lists...")

		if err := u.EnsureDirectoryExists(Logger, constants.ConsolidatedCategoriesDir); err != nil {
			Logger.Errorf("Failed to create consolidated categories directory: %v", err)
			os.Exit(1)
		}
		if err := u.EnsureDirectoryExists(Logger, constants.SummaryDir); err != nil {
			Logger.Errorf("Failed to create summary directory: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	consolidateCmd.AddCommand(consolidateCategoriesCmd)
}
