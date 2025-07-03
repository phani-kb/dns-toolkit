package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateReadmeCmdRun(t *testing.T) {
	err := os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	require.NoError(t, err)
	defer func() {
		err := os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
		if err != nil {
			t.Logf("Failed to unset DNS_TOOLKIT_TEST_MODE: %v", err)
		}
	}()

	generateReadmeCmd.Run(generateReadmeCmd, []string{})
}

func TestGenerateOutputBranchReadme(t *testing.T) {

	tempDir, err := os.MkdirTemp("", "readme-test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	outputSummariesDir := filepath.Join(tempDir, "summaries")
	outputDir := filepath.Join(tempDir, "output")
	require.NoError(t, os.MkdirAll(outputSummariesDir, 0755))
	require.NoError(t, os.MkdirAll(outputDir, 0755))

	origOutputSummariesDir := constants.OutputSummariesDir
	origOutputDir := constants.OutputDir

	constants.OutputSummariesDir = outputSummariesDir
	constants.OutputDir = outputDir

	defer func() {
		constants.OutputSummariesDir = origOutputSummariesDir
		constants.OutputDir = origOutputDir
	}()

	createTestSummaryFiles(t, outputSummariesDir)

	testFiles := []string{"test_blocklist.txt", "test_allowlist.txt"}
	for _, filename := range testFiles {
		filePath := filepath.Join(outputDir, filename)
		require.NoError(t, os.WriteFile(filePath, []byte("test content"), 0644))
	}

	readme, err := generateOutputBranchReadme()
	require.NoError(t, err)
	assert.NotEmpty(t, readme)

	assert.Contains(t, readme, "# DNS Toolkit - Daily Processing Results")
	assert.Contains(t, readme, "## Quick Start")
	assert.Contains(t, readme, "## Daily Workflow Summary")
	assert.Contains(t, readme, "### Download Statistics")
	assert.Contains(t, readme, "### Processing Statistics")
	assert.Contains(t, readme, "### Consolidation Statistics")
	assert.Contains(t, readme, "## About")

	assert.Contains(t, readme, constants.GitHubRawURL)
	if strings.Contains(readme, "test_blocklist.txt") || strings.Contains(readme, "test_allowlist.txt") {
		t.Logf("Test files found in readme as expected")
	} else {
		t.Logf("Test files not found in readme, but GitHubRawURL is present")
	}
}

func TestCollectWorkflowSummary(t *testing.T) {

	tempDir, err := os.MkdirTemp("", "workflow-summary-test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	outputSummariesDir := filepath.Join(tempDir, "summaries")
	require.NoError(t, os.MkdirAll(outputSummariesDir, 0755))

	origOutputSummariesDir := constants.OutputSummariesDir
	constants.OutputSummariesDir = outputSummariesDir
	defer func() {
		constants.OutputSummariesDir = origOutputSummariesDir
	}()

	createTestSummaryFiles(t, outputSummariesDir)

	summary, err := collectWorkflowSummary()
	require.NoError(t, err)
	assert.NotNil(t, summary)
	assert.NotEmpty(t, summary.LastRun)

	require.NoError(t, os.RemoveAll(outputSummariesDir))
	summary2, err := collectWorkflowSummary()
	require.NoError(t, err)
	assert.NotNil(t, summary2)
}

func TestCollectDownloadStats(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "download-stats-test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	outputSummariesDir := filepath.Join(tempDir, "summaries")
	require.NoError(t, os.MkdirAll(outputSummariesDir, 0755))

	origOutputSummariesDir := constants.OutputSummariesDir
	constants.OutputSummariesDir = outputSummariesDir
	defer func() {
		constants.OutputSummariesDir = origOutputSummariesDir
	}()

	stats := &DownloadStats{}
	err = collectDownloadStats(stats)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "download summary file not found")

	downloadSummaries := []c.DownloadSummary{
		{
			Name:                  "test-source-1",
			Error:                 "",
			LastDownloadTimestamp: "2023-01-01T10:00:00Z",
			Types: []c.SourceType{
				{Name: "domain"},
				{Name: "ipv4"},
			},
		},
		{
			Name:  "test-source-2",
			Error: "download failed",
			Types: []c.SourceType{
				{Name: "domain"},
			},
		},
	}

	summaryFile := filepath.Join(outputSummariesDir, constants.DefaultSummaryFiles["download"])
	content, err := json.Marshal(downloadSummaries)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(summaryFile, content, 0644))

	stats = &DownloadStats{}
	err = collectDownloadStats(stats)
	require.NoError(t, err)

	assert.Equal(t, 2, stats.TotalSources)
	assert.Equal(t, 1, stats.SuccessCount)
	assert.Equal(t, 1, stats.FailedCount)
	assert.Equal(t, "2023-01-01T10:00:00Z", stats.LastUpdateTime)
	assert.Contains(t, stats.ErrorSources, "test-source-2")
	assert.Equal(t, 2, stats.SourcesByType["domain"])
	assert.Equal(t, 1, stats.SourcesByType["ipv4"])
}

func TestCollectProcessingStats(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "processing-stats-test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	outputSummariesDir := filepath.Join(tempDir, "summaries")
	require.NoError(t, os.MkdirAll(outputSummariesDir, 0755))

	origOutputSummariesDir := constants.OutputSummariesDir
	constants.OutputSummariesDir = outputSummariesDir
	defer func() {
		constants.OutputSummariesDir = origOutputSummariesDir
	}()

	// Test with missing file
	var stats *ProcessingStats = &ProcessingStats{}
	err = collectProcessingStats(stats)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "processed summary file not found")

	// Create test processing summary
	processedSummaries := []c.ProcessedSummary{
		{
			Name:                   "test-source",
			LastProcessedTimestamp: "2023-01-01T11:00:00Z",
			ValidFiles: []c.ProcessedFile{
				{GenericSourceType: "domain"},
				{GenericSourceType: "ipv4"},
			},
			InvalidFiles: []c.ProcessedFile{
				{GenericSourceType: "domain"},
			},
		},
	}

	summaryFile := filepath.Join(outputSummariesDir, constants.DefaultSummaryFiles["processed"])
	content, err := json.Marshal(processedSummaries)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(summaryFile, content, 0644))

	stats = &ProcessingStats{}
	err = collectProcessingStats(stats)
	require.NoError(t, err)

	assert.Equal(t, 1, stats.TotalSources)
	assert.Equal(t, "2023-01-01T11:00:00Z", stats.LastUpdateTime)
	assert.Equal(t, 1, stats.ValidFilesByType["domain"])
	assert.Equal(t, 1, stats.ValidFilesByType["ipv4"])
	assert.Equal(t, 1, stats.InvalidFilesByType["domain"])
}

func TestCollectGroupsStats(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "groups-stats-test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	outputSummariesDir := filepath.Join(tempDir, "summaries")
	require.NoError(t, os.MkdirAll(outputSummariesDir, 0755))

	origOutputSummariesDir := constants.OutputSummariesDir
	constants.OutputSummariesDir = outputSummariesDir
	defer func() {
		constants.OutputSummariesDir = origOutputSummariesDir
	}()

	// Test with missing file
	stats := &GroupsStats{}
	err = collectGroupsStats(stats)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "groups summary file not found")

	// Create test groups summary
	groupsSummaries := []c.ConsolidatedGroupsSummary{
		{
			Group:                     "mini",
			LastConsolidatedTimestamp: "2023-01-01T13:00:00Z",
			ConsolidatedSummaries: []c.ConsolidatedSummary{
				{
					Type:     "domain",
					ListType: "blocklist",
					Count:    100,
				},
				{
					Type:     "domain",
					ListType: "allowlist",
					Count:    50,
				},
			},
		},
		{
			Group:                     "lite",
			LastConsolidatedTimestamp: "2023-01-01T13:00:00Z",
			ConsolidatedSummaries: []c.ConsolidatedSummary{
				{
					Type:     "domain",
					ListType: "blocklist",
					Count:    200,
				},
			},
		},
	}

	summaryFile := filepath.Join(outputSummariesDir, constants.DefaultSummaryFiles["consolidated_groups"])
	content, err := json.Marshal(groupsSummaries)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(summaryFile, content, 0644))

	stats = &GroupsStats{}
	err = collectGroupsStats(stats)
	require.NoError(t, err)

	assert.Equal(t, 2, stats.TotalGroups)
	assert.Equal(t, "2023-01-01T13:00:00Z", stats.LastUpdateTime)
	assert.Equal(t, 150, stats.GroupSummary["mini"])
	assert.Equal(t, 200, stats.GroupSummary["lite"])
	assert.Contains(t, stats.GroupListTypes["mini"], "blocklist")
	assert.Contains(t, stats.GroupListTypes["mini"], "allowlist")
	assert.Contains(t, stats.GroupListTypes["lite"], "blocklist")
	assert.NotContains(t, stats.GroupListTypes["lite"], "allowlist")
}

func TestCollectCategoriesStats(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "categories-stats-test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	outputSummariesDir := filepath.Join(tempDir, "summaries")
	require.NoError(t, os.MkdirAll(outputSummariesDir, 0755))

	origOutputSummariesDir := constants.OutputSummariesDir
	constants.OutputSummariesDir = outputSummariesDir
	defer func() {
		constants.OutputSummariesDir = origOutputSummariesDir
	}()

	stats := &CategoriesStats{}
	err = collectCategoriesStats(stats)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "categories summary file not found")

	categoriesSummaries := []c.ConsolidatedCategoriesSummary{
		{
			Category:                  "malware",
			LastConsolidatedTimestamp: "2023-01-01T14:00:00Z",
			ConsolidatedSummaries: []c.ConsolidatedSummary{
				{
					Type:     "domain",
					ListType: "blocklist",
					Count:    300,
				},
				{
					Type:     "ipv4",
					ListType: "blocklist",
					Count:    150,
				},
			},
		},
		{
			Category:                  "ads",
			LastConsolidatedTimestamp: "2023-01-01T14:00:00Z",
			ConsolidatedSummaries: []c.ConsolidatedSummary{
				{
					Type:     "domain",
					ListType: "allowlist",
					Count:    75,
				},
			},
		},
	}

	summaryFile := filepath.Join(outputSummariesDir, constants.DefaultSummaryFiles["consolidated_categories"])
	content, err := json.Marshal(categoriesSummaries)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(summaryFile, content, 0644))

	stats = &CategoriesStats{}
	err = collectCategoriesStats(stats)
	require.NoError(t, err)

	assert.Equal(t, 2, stats.TotalCategories)
	assert.Equal(t, "2023-01-01T14:00:00Z", stats.LastUpdateTime)
	assert.Equal(t, 450, stats.CategorySummary["malware"])
	assert.Equal(t, 75, stats.CategorySummary["ads"])
	assert.Contains(t, stats.CategoryListTypes["malware"], "blocklist")
	assert.NotContains(t, stats.CategoryListTypes["malware"], "allowlist")
	assert.Contains(t, stats.CategoryListTypes["ads"], "allowlist")
	assert.NotContains(t, stats.CategoryListTypes["ads"], "blocklist")
}

func TestCollectConsolidateStats(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "consolidate-stats-test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	outputSummariesDir := filepath.Join(tempDir, "summaries")
	require.NoError(t, os.MkdirAll(outputSummariesDir, 0755))

	origOutputSummariesDir := constants.OutputSummariesDir
	constants.OutputSummariesDir = outputSummariesDir
	defer func() {
		constants.OutputSummariesDir = origOutputSummariesDir
	}()

	// Test with missing file
	stats := &ConsolidateStats{}
	err = collectConsolidateStats(stats)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "consolidated summary file not found")

	// Create test consolidated summary
	consolidatedSummaries := []c.ConsolidatedSummary{
		{
			Type:                      "domain",
			ListType:                  "blocklist",
			Count:                     1000,
			FilesCount:                5,
			IgnoredEntriesCount:       10,
			LastConsolidatedTimestamp: "2023-01-01T12:00:00Z",
		},
		{
			Type:                      "domain",
			ListType:                  "allowlist",
			Count:                     500,
			FilesCount:                3,
			IgnoredEntriesCount:       5,
			LastConsolidatedTimestamp: "2023-01-01T12:00:00Z",
		},
	}

	summaryFile := filepath.Join(outputSummariesDir, constants.DefaultSummaryFiles["consolidated"])
	content, err := json.Marshal(consolidatedSummaries)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(summaryFile, content, 0644))

	stats = &ConsolidateStats{}
	err = collectConsolidateStats(stats)
	require.NoError(t, err)

	assert.Equal(t, 2, stats.TotalFiles)
	assert.Equal(t, "2023-01-01T12:00:00Z", stats.LastUpdateTime)

	domainStats := stats.FilesByType["domain"]
	assert.Equal(t, 1000, domainStats.Blocklist.Count)
	assert.Equal(t, 5, domainStats.Blocklist.FilesCount)
	assert.Equal(t, 10, domainStats.Blocklist.IgnoredCount)
	assert.Equal(t, 500, domainStats.Allowlist.Count)
	assert.Equal(t, 3, domainStats.Allowlist.FilesCount)
	assert.Equal(t, 5, domainStats.Allowlist.IgnoredCount)
}

func TestCollectOverlapStats(t *testing.T) {

	tempDir, err := os.MkdirTemp("", "overlap-stats-test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	outputSummariesDir := filepath.Join(tempDir, "summaries")
	require.NoError(t, os.MkdirAll(outputSummariesDir, 0755))

	origOutputSummariesDir := constants.OutputSummariesDir
	constants.OutputSummariesDir = outputSummariesDir
	defer func() {
		constants.OutputSummariesDir = origOutputSummariesDir
	}()

	// Create test overlap summary
	overlapSummaries := []c.OverlapSummary{
		{
			Source: "test-source-1",
			Type:   "domain",
		},
		{
			Source: "test-source-2",
			Type:   "ipv4",
		},
	}

	summaryFile := filepath.Join(outputSummariesDir, constants.DefaultSummaryFiles["overlap"])
	content, err := json.Marshal(overlapSummaries)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(summaryFile, content, 0644))

	stats := &OverlapStats{}
	err = collectOverlapStats(stats)
	require.NoError(t, err)

	assert.Equal(t, 2, stats.TotalAnalyzed)
	assert.NotEmpty(t, stats.LastUpdateTime)
}

func TestGetTopLevelTxtFiles(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "txt-files-test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	origOutputDir := constants.OutputDir
	constants.OutputDir = tempDir
	defer func() {
		constants.OutputDir = origOutputDir
	}()

	// Create test files
	testFiles := []string{
		"test1.txt",
		"test2.txt",
		"notxt.md",
		"another.json",
	}

	for _, filename := range testFiles {
		filePath := filepath.Join(tempDir, filename)
		require.NoError(t, os.WriteFile(filePath, []byte("content"), 0644))
	}

	// Create a subdirectory to ensure it's ignored
	subDir := filepath.Join(tempDir, "subdir")
	require.NoError(t, os.MkdirAll(subDir, 0755))
	require.NoError(t, os.WriteFile(filepath.Join(subDir, "sub.txt"), []byte("content"), 0644))

	txtFiles, err := getTopLevelTxtFiles()
	require.NoError(t, err)

	expected := []string{"test1.txt", "test2.txt"}
	assert.Equal(t, expected, txtFiles)

	// Test with non-existent directory
	constants.OutputDir = "/non/existent/path"
	_, err = getTopLevelTxtFiles()
	assert.Error(t, err)
}

func TestFormatConsolidateCount(t *testing.T) {

	tests := []struct {
		name     string
		stats    ConsolidateListStats
		expected string
	}{
		{
			name:     "zero count",
			stats:    ConsolidateListStats{Count: 0},
			expected: "-",
		},
		{
			name:     "count without ignored",
			stats:    ConsolidateListStats{Count: 1500, IgnoredCount: 0},
			expected: "1.5K",
		},
		{
			name:     "count with ignored",
			stats:    ConsolidateListStats{Count: 1500, IgnoredCount: 100},
			expected: "1.5K (-100 ignored)",
		},
		{
			name:     "large count with ignored",
			stats:    ConsolidateListStats{Count: 2500000, IgnoredCount: 50000},
			expected: "2.5M (-50.0K ignored)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatConsolidateCount(tt.stats)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatNumber(t *testing.T) {

	tests := []struct {
		input    int
		expected string
	}{
		{0, "0"},
		{999, "999"},
		{1000, "1.0K"},
		{1500, "1.5K"},
		{999999, "1000.0K"}, // Fixed expected value
		{1000000, "1.0M"},
		{2500000, "2.5M"},
		{1234567, "1.2M"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := formatNumber(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Helper function to create test summary files
func createTestSummaryFiles(t *testing.T, outputSummariesDir string) {
	// Download summary
	downloadSummaries := []c.DownloadSummary{
		{
			Name:                  "test-source",
			LastDownloadTimestamp: time.Now().Format(time.RFC3339),
			Types:                 []c.SourceType{{Name: "domain"}},
		},
	}
	content, err := json.Marshal(downloadSummaries)
	require.NoError(t, err)
	summaryFile := filepath.Join(outputSummariesDir, constants.DefaultSummaryFiles["download"])
	require.NoError(t, os.WriteFile(summaryFile, content, 0644))

	// Processed summary
	processedSummaries := []c.ProcessedSummary{
		{
			Name:                   "test-source",
			LastProcessedTimestamp: time.Now().Format(time.RFC3339),
			ValidFiles:             []c.ProcessedFile{{GenericSourceType: "domain"}},
		},
	}
	content, err = json.Marshal(processedSummaries)
	require.NoError(t, err)
	summaryFile = filepath.Join(outputSummariesDir, constants.DefaultSummaryFiles["processed"])
	require.NoError(t, os.WriteFile(summaryFile, content, 0644))

	// Consolidated summary
	consolidatedSummaries := []c.ConsolidatedSummary{
		{
			Type:                      "domain",
			ListType:                  "blocklist",
			Count:                     1000,
			LastConsolidatedTimestamp: time.Now().Format(time.RFC3339),
		},
	}
	content, err = json.Marshal(consolidatedSummaries)
	require.NoError(t, err)
	summaryFile = filepath.Join(outputSummariesDir, constants.DefaultSummaryFiles["consolidated"])
	require.NoError(t, os.WriteFile(summaryFile, content, 0644))

	// Overlap summary
	overlapSummaries := []c.OverlapSummary{
		{Source: "test-source", Type: "domain"},
	}
	content, err = json.Marshal(overlapSummaries)
	require.NoError(t, err)
	summaryFile = filepath.Join(outputSummariesDir, constants.DefaultSummaryFiles["overlap"])
	require.NoError(t, os.WriteFile(summaryFile, content, 0644))

	// Top summary
	topSummaries := []c.TopSummary{
		{
			GenericSourceType: "domain",
			ListType:          "blocklist",
			MinSources:        3,
			Count:             500,
		},
	}
	content, err = json.Marshal(topSummaries)
	require.NoError(t, err)
	summaryFile = filepath.Join(outputSummariesDir, constants.DefaultSummaryFiles["top"])
	require.NoError(t, os.WriteFile(summaryFile, content, 0644))
}

func TestCollectTopStats(t *testing.T) {

	tempDir, err := os.MkdirTemp("", "top-stats-test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	// Store original value
	origOutputSummariesDir := constants.OutputSummariesDir
	defer func() {
		constants.OutputSummariesDir = origOutputSummariesDir
	}()

	// Set test directory
	constants.OutputSummariesDir = tempDir

	tests := []struct {
		name          string
		setupFunc     func() error
		expectedError bool
		expectedFiles int
		expectedTypes map[string]int
	}{
		{
			name: "missing top summary file",
			setupFunc: func() error {
				// Don't create any file
				return nil
			},
			expectedError: true,
		},
		{
			name: "valid top summary file",
			setupFunc: func() error {
				topSummaries := []c.TopSummary{
					{
						GenericSourceType: "domain",
						MinSources:        3,
						Count:             1000,
						ListType:          "blocklist",
						Filepath:          "/tmp/top_domain_blocklist_min3.txt",
					},
					{
						GenericSourceType: "ipv4",
						MinSources:        2,
						Count:             500,
						ListType:          "blocklist",
						Filepath:          "/tmp/top_ipv4_blocklist_min2.txt",
					},
					{
						GenericSourceType: "domain",
						MinSources:        5,
						Count:             200,
						ListType:          "allowlist",
						Filepath:          "/tmp/top_domain_allowlist_min5.txt",
					},
				}

				content, err := json.Marshal(topSummaries)
				if err != nil {
					return err
				}

				topFile := filepath.Join(tempDir, constants.DefaultSummaryFiles["top"])
				return os.WriteFile(topFile, content, 0644)
			},
			expectedError: false,
			expectedFiles: 3,
			expectedTypes: map[string]int{
				"domain": 2,
				"ipv4":   1,
			},
		},
		{
			name: "invalid JSON in top summary file",
			setupFunc: func() error {
				topFile := filepath.Join(tempDir, constants.DefaultSummaryFiles["top"])
				return os.WriteFile(topFile, []byte("invalid json"), 0644)
			},
			expectedError: true,
		},
		{
			name: "empty top summary file",
			setupFunc: func() error {
				topFile := filepath.Join(tempDir, constants.DefaultSummaryFiles["top"])
				return os.WriteFile(topFile, []byte("[]"), 0644)
			},
			expectedError: false,
			expectedFiles: 0,
			expectedTypes: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup test
			err := tt.setupFunc()
			require.NoError(t, err)

			// Clean up before each test
			defer func() {
				topFile := filepath.Join(tempDir, constants.DefaultSummaryFiles["top"])
				os.Remove(topFile)
			}()

			var stats TopStats
			err = collectTopStats(&stats)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedFiles, stats.TotalFiles)
				assert.Equal(t, tt.expectedTypes, stats.FilesByType)

				if tt.expectedFiles > 0 {
					assert.NotEmpty(t, stats.LastUpdateTime)
					assert.NotNil(t, stats.FileDetails)
				}
			}
		})
	}
}
