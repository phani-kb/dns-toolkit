package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/multilog"
)

// SaveSummary saves a single summary object to the specified file.
// This is a convenience function that calls SaveSummaries with a slice containing only the provided summary.
//
// Type Parameters:
//   - T: The type of the summary object
//
// Parameters:
//   - logger: Logger for recording operations and errors
//   - summary: The summary object to save
//   - summaryFile: Path where the summary will be saved
//   - lessFunc: Function for sorting summaries
//
// Returns:
//   - The number of summaries written to the file
func SaveSummary[T any](
	logger *multilog.Logger,
	summary T,
	summaryFile string,
	lessFunc func(i, j T) bool,
) int {
	summariesCount, err := SaveSummaries(logger, []T{summary}, summaryFile, lessFunc)
	if err != nil {
		logger.Errorf("Failed to save summary: %v", err)
	}
	return summariesCount
}

// SaveSummaries saves multiple summary objects to a JSON file.
// The summaries are sorted using the provided comparison function before being saved.
// If the summary file already exists, it will be archived with a timestamp suffix.
// Small elements are written compactly while preserving readability.
//
// Type Parameters:
//   - T: The type of the summary objects
//
// Parameters:
//   - logger: Logger for recording operations and errors
//   - summaries: Slice of summary objects to save
//   - summaryFile: Path where the summaries will be saved
//   - lessFunc: Function for sorting summaries
//
// Returns:
//   - The number of summaries written to the file
//   - An error object if the operation failed, nil on success
func SaveSummaries[T any](
	logger *multilog.Logger,
	summaries []T,
	summaryFile string,
	lessFunc func(i, j T) bool,
) (int, error) {
	logger.Debugf("Starting SaveSummaries for file: %s", summaryFile)

	if len(summaries) == 0 {
		logger.Debugf("No summaries to write")
		return 0, nil
	}

	// Backup an existing summary file if it exists
	if existingSummaryFileInfo, err := os.Stat(summaryFile); err == nil {
		// File exists, move it to backup directory with timestamp
		// get timestamp of the file

		backupTimestamp := existingSummaryFileInfo.ModTime().Format(constants.BackupFileTimestampFormat)

		// Get the base filename and extension
		baseFileName := filepath.Base(summaryFile)
		fileExt := filepath.Ext(baseFileName)
		fileNameWithoutExt := strings.TrimSuffix(baseFileName, fileExt)

		// Create backup filename with timestamp before the extension
		backupFileName := fmt.Sprintf("%s_%s%s", fileNameWithoutExt, backupTimestamp, fileExt)

		// Full path to the backup file
		backupFilePath := filepath.Join(constants.BackupDir, backupFileName)

		// Create the backup directory if it doesn't exist
		if err := os.MkdirAll(constants.BackupDir, os.ModePerm); err != nil {
			logger.Warnf("Failed to create backup directory: %v", err)
		} else {
			// Copy the existing file to the backup
			if err := copyFile(logger, summaryFile, backupFilePath); err != nil {
				logger.Warnf("Failed to backup summary file: %v", err)
			} else {
				logger.Infof("Backed up summary file to: %s", backupFilePath)
			}
		}
	}

	logger.Debugf("Sorting summaries")
	slices.SortFunc(summaries, func(i, j T) int {
		if lessFunc(i, j) {
			return -1
		}
		if lessFunc(j, i) {
			return 1
		}
		return 0
	})

	file, err := os.Create(summaryFile)
	if err != nil {
		logger.Errorf("Creating summary file error: %v (file: %s)", err, summaryFile)
		return 0, err
	}
	defer CloseFile(logger, file)

	logger.Debugf("Marshaling summaries to JSON")
	data, err := json.Marshal(summaries)
	if err != nil {
		logger.Errorf("Marshaling summaries error: %v", err)
		return 0, err
	}

	logger.Debugf("Formatting JSON output")
	var buf bytes.Buffer
	if err := json.Indent(&buf, data, "", "  "); err != nil {
		logger.Errorf("Formatting JSON error: %v", err)
		return 0, err
	}

	logger.Debugf("Writing summaries to file")
	if _, err := file.Write(buf.Bytes()); err != nil {
		logger.Errorf("Writing summaries error: %v", err)
		return 0, err
	}

	logger.Infof("Successfully wrote summaries to: %s", summaryFile)
	return len(summaries), nil
}

func GetLastSummary[T any](logger *multilog.Logger, summaryFile string, sourceName string) (T, error) {
	var zeroValue T
	logger.Debugf("Getting last summary from file: %s", summaryFile)
	if _, err := os.Stat(summaryFile); err != nil {
		if os.IsNotExist(err) {
			logger.Debugf("Summary file does not exist: %s", summaryFile)
			return zeroValue, nil
		}
		logger.Errorf("Error checking file: %v (file: %s)", err, summaryFile)
		return zeroValue, err
	}
	content, err := os.ReadFile(summaryFile)
	if err != nil {
		logger.Errorf("Reading file error: %v (file: %s)", err, summaryFile)
		return zeroValue, err
	}

	var summaries []T
	if err := json.Unmarshal(content, &summaries); err != nil {
		logger.Errorf("Unmarshalling JSON error: %v", err)
		return zeroValue, err
	}

	if len(summaries) == 0 {
		logger.Debugf("No summaries found")
		return zeroValue, nil
	}

	for _, summary := range summaries {
		switch v := any(summary).(type) {
		case c.DownloadSummary:
			if v.Name == sourceName {
				return summary, nil
			}
		case c.ProcessedSummary:
			if v.Name == sourceName {
				return summary, nil
			}
		default:
			logger.Errorf("Unknown summary type: %T", summary)
		}
	}

	logger.Debugf("No summary found for source: %s", sourceName)
	return zeroValue, nil
}

// GetSummaryFiles is a generic function that retrieves files from a summary file
func GetSummaryFiles[T any](
	logger *multilog.Logger,
	sourceType string,
	summaryFile string,
	extractFiles func(T, string) []string,
) ([]string, error) {
	content, err := os.ReadFile(summaryFile)
	if err != nil {
		logger.Errorf("Error reading %s file: %s", summaryFile, err)
		return nil, err
	}

	var summaries []T
	err = json.Unmarshal(content, &summaries)
	if err != nil {
		logger.Errorf("Parsing summary error: %v", err)
		return nil, err
	}

	files := make([]string, 0)
	for _, summary := range summaries {
		files = append(files, extractFiles(summary, sourceType)...)
	}
	return files, nil
}

// DetermineSummaryTypeFromPath determines the summary type based on a file path
func DetermineSummaryTypeFromPath(path string) string {
	filename := filepath.Base(path)

	// Check if the filename starts with any of the known summary type prefixes
	for _, summaryType := range constants.AllSummaryTypes {
		if strings.HasPrefix(filename, summaryType) {
			// Special case for consolidated_groups vs. consolidated
			if summaryType == constants.SummaryTypeConsolidated {
				if strings.HasPrefix(filename, constants.SummaryTypeConsolidatedGroups) {
					return constants.SummaryTypeConsolidatedGroups
				} else if strings.HasPrefix(filename, constants.SummaryTypeConsolidatedCategories) {
					return constants.SummaryTypeConsolidatedCategories
				}
			}
			return summaryType
		}
	}

	return constants.SummaryTypeUnknown
}

// GetSummaryTypeFromFolder determines the summary type based on the folder name
func GetSummaryTypeFromFolder(folderName string) string {
	// Check direct mapping
	if summaryType, exists := constants.FolderToSummaryTypeMap[folderName]; exists {
		return summaryType
	}

	return constants.SummaryTypeUnknown
}

// GetFoldersToArchive returns a map of folders that should be archived
func GetFoldersToArchive(logger *multilog.Logger, folders map[string]string) map[string]string {
	result := make(map[string]string)
	// Add folders based on summary types
	for folderKey, folderPath := range folders {
		summaryType := GetSummaryTypeFromFolder(folderKey)
		logger.Debugf("Processing folder: %s, Summary Type: %s, Path: %s", folderKey, summaryType, folderPath)
		if summaryType != constants.SummaryTypeUnknown && !constants.ArchiveFoldersToSkipMap[folderPath] {
			result[folderPath] = summaryType
		} else {
			logger.Debugf("Skipping folder %s for summary type %s", folderPath, summaryType)
		}
	}

	return result
}

func GetFilesFromSummaries[T any](summaries []T, summaryType string) map[string]T {
	files := make(map[string]T)
	for _, summary := range summaries {
		if file := getFilePathForSummary(summary, summaryType); file != "" {
			files[file] = summary
		}
	}
	return files
}

func getFilePathForSummary[T any](summary T, summaryType string) string {
	switch summaryType {
	case constants.SummaryTypeConsolidated,
		constants.SummaryTypeConsolidatedGroups,
		constants.SummaryTypeConsolidatedCategories:
		if s, ok := any(summary).(c.ConsolidatedSummary); ok {
			return s.Filepath
		}
	case constants.SummaryTypeTop:
		if s, ok := any(summary).(c.TopSummary); ok {
			return s.Filepath
		}
	}
	return ""
}
