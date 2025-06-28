package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEntryContains(t *testing.T) {
	t.Parallel()

	os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	configPath := os.Getenv("DNS_TOOLKIT_TEST_CONFIG_PATH")
	if configPath == "" {
		t.Skip("DNS_TOOLKIT_TEST_CONFIG_PATH is not set, skipping test")
		return
	}

	defer func() {
		os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
	}()

	logger := multilog.NewLogger()
	Logger = logger

	tmpDir := t.TempDir()

	tests := []struct {
		name       string
		content    string
		query      string
		exactMatch bool
		expected   bool
		wantErr    bool
	}{
		{
			name:       "exact match found",
			content:    "example.com\ntest.org\ngoogle.com\n",
			query:      "test.org",
			exactMatch: true,
			expected:   true,
			wantErr:    false,
		},
		{
			name:       "exact match not found",
			content:    "example.com\ntest.org\ngoogle.com\n",
			query:      "test.net",
			exactMatch: true,
			expected:   false,
			wantErr:    false,
		},
		{
			name:       "partial match found",
			content:    "example.com\ntest.org\ngoogle.com\n",
			query:      "test",
			exactMatch: false,
			expected:   true,
			wantErr:    false,
		},
		{
			name:       "partial match not found",
			content:    "example.com\ntest.org\ngoogle.com\n",
			query:      "yahoo",
			exactMatch: false,
			expected:   false,
			wantErr:    false,
		},
		{
			name:       "case insensitive exact match",
			content:    "Example.COM\nTest.ORG\nGoogle.com\n",
			query:      "test.org",
			exactMatch: true,
			expected:   true,
			wantErr:    false,
		},
		{
			name:       "case insensitive partial match",
			content:    "Example.COM\nTest.ORG\nGoogle.com\n",
			query:      "google",
			exactMatch: false,
			expected:   true,
			wantErr:    false,
		},
		{
			name:       "skip comment lines",
			content:    "# This is a comment\nexample.com\n! Another comment\ntest.org\n",
			query:      "comment",
			exactMatch: false,
			expected:   false,
			wantErr:    false,
		},
		{
			name:       "empty file",
			content:    "",
			query:      "test",
			exactMatch: false,
			expected:   false,
			wantErr:    false,
		},
		{
			name:       "whitespace handling",
			content:    "  example.com  \n  test.org  \n",
			query:      "example.com",
			exactMatch: true,
			expected:   true,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFile := filepath.Join(tmpDir, "test_file.txt")
			err := os.WriteFile(testFile, []byte(tt.content), 0644)
			require.NoError(t, err)

			result, err := entryContains(tt.query, testFile, tt.exactMatch)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}

			os.Remove(testFile)
		})
	}
}

func TestEntryContainsFileError(t *testing.T) {
	t.Parallel()

	result, err := entryContains("test", "/nonexistent/file.txt", false)
	assert.Error(t, err)
	assert.False(t, result)
}

func TestCategorizeFileContent(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()

	tests := []struct {
		name     string
		lines    []string
		expected string
	}{
		{
			name:     "mostly domains",
			lines:    []string{"example.com", "test.org", "google.com", "1.2.3.4"},
			expected: constants.SourceTypeDomain,
		},
		{
			name:     "mostly IPv4",
			lines:    []string{"1.2.3.4", "5.6.7.8", "9.10.11.12", "example.com"},
			expected: constants.SourceTypeIpv4,
		},
		{
			name: "mostly IPv6",
			lines: []string{
				"2001:0db8:85a3:0000:0000:8a2e:0370:7334",
				"2001:0db8:85a3:0000:0000:8a2e:0370:7335",
				"2001:0db8:85a3:0000:0000:8a2e:0370:7336",
				"example.com",
			},
			expected: constants.SourceTypeIpv6,
		},
		{
			name:     "CIDR IPv4 entries (but IPv4 regex wins)",
			lines:    []string{"192.168.1.0/24", "10.0.0.0/8", "172.16.0.0/12", "192.168.2.0/24", "10.1.0.0/8"},
			expected: constants.SourceTypeCidrIpv4,
		},
		{
			name:     "IPv4 hostname entries",
			lines:    []string{"1.2.3.4 example.com", "5.6.7.8 test.org", "9.10.11.12 google.com"},
			expected: constants.SourceTypeIpv4Hostname, // IPv4 hostname regex should match first
		},
		{
			name:     "unknown format entries",
			lines:    []string{"||example.com^", "@@||whitelist.com^", "||ads.com^"},
			expected: constants.SourceTypeUnknown,
		},
		{
			name:     "mixed content - ipv4 wins",
			lines:    []string{"example.com", "1.2.3.4", "5.6.7.8", "192.168.1.0/24"},
			expected: constants.SourceTypeIpv4,
		},
		{
			name:     "empty content",
			lines:    []string{},
			expected: constants.SourceTypeUnknown,
		},
		{
			name:     "single domain",
			lines:    []string{"example.com"},
			expected: constants.SourceTypeDomain,
		},
		{
			name:     "single IPv4",
			lines:    []string{"1.2.3.4"},
			expected: constants.SourceTypeIpv4,
		},
		{
			name:     "no matches defaults to unknown",
			lines:    []string{"invalid entry", "another invalid", "not matching anything"},
			expected: constants.SourceTypeUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := categorizeFileContent(logger, tt.lines)
			assert.Equal(t, tt.expected, result)
		})
	}
}
