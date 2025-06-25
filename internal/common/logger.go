package common

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/phani-kb/multilog"
)

// InitLogger initializes the logger with the specified configuration.
// It resolves relative log file paths to absolute paths based on the project root.
func InitLogger(configPath string) *multilog.Logger {
	// Read the original config file
	configData, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	// Find project root by looking for go.mod from the config file's directory
	projectRoot, err := findProjectRootFromConfig(configPath)
	if err != nil {
		log.Fatalf("failed to find project root: %v", err)
	}

	// Replace relative log file paths in the config data
	configString := string(configData)
	configString = resolveLogFilePaths(configString, projectRoot)

	// Create a temporary config file with resolved paths
	tempConfigFile, err := os.CreateTemp("", "dns-toolkit-config-*.yml")
	if err != nil {
		log.Fatalf("failed to create temp config file: %v", err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			log.Printf("failed to remove temp config file %s: %v", name, err)
		}
	}(tempConfigFile.Name())

	if _, err := tempConfigFile.WriteString(configString); err != nil {
		log.Fatalf("failed to write temp config file: %v", err)
	}
	err = tempConfigFile.Close()
	if err != nil {
		log.Fatalf("failed to close temp config file: %v", err)
	}

	// Use the temporary config file with resolved paths
	cfg, err := multilog.NewConfig(tempConfigFile.Name())
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

// findProjectRootFromConfig finds the project root by traversing up from the config file
func findProjectRootFromConfig(configPath string) (string, error) {
	dir := filepath.Dir(configPath)
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", os.ErrNotExist
}

// resolveLogFilePaths replaces relative log file paths with absolute paths
func resolveLogFilePaths(configString, projectRoot string) string {
	lines := strings.Split(configString, "\n")
	for i, line := range lines {
		if strings.Contains(line, "file:") && !strings.Contains(line, "source_files:") {
			// Extract the file path
			parts := strings.SplitN(line, "file:", 2)
			if len(parts) == 2 {
				filePath := strings.TrimSpace(parts[1])
				// Remove quotes if present
				filePath = strings.Trim(filePath, `"'`)
				// If it's a relative path, make it absolute
				if !filepath.IsAbs(filePath) {
					absolutePath := filepath.Join(projectRoot, filePath)
					lines[i] = parts[0] + "file: " + absolutePath
				}
			}
		}
	}
	return strings.Join(lines, "\n")
}
