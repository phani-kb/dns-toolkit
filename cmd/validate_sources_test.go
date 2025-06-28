package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateConfig(t *testing.T) {
	t.Parallel()

	os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	configPath := os.Getenv("DNS_TOOLKIT_TEST_CONFIG_PATH")
	if configPath == "" {
		t.Skip("DNS_TOOLKIT_TEST_CONFIG_PATH is not set, skipping test")
		return
	}

	defer func() {
		os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
	}()

	tests := []struct {
		name        string
		configPath  string
		expectError bool
		setupFunc   func()
		cleanupFunc func()
	}{
		{
			name:        "valid config path",
			configPath:  configPath,
			expectError: false,
			setupFunc: func() {
				validationPerformed = false
			},
			cleanupFunc: func() {
				validationPerformed = false
			},
		},
		{
			name:        "already validated - should skip",
			configPath:  configPath,
			expectError: false,
			setupFunc: func() {
				validationPerformed = true
			},
			cleanupFunc: func() {
				validationPerformed = false
			},
		},
		{
			name:        "invalid config path",
			configPath:  "/nonexistent/config.yml",
			expectError: true,
			setupFunc: func() {
				validationPerformed = false
			},
			cleanupFunc: func() {
				validationPerformed = false
			},
		},
		{
			name:        "empty config path",
			configPath:  "",
			expectError: true,
			setupFunc: func() {
				validationPerformed = false
			},
			cleanupFunc: func() {
				validationPerformed = false
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupFunc != nil {
				tt.setupFunc()
			}

			defer func() {
				if tt.cleanupFunc != nil {
					tt.cleanupFunc()
				}
			}()

			err := validateConfig(tt.configPath)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "config validation error")
			} else {
				if !validationPerformed && tt.name != "already validated - should skip" {
					// Should have attempted validation
				}
			}
		})
	}
}

func TestValidateSourcesCommand(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, validateSourcesCmd)
	assert.Equal(t, "validate-sources", validateSourcesCmd.Use)
	assert.Contains(t, validateSourcesCmd.Short, "Validate")
	assert.NotNil(t, validateSourcesCmd.Run)
}
