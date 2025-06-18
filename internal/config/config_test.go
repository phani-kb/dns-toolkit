package config

import (
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
				Name: "test-source",
				URL:  "http://example.com",
			},
			wantErr: false,
		},
		{
			name: "Missing name",
			source: Source{
				URL: "http://example.com",
			},
			wantErr: true,
		},
		{
			name: "Empty name",
			source: Source{
				Name: "",
				URL:  "http://example.com",
			},
			wantErr: true,
		},
		{
			name: "Missing URL but valid (URL not validated)",
			source: Source{
				Name: "test-source",
			},
			wantErr: false,
		},
		{
			name: "Empty URL but valid (URL not validated)",
			source: Source{
				Name: "test-source",
				URL:  "",
			},
			wantErr: false,
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
				Name: "test-source",
				URL:  "http://example.com",
			},
			appConfig: appConfig,
			wantErr:   false,
		},
		{
			name: "Valid source with nil app config",
			source: Source{
				Name: "test-source",
				URL:  "http://example.com",
			},
			appConfig: nil,
			wantErr:   false,
		},
		{
			name: "Invalid source with app config",
			source: Source{
				Name: "",
				URL:  "http://example.com",
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

	logger := CreateTestLogger()
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
