package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateStatsCommand(t *testing.T) {
	err := os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	require.NoError(t, err)
	defer func() {
		err := os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
		require.NoError(t, err)
	}()

	assert.Equal(t, "stats-readme", generateStatsCmd.Use)
	assert.Equal(t, "Generate and update source statistics in README.md", generateStatsCmd.Short)
	assert.NotNil(t, generateStatsCmd.Run)
}

func TestCollectSourceStats(t *testing.T) {
	err := os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	require.NoError(t, err)
	defer func() {
		err := os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
		require.NoError(t, err)
	}()

	originalSourcesConfigs := SourcesConfigs
	defer func() {
		SourcesConfigs = originalSourcesConfigs
	}()

	mockSources := []config.SourcesConfig{
		{
			Sources: []config.Source{
				{
					Name:       "test-source-1",
					Disabled:   false,
					Categories: []string{"ads", "malware"},
					Countries:  []string{"US", "GB"},
					Types: []common.SourceType{
						{
							Name: "domain",
							ListTypes: []common.ListType{
								{Name: "blocklist"},
								{Name: "allowlist"},
							},
						},
					},
				},
				{
					Name:       "test-source-2",
					Disabled:   true,
					Categories: []string{"privacy", "ads"},
					Countries:  []string{"DE", "US"},
					Types: []common.SourceType{
						{
							Name: "ipv4",
							ListTypes: []common.ListType{
								{Name: "blocklist"},
							},
						},
					},
				},
				{
					Name:       "test-source-3",
					Disabled:   false,
					Categories: []string{"malware"},
					Countries:  []string{"FR"},
					Types: []common.SourceType{
						{
							Name: "domain_adguard",
							ListTypes: []common.ListType{
								{Name: "allowlist"},
							},
						},
					},
				},
			},
		},
	}

	SourcesConfigs = mockSources

	stats := collectSourceStats()

	assert.Equal(t, 3, stats.TotalSources)
	assert.Equal(t, 2, stats.EnabledSources)
	assert.Equal(t, 1, stats.DisabledSources)
	assert.Equal(t, 2, stats.BlocklistSources)
	assert.Equal(t, 2, stats.AllowlistSources)

	expectedCategories := []string{"ads", "malware", "privacy"}
	assert.Equal(t, expectedCategories, stats.Categories)

	expectedSourceTypes := []string{"domain", "domain_adguard", "ipv4"}
	assert.Equal(t, expectedSourceTypes, stats.SourceTypes)

	expectedCountries := []string{"DE", "FR", "GB", "US"}
	assert.Equal(t, expectedCountries, stats.Countries)

	assert.NotEmpty(t, stats.LastUpdated)
	_, parseErr := time.Parse("2006-01-02 15:04:05 UTC", stats.LastUpdated)
	assert.NoError(t, parseErr)
}

func TestGenerateStatsSection(t *testing.T) {
	stats := &SourceStats{
		LastUpdated:      "2025-01-01 12:00:00 UTC",
		Categories:       []string{"ads", "malware"},
		SourceTypes:      []string{"domain", "ipv4"},
		Countries:        []string{"US", "DE"},
		TotalSources:     10,
		EnabledSources:   8,
		DisabledSources:  2,
		BlocklistSources: 9,
		AllowlistSources: 3,
	}

	section := generateStatsSection(stats)

	assert.Contains(t, section, "<!-- STATS_START -->")
	assert.Contains(t, section, "<!-- STATS_END -->")
	assert.Contains(t, section, "## Source Statistics")
	assert.Contains(t, section, "| Metric | Count | Details |")

	assert.Contains(t, section, "| **Total Sources** | 10 | 8 enabled, 2 disabled |")
	assert.Contains(t, section, "| **Blocklist Sources** | 9 | Sources providing blocking rules |")
	assert.Contains(t, section, "| **Allowlist Sources** | 3 | Sources providing exception rules |")
	assert.Contains(t, section, "| **Categories** | 2 | ads, malware |")
	assert.Contains(t, section, "| **Source Types** | 2 | domain, ipv4 |")
	assert.Contains(t, section, "| **Geographic Coverage** | 2 countries | US, DE |")
	assert.Contains(t, section, "| **Last Updated** | 2025-01-01 12:00:00 UTC | Statistics generation time |")

	assert.True(t, strings.HasSuffix(section, "<!-- STATS_END -->"))
}

func TestUpdateReadmeWithStats(t *testing.T) {
	t.Run("Replace existing stats section", func(t *testing.T) {
		tempFile := filepath.Join(t.TempDir(), "README.md")
		testFile := filepath.Join("..", "testdata", "README_with_stats.md")

		content, err := os.ReadFile(testFile)
		require.NoError(t, err)
		err = os.WriteFile(tempFile, content, 0644)
		require.NoError(t, err)

		stats := &SourceStats{
			LastUpdated:      "2025-01-01 12:00:00 UTC",
			Categories:       []string{"ads"},
			SourceTypes:      []string{"domain"},
			Countries:        []string{"US"},
			TotalSources:     5,
			EnabledSources:   4,
			DisabledSources:  1,
			BlocklistSources: 4,
			AllowlistSources: 1,
		}

		err = updateReadmeWithStats(stats, tempFile)
		require.NoError(t, err)

		updatedContent, err := os.ReadFile(tempFile)
		require.NoError(t, err)
		contentStr := string(updatedContent)

		assert.NotContains(t, contentStr, "Source Statistics (Old)")
		assert.NotContains(t, contentStr, "| **Total Sources** | 99 |")

		assert.Contains(t, contentStr, "## Source Statistics")
		assert.Contains(t, contentStr, "| **Total Sources** | 5 | 4 enabled, 1 disabled |")
		assert.Contains(t, contentStr, "2025-01-01 12:00:00 UTC")

		assert.Contains(t, contentStr, "# DNS Toolkit Test")
		assert.Contains(t, contentStr, "## Published Outputs")
	})

	t.Run("Add new stats section", func(t *testing.T) {
		tempFile := filepath.Join(t.TempDir(), "README.md")
		testFile := filepath.Join("..", "testdata", "README.md")

		content, err := os.ReadFile(testFile)
		require.NoError(t, err)
		err = os.WriteFile(tempFile, content, 0644)
		require.NoError(t, err)

		stats := &SourceStats{
			LastUpdated:      "2025-01-01 12:00:00 UTC",
			Categories:       []string{"ads"},
			SourceTypes:      []string{"domain"},
			Countries:        []string{"US"},
			TotalSources:     3,
			EnabledSources:   2,
			DisabledSources:  1,
			BlocklistSources: 2,
			AllowlistSources: 1,
		}

		err = updateReadmeWithStats(stats, tempFile)
		require.NoError(t, err)

		updatedContent, err := os.ReadFile(tempFile)
		require.NoError(t, err)
		contentStr := string(updatedContent)

		assert.Contains(t, contentStr, "## Source Statistics")
		assert.Contains(t, contentStr, "| **Total Sources** | 3 | 2 enabled, 1 disabled |")

		statsIndex := strings.Index(contentStr, "## Source Statistics")
		outputsIndex := strings.Index(contentStr, "## Published Outputs")
		assert.Greater(t, outputsIndex, statsIndex)
	})

	t.Run("Error when no Published Outputs section", func(t *testing.T) {
		tempFile := filepath.Join(t.TempDir(), "README.md")
		testContent := `# Test Project

Some description here.
`
		err := os.WriteFile(tempFile, []byte(testContent), 0644)
		require.NoError(t, err)

		stats := &SourceStats{
			LastUpdated:      "2025-01-01 12:00:00 UTC",
			Categories:       []string{"ads"},
			SourceTypes:      []string{"domain"},
			Countries:        []string{"US"},
			TotalSources:     1,
			EnabledSources:   1,
			DisabledSources:  0,
			BlocklistSources: 1,
			AllowlistSources: 0,
		}

		err = updateReadmeWithStats(stats, tempFile)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "could not find '## Published Outputs' section")
	})

	t.Run("Error with malformed markers - only STATS_START", func(t *testing.T) {
		tempFile := filepath.Join(t.TempDir(), "README.md")
		testContent := `# Test Project

Some description here.

<!-- STATS_START -->
## Old Statistics
Old content here

## Published Outputs
Output information here.
`
		err := os.WriteFile(tempFile, []byte(testContent), 0644)
		require.NoError(t, err)

		stats := &SourceStats{
			LastUpdated:      "2025-01-01 12:00:00 UTC",
			Categories:       []string{"ads"},
			SourceTypes:      []string{"domain"},
			Countries:        []string{"US"},
			TotalSources:     1,
			EnabledSources:   1,
			DisabledSources:  0,
			BlocklistSources: 1,
			AllowlistSources: 0,
		}

		err = updateReadmeWithStats(stats, tempFile)
		require.NoError(t, err)

		updatedContent, err := os.ReadFile(tempFile)
		require.NoError(t, err)
		contentStr := string(updatedContent)

		assert.Contains(t, contentStr, "## Source Statistics")
		assert.Contains(t, contentStr, "| **Total Sources** | 1 | 1 enabled, 0 disabled |")

		assert.Contains(t, contentStr, "<!-- STATS_START -->")
		assert.Contains(t, contentStr, "<!-- STATS_END -->")
	})

	t.Run("Error with malformed markers - only STATS_END", func(t *testing.T) {
		tempFile := filepath.Join(t.TempDir(), "README.md")
		testContent := `# Test Project

Some description here.

## Old Statistics
Old content here
<!-- STATS_END -->

## Published Outputs
Output information here.
`
		err := os.WriteFile(tempFile, []byte(testContent), 0644)
		require.NoError(t, err)

		stats := &SourceStats{
			LastUpdated:      "2025-01-01 12:00:00 UTC",
			Categories:       []string{"ads"},
			SourceTypes:      []string{"domain"},
			Countries:        []string{"US"},
			TotalSources:     1,
			EnabledSources:   1,
			DisabledSources:  0,
			BlocklistSources: 1,
			AllowlistSources: 0,
		}

		err = updateReadmeWithStats(stats, tempFile)
		require.NoError(t, err)

		updatedContent, err := os.ReadFile(tempFile)
		require.NoError(t, err)
		contentStr := string(updatedContent)

		assert.Contains(t, contentStr, "## Source Statistics")
		assert.Contains(t, contentStr, "| **Total Sources** | 1 | 1 enabled, 0 disabled |")

		assert.Contains(t, contentStr, "<!-- STATS_START -->")
		assert.Contains(t, contentStr, "<!-- STATS_END -->")
	})
}

func TestGenerateStatsCommandRun(t *testing.T) {
	err := os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	require.NoError(t, err)
	defer func() {
		err := os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
		require.NoError(t, err)
	}()

	originalSourcesConfigs := SourcesConfigs
	defer func() {
		SourcesConfigs = originalSourcesConfigs
	}()

	mockSources := []config.SourcesConfig{
		{
			Sources: []config.Source{
				{
					Name:       "test-source",
					Disabled:   false,
					Categories: []string{"ads"},
					Countries:  []string{"US"},
					Types: []common.SourceType{
						{
							Name: "domain",
							ListTypes: []common.ListType{
								{Name: "blocklist"},
							},
						},
					},
				},
			},
		},
	}

	SourcesConfigs = mockSources

	tempFile := filepath.Join(t.TempDir(), "README.md")
	testFile := filepath.Join("..", "testdata", "README.md")

	content, err := os.ReadFile(testFile)
	require.NoError(t, err)
	err = os.WriteFile(tempFile, content, 0644)
	require.NoError(t, err)

	originalDir, err := os.Getwd()
	require.NoError(t, err)
	defer func() {
		err := os.Chdir(originalDir)
		require.NoError(t, err)
	}()

	err = os.Chdir(filepath.Dir(tempFile))
	require.NoError(t, err)

	err = os.Rename(filepath.Base(tempFile), "README.md")
	require.NoError(t, err)

	err = os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
	require.NoError(t, err)
	defer func() {
		err := os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
		require.NoError(t, err)
	}()

	cmd := &cobra.Command{}
	runFunc := generateStatsCmd.Run

	assert.NotPanics(t, func() {
		runFunc(cmd, []string{})
	})

	updatedContent, err := os.ReadFile("README.md")
	require.NoError(t, err)
	contentStr := string(updatedContent)

	assert.Contains(t, contentStr, "## Source Statistics")
	assert.Contains(t, contentStr, "| **Total Sources** | 1 | 1 enabled, 0 disabled |")
}

func TestCollectSourceStatsWithEmptyData(t *testing.T) {
	err := os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	require.NoError(t, err)
	defer func() {
		err := os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
		require.NoError(t, err)
	}()

	originalSourcesConfigs := SourcesConfigs
	defer func() {
		SourcesConfigs = originalSourcesConfigs
	}()

	SourcesConfigs = []config.SourcesConfig{}

	stats := collectSourceStats()

	assert.Equal(t, 0, stats.TotalSources)
	assert.Equal(t, 0, stats.EnabledSources)
	assert.Equal(t, 0, stats.DisabledSources)
	assert.Equal(t, 0, stats.BlocklistSources)
	assert.Equal(t, 0, stats.AllowlistSources)
	assert.Empty(t, stats.Categories)
	assert.Empty(t, stats.SourceTypes)
	assert.Empty(t, stats.Countries)
	assert.NotEmpty(t, stats.LastUpdated)
}

func TestCollectSourceStatsWithComplexData(t *testing.T) {
	err := os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	require.NoError(t, err)
	defer func() {
		err := os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
		require.NoError(t, err)
	}()

	originalSourcesConfigs := SourcesConfigs
	defer func() {
		SourcesConfigs = originalSourcesConfigs
	}()

	mockSources := []config.SourcesConfig{
		{
			Sources: []config.Source{
				{
					Name:       "source-1",
					Disabled:   false,
					Categories: []string{"ads", "malware", "privacy"},
					Countries:  []string{"US", "GB", "DE"},
					Types: []common.SourceType{
						{
							Name: "domain",
							ListTypes: []common.ListType{
								{Name: "blocklist"},
							},
						},
						{
							Name: "ipv4",
							ListTypes: []common.ListType{
								{Name: "allowlist"},
							},
						},
					},
				},
				{
					Name:       "source-2",
					Disabled:   true,
					Categories: []string{"", "  ", "ads"},
					Countries:  []string{"", "FR", "  "},
					Types: []common.SourceType{
						{
							Name: "domain_adguard",
							ListTypes: []common.ListType{
								{Name: "blocklist"},
								{Name: "allowlist"},
							},
						},
					},
				},
			},
		},
	}

	SourcesConfigs = mockSources

	stats := collectSourceStats()

	assert.Equal(t, 2, stats.TotalSources)
	assert.Equal(t, 1, stats.EnabledSources)
	assert.Equal(t, 1, stats.DisabledSources)
	assert.Equal(t, 2, stats.BlocklistSources)
	assert.Equal(t, 2, stats.AllowlistSources)

	expectedCategories := []string{"ads", "malware", "privacy"}
	assert.Equal(t, expectedCategories, stats.Categories)

	expectedCountries := []string{"DE", "FR", "GB", "US"}
	assert.Equal(t, expectedCountries, stats.Countries)

	expectedSourceTypes := []string{"domain", "domain_adguard", "ipv4"}
	assert.Equal(t, expectedSourceTypes, stats.SourceTypes)
}
