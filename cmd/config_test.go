package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfigPath(t *testing.T) {
	os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
	os.Unsetenv("DNS_TOOLKIT_TEST_CONFIG_PATH")

	configPath, err := GetConfigPath()
	assert.NoError(t, err)
	assert.Equal(t, "configs/config.yml", configPath)

	os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	os.Setenv("DNS_TOOLKIT_TEST_CONFIG_PATH", "/test/config.yml")

	configPath, err = GetConfigPath()
	assert.NoError(t, err)
	assert.Equal(t, "/test/config.yml", configPath)

	os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	os.Unsetenv("DNS_TOOLKIT_TEST_CONFIG_PATH")

	_, err = GetConfigPath()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "DNS_TOOLKIT_TEST_CONFIG_PATH is not set")

	os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
	os.Unsetenv("DNS_TOOLKIT_TEST_CONFIG_PATH")
}
