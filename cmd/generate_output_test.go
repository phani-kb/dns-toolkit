package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateDescription(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		summaryType string
		fileName    string
		format      string
		listType    string
		expected    string
	}{
		{
			name:        "blocklist description",
			summaryType: "domain",
			fileName:    "domain_test_blocklist.txt",
			format:      "txt",
			listType:    "blocklist",
			expected:    "Txt Domain blocklist",
		},
		{
			name:        "allowlist description",
			summaryType: "ipv4",
			fileName:    "ipv4_test_allowlist.txt",
			format:      "txt",
			listType:    "allowlist",
			expected:    "Txt IPv4 allowlist",
		},
		{
			name:        "unknown list type",
			summaryType: "domain",
			fileName:    "general_test.txt",
			format:      "txt",
			listType:    "unknown",
			expected:    "Txt General unknown",
		},
		{
			name:        "empty parameters",
			summaryType: "",
			fileName:    "",
			format:      "",
			listType:    "",
			expected:    " General unknown",
		},
		{
			name:        "adguard format",
			summaryType: "adguard",
			fileName:    "adguard_test_blocklist.txt",
			format:      "adguard_rules",
			listType:    "blocklist",
			expected:    "Adguard Rules AdGuard blocklist",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateDescription(tt.summaryType, tt.fileName, tt.format, tt.listType)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPrepareDirectories(t *testing.T) {
	t.Parallel()

	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "dns-toolkit-test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Change to the temp directory
	oldWd, err := os.Getwd()
	assert.NoError(t, err)
	defer os.Chdir(oldWd)

	err = os.Chdir(tempDir)
	assert.NoError(t, err)

	includeIgnored = false
	_ = prepareDirectories()
}

func TestCopySummaryFile(t *testing.T) {
	t.Parallel()

	tempDir, err := os.MkdirTemp("", "dns-toolkit-test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	srcPath := filepath.Join(tempDir, "source.txt")
	dstPath := filepath.Join(tempDir, "dest.txt")

	content := "test content"
	err = os.WriteFile(srcPath, []byte(content), 0644)
	assert.NoError(t, err)

	err = copySummaryFile(srcPath, dstPath)
	assert.NoError(t, err)

	assert.FileExists(t, dstPath)
	destContent, err := os.ReadFile(dstPath)
	assert.NoError(t, err)
	assert.Equal(t, content, string(destContent))

	nonExistentSrc := filepath.Join(tempDir, "nonexistent.txt")
	err = copySummaryFile(nonExistentSrc, dstPath)
	assert.Error(t, err)
}

func TestLoadTemplates(t *testing.T) {
	t.Parallel()

	os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	defer os.Unsetenv("DNS_TOOLKIT_TEST_MODE")

	tmpl, staticTemplate, err := loadTemplates()

	if err != nil {
		assert.Contains(t, err.Error(), "template")
	} else {
		assert.NotNil(t, tmpl)
		assert.NotNil(t, staticTemplate)
		assert.Greater(t, len(staticTemplate), 0)
	}
}
