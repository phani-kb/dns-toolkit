package config

import (
	"io"
	"log/slog"

	"github.com/phani-kb/multilog"
)

// CreateTestLogger creates a properly initialized logger for tests
func CreateTestLogger() *multilog.Logger {
	// Create a no-op handler that doesn't write anywhere
	handler := slog.NewTextHandler(io.Discard, nil)
	logger := multilog.Logger{
		Logger: slog.New(handler),
	}
	return &logger
}
