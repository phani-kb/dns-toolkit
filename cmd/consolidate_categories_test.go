package cmd

import (
	"os"
	"path/filepath"
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	con "github.com/phani-kb/dns-toolkit/internal/consolidators"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"

	"github.com/stretchr/testify/assert"
)

func createTempFileWithContent(t *testing.T, content string) string {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	return tmpFile
}

func TestGetUniqueCategories(t *testing.T) {
	tests := []struct {
		name           string
		processedFiles []c.ProcessedFile
		expected       []string
	}{
		{
			name:           "empty input",
			processedFiles: []c.ProcessedFile{},
			expected:       []string{},
		},
		{
			name: "single category",
			processedFiles: []c.ProcessedFile{
				{Categories: []string{"ads"}},
			},
			expected: []string{"ads"},
		},
		{
			name: "multiple categories sorted",
			processedFiles: []c.ProcessedFile{
				{Categories: []string{"malware", "ads"}},
				{Categories: []string{"privacy", "ads"}},
			},
			expected: []string{"ads", "malware", "privacy"},
		},
		{
			name: "duplicate categories removed",
			processedFiles: []c.ProcessedFile{
				{Categories: []string{"ads", "ads"}},
				{Categories: []string{"ads", "malware"}},
			},
			expected: []string{"ads", "malware"},
		},
		{
			name: "empty categories filtered out",
			processedFiles: []c.ProcessedFile{
				{Categories: []string{"", "ads", ""}},
				{Categories: []string{"malware", ""}},
			},
			expected: []string{"ads", "malware"},
		},
		{
			name: "files without categories",
			processedFiles: []c.ProcessedFile{
				{Categories: []string{}},
				{Categories: nil},
			},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getUniqueCategories(tt.processedFiles)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestConsolidateByCategory(t *testing.T) {
	logger := multilog.NewLogger()

	tests := []struct {
		entriesToIgnore   u.StringSet
		name              string
		genericSourceType string
		listType          string
		category          string
		expectedCategory  string
		processedFiles    []c.ProcessedFile
		expectedCount     int
		expectedValid     bool
	}{
		{
			name:              "empty processed files",
			genericSourceType: constants.SourceTypeDomain,
			listType:          constants.ListTypeBlocklist,
			category:          "ads",
			entriesToIgnore:   u.NewStringSet([]string{}),
			processedFiles:    []c.ProcessedFile{},
			expectedCount:     0,
			expectedValid:     false,
			expectedCategory:  "",
		},
		{
			name:              "invalid generic source type",
			genericSourceType: "invalid_type",
			listType:          constants.ListTypeBlocklist,
			category:          "ads",
			entriesToIgnore:   u.NewStringSet([]string{}),
			processedFiles:    []c.ProcessedFile{{Valid: true}},
			expectedCount:     0,
			expectedValid:     false,
			expectedCategory:  "",
		},
		{
			name:              "valid consolidation",
			genericSourceType: constants.SourceTypeDomain,
			listType:          constants.ListTypeBlocklist,
			category:          "ads",
			entriesToIgnore:   u.NewStringSet([]string{}),
			processedFiles: []c.ProcessedFile{
				{
					Valid:             true,
					GenericSourceType: constants.SourceTypeDomain,
					ListType:          constants.ListTypeBlocklist,
					Categories:        []string{"ads"},
					NumberOfEntries:   2,
					Filepath:          "/tmp/test1.txt",
				},
			},
			expectedCount:    0, // Will be 0 because consolidator can't actually process without real files
			expectedValid:    false,
			expectedCategory: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test registry to isolate from global state
			testRegistry := con.NewConsolidatorRegistry()
			origRegistry := con.Consolidators
			con.Consolidators = testRegistry
			defer func() { con.Consolidators = origRegistry }()

			entries, summary := consolidateByCategory(
				logger,
				tt.genericSourceType,
				tt.listType,
				tt.category,
				tt.entriesToIgnore,
				tt.processedFiles,
			)

			assert.Equal(t, tt.expectedCount, len(entries.ToSlice()))
			assert.Equal(t, tt.expectedValid, summary.Valid)
			assert.Equal(t, tt.expectedCategory, summary.Category)
		})
	}
}

func TestConsolidateCategoriesCommand(t *testing.T) {
	assert.NotNil(t, consolidateCategoriesCmd)
	assert.Equal(t, "categories", consolidateCategoriesCmd.Use)
	assert.Contains(t, consolidateCategoriesCmd.Short, "category-based")
	assert.NotNil(t, consolidateCategoriesCmd.Run)
}

func TestGetFilesForCategory(t *testing.T) {
	tests := []struct {
		name           string
		category       string
		processedFiles []c.ProcessedFile
		expected       int
	}{
		{
			name:           "empty input",
			processedFiles: []c.ProcessedFile{},
			category:       "ads",
			expected:       0,
		},
		{
			name: "no matching category",
			processedFiles: []c.ProcessedFile{
				{Categories: []string{"malware"}, Valid: true},
				{Categories: []string{"privacy"}, Valid: true},
			},
			category: "ads",
			expected: 0,
		},
		{
			name: "matching category",
			processedFiles: []c.ProcessedFile{
				{Categories: []string{"ads"}, Valid: true},
				{Categories: []string{"malware"}, Valid: true},
				{Categories: []string{"ads", "privacy"}, Valid: true},
			},
			category: "ads",
			expected: 2,
		},
		{
			name: "invalid files filtered out",
			processedFiles: []c.ProcessedFile{
				{Categories: []string{"ads"}, Valid: true},
				{Categories: []string{"ads"}, Valid: false}, // Invalid file
				{Categories: []string{"malware"}, Valid: true},
			},
			category: "ads",
			expected: 1,
		},
		{
			name: "multiple categories per file",
			processedFiles: []c.ProcessedFile{
				{Categories: []string{"ads", "malware", "privacy"}, Valid: true},
				{Categories: []string{"malware", "ads"}, Valid: true},
				{Categories: []string{"privacy"}, Valid: true},
			},
			category: "ads",
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getFilesForCategory(tt.processedFiles, tt.category)
			assert.Equal(t, tt.expected, len(result))

			for _, file := range result {
				assert.True(t, file.Valid)
				assert.Contains(t, file.Categories, tt.category)
			}
		})
	}
}

func TestProcessCategoryConsolidation(t *testing.T) {
	logger := multilog.NewLogger()

	con.InitForTesting()

	tests := []struct {
		name               string
		category           string
		processedFiles     []c.ProcessedFile
		genericSourceTypes []string
		expectedResults    int
	}{
		{
			name:               "empty input",
			category:           "ads",
			processedFiles:     []c.ProcessedFile{},
			genericSourceTypes: []string{constants.SourceTypeDomain},
			expectedResults:    0,
		},
		{
			name:     "no matching category",
			category: "ads",
			processedFiles: []c.ProcessedFile{
				{Categories: []string{"malware"}, Valid: true},
			},
			genericSourceTypes: []string{constants.SourceTypeDomain},
			expectedResults:    0,
		},
		{
			name:     "matching files",
			category: "ads",
			processedFiles: []c.ProcessedFile{
				{
					Categories:        []string{"ads"},
					Valid:             true,
					GenericSourceType: constants.SourceTypeDomain,
					ListType:          constants.ListTypeBlocklist,
					Filepath:          createTempFileWithContent(t, "example.com\ntest.com\n"),
				},
			},
			genericSourceTypes: []string{constants.SourceTypeDomain},
			expectedResults:    0, // TODO
		},
		{
			name:     "multiple source types",
			category: "ads",
			processedFiles: []c.ProcessedFile{
				{
					Categories:        []string{"ads"},
					Valid:             true,
					GenericSourceType: constants.SourceTypeDomain,
					ListType:          constants.ListTypeBlocklist,
					Filepath:          createTempFileWithContent(t, "example.com\ntest.com\n"),
				},
				{
					Categories:        []string{"ads"},
					Valid:             true,
					GenericSourceType: constants.SourceTypeIpv4,
					ListType:          constants.ListTypeBlocklist,
					Filepath:          createTempFileWithContent(t, "192.168.1.1\n10.0.0.1\n"),
				},
			},
			genericSourceTypes: []string{constants.SourceTypeDomain, constants.SourceTypeIpv4},
			expectedResults:    0, // TODO
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := processCategoryConsolidation(
				logger,
				tt.category,
				tt.processedFiles,
				tt.genericSourceTypes,
			)

			assert.Equal(t, tt.expectedResults, len(result))

			for sourceType, summaries := range result {
				assert.Contains(t, tt.genericSourceTypes, sourceType)
				assert.IsType(t, []c.ConsolidatedSummary{}, summaries)
			}
		})
	}
}
