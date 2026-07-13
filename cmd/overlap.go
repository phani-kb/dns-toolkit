package cmd

import (
	"github.com/spf13/cobra"
)

var overlapCmd = &cobra.Command{
	Use:   "overlap",
	Short: "Find overlap between source files",
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Infof("TODO DB-based processing")
	},
}
