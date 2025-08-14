package consolidators_test

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/consolidators"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCommonConsolidatorConsolidate tests the Consolidate method
func TestCommonConsolidatorConsolidate(t *testing.T) {
	tests := []struct {
		name           string
		sourceType     string
		listType       string
		files          []fileData
		expectedCount  int
		expectedFiles  int
		shouldHaveData bool
	}{
		{
			name:       "domain_blocklist_consolidation",
			sourceType: constants.SourceTypeDomain,
			listType:   constants.ListTypeBlocklist,
			files: []fileData{
				{
					name:    "domains1.txt",
					content: "example.com\ntest.com\nbad.domain.com\n",
					entries: 3,
				},
				{
					name:    "domains2.txt",
					content: "malware.com\nexample.com\nspam.net\n", // example.com is duplicate
					entries: 3,
				},
			},
			expectedCount:  5, // 3 + 3 - 1 duplicate = 5 unique
			expectedFiles:  2,
			shouldHaveData: true,
		},
		{
			name:       "ipv4_allowlist_consolidation",
			sourceType: constants.SourceTypeIpv4,
			listType:   constants.ListTypeAllowlist,
			files: []fileData{
				{
					name:    "ips1.txt",
					content: "192.168.1.1\n10.0.0.1\n172.16.0.1\n",
					entries: 3,
				},
				{
					name:    "ips2.txt",
					content: "8.8.8.8\n1.1.1.1\n192.168.1.1\n", // 192.168.1.1 is duplicate
					entries: 3,
				},
			},
			expectedCount:  5, // 3 + 3 - 1 duplicate = 5 unique
			expectedFiles:  2,
			shouldHaveData: true,
		},
		{
			name:       "ipv6_blocklist_consolidation",
			sourceType: constants.SourceTypeIpv6,
			listType:   constants.ListTypeBlocklist,
			files: []fileData{
				{
					name:    "ipv6s.txt",
					content: "2001:db8::1\n2001:db8::2\n",
					entries: 2,
				},
			},
			expectedCount:  2,
			expectedFiles:  1,
			shouldHaveData: true,
		},
		{
			name:       "cidr_ipv4_consolidation",
			sourceType: constants.SourceTypeCidrIpv4,
			listType:   constants.ListTypeBlocklist,
			files: []fileData{
				{
					name:    "cidrs.txt",
					content: "192.168.0.0/24\n10.0.0.0/8\n172.16.0.0/12\n",
					entries: 3,
				},
			},
			expectedCount:  3,
			expectedFiles:  1,
			shouldHaveData: true,
		},
		{
			name:           "empty_files",
			sourceType:     constants.SourceTypeDomain,
			listType:       constants.ListTypeBlocklist,
			files:          []fileData{},
			expectedCount:  0,
			expectedFiles:  0,
			shouldHaveData: false,
		},
		{
			name:       "invalid_source_type_files",
			sourceType: constants.SourceTypeDomain,
			listType:   constants.ListTypeBlocklist,
			files: []fileData{
				{
					name:         "wrong_type.txt",
					content:      "example.com\ntest.com\n",
					entries:      2,
					sourceType:   constants.SourceTypeIpv4, // Wrong source type
					listType:     constants.ListTypeBlocklist,
					shouldIgnore: true,
				},
			},
			expectedCount:  0,
			expectedFiles:  0,
			shouldHaveData: false,
		},
		{
			name:       "entry_count_mismatch",
			sourceType: constants.SourceTypeDomain,
			listType:   constants.ListTypeBlocklist,
			files: []fileData{
				{
					name:         "mismatch.txt",
					content:      "example.com\ntest.com\n",
					entries:      5, // Wrong count, should be 2
					shouldIgnore: true,
				},
			},
			expectedCount:  0,
			expectedFiles:  0,
			shouldHaveData: false,
		},
		{
			name:       "mixed_valid_invalid_files",
			sourceType: constants.SourceTypeDomain,
			listType:   constants.ListTypeBlocklist,
			files: []fileData{
				{
					name:    "valid.txt",
					content: "good.com\nvalid.com\n",
					entries: 2,
				},
				{
					name:         "invalid_source.txt",
					content:      "bad.com\nevil.com\n",
					entries:      2,
					sourceType:   constants.SourceTypeIpv4, // Wrong source type
					shouldIgnore: true,
				},
				{
					name:         "invalid_count.txt",
					content:      "wrong.com\ncount.com\n",
					entries:      10, // Wrong count
					shouldIgnore: true,
				},
			},
			expectedCount:  2, // Only from valid.txt
			expectedFiles:  1,
			shouldHaveData: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := multilog.NewLogger()
			cc := consolidators.NewCommonConsolidator(tt.sourceType, tt.listType)

			tempDir := t.TempDir()
			var processedFiles []common.ProcessedFile

			for _, file := range tt.files {
				filePath := filepath.Join(tempDir, file.name)
				require.NoError(t, os.WriteFile(filePath, []byte(file.content), 0644))

				sourceType := file.sourceType
				if sourceType == "" {
					sourceType = tt.sourceType
				}
				listType := file.listType
				if listType == "" {
					listType = tt.listType
				}

				processedFiles = append(processedFiles, common.ProcessedFile{
					Name:              file.name,
					Filepath:          filePath,
					GenericSourceType: sourceType,
					ListType:          listType,
					NumberOfEntries:   file.entries,
					MustConsider:      file.mustConsider,
				})
			}

			consolidatedSet, fileInfos := cc.Consolidate(logger, processedFiles)

			assert.Equal(t, tt.expectedCount, len(consolidatedSet), "Consolidated entries count should match expected")
			assert.Equal(t, tt.expectedFiles, len(fileInfos), "File info count should match expected")

			if tt.shouldHaveData {
				assert.Greater(t, len(consolidatedSet), 0, "Should have consolidated data")
			} else {
				assert.Equal(t, 0, len(consolidatedSet), "Should have no consolidated data")
			}

			validFileCount := 0
			for _, file := range tt.files {
				if !file.shouldIgnore {
					validFileCount++
				}
			}
			assert.Equal(t, validFileCount, len(fileInfos), "File info should only include valid files")
		})
	}
}

// TestCommonConsolidatorFilterEntries tests the FilterEntries method
func TestCommonConsolidatorFilterEntries(t *testing.T) {
	tests := []struct {
		name                string
		sourceType          string
		listType            string
		entries             []entryData
		filterEntries       []entryData
		expectedFiltered    []string
		expectedIgnored     []string
		expectedFilteredLen int
		expectedIgnoredLen  int
	}{
		{
			name:       "domain_basic_filtering",
			sourceType: constants.SourceTypeDomain,
			listType:   constants.ListTypeBlocklist,
			entries: []entryData{
				{value: "example.com", mustConsider: false},
				{value: "test.com", mustConsider: false},
				{value: "important.com", mustConsider: true},
				{value: "blocked.com", mustConsider: false},
			},
			filterEntries: []entryData{
				{value: "test.com", mustConsider: true},     // Should be filtered
				{value: "blocked.com", mustConsider: false}, // Should be ignored
			},
			expectedFiltered:    []string{"example.com", "important.com"},
			expectedIgnored:     []string{"test.com", "blocked.com"},
			expectedFilteredLen: 2,
			expectedIgnoredLen:  2,
		},
		{
			name:       "ipv4_must_consider_override",
			sourceType: constants.SourceTypeIpv4,
			listType:   constants.ListTypeBlocklist,
			entries: []entryData{
				{value: "192.168.1.1", mustConsider: true}, // Must consider, should stay even if filtered
				{value: "10.0.0.1", mustConsider: false},   // Normal entry
				{value: "172.16.0.1", mustConsider: false}, // Will be filtered
			},
			filterEntries: []entryData{
				{value: "192.168.1.1", mustConsider: true}, // Try to filter must-consider entry
				{value: "172.16.0.1", mustConsider: false}, // Normal filter
			},
			expectedFiltered:    []string{"192.168.1.1", "10.0.0.1"}, // Must consider entry stays
			expectedIgnored:     []string{"172.16.0.1"},
			expectedFilteredLen: 2,
			expectedIgnoredLen:  1,
		},
		{
			name:       "ipv6_no_filtering_needed",
			sourceType: constants.SourceTypeIpv6,
			listType:   constants.ListTypeAllowlist,
			entries: []entryData{
				{value: "2001:db8::1", mustConsider: false},
				{value: "2001:db8::2", mustConsider: false},
			},
			filterEntries:       []entryData{}, // No filter entries
			expectedFiltered:    []string{"2001:db8::1", "2001:db8::2"},
			expectedIgnored:     []string{},
			expectedFilteredLen: 2,
			expectedIgnoredLen:  0,
		},
		{
			name:       "cidr_empty_entries",
			sourceType: constants.SourceTypeCidrIpv4,
			listType:   constants.ListTypeBlocklist,
			entries:    []entryData{}, // No entries
			filterEntries: []entryData{
				{value: "192.168.0.0/24", mustConsider: false},
			},
			expectedFiltered:    []string{},
			expectedIgnored:     []string{},
			expectedFilteredLen: 0,
			expectedIgnoredLen:  0,
		},
		{
			name:       "domain_complex_filtering_scenario",
			sourceType: constants.SourceTypeDomain,
			listType:   constants.ListTypeBlocklist,
			entries: []entryData{
				{value: "keep1.com", mustConsider: false},
				{value: "keep2.com", mustConsider: true},    // Must consider
				{value: "filter1.com", mustConsider: false}, // Will be filtered as must-filter
				{value: "filter2.com", mustConsider: true},  // Must consider but also must filter
				{value: "ignore1.com", mustConsider: false}, // Will be ignored
			},
			filterEntries: []entryData{
				{value: "filter1.com", mustConsider: true},  // Must filter
				{value: "filter2.com", mustConsider: true},  // Must filter but entry is must consider
				{value: "ignore1.com", mustConsider: false}, // Will be ignored
			},
			expectedFiltered: []string{
				"keep1.com",
				"keep2.com",
				"filter2.com",
			}, // filter2.com kept due to must consider
			expectedIgnored:     []string{"filter1.com", "ignore1.com"},
			expectedFilteredLen: 3,
			expectedIgnoredLen:  2,
		},
		{
			name:       "domain_blocklist_vs_allowlist_must_consider_true",
			sourceType: constants.SourceTypeDomain,
			listType:   constants.ListTypeBlocklist,
			entries: []entryData{
				{value: "example.com", mustConsider: true}, // Blocklist mustConsider
				{value: "other.com", mustConsider: false},
			},
			filterEntries: []entryData{
				{value: "example.com", mustConsider: true}, // Allowlist mustConsider
				{value: "someother.com", mustConsider: false},
			},
			expectedFiltered:    []string{"example.com", "other.com"},
			expectedIgnored:     []string{},
			expectedFilteredLen: 2,
			expectedIgnoredLen:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := multilog.NewLogger()
			cc := consolidators.NewCommonConsolidator(tt.sourceType, tt.listType)

			entrySet := utils.NewStringSet([]string{})
			for _, entry := range tt.entries {
				entrySet.AddWithConsider(entry.value, entry.mustConsider)
			}

			filterSet := utils.NewStringSet([]string{})
			for _, entry := range tt.filterEntries {
				filterSet.AddWithConsider(entry.value, entry.mustConsider)
			}

			filteredSet, ignoredSet := cc.FilterEntries(logger, entrySet, filterSet)

			assert.Equal(t, tt.expectedFilteredLen, len(filteredSet), "Filtered set size should match expected")
			assert.Equal(t, tt.expectedIgnoredLen, len(ignoredSet), "Ignored set size should match expected")

			for _, expectedEntry := range tt.expectedFiltered {
				assert.True(t, filteredSet.Contains(expectedEntry), "Filtered set should contain: %s", expectedEntry)
			}

			for _, expectedEntry := range tt.expectedIgnored {
				assert.True(t, ignoredSet.Contains(expectedEntry), "Ignored set should contain: %s", expectedEntry)
			}

			for entry := range filteredSet {
				assert.False(
					t,
					ignoredSet.Contains(entry),
					"Entry should not be in both filtered and ignored sets: %s",
					entry,
				)
			}
		})
	}
}

// TestCommonConsolidatorSaveEntries tests the SaveEntries method
func TestCommonConsolidatorSaveEntries(t *testing.T) {
	tests := []struct {
		name          string
		sourceType    string
		listType      string
		entries       []string
		expectedError bool
		validateFunc  func(t *testing.T, filePath string, entries []string)
	}{
		{
			name:       "domain_save_success",
			sourceType: constants.SourceTypeDomain,
			listType:   constants.ListTypeBlocklist,
			entries:    []string{"example.com", "test.com", "bad.domain.com"},
			validateFunc: func(t *testing.T, filePath string, entries []string) {
				content, err := os.ReadFile(filePath)
				require.NoError(t, err, "Should be able to read saved file")

				fileContent := string(content)
				for _, entry := range entries {
					assert.Contains(t, fileContent, entry, "File should contain entry: %s", entry)
				}

				// Check that entries are properly formatted (one per line)
				lines := len(entries)
				actualLines := len(strings.Split(strings.TrimSpace(fileContent), "\n"))
				assert.Equal(t, lines, actualLines, "File should have correct number of lines")
			},
		},
		{
			name:       "ipv4_save_success",
			sourceType: constants.SourceTypeIpv4,
			listType:   constants.ListTypeAllowlist,
			entries:    []string{"192.168.1.1", "10.0.0.1", "172.16.0.1", "8.8.8.8"},
			validateFunc: func(t *testing.T, filePath string, entries []string) {
				content, err := os.ReadFile(filePath)
				require.NoError(t, err, "Should be able to read saved file")

				fileContent := string(content)
				for _, entry := range entries {
					assert.Contains(t, fileContent, entry, "File should contain IP: %s", entry)
				}
			},
		},
		{
			name:       "ipv6_save_success",
			sourceType: constants.SourceTypeIpv6,
			listType:   constants.ListTypeBlocklist,
			entries:    []string{"2001:db8::1", "2001:db8::2", "fe80::1"},
			validateFunc: func(t *testing.T, filePath string, entries []string) {
				content, err := os.ReadFile(filePath)
				require.NoError(t, err, "Should be able to read saved file")

				fileContent := string(content)
				for _, entry := range entries {
					assert.Contains(t, fileContent, entry, "File should contain IPv6: %s", entry)
				}
			},
		},
		{
			name:       "cidr_save_success",
			sourceType: constants.SourceTypeCidrIpv4,
			listType:   constants.ListTypeBlocklist,
			entries:    []string{"192.168.0.0/24", "10.0.0.0/8", "172.16.0.0/12"},
			validateFunc: func(t *testing.T, filePath string, entries []string) {
				content, err := os.ReadFile(filePath)
				require.NoError(t, err, "Should be able to read saved file")

				fileContent := string(content)
				for _, entry := range entries {
					assert.Contains(t, fileContent, entry, "File should contain CIDR: %s", entry)
				}
			},
		},
		{
			name:       "empty_entries_save",
			sourceType: constants.SourceTypeDomain,
			listType:   constants.ListTypeBlocklist,
			entries:    []string{},
			validateFunc: func(t *testing.T, filePath string, entries []string) {
				// When there are no entries, the file should not be created
				_, err := os.Stat(filePath)
				assert.True(t, os.IsNotExist(err), "Empty entries should not create a file")
			},
		},
		{
			name:       "single_entry_save",
			sourceType: constants.SourceTypeDomain,
			listType:   constants.ListTypeAllowlist,
			entries:    []string{"single.com"},
			validateFunc: func(t *testing.T, filePath string, entries []string) {
				content, err := os.ReadFile(filePath)
				require.NoError(t, err, "Should be able to read saved file")

				fileContent := strings.TrimSpace(string(content))
				assert.Equal(t, "single.com", fileContent, "Single entry should be saved correctly")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := multilog.NewLogger()
			cc := consolidators.NewCommonConsolidator(tt.sourceType, tt.listType)

			entrySet := utils.NewStringSet(tt.entries)

			tempDir := t.TempDir()
			filePath := filepath.Join(tempDir, "output.txt")

			err := cc.SaveEntries(logger, entrySet, filePath)

			if tt.expectedError {
				assert.Error(t, err, "SaveEntries should return an error")
				return
			}

			assert.NoError(t, err, "SaveEntries should not return an error")

			if len(tt.entries) > 0 {
				assert.FileExists(t, filePath, "Output file should exist")
			}

			if tt.validateFunc != nil {
				tt.validateFunc(t, filePath, tt.entries)
			}
		})
	}
}

// TestCommonConsolidatorSaveEntriesErrorCases tests error scenarios for SaveEntries
func TestCommonConsolidatorSaveEntriesErrorCases(t *testing.T) {
	logger := multilog.NewLogger()
	cc := consolidators.NewCommonConsolidator(constants.SourceTypeDomain, constants.ListTypeBlocklist)

	entrySet := utils.NewStringSet([]string{"test.com"})

	if runtime.GOOS != "windows" { // Skip on Windows
		readOnlyDir := t.TempDir()
		require.NoError(t, os.Chmod(readOnlyDir, 0444)) // Read-only
		defer os.Chmod(readOnlyDir, 0755)               // Restore permissions for cleanup

		invalidPath := filepath.Join(readOnlyDir, "output.txt")
		_ = cc.SaveEntries(logger, entrySet, invalidPath)
	}
}

// TestCommonConsolidatorIsValid tests the IsValid method
func TestCommonConsolidatorIsValid(t *testing.T) {
	tests := []struct {
		name             string
		consolidatorType string
		consolidatorList string
		file             common.ProcessedFile
		expected         bool
	}{
		{
			name:             "valid_domain_blocklist",
			consolidatorType: constants.SourceTypeDomain,
			consolidatorList: constants.ListTypeBlocklist,
			file: common.ProcessedFile{
				GenericSourceType: constants.SourceTypeDomain,
				ListType:          constants.ListTypeBlocklist,
			},
			expected: true,
		},
		{
			name:             "valid_ipv4_allowlist",
			consolidatorType: constants.SourceTypeIpv4,
			consolidatorList: constants.ListTypeAllowlist,
			file: common.ProcessedFile{
				GenericSourceType: constants.SourceTypeIpv4,
				ListType:          constants.ListTypeAllowlist,
			},
			expected: true,
		},
		{
			name:             "invalid_source_type",
			consolidatorType: constants.SourceTypeDomain,
			consolidatorList: constants.ListTypeBlocklist,
			file: common.ProcessedFile{
				GenericSourceType: constants.SourceTypeIpv4, // Wrong source type
				ListType:          constants.ListTypeBlocklist,
			},
			expected: false,
		},
		{
			name:             "invalid_list_type",
			consolidatorType: constants.SourceTypeDomain,
			consolidatorList: constants.ListTypeBlocklist,
			file: common.ProcessedFile{
				GenericSourceType: constants.SourceTypeDomain,
				ListType:          constants.ListTypeAllowlist, // Wrong list type
			},
			expected: false,
		},
		{
			name:             "invalid_both_types",
			consolidatorType: constants.SourceTypeIpv6,
			consolidatorList: constants.ListTypeAllowlist,
			file: common.ProcessedFile{
				GenericSourceType: constants.SourceTypeDomain,  // Wrong source type
				ListType:          constants.ListTypeBlocklist, // Wrong list type
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := consolidators.NewCommonConsolidator(tt.consolidatorType, tt.consolidatorList)
			result := cc.IsValid(tt.file)
			assert.Equal(t, tt.expected, result, "IsValid should return expected result")
		})
	}
}

// Helper types for test data
type fileData struct {
	name         string
	content      string
	entries      int
	sourceType   string
	listType     string
	mustConsider bool
	shouldIgnore bool
}

type entryData struct {
	value        string
	mustConsider bool
}
