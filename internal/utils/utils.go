package utils

import (
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/multilog"
)

// CloseFile safely closes the given file and logs an error if it fails.
//
// Parameters:
//   - logger: Logger for recording errors
//   - file: The file handle to close
func CloseFile(logger *multilog.Logger, file *os.File) {
	err := file.Close()
	if err != nil {
		logger.Errorf("Closing file error: %v (file: %s)", err, file.Name())
	}
}

// IsComment determines if a line is a comment or an empty line.
// A line is considered a comment if it's empty or starts with any of the common comment prefixes.
//
// Parameters:
//   - line: The string to check
//
// Returns:
//   - true if the line is a comment or empty, false otherwise
func IsComment(line string) bool {
	trimmedLine := strings.TrimSpace(line)
	if trimmedLine == "" {
		return true
	}

	for _, prefix := range constants.CommentPrefixes {
		if strings.HasPrefix(trimmedLine, prefix) {
			return true
		}
	}
	return false
}

// GetTimestamp returns the current time formatted according to the application's standard timestamp format.
//
// Returns:
//   - A string representation of the current time
func GetTimestamp() string {
	return time.Now().Format(constants.TimestampFormat)
}

// IsAlphanumericWithUnderscoresAndDashes checks if a string contains only alphanumeric characters, underscores, and dashes.
func IsAlphanumericWithUnderscoresAndDashes(s string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", s)
	return match
}

// GetFileLastModifiedTime retrieves the last modified time of a file and formats it according to the application's standard timestamp format.
func GetFileLastModifiedTime(logger *multilog.Logger, filePath string) (string, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		logger.Errorf("Getting file info error: %v", err)
		return "", err
	}
	return fileInfo.ModTime().Format(constants.TimestampFormat), nil
}

func LogMemStats(logger *multilog.Logger, prefix string) {
	logger.Perff(prefix)
}

// EnsureDirectoryExists creates a directory if it doesn't exist
func EnsureDirectoryExists(logger *multilog.Logger, dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		logger.Infof("Creating directory: %s", dir)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			logger.Errorf("Error creating directory %s: %v", dir, err)
			return err
		}
	}
	return nil
}
