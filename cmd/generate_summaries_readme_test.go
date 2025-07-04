package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateSummariesReadmeCmdRun(t *testing.T) {
	oldLogger := Logger
	Logger = config.CreateTestLogger()
	defer func() { Logger = oldLogger }()

	err := os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	require.NoError(t, err)
	defer func() {
		err := os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
		if err != nil {
			t.Logf("Failed to unset DNS_TOOLKIT_TEST_MODE: %v", err)
		}
	}()

	generateSummariesReadmeCmd.Run(generateSummariesReadmeCmd, []string{})
}

func TestGenerateSummariesReadme(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "summaries-readme-test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	summariesDir := filepath.Join(tempDir, "summaries")
	require.NoError(t, os.MkdirAll(summariesDir, 0755))

	origOutputDir := constants.OutputDir
	constants.OutputDir = tempDir
	defer func() {
		constants.OutputDir = origOutputDir
	}()

	createTestSummaryFilesForSummariesReadme(t, summariesDir)

	readme := generateSummariesReadme()

	assert.NotEmpty(t, readme)
	assert.Contains(t, readme, "# DNS Toolkit - Summary Files")
	assert.Contains(t, readme, "## Overview")
	assert.Contains(t, readme, "## Overall Statistics")
	assert.Contains(t, readme, "## Summary Files")
	assert.Contains(t, readme, "## About")
	assert.Contains(t, readme, "Last Generated:")
	assert.Contains(t, readme, "github.com/phani-kb/dns-toolkit")
}

func TestCollectSummariesInfo(t *testing.T) {
	oldLogger := Logger
	Logger = config.CreateTestLogger()
	defer func() { Logger = oldLogger }()

	tempDir, err := os.MkdirTemp("", "summaries-info-test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	nonExistentDir := filepath.Join(tempDir, "nonexistent")
	origOutputDir := constants.OutputDir
	constants.OutputDir = nonExistentDir
	defer func() {
		constants.OutputDir = origOutputDir
	}()

	info := collectSummariesInfo()
	assert.NotNil(t, info)
	assert.Equal(t, 0, info.TotalFiles)
	assert.Equal(t, 0, len(info.SummaryTypes))
	assert.NotEmpty(t, info.LastGenerated)

	summariesDir := filepath.Join(tempDir, "summaries")
	require.NoError(t, os.MkdirAll(summariesDir, 0755))
	constants.OutputDir = tempDir

	createTestSummaryFilesForSummariesReadme(t, summariesDir)

	info = collectSummariesInfo()
	assert.NotNil(t, info)
	assert.True(t, info.TotalFiles > 0)
	assert.True(t, len(info.SummaryTypes) > 0)
	assert.NotEmpty(t, info.LastGenerated)

	assert.True(t, info.OverallStats.TotalSources > 0)
	assert.True(t, info.OverallStats.TotalDownloads > 0)
}

func TestGetDetailedStatsForSummaryType(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "detailed-stats-test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	summariesDir := filepath.Join(tempDir, "summaries")
	require.NoError(t, os.MkdirAll(summariesDir, 0755))

	origOutputDir := constants.OutputDir
	constants.OutputDir = tempDir
	defer func() {
		constants.OutputDir = origOutputDir
	}()

	tests := []struct {
		name         string
		filename     string
		createFile   func(string) error
		expectStats  bool
		expectedText []string
	}{
		{
			name:     "non-existent file",
			filename: "non_existent.json",
			createFile: func(path string) error {
				return nil
			},
			expectStats: false,
		},
		{
			name:     "download_summary.json",
			filename: "download_summary.json",
			createFile: func(path string) error {
				downloadSummaries := []c.DownloadSummary{
					{
						Name:  "test-source-1",
						Error: "",
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
				content, err := json.Marshal(downloadSummaries)
				if err != nil {
					return err
				}
				return os.WriteFile(path, content, 0644)
			},
			expectStats:  true,
			expectedText: []string{"Sources:", "2 total", "1 successful", "1 failed", "Types:", "domain (2)", "ipv4 (1)"},
		},
		{
			name:     "processed_summary.json",
			filename: "processed_summary.json",
			createFile: func(path string) error {
				processedSummaries := []c.ProcessedSummary{
					{
						Name: "test-source",
						ValidFiles: []c.ProcessedFile{
							{GenericSourceType: "domain"},
							{GenericSourceType: "ipv4"},
						},
						InvalidFiles: []c.ProcessedFile{
							{GenericSourceType: "domain"},
						},
					},
				}
				content, err := json.Marshal(processedSummaries)
				if err != nil {
					return err
				}
				return os.WriteFile(path, content, 0644)
			},
			expectStats:  true,
			expectedText: []string{"Sources:", "1 processed", "Files:", "2 valid, 1 invalid", "Types:", "domain (1)", "ipv4 (1)"},
		},
		{
			name:     "consolidated_summary.json",
			filename: "consolidated_summary.json",
			createFile: func(path string) error {
				consolidatedSummaries := []c.ConsolidatedSummary{
					{
						Type:       "domain",
						FilesCount: 5,
						Count:      1000,
					},
					{
						Type:       "ipv4",
						FilesCount: 3,
						Count:      500,
					},
				}
				content, err := json.Marshal(consolidatedSummaries)
				if err != nil {
					return err
				}
				return os.WriteFile(path, content, 0644)
			},
			expectStats:  true,
			expectedText: []string{"Files:", "8 consolidated", "Entries:", "1.5K total", "Types:", "domain (1)", "ipv4 (1)"},
		},
		{
			name:     "consolidated_categories_summary.json",
			filename: "consolidated_categories_summary.json",
			createFile: func(path string) error {
				categoriesSummaries := []c.ConsolidatedSummary{
					{
						Category:   "malware",
						FilesCount: 3,
						Count:      800,
					},
					{
						Category:   "ads",
						FilesCount: 2,
						Count:      200,
					},
				}
				content, err := json.Marshal(categoriesSummaries)
				if err != nil {
					return err
				}
				return os.WriteFile(path, content, 0644)
			},
			expectStats:  true,
			expectedText: []string{"Categories:", "2 processed", "Files:", "5 consolidated", "Entries:", "1.0K total", "Categories:", "ads (1)", "malware (1)"},
		},
		{
			name:     "consolidated_groups_summary.json",
			filename: "consolidated_groups_summary.json",
			createFile: func(path string) error {
				groupsSummaries := []c.ConsolidatedSummary{
					{
						Group:      "mini",
						FilesCount: 2,
						Count:      300,
					},
					{
						Group:      "lite",
						FilesCount: 4,
						Count:      700,
					},
				}
				content, err := json.Marshal(groupsSummaries)
				if err != nil {
					return err
				}
				return os.WriteFile(path, content, 0644)
			},
			expectStats:  true,
			expectedText: []string{"Groups:", "2 processed", "Files:", "6 consolidated", "Entries:", "1.0K total", "Groups:", "lite (1)", "mini (1)"},
		},
		{
			name:     "top_summary.json",
			filename: "top_summary.json",
			createFile: func(path string) error {
				topSummaries := []c.TopSummary{
					{
						GenericSourceType: "domain",
						MinSources:        3,
						Count:             500,
					},
					{
						GenericSourceType: "ipv4",
						MinSources:        2,
						Count:             250,
					},
				}
				content, err := json.Marshal(topSummaries)
				if err != nil {
					return err
				}
				return os.WriteFile(path, content, 0644)
			},
			expectStats:  true,
			expectedText: []string{"Types:", "2 analyzed", "Top Entries:", "0 total", "Details:", "domain (3+ min (0))", "ipv4 (2+ min (0))"},
		},
		{
			name:     "overlap_summary.json",
			filename: "overlap_summary.json",
			createFile: func(path string) error {
				overlapSummaries := []c.OverlapSummary{
					{
						Type:   "domain",
						Count:  1000,
						Unique: 800,
					},
					{
						Type:   "ipv4",
						Count:  500,
						Unique: 450,
					},
				}
				content, err := json.Marshal(overlapSummaries)
				if err != nil {
					return err
				}
				return os.WriteFile(path, content, 0644)
			},
			expectStats:  true,
			expectedText: []string{"Sources:", "2 analyzed", "Total Entries:", "1.5K", "Unique Entries:", "1.2K (83.3%)", "Types:", "domain (1)", "ipv4 (1)"},
		},
		{
			name:     "unknown file type",
			filename: "unknown_summary.json",
			createFile: func(path string) error {
				return os.WriteFile(path, []byte("{}"), 0644)
			},
			expectStats: false,
		},
		{
			name:     "invalid JSON",
			filename: "download_summary.json",
			createFile: func(path string) error {
				return os.WriteFile(path, []byte("invalid json"), 0644)
			},
			expectStats: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := filepath.Join(summariesDir, tt.filename)
			err := tt.createFile(filePath)
			require.NoError(t, err)

			defer func() {
				_ = os.Remove(filePath)
			}()

			stats := getDetailedStatsForSummaryType(tt.filename)

			if tt.expectStats {
				assert.NotEmpty(t, stats)
				for _, expectedText := range tt.expectedText {
					assert.Contains(t, stats, expectedText, "Expected text not found: %s", expectedText)
				}
			} else {
				assert.Empty(t, stats)
			}
		})
	}
}

func TestCollectOverallStatsFromFile(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "overall-stats-test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	tests := []struct {
		name            string
		filename        string
		createFile      func(string) error
		expectedUpdates func(*OverallSummaryStats) bool
		expectNoChange  bool
	}{
		{
			name:     "download_summary.json",
			filename: "download_summary.json",
			createFile: func(path string) error {
				downloadSummaries := []c.DownloadSummary{
					{Name: "source1"},
					{Name: "source2"},
					{Name: "source1"},
				}
				content, err := json.Marshal(downloadSummaries)
				if err != nil {
					return err
				}
				return os.WriteFile(path, content, 0644)
			},
			expectedUpdates: func(stats *OverallSummaryStats) bool {
				return stats.TotalDownloads == 3 && stats.TotalSources == 2
			},
		},
		{
			name:     "processed_summary.json",
			filename: "processed_summary.json",
			createFile: func(path string) error {
				processedSummaries := []c.ProcessedSummary{
					{Name: "source1"},
					{Name: "source2"},
				}
				content, err := json.Marshal(processedSummaries)
				if err != nil {
					return err
				}
				return os.WriteFile(path, content, 0644)
			},
			expectedUpdates: func(stats *OverallSummaryStats) bool {
				return stats.TotalProcessed == 2
			},
		},
		{
			name:     "consolidated_summary.json",
			filename: "consolidated_summary.json",
			createFile: func(path string) error {
				consolidatedSummaries := []c.ConsolidatedSummary{
					{Type: "domain"},
					{Type: "ipv4"},
				}
				content, err := json.Marshal(consolidatedSummaries)
				if err != nil {
					return err
				}
				return os.WriteFile(path, content, 0644)
			},
			expectedUpdates: func(stats *OverallSummaryStats) bool {
				return stats.TotalConsolidated == 2
			},
		},
		{
			name:     "consolidated_groups_summary.json",
			filename: "consolidated_groups_summary.json",
			createFile: func(path string) error {
				groupsSummaries := []c.ConsolidatedSummary{
					{Group: "mini"},
					{Group: "lite"},
					{Group: "mini"},
					{Group: ""},
				}
				content, err := json.Marshal(groupsSummaries)
				if err != nil {
					return err
				}
				return os.WriteFile(path, content, 0644)
			},
			expectedUpdates: func(stats *OverallSummaryStats) bool {
				return stats.TotalGroups == 2
			},
		},
		{
			name:     "consolidated_categories_summary.json",
			filename: "consolidated_categories_summary.json",
			createFile: func(path string) error {
				categoriesSummaries := []c.ConsolidatedSummary{
					{Category: "malware"},
					{Category: "ads"},
					{Category: "malware"},
					{Category: ""},
				}
				content, err := json.Marshal(categoriesSummaries)
				if err != nil {
					return err
				}
				return os.WriteFile(path, content, 0644)
			},
			expectedUpdates: func(stats *OverallSummaryStats) bool {
				return stats.TotalCategories == 2
			},
		},
		{
			name:     "top_summary.json",
			filename: "top_summary.json",
			createFile: func(path string) error {
				topSummaries := []c.TopSummary{
					{GenericSourceType: "domain"},
					{GenericSourceType: "ipv4"},
				}
				content, err := json.Marshal(topSummaries)
				if err != nil {
					return err
				}
				return os.WriteFile(path, content, 0644)
			},
			expectedUpdates: func(stats *OverallSummaryStats) bool {
				return stats.TotalTopLists == 2
			},
		},
		{
			name:     "overlap_summary.json",
			filename: "overlap_summary.json",
			createFile: func(path string) error {
				overlapSummaries := []c.OverlapSummary{
					{Source: "source1"},
					{Source: "source2"},
				}
				content, err := json.Marshal(overlapSummaries)
				if err != nil {
					return err
				}
				return os.WriteFile(path, content, 0644)
			},
			expectedUpdates: func(stats *OverallSummaryStats) bool {
				return stats.TotalOverlapAnalyzed == 2
			},
		},
		{
			name:     "unknown file type",
			filename: "unknown.json",
			createFile: func(path string) error {
				return os.WriteFile(path, []byte("{}"), 0644)
			},
			expectNoChange: true,
		},
		{
			name:     "invalid JSON",
			filename: "download_summary.json",
			createFile: func(path string) error {
				return os.WriteFile(path, []byte("invalid json"), 0644)
			},
			expectNoChange: true,
		},
		{
			name:     "non-existent file",
			filename: "non_existent.json",
			createFile: func(path string) error {
				return nil
			},
			expectNoChange: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := filepath.Join(tempDir, tt.filename)
			err := tt.createFile(filePath)
			require.NoError(t, err)

			defer func() {
				_ = os.Remove(filePath)
			}()

			stats := &OverallSummaryStats{}
			collectOverallStatsFromFile(filePath, tt.filename, stats)

			if tt.expectNoChange {
				assert.Equal(t, 0, stats.TotalSources)
				assert.Equal(t, 0, stats.TotalDownloads)
				assert.Equal(t, 0, stats.TotalProcessed)
				assert.Equal(t, 0, stats.TotalConsolidated)
				assert.Equal(t, 0, stats.TotalGroups)
				assert.Equal(t, 0, stats.TotalCategories)
				assert.Equal(t, 0, stats.TotalTopLists)
				assert.Equal(t, 0, stats.TotalOverlapAnalyzed)
			} else {
				assert.True(t, tt.expectedUpdates(stats), "Expected stats updates not found")
			}
		})
	}
}

// Helper function to create test summary files for summaries readme tests
func createTestSummaryFilesForSummariesReadme(t *testing.T, summariesDir string) {
	downloadSummaries := []c.DownloadSummary{
		{
			Name:                  "test-source-1",
			LastDownloadTimestamp: time.Now().Format(time.RFC3339),
			Types: []c.SourceType{
				{Name: "domain"},
				{Name: "ipv4"},
			},
		},
		{
			Name:                  "test-source-2",
			Error:                 "download failed",
			LastDownloadTimestamp: time.Now().Format(time.RFC3339),
			Types: []c.SourceType{
				{Name: "domain"},
			},
		},
	}
	content, err := json.Marshal(downloadSummaries)
	require.NoError(t, err)
	summaryFile := filepath.Join(summariesDir, "download_summary.json")
	require.NoError(t, os.WriteFile(summaryFile, content, 0644))

	processedSummaries := []c.ProcessedSummary{
		{
			Name:                   "test-source-1",
			LastProcessedTimestamp: time.Now().Format(time.RFC3339),
			ValidFiles: []c.ProcessedFile{
				{GenericSourceType: "domain"},
				{GenericSourceType: "ipv4"},
			},
		},
	}
	content, err = json.Marshal(processedSummaries)
	require.NoError(t, err)
	summaryFile = filepath.Join(summariesDir, "processed_summary.json")
	require.NoError(t, os.WriteFile(summaryFile, content, 0644))

	consolidatedSummaries := []c.ConsolidatedSummary{
		{
			Type:                      "domain",
			ListType:                  "blocklist",
			Count:                     1000,
			FilesCount:                5,
			LastConsolidatedTimestamp: time.Now().Format(time.RFC3339),
		},
	}
	content, err = json.Marshal(consolidatedSummaries)
	require.NoError(t, err)
	summaryFile = filepath.Join(summariesDir, "consolidated_summary.json")
	require.NoError(t, os.WriteFile(summaryFile, content, 0644))

	consolidatedGroupsSummaries := []c.ConsolidatedSummary{
		{
			Type:                      "domain",
			Group:                     "mini",
			Count:                     500,
			FilesCount:                2,
			LastConsolidatedTimestamp: time.Now().Format(time.RFC3339),
		},
	}
	content, err = json.Marshal(consolidatedGroupsSummaries)
	require.NoError(t, err)
	summaryFile = filepath.Join(summariesDir, "consolidated_groups_summary.json")
	require.NoError(t, os.WriteFile(summaryFile, content, 0644))

	consolidatedCategoriesSummaries := []c.ConsolidatedSummary{
		{
			Type:                      "domain",
			Category:                  "malware",
			Count:                     800,
			FilesCount:                3,
			LastConsolidatedTimestamp: time.Now().Format(time.RFC3339),
		},
	}
	content, err = json.Marshal(consolidatedCategoriesSummaries)
	require.NoError(t, err)
	summaryFile = filepath.Join(summariesDir, "consolidated_categories_summary.json")
	require.NoError(t, os.WriteFile(summaryFile, content, 0644))

	topSummaries := []c.TopSummary{
		{
			GenericSourceType: "domain",
			ListType:          "blocklist",
			MinSources:        3,
			Count:             400,
		},
	}
	content, err = json.Marshal(topSummaries)
	require.NoError(t, err)
	summaryFile = filepath.Join(summariesDir, "top_summary.json")
	require.NoError(t, os.WriteFile(summaryFile, content, 0644))

	overlapSummaries := []c.OverlapSummary{
		{
			Source: "test-source-1",
			Type:   "domain",
			Count:  1000,
			Unique: 800,
		},
	}
	content, err = json.Marshal(overlapSummaries)
	require.NoError(t, err)
	summaryFile = filepath.Join(summariesDir, "overlap_summary.json")
	require.NoError(t, os.WriteFile(summaryFile, content, 0644))
}
