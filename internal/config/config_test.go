package config

import (
	"encoding/json"
	"os"
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestSourcesConfigValidate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		config  SourcesConfig
		wantErr bool
	}{
		{
			name: "Valid sources config",
			config: SourcesConfig{
				Sources: []Source{
					{
						Name:  "test-source",
						URL:   "http://example.com",
						Types: []c.SourceType{{Name: "domain"}},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Empty sources",
			config: SourcesConfig{
				Sources: []Source{},
			},
			wantErr: true,
		},
		{
			name: "Nil sources",
			config: SourcesConfig{
				Sources: nil,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSourcesConfigValidateWithConfig(t *testing.T) {
	t.Parallel()

	appConfig := &AppConfig{
		Application: ApplicationConfig{
			Name:        "TestApp",
			Version:     "1.0.0",
			Description: "Test application",
		},
	}

	tests := []struct {
		name      string
		config    SourcesConfig
		appConfig *AppConfig
		wantErr   bool
	}{
		{
			name: "Valid sources with app config",
			config: SourcesConfig{
				Sources: []Source{
					{
						Name:  "test-source",
						URL:   "http://example.com",
						Types: []c.SourceType{{Name: "domain"}},
					},
				},
			},
			appConfig: appConfig,
			wantErr:   false,
		},
		{
			name: "Valid sources with nil app config",
			config: SourcesConfig{
				Sources: []Source{
					{
						Name:  "test-source",
						URL:   "http://example.com",
						Types: []c.SourceType{{Name: "domain"}},
					},
				},
			},
			appConfig: nil,
			wantErr:   false,
		},
		{
			name: "Empty sources with app config",
			config: SourcesConfig{
				Sources: []Source{},
			},
			appConfig: appConfig,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.ValidateWithConfig(tt.appConfig)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSourceValidate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		source  Source
		wantErr bool
	}{
		{
			name: "Valid source",
			source: Source{
				Name:  "test-source",
				URL:   "http://example.com",
				Types: []c.SourceType{{Name: "domain"}},
			},
			wantErr: false,
		},
		{
			name: "Missing name",
			source: Source{
				URL:   "http://example.com",
				Types: []c.SourceType{{Name: "domain"}},
			},
			wantErr: true,
		},
		{
			name: "Empty name",
			source: Source{
				Name:  "",
				URL:   "http://example.com",
				Types: []c.SourceType{{Name: "domain"}},
			},
			wantErr: true,
		},
		{
			name: "Missing URL",
			source: Source{
				Name:  "test-source",
				Types: []c.SourceType{{Name: "domain"}},
			},
			wantErr: true,
		},
		{
			name: "Empty URL",
			source: Source{
				Name:  "test-source",
				URL:   "",
				Types: []c.SourceType{{Name: "domain"}},
			},
			wantErr: true,
		},
		{
			name: "Missing types",
			source: Source{
				Name: "test-source",
				URL:  "http://example.com",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.source.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSourceValidateWithConfig(t *testing.T) {
	t.Parallel()

	appConfig := &AppConfig{
		Application: ApplicationConfig{
			Name:        "TestApp",
			Version:     "1.0.0",
			Description: "Test application",
		},
	}

	tests := []struct {
		name      string
		source    Source
		appConfig *AppConfig
		wantErr   bool
	}{
		{
			name: "Valid source with app config",
			source: Source{
				Name:  "test-source",
				URL:   "http://example.com",
				Types: []c.SourceType{{Name: "domain"}},
			},
			appConfig: appConfig,
			wantErr:   false,
		},
		{
			name: "Valid source with nil app config",
			source: Source{
				Name:  "test-source",
				URL:   "http://example.com",
				Types: []c.SourceType{{Name: "domain"}},
			},
			appConfig: nil,
			wantErr:   false,
		},
		{
			name: "Invalid source with app config",
			source: Source{
				Name:  "",
				URL:   "http://example.com",
				Types: []c.SourceType{{Name: "domain"}},
			},
			appConfig: appConfig,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.source.ValidateWithConfig(tt.appConfig)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSourcesConfigGetEnabledSources(t *testing.T) {
	t.Parallel()

	config := SourcesConfig{
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

	// Empty SourceFilters
	filters := SourceFilters{}

	enabled := config.GetEnabledSources(filters)
	assert.Len(t, enabled, 1)
	assert.Equal(t, "enabled-source", enabled[0].Name)
}

func TestSourcesConfigGetSourceByName(t *testing.T) {
	t.Parallel()

	config := SourcesConfig{
		Sources: []Source{
			{Name: "test-source-1", URL: "http://example.com"},
			{Name: "test-source-2", URL: "http://example.org"},
		},
	}

	source, found := config.GetSourceByName("test-source-1")
	assert.True(t, found)
	assert.Equal(t, "test-source-1", source.Name)
	assert.Equal(t, "http://example.com", source.URL)

	// Non-existing source
	_, found = config.GetSourceByName("non-existent")
	assert.False(t, found)
}

func TestSourceGetDownloadFile(t *testing.T) {
	t.Parallel()

	source := Source{
		Name: "test-source",
		URL:  "http://example.com/file.txt",
	}

	logger := createTestLogger(t)
	downloadDir := "/tmp/downloads"

	downloadFile, err := source.GetDownloadFile(logger, downloadDir)
	assert.NoError(t, err)
	assert.Equal(t, "test-source", downloadFile.Name)
	assert.Equal(t, "http://example.com/file.txt", downloadFile.URL)
}

func TestValidateDNSToolkitAppConfig(t *testing.T) {
	t.Parallel()

	// Valid config without source files
	validConfig := AppConfig{
		Application: ApplicationConfig{
			Name:        "TestApp",
			Version:     "1.0.0",
			Description: "Test application",
		},
		DNSToolkit: DNSToolkitConfig{
			MaxWorkers: 4,
		},
		Multilog: map[string]interface{}{
			"level": "info",
		},
	}

	err := ValidateDNSToolkitAppConfig(validConfig)
	assert.Error(t, err)

	// Invalid config
	invalidConfig := AppConfig{
		Application: ApplicationConfig{
			// Missing required fields
		},
	}

	err = ValidateDNSToolkitAppConfig(invalidConfig)
	assert.Error(t, err)
}

func TestValidateAppConfig(t *testing.T) {
	t.Parallel()

	validConfig := AppConfig{
		Application: ApplicationConfig{
			Name:        "TestApp",
			Version:     "1.0.0",
			Description: "Test application",
		},
		DNSToolkit: DNSToolkitConfig{
			MaxWorkers: 4,
			// SourceFiles missing
		},
		Multilog: map[string]interface{}{
			"level": "info",
		},
	}

	err := ValidateAppConfig(validConfig)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "application validation error")

	// Invalid config
	invalidConfig := AppConfig{
		Application: ApplicationConfig{
			// Missing required fields
		},
	}

	err = ValidateAppConfig(invalidConfig)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "application validation error")
}

func TestSourcesConfigIsEnabled(t *testing.T) {
	t.Parallel()

	config := SourcesConfig{
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

	assert.True(t, config.IsEnabled("enabled-source"))
	assert.False(t, config.IsEnabled("disabled-source"))
	assert.False(t, config.IsEnabled("non-existent"))
}

func TestSourcesConfigGetSourceByField(t *testing.T) {
	t.Parallel()

	config := SourcesConfig{
		Sources: []Source{
			{
				Name:      "source1",
				URL:       "http://example.com",
				Frequency: "daily",
				Types: []c.SourceType{
					{
						Name: "domain",
						ListTypes: []c.ListType{
							{Name: "blocklist"},
						},
					},
				},
				Categories: []string{"malware"},
				Countries:  []string{"US"},
			},
			{
				Name:      "source2",
				URL:       "http://example.org",
				Frequency: "weekly",
				Types: []c.SourceType{
					{
						Name: "ip",
					},
				},
				Categories: []string{"phishing"},
				Countries:  []string{"CA"},
			},
		},
	}

	sources := config.GetSourceByField("listType", "blocklist")
	assert.Len(t, sources, 1)
	assert.Equal(t, "source1", sources[0].Name)

	sources = config.GetSourceByField("type", "domain")
	assert.Len(t, sources, 1)
	assert.Equal(t, "source1", sources[0].Name)

	sources = config.GetSourceByField("frequency", "daily")
	assert.Len(t, sources, 1)
	assert.Equal(t, "source1", sources[0].Name)

	sources = config.GetSourceByField("category", "malware")
	assert.Len(t, sources, 1)
	assert.Equal(t, "source1", sources[0].Name)

	sources = config.GetSourceByField("country", "US")
	assert.Len(t, sources, 1)
	assert.Equal(t, "source1", sources[0].Name)

	// Test no matches
	sources = config.GetSourceByField("type", "nonexistent")
	assert.Len(t, sources, 0)
}

func TestLoadSourcesConfig(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)

	_, err := LoadSourcesConfig(logger, "/nonexistent/file.json")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error opening sources file")

	// Test with invalid JSON
	tempFile, err := os.CreateTemp("", "invalid_*.json")
	assert.NoError(t, err)
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

	_, err = tempFile.WriteString("invalid json")
	assert.NoError(t, err)
	err = tempFile.Close()
	assert.NoError(t, err)

	_, err = LoadSourcesConfig(logger, tempFile.Name())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error unmarshalling JSON")

	// Test with valid JSON
	validJSON := `{
		"sources": [
			{
				"name": "test-source",
				"url": "http://example.com",
				"types": [
					{
						"name": "domain"
					}
				]
			}
		]
	}`

	tempFile2, err := os.CreateTemp("", "valid_*.json")
	assert.NoError(t, err)
	defer func() {
		if err := os.Remove(tempFile2.Name()); err != nil {
			t.Logf("Failed to remove temp file: %v", err)
		}
	}()
	defer func() {
		err := tempFile2.Close()
		if err != nil {
			t.Logf("Failed to close tempFile2: %v", err)
		}
	}()

	_, err = tempFile2.WriteString(validJSON)
	assert.NoError(t, err)
	err = tempFile2.Close()
	assert.NoError(t, err)

	config, err := LoadSourcesConfig(logger, tempFile2.Name())
	assert.NoError(t, err)
	assert.Len(t, config.Sources, 1)
	assert.Equal(t, "test-source", config.Sources[0].Name)
}

func TestSourceUnmarshalJSON(t *testing.T) {
	t.Parallel()

	jsonData := `{
		"name": "test-source",
		"url": "http://example.com",
		"files": "file1.txt, file2.txt ,file3.txt"
	}`

	var source Source
	err := json.Unmarshal([]byte(jsonData), &source)
	assert.NoError(t, err)
	assert.Equal(t, "test-source", source.Name)
	assert.Equal(t, "http://example.com", source.URL)
	assert.Len(t, source.Files, 3)
	assert.Equal(t, "file2.txt", source.Files[1]) // Should be trimmed

	// Test with missing files field
	jsonData2 := `{
		"name": "test-source",
		"url": "http://example.com"
	}`

	var source2 Source
	err = json.Unmarshal([]byte(jsonData2), &source2)
	assert.NoError(t, err)
	assert.Len(t, source2.Files, 0)

	// Test invalid JSON
	invalidJSON := `{invalid json`
	var source3 Source
	err = json.Unmarshal([]byte(invalidJSON), &source3)
	assert.Error(t, err)
}
