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
	defer os.Remove(tempFile.Name())

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
	defer os.Remove(tempFile.Name())

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
