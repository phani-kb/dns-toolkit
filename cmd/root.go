package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/spf13/cobra"
)

var (
	startTime time.Time
	endTime   time.Time
)

// create a map of cobra commands to be ignored for completion
var ignoreCmds = map[string]bool{
	"help":    true,
	"sts":     true,
	"version": true,
}

func validateAndSetDirs() {
	if AppConfig == nil {
		return
	}

	for key, defaultDir := range constants.Folders {
		var dir string
		switch key {
		case "download":
			dir = AppConfig.DNSToolkit.Folders.Download
		}

		if dir == "" {
			dir = defaultDir
		}

		switch key {
		case "download":
			constants.DownloadDir = dir
		}
	}
}

var rootCmd = &cobra.Command{
	Use:   "dns-toolkit",
	Short: "DNS Toolkit",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if ignoreCmds[cmd.Name()] {
			return
		}
		startTime = time.Now()
		Logger.Infof("Command %s started", cmd.Name())
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if ignoreCmds[cmd.Name()] {
			return
		}
		endTime = time.Now()
		duration := endTime.Sub(startTime)
		Logger.Infof("Command %s completed in %s", cmd.Name(), duration)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	configPath := filepath.Join("configs", "config.yml")
	Logger = common.InitLogger(configPath)

	// Skip validation for the help command
	if len(os.Args) <= 1 || (os.Args[1] != "help" && os.Args[1] != "--help" && os.Args[1] != "-h") {
		if err := validateConfig(); err != nil {
			slog.Error("Config validation failed", "error", err)
			os.Exit(1)
		}
	}

	cobra.OnInitialize(validateAndSetDirs)

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(processCmd)
}
