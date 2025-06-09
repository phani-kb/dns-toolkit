package cmd

import (
	"runtime"

	"github.com/spf13/cobra"
)

var minSources int
var maxEntries int
var maxWorkers int

var topEntriesCmd = &cobra.Command{
	Use:   "top",
	Short: "Find top entry(s) in each generic source type",
	Run: func(cmd *cobra.Command, args []string) {
		maxWorkers = AppConfig.DNSToolkit.MaxWorkers
		if maxWorkers <= 0 {
			maxWorkers = runtime.GOMAXPROCS(0)
		}
		Logger.Infof("Using %d worker(s) for top entry(s) processing", maxWorkers)

	},
}

func init() {
	topEntriesCmd.Flags().IntVarP(&minSources, "min-sources", "m", 0, "Minimum sources (3-12)")
	topEntriesCmd.Flags().IntVarP(&maxEntries, "max-entries", "x", int(^uint(0)>>1), "Max entries")
}
