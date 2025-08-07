package cmd

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/multilog"
)

var (
	Logger         *multilog.Logger
	AppConfig      *config.AppConfig
	SourcesConfigs []config.SourcesConfig
)

// GetConfigPath returns the path to the configuration file.
func GetConfigPath() (string, error) {
	configPath := filepath.Join("configs", "config.yml")
	inTestMode := os.Getenv("DNS_TOOLKIT_TEST_MODE") == "true"
	if inTestMode {
		configPath = os.Getenv("DNS_TOOLKIT_TEST_CONFIG_PATH")
		if configPath == "" {
			return "", errors.New("DNS_TOOLKIT_TEST_CONFIG_PATH is not set")
		}
	}
	return configPath, nil
}
