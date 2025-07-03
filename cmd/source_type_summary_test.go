package cmd

import (
	"os"
	"path/filepath"
	"sync"
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSourceTypesSummaryCommand(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, sourceTypesSummaryCmd)
	assert.Equal(t, "sts", sourceTypesSummaryCmd.Use)
	assert.Contains(t, sourceTypesSummaryCmd.Short, "source types summary")
	assert.NotNil(t, sourceTypesSummaryCmd.Run)
}

func TestProcessSource(t *testing.T) {
	t.Parallel()

	tempDir, err := os.MkdirTemp("", "process-source-test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	// Store original download directory
	origDownloadDir := constants.DownloadDir
	defer func() {
		constants.DownloadDir = origDownloadDir
	}()

	downloadDir := filepath.Join(tempDir, "download")
	constants.DownloadDir = downloadDir
	require.NoError(t, os.MkdirAll(downloadDir, 0755))

	logger := multilog.NewLogger()

	tests := []struct {
		setupFunc      func() error
		expectedCounts map[string]int
		name           string
		source         config.Source
		expectedTotal  int
	}{
		{
			name: "invalid source with missing download file",
			source: config.Source{
				Name: "Invalid Source",
				URL:  "http://invalid.com/test.txt",
				Types: []c.SourceType{
					{Name: constants.SourceTypeDomain},
				},
				TypeCount: 1,
			},
			setupFunc:      func() error { return nil }, // Don't create any files
			expectedCounts: map[string]int{},
			expectedTotal:  0,
		},
		{
			name: "valid source with domain content",
			source: config.Source{
				Name: "Test Domain Source",
				URL:  "http://example.com/domains.txt",
				Types: []c.SourceType{
					{Name: constants.SourceTypeDomain},
				},
				TypeCount: 1,
			},
			setupFunc: func() error {
				downloadFile := filepath.Join(downloadDir, "Test Domain Source.txt")
				content := "example.com\nbadsite.net\nmalware.org\n"
				return os.WriteFile(downloadFile, []byte(content), 0644)
			},
			expectedCounts: map[string]int{constants.SourceTypeDomain: 1},
			expectedTotal:  1,
		},
		{
			name: "valid source with IPv4 content",
			source: config.Source{
				Name: "Test IPv4 Source",
				URL:  "http://example.com/ips.txt",
				Types: []c.SourceType{
					{Name: constants.SourceTypeIpv4},
				},
				TypeCount: 1,
			},
			setupFunc: func() error {
				downloadFile := filepath.Join(downloadDir, "Test IPv4 Source.txt")
				content := "192.168.1.1\n10.0.0.1\n172.16.0.1\n"
				return os.WriteFile(downloadFile, []byte(content), 0644)
			},
			expectedCounts: map[string]int{constants.SourceTypeIpv4: 1},
			expectedTotal:  1,
		},
		{
			name: "source with multiple types",
			source: config.Source{
				Name: "Multi Type Source",
				URL:  "http://example.com/mixed.txt",
				Types: []c.SourceType{
					{Name: constants.SourceTypeDomain},
					{Name: constants.SourceTypeIpv4},
				},
				TypeCount: 2,
			},
			setupFunc: func() error {
				downloadFile := filepath.Join(downloadDir, "Multi Type Source.txt")
				content := "example.com\n192.168.1.1\nbadsite.net\n10.0.0.1\n"
				return os.WriteFile(downloadFile, []byte(content), 0644)
			},
			expectedCounts: map[string]int{
				constants.SourceTypeDomain: 1,
				constants.SourceTypeIpv4:   1,
			}, // Multi-type sources count each type
			expectedTotal: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setupFunc()
			require.NoError(t, err)

			counts := make(map[string]int)
			var mismatches []string
			var mu sync.Mutex

			processSource(logger, tt.source, counts, &mismatches, &mu)

			assert.Equal(t, len(tt.expectedCounts), len(counts))
			for expectedType, expectedCount := range tt.expectedCounts {
				assert.Equal(t, expectedCount, counts[expectedType], "Count mismatch for type %s", expectedType)
			}

			// Calculate total
			total := 0
			for _, count := range counts {
				total += count
			}
			assert.Equal(t, tt.expectedTotal, total)
		})
	}
}
