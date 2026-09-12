package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	d "github.com/phani-kb/dns-toolkit/internal/downloaders"
	"github.com/phani-kb/dns-toolkit/internal/utils"

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

func TestDownloadCommand_InitConfigurations(t *testing.T) {
	projectRoot, err := utils.FindProjectRoot("")
	assert.NoError(t, err)

	originalTestMode := os.Getenv("DNS_TOOLKIT_TEST_MODE")
	originalTestConfigPath := os.Getenv("DNS_TOOLKIT_TEST_CONFIG_PATH")
	defer func() {
		assert.NoError(t, os.Setenv("DNS_TOOLKIT_TEST_MODE", originalTestMode))
		assert.NoError(t, os.Setenv("DNS_TOOLKIT_TEST_CONFIG_PATH", originalTestConfigPath))
	}()
	assert.NoError(t, os.Setenv("DNS_TOOLKIT_TEST_MODE", constants.BooleanTrue))
	assert.NoError(t, os.Setenv("DNS_TOOLKIT_TEST_CONFIG_PATH", filepath.Join(projectRoot, "testdata", "config.yml")))

	InitForTesting()

	oldSources := SourcesConfigs
	SourcesConfigs = []config.SourcesConfig{}
	defer func() { SourcesConfigs = oldSources }()

	tests := []struct {
		name      string
		appConfig *config.AppConfig
	}{
		{
			name:      "nil app config uses defaults",
			appConfig: nil,
		},
		{
			name: "custom app config applies retry settings",
			appConfig: &config.AppConfig{
				DNSToolkit: config.DNSToolkitConfig{
					MaxRetries: 5,
					MaxWorkers: 2,
				},
			},
		},
	}

	oldConfig := AppConfig
	defer func() { AppConfig = oldConfig }()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AppConfig = tt.appConfig

			runFunc := downloadCmd.Run
			runFunc(downloadCmd, []string{})

			defaultDownloader, exists := d.GetDownloader("default")
			assert.True(t, exists, "Default downloader should be registered")
			assert.NotNil(t, defaultDownloader, "Default downloader should not be nil")
		})
	}
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

	domainTopDownloader, exists := d.GetDownloader("tranco")
	assert.True(t, exists, "Domain top downloader should be registered")
	assert.NotNil(t, domainTopDownloader, "Domain top downloader should not be nil")
	assert.Equal(t, "tranco", domainTopDownloader.Name(), "Domain top downloader should have correct name")
}
