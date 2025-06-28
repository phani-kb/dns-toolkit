package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// createTestSources creates a sample sources configuration for testing
func createTestSources() SourcesConfig {
	return SourcesConfig{
		Sources: []Source{
			{
				Name:       "test-source1",
				URL:        "http://example.com/list1.txt",
				Frequency:  "daily",
				Categories: []string{"malware", "security"},
				Types: []c.SourceType{
					{
						Name: "domain",
						ListTypes: []c.ListType{
							{
								Name:         "blocklist",
								Groups:       []string{"big", "normal"},
								MustConsider: true,
							},
						},
					},
				},
			},
			{
				Name:      "test-source2",
				URL:       "http://example.com/list2.txt",
				Disabled:  true,
				Frequency: "weekly",
				Types: []c.SourceType{
					{
						Name: "ipv4",
						ListTypes: []c.ListType{
							{
								Name:   "allowlist",
								Groups: []string{"mini"},
							},
						},
					},
				},
			},
			{
				Name:      "test-source3",
				URL:       "http://example.com/list3.txt",
				Frequency: "monthly",
				Countries: []string{"US", "CA"},
				Types: []c.SourceType{
					{
						Name: "domain",
						ListTypes: []c.ListType{
							{
								Name:   "blocklist",
								Groups: []string{"lite"},
							},
						},
					},
					{
						Name: "ipv4",
						ListTypes: []c.ListType{
							{
								Name:   "blocklist",
								Groups: []string{"normal"},
							},
						},
					},
				},
			},
		},
	}
}

// TestSourcesConfigValidation tests the validation of sources config
func TestSourcesConfigValidation(t *testing.T) {
	t.Run("Valid config validates successfully", func(t *testing.T) {
		config := createTestSources()
		err := config.Validate()
		assert.NoError(t, err)
	})

	t.Run("Empty sources array fails validation", func(t *testing.T) {
		config := SourcesConfig{
			Sources: []Source{},
		}
		err := config.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "at least one source is required")
	})

	t.Run("Duplicate source names fail validation", func(t *testing.T) {
		config := SourcesConfig{
			Sources: []Source{
				{
					Name: "duplicate",
					URL:  "http://example.com/list1.txt",
					Types: []c.SourceType{
						{Name: "domain"},
					},
				},
				{
					Name: "duplicate",
					URL:  "http://example.com/list2.txt",
					Types: []c.SourceType{
						{Name: "domain"},
					},
				},
			},
		}
		err := config.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "duplicate source name")
	})
}

// TestSourceValidation tests validation of individual source entries
func TestSourceValidation(t *testing.T) {
	t.Run("Valid source validates successfully", func(t *testing.T) {
		source := Source{
			Name: "test-source",
			URL:  "http://example.com/list.txt",
			Types: []c.SourceType{
				{Name: "domain"},
			},
		}
		err := source.Validate()
		assert.NoError(t, err)
	})

	t.Run("Missing name fails validation", func(t *testing.T) {
		source := Source{
			URL: "http://example.com/list.txt",
			Types: []c.SourceType{
				{Name: "domain"},
			},
		}
		err := source.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "name is required")
	})

	t.Run("Missing URL fails validation", func(t *testing.T) {
		source := Source{
			Name: "test-source",
			Types: []c.SourceType{
				{Name: "domain"},
			},
		}
		err := source.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "url is required")
	})

	t.Run("Missing types fails validation", func(t *testing.T) {
		source := Source{
			Name: "test-source",
			URL:  "http://example.com/list.txt",
		}
		err := source.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "at least one type is required")
	})

	t.Run("Invalid country code fails validation", func(t *testing.T) {
		source := Source{
			Name:      "test-source",
			URL:       "http://example.com/list.txt",
			Types:     []c.SourceType{{Name: "domain"}},
			Countries: []string{"USA"}, // Should be 2-letter code
		}
		err := source.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid country code")
	})

	t.Run("Duplicate country codes fail validation", func(t *testing.T) {
		source := Source{
			Name:      "test-source",
			URL:       "http://example.com/list.txt",
			Types:     []c.SourceType{{Name: "domain"}},
			Countries: []string{"US", "US"}, // Duplicate
		}
		err := source.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "duplicate country")
	})

	t.Run("Duplicate type names fail validation", func(t *testing.T) {
		source := Source{
			Name: "test-source",
			URL:  "http://example.com/list.txt",
			Types: []c.SourceType{
				{Name: "domain"},
				{Name: "domain"}, // Duplicate
			},
		}
		err := source.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "duplicate type")
	})
}

// TestSourceFiltering tests filtering sources with various criteria
func TestSourceFiltering(t *testing.T) {
	config := createTestSources()

	t.Run("Get all enabled sources", func(t *testing.T) {
		sources := config.GetEnabledSources(SourceFilters{})
		assert.Equal(t, 2, len(sources))
		// test-source2 should be excluded as it's disabled
		sourceNames := []string{sources[0].Name, sources[1].Name}
		assert.Contains(t, sourceNames, "test-source1")
		assert.Contains(t, sourceNames, "test-source3")
		assert.NotContains(t, sourceNames, "test-source2")
	})

	t.Run("Filter by name contains", func(t *testing.T) {
		filters := SourceFilters{
			Name: NameFilter{
				Contains: []string{"source1"},
			},
		}
		sources := config.GetEnabledSources(filters)
		found := false
		for _, s := range sources {
			if s.Name == "test-source1" {
				found = true
				break
			}
		}
		assert.True(t, found, "Should contain test-source1")
	})

	t.Run("Filter by name not contains", func(t *testing.T) {
		filters := SourceFilters{
			Name: NameFilter{
				NotContains: []string{"source1"},
			},
		}
		sources := config.GetEnabledSources(filters)

		found := false
		for _, s := range sources {
			if s.Name == "test-source3" {
				found = true
				break
			}
		}
		assert.True(t, found, "Should contain test-source3")
	})

	t.Run("Filter by type", func(t *testing.T) {
		filters := SourceFilters{
			Type: "ipv4",
		}
		sources := config.GetEnabledSources(filters)

		found := false
		for _, s := range sources {
			if s.Name == "test-source3" {
				found = true
				break
			}
		}
		assert.True(t, found, "Should contain test-source3")
	})

	t.Run("Complex filtering", func(t *testing.T) {
		filters := SourceFilters{
			Type: "domain",
			Name: NameFilter{
				Contains:    []string{"test"},
				NotContains: []string{"source2"},
			},
		}
		sources := config.GetEnabledSources(filters)
		assert.Equal(t, 2, len(sources))
		sourceNames := []string{sources[0].Name, sources[1].Name}
		assert.Contains(t, sourceNames, "test-source1")
		assert.Contains(t, sourceNames, "test-source3")
	})
}

// TestGetSourceByField tests retrieving sources by specific fields
func TestGetSourceByField(t *testing.T) {
	config := createTestSources()

	t.Run("By list type", func(t *testing.T) {
		sources := config.GetSourceByField("listType", "blocklist")

		foundSource1 := false
		foundSource3 := false
		for _, s := range sources {
			if s.Name == "test-source1" {
				foundSource1 = true
			}
			if s.Name == "test-source3" {
				foundSource3 = true
			}
		}
		assert.True(t, foundSource1, "Should contain test-source1")
		assert.True(t, foundSource3, "Should contain test-source3")
	})

	t.Run("By source type", func(t *testing.T) {
		sources := config.GetSourceByField("type", "ipv4")
		assert.Equal(t, 2, len(sources))
		sourceNames := map[string]bool{sources[0].Name: true, sources[1].Name: true}
		assert.True(t, sourceNames["test-source2"])
		assert.True(t, sourceNames["test-source3"])

		sources = config.GetSourceByField("type", "domain")
		assert.Equal(t, 2, len(sources))
		sourceNames = map[string]bool{sources[0].Name: true, sources[1].Name: true}
		assert.True(t, sourceNames["test-source1"])
		assert.True(t, sourceNames["test-source3"])
	})

	t.Run("By frequency", func(t *testing.T) {
		sources := config.GetSourceByField("frequency", "daily")
		assert.Equal(t, 1, len(sources))
		assert.Equal(t, "test-source1", sources[0].Name)

		sources = config.GetSourceByField("frequency", "monthly")
		assert.Equal(t, 1, len(sources))
		assert.Equal(t, "test-source3", sources[0].Name)
	})

	t.Run("By category", func(t *testing.T) {
		sources := config.GetSourceByField("category", "security")
		assert.Equal(t, 1, len(sources))
		assert.Equal(t, "test-source1", sources[0].Name)
	})

	t.Run("By country", func(t *testing.T) {
		sources := config.GetSourceByField("country", "US")
		assert.Equal(t, 1, len(sources))
		assert.Equal(t, "test-source3", sources[0].Name)
	})

	t.Run("No matching field", func(t *testing.T) {
		sources := config.GetSourceByField("nonexistent", "value")
		assert.Equal(t, 0, len(sources))
	})

	t.Run("No matching value", func(t *testing.T) {
		sources := config.GetSourceByField("type", "nonexistent")
		assert.Equal(t, 0, len(sources))
	})
}

// TestGetSourceByName tests retrieving sources by name
func TestGetSourceByName(t *testing.T) {
	config := createTestSources()

	t.Run("Existing source", func(t *testing.T) {
		source, found := config.GetSourceByName("test-source1")
		assert.True(t, found)
		assert.Equal(t, "test-source1", source.Name)
	})

	t.Run("Non-existent source", func(t *testing.T) {
		_, found := config.GetSourceByName("non-existent")
		assert.False(t, found)
	})
}

// TestConfigIsEnabled tests the IsEnabled method on SourcesConfig
func TestConfigIsEnabled(t *testing.T) {
	config := createTestSources()

	t.Run("Enabled source", func(t *testing.T) {
		isEnabled := config.IsEnabled("test-source1")
		assert.True(t, isEnabled)
	})

	t.Run("Disabled source", func(t *testing.T) {
		isEnabled := config.IsEnabled("test-source2")
		assert.False(t, isEnabled)
	})

	t.Run("Non-existent source", func(t *testing.T) {
		isEnabled := config.IsEnabled("non-existent")
		assert.False(t, isEnabled)
	})
}

// TestSourceIsEnabled tests the IsEnabled method on Source
func TestSourceIsEnabled(t *testing.T) {
	t.Run("Enabled source", func(t *testing.T) {
		source := Source{
			Name: "test-source",
		}
		assert.True(t, source.IsEnabled())
	})

	t.Run("Disabled source", func(t *testing.T) {
		source := Source{
			Name:     "test-source",
			Disabled: true,
		}
		assert.False(t, source.IsEnabled())
	})
}

// TestSourceJSONUnmarshaling tests unmarshaling of JSON into the Source struct
func TestSourceJSONUnmarshaling(t *testing.T) {
	t.Run("Basic source", func(t *testing.T) {
		jsonData := `{
			"name": "test-source",
			"url": "http://example.com/list.txt",
			"types": [
				{
					"name": "domain"
				}
			]
		}`

		var source Source
		err := json.Unmarshal([]byte(jsonData), &source)
		assert.NoError(t, err)
		assert.Equal(t, "test-source", source.Name)
		assert.Equal(t, "http://example.com/list.txt", source.URL)
		assert.Equal(t, 1, len(source.Types))
		assert.Equal(t, "domain", source.Types[0].Name)
		assert.Equal(t, 1, source.TypeCount)
	})

	t.Run("Source with list types", func(t *testing.T) {
		jsonData := `{
			"name": "test-source",
			"url": "http://example.com/list.txt",
			"types": [
				{
					"name": "domain",
					"list_types": [
						{
							"name": "blocklist",
							"groups": "big,normal",
							"must_consider": true
						}
					]
				}
			]
		}`

		var source Source
		err := json.Unmarshal([]byte(jsonData), &source)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(source.Types))
		assert.Equal(t, 1, len(source.Types[0].ListTypes))
		assert.Equal(t, "blocklist", source.Types[0].ListTypes[0].Name)
		assert.Equal(t, 2, len(source.Types[0].ListTypes[0].Groups))
		assert.Contains(t, source.Types[0].ListTypes[0].Groups, "big")
		assert.Contains(t, source.Types[0].ListTypes[0].Groups, "normal")
		assert.True(t, source.Types[0].ListTypes[0].MustConsider)
	})

	t.Run("Source with comma-separated fields", func(t *testing.T) {
		jsonData := `{
			"name": "test-source",
			"url": "http://example.com/list.txt",
			"types": [{"name": "domain"}],
			"categories": "malware, ads",
			"countries": "US, CA",
			"files": "file1, file2"
		}`

		var source Source
		err := json.Unmarshal([]byte(jsonData), &source)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(source.Categories))
		assert.Contains(t, source.Categories, "malware")
		assert.Contains(t, source.Categories, "ads")
		assert.Equal(t, 2, len(source.Countries))
		assert.Contains(t, source.Countries, "US")
		assert.Contains(t, source.Countries, "CA")
		assert.Equal(t, 2, len(source.Files))
		assert.Contains(t, source.Files, "file1")
		assert.Contains(t, source.Files, "file2")
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		jsonData := `{invalid json`

		var source Source
		err := json.Unmarshal([]byte(jsonData), &source)
		assert.Error(t, err)
	})
}

// TestLoadingSourcesWithDefaultValues tests that default values are properly applied when loading a sources config
func TestLoadingSourcesWithDefaultValues(t *testing.T) {
	testDir, err := os.MkdirTemp("", "dns-toolkit-test")
	require.NoError(t, err)
	defer os.RemoveAll(testDir)

	// Create source file with minimal values that should get defaults applied
	minimalSourceContent := `{
		"sources": [
			{
				"name": "minimal-source",
				"url": "http://example.com/list.txt",
				"types": [
					{
						"name": "domain",
						"list_types": []
					}
				]
			},
			{
				"name": "minimal-source-no-listtypes",
				"url": "http://example.com/list2.txt",
				"types": [
					{
						"name": "domain"
					}
				]
			},
			{
				"name": "source-with-mini-group",
				"url": "http://example.com/list3.txt",
				"types": [
					{
						"name": "domain",
						"list_types": [
							{
								"name": "blocklist",
								"groups": "mini"
							}
						]
					}
				]
			}
		]
	}`

	sourcesPath := filepath.Join(testDir, "sources_with_defaults.json")
	err = os.WriteFile(sourcesPath, []byte(minimalSourceContent), 0644)
	require.NoError(t, err)

	logger := createTestLogger(t)
	config, err := LoadSourcesConfig(logger, sourcesPath)
	require.NoError(t, err)

	source1 := config.Sources[0]
	assert.Equal(t, 1, len(source1.Types))
	assert.Equal(t, "domain", source1.Types[0].Name)
	assert.Equal(t, 1, len(source1.Types[0].ListTypes))
	assert.Equal(t, "blocklist", source1.Types[0].ListTypes[0].Name)
	assert.Contains(t, source1.Types[0].ListTypes[0].Groups, "big") // Default group

	source2 := config.Sources[1]
	assert.Equal(t, 1, len(source2.Types))
	assert.Equal(t, "domain", source2.Types[0].Name)
	assert.Equal(t, 1, len(source2.Types[0].ListTypes))
	assert.Equal(t, "blocklist", source2.Types[0].ListTypes[0].Name)
	assert.Contains(t, source2.Types[0].ListTypes[0].Groups, "big") // Default group

	source3 := config.Sources[2]
	assert.Equal(t, 1, len(source3.Types))
	assert.Equal(t, "domain", source3.Types[0].Name)
	assert.Equal(t, 1, len(source3.Types[0].ListTypes))
	assert.Equal(t, "blocklist", source3.Types[0].ListTypes[0].Name)

	groups := source3.Types[0].ListTypes[0].Groups
	assert.Contains(t, groups, "mini")

	assert.True(t, len(groups) > 1, "Mini group should be expanded to include less restrictive groups")
}

// TestGetDownloadFileWithExtensions tests the GetDownloadFile method with various URL extensions
func TestGetDownloadFileWithExtensions(t *testing.T) {
	testDir, err := os.MkdirTemp("", "dns-toolkit-test")
	require.NoError(t, err)
	defer os.RemoveAll(testDir)

	downloadDir := filepath.Join(testDir, "download")
	err = os.MkdirAll(downloadDir, 0755)
	require.NoError(t, err)

	logger := createTestLogger(t)

	testCases := []struct {
		name          string
		url           string
		expectedExt   string
		expectArchive bool
	}{
		{
			name:          "Text file extension",
			url:           "http://example.com/list.txt",
			expectedExt:   ".txt",
			expectArchive: false,
		},
		{
			name:          "Data file extension",
			url:           "http://example.com/list.dat",
			expectedExt:   ".txt", // Expected to be converted to .txt
			expectArchive: false,
		},
		{
			name:          "No extension",
			url:           "http://example.com/list",
			expectedExt:   ".txt",
			expectArchive: false,
		},
		{
			name:          "Custom extension",
			url:           "http://example.com/list.custom",
			expectedExt:   ".txt",
			expectArchive: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			source := Source{
				Name: "test-source",
				URL:  tc.url,
				Types: []c.SourceType{
					{Name: "domain"},
				},
			}

			downloadFile, err := source.GetDownloadFile(logger, downloadDir)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectArchive, downloadFile.IsArchive)
			assert.Equal(t, "test-source"+tc.expectedExt, downloadFile.Filename)
			assert.Equal(t, 1, len(downloadFile.Targets))
		})
	}
}

// TestGetDownloadFileDetailedCases tests more edge cases of the GetDownloadFile method
func TestGetDownloadFileDetailedCases(t *testing.T) {
	testDir, err := os.MkdirTemp("", "dns-toolkit-test")
	require.NoError(t, err)
	defer os.RemoveAll(testDir)

	downloadDir := filepath.Join(testDir, "download")
	err = os.MkdirAll(downloadDir, 0755)
	require.NoError(t, err)

	logger := createTestLogger(t)

	t.Run("Source with files specified causes error", func(t *testing.T) {
		source := Source{
			Name:  "source-with-files",
			URL:   "http://example.com/list.txt",
			Files: []string{"file1.txt", "file2.txt"},
			Types: []c.SourceType{
				{Name: "domain"},
			},
		}

		_, err := source.GetDownloadFile(logger, downloadDir)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "files are not required")
	})

	t.Run("Verify download target properties", func(t *testing.T) {
		source := Source{
			Name: "target-test-source",
			URL:  "http://example.com/list.txt",
			Types: []c.SourceType{
				{Name: "domain"},
			},
		}

		downloadFile, err := source.GetDownloadFile(logger, downloadDir)
		assert.NoError(t, err)
		assert.Equal(t, "target-test-source", downloadFile.Name)
		assert.Equal(t, "http://example.com/list.txt", downloadFile.URL)
		assert.Equal(t, downloadDir, downloadFile.Folder)
		assert.Equal(t, "target-test-source.txt", downloadFile.Filename)
		assert.Equal(t, 1, len(downloadFile.Targets))

		target := downloadFile.Targets[0]
		assert.Equal(t, downloadDir, target.SourceFolder)
		assert.Equal(t, "target-test-source.txt", target.SourceFile)
		assert.Equal(t, "target-test-source.txt", target.TargetFile)
		assert.Equal(t, downloadDir, target.TargetFolder)
	})
}

func TestCfgGetUserAgentFunction(t *testing.T) {
	t.Parallel()

	logFile, err := os.CreateTemp("", "test-logger-*.log")
	assert.NoError(t, err)
	defer os.Remove(logFile.Name())
	defer logFile.Close()

	logger := createTestLogger(t)

	appConfig := ApplicationConfig{
		Name:        "test-app",
		Version:     "1.2.3",
		Description: "Test Description",
	}

	userAgent := GetUserAgent(logger, appConfig)

	assert.Contains(t, userAgent, "test-app")
	assert.Contains(t, userAgent, "1.2.3")
	assert.Contains(t, userAgent, "Test Description")

	emptyConfig := ApplicationConfig{}
	defaultUserAgent := GetUserAgent(logger, emptyConfig)

	assert.Contains(t, defaultUserAgent, constants.AppName)
	assert.Contains(t, defaultUserAgent, constants.AppVersion)
	assert.Contains(t, defaultUserAgent, constants.AppDescription)
}
