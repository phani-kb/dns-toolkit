package cmd

import (
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of DNS Toolkit",
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Infof("%s %s", AppConfig.Application.Name, AppConfig.Application.Version)
	},
}
