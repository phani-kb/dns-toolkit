package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateOverlapCmdRun(t *testing.T) {
	err := os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	require.NoError(t, err)
	defer func() {
		err := os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
		if err != nil {
			t.Logf("Failed to unset DNS_TOOLKIT_TEST_MODE: %v", err)
		}
	}()

	generateOverlapCmd.Run(generateOverlapCmd, []string{})
}

func TestGenerateDetailedOverlapAnalysis(t *testing.T) {
	t.Parallel()

	tempDir, err := os.MkdirTemp("", "overlap-test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	summaryDir := filepath.Join(tempDir, "summaries")
	outputDir := filepath.Join(tempDir, "output")
	require.NoError(t, os.MkdirAll(summaryDir, 0755))
	require.NoError(t, os.MkdirAll(outputDir, 0755))

	origSummaryDir := constants.SummaryDir
	origOutputDir := constants.OutputDir

	constants.SummaryDir = summaryDir
	constants.OutputDir = outputDir

	defer func() {
		constants.SummaryDir = origSummaryDir
		constants.OutputDir = origOutputDir
	}()

	_, err = generateDetailedOverlapAnalysis()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overlap summary file not found")

	createTestOverlapSummary(t, summaryDir)

	overlapMd, err := generateDetailedOverlapAnalysis()
	require.NoError(t, err)
	assert.NotEmpty(t, overlapMd)

	assert.Contains(t, overlapMd, "# DNS Toolkit - Detailed Overlap Analysis")
	assert.Contains(t, overlapMd, "## Overview")
	assert.Contains(t, overlapMd, "## Detailed Source Analysis")
	assert.Contains(t, overlapMd, "### test-source-1")
	assert.Contains(t, overlapMd, "### test-source-2")
	assert.Contains(t, overlapMd, "**Last Updated:**")
	assert.Contains(t, overlapMd, "## About")

	emptyOverlapSummaries := []c.OverlapSummary{}
	content, err := json.Marshal(emptyOverlapSummaries)
	require.NoError(t, err)
	summaryFile := filepath.Join(summaryDir, constants.DefaultSummaryFiles["overlap"])
	require.NoError(t, os.WriteFile(summaryFile, content, 0644))

	_, err = generateDetailedOverlapAnalysis()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no overlap data found")
}

func TestParseTargetString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		expected *TargetDetail
		name     string
		input    string
	}{
		{
			name:  "valid target string",
			input: "abpvn_hosts, lt: blocklist, type: adguard, count: 1051, overlap: 10, percent: 1.0",
			expected: &TargetDetail{
				Name:           "abpvn_hosts",
				ListType:       "blocklist",
				SourceType:     "adguard",
				Count:          1051,
				OverlapCount:   10,
				OverlapPercent: 1.0,
			},
		},
		{
			name:  "target string with allowlist",
			input: "test_source, lt: allowlist, type: domain, count: 500, overlap: 25, percent: 5.0",
			expected: &TargetDetail{
				Name:           "test_source",
				ListType:       "allowlist",
				SourceType:     "domain",
				Count:          500,
				OverlapCount:   25,
				OverlapPercent: 5.0,
			},
		},
		{
			name:     "invalid target string with insufficient parts",
			input:    "incomplete, lt: blocklist",
			expected: nil,
		},
		{
			name:     "empty string",
			input:    "",
			expected: nil,
		},
		{
			name:  "target string with zero values",
			input: "zero_source, lt: blocklist, type: domain, count: 0, overlap: 0, percent: 0.0",
			expected: &TargetDetail{
				Name:           "zero_source",
				ListType:       "blocklist",
				SourceType:     "domain",
				Count:          0,
				OverlapCount:   0,
				OverlapPercent: 0.0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseTargetString(tt.input)
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				require.NotNil(t, result)
				assert.Equal(t, tt.expected.Name, result.Name)
				assert.Equal(t, tt.expected.ListType, result.ListType)
				assert.Equal(t, tt.expected.SourceType, result.SourceType)
				assert.Equal(t, tt.expected.Count, result.Count)
				assert.Equal(t, tt.expected.OverlapCount, result.OverlapCount)
				assert.Equal(t, tt.expected.OverlapPercent, result.OverlapPercent)
			}
		})
	}
}

func TestParseIntFromString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input       string
		expected    int
		expectError bool
	}{
		{"123", 123, false},
		{"0", 0, false},
		{"999999", 999999, false},
		{"abc", 0, true},
		{"12.34", 12, false}, // Parses only the integer part
		{"", 0, true},
		{"-456", -456, false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := parseIntFromString(tt.input)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestParseFloatFromString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input       string
		expected    float64
		expectError bool
	}{
		{"123.45", 123.45, false},
		{"0.0", 0.0, false},
		{"999.999", 999.999, false},
		{"abc", 0.0, true},
		{"", 0.0, true},
		{"-456.78", -456.78, false},
		{"123", 123.0, false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := parseFloatFromString(tt.input)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestOverlapAnalysisWithComplexData(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "complex-overlap-test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	summaryDir := filepath.Join(tempDir, "summaries")
	require.NoError(t, os.MkdirAll(summaryDir, 0755))

	origSummaryDir := constants.SummaryDir
	constants.SummaryDir = summaryDir
	defer func() {
		constants.SummaryDir = origSummaryDir
	}()

	overlapSummaries := []c.OverlapSummary{
		{
			Type:         "domain",
			Source:       "source_a",
			ListType:     "blocklist",
			Count:        1000,
			Unique:       800,
			TargetsCount: 2,
			TargetsList: []string{
				"source_b, lt: blocklist, type: domain, count: 1500, overlap: 150, percent: 15.0",
				"source_c, lt: allowlist, type: domain, count: 500, overlap: 50, percent: 5.0",
			},
		},
		{
			Type:         "ipv4",
			Source:       "source_d",
			ListType:     "blocklist",
			Count:        2000,
			Unique:       1800,
			TargetsCount: 1,
			TargetsList: []string{
				"source_e, lt: blocklist, type: ipv4, count: 3000, overlap: 200, percent: 10.0",
			},
		},
		{
			Type:         "domain",
			Source:       "source_f",
			ListType:     "allowlist",
			Count:        500,
			Unique:       500,
			TargetsCount: 0,
			TargetsList:  []string{},
		},
	}

	content, err := json.Marshal(overlapSummaries)
	require.NoError(t, err)
	summaryFile := filepath.Join(summaryDir, constants.DefaultSummaryFiles["overlap"])
	require.NoError(t, os.WriteFile(summaryFile, content, 0644))

	overlapMd, err := generateDetailedOverlapAnalysis()
	require.NoError(t, err)
	assert.NotEmpty(t, overlapMd)

	assert.Contains(t, overlapMd, "Total Sources Analyzed | 3")
	assert.Contains(t, overlapMd, "source_a")
	assert.Contains(t, overlapMd, "source_d")
	assert.Contains(t, overlapMd, "source_f")
	assert.Contains(t, overlapMd, "| blocklist | 2 |")
	assert.Contains(t, overlapMd, "| allowlist | 1 |")
	assert.Contains(t, overlapMd, "| domain | 2 |")
	assert.Contains(t, overlapMd, "| ipv4 | 1 |")
	assert.Contains(t, overlapMd, "*No overlaps found with other sources.*")
}

func TestFormatNumberInOverlap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    int
		expected string
	}{
		{0, "0"},
		{500, "500"},
		{1000, "1.0K"},
		{1500, "1.5K"},
		{1000000, "1.0M"},
		{2500000, "2.5M"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := formatNumber(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Helper function to create test overlap summary
func createTestOverlapSummary(t *testing.T, summaryDir string) {
	overlapSummaries := []c.OverlapSummary{
		{
			Type:         "domain",
			Source:       "test-source-1",
			ListType:     "blocklist",
			Count:        1000,
			Unique:       900,
			TargetsCount: 1,
			TargetsList: []string{
				"test-source-2, lt: blocklist, type: domain, count: 1500, overlap: 100, percent: 10.0",
			},
		},
		{
			Type:         "ipv4",
			Source:       "test-source-2",
			ListType:     "blocklist",
			Count:        1500,
			Unique:       1400,
			TargetsCount: 1,
			TargetsList: []string{
				"test-source-1, lt: blocklist, type: domain, count: 1000, overlap: 100, percent: 6.7",
			},
		},
	}

	content, err := json.Marshal(overlapSummaries)
	require.NoError(t, err)
	summaryFile := filepath.Join(summaryDir, constants.DefaultSummaryFiles["overlap"])
	require.NoError(t, os.WriteFile(summaryFile, content, 0644))
}

func TestFilterOverlapSummariesByMinPercent(t *testing.T) {
	overlapSummaries := []c.OverlapSummary{
		{
			Type:         "domain",
			Source:       "source_a",
			ListType:     "blocklist",
			Count:        1000,
			Unique:       800,
			TargetsCount: 2,
			TargetsList: []string{
				"source_b, lt: blocklist, type: domain, count: 1500, overlap: 150, percent: 15.0",
				"source_c, lt: allowlist, type: domain, count: 500, overlap: 50, percent: 5.0",
			},
		},
		{
			Type:         "domain",
			Source:       "source_d",
			ListType:     "blocklist",
			Count:        500,
			Unique:       400,
			TargetsCount: 1,
			TargetsList: []string{
				"source_e, lt: blocklist, type: domain, count: 700, overlap: 20, percent: 2.0",
			},
		},
	}

	minPercent := 10.0
	origMin := AppConfig.DNSToolkit.MinOverlapPercent
	AppConfig.DNSToolkit.MinOverlapPercent = minPercent
	defer func() { AppConfig.DNSToolkit.MinOverlapPercent = origMin }()

	filteredSummaries := FilterOverlapSummariesByMinPercent(
		overlapSummaries,
		AppConfig.DNSToolkit.GetMinOverlapPercent(),
	)

	require.Len(t, filteredSummaries, 1)
	require.Len(t, filteredSummaries[0].TargetsList, 1)
	assert.Contains(t, filteredSummaries[0].TargetsList[0], "percent: 15.0")
	assert.NotContains(t, filteredSummaries[0].TargetsList[0], "percent: 5.0")
}
