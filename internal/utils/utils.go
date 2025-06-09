package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	cfg "github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/multilog"
)

// SaveFile saves the content from the reader to the specified destination folder and file name.
// It creates the destination folder if it doesn't exist.
//
// Parameters:
//   - logger: Logger for recording operations and errors
//   - destFolder: Target directory where the file will be saved
//   - fileName: Name of the file to create
//   - reader: Source of content to be saved
//
// Returns:
//   - The absolute path of the saved file or an empty string on error
//   - An error object if the operation failed, nil on success
func SaveFile(logger *multilog.Logger, destFolder, fileName string, reader io.Reader) (string, error) {
	// Ensure the destination folder exists
	if _, err := os.Stat(destFolder); os.IsNotExist(err) {
		err := os.MkdirAll(destFolder, os.ModePerm)
		if err != nil {
			logger.Errorf("Destination folder creation error: %v (folder: %s)", err, destFolder)
			return "", err
		}
	}

	// Save file
	filePath := filepath.Join(destFolder, fileName)
	if _, err := os.Stat(filePath); err == nil {
		logger.Debugf("File already exists: %s", filePath)
	}

	out, err := os.Create(filePath)
	if err != nil {
		logger.Errorf("Creating file error: %v (file: %s)", err, filePath)
		return "", err
	}
	defer CloseFile(logger, out)

	_, err = io.Copy(out, reader)
	if err != nil {
		logger.Errorf("Saving file error: %v (file: %s)", err, filePath)
		return "", err
	}

	return filePath, nil
}

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

// ArchiveExtensions contains the supported archive file extensions.
var ArchiveExtensions = []string{".tar.gz", ".zip"}

// IsArchive checks if a file is an archive based on its extension.
func IsArchive(filePath string) bool {
	for _, ext := range ArchiveExtensions {
		if strings.HasSuffix(filePath, ext) {
			return true
		}
	}
	return false
}

// ExtractArchive extracts the contents of an archive file (either .tar.gz or .zip) to the specified destination folder.
//
// Parameters:
//   - archivePath: Path to the archive file
//   - destFolder: Destination directory where the contents will be extracted
//
// Returns:
//   - An error object if the extraction fails, nil on success
func ExtractArchive(logger *multilog.Logger, archivePath, destFolder string) error {
	return fmt.Errorf("unsupported archive format: %s", archivePath)
}

func CopySourceToTarget(logger *multilog.Logger, target c.DownloadTarget) error {
	sourceFilepath := filepath.Join(target.SourceFolder, target.SourceFile)
	if _, err := os.Stat(sourceFilepath); os.IsNotExist(err) {
		return fmt.Errorf("source file not found: %s", sourceFilepath)
	}
	if _, err := os.Stat(target.TargetFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(target.TargetFolder, os.ModePerm); err != nil {
			return err
		}
	}
	targetFilepath := filepath.Join(target.TargetFolder, target.TargetFile)
	if _, err := os.Stat(targetFilepath); err == nil {
		logger.Debugf("Target file already exists: %s", targetFilepath)
		return nil
	}

	// copy a source file to a target file
	sourceFile, err := os.Open(sourceFilepath)
	if err != nil {
		return err
	}
	defer CloseFile(logger, sourceFile)

	targetFile, err := os.Create(targetFilepath)
	if err != nil {
		return err
	}

	if _, err := io.Copy(targetFile, sourceFile); err != nil {
		return err
	}

	return nil
}
func CloseBody(logger *multilog.Logger, body io.Closer) {
	if err := body.Close(); err != nil {
		logger.Errorf("Closing body error: %v", err)
	}
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

// GetUserAgent returns a user agent string.
func GetUserAgent(logger *multilog.Logger, applicationConfig cfg.ApplicationConfig) string {
	appName := applicationConfig.Name
	if appName == "" {
		appName = constants.AppName
	}
	appVersion := applicationConfig.Version
	if appVersion == "" {
		appVersion = constants.AppVersion
	}
	appDesc := applicationConfig.Description
	if appDesc == "" {
		appDesc = constants.AppDescription
	}
	userAgent := fmt.Sprintf("%s/%s (%s)", appName, appVersion, appDesc)
	return userAgent
}
