package overlap

import (
	"fmt"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"sort"
	"strings"
	"sync"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

// DefaultOverlapService is the standard implementation of the overlap Service interface
type DefaultOverlapService struct {
	// OverlapDir is the directory where overlap files will be stored
	OverlapDir string

	// SummaryDir is the directory where summary files will be stored
	SummaryDir string
}

// NewDefaultService creates a new instance of DefaultService with the given directories
func NewDefaultService(overlapDir, summaryDir string) *DefaultOverlapService {
	return &DefaultOverlapService{
		OverlapDir: overlapDir,
		SummaryDir: summaryDir,
	}
}

// FindOverlap processes a collection of files of the same generic source type to identify overlaps
func (s *DefaultOverlapService) FindOverlap(
	logger *multilog.Logger,
	genericSourceType string,
	files []c.ProcessedFile,
) c.OverlapSourceType {
	stringPool := u.NewDTEntryPool() // Initialize the string pool for this specific genericSourceType processing
	iGenericSourceType := stringPool.Intern(genericSourceType)

	logger.Infof("Finding overlap for %s", iGenericSourceType)

	validFiles := make(map[string]string)        // Filepath -> Interned Name
	listTypes := make(map[string]string)         // Filepath -> Interned ListType
	actualSourceTypes := make(map[string]string) // Filepath -> Interned ActualSourceType

	for _, processedFile := range files {
		if !processedFile.Valid {
			continue
		}
		// Intern strings that will be used repeatedly
		validFiles[processedFile.Filepath] = stringPool.Intern(processedFile.Name)
		actualSourceTypes[processedFile.Filepath] = stringPool.Intern(processedFile.ActualSourceType)
		listTypes[processedFile.Filepath] = stringPool.Intern(processedFile.ListType)
	}

	logger.Infof("Valid Processed files count: %d for %s", len(validFiles), iGenericSourceType)

	overlapSourceType := c.OverlapSourceType{
		Type:  iGenericSourceType, // Use interned string
		Pairs: make([]c.OverlapPair, 0),
	}

	var mu sync.Mutex
	maxWorkers := runtime.GOMAXPROCS(0)
	pairProcessingPool := c.NewDTWorkerPool(maxWorkers) // Worker pool for processing file pairs

	type filePair struct {
		file1, file2 string
		name1, name2 string
	}

	var pairs []filePair
	// Create pairs to process using interned names
	var filePaths []string
	for fp := range validFiles {
		filePaths = append(filePaths, fp)
	}
	sort.Strings(filePaths)

	for i := 0; i < len(filePaths); i++ {
		for j := i + 1; j < len(filePaths); j++ {
			file1 := filePaths[i]
			file2 := filePaths[j]
			pairs = append(pairs, filePair{file1, file2, validFiles[file1], validFiles[file2]})
		}
	}

	logger.Infof("Processing %d file pairs for overlap calculation for %s", len(pairs), iGenericSourceType)

	for _, pair := range pairs {
		currentPair := pair
		pairProcessingPool.Submit(func() {
			file1, file2 := currentPair.file1, currentPair.file2
			name1, name2 := currentPair.name1, currentPair.name2

			logger.Debugf("Finding overlap between %s (%s) and %s (%s)", name1, file1, name2, file2)

			rawOverlap, sourceCount, targetCount := u.FindOverlap(
				logger,
				file1,
				file2,
			)

			internedOverlap := make([]string, len(rawOverlap))
			for k, s := range rawOverlap {
				internedOverlap[k] = stringPool.Intern(s)
			}

			logger.Debugf("Overlap found: %d Source: %d Target: %d", len(internedOverlap), sourceCount, targetCount)

			// Intern list types and generic source type for filename generation
			internedListType1 := listTypes[file1]
			internedListType2 := listTypes[file2]

			filename := s.GetOverlapFilename(
				name1,
				internedListType1,
				name2,
				internedListType2,
				iGenericSourceType,
			)
			_, err := s.SaveOverlap(logger, internedOverlap, filename)
			if err != nil {
				logger.Errorf("Error saving overlap file %s: %v", filename, err)
				return
			}

			if len(internedOverlap) == 0 {
				return
			}

			// Intern types for OverlapFileInfo
			internedActualSourceType1 := actualSourceTypes[file1]
			internedActualSourceType2 := actualSourceTypes[file2]

			overlapPair1 := c.OverlapPair{
				Source: c.OverlapFileInfo{
					Filename: file1,
					Name:     name1,
					Type:     internedActualSourceType1,
					ListType: internedListType1,
					Count:    sourceCount,
					Percent:  fmt.Sprintf("%.1f", float64(len(internedOverlap))/float64(sourceCount)*100),
				},
				Target: c.OverlapFileInfo{
					Filename: file2,
					Name:     name2,
					Type:     internedActualSourceType2,
					ListType: internedListType2,
					Count:    targetCount,
					Percent:  fmt.Sprintf("%.1f", float64(len(internedOverlap))/float64(targetCount)*100),
				},
				Overlap:  len(internedOverlap),
				Filepath: filename,
			}

			overlapPair2 := c.OverlapPair{
				Source: c.OverlapFileInfo{
					Filename: file2,
					Name:     name2,
					Type:     internedActualSourceType2,
					ListType: internedListType2,
					Count:    targetCount,
					Percent:  fmt.Sprintf("%.1f", float64(len(internedOverlap))/float64(targetCount)*100),
				},
				Target: c.OverlapFileInfo{
					Filename: file1,
					Name:     name1,
					Type:     internedActualSourceType1,
					ListType: internedListType1,
					Count:    sourceCount,
					Percent:  fmt.Sprintf("%.1f", float64(len(internedOverlap))/float64(sourceCount)*100),
				},
				Overlap:  len(internedOverlap),
				Filepath: filename,
			}

			mu.Lock()
			overlapSourceType.Pairs = append(overlapSourceType.Pairs, overlapPair1, overlapPair2)
			overlapSourceType.PairsCount = len(overlapSourceType.Pairs)
			mu.Unlock()
		})
	}

	pairProcessingPool.Wait()
	// Free memory after processing all pairs for this genericSourceType
	logger.Infof("Processed all %d pairs for %s.", len(pairs), iGenericSourceType)
	runtime.GC()
	debug.FreeOSMemory()
	logger.Debugf("Memory freed after processing pairs for %s", iGenericSourceType)

	logger.Infof("Overlap found for %s: %d pairs", iGenericSourceType, len(overlapSourceType.Pairs))
	return overlapSourceType
}

// WriteCompactOverlapSummaries generates and writes overlap summaries from processed files
func (s *DefaultOverlapService) WriteCompactOverlapSummaries(
	logger *multilog.Logger,
	processedFiles []c.ProcessedFile,
	genericSourceTypes []string,
	maxWorkers int,
) ([]c.OverlapSummary, error) {
	var overlapSources []c.OverlapSourceType
	var mu sync.Mutex

	workerPool := c.NewDTWorkerPool(maxWorkers)

	processedFilesMap := make(map[string][]c.ProcessedFile)
	// group processed files by generic source type
	for _, file := range processedFiles {
		processedFilesMap[file.GenericSourceType] = append(processedFilesMap[file.GenericSourceType], file)
	}

	for _, genericSourceType := range genericSourceTypes {
		currentGenericSourceType := genericSourceType
		workerPool.Submit(func() {
			// Each call to findOverlap will manage its own string pool and memory internally
			overlapSourceType := s.FindOverlap(
				logger,
				currentGenericSourceType,
				processedFilesMap[currentGenericSourceType],
			)
			if len(overlapSourceType.Pairs) > 0 { // Only append if there's actual overlap data
				mu.Lock()
				overlapSources = append(overlapSources, overlapSourceType)
				mu.Unlock()
			}
		})
	}
	workerPool.Wait()
	u.LogMemStats(logger, "After overlap calculation")
	runtime.GC()
	debug.FreeOSMemory()
	logger.Debugf("Memory freed after all generic source type overlap calculations.")

	sourceTypeMap := make(map[string]map[string][]c.OverlapPair)

	// Group pairs by source type and source name
	for _, sourceType := range overlapSources {
		if _, exists := sourceTypeMap[sourceType.Type]; !exists {
			sourceTypeMap[sourceType.Type] = make(map[string][]c.OverlapPair)
		}

		for _, pair := range sourceType.Pairs {
			sourceName := pair.Source.Name
			sourceTypeMap[sourceType.Type][sourceName] = append(
				sourceTypeMap[sourceType.Type][sourceName],
				pair,
			)
		}
	}

	compactSummaries := make([]c.OverlapSummary, 0)

	// Process each source type
	for sourceTypeName, sourceNameMap := range sourceTypeMap {
		// Process each source name within this type
		for sourceName, pairs := range sourceNameMap {
			cs := computeCompactSummaryFromPairs(sourceTypeName, sourceName, pairs)
			if cs == nil {
				continue
			}
			logger.Infof("Source: %s targets count: %d for %s", cs.Source, cs.TargetsCount, sourceTypeName)
			compactSummaries = append(compactSummaries, *cs)
		}
		logger.Infof("Overlap summaries count: %d for %s", len(compactSummaries), sourceTypeName)
	}
	summaryFile := filepath.Join(s.SummaryDir, constants.DefaultSummaryFiles["overlap"])
	savedCount, err := u.SaveSummaries(logger, compactSummaries, summaryFile, c.OverlapSummaryLessFunc)
	if err != nil {
		logger.Errorf("Saving summaries error: %v", err)
		return compactSummaries, err
	}

	logger.Infof("Saved %d overlap summaries to %s", savedCount, summaryFile)
	return compactSummaries, nil
}

// computeCompactSummaryFromPairs builds a compact OverlapSummary for a given source from its pairs.
// It deduplicates targets by name, keeps the highest overlap per target name, and computes Unique
// as Count - sum(deduplicated overlaps).
func computeCompactSummaryFromPairs(sourceTypeName, sourceName string, pairs []c.OverlapPair) *c.OverlapSummary {
	if len(pairs) == 0 {
		return nil
	}
	firstPair := pairs[0]
	cs := &c.OverlapSummary{
		Type:        sourceTypeName,
		Source:      sourceName,
		ListType:    firstPair.Source.ListType,
		Count:       firstPair.Source.Count,
		TargetsList: make([]string, 0),
		Targets:     make([]c.OverlapTargetFileInfo, 0),
	}

	targetMap := make(map[string]c.OverlapTargetFileInfo)

	for _, targetPair := range pairs {
		if targetPair.Source.Name != sourceName {
			continue
		}
		targetFile := targetPair.Target
		if targetPair.Overlap <= 0 {
			continue
		}
		ot := c.OverlapTargetFileInfo{
			Name:     targetFile.Name,
			ListType: targetFile.ListType,
			Type:     targetFile.Type,
			Count:    targetFile.Count,
			Percent:  targetFile.Percent,
			Overlap:  targetPair.Overlap,
		}

		if existing, exists := targetMap[targetFile.Name]; !exists || ot.Overlap > existing.Overlap {
			targetMap[targetFile.Name] = ot
		}
	}

	for _, target := range targetMap {
		cs.Targets = append(cs.Targets, target)
	}

	if len(cs.Targets) == 0 {
		return nil
	}

	sort.Slice(cs.Targets, func(i, j int) bool {
		return u.ParsePercent(cs.Targets[i].Percent) > u.ParsePercent(cs.Targets[j].Percent)
	})

	for _, target := range cs.Targets {
		cs.TargetsList = append(cs.TargetsList, target.GetString())
	}

	cs.TargetsCount = len(cs.Targets)

	totalOverlapCount := 0
	for _, t := range cs.Targets {
		totalOverlapCount += t.Overlap
	}

	uniqueEntries := cs.Count - totalOverlapCount
	if uniqueEntries < 0 {
		uniqueEntries = 0
	}
	cs.Unique = uniqueEntries

	return cs
}

// GetOverlapFilename generates a filename for an overlap file based on source and target file information
func (s *DefaultOverlapService) GetOverlapFilename(name1, listType1, name2, listType2, sourceType string) string {
	var listType string

	if val, ok := constants.ListTypeMap[listType1]; ok {
		listType = val
	} else {
		// If it's not in the map, use the string directly
		listType = listType1
	}

	if listType1 != listType2 {
		listType = "mixed"
	}

	return fmt.Sprintf("%s_%s_%s_%s_overlap.txt", name1, name2, sourceType, listType)
}

// SaveOverlap saves overlap entries to a file
func (s *DefaultOverlapService) SaveOverlap(
	logger *multilog.Logger,
	overlap []string,
	filename string,
) (string, error) {
	if len(overlap) == 0 {
		logger.Debugf("No overlap found, skipping file: %s", filename)
		return "", nil
	}

	logger.Debugf("Saving overlap file: %s", filename)
	overlapReader := strings.NewReader(strings.Join(overlap, "\n"))
	fp, err := u.SaveFile(logger, s.OverlapDir, filename, overlapReader)
	if err != nil {
		logger.Errorf("Error saving overlap file: %s", err)
		return "", err
	}

	logger.Debugf("Overlap saved to %s", fp)
	return fp, nil
}
