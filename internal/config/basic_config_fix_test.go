package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBasicLoadAppConfigWithProperLogger(t *testing.T) {
	testDir := createTestConfigFiles(t)
	defer os.RemoveAll(testDir)

	logger := CreateTestLogger()

	validConfigPath := filepath.Join(testDir, "valid_config.yaml")

	// Test loading valid config
	appConfig, sourcesConfig, err := LoadAppConfig(logger, validConfigPath)
	if err != nil {
		t.Errorf("LoadAppConfig with valid file should not return error: %v", err)
		return
	}

	// Validate the config was loaded correctly
	if appConfig.Application.Name != "TestApp" {
		t.Errorf("Expected application name to be 'TestApp', got '%s'", appConfig.Application.Name)
	}

	if len(sourcesConfig) == 0 {
		t.Error("LoadAppConfig should return non-empty SourcesConfig slice")
		return
	}

	if len(sourcesConfig[0].Sources) != 1 {
		t.Errorf("Expected 1 source, got %d", len(sourcesConfig[0].Sources))
		return
	}

	if sourcesConfig[0].Sources[0].Name != "test_source" {
		t.Errorf("Expected source name to be 'test_source', got '%s'", sourcesConfig[0].Sources[0].Name)
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
