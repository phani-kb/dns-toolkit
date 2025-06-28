package cmd

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/multilog"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestMergeSummaries(t *testing.T) {
	t.Parallel()

	existing := &c.ProcessedSummary{
		Types: []c.SourceType{
			{
				Name: "domain",
			},
		},
		ValidFiles: []c.ProcessedFile{
			{
				Name:              "test1",
				GenericSourceType: "domain",
				ListType:          "blocklist",
				Filepath:          "/path/to/test1.txt",
				NumberOfEntries:   5,
				Valid:             true,
			},
		},
		InvalidFiles: []c.ProcessedFile{
			{
				Name:              "test1",
				GenericSourceType: "domain",
				ListType:          "blocklist",
				Filepath:          "/path/to/invalid1.txt",
				NumberOfEntries:   2,
				Valid:             false,
			},
		},
		LastProcessedTimestamp: "2023-01-01T00:00:00Z",
	}

	newProcessedSummary := &c.ProcessedSummary{
		Types: []c.SourceType{
			{
				Name: "ipv4",
			},
		},
		ValidFiles: []c.ProcessedFile{
			{
				Name:              "test2",
				GenericSourceType: "ipv4",
				ListType:          "allowlist",
				Filepath:          "/path/to/test2.txt",
				NumberOfEntries:   8,
				Valid:             true,
			},
		},
		InvalidFiles: []c.ProcessedFile{
			{
				Name:              "test2",
				GenericSourceType: "ipv4",
				ListType:          "allowlist",
				Filepath:          "/path/to/invalid2.txt",
				NumberOfEntries:   3,
				Valid:             false,
			},
		},
		LastProcessedTimestamp: "2023-01-02T00:00:00Z",
	}

	mergeSummaries(existing, newProcessedSummary)

	assert.Len(t, existing.Types, 2, "Should have 2 source types after merge")
	assert.Len(t, existing.ValidFiles, 2, "Should have 2 valid files after merge")
	assert.Len(t, existing.InvalidFiles, 2, "Should have 2 invalid files after merge")
	assert.Equal(t, "2023-01-02T00:00:00Z", existing.LastProcessedTimestamp, "Timestamp should be updated to latest")
}

func TestMergeSourceTypes(t *testing.T) {
	t.Parallel()

	existing := []c.SourceType{
		{
			Name: "domain",
		},
	}

	sourceTypes := []c.SourceType{
		{
			Name: "ipv4",
		},
		{
			Name: "domain",
		},
	}

	result := mergeSourceTypes(existing, sourceTypes)

	assert.Len(t, result, 2, "Should have 2 source types after merge")

	typeNames := make([]string, len(result))
	for i, st := range result {
		typeNames[i] = st.Name
	}
	assert.Contains(t, typeNames, "domain")
	assert.Contains(t, typeNames, "ipv4")
}

func TestMergeProcessedFiles(t *testing.T) {
	t.Parallel()

	existing := []c.ProcessedFile{
		{
			Name:              "test1",
			GenericSourceType: "domain",
			ListType:          "blocklist",
			Filepath:          "/path/to/test1.txt",
			NumberOfEntries:   5,
			Valid:             true,
		},
	}

	processedFiles := []c.ProcessedFile{
		{
			Name:              "test2",
			GenericSourceType: "ipv4",
			ListType:          "allowlist",
			Filepath:          "/path/to/test2.txt",
			NumberOfEntries:   8,
			Valid:             true,
		},
		{
			Name:              "test1",
			GenericSourceType: "domain",
			ListType:          "blocklist",
			Filepath:          "/path/to/test1_updated.txt",
			NumberOfEntries:   3,
			Valid:             false,
		},
	}

	result := mergeProcessedFiles(existing, processedFiles)

	assert.Len(t, result, 3, "Should have 3 processed files after merge")

	filepaths := make([]string, len(result))
	for i, pf := range result {
		filepaths[i] = pf.Filepath
	}
	assert.Contains(t, filepaths, "/path/to/test1.txt")
	assert.Contains(t, filepaths, "/path/to/test2.txt")
	assert.Contains(t, filepaths, "/path/to/test1_updated.txt")
}

func TestCreateSummary(t *testing.T) {
	tests := []struct {
		name           string
		sourceName     string
		sourceTypes    []c.SourceType
		validFiles     []c.ProcessedFile
		invalidFiles   []c.ProcessedFile
		expectedFields map[string]interface{}
	}{
		{
			name:       "Basic summary creation",
			sourceName: "test-source",
			sourceTypes: []c.SourceType{
				{
					Name: "domain",
					ListTypes: []c.ListType{
						{Name: "blocklist", Disabled: false},
					},
				},
			},
			validFiles: []c.ProcessedFile{
				{
					Name:              "test-source",
					GenericSourceType: "domain",
					ListType:          "blocklist",
					Filepath:          "/path/to/test.txt",
					NumberOfEntries:   950,
					Valid:             true,
				},
			},
			invalidFiles: []c.ProcessedFile{},
			expectedFields: map[string]interface{}{
				"source_name": "test-source",
			},
		},
		{
			name:       "Empty summary creation",
			sourceName: "empty-source",
			sourceTypes: []c.SourceType{
				{
					Name: "domain",
					ListTypes: []c.ListType{
						{Name: "blocklist", Disabled: false},
					},
				},
			},
			validFiles:   []c.ProcessedFile{},
			invalidFiles: []c.ProcessedFile{},
			expectedFields: map[string]interface{}{
				"source_name": "empty-source",
			},
		},
		{
			name:       "Large numbers",
			sourceName: "large-source",
			sourceTypes: []c.SourceType{
				{
					Name: "domain",
					ListTypes: []c.ListType{
						{Name: "blocklist", Disabled: false},
					},
				},
				{
					Name: "ipv4",
					ListTypes: []c.ListType{
						{Name: "blocklist", Disabled: false},
					},
				},
			},
			validFiles: []c.ProcessedFile{
				{
					Name:              "large-source",
					GenericSourceType: "domain",
					ListType:          "blocklist",
					Filepath:          "/path/to/large.txt",
					NumberOfEntries:   999500,
					Valid:             true,
				},
			},
			invalidFiles: []c.ProcessedFile{},
			expectedFields: map[string]interface{}{
				"source_name": "large-source",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			summary := createSummary(tt.sourceName, tt.sourceTypes, tt.validFiles, tt.invalidFiles)

			if summary.Name == "" {
				t.Fatal("createSummary returned empty summary")
			}

			if summary.LastProcessedTimestamp == "" {
				t.Error("LastProcessedTimestamp should not be empty")
			}

			_, err := time.Parse(constants.TimestampFormat, summary.LastProcessedTimestamp)
			if err != nil {
				t.Errorf("Invalid timestamp format: %v", err)
			}

			if summary.Name != tt.sourceName {
				t.Errorf("Expected source name %s, got %s", tt.sourceName, summary.Name)
			}
			if len(summary.Types) != len(tt.sourceTypes) {
				t.Errorf("Expected %d source types, got %d", len(tt.sourceTypes), len(summary.Types))
			}
			if len(summary.ValidFiles) != len(tt.validFiles) {
				t.Errorf("Expected %d valid files, got %d", len(tt.validFiles), len(summary.ValidFiles))
			}
			if len(summary.InvalidFiles) != len(tt.invalidFiles) {
				t.Errorf("Expected %d invalid files, got %d", len(tt.invalidFiles), len(summary.InvalidFiles))
			}
		})
	}
}

func TestSaveEntries(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)

	tempDir, err := os.MkdirTemp("", "save_entries_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tests := []struct {
		name              string
		validEntries      []string
		invalidEntries    []string
		processedDir      string
		sourceName        string
		sourceType        string
		listType          string
		expectValidFile   bool
		expectInvalidFile bool
	}{
		{
			name:              "Save valid and invalid entries",
			validEntries:      []string{"example.com", "test.org", "sample.net"},
			invalidEntries:    []string{"invalid1", "invalid2"},
			processedDir:      tempDir,
			sourceName:        "test-source",
			sourceType:        "domain",
			listType:          "blocklist",
			expectValidFile:   true,
			expectInvalidFile: true,
		},
		{
			name:              "Save only valid entries",
			validEntries:      []string{"valid.com"},
			invalidEntries:    []string{},
			processedDir:      tempDir,
			sourceName:        "valid-source",
			sourceType:        "domain",
			listType:          "blocklist",
			expectValidFile:   true,
			expectInvalidFile: false,
		},
		{
			name:              "Save only invalid entries",
			validEntries:      []string{},
			invalidEntries:    []string{"invalid.com"},
			processedDir:      tempDir,
			sourceName:        "invalid-source",
			sourceType:        "domain",
			listType:          "blocklist",
			expectValidFile:   false,
			expectInvalidFile: true,
		},
		{
			name:              "Save empty entries",
			validEntries:      []string{},
			invalidEntries:    []string{},
			processedDir:      tempDir,
			sourceName:        "empty-source",
			sourceType:        "domain",
			listType:          "blocklist",
			expectValidFile:   false,
			expectInvalidFile: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validFilePath, invalidFilePath := saveEntries(logger, tt.processedDir, tt.sourceName,
				tt.sourceType, tt.listType, tt.validEntries, tt.invalidEntries)

			if tt.expectValidFile {
				if validFilePath == "" {
					t.Error("Expected valid file path but got empty string")
				} else if _, err := os.Stat(validFilePath); os.IsNotExist(err) {
					t.Errorf("Expected valid file to exist at %s but it doesn't", validFilePath)
				}
			} else {
				if validFilePath != "" {
					t.Errorf("Expected no valid file but got path: %s", validFilePath)
				}
			}

			if tt.expectInvalidFile {
				if invalidFilePath == "" {
					t.Error("Expected invalid file path but got empty string")
				} else if _, err := os.Stat(invalidFilePath); os.IsNotExist(err) {
					t.Errorf("Expected invalid file to exist at %s but it doesn't", invalidFilePath)
				}
			} else {
				if invalidFilePath != "" {
					t.Errorf("Expected no invalid file but got path: %s", invalidFilePath)
				}
			}

			if tt.expectValidFile && validFilePath != "" {
				content, err := os.ReadFile(validFilePath)
				if err != nil {
					t.Errorf("Failed to read valid file: %v", err)
				} else {
					lines := strings.Split(strings.TrimSpace(string(content)), "\n")
					// Remove empty lines
					var nonEmptyLines []string
					for _, line := range lines {
						if strings.TrimSpace(line) != "" {
							nonEmptyLines = append(nonEmptyLines, line)
						}
					}
					if len(nonEmptyLines) != len(tt.validEntries) {
						t.Errorf("Expected %d valid entries in file, got %d", len(tt.validEntries), len(nonEmptyLines))
					}
				}
			}

			if tt.expectInvalidFile && invalidFilePath != "" {
				content, err := os.ReadFile(invalidFilePath)
				if err != nil {
					t.Errorf("Failed to read invalid file: %v", err)
				} else {
					lines := strings.Split(strings.TrimSpace(string(content)), "\n")
					// Remove empty lines
					var nonEmptyLines []string
					for _, line := range lines {
						if strings.TrimSpace(line) != "" {
							nonEmptyLines = append(nonEmptyLines, line)
						}
					}
					if len(nonEmptyLines) != len(tt.invalidEntries) {
						t.Errorf("Expected %d invalid entries in file, got %d", len(tt.invalidEntries), len(nonEmptyLines))
					}
				}
			}
		})
	}
}

func TestSaveToFile(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)

	tempDir, err := os.MkdirTemp("", "save_to_file_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tests := []struct {
		name            string
		entries         []string
		filePath        string
		expectError     bool
		expectedContent []string
		setupFunc       func() // Optional setup function
	}{
		{
			name:            "Save valid entries to file",
			entries:         []string{"domain1.com", "domain2.org", "domain3.net"},
			filePath:        filepath.Join(tempDir, "valid_domains.txt"),
			expectError:     false,
			expectedContent: []string{"domain1.com", "domain2.org", "domain3.net"},
		},
		{
			name:            "Save empty entries",
			entries:         []string{},
			filePath:        filepath.Join(tempDir, "empty.txt"),
			expectError:     false,
			expectedContent: []string{},
		},
		{
			name:            "Save single entry",
			entries:         []string{"single.com"},
			filePath:        filepath.Join(tempDir, "single.txt"),
			expectError:     false,
			expectedContent: []string{"single.com"},
		},
		{
			name:            "Save entries with duplicates (should be sorted and deduplicated)",
			entries:         []string{"zebra.com", "alpha.com", "beta.org", "alpha.com"},
			filePath:        filepath.Join(tempDir, "sorted.txt"),
			expectError:     false,
			expectedContent: []string{"alpha.com", "beta.org", "zebra.com"}, // Sorted and deduplicated
		},
		{
			name:     "Save to read-only directory",
			entries:  []string{"test.com"},
			filePath: filepath.Join(tempDir, "readonly", "test.txt"),
			setupFunc: func() {
				roDir := filepath.Join(tempDir, "readonly")
				os.MkdirAll(roDir, 0755)
				os.Chmod(roDir, 0444) // Read-only
			},
			expectError: true,
		},
		{
			name:            "Save to file in nested directory",
			entries:         []string{"nested.com"},
			filePath:        filepath.Join(tempDir, "level1", "level2", "nested.txt"),
			expectError:     false,
			expectedContent: []string{"nested.com"},
			setupFunc: func() {
				nestedDir := filepath.Join(tempDir, "level1", "level2")
				os.MkdirAll(nestedDir, 0755)
			},
		},
		{
			name:            "Overwrite existing file",
			entries:         []string{"new.com"},
			filePath:        filepath.Join(tempDir, "overwrite.txt"),
			expectError:     false,
			expectedContent: []string{"new.com"},
			setupFunc: func() {
				existingPath := filepath.Join(tempDir, "overwrite.txt")
				os.WriteFile(existingPath, []byte("old.com\n"), 0644)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupFunc != nil {
				tt.setupFunc()
			}

			saveToFile(logger, tt.filePath, tt.entries)

			fileExists := false
			if _, err := os.Stat(tt.filePath); err == nil {
				fileExists = true
			}

			if !tt.expectError && len(tt.entries) > 0 && !fileExists {
				t.Errorf("Expected file %s to exist but it doesn't", tt.filePath)
				return
			}

			if tt.expectError {
				return
			}

			if len(tt.entries) > 0 && !tt.expectError {
				content, err := os.ReadFile(tt.filePath)
				if err != nil {
					t.Errorf("Failed to read file: %v", err)
					return
				}

				lines := strings.Split(strings.TrimSpace(string(content)), "\n")

				var nonEmptyLines []string
				for _, line := range lines {
					if strings.TrimSpace(line) != "" {
						nonEmptyLines = append(nonEmptyLines, line)
					}
				}

				expectedSorted := make([]string, len(tt.expectedContent))
				copy(expectedSorted, tt.expectedContent)
				sort.Strings(expectedSorted)

				actualSorted := make([]string, len(nonEmptyLines))
				copy(actualSorted, nonEmptyLines)
				sort.Strings(actualSorted)

				if len(expectedSorted) != len(actualSorted) {
					t.Errorf("Expected %d lines, got %d", len(expectedSorted), len(actualSorted))
					return
				}

				for i, expected := range expectedSorted {
					if i < len(actualSorted) && actualSorted[i] != expected {
						t.Errorf("Line %d: expected %q, got %q", i, expected, actualSorted[i])
					}
				}
			}

			if tt.setupFunc != nil {
				roDir := filepath.Join(tempDir, "readonly")
				if _, err := os.Stat(roDir); err == nil {
					os.Chmod(roDir, 0755)
				}
			}
		})
	}
}
