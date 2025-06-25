package cmd

import (
	"fmt"
	"log/slog"

	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/spf13/cobra"
)

var validationPerformed bool

func validateConfig(configPath string) error {
	if validationPerformed {
		slog.Debug("Skipping validation as it has already been performed")
		return nil
	}

	appConfig, sourcesConfigs, err := config.LoadAppConfig(Logger, configPath)
	if err != nil {
		return fmt.Errorf("config validation error: %w", err)
	}

	AppConfig = &appConfig
	SourcesConfigs = sourcesConfigs
	validationPerformed = true

	return nil
}

var validateSourcesCmd = &cobra.Command{
	Use:   "validate-sources",
	Short: "Validate the sources configuration",
	Run: func(cmd *cobra.Command, args []string) {
		configPath, err := GetConfigPath()
		if err != nil {
			slog.Error("Failed to get config path", "error", err)
			return
		}
		if err := validateConfig(configPath); err != nil {
			slog.Error("Validation failed", "error", err)
			return
		}
		Logger.Infof("Successfully loaded and validated configuration")
	},
}
