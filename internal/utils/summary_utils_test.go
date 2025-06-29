package utils

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	c "github.com/phani-kb/dns-toolkit/internal/common"
)

func TestGetLastSummary(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()

	summary, err := GetLastSummary[c.DownloadSummary](logger, "/nonexistent/file", "test-source")
	assert.NoError(t, err)
	assert.Equal(t, c.DownloadSummary{}, summary)

	tempFile, err := os.CreateTemp("", "test_summary")
	require.NoError(t, err)
	defer func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Logf("Failed to remove temp file: %v", err)
		}
	}()

	err = os.WriteFile(tempFile.Name(), []byte("[]"), 0644)
	require.NoError(t, err)

	summary, err = GetLastSummary[c.DownloadSummary](logger, tempFile.Name(), "test-source")
	assert.NoError(t, err)
	assert.Equal(t, c.DownloadSummary{}, summary)

	summaries := []c.DownloadSummary{
		{
			Name:                  "test-source",
			LastDownloadTimestamp: "2023-06-01T10:00:00Z",
			Frequency:             "daily",
		},
		{
			Name:                  "other-source",
			LastDownloadTimestamp: "2023-06-01T11:00:00Z",
			Frequency:             "weekly",
		},
	}

	data, err := json.Marshal(summaries)
	require.NoError(t, err)

	err = os.WriteFile(tempFile.Name(), data, 0644)
	require.NoError(t, err)

	summary, err = GetLastSummary[c.DownloadSummary](logger, tempFile.Name(), "test-source")
	assert.NoError(t, err)
	assert.Equal(t, "test-source", summary.Name)
	assert.Equal(t, "daily", summary.Frequency)

	summary, err = GetLastSummary[c.DownloadSummary](logger, tempFile.Name(), "non-existent")
	assert.NoError(t, err)
	assert.Equal(t, c.DownloadSummary{}, summary)

	err = os.WriteFile(tempFile.Name(), []byte("invalid json"), 0644)
	require.NoError(t, err)

	summary, err = GetLastSummary[c.DownloadSummary](logger, tempFile.Name(), "test-source")
	assert.Error(t, err)
	assert.Equal(t, c.DownloadSummary{}, summary)

	processedSummaries := []c.ProcessedSummary{
		{
			Name: "processed-source",
		},
	}

	processedData, err := json.Marshal(processedSummaries)
	require.NoError(t, err)

	err = os.WriteFile(tempFile.Name(), processedData, 0644)
	require.NoError(t, err)

	processedSummary, err := GetLastSummary[c.ProcessedSummary](logger, tempFile.Name(), "processed-source")
	assert.NoError(t, err)
	assert.Equal(t, "processed-source", processedSummary.Name)
}

func TestGetSummaryFiles(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()

	// Test with non-existent file
	files, err := GetSummaryFiles(
		logger,
		"domain",
		"/nonexistent/file",
		func(summary c.DownloadSummary, sourceType string) []string {
			return []string{"test.txt"}
		},
	)
	assert.Error(t, err)
	assert.Nil(t, files)

	// Create a temporary file with test data
	tempFile, err := os.CreateTemp("", "test_summary_files")
	require.NoError(t, err)
	defer func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Logf("Failed to remove temp file: %v", err)
		}
	}()

	summaries := []c.DownloadSummary{
		{
			Name: "test-source1",
		},
		{
			Name: "test-source2",
		},
	}

	data, err := json.Marshal(summaries)
	require.NoError(t, err)

	err = os.WriteFile(tempFile.Name(), data, 0644)
	require.NoError(t, err)

	files, err = GetSummaryFiles(
		logger,
		"domain",
		tempFile.Name(),
		func(summary c.DownloadSummary, sourceType string) []string {
			return []string{summary.Name + "_" + sourceType + ".txt"}
		},
	)
	assert.NoError(t, err)
	assert.Len(t, files, 2)
	assert.Contains(t, files, "test-source1_domain.txt")
	assert.Contains(t, files, "test-source2_domain.txt")

	files, err = GetSummaryFiles(
		logger,
		"ipv4",
		tempFile.Name(),
		func(summary c.DownloadSummary, sourceType string) []string {
			return []string{}
		},
	)
	assert.NoError(t, err)
	assert.Len(t, files, 0)

	err = os.WriteFile(tempFile.Name(), []byte("invalid json"), 0644)
	require.NoError(t, err)

	files, err = GetSummaryFiles(
		logger,
		"domain",
		tempFile.Name(),
		func(summary c.DownloadSummary, sourceType string) []string {
			return []string{"test.txt"}
		},
	)
	assert.Error(t, err)
	assert.Nil(t, files)
}

func TestGetSummaryTypeFromFolder(t *testing.T) {
	tests := []struct {
		folderName   string
		expectedType string
	}{
		{"download", "download"},
		{"processed", "processed"},
		{"consolidated", "consolidated"},
		{"consolidated_groups", "consolidated_groups"},
		{"consolidated_categories", "consolidated_categories"},
		{"top", "top"},
		{"unknown_folder", "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.folderName, func(t *testing.T) {
			got := GetSummaryTypeFromFolder(tt.folderName)
			assert.Equal(t, tt.expectedType, got)
		})
	}
}

func TestDetermineSummaryTypeFromPath(t *testing.T) {
	tests := []struct {
		path         string
		expectedType string
	}{
		{"/tmp/download_summary.json", "download"},
		{"/tmp/processed_summary.json", "processed"},
		{"/tmp/consolidated_summary.json", "consolidated"},
		{"/tmp/consolidated_groups_summary.json", "consolidated_groups"},
		{"/tmp/consolidated_categories_summary.json", "consolidated_categories"},
		{"/tmp/top_summary.json", "top"},
		{"/tmp/unknown_summary.json", "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			got := DetermineSummaryTypeFromPath(tt.path)
			assert.Equal(t, tt.expectedType, got)
		})
	}
}
