package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/stretchr/testify/assert"
)

// TestConstantsSummaryDirSetProperly tests that the constants.SummaryDir is set correctly
func TestConstantsSummaryDirSetProperly(t *testing.T) {
	oldConfig := AppConfig
	defer func() { AppConfig = oldConfig }()

	// Set up test config
	AppConfig = &config.AppConfig{
		DNSToolkit: config.DNSToolkitConfig{
			Folders: config.FoldersConfig{
				Summary: "testdata/summary",
			},
		},
	}

	// Set test mode
	err := os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	assert.NoError(t, err)
	defer func() {
		err := os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
		if err != nil {
			t.Logf("Failed to unset DNS_TOOLKIT_TEST_MODE: %v", err)
		}
	}()

	InitForTesting()

	assert.True(t, strings.HasSuffix(constants.SummaryDir, "testdata/summary"),
		"constants.SummaryDir should end with testdata/summary from test config, got: %s", constants.SummaryDir)
	assert.NotEqual(t, "data", constants.SummaryDir,
		"constants.SummaryDir should not be the default 'data'")
	assert.True(t, filepath.IsAbs(constants.SummaryDir),
		"constants.SummaryDir should be an absolute path in test mode")
}
