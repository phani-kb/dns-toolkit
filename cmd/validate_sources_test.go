package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateConfig(t *testing.T) {
	t.Parallel()

	err := os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	assert.NoError(t, err)
	configPath := os.Getenv("DNS_TOOLKIT_TEST_CONFIG_PATH")
	if configPath == "" {
		t.Skip("DNS_TOOLKIT_TEST_CONFIG_PATH is not set, skipping test")
		return
	}

	defer func() {
		err := os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
		if err != nil {
			t.Logf("Failed to unset DNS_TOOLKIT_TEST_MODE: %v", err)
		}
	}()

	tests := []struct {
		setupFunc   func()
		cleanupFunc func()
		name        string
		configPath  string
		expectError bool
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

			// Capture initial state
			initialValidationState := validationPerformed

			err := validateConfig(tt.configPath)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "config validation error")
			} else {
				assert.NoError(t, err)
				if tt.name == "already validated - should skip" {
					assert.True(t, initialValidationState, "Initial state should be true for skip case")
					assert.True(t, validationPerformed, "Validation should remain true")
				} else {
					assert.False(t, initialValidationState, "Initial state should be false")
					assert.True(t, validationPerformed, "Validation should be performed and become true")
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
