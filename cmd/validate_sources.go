package cmd

import (
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/phani-kb/dns-toolkit/internal/config"
)

var (
	validationPerformed bool
	configPath          = filepath.Join("configs", "config.yml")
)

func validateConfig() error {
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
