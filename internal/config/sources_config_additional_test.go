package config

import (
	"os"
	"path/filepath"
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResolveFilePath(t *testing.T) {
	t.Parallel()

	absPath := "/absolute/path/to/file.json"
	result := resolveFilePath(absPath)
	assert.Equal(t, absPath, result)

	relPath := "relative/path/file.json"
	result = resolveFilePath(relPath)
	assert.Equal(t, relPath, result)

	originalTestMode := os.Getenv("DNS_TOOLKIT_TEST_MODE")
	defer os.Setenv("DNS_TOOLKIT_TEST_MODE", originalTestMode)

	os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")

	tempDir, err := os.MkdirTemp("", "test_resolve_*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	goModFile := filepath.Join(tempDir, "go.mod")
	err = os.WriteFile(goModFile, []byte("module test"), 0644)
	require.NoError(t, err)

	testFile := filepath.Join(tempDir, "testfile.json")
	err = os.WriteFile(testFile, []byte("{}"), 0644)
	require.NoError(t, err)

	originalWd, err := os.Getwd()
	require.NoError(t, err)
	defer os.Chdir(originalWd)

	err = os.Chdir(tempDir)
	require.NoError(t, err)

	result = resolveFilePath("testfile.json")
	assert.Equal(t, testFile, result)

	result = resolveFilePath("nonexistent.json")
	assert.Equal(t, "nonexistent.json", result)
}

func TestMatchesFilters(t *testing.T) {
	t.Parallel()

	source := Source{
		Name:      "test-source",
		URL:       "http://example.com",
		Frequency: "daily",
		Types: []c.SourceType{
			{
				Name: "domain",
				ListTypes: []c.ListType{
					{Name: "blocklist", Groups: []string{"security", "normal"}},
					{Name: "allowlist", Groups: []string{"security"}},
				},
			},
			{
				Name: "ip",
				ListTypes: []c.ListType{
					{Name: "blocklist", Groups: []string{"security"}},
				},
			},
		},
		Categories: []string{"malware", "phishing"},
		Countries:  []string{"US", "CA"},
	}

	tests := []struct {
		name     string
		filters  SourceFilters
		expected bool
	}{
		{
			name:     "Empty filters should match",
			filters:  SourceFilters{},
			expected: true,
		},
		{
			name: "Name contains filter matches",
			filters: SourceFilters{
				Name: NameFilter{
					Contains: []string{"test"},
				},
			},
			expected: true,
		},
		{
			name: "Name contains filter no match",
			filters: SourceFilters{
				Name: NameFilter{
					Contains: []string{"nomatch"},
				},
			},
			expected: false,
		},
		{
			name: "Name not contains filter matches",
			filters: SourceFilters{
				Name: NameFilter{
					NotContains: []string{"nomatch"},
				},
			},
			expected: true,
		},
		{
			name: "Name not contains filter excludes",
			filters: SourceFilters{
				Name: NameFilter{
					NotContains: []string{"test"},
				},
			},
			expected: false,
		},
		{
			name: "Type filter matches",
			filters: SourceFilters{
				Type: "domain",
			},
			expected: true,
		},
		{
			name: "Type filter no match",
			filters: SourceFilters{
				Type: "url",
			},
			expected: false,
		},
		{
			name: "Frequency filter matches",
			filters: SourceFilters{
				Frequency: "daily",
			},
			expected: true,
		},
		{
			name: "Frequency filter no match",
			filters: SourceFilters{
				Frequency: "weekly",
			},
			expected: false,
		},
		{
			name: "Category contains filter matches",
			filters: SourceFilters{
				Category: NameFilter{
					Contains: []string{"malware"},
				},
			},
			expected: true,
		},
		{
			name: "Category contains filter no match",
			filters: SourceFilters{
				Category: NameFilter{
					Contains: []string{"spam"},
				},
			},
			expected: false,
		},
		{
			name: "ListType filter matches",
			filters: SourceFilters{
				ListType: "blocklist",
			},
			expected: true,
		},
		{
			name: "ListType filter no match",
			filters: SourceFilters{
				ListType: "greylist",
			},
			expected: false,
		},
		{
			name: "Countries contains filter matches",
			filters: SourceFilters{
				Countries: NameFilter{
					Contains: []string{"US"},
				},
			},
			expected: true,
		},
		{
			name: "Countries contains filter no match",
			filters: SourceFilters{
				Countries: NameFilter{
					Contains: []string{"DE"},
				},
			},
			expected: false,
		},
		{
			name: "Multiple filters all match",
			filters: SourceFilters{
				Type:      "domain",
				Frequency: "daily",
				Group:     "security",
			},
			expected: true,
		},
		{
			name: "Multiple filters partial match",
			filters: SourceFilters{
				Type:      "domain",
				Frequency: "weekly", // This doesn't match
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := matchesFilters(source, tt.filters)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetDownloadFileEdgeCases(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)

	tests := []struct {
		name        string
		source      Source
		expectError bool
	}{
		{
			name: "Source with files should error",
			source: Source{
				Name:  "test-source",
				URL:   "http://example.com",
				Files: []string{"file1.txt", "file2.txt"},
			},
			expectError: true,
		},
		{
			name: "Source with URL ending in slash",
			source: Source{
				Name: "test-source",
				URL:  "http://example.com/",
			},
			expectError: false,
		},
		{
			name: "Source with complex URL",
			source: Source{
				Name: "test-source",
				URL:  "http://example.com/path/to/file.txt?param=value",
			},
			expectError: false,
		},
		{
			name: "Source with different extensions",
			source: Source{
				Name: "test-source",
				URL:  "http://example.com/file.dat",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			downloadFile, err := tt.source.GetDownloadFile(logger, "/tmp")
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.source.Name, downloadFile.Name)
				assert.Equal(t, tt.source.URL, downloadFile.URL)
			}
		})
	}
}

func TestSourcesConfigComplexFiltering(t *testing.T) {
	t.Parallel()

	config := SourcesConfig{
		Sources: []Source{
			{
				Name:      "malware-domains",
				URL:       "http://example.com/malware",
				Frequency: "daily",
				Types: []c.SourceType{
					{
						Name: "domain",
						ListTypes: []c.ListType{
							{Name: "blocklist"},
						},
					},
				},
				Categories: []string{"malware"},
				Countries:  []string{"US"},
				Disabled:   false,
			},
			{
				Name:      "phishing-ips",
				URL:       "http://example.com/phishing",
				Frequency: "weekly",
				Types: []c.SourceType{
					{
						Name: "ip",
						ListTypes: []c.ListType{
							{Name: "blocklist"},
						},
					},
				},
				Categories: []string{"phishing"},
				Countries:  []string{"CA"},
				Disabled:   false,
			},
			{
				Name:      "allowlist-domains",
				URL:       "http://example.com/allowlist",
				Frequency: "daily",
				Types: []c.SourceType{
					{
						Name: "domain",
						ListTypes: []c.ListType{
							{Name: "allowlist"},
						},
					},
				},
				Categories: []string{"whitelist"},
				Countries:  []string{"US"},
				Disabled:   false,
			},
			{
				Name:     "disabled-source",
				URL:      "http://example.com/disabled",
				Disabled: true,
			},
		},
	}

	// Test complex filtering scenarios
	tests := []struct {
		name          string
		filters       SourceFilters
		expectedCount int
		expectedNames []string
	}{
		{
			name:          "All enabled sources",
			filters:       SourceFilters{},
			expectedCount: 3, // Excluding disabled source
			expectedNames: []string{"malware-domains", "phishing-ips", "allowlist-domains"},
		},
		{
			name: "Filter by daily frequency",
			filters: SourceFilters{
				Frequency: "daily",
			},
			expectedCount: 2,
			expectedNames: []string{"malware-domains", "allowlist-domains"},
		},
		{
			name: "Filter by domain type",
			filters: SourceFilters{
				Type: "domain",
			},
			expectedCount: 2,
			expectedNames: []string{"malware-domains", "allowlist-domains"},
		},
		{
			name: "Filter by blocklist type",
			filters: SourceFilters{
				ListType: "blocklist",
			},
			expectedCount: 2,
			expectedNames: []string{"malware-domains", "phishing-ips"},
		},
		{
			name: "Filter by US country",
			filters: SourceFilters{
				Countries: NameFilter{
					Contains: []string{"US"},
				},
			},
			expectedCount: 2,
			expectedNames: []string{"malware-domains", "allowlist-domains"},
		},
		{
			name: "Complex filter: domain + daily frequency",
			filters: SourceFilters{
				Type:      "domain",
				Frequency: "daily",
			},
			expectedCount: 2,
			expectedNames: []string{"malware-domains", "allowlist-domains"},
		},
		{
			name: "Filter excluding certain categories",
			filters: SourceFilters{
				Category: NameFilter{
					NotContains: []string{"whitelist"},
				},
			},
			expectedCount: 2,
			expectedNames: []string{"malware-domains", "phishing-ips"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enabled := config.GetEnabledSources(tt.filters)
			assert.Len(t, enabled, tt.expectedCount)

			actualNames := make([]string, len(enabled))
			for i, source := range enabled {
				actualNames[i] = source.Name
			}

			for _, expectedName := range tt.expectedNames {
				assert.Contains(t, actualNames, expectedName)
			}
		})
	}
}
