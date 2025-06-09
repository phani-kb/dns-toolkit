package cmd

import (
	"runtime"

	"github.com/spf13/cobra"
)

var overlapMaxWorkers int

var overlapCmd = &cobra.Command{
	Use:   "overlap",
	Short: "Find overlap between source files",
	Run: func(cmd *cobra.Command, args []string) {

		overlapMaxWorkers = AppConfig.DNSToolkit.MaxWorkers
		if overlapMaxWorkers <= 0 {
			overlapMaxWorkers = runtime.GOMAXPROCS(0)
		}
		Logger.Infof("Using %d worker(s) for overlap entry processing", overlapMaxWorkers)
	},
}
