package cmd

import (
	"github.com/spf13/cobra"
)

var (
	minSources int
	maxEntries int
)

var topEntriesCmd = &cobra.Command{
	Use:   "top",
	Short: "Find top entry(s) in each generic source type",
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Infof("TODO DB-based processing")
	},
}

func init() {
	topEntriesCmd.Flags().IntVarP(&minSources, "min-sources", "m", 0, "Minimum sources (3-12)")
	topEntriesCmd.Flags().IntVarP(&maxEntries, "max-entries", "x", int(^uint(0)>>1), "Max entries")
}
