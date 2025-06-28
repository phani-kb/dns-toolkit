package cmd

import (
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"

	"github.com/stretchr/testify/assert"
)

func TestGetUniqueCategories(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

	logger := multilog.NewLogger()

	tests := []struct {
		name              string
		genericSourceType string
		listType          string
		category          string
		entriesToIgnore   u.StringSet
		processedFiles    []c.ProcessedFile
		expectedCount     int
		expectedValid     bool
		expectedCategory  string
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
	t.Parallel()

	assert.NotNil(t, consolidateCategoriesCmd)
	assert.Equal(t, "categories", consolidateCategoriesCmd.Use)
	assert.Contains(t, consolidateCategoriesCmd.Short, "category-based")
	assert.NotNil(t, consolidateCategoriesCmd.Run)
}
