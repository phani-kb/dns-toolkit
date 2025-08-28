package cmd

import (
	"strings"
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	con "github.com/phani-kb/dns-toolkit/internal/consolidators"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
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
					Name:       "test.txt",
					SourceType: "domain",
					Filepath:   "/path/to/test.txt",
					Count:      100,
				},
			},
			expected: []string{"test.txt [domain] [/path/to/test.txt] [100]"},
		},
		{
			name: "multiple file infos",
			fileInfos: []c.FileInfo{
				{
					Name:       "file1.txt",
					SourceType: "domain",
					Filepath:   "/path/to/file1.txt",
					Count:      50,
				},
				{
					Name:         "file2.txt",
					SourceType:   "ipv4",
					Filepath:     "/path/to/file2.txt",
					Count:        250,
					MustConsider: true,
				},
				{
					Name:       "file3.txt",
					SourceType: "domain",
					Filepath:   "/path/to/file3.txt",
					Count:      75,
				},
			},
			expected: []string{
				"file1.txt [domain] [/path/to/file1.txt] [50]",
				"file2.txt [ipv4] [/path/to/file2.txt] [250] [must consider]",
				"file3.txt [domain] [/path/to/file3.txt] [75]",
			},
		},
		{
			name: "file info with must consider flag",
			fileInfos: []c.FileInfo{
				{
					Name:         "important.txt",
					SourceType:   "ipv4",
					Filepath:     "/path/to/important.txt",
					Count:        500,
					MustConsider: true,
				},
			},
			expected: []string{"important.txt [ipv4] [/path/to/important.txt] [500] [must consider]"},
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
	mockEntries u.StringSet
	mockFiles   []c.FileInfo
}

func (m *mockConsolidator) Consolidate(_ *multilog.Logger, _ []c.ProcessedFile) (u.StringSet, []c.FileInfo) {
	return m.mockEntries, m.mockFiles
}

func (m *mockConsolidator) FilterEntries(
	_ *multilog.Logger,
	entries, filterEntries u.StringSet,
) (u.StringSet, u.StringSet) {
	if len(filterEntries) == 0 {
		// No filtering
		return entries, u.NewStringSet([]string{})
	}

	filteredSet := u.NewStringSet([]string{})
	ignoredSet := u.NewStringSet([]string{})

	for entry := range entries {
		if filterEntries.Contains(entry) {
			ignoredSet.Add(entry)
		} else {
			filteredSet.Add(entry)
		}
	}

	return filteredSet, ignoredSet
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
		mockEntries = u.NewStringSet([]string{"a.com", "b.com"})
		mockFiles   = []c.FileInfo{{Name: "file1.txt", SourceType: "domain", Filepath: "/tmp/file1.txt", Count: 10}}
	)

	mock := &mockConsolidator{
		mockEntries: mockEntries,
		mockFiles:   mockFiles,
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

	// When no filtering is applied, all entries should be returned
	assert.ElementsMatch(t, mockEntries.ToSlice(), gotEntries.ToSlice())
	assert.Equal(t, "domain", gotSummary.Type)
	assert.Equal(t, "blocklist", gotSummary.ListType)
	assert.Equal(t, len(mockFiles), gotSummary.FilesCount)
	assert.Equal(t, len(mockEntries), gotSummary.Count)
}

func TestProcessAllowlists(t *testing.T) {
	// Mock consolidators for allowlist and blocklist
	allowlistMock := &mockConsolidator{
		mockEntries: u.NewStringSet([]string{"allow1.com", "allow2.com", "blocked.com"}),
		mockFiles:   []c.FileInfo{{Name: "allowlist.txt", SourceType: "domain", Count: 3}},
	}

	blocklistMock := &mockConsolidator{
		mockEntries: u.NewStringSet([]string{"blocked.com"}),
		mockFiles:   []c.FileInfo{{Name: "local_blocklist.txt", SourceType: "domain", Count: 1}},
	}

	// Set up test registry
	testRegistry := con.NewConsolidatorRegistry()
	testRegistry.RegisterConsolidator("domain", constants.ListTypeAllowlist, allowlistMock)
	testRegistry.RegisterConsolidator("domain", constants.ListTypeBlocklist, blocklistMock)

	origRegistry := con.Consolidators
	con.Consolidators = testRegistry
	defer func() { con.Consolidators = origRegistry }()

	// Test data
	processedFiles := []c.ProcessedFile{
		{
			Name:              "Regular Allowlist",
			GenericSourceType: "domain",
			ListType:          constants.ListTypeAllowlist,
			Valid:             true,
		},
		{
			Name:              "Local Blocklist", // Name starts with "Local"
			GenericSourceType: "domain",
			ListType:          constants.ListTypeBlocklist,
			Valid:             true,
		},
	}

	genericSourceTypes := []string{"domain"}
	allowlistEntriesByType := make(map[string]u.StringSet)
	var allConsolidatedSummaries []c.ConsolidatedSummary

	// Execute the function
	processAllowlists(
		genericSourceTypes,
		processedFiles,
		allowlistEntriesByType,
		&allConsolidatedSummaries,
	)

	// Verify results
	assert.Len(t, allowlistEntriesByType, 1)
	assert.Contains(t, allowlistEntriesByType, "domain")

	domainEntries := allowlistEntriesByType["domain"]
	assert.Equal(t, 2, len(domainEntries)) // Should have 2 entries after filtering
	assert.True(t, domainEntries.Contains("allow1.com"))
	assert.True(t, domainEntries.Contains("allow2.com"))
	assert.False(t, domainEntries.Contains("blocked.com")) // Should be filtered out by local blocklist

	// Check that we have at least one consolidated summary (may be more if includeInvalid is processed)
	assert.GreaterOrEqual(t, len(allConsolidatedSummaries), 1)

	// Find the summary for allowlist (there might be multiple summaries)
	var allowlistSummary *c.ConsolidatedSummary
	for i := range allConsolidatedSummaries {
		if allConsolidatedSummaries[i].ListType == constants.ListTypeAllowlist {
			allowlistSummary = &allConsolidatedSummaries[i]
			break
		}
	}

	assert.NotNil(t, allowlistSummary, "Should have an allowlist summary")
	if allowlistSummary != nil {
		assert.Equal(t, "domain", allowlistSummary.Type)
		assert.Equal(t, constants.ListTypeAllowlist, allowlistSummary.ListType)
	}
}

func TestGetLocalBlocklistEntries(t *testing.T) {
	t.Parallel()

	// Mock consolidator for local blocklist
	blocklistMock := &mockConsolidator{
		mockEntries: u.NewStringSet([]string{"local-blocked1.com", "local-blocked2.com"}),
		mockFiles:   []c.FileInfo{{Name: "local_blocklist.txt", SourceType: "domain", Count: 2}},
	}

	// Set up test registry
	testRegistry := con.NewConsolidatorRegistry()
	testRegistry.RegisterConsolidator("domain", constants.ListTypeBlocklist, blocklistMock)

	origRegistry := con.Consolidators
	con.Consolidators = testRegistry
	defer func() { con.Consolidators = origRegistry }()

	tests := []struct {
		name              string
		genericSourceType string
		processedFiles    []c.ProcessedFile
		expectedEntries   []string
		expectedCount     int
	}{
		{
			name:              "local blocklist found",
			genericSourceType: "domain",
			processedFiles: []c.ProcessedFile{
				{
					Name:              "Local Blocklist (Domain)",
					GenericSourceType: "domain",
					ListType:          constants.ListTypeBlocklist,
					Valid:             true,
				},
				{
					Name:              "Regular Allowlist",
					GenericSourceType: "domain",
					ListType:          constants.ListTypeAllowlist,
					Valid:             true,
				},
			},
			expectedCount:   2,
			expectedEntries: []string{"local-blocked1.com", "local-blocked2.com"},
		},
		{
			name:              "no local blocklist found",
			genericSourceType: "domain",
			processedFiles: []c.ProcessedFile{
				{
					Name:              "Regular Blocklist",
					GenericSourceType: "domain",
					ListType:          constants.ListTypeBlocklist,
					Valid:             true,
				},
			},
			expectedCount:   0,
			expectedEntries: []string{},
		},
		{
			name:              "invalid local blocklist ignored",
			genericSourceType: "domain",
			processedFiles: []c.ProcessedFile{
				{
					Name:              "Local Blocklist (Domain)",
					GenericSourceType: "domain",
					ListType:          constants.ListTypeBlocklist,
					Valid:             false, // Invalid file should be ignored
				},
			},
			expectedCount:   0,
			expectedEntries: []string{},
		},
		{
			name:              "wrong source type ignored",
			genericSourceType: "ipv4",
			processedFiles: []c.ProcessedFile{
				{
					Name:              "Local Blocklist (Domain)",
					GenericSourceType: "domain", // Different source type
					ListType:          constants.ListTypeBlocklist,
					Valid:             true,
				},
			},
			expectedCount:   0,
			expectedEntries: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getLocalBlocklistEntries(tt.genericSourceType, tt.processedFiles)

			assert.Equal(t, tt.expectedCount, len(result))
			for _, expectedEntry := range tt.expectedEntries {
				assert.True(t, result.Contains(expectedEntry), "Should contain %s", expectedEntry)
			}
		})
	}
}
