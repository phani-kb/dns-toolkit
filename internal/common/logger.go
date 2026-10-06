package common

import (
	"log"
	"log/slog"
	"path/filepath"

	"github.com/phani-kb/multilog"
)

// InitLogger initializes the logger with the specified configuration.
func InitLogger(configPath string) *multilog.Logger {
	cfg, err := multilog.NewConfig(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	baseDir := filepath.Dir(configPath)
	switch filepath.Base(baseDir) {
	case "configs", "testdata":
		baseDir = filepath.Dir(baseDir)
	}

	for i := range cfg.Multilog.Handlers {
		handler := &cfg.Multilog.Handlers[i]
		if handler.File == "" || filepath.IsAbs(handler.File) {
			continue
		}
		handler.File = filepath.Join(baseDir, handler.File)
	}

	hs, err := multilog.CreateHandlers(cfg)
	if err != nil {
		log.Fatalf("failed to create handlers: %v", err)
	}

	logger := multilog.NewLogger(hs...)
	slog.SetDefault(logger.Logger)

	return logger
}
