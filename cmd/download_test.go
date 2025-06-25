package cmd

import (
	"os"
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"

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

	// Create a test source config with a disabled source
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
