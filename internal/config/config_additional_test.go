package config

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

func TestGetGenericSourceType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		sourceType string
		expected   string
	}{
		{
			name:       "Exact alias match",
			sourceType: "domain",
			expected:   "domain",
		},
		{
			name:       "Prefix match",
			sourceType: "domain_malware",
			expected:   "domain",
		},
		{
			name:       "No match returns original",
			sourceType: "unknown_type",
			expected:   "unknown_type",
		},
		{
			name:       "Empty string",
			sourceType: "",
			expected:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetGenericSourceType(tt.sourceType)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsEnabledSource(t *testing.T) {
	t.Parallel()

	sourceConfig := SourcesConfig{
		Sources: []Source{
			{
				Name:     "enabled-source",
				URL:      "http://example.com",
				Types:    []c.SourceType{{Name: "domain"}},
				Disabled: false,
			},
			{
				Name:     "disabled-source",
				URL:      "http://example.com",
				Types:    []c.SourceType{{Name: "domain"}},
				Disabled: true,
			},
		},
	}

	appConfig := AppConfig{
		DNSToolkit: DNSToolkitConfig{
			SourceFilters: SourceFilters{},
		},
	}

	tests := []struct {
		name       string
		sourceName string
		expected   bool
	}{
		{
			name:       "Enabled source found",
			sourceName: "enabled-source",
			expected:   true,
		},
		{
			name:       "Disabled source not found",
			sourceName: "disabled-source",
			expected:   false,
		},
		{
			name:       "Non-existent source",
			sourceName: "non-existent",
			expected:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEnabledSource(tt.sourceName, []SourcesConfig{sourceConfig}, appConfig)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetProcessedSummariesForConsolidation(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)

	tempDir, err := os.MkdirTemp("", "test_summaries_*")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temp directory: %v", err)
		}
	}()

	originalSummaryDir := constants.SummaryDir
	constants.SummaryDir = tempDir
	defer func() {
		constants.SummaryDir = originalSummaryDir
	}()

	testSummaries := []c.ProcessedSummary{
		{
			Name: "enabled-source",
			ValidFiles: []c.ProcessedFile{
				{
					GenericSourceType: "domain",
				},
			},
			InvalidFiles: []c.ProcessedFile{
				{
					GenericSourceType: "ip",
				},
			},
		},
		{
			Name: "disabled-source",
			ValidFiles: []c.ProcessedFile{
				{
					GenericSourceType: "domain",
				},
			},
		},
	}

	summaryFile := filepath.Join(tempDir, constants.DefaultSummaryFiles["processed"])
	err = os.MkdirAll(filepath.Dir(summaryFile), 0755)
	require.NoError(t, err)

	summaryData, err := json.Marshal(testSummaries)
	require.NoError(t, err)

	err = os.WriteFile(summaryFile, summaryData, 0644)
	require.NoError(t, err)

	sourceConfig := SourcesConfig{
		Sources: []Source{
			{
				Name:     "enabled-source",
				URL:      "http://example.com",
				Types:    []c.SourceType{{Name: "domain"}},
				Disabled: false,
			},
			{
				Name:     "disabled-source",
				URL:      "http://example.com",
				Types:    []c.SourceType{{Name: "domain"}},
				Disabled: true,
			},
		},
	}

	appConfig := AppConfig{
		DNSToolkit: DNSToolkitConfig{
			SourceFilters: SourceFilters{},
		},
	}

	summaries, genericTypes, processedFiles := GetProcessedSummariesForConsolidation(
		logger,
		[]SourcesConfig{sourceConfig},
		appConfig,
		"general",
	)

	assert.Len(t, summaries, 1) // Only enabled source
	assert.Equal(t, "enabled-source", summaries[0].Name)
	assert.Contains(t, genericTypes, "domain")
	assert.Contains(t, genericTypes, "ip")
	assert.Len(t, processedFiles, 2)

	err = os.Remove(summaryFile)
	require.NoError(t, err)

	summaries, genericTypes, processedFiles = GetProcessedSummariesForConsolidation(
		logger,
		[]SourcesConfig{sourceConfig},
		appConfig,
		"general",
	)
	assert.Nil(t, summaries)
	assert.Nil(t, genericTypes)
	assert.Nil(t, processedFiles)

	err = os.WriteFile(summaryFile, []byte("invalid json"), 0644)
	require.NoError(t, err)

	summaries, genericTypes, processedFiles = GetProcessedSummariesForConsolidation(
		logger,
		[]SourcesConfig{sourceConfig},
		appConfig,
		"general",
	)
	assert.Nil(t, summaries)
	assert.Nil(t, genericTypes)
	assert.Nil(t, processedFiles)
}

func TestGetAllProcessedFiles(t *testing.T) {
	t.Parallel()

	testSummaries := []c.ProcessedSummary{
		{
			Name: "source1",
			ValidFiles: []c.ProcessedFile{
				{GenericSourceType: "domain"},
				{GenericSourceType: "ip"},
			},
			InvalidFiles: []c.ProcessedFile{
				{GenericSourceType: "domain"},
			},
		},
		{
			Name: "source2",
			ValidFiles: []c.ProcessedFile{
				{GenericSourceType: "url"},
			},
			InvalidFiles: []c.ProcessedFile{},
		},
	}

	files := GetAllProcessedFiles(testSummaries)

	assert.Len(t, files, 4) // 3 from source1 + 1 from source2
}

func TestFilterEnabledSummariesForConsolidation(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)

	testSummaries := []c.ProcessedSummary{
		{Name: "enabled-source"},
		{Name: "disabled-source"},
		{Name: "non-existent"},
	}

	sourceConfig := SourcesConfig{
		Sources: []Source{
			{
				Name:     "enabled-source",
				URL:      "http://example.com",
				Types:    []c.SourceType{{Name: "domain"}},
				Disabled: false,
			},
			{
				Name:     "disabled-source",
				URL:      "http://example.com",
				Types:    []c.SourceType{{Name: "domain"}},
				Disabled: true,
			},
		},
	}

	appConfig := AppConfig{
		DNSToolkit: DNSToolkitConfig{
			SourceFilters: SourceFilters{},
		},
	}

	enabled := filterEnabledSummariesForConsolidation(
		logger,
		testSummaries,
		[]SourcesConfig{sourceConfig},
		appConfig,
		"general",
	)

	assert.Len(t, enabled, 1)
	assert.Equal(t, "enabled-source", enabled[0].Name)
}

func TestExtractGenericSourceTypes(t *testing.T) {
	t.Parallel()

	testSummaries := []c.ProcessedSummary{
		{
			ValidFiles: []c.ProcessedFile{
				{GenericSourceType: "domain"},
				{GenericSourceType: "ip"},
			},
			InvalidFiles: []c.ProcessedFile{
				{GenericSourceType: "domain"}, // Duplicate
				{GenericSourceType: "url"},
			},
		},
		{
			ValidFiles: []c.ProcessedFile{
				{GenericSourceType: "ip"}, // Duplicate
			},
		},
	}

	types := extractGenericSourceTypes(testSummaries)

	assert.Len(t, types, 3)
	assert.Contains(t, types, "domain")
	assert.Contains(t, types, "ip")
	assert.Contains(t, types, "url")

	for i := 1; i < len(types); i++ {
		assert.True(t, types[i-1] < types[i], "Types should be sorted")
	}
}

func TestDNSToolkitConfigValidateEdgeCases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		config  DNSToolkitConfig
		wantErr bool
		setup   func()
		cleanup func()
	}{
		{
			name: "Config with non-existent source file",
			config: DNSToolkitConfig{
				SourceFiles: []string{"/non/existent/file.json"},
				MaxWorkers:  4,
			},
			wantErr: true,
		},
		{
			name: "Config with existing source file",
			config: DNSToolkitConfig{
				SourceFiles: []string{},
				MaxWorkers:  4,
			},
			wantErr: false,
			setup: func() {
			},
		},
		{
			name: "Config with max workers exceeding GOMAXPROCS",
			config: DNSToolkitConfig{
				SourceFiles: []string{},
				MaxWorkers:  1000,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if tt.name == "Config with existing source file" {
					tempFile, err := os.CreateTemp("", "test_source_*.json")
					require.NoError(t, err)
					defer func() {
						if err := os.Remove(tempFile.Name()); err != nil {
							t.Logf("Failed to remove temp file: %v", err)
						}
					}()
					defer func() {
						if err := tempFile.Close(); err != nil {
							t.Logf("Failed to close temp file: %v", err)
						}
					}()

					tt.config.SourceFiles = []string{tempFile.Name()}
				}
			}

			err := tt.config.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}

func TestLoadAppConfigEdgeCases(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)

	_, _, err := LoadAppConfig(logger, "/non/existent/config.yml")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "opening config file")

	tempFile, err := os.CreateTemp("", "invalid_config_*.yml")
	require.NoError(t, err)
	defer func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Logf("Failed to remove temp file: %v", err)
		}
	}()
	defer func() {
		if err := tempFile.Close(); err != nil {
			t.Logf("Failed to close temp file: %v", err)
		}
	}()

	_, err = tempFile.WriteString("invalid: yaml: content: [")
	require.NoError(t, err)
	err = tempFile.Close()
	require.NoError(t, err)

	_, _, err = LoadAppConfig(logger, tempFile.Name())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "decoding config file")

	validYAML := `
application:
  name: TestApp
  version: 0.4.0
  description: Test application
dns_toolkit:
  source_files: []
  max_workers: 2
multilog:
  level: info
`

	tempFile2, err := os.CreateTemp("", "valid_config_*.yml")
	require.NoError(t, err)
	defer func() {
		if err := os.Remove(tempFile2.Name()); err != nil {
			t.Logf("Failed to remove temp file: %v", err)
		}
	}()
	defer func() {
		if err := tempFile2.Close(); err != nil {
			t.Logf("Failed to close temp file: %v", err)
		}
	}()

	_, err = tempFile2.WriteString(validYAML)
	require.NoError(t, err)
	err = tempFile2.Close()
	require.NoError(t, err)

	_, _, err = LoadAppConfig(logger, tempFile2.Name())
	assert.Error(t, err)
}
