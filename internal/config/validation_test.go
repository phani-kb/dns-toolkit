package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateAppConfig_Validation(t *testing.T) {
	t.Run("Valid app config validates successfully", func(t *testing.T) {
		config := AppConfig{
			Application: ApplicationConfig{
				Name:        "dns-toolkit",
				Version:     "1.0.0",
				Description: "DNS Toolkit Application",
			},
			DNSToolkit: DNSToolkitConfig{
				SourceFiles: []string{
					"/tmp/sources1.json",
					"/tmp/sources2.json",
				},
				Folders: FoldersConfig{
					Download:     "/tmp/download",
					Processed:    "/tmp/processed",
					Consolidated: "/tmp/consolidated",
					Summary:      "/tmp/summary",
					Overlap:      "/tmp/overlap",
					Top:          "/tmp/top",
				},
			},
			Multilog: map[string]interface{}{
				"log_level": "info",
				"log_file":  "test.log",
			},
		}

		// Create temporary files to satisfy validation
		createTempFiles(t, config.DNSToolkit.SourceFiles)
		defer removeTempFiles(config.DNSToolkit.SourceFiles)

		err := ValidateAppConfig(config)
		assert.NoError(t, err)
	})

	t.Run("Invalid app config returns validation error", func(t *testing.T) {
		config := AppConfig{
			Application: ApplicationConfig{
				Name:    "", // Missing required name
				Version: "1.0.0",
			},
			DNSToolkit: DNSToolkitConfig{
				SourceFiles: []string{"/non-existent-file.json"},
			},
		}

		err := ValidateAppConfig(config)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "application validation failed")
	})
}

// Helper functions to create and remove temporary files for testing
func createTempFiles(t *testing.T, files []string) {
	for _, file := range files {
		f, err := os.Create(file)
		assert.NoError(t, err)
		f.Close()
	}
}

func removeTempFiles(files []string) {
	for _, file := range files {
		os.Remove(file)
	}
}
