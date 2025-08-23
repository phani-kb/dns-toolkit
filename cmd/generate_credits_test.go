package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateCreditsCommand(t *testing.T) {
	err := os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	require.NoError(t, err)
	defer func() {
		err := os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
		require.NoError(t, err)
	}()

	assert.Equal(t, "credits", generateCreditsCmd.Use)
	assert.Equal(t, "Generate and update source credits in README.md", generateCreditsCmd.Short)
	assert.NotNil(t, generateCreditsCmd.Run)

	cmd := &cobra.Command{}
	runFunc := generateCreditsCmd.Run
	assert.NotPanics(t, func() {
		runFunc(cmd, []string{})
	})
}

func TestGenerateCreditsSection(t *testing.T) {
	originalAppConfig := AppConfig
	originalSourcesConfigs := SourcesConfigs
	defer func() {
		AppConfig = originalAppConfig
		SourcesConfigs = originalSourcesConfigs
	}()

	t.Run("Basic credits generation", func(t *testing.T) {
		mockAppConfig := &config.AppConfig{
			DNSToolkit: config.DNSToolkitConfig{
				SourceFiles: []string{"test-sources.yml"},
			},
		}
		AppConfig = mockAppConfig

		mockSourcesConfigs := []config.SourcesConfig{
			{
				Sources: []config.Source{
					{
						Name:       "test-source",
						URL:        "https://example.com/list",
						Disabled:   false,
						Categories: []string{"ads"},
						Notes:      "Test source",
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
		SourcesConfigs = mockSourcesConfigs

		section := generateCreditsSection()

		assert.Contains(t, section, "<!-- CREDITS_START -->")
		assert.Contains(t, section, "<!-- CREDITS_END -->")
		assert.Contains(t, section, "## Source Credits")
		assert.Contains(t, section, "<details>")
		assert.Contains(t, section, "test-sources.yml")
		assert.Contains(t, section, "[test-source](https://example.com/list)")
		assert.Contains(t, section, "✅ Enabled")
		assert.Contains(t, section, "ads")
		assert.Contains(t, section, "Test source")
		assert.True(t, strings.HasSuffix(section, "<!-- CREDITS_END -->\n"))
	})

	t.Run("Special characters and edge cases", func(t *testing.T) {
		mockAppConfig := &config.AppConfig{
			DNSToolkit: config.DNSToolkitConfig{
				SourceFiles: []string{"test-sources.yml"},
			},
		}
		AppConfig = mockAppConfig

		mockSourcesConfigs := []config.SourcesConfig{
			{
				Sources: []config.Source{
					{
						Name:       "special-source",
						URL:        "file:///local/path",
						Disabled:   true,
						Categories: []string{},
						Notes:      "Notes with | pipes | and\nnewlines",
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
		SourcesConfigs = mockSourcesConfigs

		section := generateCreditsSection()

		assert.Contains(t, section, "special-source")
		assert.NotContains(t, section, "[special-source](file://")
		assert.Contains(t, section, "❌ Disabled")
		assert.Contains(t, section, "| - |")
		assert.Contains(t, section, "Notes with \\| pipes \\|")
		assert.Contains(t, section, "and newlines")
	})

	t.Run("No sources configured", func(t *testing.T) {
		AppConfig = nil
		SourcesConfigs = []config.SourcesConfig{}

		section := generateCreditsSection()

		assert.Contains(t, section, "<!-- CREDITS_START -->")
		assert.Contains(t, section, "<!-- CREDITS_END -->")
		assert.Contains(t, section, "## Source Credits")
		assert.NotContains(t, section, "<details>")
	})
}

func TestUpdateReadmeWithCredits(t *testing.T) {
	originalAppConfig := AppConfig
	originalSourcesConfigs := SourcesConfigs
	defer func() {
		AppConfig = originalAppConfig
		SourcesConfigs = originalSourcesConfigs
	}()

	mockAppConfig := &config.AppConfig{
		DNSToolkit: config.DNSToolkitConfig{
			SourceFiles: []string{"test-sources.yml"},
		},
	}
	AppConfig = mockAppConfig

	mockSourcesConfigs := []config.SourcesConfig{
		{
			Sources: []config.Source{
				{
					Name:       "test-source",
					URL:        "https://example.com/list",
					Disabled:   false,
					Categories: []string{"ads"},
					Notes:      "Test source",
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
	SourcesConfigs = mockSourcesConfigs

	t.Run("Replace existing credits section", func(t *testing.T) {
		tempFile := filepath.Join(t.TempDir(), "README.md")
		testContent := `# Test Project

<!-- CREDITS_START -->
## Old Credits
Old content here
<!-- CREDITS_END -->

## Installation
Installation information here.
`
		err := os.WriteFile(tempFile, []byte(testContent), 0644)
		require.NoError(t, err)

		err = updateReadmeWithCredits(Logger, tempFile)
		require.NoError(t, err)

		updatedContent, err := os.ReadFile(tempFile)
		require.NoError(t, err)
		contentStr := string(updatedContent)

		assert.NotContains(t, contentStr, "Old Credits")
		assert.NotContains(t, contentStr, "Old content here")

		assert.Contains(t, contentStr, "## Source Credits")
		assert.Contains(t, contentStr, "test-sources.yml")
		assert.Contains(t, contentStr, "[test-source](https://example.com/list)")

		assert.Contains(t, contentStr, "# Test Project")
		assert.Contains(t, contentStr, "## Installation")
	})

	t.Run("Replace existing credits section with testdata file", func(t *testing.T) {
		tempFile := filepath.Join(t.TempDir(), "README.md")
		testFile := filepath.Join("..", "testdata", "README_with_credits.md")

		content, err := os.ReadFile(testFile)
		require.NoError(t, err)
		err = os.WriteFile(tempFile, content, 0644)
		require.NoError(t, err)

		err = updateReadmeWithCredits(Logger, tempFile)
		require.NoError(t, err)

		updatedContent, err := os.ReadFile(tempFile)
		require.NoError(t, err)
		contentStr := string(updatedContent)

		assert.NotContains(t, contentStr, "Source Credits (Old)")
		assert.NotContains(t, contentStr, "old-sources-1.yml")
		assert.NotContains(t, contentStr, "old-source-1")

		assert.Contains(t, contentStr, "## Source Credits")
		assert.Contains(t, contentStr, "test-sources.yml")
		assert.Contains(t, contentStr, "[test-source](https://example.com/list)")

		assert.Contains(t, contentStr, "# DNS Toolkit Test")
		assert.Contains(t, contentStr, "## Installation")
		assert.Contains(t, contentStr, "## Source Statistics")
		assert.Contains(t, contentStr, "## Branch Sizes")
	})
}
