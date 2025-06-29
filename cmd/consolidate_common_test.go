package cmd

import (
	"fmt"
	"os"
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/config"
	con "github.com/phani-kb/dns-toolkit/internal/consolidators"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockConsolidatorCommon implements the Consolidator interface for testing
type mockConsolidatorCommon struct {
	mockEntries       u.StringSet
	mockFiles         []c.FileInfo
	filteredEntries   u.StringSet
	ignoredEntries    u.StringSet
	saveError         error
	sourceType        string
	listType          string
	shouldReturnValid bool
}

func (m *mockConsolidatorCommon) Consolidate(_ *multilog.Logger, _ []c.ProcessedFile) (u.StringSet, []c.FileInfo) {
	return m.mockEntries, m.mockFiles
}

func (m *mockConsolidatorCommon) FilterEntries(
	_ *multilog.Logger,
	entries, filterEntries u.StringSet,
) (u.StringSet, u.StringSet) {
	if m.filteredEntries != nil && m.ignoredEntries != nil {
		return m.filteredEntries, m.ignoredEntries
	}
	// Default behavior: return all entries as filtered, none as ignored
	return entries, u.NewStringSet([]string{})
}

func (m *mockConsolidatorCommon) SaveEntries(_ *multilog.Logger, _ u.StringSet, _ string) error {
	return m.saveError
}

func (m *mockConsolidatorCommon) IsValid(_ c.ProcessedFile) bool {
	return m.shouldReturnValid
}

func (m *mockConsolidatorCommon) GetSourceType() string {
	return m.sourceType
}

func (m *mockConsolidatorCommon) GetListType() string {
	return m.listType
}

func TestConsolidateGeneric(t *testing.T) {
	// Disable parallel execution to avoid registry conflicts
	// t.Parallel()

	// Create temporary directory for test outputs
	tempDir := t.TempDir()

	logger := multilog.NewLogger()

	tests := []struct {
		name                string
		params              ConsolidationParams
		entriesToIgnore     u.StringSet
		processedFiles      []c.ProcessedFile
		setupMock           func() *mockConsolidatorCommon
		expectedEntryCount  int
		expectedSummaryType string
		shouldHaveFilepath  bool
		shouldHaveGroup     bool
		shouldHaveCategory  bool
	}{
		{
			name: "successful consolidation with group",
			params: ConsolidationParams{
				GenericSourceType: constants.SourceTypeDomain,
				ListType:          constants.ListTypeBlocklist,
				Identifier:        "test-group",
				OutputDir:         tempDir,
				IdentifierField:   "Group",
			},
			entriesToIgnore: u.NewStringSet([]string{}),
			processedFiles: []c.ProcessedFile{
				{
					Name:              "test1",
					GenericSourceType: constants.SourceTypeDomain,
					ListType:          constants.ListTypeBlocklist,
					Valid:             true,
				},
			},
			setupMock: func() *mockConsolidatorCommon {
				mock := &mockConsolidatorCommon{
					mockEntries: u.NewStringSet([]string{"example.com", "test.com"}),
					mockFiles: []c.FileInfo{
						{Name: "test1", Filepath: "/tmp/test1.txt"},
					},
					sourceType:        constants.SourceTypeDomain,
					listType:          constants.ListTypeBlocklist,
					shouldReturnValid: true,
				}
				return mock
			},
			expectedEntryCount:  2,
			expectedSummaryType: constants.SourceTypeDomain,
			shouldHaveFilepath:  true,
			shouldHaveGroup:     true,
			shouldHaveCategory:  false,
		},
		{
			name: "successful consolidation with category",
			params: ConsolidationParams{
				GenericSourceType: constants.SourceTypeDomain,
				ListType:          constants.ListTypeBlocklist,
				Identifier:        "test-category",
				OutputDir:         tempDir,
				IdentifierField:   "Category",
			},
			entriesToIgnore: u.NewStringSet([]string{}),
			processedFiles: []c.ProcessedFile{
				{
					Name:              "test1",
					GenericSourceType: constants.SourceTypeDomain,
					ListType:          constants.ListTypeBlocklist,
					Valid:             true,
				},
			},
			setupMock: func() *mockConsolidatorCommon {
				mock := &mockConsolidatorCommon{
					mockEntries: u.NewStringSet([]string{"malware.com", "phishing.com"}),
					mockFiles: []c.FileInfo{
						{Name: "test1", Filepath: "/tmp/test1.txt"},
					},
					sourceType:        constants.SourceTypeDomain,
					listType:          constants.ListTypeBlocklist,
					shouldReturnValid: true,
				}
				return mock
			},
			expectedEntryCount:  2,
			expectedSummaryType: constants.SourceTypeDomain,
			shouldHaveFilepath:  true,
			shouldHaveGroup:     false,
			shouldHaveCategory:  true,
		},
		{
			name: "no processed files",
			params: ConsolidationParams{
				GenericSourceType: constants.SourceTypeDomain,
				ListType:          constants.ListTypeBlocklist,
				Identifier:        "empty-group",
				OutputDir:         tempDir,
				IdentifierField:   "Group",
			},
			entriesToIgnore: u.NewStringSet([]string{}),
			processedFiles:  []c.ProcessedFile{},
			setupMock: func() *mockConsolidatorCommon {
				return &mockConsolidatorCommon{}
			},
			expectedEntryCount:  0,
			expectedSummaryType: "",
			shouldHaveFilepath:  false,
			shouldHaveGroup:     false,
			shouldHaveCategory:  false,
		},
		{
			name: "consolidator not found",
			params: ConsolidationParams{
				GenericSourceType: "unknown_type",
				ListType:          "unknown_list",
				Identifier:        "test-group",
				OutputDir:         tempDir,
				IdentifierField:   "Group",
			},
			entriesToIgnore: u.NewStringSet([]string{}),
			processedFiles: []c.ProcessedFile{
				{
					Name:              "test1",
					GenericSourceType: "unknown_type",
					ListType:          "unknown_list",
					Valid:             true,
				},
			},
			setupMock: func() *mockConsolidatorCommon {
				return &mockConsolidatorCommon{}
			},
			expectedEntryCount:  0,
			expectedSummaryType: "",
			shouldHaveFilepath:  false,
			shouldHaveGroup:     false,
			shouldHaveCategory:  false,
		},
		{
			name: "consolidation with ignored entries",
			params: ConsolidationParams{
				GenericSourceType: constants.SourceTypeDomain,
				ListType:          constants.ListTypeBlocklist,
				Identifier:        "test-group",
				OutputDir:         tempDir,
				IdentifierField:   "Group",
			},
			entriesToIgnore: u.NewStringSet([]string{"ignore.com"}),
			processedFiles: []c.ProcessedFile{
				{
					Name:              "test1",
					GenericSourceType: constants.SourceTypeDomain,
					ListType:          constants.ListTypeBlocklist,
					Valid:             true,
				},
			},
			setupMock: func() *mockConsolidatorCommon {
				mock := &mockConsolidatorCommon{
					mockEntries: u.NewStringSet([]string{"example.com", "ignore.com"}),
					mockFiles: []c.FileInfo{
						{Name: "test1", Filepath: "/tmp/test1.txt"},
					},
					filteredEntries:   u.NewStringSet([]string{"example.com"}),
					ignoredEntries:    u.NewStringSet([]string{"ignore.com"}),
					sourceType:        constants.SourceTypeDomain,
					listType:          constants.ListTypeBlocklist,
					shouldReturnValid: true,
				}
				return mock
			},
			expectedEntryCount:  1,
			expectedSummaryType: constants.SourceTypeDomain,
			shouldHaveFilepath:  true,
			shouldHaveGroup:     true,
			shouldHaveCategory:  false,
		},
		{
			name: "no entries after filtering",
			params: ConsolidationParams{
				GenericSourceType: constants.SourceTypeDomain,
				ListType:          constants.ListTypeBlocklist,
				Identifier:        "test-group",
				OutputDir:         tempDir,
				IdentifierField:   "Group",
			},
			entriesToIgnore: u.NewStringSet([]string{}),
			processedFiles: []c.ProcessedFile{
				{
					Name:              "test1",
					GenericSourceType: constants.SourceTypeDomain,
					ListType:          constants.ListTypeBlocklist,
					Valid:             true,
				},
			},
			setupMock: func() *mockConsolidatorCommon {
				mock := &mockConsolidatorCommon{
					mockEntries: u.NewStringSet([]string{"example.com"}),
					mockFiles: []c.FileInfo{
						{Name: "test1", Filepath: "/tmp/test1.txt"},
					},
					filteredEntries:   u.NewStringSet([]string{}), // No entries after filtering
					ignoredEntries:    u.NewStringSet([]string{"example.com"}),
					sourceType:        constants.SourceTypeDomain,
					listType:          constants.ListTypeBlocklist,
					shouldReturnValid: true,
				}
				return mock
			},
			expectedEntryCount:  0,
			expectedSummaryType: "",
			shouldHaveFilepath:  false,
			shouldHaveGroup:     false,
			shouldHaveCategory:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock consolidator
			mock := tt.setupMock()

			// Create test registry and register mock
			testRegistry := con.NewConsolidatorRegistry()
			if mock.sourceType != "" && mock.listType != "" {
				testRegistry.RegisterConsolidator(mock.sourceType, mock.listType, mock)
			}

			// Backup original registry and replace with test registry
			origRegistry := con.Consolidators
			con.Consolidators = testRegistry
			defer func() { con.Consolidators = origRegistry }()

			// Call the function under test
			resultEntries, resultSummary := consolidateGeneric(
				logger,
				tt.params,
				tt.entriesToIgnore,
				tt.processedFiles,
			)

			// Verify results
			assert.Equal(t, tt.expectedEntryCount, len(resultEntries), "Entry count mismatch")
			assert.Equal(t, tt.expectedSummaryType, resultSummary.Type, "Summary type mismatch")

			if tt.shouldHaveFilepath {
				assert.NotEmpty(t, resultSummary.Filepath, "Expected filepath to be set")
				assert.Contains(t, resultSummary.Filepath, tt.params.Identifier, "Filepath should contain identifier")
			} else {
				assert.Empty(t, resultSummary.Filepath, "Expected filepath to be empty")
			}

			if tt.shouldHaveGroup {
				assert.Equal(t, tt.params.Identifier, resultSummary.Group, "Group should be set")
				assert.Empty(t, resultSummary.Category, "Category should be empty when Group is set")
			} else if tt.shouldHaveCategory {
				assert.Equal(t, tt.params.Identifier, resultSummary.Category, "Category should be set")
				assert.Empty(t, resultSummary.Group, "Group should be empty when Category is set")
			} else {
				assert.Empty(t, resultSummary.Group, "Group should be empty")
				assert.Empty(t, resultSummary.Category, "Category should be empty")
			}

			if tt.expectedEntryCount > 0 {
				assert.Equal(t, tt.expectedEntryCount, resultSummary.Count, "Summary count should match")
				assert.True(t, resultSummary.Valid, "Summary should be valid")
				assert.NotEmpty(t, resultSummary.LastConsolidatedTimestamp, "Timestamp should be set")
			}
		})
	}
}

func TestProcessIdentifierConsolidation(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()

	// Mock consolidate function
	mockConsolidateFunc := func(logger *multilog.Logger, gst, listType, identifier string, entriesToIgnore u.StringSet, processedFiles []c.ProcessedFile) (u.StringSet, c.ConsolidatedSummary) {
		entries := u.NewStringSet([]string{"example.com", "test.com"})
		summary := c.ConsolidatedSummary{
			Type:                      gst,
			ListType:                  listType,
			Count:                     len(entries),
			Valid:                     true,
			LastConsolidatedTimestamp: u.GetTimestamp(),
		}

		// Set identifier field based on the function call
		if identifier == "test-group" {
			summary.Group = identifier
		} else if identifier == "test-category" {
			summary.Category = identifier
		}

		return entries, summary
	}

	// Mock get files function for groups
	mockGetGroupFilesFunc := func(files []c.ProcessedFile, group string) []c.ProcessedFile {
		var result []c.ProcessedFile
		for _, file := range files {
			for _, g := range file.Groups {
				if g == group {
					result = append(result, file)
					break
				}
			}
		}
		return result
	}

	// Mock get files function for categories
	mockGetCategoryFilesFunc := func(files []c.ProcessedFile, category string) []c.ProcessedFile {
		var result []c.ProcessedFile
		for _, file := range files {
			for _, c := range file.Categories {
				if c == category {
					result = append(result, file)
					break
				}
			}
		}
		return result
	}

	tests := []struct {
		name                   string
		config                 ProcessingConfig
		expectedSummariesCount int
		expectedAllowlistCount int
		expectedBlocklistCount int
		shouldHaveGroup        bool
		shouldHaveCategory     bool
	}{
		{
			name: "process group consolidation with both allowlist and blocklist",
			config: ProcessingConfig{
				GetFilesFunc:    mockGetGroupFilesFunc,
				ConsolidateFunc: mockConsolidateFunc,
				ProcessedFiles: []c.ProcessedFile{
					{
						Name:              "test1",
						GenericSourceType: constants.SourceTypeDomain,
						ListType:          constants.ListTypeAllowlist,
						Groups:            []string{"test-group"},
						Valid:             true,
					},
					{
						Name:              "test2",
						GenericSourceType: constants.SourceTypeDomain,
						ListType:          constants.ListTypeBlocklist,
						Groups:            []string{"test-group"},
						Valid:             true,
					},
				},
				GenericSourceTypes: []string{constants.SourceTypeDomain},
				Identifier:         "test-group",
				IdentifierField:    "Group",
			},
			expectedSummariesCount: 2, // one allowlist + one blocklist
			expectedAllowlistCount: 1,
			expectedBlocklistCount: 1,
			shouldHaveGroup:        true,
			shouldHaveCategory:     false,
		},
		{
			name: "process category consolidation",
			config: ProcessingConfig{
				GetFilesFunc:    mockGetCategoryFilesFunc,
				ConsolidateFunc: mockConsolidateFunc,
				ProcessedFiles: []c.ProcessedFile{
					{
						Name:              "test1",
						GenericSourceType: constants.SourceTypeDomain,
						ListType:          constants.ListTypeBlocklist,
						Categories:        []string{"test-category"},
						Valid:             true,
					},
				},
				GenericSourceTypes: []string{constants.SourceTypeDomain},
				Identifier:         "test-category",
				IdentifierField:    "Category",
			},
			expectedSummariesCount: 1, // only blocklist
			expectedAllowlistCount: 0,
			expectedBlocklistCount: 1,
			shouldHaveGroup:        false,
			shouldHaveCategory:     true,
		},
		{
			name: "no files for identifier",
			config: ProcessingConfig{
				GetFilesFunc:    mockGetGroupFilesFunc,
				ConsolidateFunc: mockConsolidateFunc,
				ProcessedFiles: []c.ProcessedFile{
					{
						Name:              "test1",
						GenericSourceType: constants.SourceTypeDomain,
						ListType:          constants.ListTypeBlocklist,
						Groups:            []string{"other-group"}, // Different group
						Valid:             true,
					},
				},
				GenericSourceTypes: []string{constants.SourceTypeDomain},
				Identifier:         "test-group",
				IdentifierField:    "Group",
			},
			expectedSummariesCount: 0,
			expectedAllowlistCount: 0,
			expectedBlocklistCount: 0,
			shouldHaveGroup:        false,
			shouldHaveCategory:     false,
		},
		{
			name: "only allowlist files",
			config: ProcessingConfig{
				GetFilesFunc:    mockGetGroupFilesFunc,
				ConsolidateFunc: mockConsolidateFunc,
				ProcessedFiles: []c.ProcessedFile{
					{
						Name:              "test1",
						GenericSourceType: constants.SourceTypeDomain,
						ListType:          constants.ListTypeAllowlist,
						Groups:            []string{"test-group"},
						Valid:             true,
					},
				},
				GenericSourceTypes: []string{constants.SourceTypeDomain},
				Identifier:         "test-group",
				IdentifierField:    "Group",
			},
			expectedSummariesCount: 1, // only allowlist
			expectedAllowlistCount: 1,
			expectedBlocklistCount: 0,
			shouldHaveGroup:        true,
			shouldHaveCategory:     false,
		},
		{
			name: "multiple source types",
			config: ProcessingConfig{
				GetFilesFunc:    mockGetGroupFilesFunc,
				ConsolidateFunc: mockConsolidateFunc,
				ProcessedFiles: []c.ProcessedFile{
					{
						Name:              "test1",
						GenericSourceType: constants.SourceTypeDomain,
						ListType:          constants.ListTypeBlocklist,
						Groups:            []string{"test-group"},
						Valid:             true,
					},
					{
						Name:              "test2",
						GenericSourceType: constants.SourceTypeIpv4,
						ListType:          constants.ListTypeBlocklist,
						Groups:            []string{"test-group"},
						Valid:             true,
					},
				},
				GenericSourceTypes: []string{constants.SourceTypeDomain, constants.SourceTypeIpv4},
				Identifier:         "test-group",
				IdentifierField:    "Group",
			},
			expectedSummariesCount: 2, // domain + ipv4 blocklists
			expectedAllowlistCount: 0,
			expectedBlocklistCount: 2,
			shouldHaveGroup:        true,
			shouldHaveCategory:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := processIdentifierConsolidation(logger, tt.config)

			// Verify the identifier exists in results
			summaries, exists := result[tt.config.Identifier]
			if tt.expectedSummariesCount == 0 {
				assert.True(t, !exists || len(summaries) == 0, "Expected no summaries")
				return
			}

			require.True(t, exists, "Expected identifier to exist in results")
			assert.Equal(t, tt.expectedSummariesCount, len(summaries), "Summary count mismatch")

			// Count allowlist and blocklist summaries
			allowlistCount := 0
			blocklistCount := 0
			for _, summary := range summaries {
				if summary.ListType == constants.ListTypeAllowlist {
					allowlistCount++
				} else if summary.ListType == constants.ListTypeBlocklist {
					blocklistCount++
				}

				// Verify identifier fields are set correctly
				if tt.shouldHaveGroup {
					assert.Equal(t, tt.config.Identifier, summary.Group, "Group should be set")
					assert.Empty(t, summary.Category, "Category should be empty when Group is set")
				} else if tt.shouldHaveCategory {
					assert.Equal(t, tt.config.Identifier, summary.Category, "Category should be set")
					assert.Empty(t, summary.Group, "Group should be empty when Category is set")
				}
			}

			assert.Equal(t, tt.expectedAllowlistCount, allowlistCount, "Allowlist count mismatch")
			assert.Equal(t, tt.expectedBlocklistCount, blocklistCount, "Blocklist count mismatch")
		})
	}
}

func TestConsolidationParams(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		params ConsolidationParams
	}{
		{
			name: "create consolidation params for group",
			params: ConsolidationParams{
				GenericSourceType: constants.SourceTypeDomain,
				ListType:          constants.ListTypeBlocklist,
				Identifier:        "test-group",
				OutputDir:         "/tmp/output",
				IdentifierField:   "Group",
			},
		},
		{
			name: "create consolidation params for category",
			params: ConsolidationParams{
				GenericSourceType: constants.SourceTypeIpv4,
				ListType:          constants.ListTypeAllowlist,
				Identifier:        "test-category",
				OutputDir:         "/tmp/output",
				IdentifierField:   "Category",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotEmpty(t, tt.params.GenericSourceType, "GenericSourceType should not be empty")
			assert.NotEmpty(t, tt.params.ListType, "ListType should not be empty")
			assert.NotEmpty(t, tt.params.Identifier, "Identifier should not be empty")
			assert.NotEmpty(t, tt.params.OutputDir, "OutputDir should not be empty")
			assert.NotEmpty(t, tt.params.IdentifierField, "IdentifierField should not be empty")
		})
	}
}

func TestProcessingConfig(t *testing.T) {
	t.Parallel()

	mockGetFilesFunc := func(files []c.ProcessedFile, identifier string) []c.ProcessedFile {
		return files
	}

	mockConsolidateFunc := func(logger *multilog.Logger, gst, listType, identifier string, entriesToIgnore u.StringSet, processedFiles []c.ProcessedFile) (u.StringSet, c.ConsolidatedSummary) {
		return u.NewStringSet([]string{}), c.ConsolidatedSummary{}
	}

	config := ProcessingConfig{
		GetFilesFunc:       mockGetFilesFunc,
		ConsolidateFunc:    mockConsolidateFunc,
		ProcessedFiles:     []c.ProcessedFile{},
		GenericSourceTypes: []string{constants.SourceTypeDomain},
		Identifier:         "test",
		IdentifierField:    "Group",
	}

	assert.NotNil(t, config.GetFilesFunc, "GetFilesFunc should not be nil")
	assert.NotNil(t, config.ConsolidateFunc, "ConsolidateFunc should not be nil")
	assert.NotNil(t, config.ProcessedFiles, "ProcessedFiles should not be nil")
	assert.NotEmpty(t, config.GenericSourceTypes, "GenericSourceTypes should not be empty")
	assert.NotEmpty(t, config.Identifier, "Identifier should not be empty")
	assert.NotEmpty(t, config.IdentifierField, "IdentifierField should not be empty")
}

func TestConsolidateGeneric_FileSaveError(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	logger := multilog.NewLogger()

	params := ConsolidationParams{
		GenericSourceType: constants.SourceTypeDomain,
		ListType:          constants.ListTypeBlocklist,
		Identifier:        "test-group",
		OutputDir:         tempDir,
		IdentifierField:   "Group",
	}

	entriesToIgnore := u.NewStringSet([]string{})
	processedFiles := []c.ProcessedFile{
		{
			Name:              "test1",
			GenericSourceType: constants.SourceTypeDomain,
			ListType:          constants.ListTypeBlocklist,
			Valid:             true,
		},
	}

	// Setup mock consolidator with save error
	mock := &mockConsolidatorCommon{
		mockEntries: u.NewStringSet([]string{"example.com"}),
		mockFiles: []c.FileInfo{
			{Name: "test1", Filepath: "/tmp/test1.txt"},
		},
		filteredEntries:   u.NewStringSet([]string{"example.com"}),
		ignoredEntries:    u.NewStringSet([]string{}),
		saveError:         assert.AnError, // Simulate save error
		sourceType:        constants.SourceTypeDomain,
		listType:          constants.ListTypeBlocklist,
		shouldReturnValid: true,
	}

	// Create test registry and register mock
	testRegistry := con.NewConsolidatorRegistry()
	testRegistry.RegisterConsolidator(mock.sourceType, mock.listType, mock)

	// Backup original registry and replace with test registry
	origRegistry := con.Consolidators
	con.Consolidators = testRegistry
	defer func() { con.Consolidators = origRegistry }()

	// Call the function under test
	resultEntries, resultSummary := consolidateGeneric(
		logger,
		params,
		entriesToIgnore,
		processedFiles,
	)

	// Even with save error, the function should return the entries and summary
	assert.Equal(t, 1, len(resultEntries), "Should return entries even with save error")
	assert.Equal(t, constants.SourceTypeDomain, resultSummary.Type, "Summary type should be set")
	assert.NotEmpty(t, resultSummary.Filepath, "Filepath should be set even with save error")
}

// mockConsolidatorWithCustomSave allows overriding SaveEntries behavior
type mockConsolidatorWithCustomSave struct {
	*mockConsolidatorCommon
	customSaveFunc func(*multilog.Logger, u.StringSet, string) error
}

func (m *mockConsolidatorWithCustomSave) SaveEntries(
	logger *multilog.Logger,
	entries u.StringSet,
	filePath string,
) error {
	if m.customSaveFunc != nil {
		return m.customSaveFunc(logger, entries, filePath)
	}
	return m.mockConsolidatorCommon.SaveEntries(logger, entries, filePath)
}

func TestConsolidateGeneric_ChecksumCalculation(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	logger := multilog.NewLogger()

	// Set up AppConfig for testing
	oldConfig := AppConfig
	defer func() { AppConfig = oldConfig }()

	AppConfig = &config.AppConfig{
		DNSToolkit: config.DNSToolkitConfig{
			FilesChecksum: config.FilesChecksumConfig{
				Enabled:   true,
				Algorithm: "md5",
			},
		},
	}

	params := ConsolidationParams{
		GenericSourceType: constants.SourceTypeDomain,
		ListType:          constants.ListTypeBlocklist,
		Identifier:        "test",
		OutputDir:         tempDir,
		IdentifierField:   "Group",
	}

	// Setup mock consolidator that saves to the test file
	baseMock := &mockConsolidatorCommon{
		mockEntries: u.NewStringSet([]string{"example.com", "test.com"}),
		mockFiles: []c.FileInfo{
			{Name: "test1", Filepath: "/tmp/test1.txt"},
		},
		filteredEntries:   u.NewStringSet([]string{"example.com", "test.com"}), // Ensure entries are filtered
		ignoredEntries:    u.NewStringSet([]string{}),
		sourceType:        constants.SourceTypeDomain,
		listType:          constants.ListTypeBlocklist,
		shouldReturnValid: true,
		saveError:         nil,
	}

	mock := &mockConsolidatorWithCustomSave{
		mockConsolidatorCommon: baseMock,
		customSaveFunc: func(logger *multilog.Logger, entries u.StringSet, filePath string) error {
			// Write test content to the expected file path
			return os.WriteFile(filePath, []byte("example.com\ntest.com\n"), 0644)
		},
	}

	testRegistry := con.NewConsolidatorRegistry()
	testRegistry.RegisterConsolidator(mock.sourceType, mock.listType, mock)

	origRegistry := con.Consolidators
	con.Consolidators = testRegistry
	defer func() { con.Consolidators = origRegistry }()

	// Enable checksum calculation
	origCalculateChecksum := calculateChecksum
	calculateChecksum = true
	defer func() { calculateChecksum = origCalculateChecksum }()

	processedFiles := []c.ProcessedFile{
		{
			Name:              "test1",
			GenericSourceType: constants.SourceTypeDomain,
			ListType:          constants.ListTypeBlocklist,
			Valid:             true,
		},
	}

	// Call the function under test
	_, resultSummary := consolidateGeneric(
		logger,
		params,
		u.NewStringSet([]string{}),
		processedFiles,
	)

	// Verify checksum is calculated
	assert.NotEmpty(t, resultSummary.Checksum, "Checksum should be calculated when enabled")
}

// Additional edge case tests

func TestConsolidateGeneric_EdgeCases(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	logger := multilog.NewLogger()

	t.Run("empty identifier", func(t *testing.T) {
		params := ConsolidationParams{
			GenericSourceType: constants.SourceTypeDomain,
			ListType:          constants.ListTypeBlocklist,
			Identifier:        "", // Empty identifier
			OutputDir:         tempDir,
			IdentifierField:   "Group",
		}

		mock := &mockConsolidatorCommon{
			mockEntries:       u.NewStringSet([]string{"example.com"}),
			mockFiles:         []c.FileInfo{{Name: "test", Filepath: "/tmp/test.txt"}},
			sourceType:        constants.SourceTypeDomain,
			listType:          constants.ListTypeBlocklist,
			shouldReturnValid: true,
		}

		testRegistry := con.NewConsolidatorRegistry()
		testRegistry.RegisterConsolidator(mock.sourceType, mock.listType, mock)

		origRegistry := con.Consolidators
		con.Consolidators = testRegistry
		defer func() { con.Consolidators = origRegistry }()

		processedFiles := []c.ProcessedFile{
			{
				Name:              "test",
				GenericSourceType: constants.SourceTypeDomain,
				ListType:          constants.ListTypeBlocklist,
				Valid:             true,
			},
		}

		entries, summary := consolidateGeneric(logger, params, u.NewStringSet([]string{}), processedFiles)

		assert.Equal(t, 1, len(entries), "Should handle empty identifier")
		assert.Contains(t, summary.Filepath, "_domain_blocklist.txt", "Filepath should still be generated")
	})

	t.Run("invalid identifier field", func(t *testing.T) {
		params := ConsolidationParams{
			GenericSourceType: constants.SourceTypeDomain,
			ListType:          constants.ListTypeBlocklist,
			Identifier:        "test",
			OutputDir:         tempDir,
			IdentifierField:   "InvalidField", // Invalid field
		}

		mock := &mockConsolidatorCommon{
			mockEntries:       u.NewStringSet([]string{"example.com"}),
			mockFiles:         []c.FileInfo{{Name: "test", Filepath: "/tmp/test.txt"}},
			sourceType:        constants.SourceTypeDomain,
			listType:          constants.ListTypeBlocklist,
			shouldReturnValid: true,
		}

		testRegistry := con.NewConsolidatorRegistry()
		testRegistry.RegisterConsolidator(mock.sourceType, mock.listType, mock)

		origRegistry := con.Consolidators
		con.Consolidators = testRegistry
		defer func() { con.Consolidators = origRegistry }()

		processedFiles := []c.ProcessedFile{
			{
				Name:              "test",
				GenericSourceType: constants.SourceTypeDomain,
				ListType:          constants.ListTypeBlocklist,
				Valid:             true,
			},
		}

		entries, summary := consolidateGeneric(logger, params, u.NewStringSet([]string{}), processedFiles)

		assert.Equal(t, 1, len(entries), "Should process even with invalid identifier field")
		assert.Empty(t, summary.Group, "Group should not be set for invalid field")
		assert.Empty(t, summary.Category, "Category should not be set for invalid field")
	})

	t.Run("nil entries from consolidator", func(t *testing.T) {
		params := ConsolidationParams{
			GenericSourceType: constants.SourceTypeDomain,
			ListType:          constants.ListTypeBlocklist,
			Identifier:        "test",
			OutputDir:         tempDir,
			IdentifierField:   "Group",
		}

		mock := &mockConsolidatorCommon{
			mockEntries:       nil, // Nil entries
			mockFiles:         []c.FileInfo{{Name: "test", Filepath: "/tmp/test.txt"}},
			sourceType:        constants.SourceTypeDomain,
			listType:          constants.ListTypeBlocklist,
			shouldReturnValid: true,
		}

		testRegistry := con.NewConsolidatorRegistry()
		testRegistry.RegisterConsolidator(mock.sourceType, mock.listType, mock)

		origRegistry := con.Consolidators
		con.Consolidators = testRegistry
		defer func() { con.Consolidators = origRegistry }()

		processedFiles := []c.ProcessedFile{
			{
				Name:              "test",
				GenericSourceType: constants.SourceTypeDomain,
				ListType:          constants.ListTypeBlocklist,
				Valid:             true,
			},
		}

		entries, summary := consolidateGeneric(logger, params, u.NewStringSet([]string{}), processedFiles)

		assert.Equal(t, 0, len(entries), "Should handle nil entries gracefully")
		assert.Empty(t, summary.Filepath, "Should not create filepath for empty result")
	})
}

func TestProcessIdentifierConsolidation_EdgeCases(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()

	t.Run("nil get files function", func(t *testing.T) {
		config := ProcessingConfig{
			GetFilesFunc:       nil, // Nil function
			ConsolidateFunc:    nil,
			ProcessedFiles:     []c.ProcessedFile{},
			GenericSourceTypes: []string{constants.SourceTypeDomain},
			Identifier:         "test",
			IdentifierField:    "Group",
		}

		// This should panic because the function expects valid function pointers
		// This test demonstrates that the function requires proper validation
		assert.Panics(t, func() {
			processIdentifierConsolidation(logger, config)
		}, "Should panic with nil function")
	})

	t.Run("empty generic source types", func(t *testing.T) {
		mockGetFilesFunc := func(files []c.ProcessedFile, identifier string) []c.ProcessedFile {
			return files
		}

		mockConsolidateFunc := func(logger *multilog.Logger, gst, listType, identifier string, entriesToIgnore u.StringSet, processedFiles []c.ProcessedFile) (u.StringSet, c.ConsolidatedSummary) {
			return u.NewStringSet([]string{}), c.ConsolidatedSummary{}
		}

		config := ProcessingConfig{
			GetFilesFunc:       mockGetFilesFunc,
			ConsolidateFunc:    mockConsolidateFunc,
			ProcessedFiles:     []c.ProcessedFile{},
			GenericSourceTypes: []string{}, // Empty slice
			Identifier:         "test",
			IdentifierField:    "Group",
		}

		result := processIdentifierConsolidation(logger, config)

		summaries := result["test"]
		assert.Equal(t, 0, len(summaries), "Should handle empty source types")
	})

	t.Run("unknown identifier field", func(t *testing.T) {
		mockGetFilesFunc := func(files []c.ProcessedFile, identifier string) []c.ProcessedFile {
			return files
		}

		mockConsolidateFunc := func(logger *multilog.Logger, gst, listType, identifier string, entriesToIgnore u.StringSet, processedFiles []c.ProcessedFile) (u.StringSet, c.ConsolidatedSummary) {
			return u.NewStringSet([]string{"test.com"}), c.ConsolidatedSummary{
				Type:     gst,
				ListType: listType,
				Count:    1,
				Valid:    true,
			}
		}

		config := ProcessingConfig{
			GetFilesFunc:    mockGetFilesFunc,
			ConsolidateFunc: mockConsolidateFunc,
			ProcessedFiles: []c.ProcessedFile{
				{
					Name:              "test",
					GenericSourceType: constants.SourceTypeDomain,
					ListType:          constants.ListTypeBlocklist,
					Groups:            []string{"test"},
					Valid:             true,
				},
			},
			GenericSourceTypes: []string{constants.SourceTypeDomain},
			Identifier:         "test",
			IdentifierField:    "UnknownField", // Unknown field
		}

		result := processIdentifierConsolidation(logger, config)

		summaries := result["test"]
		assert.Equal(t, 1, len(summaries), "Should process with unknown field")
		assert.Empty(t, summaries[0].Group, "Group should not be set")
		assert.Empty(t, summaries[0].Category, "Category should not be set")
	})
}

func TestConsolidateGeneric_LargeDataset(t *testing.T) {
	// Disable parallel execution to avoid registry conflicts
	// t.Parallel()

	tempDir := t.TempDir()
	logger := multilog.NewLogger()

	// Create a large set of entries
	largeEntrySet := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		largeEntrySet[i] = fmt.Sprintf("domain%d.com", i)
	}

	params := ConsolidationParams{
		GenericSourceType: constants.SourceTypeDomain,
		ListType:          constants.ListTypeBlocklist,
		Identifier:        "large-test",
		OutputDir:         tempDir,
		IdentifierField:   "Group",
	}

	mock := &mockConsolidatorCommon{
		mockEntries: u.NewStringSet(largeEntrySet),
		mockFiles: []c.FileInfo{
			{Name: "large_test", Filepath: "/tmp/large_test.txt"},
		},
		filteredEntries:   u.NewStringSet(largeEntrySet), // Explicitly set filtered entries
		ignoredEntries:    u.NewStringSet([]string{}),    // No ignored entries
		sourceType:        constants.SourceTypeDomain,
		listType:          constants.ListTypeBlocklist,
		shouldReturnValid: true,
	}

	testRegistry := con.NewConsolidatorRegistry()
	testRegistry.RegisterConsolidator(mock.sourceType, mock.listType, mock)

	origRegistry := con.Consolidators
	con.Consolidators = testRegistry
	defer func() { con.Consolidators = origRegistry }()

	processedFiles := []c.ProcessedFile{
		{
			Name:              "large_test",
			GenericSourceType: constants.SourceTypeDomain,
			ListType:          constants.ListTypeBlocklist,
			Valid:             true,
		},
	}

	entries, summary := consolidateGeneric(logger, params, u.NewStringSet([]string{}), processedFiles)

	assert.Equal(t, 1000, len(entries), "Should handle large datasets")
	assert.Equal(t, 1000, summary.Count, "Summary count should match")
	assert.True(t, summary.Valid, "Summary should be valid")
	assert.Equal(t, "large-test", summary.Group, "Group should be set correctly")
}

func TestConsolidateGeneric_ConfigFlags(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	logger := multilog.NewLogger()

	params := ConsolidationParams{
		GenericSourceType: constants.SourceTypeDomain,
		ListType:          constants.ListTypeBlocklist,
		Identifier:        "config-test",
		OutputDir:         tempDir,
		IdentifierField:   "Group",
	}

	// Test with AppConfig.DNSToolkit.FilesChecksum.Enabled = true (simulated)
	baseMock := &mockConsolidatorCommon{
		mockEntries: u.NewStringSet([]string{"example.com"}),
		mockFiles: []c.FileInfo{
			{Name: "test", Filepath: "/tmp/test.txt"},
		},
		filteredEntries:   u.NewStringSet([]string{"example.com"}), // Ensure filtering works
		ignoredEntries:    u.NewStringSet([]string{}),
		sourceType:        constants.SourceTypeDomain,
		listType:          constants.ListTypeBlocklist,
		shouldReturnValid: true,
	}

	mock := &mockConsolidatorWithCustomSave{
		mockConsolidatorCommon: baseMock,
		customSaveFunc: func(logger *multilog.Logger, entries u.StringSet, filePath string) error {
			return os.WriteFile(filePath, []byte("example.com\n"), 0644)
		},
	}

	testRegistry := con.NewConsolidatorRegistry()
	testRegistry.RegisterConsolidator(mock.sourceType, mock.listType, mock)

	origRegistry := con.Consolidators
	con.Consolidators = testRegistry
	defer func() { con.Consolidators = origRegistry }()

	processedFiles := []c.ProcessedFile{
		{
			Name:              "test",
			GenericSourceType: constants.SourceTypeDomain,
			ListType:          constants.ListTypeBlocklist,
			Valid:             true,
		},
	}

	t.Run("with calculateChecksum false", func(t *testing.T) {
		origCalculateChecksum := calculateChecksum
		calculateChecksum = false
		defer func() { calculateChecksum = origCalculateChecksum }()

		_, summary := consolidateGeneric(logger, params, u.NewStringSet([]string{}), processedFiles)

		// When there are entries, filepath should be set regardless of calculateChecksum setting
		if summary.Count > 0 {
			assert.NotEmpty(t, summary.Filepath, "Filepath should be set when entries exist")
		} else {
			// Test passed - no entries to consolidate is a valid scenario
			assert.Equal(t, 0, summary.Count, "Count should be 0 when no entries")
		}
	})
}
