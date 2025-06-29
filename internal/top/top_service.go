package top

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"sync"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

// DefaultTopEntriesService is the default implementation of the EntriesService interface
type DefaultTopEntriesService struct {
	topDir     string
	summaryDir string
}

// NewDefaultService creates a new instance of the DefaultTopEntriesService
func NewDefaultService(topDir string, summaryDir string) *DefaultTopEntriesService {
	return &DefaultTopEntriesService{
		topDir:     topDir,
		summaryDir: summaryDir,
	}
}

// FindTopEntries processes files to find the top entries that appear in at least minSources
// sources for a specific generic source type and list type
func (s *DefaultTopEntriesService) FindTopEntries(
	logger *multilog.Logger,
	gst string,
	listType string,
	allProcessedFiles []common.ProcessedFile,
	minSources, maxEntries int,
	stringPool *utils.DTEntryPool,
) (common.TopSummary, error) {
	logger.Debugf("Starting findTopEntries for %s, %s, MinSources: %d", gst, listType, minSources)
	startTime := time.Now()

	var relevantFiles []common.ProcessedFile
	for _, pf := range allProcessedFiles {
		if pf.Valid && pf.GenericSourceType == gst && pf.ListType == listType {
			relevantFiles = append(relevantFiles, pf)
		}
	}

	if len(relevantFiles) == 0 {
		logger.Debugf("No relevant files found for %s, %s", gst, listType)
		return common.TopSummary{GenericSourceType: gst, ListType: listType, MinSources: minSources, Count: 0}, nil
	}
	logger.Debugf("Found %d relevant file(s) for %s (%s)", len(relevantFiles), gst, listType)

	sourceNameToID := make(map[string]uint16)
	sourceIDToName := make([]string, 0)
	for _, pf := range relevantFiles {
		internedName := stringPool.Intern(pf.Name)
		if _, exists := sourceNameToID[internedName]; !exists {
			newID := uint16(len(sourceIDToName))
			sourceNameToID[internedName] = newID
			sourceIDToName = append(sourceIDToName, internedName)
		}
	}
	logger.Debugf("Created source name to ID mapping with %d unique sources", len(sourceIDToName))

	entrySources := make(map[string]map[uint16]struct{})
	var mapMutex sync.Mutex

	fileProcessingWorkers := runtime.GOMAXPROCS(0)

	processFileFunc := func(processedFile common.ProcessedFile, pool *utils.DTEntryPool) {
		// Intern the source name once per file, not per line
		internedSourceNameFromFile := pool.Intern(processedFile.Name)
		sourceID, ok := sourceNameToID[internedSourceNameFromFile]
		if !ok {
			logger.Errorf(
				"Logic error: interned source name %s not found in sourceNameToID map during setup for file %s",
				internedSourceNameFromFile,
				processedFile.Filepath,
			)
			return
		}

		file, err := os.Open(processedFile.Filepath)
		if err != nil {
			logger.Errorf("Error opening file %s: %v", processedFile.Filepath, err)
			return
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				logger.Errorf("Error closing file %s: %v", processedFile.Filepath, err)
			}
		}(file)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lineBytes := scanner.Bytes()

			// Trim leading space
			startIdx := 0
			for startIdx < len(lineBytes) {
				r, size := utf8.DecodeRune(lineBytes[startIdx:])
				if !unicode.IsSpace(r) {
					break
				}
				startIdx += size
			}

			// Trim trailing space
			endIdx := len(lineBytes)
			for endIdx > startIdx {
				r, size := utf8.DecodeLastRune(lineBytes[startIdx:endIdx])
				if !unicode.IsSpace(r) {
					break
				}
				endIdx -= size
			}

			if startIdx >= endIdx {
				continue
			}

			// string() makes a copy, which is necessary because lineBytes is overwritten by the next line
			entry := pool.Intern(string(lineBytes[startIdx:endIdx]))

			mapMutex.Lock()
			if _, ok := entrySources[entry]; !ok {
				entrySources[entry] = make(map[uint16]struct{})
			}
			entrySources[entry][sourceID] = struct{}{}
			mapMutex.Unlock()
		}
		if err := scanner.Err(); err != nil {
			logger.Errorf("Error scanning file %s: %v", processedFile.Filepath, err)
		}
	}

	if len(relevantFiles) > constants.MinFilesForParallelProcessing && fileProcessingWorkers > 1 {
		filePool := common.NewDTWorkerPool(fileProcessingWorkers)
		logger.Debugf(
			"Processing %d files in parallel for %s (%s) using %d workers",
			len(relevantFiles),
			gst,
			listType,
			fileProcessingWorkers,
		)
		for _, pf := range relevantFiles {
			currentFile := pf
			filePool.Submit(func() {
				processFileFunc(currentFile, stringPool)
			})
		}
		filePool.Wait()
	} else {
		logger.Debugf("Processing %d files sequentially for %s (%s)", len(relevantFiles), gst, listType)
		for _, pf := range relevantFiles {
			processFileFunc(pf, stringPool)
		}
	}

	logger.Debugf(
		"Finished populating entrySources for %s (%s). Unique entries: %d. Time taken: %s",
		gst,
		listType,
		len(entrySources),
		time.Since(startTime),
	)

	topNEntries := s.GetTopNEntries(entrySources, minSources, maxEntries)
	entrySources = nil
	if len(relevantFiles) > 0 {
		runtime.GC()
		debug.FreeOSMemory()
	}

	if len(topNEntries) == 0 {
		logger.Debugf("No top entry found for %s (%s) with min %d after filtering.", gst, listType, minSources)
		return common.TopSummary{GenericSourceType: gst, ListType: listType, MinSources: minSources, Count: 0}, nil
	}

	fp, err := s.SaveTopEntries(logger, gst, listType, minSources, topNEntries)
	if err != nil {
		return common.TopSummary{}, fmt.Errorf("saving top entry(s) for %s (%s, min %d): %w", gst,
			listType,
			minSources,
			err,
		)
	}
	logger.Debugf(
		"Finished saving top entry(s) for %s (%s) to %s. Time taken: %s",
		gst,
		listType,
		fp,
		time.Since(startTime),
	)

	summary := common.TopSummary{
		GenericSourceType: gst,
		ListType:          listType,
		MinSources:        minSources,
		Filepath:          fp,
		TopEntries:        topNEntries,
		Count:             len(topNEntries),
	}
	logger.Infof(
		"%d top entry(s) for %s (%s, min %d). Saved to %s. Time taken: %s",
		summary.Count,
		gst,
		listType,
		minSources,
		fp,
		time.Since(startTime),
	)
	return summary, nil
}

// ProcessTopEntries processes all processed files to generate top entries for different
// combinations of generic source types, list types and minimum sources thresholds
func (s *DefaultTopEntriesService) ProcessTopEntries(
	logger *multilog.Logger,
	genericSourceTypes []string,
	processedFiles []common.ProcessedFile,
	minSourcesValues []int,
	maxEntries int,
	maxWorkers int,
) ([]common.TopSummary, error) {
	var mu sync.Mutex
	topSummaries := make([]common.TopSummary, 0)

	workerPool := common.NewDTWorkerPool(maxWorkers)

	stringPool := utils.NewDTEntryPool()

	for _, gst := range genericSourceTypes {
		for _, listType := range constants.ListTypes {
			for _, minSrc := range minSourcesValues {
				currentGst, currentListType, currentMinSrc := gst, listType, minSrc
				workerPool.Submit(func() {
					topSummary, err := s.FindTopEntries(
						logger,
						currentGst,
						currentListType,
						processedFiles,
						currentMinSrc,
						maxEntries,
						stringPool,
					)
					if err != nil {
						logger.Errorf(
							"Error finding top entry(s) for %s (%s) with minSources %d: %v",
							currentGst,
							currentListType,
							currentMinSrc,
							err,
						)
						return
					}

					if topSummary.Count > 0 {
						mu.Lock()
						topSummaries = append(topSummaries, topSummary)
						mu.Unlock()
					}
				})
			}
		}
		utils.LogMemStats(logger, "After processing "+gst)
	}
	workerPool.Wait()

	filteredSummaries := s.FilterTopSummaries(topSummaries)

	_, err := s.SaveTopSummaries(logger, filteredSummaries)
	if err != nil {
		return filteredSummaries, err
	}

	return filteredSummaries, nil
}

// GetTopNEntries returns the top N entries that appear in at least minSources sources
func (s *DefaultTopEntriesService) GetTopNEntries(
	entrySources map[string]map[uint16]struct{},
	minSources int,
	maxEntries int,
) []common.EntryCountPair {
	if maxEntries <= 0 {
		return []common.EntryCountPair{}
	}

	h := &common.EntryHeap{}
	heap.Init(h)

	for entry, sources := range entrySources {
		sourceCount := len(sources)
		if sourceCount >= minSources {
			heap.Push(h, common.EntryCountPair{Entry: entry, Count: sourceCount})
			if h.Len() > maxEntries {
				heap.Pop(h)
			}
		}
	}

	result := make([]common.EntryCountPair, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		if ecp, ok := heap.Pop(h).(common.EntryCountPair); ok {
			result[i] = ecp
		}
	}
	return result
}

// SaveTopEntries saves the top entries to a file and returns the filepath
func (s *DefaultTopEntriesService) SaveTopEntries(
	logger *multilog.Logger,
	gst, listType string,
	minSources int,
	entries []common.EntryCountPair,
) (string, error) {
	if len(entries) == 0 {
		logger.Debugf("No entry to save for %s (%s) with minSources %d", gst, listType, minSources)
		return "", nil
	}

	fileName := fmt.Sprintf("top_%s_%s_min%d.txt", gst, listType, minSources)
	filePath := filepath.Join(s.topDir, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create top entry file %s: %w", filePath, err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.Errorf("Error closing file %s: %v", filePath, err)
		}
	}(file)

	writer := bufio.NewWriter(file)
	for _, pair := range entries {
		if _, err := fmt.Fprintf(writer, "%s\n", pair.Entry); err != nil {
			return "", fmt.Errorf("failed to write entry to %s: %w", filePath, err)
		}
	}

	if err := writer.Flush(); err != nil {
		return "", fmt.Errorf("failed to flush writer for %s: %w", filePath, err)
	}
	logger.Debugf("Saved %d top entry(s) to %s", len(entries), filePath)
	return filePath, nil
}

// FilterTopSummaries removes empty or invalid top summaries
func (s *DefaultTopEntriesService) FilterTopSummaries(topSummaries []common.TopSummary) []common.TopSummary {
	var filteredSummaries []common.TopSummary
	for _, summary := range topSummaries {
		if summary.Count > 0 && len(summary.TopEntries) > 0 {
			filteredSummaries = append(filteredSummaries, summary)
		}
	}
	return filteredSummaries
}

// SaveTopSummaries saves the top summaries to the summary file
func (s *DefaultTopEntriesService) SaveTopSummaries(
	logger *multilog.Logger,
	topSummaries []common.TopSummary,
) (int, error) {
	summariesByListType := make(map[string][]common.TopSummary)
	for _, summary := range topSummaries {
		if summary.ListType == "" {
			summary.ListType = constants.ListTypeBlocklist
		}
		summariesByListType[summary.ListType] = append(summariesByListType[summary.ListType], summary)
	}

	summaryFile := filepath.Join(s.summaryDir, constants.DefaultSummaryFiles["top"])
	summariesCount, err := utils.SaveSummaries(logger, topSummaries, summaryFile, common.TopSummaryLessFunc)
	if err != nil {
		logger.Errorf("Error saving top summaries: %v", err)
		return 0, err
	}
	if summariesCount > 0 {
		logger.Infof("Top summaries saved to %s", summaryFile)
	}

	return summariesCount, nil
}
