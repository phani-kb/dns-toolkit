package config

import (
	"testing"
)

func TestAppConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		config  AppConfig
		wantErr bool
	}{
		{
			name: "Valid config",
			config: AppConfig{
				Application: ApplicationConfig{
					Name:        "TestApp",
					Version:     "1.0.0",
					Description: "Test application",
				},
				DNSToolkit: DNSToolkitConfig{
					SourceFiles: []string{"../../testdata/valid_sources.json"},
					MaxWorkers:  4,
				},
				Multilog: map[string]interface{}{
					"level": "info",
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid Application config",
			config: AppConfig{
				Application: ApplicationConfig{
					// Missing required fields
				},
				DNSToolkit: DNSToolkitConfig{
					SourceFiles: []string{"testdata/valid_sources.json"},
					MaxWorkers:  4,
				},
				Multilog: map[string]interface{}{
					"level": "info",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("AppConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApplicationConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		config  ApplicationConfig
		wantErr bool
	}{
		{
			name: "Valid config",
			config: ApplicationConfig{
				Name:        "TestApp",
				Version:     "1.0.0",
				Description: "Test application",
			},
			wantErr: false,
		},
		{
			name: "Missing name",
			config: ApplicationConfig{
				Version:     "1.0.0",
				Description: "Test application",
			},
			wantErr: true,
		},
		{
			name: "Missing version",
			config: ApplicationConfig{
				Name:        "TestApp",
				Description: "Test application",
			},
			wantErr: true,
		},
		{
			name: "Missing description",
			config: ApplicationConfig{
				Name:    "TestApp",
				Version: "1.0.0",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplicationConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
