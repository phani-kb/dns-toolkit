package cmd

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"sync"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	cfg "github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	r "github.com/phani-kb/dns-toolkit/internal/processors"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/spf13/cobra"
)

// processCmd is the cobra command for processing downloaded files.
// It extracts and validates content from the files, categorizing them into valid and invalid entries.
var processCmd = &cobra.Command{
	Use:   "process",
	Short: "Process downloaded files",
	Run: func(cmd *cobra.Command, args []string) {
		if err := u.EnsureDirectoryExists(Logger, constants.ProcessedDir); err != nil {
			Logger.Errorf("Failed to create processed directory: %v", err)
			os.Exit(1)
		}
		if err := u.EnsureDirectoryExists(Logger, constants.SummaryDir); err != nil {
			Logger.Errorf("Failed to create summary directory: %v", err)
			os.Exit(1)
		}

		ctx := context.Background()
		processAllSources(ctx, Logger, constants.ProcessedDir)
	},
}

// processAllSources processes all downloaded source files.
//
// Parameters:
//   - ctx: Context for cancellation
//   - logger: Logger for recording operations and errors
//   - processedDir: Directory to save processed results
func processAllSources(ctx context.Context, logger *multilog.Logger, processedDir string) {
	dsf := filepath.Join(constants.SummaryDir, constants.DefaultSummaryFiles["download"])
	content, err := os.ReadFile(dsf)
	if err != nil {
		logger.Errorf("Error reading %s file: %s", dsf, err)
		return
	}

	var downloadSummaries []c.DownloadSummary
	err = json.Unmarshal(content, &downloadSummaries)
	if err != nil {
		logger.Errorf("Parsing download summary error: %v", err)
		return
	}

	processedSummariesMap := make(map[string]c.ProcessedSummary)
	var mu sync.Mutex

	// Create a worker pool for controlled concurrency
	maxWorkers := AppConfig.DNSToolkit.MaxWorkers
	logger.Infof("Using worker pool with %d worker(s) for processing", maxWorkers)
	workerPool := c.NewDTWorkerPool(maxWorkers)

	for _, summary := range downloadSummaries {
		if summary.Error != "" {
			logger.Warnf("Skipping processing for %s: %s", summary.Name, summary.Error)
			continue
		}
		if !cfg.IsEnabledSource(summary.Name, SourcesConfigs, *AppConfig) {
			logger.Debugf("Skipping processing for %s: source is disabled", summary.Name)
			continue
		}

		summary := summary // Create a local copy for goroutine
		workerPool.Submit(func() {

			// Check for context cancellation
			select {
			case <-ctx.Done():
				logger.Warnf("Processing cancelled for %s: %v", summary.Name, ctx.Err())
				return
			default:
				// Continue processing
			}

			processedSummaries := processSourceFile(ctx, logger, summary, processedDir)

			mu.Lock()
			for _, processedSummary := range processedSummaries {
				if existingSummary, exists := processedSummariesMap[summary.Name]; exists {
					mergeSummaries(&existingSummary, &processedSummary)
					processedSummariesMap[summary.Name] = existingSummary
				} else {
					processedSummariesMap[summary.Name] = processedSummary
				}
			}
			mu.Unlock()
		})
	}

	workerPool.Wait()

	var processedSummaries []c.ProcessedSummary
	for _, summary := range processedSummariesMap {
		processedSummaries = append(processedSummaries, summary)
	}

	summaryFile := filepath.Join(constants.SummaryDir, constants.DefaultSummaryFiles["processed"])
	_, err = u.SaveSummaries(logger, processedSummaries, summaryFile, c.ProcessedSummaryLessFunc)
	if err != nil {
		logger.Errorf("Saving processed summaries error: %v", err)
	}
}

// processSourceFile processes a single source file.
// It reads the file content and processes it according to its source types.
//
// Parameters:
//   - ctx: Context for cancellation
//   - logger: Logger for recording operations and errors
//   - summary: Download summary containing file information
//   - processedDir: Directory to save processed results
//
// Returns:
//   - A ProcessedSummary containing information about the processing results
func processSourceFile(
	_ context.Context,
	logger *multilog.Logger,
	summary c.DownloadSummary,
	processedDir string,
) []c.ProcessedSummary {
	processedSummaries := make([]c.ProcessedSummary, 0)
	content, err := os.ReadFile(summary.Filepath)
	if err != nil {
		logger.Errorf("Reading file error: %v (file: %s)", err, summary.Filepath)
		return processedSummaries
	}

	validFiles := make(map[string]c.ProcessedFile)
	invalidFiles := make(map[string]c.ProcessedFile)

	for _, sourceTypeObj := range summary.GetSourceTypes() {
		sourceTypeName := sourceTypeObj.Name
		for _, listTypeObj := range sourceTypeObj.GetListTypes() {
			listTypeName := listTypeObj.Name
			mustConsider := listTypeObj.MustConsider
			validEntries, invalidEntries := extractEntriesByType(logger, string(content), sourceTypeName, listTypeName)
			validFilePath, invalidFilePath := saveEntries(
				logger,
				processedDir,
				summary.Name,
				sourceTypeName,
				listTypeName,
				validEntries,
				invalidEntries,
			)

			key := fmt.Sprintf("%s_%s", sourceTypeName, listTypeName)
			if validFilePath != "" {
				validFiles[key] = createProcessedFile(
					logger,
					summary.Name,
					validFilePath,
					sourceTypeName,
					listTypeName,
					validEntries,
					mustConsider,
					true,
					listTypeObj.Groups,
					summary.Categories,
				)
			}
			if invalidFilePath != "" {
				invalidFiles[key] = createProcessedFile(
					logger,
					summary.Name,
					invalidFilePath,
					sourceTypeName,
					listTypeName,
					invalidEntries,
					mustConsider,
					false,
					listTypeObj.Groups,
					summary.Categories,
				)
			}

			logger.Infof(
				"Processed %s (%s-%s): %d valid, %d invalid",
				summary.Name,
				sourceTypeName,
				listTypeName,
				len(validEntries),
				len(invalidEntries),
			)
		}

		// get the valid and invalid files arrays
		var validFilesArray, invalidFilesArray []c.ProcessedFile
		for _, file := range validFiles {
			validFilesArray = append(validFilesArray, file)
		}
		for _, file := range invalidFiles {
			invalidFilesArray = append(invalidFilesArray, file)
		}

		ps := createSummary(summary.Name, summary.GetSourceTypes(), validFilesArray, invalidFilesArray)
		processedSummaries = append(processedSummaries, ps)
	}

	return processedSummaries
}

// createProcessedFile creates a ProcessedFile object for the given file.
// It includes metadata about the file and calculates a checksum if enabled.
//
// Parameters:
//   - logger: Logger for recording operations and errors
//   - fileName: Name of the processed file
//   - sourceType: Type of source (domain, ipv4, etc.)
//   - listType: Type of list (blocklist, allowlist, etc.)
//   - entries: The entries contained in the file
//
// Returns:
//   - A ProcessedFile struct with file metadata
func createProcessedFile(
	logger *multilog.Logger,
	name, filePath, sourceType, listType string,
	entries []string,
	mustConsider bool,
	valid bool,
	groups []string,
	categories []string,
) c.ProcessedFile {
	var checksum string
	if AppConfig.DNSToolkit.FilesChecksum.Enabled {
		checksum = u.CalculateChecksum(
			logger,
			filePath,
			AppConfig.DNSToolkit.FilesChecksum.Algorithm,
		)
	}
	return c.ProcessedFile{
		Name:              name,
		GenericSourceType: cfg.GetGenericSourceType(sourceType),
		ActualSourceType:  sourceType,
		ListType:          listType,
		Filepath:          filePath,
		NumberOfEntries:   len(entries),
		Checksum:          checksum,
		MustConsider:      mustConsider,
		Valid:             valid,
		Groups:            groups,
		Categories:        categories,
	}
}

// extractEntriesByType extracts entries from content based on the source type.
// It uses either built-in regex patterns or custom processors to extract entries.
//
// Parameters:
//   - logger: Logger for recording operations and errors
//   - content: The content to process
//   - sourceType: Type of source (domain, ipv4, etc.)
//   - listType: Type of list (allowlist, blocklist, etc.)
//
// Returns:
//   - A slice of valid entries
//   - A slice of invalid entries
func extractEntriesByType(
	logger *multilog.Logger,
	content string,
	sourceType string,
	listType string,
) ([]string, []string) {
	var validEntries, invalidEntries []string

	if regex, exists := constants.SourceTypeRegexMap[sourceType]; exists {
		vEntries, iEntries := extractEntriesWithRegex(content, regex)
		validEntries = append(validEntries, vEntries...)
		invalidEntries = append(invalidEntries, iEntries...)
	} else if processor, exists := r.Processors.GetProcessor(sourceType, listType); exists {
		vEntries, iEntries := processor.Process(logger, content)
		validEntries = append(validEntries, vEntries...)
		invalidEntries = append(invalidEntries, iEntries...)
	} else {
		logger.Warnf("Unsupported source type: %s", sourceType)
	}

	return u.RemoveDuplicates(validEntries), u.RemoveDuplicates(invalidEntries)
}

// createSummary creates a ProcessedSummary from the processing results.
// It includes information about the source, the valid and invalid files,
// and when the processing was completed.
//
// Parameters:
//   - summary: The original download summary
//   - validFiles: Map of valid processed files by source type
//   - invalidFiles: Map of invalid processed files by source type
//
// Returns:
//   - A ProcessedSummary struct with processing metadata
func createSummary(
	name string, types []c.SourceType,
	validFiles []c.ProcessedFile,
	invalidFiles []c.ProcessedFile,
) c.ProcessedSummary {
	// Filter out disabled list types from each source type
	filteredTypes := make([]c.SourceType, 0)
	for _, sourceType := range types {
		filteredType := sourceType
		filteredListTypes := make([]c.ListType, 0)

		for _, listType := range sourceType.ListTypes {
			if !listType.Disabled {
				filteredListTypes = append(filteredListTypes, listType)
			}
		}

		if len(filteredListTypes) > 0 {
			filteredType.ListTypes = filteredListTypes
			filteredTypes = append(filteredTypes, filteredType)
		}
	}

	return c.ProcessedSummary{
		Name:                   name,
		Types:                  filteredTypes,
		ValidFiles:             validFiles,
		InvalidFiles:           invalidFiles,
		LastProcessedTimestamp: u.GetTimestamp(),
	}
}

// extractEntriesWithRegex extracts entries from content using a regex pattern.
// Lines that match the regex are considered valid, others invalid.
//
// Parameters:
//   - content: The content to process
//   - regex: The regex pattern to match against
//
// Returns:
//   - A slice of valid entries (match the regex)
//   - A slice of invalid entries (don't match the regex)
func extractEntriesWithRegex(content string, regex *regexp.Regexp) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if u.IsComment(line) {
			continue
		}
		matchedString := regex.FindString(line)
		if matchedString != "" {
			validEntries = append(validEntries, matchedString)
		} else {
			invalidEntries = append(invalidEntries, line)
		}
	}
	return u.RemoveDuplicates(validEntries), u.RemoveDuplicates(invalidEntries)
}

// saveEntries saves valid and invalid entries to separate files.
// It generates filenames based on the source and entry type.
//
// Parameters:
//   - logger: Logger for recording operations and errors
//   - processedDir: Directory to save processed files
//   - name: Name of the source
//   - sourceType: Type of source (domain, ipv4, etc.)
//   - validEntries: The valid entries to save
//   - invalidEntries: The invalid entries to save
//
// Returns:
//   - The name of the valid entries file, or empty if none
//   - The name of the invalid entries file, or empty if none
func saveEntries(
	logger *multilog.Logger,
	processedDir, name, sourceType, listType string,
	validEntries, invalidEntries []string,
) (string, string) {
	var validFilePath, invalidFilePath string

	save := func(entries []string, entryType string) string {
		if len(entries) == 0 {
			return ""
		}
		fileName := generateFileName(logger, name, sourceType, listType, entryType)
		filePath := filepath.Join(processedDir, fileName)
		saveToFile(logger, filePath, entries)
		return filePath
	}

	validFilePath = save(validEntries, "valid")
	invalidFilePath = save(invalidEntries, "invalid")

	return validFilePath, invalidFilePath
}

// saveToFile saves entries to a file in the processed directory.
// It deduplicates and sorts the entries before saving.
//
// Parameters:
//   - logger: Logger for recording operations and errors
//   - filePath: Path of the file to create
//   - entries: Entries to write to the file
func saveToFile(logger *multilog.Logger, filePath string, entries []string) {
	uniqueEntries := u.RemoveDuplicates(entries)
	if len(uniqueEntries) == 0 {
		return
	}

	sorted := slices.Clone(uniqueEntries)
	slices.Sort(sorted)

	content := strings.Join(sorted, "\n") + "\n"
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		logger.Errorf("Error writing file: %v (path: %s)", err, filePath)
	}
}

// generateFileName generates a unique filename for a processed file.
// The filename includes the source name, source type, entry type (valid/invalid),
// and an MD5 hash to ensure uniqueness.
//
// Parameters:
//   - name: Name of the source
//   - sourceType: Type of source (domain, ipv4, etc.)
//   - listType: Type of list ("blocklist" or "allowlist")
//   - entryType: Type of entries ("valid" or "invalid")
//
// Returns:
//   - A unique filename for the processed file
func generateFileName(logger *multilog.Logger, name, sourceType, listType, entryType string) string {
	hash := md5.Sum([]byte(name + sourceType + listType + entryType))
	lt := constants.ListTypeMap[strings.ToLower(listType)]
	if lt == "" {
		logger.Warnf("Unknown list type: %s", listType)
		lt = listType
	}
	return fmt.Sprintf("%s_%s_%s_%s_%s.txt", name, sourceType, lt, entryType, hex.EncodeToString(hash[:]))
}

// mergeSummaries merges two ProcessedSummary objects.
// It combines their valid and invalid files maps.
//
// Parameters:
//   - existingSummary: The summary to merge into
//   - newSummary: The summary to merge from
func mergeSummaries(existingSummary, newSummary *c.ProcessedSummary) {
	// Correctly merge valid with valid and invalid with invalid
	// merge the arrays
	existingSummary.Types = mergeSourceTypes(existingSummary.Types, newSummary.Types)
	existingSummary.ValidFiles = mergeProcessedFiles(existingSummary.ValidFiles, newSummary.ValidFiles)
	existingSummary.InvalidFiles = mergeProcessedFiles(existingSummary.InvalidFiles, newSummary.InvalidFiles)

	// Update the timestamp to the latest
	if newSummary.LastProcessedTimestamp > existingSummary.LastProcessedTimestamp {
		existingSummary.LastProcessedTimestamp = newSummary.LastProcessedTimestamp
	}
}

// mergeSourceTypes combines two slices of SourceType objects,
// avoiding duplicates by using the name as a unique identifier.
//
// Parameters:
//   - existing: The existing slice of SourceType objects
//   - new: The new slice of SourceType objects to merge in
//
// Returns:
//   - A merged slice of SourceType objects without duplicates
func mergeSourceTypes(existing, new []c.SourceType) []c.SourceType {
	// Create a map for a quick lookup of existing source types by name
	typeMap := make(map[string]c.SourceType)

	// Add all existing source types to the map
	for _, t := range existing {
		typeMap[t.Name] = t
	}

	// Add or overwrite with new source types
	for _, t := range new {
		typeMap[t.Name] = t
	}

	// Convert a map back to slice
	result := make([]c.SourceType, 0, len(typeMap))
	for _, t := range typeMap {
		result = append(result, t)
	}

	return result
}

// mergeProcessedFiles combines two slices of ProcessedFile objects,
// avoiding duplicates by using the filename as a unique identifier.
//
// Parameters:
//   - existing: The existing slice of ProcessedFile objects
//   - new: The new slice of ProcessedFile objects to merge in
//
// Returns:
//   - A merged slice of ProcessedFile objects without duplicates
func mergeProcessedFiles(existing, new []c.ProcessedFile) []c.ProcessedFile {
	// Create a map for a quick lookup of existing files by filename
	fileMap := make(map[string]c.ProcessedFile)

	// Add all existing files to the map
	for _, file := range existing {
		fileMap[file.Filepath] = file
	}

	// Add or overwrite with new files
	for _, file := range new {
		fileMap[file.Filepath] = file
	}

	// Convert a map back to slice
	result := make([]c.ProcessedFile, 0, len(fileMap))
	for _, file := range fileMap {
		result = append(result, file)
	}

	return result
}
