package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/spf13/cobra"
)

var includeIgnored bool
var deleteFolders bool

// prepareDirectories creates necessary output directories
func prepareDirectories() error {
	// Create an output directory
	if err := os.MkdirAll(constants.OutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create subdirectories for different consolidated types
	for summaryType, dirName := range constants.SummaryTypesOutputDirMap {
		if constants.SummaryTypesOutputToSkipMap[summaryType] {
			continue
		}

		dir := filepath.Join(dirName)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory for summary type %s: %w", summaryType, err)
		}
	}

	// Create an ignored directory if the flag is set
	if includeIgnored {
		if err := os.MkdirAll(constants.OutputIgnoredDir, 0755); err != nil {
			return fmt.Errorf("failed to create ignored directory: %w", err)
		}
	}

	// Create a summaries directory
	if err := os.MkdirAll(constants.OutputSummariesDir, 0755); err != nil {
		return fmt.Errorf("failed to create summaries directory: %w", err)
	}

	// Create an archive directory
	if err := os.MkdirAll(constants.ArchiveDir, 0755); err != nil {
		return fmt.Errorf("failed to create archive directory: %w", err)
	}

	return nil
}

// loadTemplates loads and parses template files
func loadTemplates() (*template.Template, []byte, error) {
	configsDir := "configs"

	// In test mode, resolve relative paths to project root
	if os.Getenv("DNS_TOOLKIT_TEST_MODE") == "true" {
		if projectRoot, err := u.FindProjectRoot(""); err == nil {
			configsDir = filepath.Join(projectRoot, "configs")
		}
	}

	staticTemplatePath := filepath.Join(configsDir, "templates", "static_template.txt")
	dynamicTemplatePath := filepath.Join(configsDir, "templates", "dynamic_template.txt")

	staticTemplate, err := os.ReadFile(staticTemplatePath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read static template: %w", err)
	}

	dynamicTemplateTxt, err := os.ReadFile(dynamicTemplatePath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read dynamic template: %w", err)
	}

	tmpl, err := template.New("dynamic").Parse(string(dynamicTemplateTxt))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse dynamic template: %w", err)
	}

	return tmpl, staticTemplate, nil
}

// processFilesForSummaryType extracts files information from summary data
func processFilesForSummaryType(
	summaryType string,
	summaryData []byte,
) (map[string]string, map[string]int, map[string]int) {
	typeFiles := make(map[string]string)
	fileCount := make(map[string]int)
	ignoredFilesCount := make(map[string]int)

	switch summaryType {
	case constants.SummaryTypeConsolidated:
		var summaries []common.ConsolidatedSummary
		if err := json.Unmarshal(summaryData, &summaries); err != nil {
			Logger.Error("Failed to unmarshal consolidated summary", "error", err)
			return typeFiles, fileCount, ignoredFilesCount
		}
		files := u.GetFilesFromSummaries(summaries, constants.SummaryTypeConsolidated)
		for key, value := range files {
			typeFiles[key] = value.ListType
			fileCount[key] = value.Count
			if value.IgnoredEntriesCount > 0 && value.IgnoredFilepath != "" {
				ignoredFilesCount[value.IgnoredFilepath] = value.IgnoredEntriesCount
			}
		}

	case constants.SummaryTypeConsolidatedGroups:
		var sizedSummaries []common.ConsolidatedGroupsSummary
		if err := json.Unmarshal(summaryData, &sizedSummaries); err != nil {
			Logger.Error("Failed to unmarshal consolidated groups summary", "error", err)
			return typeFiles, fileCount, ignoredFilesCount
		}
		for _, summary := range sizedSummaries {
			files := u.GetFilesFromSummaries(
				summary.ConsolidatedSummaries,
				constants.SummaryTypeConsolidatedGroups,
			)
			for key, value := range files {
				typeFiles[key] = value.ListType
				fileCount[key] = value.Count
				if value.IgnoredEntriesCount > 0 && value.IgnoredFilepath != "" {
					ignoredFilesCount[value.IgnoredFilepath] = value.IgnoredEntriesCount
				}
			}
		}

	case constants.SummaryTypeConsolidatedCategories:
		var categorySummaries []common.ConsolidatedCategoriesSummary
		if err := json.Unmarshal(summaryData, &categorySummaries); err != nil {
			Logger.Error("Failed to unmarshal consolidated categories summary", "error", err)
			return typeFiles, fileCount, ignoredFilesCount
		}
		for _, summary := range categorySummaries {
			files := u.GetFilesFromSummaries(
				summary.ConsolidatedSummaries,
				constants.SummaryTypeConsolidatedCategories,
			)
			for key, value := range files {
				typeFiles[key] = value.ListType
				fileCount[key] = value.Count
				if value.IgnoredEntriesCount > 0 && value.IgnoredFilepath != "" {
					ignoredFilesCount[value.IgnoredFilepath] = value.IgnoredEntriesCount
				}
			}
		}

	case constants.SummaryTypeTop:
		var topSummaries []common.TopSummary
		if err := json.Unmarshal(summaryData, &topSummaries); err != nil {
			Logger.Error("Failed to unmarshal top summary", "error", err)
			return typeFiles, fileCount, ignoredFilesCount
		}
		files := u.GetFilesFromSummaries(topSummaries, constants.SummaryTypeTop)
		for key, value := range files {
			typeFiles[key] = value.ListType
			fileCount[key] = value.Count
		}

	default:
		Logger.Error("Unknown summary type", "type", summaryType)
	}

	return typeFiles, fileCount, ignoredFilesCount
}

// createOutputFromFile creates an output file with template headers
func createOutputFromFile(
	tmpl *template.Template,
	staticTemplate []byte,
	filePath string,
	fileName string,
	description string,
	count int,
	outputPath string,
) error {
	// Get last updated time
	lastUpdated := time.Now().Format(constants.TimestampFormat)
	if info, err := os.Stat(filePath); err == nil {
		lastUpdated = info.ModTime().Format(constants.TimestampFormat)
	} else {
		Logger.Error("Getting file info error", "error", err)
	}

	// Execute dynamic template
	var dynamicOutput bytes.Buffer
	err := tmpl.Execute(&dynamicOutput, common.TemplateData{
		AppName:        AppConfig.Application.Name,
		AppVersion:     AppConfig.Application.Version,
		AppDescription: AppConfig.Application.Description,
		FileName:       fileName,
		LastUpdated:    lastUpdated,
		Description:    description,
		Count:          count,
	})

	if err != nil {
		return fmt.Errorf("failed to execute dynamic template: %w", err)
	}

	// Read file content
	dataContent, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Combine templates and data
	output := fmt.Sprintf("%s\n%s\n%s\n%s",
		dynamicOutput.String(),
		string(staticTemplate),
		constants.ContentSeparator,
		string(dataContent))

	// Write an output file
	if err := os.WriteFile(outputPath, []byte(output), 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}

// processRegularFiles processes and generates output for regular files
func processRegularFiles(
	tmpl *template.Template,
	staticTemplate []byte,
	summaryType string,
	typeFiles map[string]string,
	fileCount map[string]int,
) {
	for filePath, listType := range typeFiles {
		fileName := filepath.Base(filePath)
		format := constants.SummaryTypesMap[summaryType]
		description := generateDescription(summaryType, fileName, format, listType)
		outDir := constants.SummaryTypesOutputDirMap[summaryType]
		outputFilePath := filepath.Join(outDir, fileName)

		err := createOutputFromFile(
			tmpl,
			staticTemplate,
			filePath,
			fileName,
			description,
			fileCount[filePath],
			outputFilePath,
		)

		if err != nil {
			Logger.Error("Failed to create output file", "error", err)
			continue
		}

		Logger.Debug("Successfully generated output file", "path", outputFilePath, "from", filePath)
	}
}

// processIgnoredFiles processes and generates output for ignored files
func processIgnoredFiles(
	tmpl *template.Template,
	staticTemplate []byte,
	summaryType string,
	ignoredFilesCount map[string]int,
) {
	if !includeIgnored || len(ignoredFilesCount) == 0 {
		return
	}

	Logger.Info("Processing ignored files", "count", len(ignoredFilesCount))

	for ignoredFilePath, ignoredCount := range ignoredFilesCount {
		if _, err := os.Stat(ignoredFilePath); os.IsNotExist(err) {
			Logger.Error("Ignored file does not exist", "file", ignoredFilePath)
			continue
		}

		ignoredFileName := filepath.Base(ignoredFilePath)
		ignoredOutputPath := filepath.Join(constants.OutputIgnoredDir, ignoredFileName)
		ignoredDescription := generateDescription(
			summaryType,
			ignoredFileName,
			constants.SummaryTypesMap[summaryType],
			"Ignored",
		)

		err := createOutputFromFile(
			tmpl,
			staticTemplate,
			ignoredFilePath,
			ignoredFileName,
			ignoredDescription,
			ignoredCount,
			ignoredOutputPath,
		)

		if err != nil {
			Logger.Error("Failed to create ignored output file", "error", err)
			continue
		}

		Logger.Debug(
			"Successfully generated ignored output file",
			"path",
			ignoredOutputPath,
			"from",
			ignoredFilePath,
		)
	}
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate different types of outputs",
	Long:  "Generate various outputs from DNS toolkit data",
}

var generateOutputCmd = &cobra.Command{
	Use:   "output",
	Short: "Generate output files with templates prefixed to them",
	Long:  "Generate output files with static and dynamic templates prefixed to the summary types defined in SummaryTypesWithTemplateMap",
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getenv("DNS_TOOLKIT_TEST_MODE") == "true" {
			return
		}

		Logger.Info("Starting generate prefixes command...")

		if err := u.EnsureDirectoryExists(Logger, constants.OutputDir); err != nil {
			Logger.Errorf("Failed to create output directory: %v", err)
			os.Exit(1)
		}

		// Prepare directories
		if err := prepareDirectories(); err != nil {
			Logger.Error("Failed to prepare directories", "error", err)
			return
		}

		// Load templates
		tmpl, staticTemplate, err := loadTemplates()
		if err != nil {
			Logger.Error("Failed to load templates", "error", err)
			return
		}

		var processedSummaryFiles = make(map[string]string)
		// Process each summary type
		for summaryType, summaryFile := range constants.SummaryTypesWithTemplateMap {
			summaryFilePath := filepath.Join(constants.SummaryDir, summaryFile)
			Logger.Info("Processing summary type", "type", summaryType, "file", summaryFilePath)

			if _, err := os.Stat(summaryFilePath); os.IsNotExist(err) {
				Logger.Error("Summary file does not exist", "file", summaryFilePath)
				continue
			}

			summaryData, err := os.ReadFile(summaryFilePath)
			if err != nil {
				Logger.Error("Failed to read summary file", "error", err)
				continue
			}

			// Process files based on a summary type
			typeFiles, fileCount, ignoredFilesCount := processFilesForSummaryType(summaryType, summaryData)
			Logger.Info("Extracted files from summary", "count", len(typeFiles))

			// Process regular files
			processRegularFiles(tmpl, staticTemplate, summaryType, typeFiles, fileCount)

			// Process ignored files
			processIgnoredFiles(tmpl, staticTemplate, summaryType, ignoredFilesCount)

			processedSummaryFiles[summaryType] = summaryFilePath
		}

		Logger.Info("Processed summary files", "count", len(processedSummaryFiles))

		// Copy summary files to the output directory without timestamps
		Logger.Info("Copying summary files to output directory without timestamps...")
		copySummaryFiles(processedSummaryFiles, constants.OutputSummariesDir)

		// Archive summary files with timestamps
		Logger.Info("Archiving summary files with timestamps...")
		archiveSummaryFiles(processedSummaryFiles)

		deleteFilesAndFoldersAfterGeneration()

		Logger.Info("Finished generate prefixes command.")
	},
}

// generateDescription creates a human-readable description
func generateDescription(_, fileName string, format string, listType string) string {
	var description string

	entryType := "General"
	if strings.Contains(fileName, "ipv4") {
		entryType = "IPv4"
	} else if strings.Contains(fileName, "domain") {
		entryType = "Domain"
	} else if strings.Contains(fileName, "ipv6") {
		entryType = "IPv6"
	} else if strings.Contains(fileName, "adguard") {
		entryType = "AdGuard"
	} else if strings.Contains(fileName, "cidr") {
		entryType = "CIDR"
	}

	// check if the fileName has the suffix "_ignored"
	if strings.HasSuffix(fileName, "_ignored") {
		entryType = "Ignored " + entryType
	}

	var minSourcesInfo string
	if strings.Contains(fileName, "top") {
		// Look for the "min" pattern in the filename
		if strings.Contains(fileName, "min") {
			parts := strings.Split(fileName, "_")
			for _, part := range parts {
				if strings.HasPrefix(part, "min") {
					minSuffix := strings.TrimPrefix(part, "min")
					// Remove file extension if present
					minSuffix = strings.TrimSuffix(minSuffix, filepath.Ext(minSuffix))
					if minSuffix != "" {
						minSourcesInfo = fmt.Sprintf(" (minimum %s sources)", minSuffix)
						break
					}
				}
			}
		}
	}

	if listType == "" {
		listType = "unknown"
	}

	format = strings.ReplaceAll(format, "_", " ")
	format = cases.Title(language.English).String(format)

	description = fmt.Sprintf("%s %s %s%s", format, entryType, listType, minSourcesInfo)

	return description
}

// copySummaryFiles copies summary files to the output directory without timestamps
func copySummaryFiles(processedFiles map[string]string, outputDir string) {
	for summaryType, summaryFile := range constants.SummaryTypesOutputSummaryFileMap {
		if constants.SummaryTypesOutputSummaryFileToSkipMap[summaryType] {
			continue
		}
		summaryFilePath := processedFiles[summaryType]
		if summaryFilePath == "" {
			summaryFilePath = filepath.Join(constants.SummaryDir, summaryFile)
		}
		// Check if the summary file exists
		if _, err := os.Stat(summaryFilePath); os.IsNotExist(err) {
			Logger.Error("Summary file does not exist", "file", summaryFilePath)
			continue
		}

		summaryFileNameNoTimestamp := filepath.Base(summaryFile)
		destFilePathNoTimestamp := filepath.Join(outputDir, summaryFileNameNoTimestamp)

		Logger.Debug("Copying summary file without timestamp",
			"source", summaryFilePath,
			"destination", destFilePathNoTimestamp)

		if err := copySummaryFile(summaryFilePath, destFilePathNoTimestamp); err != nil {
			Logger.Error("Failed to copy summary file to output directory", "error", err)
		} else {
			Logger.Info("Successfully copied summary file to output directory",
				"file", destFilePathNoTimestamp)
		}
	}
}

// archiveSummaryFiles copies summary files to the archive directory with timestamps
func archiveSummaryFiles(processedFiles map[string]string) {
	for summaryType, summaryFile := range constants.SummaryTypesOutputSummaryFileMap {
		if constants.SummaryTypesOutputSummaryFileToSkipMap[summaryType] {
			continue
		}
		summaryFilePath := processedFiles[summaryType]
		if summaryFilePath == "" {
			summaryFilePath = filepath.Join(constants.SummaryDir, summaryFile)
		}
		// Check if the summary file exists
		if _, err := os.Stat(summaryFilePath); os.IsNotExist(err) {
			Logger.Error("Summary file does not exist for archiving", "file", summaryFilePath)
			continue
		}

		// Get file modification time for archive copy
		modTimestamp, err := u.GetFileLastModifiedTime(Logger, summaryFilePath)
		if err != nil {
			Logger.Error("Failed to get file last modified time", "error", err)
			// use the current date
			modTimestamp = time.Now().Format(constants.TimestampFormat)
		}

		summaryFileNameNoTimestamp := filepath.Base(summaryFile)
		archiveFileName := strings.TrimSuffix(summaryFileNameNoTimestamp, filepath.Ext(summaryFileNameNoTimestamp))
		archiveFileName = fmt.Sprintf("%s_%s.json", archiveFileName, modTimestamp)
		archiveFilePath := filepath.Join(constants.ArchiveDir, archiveFileName)

		Logger.Debug("Copying summary file with timestamp to archive",
			"source", summaryFilePath,
			"destination", archiveFilePath)

		if err := copySummaryFile(summaryFilePath, archiveFilePath); err != nil {
			Logger.Error("Failed to copy summary file to archive directory", "error", err)
		} else {
			Logger.Info("Successfully archived summary file", "file", archiveFilePath)
		}
	}
}

func copySummaryFile(srcPath, dstPath string) error {
	input, err := os.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("failed to read summary file: %w", err)
	}
	if err := os.WriteFile(dstPath, input, 0644); err != nil {
		return fmt.Errorf("failed to write summary file: %w", err)
	}
	return nil
}

func deleteFilesAndFoldersAfterGeneration() {
	if !deleteFolders {
		return
	}

	Logger.Info("Deleting folders after output generation...")

	for summaryType, shouldDelete := range constants.SummaryTypesToDeleteAfterOutputGenerationMap {
		if shouldDelete {
			// get the summary file path
			filePath := filepath.Join(constants.SummaryDir, constants.SummaryTypesOutputSummaryFileMap[summaryType])
			Logger.Info("Deleting file", "file", filePath)
			// Check if a file exists
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				Logger.Debug("File does not exist, skipping", "file", filePath)
				continue
			}
			// Delete the file
			if err := os.Remove(filePath); err != nil {
				Logger.Error("Failed to delete file", "file", filePath, "error", err)
				continue
			}

			Logger.Info("Successfully deleted file", "file", filePath)

			folderPath := constants.SummaryTypesDirMap[summaryType]

			Logger.Info("Deleting folder", "folder", folderPath)

			// Check if the folder exists
			if _, err := os.Stat(folderPath); os.IsNotExist(err) {
				Logger.Debug("Folder does not exist, skipping", "folder", folderPath)
				continue
			}
			// Delete the folder
			if err := os.RemoveAll(folderPath); err != nil {
				Logger.Error("Failed to delete folder", "folder", folderPath, "error", err)
				continue
			}

			Logger.Info("Successfully deleted folder", "folder", folderPath)
		}
	}
}

func init() {
	generateCmd.AddCommand(generateOutputCmd)

	// Add the includeIgnored flag
	generateOutputCmd.Flags().BoolVarP(&includeIgnored, "include-ignored", "i", false,
		"Include ignored files in the output by copying to the ignored subfolder")

	// Add the deleteFolders flag
	generateOutputCmd.Flags().BoolVarP(&deleteFolders, "delete-folders", "d", false,
		"Delete folders specified in FoldersToDeleteAfterOutputGeneration after output generation")
}
