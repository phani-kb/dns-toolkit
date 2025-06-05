package common

import (
	"log"
	"log/slog"

	"github.com/phani-kb/multilog"
)

// InitLogger initializes the logger with the specified configuration.
func InitLogger(configPath string) *multilog.Logger {
	cfg, err := multilog.NewConfig(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	hs, err := multilog.CreateHandlers(cfg)
	if err != nil {
		log.Fatalf("failed to create handlers: %v", err)
	}

	logger := multilog.NewLogger(hs...)
	slog.SetDefault(logger.Logger)

	return logger
}
