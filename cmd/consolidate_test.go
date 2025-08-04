package cmd

import (
	"strings"
	"testing"

	con "github.com/phani-kb/dns-toolkit/internal/consolidators"
	u "github.com/phani-kb/dns-toolkit/internal/utils"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
)

func TestIsConsolidatedSummaryValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		summary  c.ConsolidatedSummary
		expected bool
	}{
		{
			name: "valid summary with count > 0",
			summary: c.ConsolidatedSummary{
				Count: 10,
			},
			expected: true,
		},
		{
			name: "invalid summary with count = 0",
			summary: c.ConsolidatedSummary{
				Count: 0,
			},
			expected: false,
		},
		{
			name: "valid summary with large count",
			summary: c.ConsolidatedSummary{
				Count: 999999,
			},
			expected: true,
		},
		{
			name: "invalid summary with negative count",
			summary: c.ConsolidatedSummary{
				Count: -1,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsConsolidatedSummaryValid(tt.summary)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetFileStrings(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		fileInfos []c.FileInfo
		expected  []string
	}{
		{
			name:      "empty file infos",
			fileInfos: []c.FileInfo{},
			expected:  []string{},
		},
		{
			name: "single file info",
			fileInfos: []c.FileInfo{
				{
					Name:     "test.txt",
					Filepath: "/path/to/test.txt",
					Count:    100,
				},
			},
			expected: []string{"test.txt [/path/to/test.txt] [100]"},
		},
		{
			name: "multiple file infos",
			fileInfos: []c.FileInfo{
				{
					Name:     "file1.txt",
					Filepath: "/path/to/file1.txt",
					Count:    50,
				},
				{
					Name:         "file2.txt",
					Filepath:     "/path/to/file2.txt",
					Count:        250,
					MustConsider: true,
				},
				{
					Name:     "file3.txt",
					Filepath: "/path/to/file3.txt",
					Count:    75,
				},
			},
			expected: []string{
				"file1.txt [/path/to/file1.txt] [50]",
				"file2.txt [/path/to/file2.txt] [250] [must consider]",
				"file3.txt [/path/to/file3.txt] [75]",
			},
		},
		{
			name: "file info with must consider flag",
			fileInfos: []c.FileInfo{
				{
					Name:         "important.txt",
					Filepath:     "/path/to/important.txt",
					Count:        500,
					MustConsider: true,
				},
			},
			expected: []string{"important.txt [/path/to/important.txt] [500] [must consider]"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getFileStrings(tt.fileInfos)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGenerateFileName(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()

	tests := []struct {
		name       string
		fileName   string
		sourceType string
		listType   string
		entryType  string
		expected   string
	}{
		{
			name:       "standard filename generation",
			fileName:   "example",
			sourceType: "domain",
			listType:   "blocklist",
			entryType:  "valid",
			expected:   "example_domain_BL_valid_",
		},
		{
			name:       "filename with allowlist",
			fileName:   "test",
			sourceType: "ip",
			listType:   "allowlist",
			entryType:  "invalid",
			expected:   "test_ip_AL_invalid_",
		},
		{
			name:       "unknown list type",
			fileName:   "unknown",
			sourceType: "custom",
			listType:   "unknown_type",
			entryType:  "valid",
			expected:   "unknown_custom_unknown_type_valid_",
		},
		{
			name:       "empty inputs",
			fileName:   "",
			sourceType: "",
			listType:   "",
			entryType:  "",
			expected:   "_____",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateFileName(logger, tt.fileName, tt.sourceType, tt.listType, tt.entryType)
			// Check that the result starts with the expected prefix
			assert.True(t, len(result) > len(tt.expected), "Result should include hash suffix")
			assert.True(t, len(result) > 30, "Result should include MD5 hash (32 chars)")
			assert.Contains(t, result, tt.expected[:len(tt.expected)-1], "Result should contain expected prefix")
			assert.True(t, strings.HasSuffix(result, ".txt"), "Result should end with .txt")
		})
	}
}

type mockConsolidator struct {
	mockEntries     u.StringSet
	filteredEntries u.StringSet
	ignoredEntries  u.StringSet
	mockFiles       []c.FileInfo
}

func (m *mockConsolidator) Consolidate(_ *multilog.Logger, _ []c.ProcessedFile) (u.StringSet, []c.FileInfo) {
	return m.mockEntries, m.mockFiles
}

func (m *mockConsolidator) FilterEntries(_ *multilog.Logger, entries, _ u.StringSet) (u.StringSet, u.StringSet) {
	return m.filteredEntries, m.ignoredEntries
}

func (m *mockConsolidator) SaveEntries(_ *multilog.Logger, _ u.StringSet, _ string) error {
	return nil
}

func (m *mockConsolidator) IsValid(_ c.ProcessedFile) bool {
	return true
}

func (m *mockConsolidator) GetSourceType() string {
	return "domain"
}

func (m *mockConsolidator) GetListType() string {
	return "blocklist"
}

func TestConsolidateFilesBasedOnSTLT(t *testing.T) {
	var (
		mockEntries     = u.NewStringSet([]string{"a.com", "b.com"})
		mockFiles       = []c.FileInfo{{Name: "file1.txt", Filepath: "/tmp/file1.txt", Count: 10}}
		filteredEntries = u.NewStringSet([]string{"a.com"})
		ignoredEntries  = u.NewStringSet([]string{"b.com"})
	)

	mock := &mockConsolidator{
		mockEntries:     mockEntries,
		mockFiles:       mockFiles,
		filteredEntries: filteredEntries,
		ignoredEntries:  ignoredEntries,
	}

	var _ con.Consolidator = mock

	testRegistry := con.NewConsolidatorRegistry()
	testRegistry.RegisterConsolidator("domain", "blocklist", mock)

	origRegistry := con.Consolidators
	con.Consolidators = testRegistry
	defer func() { con.Consolidators = origRegistry }()

	logger := multilog.NewLogger()
	processedFiles := []c.ProcessedFile{
		{Valid: true},
		{Valid: false},
	}
	entriesToIgnore := u.NewStringSet([]string{})

	gotEntries, gotSummary := consolidateFilesBasedOnSTLT(
		logger,
		"domain",
		"blocklist",
		true,
		entriesToIgnore,
		processedFiles,
	)

	assert.ElementsMatch(t, filteredEntries.ToSlice(), gotEntries.ToSlice())
	assert.Equal(t, "domain", gotSummary.Type)
	assert.Equal(t, "blocklist", gotSummary.ListType)
	assert.Equal(t, len(mockFiles), gotSummary.FilesCount)
	assert.Equal(t, len(filteredEntries), gotSummary.Count)
}
