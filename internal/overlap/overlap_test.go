package overlap_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/overlap"
	"github.com/phani-kb/dns-toolkit/internal/overlap/mocks"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// createTestLogger creates a logger for testing
func createTestLogger() *multilog.Logger {
	return multilog.NewLogger()
}

// setupTestDirs creates temporary directories for testing
func setupTestDirs(t *testing.T) (string, string) {
	tmpDir := t.TempDir()
	overlapDir := filepath.Join(tmpDir, "overlap")
	summaryDir := filepath.Join(tmpDir, "summary")

	err := os.MkdirAll(overlapDir, 0755)
	require.NoError(t, err)

	err = os.MkdirAll(summaryDir, 0755)
	require.NoError(t, err)

	return overlapDir, summaryDir
}

// createTestFiles creates test files with content for testing overlap functions
func createTestFiles(t *testing.T, baseDir string) ([]common.ProcessedFile, []string) {
	processedDir := filepath.Join(baseDir, "processed")
	err := os.MkdirAll(processedDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create directory %s: %v", processedDir, err)
	}

	file1Path := filepath.Join(processedDir, "source1.txt")
	file1Content := "example.com\nmalicious.org\nads.com\ntracker.net\n"
	err = os.WriteFile(file1Path, []byte(file1Content), 0644)
	if err != nil {
		t.Fatalf("Failed to write file %s: %v", file1Path, err)
	}

	file2Path := filepath.Join(processedDir, "source2.txt")
	file2Content := "example.com\nphishing.org\ntracker.net\nspam.com\n"
	err = os.WriteFile(file2Path, []byte(file2Content), 0644)
	if err != nil {
		t.Fatalf("Failed to write file %s: %v", file2Path, err)
	}

	return []common.ProcessedFile{
		{
			Name:              "source1",
			GenericSourceType: "domain",
			ActualSourceType:  "domain",
			ListType:          constants.ListTypeBlocklist,
			Filepath:          file1Path,
			NumberOfEntries:   4,
			Valid:             true,
		},
		{
			Name:              "source2",
			GenericSourceType: "domain",
			ActualSourceType:  "domain",
			ListType:          constants.ListTypeAllowlist,
			Filepath:          file2Path,
			NumberOfEntries:   4,
			Valid:             true,
		},
	}, []string{"example.com", "tracker.net"}
}

// TestOverlapService_GetOverlapFilename tests the GetOverlapFilename method
func TestOverlapService_GetOverlapFilename(t *testing.T) {
	service := overlap.NewDefaultService("test_overlap_dir", "test_summary_dir")

	testCases := []struct {
		name         string
		name1        string
		listType1    string
		name2        string
		listType2    string
		sourceType   string
		expectedFile string
	}{
		{
			name:         "Both blocklists",
			name1:        "source1",
			listType1:    constants.ListTypeBlocklist,
			name2:        "source2",
			listType2:    constants.ListTypeBlocklist,
			sourceType:   "domain",
			expectedFile: "source1_source2_domain_BL_overlap.txt",
		},
		{
			name:         "Both allowlists",
			name1:        "source1",
			listType1:    constants.ListTypeAllowlist,
			name2:        "source2",
			listType2:    constants.ListTypeAllowlist,
			sourceType:   "domain",
			expectedFile: "source1_source2_domain_AL_overlap.txt",
		},
		{
			name:         "Mixed list types",
			name1:        "source1",
			listType1:    constants.ListTypeBlocklist,
			name2:        "source2",
			listType2:    constants.ListTypeAllowlist,
			sourceType:   "domain",
			expectedFile: "source1_source2_domain_mixed_overlap.txt",
		},
		{
			name:         "Different source type",
			name1:        "source1",
			listType1:    constants.ListTypeBlocklist,
			name2:        "source2",
			listType2:    constants.ListTypeBlocklist,
			sourceType:   "ipv4",
			expectedFile: "source1_source2_ipv4_BL_overlap.txt",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			filename := service.GetOverlapFilename(
				tc.name1,
				tc.listType1,
				tc.name2,
				tc.listType2,
				tc.sourceType,
			)
			assert.Equal(t, tc.expectedFile, filename)
		})
	}
}

// TestOverlapService_SaveOverlap tests the SaveOverlap method
func TestOverlapService_SaveOverlap(t *testing.T) {
	overlapDir, _ := setupTestDirs(t)
	service := overlap.NewDefaultService(overlapDir, "")
	logger := createTestLogger()

	t.Run("Save valid overlap", func(t *testing.T) {
		ol := []string{"example.com", "malicious.org"}
		filename := "test_overlap.txt"

		savedPath, err := service.SaveOverlap(logger, ol, filename)
		assert.NoError(t, err)
		assert.NotEmpty(t, savedPath)

		content, err := os.ReadFile(savedPath)
		assert.NoError(t, err)
		assert.Equal(t, "example.com\nmalicious.org", strings.TrimSpace(string(content)))
	})

	t.Run("Empty overlap should not create a file", func(t *testing.T) {
		var ol []string
		filename := "empty_overlap.txt"

		savedPath, err := service.SaveOverlap(logger, ol, filename)
		assert.NoError(t, err)
		assert.Empty(t, savedPath)

		expectedPath := filepath.Join(overlapDir, filename)
		_, err = os.Stat(expectedPath)
		assert.True(t, os.IsNotExist(err))
	})
}

// TestOverlapService_SaveOverlap_ErrorCase tests error handling in SaveOverlap
func TestOverlapService_SaveOverlap_ErrorCase(t *testing.T) {
	// Create a non-writable directory to force an error
	tmpDir := t.TempDir()
	nonWritableDir := filepath.Join(tmpDir, "non-writable")
	err := os.MkdirAll(nonWritableDir, 0755)
	require.NoError(t, err)

	err = os.Chmod(nonWritableDir, 0000)
	require.NoError(t, err)

	defer func(name string, mode os.FileMode) {
		err := os.Chmod(name, mode)
		if err != nil {
			t.Errorf("Failed to restore permissions for %s: %v", name, err)
		}
	}(nonWritableDir, 0755)

	service := overlap.NewDefaultService(nonWritableDir, "")
	logger := createTestLogger()

	ol := []string{"example.com", "malicious.org"}
	filename := "test_overlap.txt"

	savedPath, err := service.SaveOverlap(logger, ol, filename)

	if err == nil {
		assert.Empty(t, savedPath, "Path should be empty when SaveOverlap fails")
	} else {
		assert.Error(t, err, "Expected an error when saving to a non-writable directory")
	}
}

// TestOverlapService_FindOverlap tests the FindOverlap method with actual file processing
func TestOverlapService_FindOverlap(t *testing.T) {
	tmpDir := t.TempDir()
	overlapDir, summaryDir := setupTestDirs(t)
	service := overlap.NewDefaultService(overlapDir, summaryDir)
	logger := createTestLogger()

	processedFiles, _ := createTestFiles(t, tmpDir)

	result := service.FindOverlap(logger, "domain", processedFiles)

	assert.Equal(t, "domain", result.Type)
	assert.Greater(t, result.PairsCount, 0)
	assert.NotEmpty(t, result.Pairs)

	expectedPairsCount := len(processedFiles) * (len(processedFiles) - 1)
	assert.Equal(t, expectedPairsCount, result.PairsCount)

	var foundPair bool
	for _, pair := range result.Pairs {
		if pair.Source.Name == "source1" && pair.Target.Name == "source2" {
			foundPair = true
			assert.Equal(t, "domain", pair.Source.Type)
			assert.Equal(t, constants.ListTypeBlocklist, pair.Source.ListType)
			assert.Equal(t, "domain", pair.Target.Type)
			assert.Equal(t, constants.ListTypeAllowlist, pair.Target.ListType)
			assert.Equal(t, 2, pair.Overlap)
			assert.Contains(t, pair.Filepath, "source1_source2_domain_mixed_overlap.txt")
		}
	}
	assert.True(t, foundPair, "Expected pair not found in results")

	for _, pair := range result.Pairs {
		fileExists := false
		fullPath := filepath.Join(overlapDir, filepath.Base(pair.Filepath))
		_, err := os.Stat(fullPath)
		fileExists = !os.IsNotExist(err)
		assert.True(t, fileExists, "Overlap file not created: %s", pair.Filepath)
	}
}

// TestOverlapService_WriteCompactOverlapSummaries tests the WriteCompactOverlapSummaries method
func TestOverlapService_WriteCompactOverlapSummaries(t *testing.T) {
	tmpDir := t.TempDir()
	overlapDir, summaryDir := setupTestDirs(t)
	service := overlap.NewDefaultService(overlapDir, summaryDir)
	logger := createTestLogger()

	processedFiles, _ := createTestFiles(t, tmpDir)

	summaries, err := service.WriteCompactOverlapSummaries(
		logger,
		processedFiles,
		[]string{"domain"},
		4,
	)

	assert.NoError(t, err)
	assert.NotEmpty(t, summaries)

	sourceNames := make(map[string]bool)
	for _, summary := range summaries {
		sourceNames[summary.Source] = true
		assert.Equal(t, "domain", summary.Type)
		assert.Greater(t, summary.TargetsCount, 0)
		assert.NotEmpty(t, summary.TargetsList)
	}

	assert.True(t, sourceNames["source1"], "source1 not found in summaries")
	assert.True(t, sourceNames["source2"], "source2 not found in summaries")

	summaryFilePath := filepath.Join(summaryDir, constants.DefaultSummaryFiles["overlap"])
	_, err = os.Stat(summaryFilePath)
	assert.NoError(t, err, "Summary file was not created")
}

// TestOverlapService_WriteCompactOverlapSummaries_EmptySourceTypes tests WriteCompactOverlapSummaries with empty source types
func TestOverlapService_WriteCompactOverlapSummaries_EmptySourceTypes(t *testing.T) {
	tmpDir := t.TempDir()
	overlapDir, summaryDir := setupTestDirs(t)
	service := overlap.NewDefaultService(overlapDir, summaryDir)
	logger := createTestLogger()

	processedFiles, _ := createTestFiles(t, tmpDir)

	summaries, err := service.WriteCompactOverlapSummaries(
		logger,
		processedFiles,
		[]string{}, // Empty source types list
		4,
	)

	assert.NoError(t, err)
	assert.Empty(t, summaries)
}

// TestIsOverlapSourceTypeValid tests the IsOverlapSourceTypeValid function
func TestIsOverlapSourceTypeValid(t *testing.T) {
	testCases := []struct {
		name     string
		input    common.OverlapSourceType
		expected bool
	}{
		{
			name: "Valid overlap source type",
			input: common.OverlapSourceType{
				Type:       "test",
				PairsCount: 1,
				Pairs: []common.OverlapPair{
					{
						Source:  common.OverlapFileInfo{Name: "source1"},
						Target:  common.OverlapFileInfo{Name: "source2"},
						Overlap: 10,
					},
				},
			},
			expected: true,
		},
		{
			name: "Empty pairs",
			input: common.OverlapSourceType{
				Type:       "test",
				PairsCount: 0,
				Pairs:      []common.OverlapPair{},
			},
			expected: false,
		},
		{
			name: "Nil pairs",
			input: common.OverlapSourceType{
				Type:       "test",
				PairsCount: 0,
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := overlap.IsOverlapSourceTypeValid(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

// TestIsOverlapDetailedSummaryValid tests the IsOverlapDetailedSummaryValid function
func TestIsOverlapDetailedSummaryValid(t *testing.T) {
	validSourceType := common.OverlapSourceType{
		Type:       "test",
		PairsCount: 1,
		Pairs: []common.OverlapPair{
			{
				Source:  common.OverlapFileInfo{Name: "source1"},
				Target:  common.OverlapFileInfo{Name: "source2"},
				Overlap: 10,
			},
		},
	}

	testCases := []struct {
		name     string
		input    common.OverlapDetailedSummary
		expected bool
	}{
		{
			name: "Valid detailed summary",
			input: common.OverlapDetailedSummary{
				SourceTypes:      []common.OverlapSourceType{validSourceType},
				SourceTypesCount: 1,
			},
			expected: true,
		},
		{
			name: "Empty source types",
			input: common.OverlapDetailedSummary{
				SourceTypes:      []common.OverlapSourceType{},
				SourceTypesCount: 0,
			},
			expected: false,
		},
		{
			name: "Nil source types",
			input: common.OverlapDetailedSummary{
				SourceTypesCount: 0,
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := overlap.IsOverlapDetailedSummaryValid(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

// TestWithMockOverlapService tests using the mock FileOverlapService
func TestWithMockOverlapService(t *testing.T) {
	mockService := mocks.NewFileOverlapService(t)

	var logger *multilog.Logger

	mockService.On("GetOverlapFilename",
		"source1", constants.ListTypeBlocklist,
		"source2", constants.ListTypeBlocklist,
		"domain").Return("source1_source2_domain_BL_overlap.txt")

	mockService.On("GetOverlapFilename",
		"source1", constants.ListTypeBlocklist,
		"source2", constants.ListTypeAllowlist,
		"domain").Return("source1_source2_domain_mixed_overlap.txt")

	filename := mockService.GetOverlapFilename(
		"source1",
		constants.ListTypeBlocklist,
		"source2",
		constants.ListTypeBlocklist,
		"domain",
	)
	assert.Equal(t, "source1_source2_domain_BL_overlap.txt", filename)

	filename = mockService.GetOverlapFilename(
		"source1",
		constants.ListTypeBlocklist,
		"source2",
		constants.ListTypeAllowlist,
		"domain",
	)
	assert.Equal(t, "source1_source2_domain_mixed_overlap.txt", filename)

	mockService.On("SaveOverlap",
		logger,
		[]string{"example.com", "malicious.org"},
		"test_overlap.txt").Return("/path/to/test_overlap.txt", nil)

	savedPath, err := mockService.SaveOverlap(
		logger,
		[]string{"example.com", "malicious.org"},
		"test_overlap.txt",
	)
	assert.NoError(t, err)
	assert.Equal(t, "/path/to/test_overlap.txt", savedPath)

	processedFiles := []common.ProcessedFile{
		{
			Name:              "source1",
			GenericSourceType: "domain",
			ActualSourceType:  "domain",
			ListType:          constants.ListTypeBlocklist,
			Filepath:          "/path/to/source1.txt",
			NumberOfEntries:   100,
			Valid:             true,
		},
		{
			Name:              "source2",
			GenericSourceType: "domain",
			ActualSourceType:  "domain",
			ListType:          constants.ListTypeAllowlist,
			Filepath:          "/path/to/source2.txt",
			NumberOfEntries:   200,
			Valid:             true,
		},
	}

	expectedOverlapSourceType := common.OverlapSourceType{
		Type:       "domain",
		PairsCount: 1,
		Pairs: []common.OverlapPair{
			{
				Source: common.OverlapFileInfo{
					Name:     "source1",
					Type:     "domain",
					ListType: constants.ListTypeBlocklist,
					Count:    100,
					Percent:  "10.0",
				},
				Target: common.OverlapFileInfo{
					Name:     "source2",
					Type:     "domain",
					ListType: constants.ListTypeAllowlist,
					Count:    200,
					Percent:  "5.0",
				},
				Overlap:  10,
				Filepath: "source1_source2_domain_mixed_overlap.txt",
			},
		},
	}

	mockService.On("FindOverlap",
		logger,
		"domain",
		processedFiles).Return(expectedOverlapSourceType)

	result := mockService.FindOverlap(logger, "domain", processedFiles)
	assert.Equal(t, expectedOverlapSourceType.Type, result.Type)
	assert.Equal(t, expectedOverlapSourceType.PairsCount, result.PairsCount)
	assert.Equal(t, len(expectedOverlapSourceType.Pairs), len(result.Pairs))
	assert.Equal(t, expectedOverlapSourceType.Pairs[0].Source.Name, result.Pairs[0].Source.Name)

	expectedSummaries := []common.OverlapSummary{
		{
			Source:       "source1",
			ListType:     constants.ListTypeBlocklist,
			Type:         "domain",
			Count:        100,
			TargetsCount: 1,
			TargetsList:  []string{"source2, lt: allowlist, type: domain, count: 200, overlap: 10, percent: 5.0"},
		},
	}

	mockService.On("WriteCompactOverlapSummaries",
		logger,
		processedFiles,
		[]string{"domain"},
		4).Return(expectedSummaries, nil)

	summaries, err := mockService.WriteCompactOverlapSummaries(
		logger,
		processedFiles,
		[]string{"domain"},
		4,
	)
	assert.NoError(t, err)
	assert.Equal(t, len(expectedSummaries), len(summaries))
	assert.Equal(t, expectedSummaries[0].Source, summaries[0].Source)
	assert.Equal(t, expectedSummaries[0].ListType, summaries[0].ListType)
	assert.Equal(t, expectedSummaries[0].TargetsCount, summaries[0].TargetsCount)

	mockService.AssertExpectations(t)
}

// TestOverlapService_GetOverlapFilename_ExtendedCases tests additional cases for GetOverlapFilename
func TestOverlapService_GetOverlapFilename_ExtendedCases(t *testing.T) {
	service := overlap.NewDefaultService("test_overlap_dir", "test_summary_dir")

	filename := service.GetOverlapFilename(
		"source1",
		"custom-blocklist",
		"source2",
		"custom-blocklist",
		"domain",
	)
	assert.Equal(t, "source1_source2_domain_custom-blocklist_overlap.txt", filename)
}

// TestOverlapService_FindOverlap_InvalidFiles tests FindOverlap with invalid files
func TestOverlapService_FindOverlap_InvalidFiles(t *testing.T) {
	tmpDir := t.TempDir()
	overlapDir, summaryDir := setupTestDirs(t)
	service := overlap.NewDefaultService(overlapDir, summaryDir)
	logger := createTestLogger()

	processedFiles, _ := createTestFiles(t, tmpDir)

	invalidFile := common.ProcessedFile{
		Name:              "invalid_source",
		GenericSourceType: "domain",
		ActualSourceType:  "domain",
		ListType:          constants.ListTypeBlocklist,
		Filepath:          "/path/to/nonexistent/file.txt", // This file doesn't exist
		NumberOfEntries:   10,
		Valid:             false, // Marked as invalid
	}

	processedFiles = append(processedFiles, invalidFile)

	result := service.FindOverlap(logger, "domain", processedFiles)

	assert.Equal(t, "domain", result.Type)

	expectedPairsCount := 2 // For 2 files, we get 2 pairs (n*(n-1))
	assert.Equal(t, expectedPairsCount, result.PairsCount)
}

// TestOverlapService_WriteCompactOverlapSummaries_ErrorCase tests error handling in WriteCompactOverlapSummaries
func TestOverlapService_WriteCompactOverlapSummaries_ErrorCase(t *testing.T) {
	tmpDir := t.TempDir()

	nonWritableDir := filepath.Join(tmpDir, "non_writable_summary")
	err := os.MkdirAll(nonWritableDir, 0755)
	require.NoError(t, err)

	err = os.Chmod(nonWritableDir, 0000)
	require.NoError(t, err)

	defer func(name string, mode os.FileMode) {
		err := os.Chmod(name, mode)
		if err != nil {
			t.Errorf("Failed to restore permissions for %s: %v", name, err)
		}
	}(nonWritableDir, 0755)

	overlapDir := filepath.Join(tmpDir, "overlap")
	err = os.MkdirAll(overlapDir, 0755)
	require.NoError(t, err)

	service := overlap.NewDefaultService(overlapDir, nonWritableDir)
	logger := createTestLogger()

	processedFiles, _ := createTestFiles(t, tmpDir)

	_, err = service.WriteCompactOverlapSummaries(
		logger,
		processedFiles,
		[]string{"domain"},
		4,
	)

	if err != nil {
		assert.Error(t, err, "Expected an error when writing to a non-writable directory")
	}
}
