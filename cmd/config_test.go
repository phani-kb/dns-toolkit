package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfigPath(t *testing.T) {
	err := os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
	assert.NoError(t, err)
	err = os.Unsetenv("DNS_TOOLKIT_TEST_CONFIG_PATH")
	assert.NoError(t, err)

	configPath, err := GetConfigPath()
	assert.NoError(t, err)
	assert.Equal(t, "configs/config.yml", configPath)

	err = os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	assert.NoError(t, err)
	err = os.Setenv("DNS_TOOLKIT_TEST_CONFIG_PATH", "/test/config.yml")
	assert.NoError(t, err)

	configPath, err = GetConfigPath()
	assert.NoError(t, err)
	assert.Equal(t, "/test/config.yml", configPath)

	err = os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	assert.NoError(t, err)
	err = os.Unsetenv("DNS_TOOLKIT_TEST_CONFIG_PATH")
	assert.NoError(t, err)

	_, err = GetConfigPath()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "DNS_TOOLKIT_TEST_CONFIG_PATH is not set")

	err = os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
	assert.NoError(t, err)
	err = os.Unsetenv("DNS_TOOLKIT_TEST_CONFIG_PATH")
	assert.NoError(t, err)
}

func TestMinOverlapPercent(t *testing.T) {
	assert.Equal(t, 0.1, AppConfig.DNSToolkit.GetMinOverlapPercent())

	AppConfig.DNSToolkit.MinOverlapPercent = 0.2
	assert.Equal(t, 0.2, AppConfig.DNSToolkit.GetMinOverlapPercent())

	AppConfig.DNSToolkit.MinOverlapPercent = 0
	assert.Equal(t, 0.1, AppConfig.DNSToolkit.GetMinOverlapPercent())
}
