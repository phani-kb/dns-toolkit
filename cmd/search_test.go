package cmd

import (
	"net"
	"os"
	"path/filepath"
	"testing"

	u "github.com/phani-kb/dns-toolkit/internal/utils"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEntryContains(t *testing.T) {
	err := os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	assert.NoError(t, err)
	configPath := os.Getenv("DNS_TOOLKIT_TEST_CONFIG_PATH")
	if configPath == "" {
		t.Skip("DNS_TOOLKIT_TEST_CONFIG_PATH is not set, skipping test")
		return
	}

	defer func() {
		err := os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
		if err != nil {
			t.Logf("Failed to unset DNS_TOOLKIT_TEST_MODE: %v", err)
		}
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

			if err := os.Remove(testFile); err != nil {
				t.Logf("Failed to remove test file: %v", err)
			}
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
		expected string
		lines    []string
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

func Test_resolveDomainToIPs(t *testing.T) {
	tests := []struct {
		name         string
		domain       string
		lookupResult []net.IP
		wantIPs      []string
	}{
		{
			name:   "single IPv4",
			domain: "example.com",
			lookupResult: []net.IP{
				net.ParseIP("1.2.3.4"),
			},
			wantIPs: []string{"1.2.3.4"},
		},
		{
			name:   "IPv4 and IPv6",
			domain: "mixed.com",
			lookupResult: []net.IP{
				net.ParseIP("5.6.7.8"),
				net.ParseIP("2001:db8::1"),
			},
			wantIPs: []string{"5.6.7.8"},
		},
		{
			name:   "no IPv4",
			domain: "ipv6only.com",
			lookupResult: []net.IP{
				net.ParseIP("2001:db8::2"),
			},
			wantIPs: nil,
		},
		{
			name:   "multiple IPv4",
			domain: "multi.com",
			lookupResult: []net.IP{
				net.ParseIP("8.8.8.8"),
				net.ParseIP("8.8.4.4"),
			},
			wantIPs: []string{"8.8.8.8", "8.8.4.4"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLookup := func(host string) ([]net.IP, error) {
				return tt.lookupResult, nil
			}

			ipSet := u.NewStringSet(nil)
			resolveDomainToIPs(tt.domain, ipSet, mockLookup)
			got := ipSet.ToSlice()
			assert.ElementsMatch(t, tt.wantIPs, got)
		})
	}
}

func TestGetProcessedFiles(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "processed-files-test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	// Create test processed summary file
	summaryFile := filepath.Join(tempDir, "processed_summary.json")
	summaryContent := `[
		{
			"name": "test-source",
			"valid_files": [
				{
					"filepath": "/path/to/domain1.txt",
					"generic_source_type": "domain"
				},
				{
					"filepath": "/path/to/domain2.txt", 
					"generic_source_type": "domain"
				},
				{
					"filepath": "/path/to/ipv4.txt",
					"generic_source_type": "ipv4"
				}
			]
		}
	]`
	require.NoError(t, os.WriteFile(summaryFile, []byte(summaryContent), 0644))

	// Test getting domain files
	files, err := getProcessedFiles("domain", summaryFile)
	require.NoError(t, err)
	assert.Len(t, files, 2)
	assert.Contains(t, files, "/path/to/domain1.txt")
	assert.Contains(t, files, "/path/to/domain2.txt")

	// Test getting ipv4 files
	files, err = getProcessedFiles("ipv4", summaryFile)
	require.NoError(t, err)
	assert.Len(t, files, 1)
	assert.Contains(t, files, "/path/to/ipv4.txt")

	// Test with non-existent source type
	files, err = getProcessedFiles("unknown", summaryFile)
	require.NoError(t, err)
	assert.Len(t, files, 0)

	// Test with non-existent file
	_, err = getProcessedFiles("domain", "/non/existent/file.json")
	assert.Error(t, err)
}

func TestGetConsolidatedFiles(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "consolidated-files-test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	// Create test consolidated summary file
	summaryFile := filepath.Join(tempDir, "consolidated_summary.json")
	summaryContent := `[
		{
			"type": "domain",
			"list_type": "blocklist",
			"filepath": "/path/to/domain_blocklist.txt"
		},
		{
			"type": "domain",
			"list_type": "allowlist", 
			"filepath": "/path/to/domain_allowlist.txt"
		},
		{
			"type": "ipv4",
			"list_type": "blocklist",
			"filepath": "/path/to/ipv4_blocklist.txt"
		}
	]`
	require.NoError(t, os.WriteFile(summaryFile, []byte(summaryContent), 0644))

	// Test getting domain files
	files, err := getConsolidatedFiles("domain", summaryFile)
	require.NoError(t, err)
	assert.Len(t, files, 2)
	assert.Contains(t, files, "/path/to/domain_blocklist.txt")
	assert.Contains(t, files, "/path/to/domain_allowlist.txt")

	// Test getting ipv4 files
	files, err = getConsolidatedFiles("ipv4", summaryFile)
	require.NoError(t, err)
	assert.Len(t, files, 1)
	assert.Contains(t, files, "/path/to/ipv4_blocklist.txt")

	files, err = getConsolidatedFiles("unknown", summaryFile)
	require.NoError(t, err)
	assert.Len(t, files, 0)

	_, err = getConsolidatedFiles("domain", "/non/existent/file.json")
	assert.Error(t, err)
}

func TestDisplayCnameResults(t *testing.T) {
	displayCnameResults([]string{}, map[string][]string{})

	cnames := []string{"cdn.example.com"}
	cnameResults := map[string][]string{
		"source1": {"/path/to/file1.txt", "/path/to/file2.txt"},
		"source2": {"/path/to/file3.txt"},
	}

	displayCnameResults(cnames, cnameResults)
	displayCnameResults([]string{"test.com"}, map[string][]string{})
	displayCnameResults([]string{}, map[string][]string{"source1": {"/file.txt"}})
}

func TestDisplayDomainResults(t *testing.T) {
	domainResults := map[string][]string{
		"source1": {"/path/to/file1.txt", "/path/to/file2.txt"},
		"source2": {"/path/to/file3.txt"},
	}
	displayDomainResults("example.com", domainResults)
	displayDomainResults("notfound.com", map[string][]string{})
	displayDomainResults("", map[string][]string{})
}
