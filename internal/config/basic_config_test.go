package config

import (
	"os"
	"path/filepath"
	"testing"
)

func createTestConfigFiles(t *testing.T) string {
	testDir := filepath.Join("testdata")
	if err := os.MkdirAll(testDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create a valid sources file for testing
	validSourcesContent := `{
		"sources": [
			{
				"name": "test_source",
				"url": "http://test.com",
				"types": [{"name": "domain"}]
			}
		]
	}`
	validSourcesPath := filepath.Join(testDir, "valid_sources.json")
	if err := os.WriteFile(validSourcesPath, []byte(validSourcesContent), 0644); err != nil {
		t.Fatal(err)
	}

	validConfigContent := `
application:
  name: TestApp
  version: 1.0.0
  description: Test application

dns_toolkit:
  source_files:
    - testdata/valid_sources.json
  max_workers: 4

multilog:
  level: info
`
	validConfigPath := filepath.Join(testDir, "valid_config.yaml")
	if err := os.WriteFile(validConfigPath, []byte(validConfigContent), 0644); err != nil {
		t.Fatal(err)
	}

	return testDir
}

func TestBasicApplicationConfigValidate(t *testing.T) {
	// Test valid config
	validConfig := ApplicationConfig{
		Name:        "TestApp",
		Version:     "1.0.0",
		Description: "Test application",
	}
	if err := validConfig.Validate(); err != nil {
		t.Errorf("Valid ApplicationConfig should not return error: %v", err)
	}

	// Test missing name
	missingNameConfig := ApplicationConfig{
		Version:     "1.0.0",
		Description: "Test application",
	}
	if err := missingNameConfig.Validate(); err == nil {
		t.Error("ApplicationConfig with missing name should return error")
	}

	// Test missing version
	missingVersionConfig := ApplicationConfig{
		Name:        "TestApp",
		Description: "Test application",
	}
	if err := missingVersionConfig.Validate(); err == nil {
		t.Error("ApplicationConfig with missing version should return error")
	}

	// Test missing description
	missingDescriptionConfig := ApplicationConfig{
		Name:    "TestApp",
		Version: "1.0.0",
	}
	if err := missingDescriptionConfig.Validate(); err == nil {
		t.Error("ApplicationConfig with missing description should return error")
	}
}

func TestBasicDNSToolkitConfigValidate(t *testing.T) {
	testDir := createTestConfigFiles(t)
	defer os.RemoveAll(testDir)

	// Valid config
	validSourcesPath := filepath.Join(testDir, "valid_sources.json")
	validConfig := DNSToolkitConfig{
		SourceFiles: []string{validSourcesPath},
		MaxWorkers:  4,
	}
	if err := validConfig.Validate(); err != nil {
		t.Errorf("Valid DNSToolkitConfig should not return error: %v", err)
	}

	// Missing source files
	missingSourcesConfig := DNSToolkitConfig{
		MaxWorkers: 4,
	}
	if err := missingSourcesConfig.Validate(); err == nil {
		t.Error("DNSToolkitConfig with no source files should return error")
	}

	// Non-existent source file
	nonExistentConfig := DNSToolkitConfig{
		SourceFiles: []string{"non_existent.json"},
		MaxWorkers:  4,
	}
	if err := nonExistentConfig.Validate(); err == nil {
		t.Error("DNSToolkitConfig with non-existent source file should return error")
	}
}

func TestBasicAppConfigValidate(t *testing.T) {
	testDir := createTestConfigFiles(t)
	defer os.RemoveAll(testDir)

	validSourcesPath := filepath.Join(testDir, "valid_sources.json")

	// Valid config
	validConfig := AppConfig{
		Application: ApplicationConfig{
			Name:        "TestApp",
			Version:     "1.0.0",
			Description: "Test application",
		},
		DNSToolkit: DNSToolkitConfig{
			SourceFiles: []string{validSourcesPath},
			MaxWorkers:  4,
		},
		Multilog: map[string]interface{}{
			"level": "info",
		},
	}
	if err := validConfig.Validate(); err != nil {
		t.Errorf("Valid AppConfig should not return error: %v", err)
	}

	// Invalid Application config
	invalidAppConfig := AppConfig{
		Application: ApplicationConfig{
			// Missing required fields
		},
		DNSToolkit: DNSToolkitConfig{
			SourceFiles: []string{validSourcesPath},
			MaxWorkers:  4,
		},
		Multilog: map[string]interface{}{
			"level": "info",
		},
	}
	if err := invalidAppConfig.Validate(); err == nil {
		t.Error("AppConfig with invalid Application config should return error")
	}

	// Invalid DNSToolkit config
	invalidDNSConfig := AppConfig{
		Application: ApplicationConfig{
			Name:        "TestApp",
			Version:     "1.0.0",
			Description: "Test application",
		},
		DNSToolkit: DNSToolkitConfig{
			// Missing source files
			MaxWorkers: 4,
		},
		Multilog: map[string]interface{}{
			"level": "info",
		},
	}
	if err := invalidDNSConfig.Validate(); err == nil {
		t.Error("AppConfig with invalid DNSToolkit config should return error")
	}

	// Missing Multilog config
	missingMultilogConfig := AppConfig{
		Application: ApplicationConfig{
			Name:        "TestApp",
			Version:     "1.0.0",
			Description: "Test application",
		},
		DNSToolkit: DNSToolkitConfig{
			SourceFiles: []string{validSourcesPath},
			MaxWorkers:  4,
		},
		// Missing Multilog
	}
	if err := missingMultilogConfig.Validate(); err == nil {
		t.Error("AppConfig with missing Multilog config should return error")
	}
}

// Test LoadAppConfig
func TestBasicLoadAppConfig(t *testing.T) {
	testDir := createTestConfigFiles(t)
	defer os.RemoveAll(testDir)

	logger := CreateTestLogger()
	validConfigPath := filepath.Join(testDir, "valid_config.yaml")

	// Test loading valid config
	_, _, err := LoadAppConfig(logger, validConfigPath)
	if err != nil {
		t.Errorf("LoadAppConfig with valid file should not return error: %v", err)
	}

	// Test loading non-existent config
	_, _, err = LoadAppConfig(logger, "non_existent_config.yaml")
	if err == nil {
		t.Error("LoadAppConfig with non-existent file should return error")
	}

	// Create an invalid config file for testing
	invalidConfigPath := filepath.Join(testDir, "invalid_config.yaml")
	if err := os.WriteFile(invalidConfigPath, []byte("invalid yaml"), 0644); err != nil {
		t.Fatal(err)
	}

	// Test loading invalid config
	_, _, err = LoadAppConfig(logger, invalidConfigPath)
	if err == nil {
		t.Error("LoadAppConfig with invalid yaml should return error")
	}
}

func TestBasicSourceEnabled(t *testing.T) {
	enabledSource := Source{
		Name: "enabled_source",
	}
	if !enabledSource.IsEnabled() {
		t.Error("Source with Disabled=false should return true for IsEnabled()")
	}

	disabledSource := Source{
		Name:     "disabled_source",
		Disabled: true,
	}
	if disabledSource.IsEnabled() {
		t.Error("Source with Disabled=true should return false for IsEnabled()")
	}
}
