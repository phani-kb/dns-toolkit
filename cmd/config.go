package cmd

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/db"
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
	inTestMode := os.Getenv("DNS_TOOLKIT_TEST_MODE") == constants.BooleanTrue
	if inTestMode {
		configPath = os.Getenv("DNS_TOOLKIT_TEST_CONFIG_PATH")
		if configPath == "" {
			return "", errors.New("DNS_TOOLKIT_TEST_CONFIG_PATH is not set")
		}
	}
	return configPath, nil
}

// openDB opens the database in read/write mode, exiting if it fails.
func openDB(ctx context.Context) *db.DB {
	dbPath := getDBPath()
	database, err := db.Open(ctx, Logger, dbPath, false)
	if err != nil {
		Logger.Errorf("Failed to open database: %v", err)
		os.Exit(1)
	}
	return database
}
