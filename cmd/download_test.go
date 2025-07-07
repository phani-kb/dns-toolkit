package cmd

import (
	"os"
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	d "github.com/phani-kb/dns-toolkit/internal/downloaders"

	"github.com/stretchr/testify/assert"
)

// TestDownloadCommand tests the download command's Run function
func TestDownloadCommand(t *testing.T) {
	InitForTesting()

	oldSources := SourcesConfigs
	SourcesConfigs = []config.SourcesConfig{}
	defer func() { SourcesConfigs = oldSources }()

	cmd := downloadCmd

	runFunc := downloadCmd.Run
	runFunc(cmd, []string{})

	_, err := os.Stat(constants.DownloadDir)
	assert.NoError(t, err, "Download directory should exist")

	_, err = os.Stat(constants.SummaryDir)
	assert.NoError(t, err, "Summary directory should exist")
}

// TestDownloadCommand_WithSources tests the download command with sources
func TestDownloadCommand_WithSources(t *testing.T) {
	InitForTesting()

	testSource := config.Source{
		Name:     "test-source",
		URL:      "http://example.com/list.txt",
		Disabled: true, // Disabled to avoid actual downloads
		Types:    []common.SourceType{{Name: "domain"}},
	}

	sourceConfig := config.SourcesConfig{
		Sources: []config.Source{testSource},
	}

	oldSources := SourcesConfigs
	SourcesConfigs = []config.SourcesConfig{sourceConfig}
	defer func() { SourcesConfigs = oldSources }()

	runFunc := downloadCmd.Run
	runFunc(downloadCmd, []string{})
}

// TestValidateAndInitDownloader tests downloader initialization logic
func TestValidateAndInitDownloader(t *testing.T) {
	InitForTesting()

	oldConfig := AppConfig
	AppConfig = nil

	runFunc := downloadCmd.Run
	runFunc(downloadCmd, []string{})

	AppConfig = oldConfig
}

// TestDownloadCommand_DownloaderRegistration tests that both default and domain top downloaders are registered
func TestDownloadCommand_DownloaderRegistration(t *testing.T) {
	InitForTesting()

	oldSources := SourcesConfigs
	SourcesConfigs = []config.SourcesConfig{}
	defer func() { SourcesConfigs = oldSources }()

	runFunc := downloadCmd.Run
	runFunc(downloadCmd, []string{})

	defaultDownloader, exists := d.GetDownloader("default")
	assert.True(t, exists, "Default downloader should be registered")
	assert.NotNil(t, defaultDownloader, "Default downloader should not be nil")

	domainTopDownloader, exists := d.GetDownloader("transco")
	assert.True(t, exists, "Domain top downloader should be registered")
	assert.NotNil(t, domainTopDownloader, "Domain top downloader should not be nil")
	assert.Equal(t, "transco", domainTopDownloader.Name(), "Domain top downloader should have correct name")
}

// TestDownloadCommand_WithRetryConfiguration tests downloader initialization with retry settings
func TestDownloadCommand_WithRetryConfiguration(t *testing.T) {
	InitForTesting()

	oldConfig := AppConfig
	AppConfig = &config.AppConfig{
		DNSToolkit: config.DNSToolkitConfig{
			MaxRetries: 5,
			MaxWorkers: 2,
		},
	}
	defer func() { AppConfig = oldConfig }()

	oldSources := SourcesConfigs
	SourcesConfigs = []config.SourcesConfig{}
	defer func() { SourcesConfigs = oldSources }()

	// Run the download command
	runFunc := downloadCmd.Run
	runFunc(downloadCmd, []string{})

	// Verify directories are created
	_, err := os.Stat(constants.DownloadDir)
	assert.NoError(t, err, "Download directory should exist")

	_, err = os.Stat(constants.SummaryDir)
	assert.NoError(t, err, "Summary directory should exist")
}
