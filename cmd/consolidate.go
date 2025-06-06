package cmd

import (
	"os"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/spf13/cobra"
)

var ignoreAllowlist bool
var includeInvalid bool

var consolidateCmd = &cobra.Command{
	Use:   "consolidate",
	Short: "Consolidate processed files",
	Run: func(cmd *cobra.Command, args []string) {
		consolidateAllCmd.Run(cmd, args)
	},
}

var consolidateAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Consolidate all processed files",
	Run: func(cmd *cobra.Command, args []string) {
		if err := u.EnsureDirectoryExists(Logger, constants.ConsolidatedDir); err != nil {
			Logger.Errorf("Failed to create consolidated directory: %v", err)
			os.Exit(1)
		}
		if err := u.EnsureDirectoryExists(Logger, constants.SummaryDir); err != nil {
			Logger.Errorf("Failed to create summary directory: %v", err)
			os.Exit(1)
		}

		// First phase: Process all allowlists synchronously

		// Second phase: Process all blocklists in parallel, now that we have all allowlist entries
		Logger.Infof("Processing blocklists...")
	},
}

func init() {
	consolidateCmd.PersistentFlags().
		BoolVar(&ignoreAllowlist, "ignore-allowlist", false, "Ignore allowlist during consolidation where applicable")
	consolidateCmd.PersistentFlags().
		BoolVar(&includeInvalid, "include-invalid", false, "Include invalid entry(s) during consolidation")
}
