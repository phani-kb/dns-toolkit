package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/db"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateOutputCmdRun(t *testing.T) {
	err := os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	require.NoError(t, err)
	defer func() {
		_ = os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
	}()

	generateOutputCmd.Run(generateOutputCmd, []string{})
}

func TestPrepareDirectories(t *testing.T) {
	tempDir := t.TempDir()

	origOutputDir := constants.OutputDir
	origOutputIgnoredDir := constants.OutputIgnoredDir
	origOutputSummariesDir := constants.OutputSummariesDir
	origArchiveDir := constants.ArchiveDir
	origIncludeIgnored := includeIgnored
	defer func() {
		constants.OutputDir = origOutputDir
		constants.OutputIgnoredDir = origOutputIgnoredDir
		constants.OutputSummariesDir = origOutputSummariesDir
		constants.ArchiveDir = origArchiveDir
		includeIgnored = origIncludeIgnored
	}()

	constants.OutputDir = filepath.Join(tempDir, "output")
	constants.OutputIgnoredDir = filepath.Join(tempDir, "ignored")
	constants.OutputSummariesDir = filepath.Join(tempDir, "summaries")
	constants.ArchiveDir = filepath.Join(tempDir, "archive")

	includeIgnored = false
	err := prepareDirectories()
	assert.NoError(t, err)

	assert.DirExists(t, constants.OutputDir)
	assert.DirExists(t, constants.OutputSummariesDir)
	assert.DirExists(t, constants.ArchiveDir)
	assert.NoDirExists(t, constants.OutputIgnoredDir)

	includeIgnored = true
	err = prepareDirectories()
	assert.NoError(t, err)
	assert.DirExists(t, constants.OutputIgnoredDir)
}

func TestLoadTemplates(t *testing.T) {
	t.Parallel()

	err := os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	assert.NoError(t, err)
	defer func() {
		_ = os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
	}()

	tmpl, staticTemplate, err := loadTemplates()
	if err != nil {
		assert.Contains(t, err.Error(), "template")
	} else {
		assert.NotNil(t, tmpl)
		assert.NotNil(t, staticTemplate)
		assert.Greater(t, len(staticTemplate), 0)
	}
}

func TestBuildOutputFileName(t *testing.T) {
	tests := []struct {
		name              string
		consolidationType string
		group             db.ConsolidatedGroup
		expected          string
	}{
		{
			name:              "general blocklist",
			consolidationType: "general",
			group: db.ConsolidatedGroup{
				GenericSourceType: "domain",
				ListType:          "blocklist",
			},
			expected: "domain_blocklist.txt",
		},
		{
			name:              "group with name",
			consolidationType: "group",
			group: db.ConsolidatedGroup{
				GenericSourceType: "domain",
				ListType:          "blocklist",
				GroupName:         "mini",
			},
			expected: "mini_domain_blocklist.txt",
		},
		{
			name:              "category with name",
			consolidationType: "category",
			group: db.ConsolidatedGroup{
				GenericSourceType: "adguard",
				ListType:          "allowlist",
				Category:          "ads",
			},
			expected: "ads_adguard_allowlist.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildOutputFileName(tt.consolidationType, tt.group)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestBuildOutputDescription(t *testing.T) {
	tests := []struct {
		name              string
		consolidationType string
		group             db.ConsolidatedGroup
		expected          string
	}{
		{
			name:              "general",
			consolidationType: "general",
			group: db.ConsolidatedGroup{
				GenericSourceType: "domain",
				ListType:          "blocklist",
			},
			expected: "domain blocklist",
		},
		{
			name:              "group",
			consolidationType: "group",
			group: db.ConsolidatedGroup{
				GenericSourceType: "ipv4",
				ListType:          "blocklist",
				GroupName:         "big",
			},
			expected: "ipv4 blocklist [big]",
		},
		{
			name:              "category",
			consolidationType: "category",
			group: db.ConsolidatedGroup{
				GenericSourceType: "domain",
				ListType:          "allowlist",
				Category:          "malware",
			},
			expected: "domain allowlist [malware]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildOutputDescription(tt.consolidationType, tt.group)
			assert.Equal(t, tt.expected, result)
		})
	}
}
