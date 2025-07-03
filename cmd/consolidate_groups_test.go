package cmd

import (
	"os"
	"path/filepath"
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	con "github.com/phani-kb/dns-toolkit/internal/consolidators"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"

	"github.com/stretchr/testify/assert"
)

func createTempFileWithContentGroups(t *testing.T, content string) string {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	return tmpFile
}

func TestConsolidateGroupsCommand(t *testing.T) {

	assert.NotNil(t, consolidateGroupsCmd)
	assert.Equal(t, "groups", consolidateGroupsCmd.Use)
	assert.Contains(t, consolidateGroupsCmd.Short, "sized consolidated lists")
	assert.NotNil(t, consolidateGroupsCmd.Run)
}

func TestConsolidateByGroup(t *testing.T) {

	logger := multilog.NewLogger()

	tests := []struct {
		entriesToIgnore   u.StringSet
		name              string
		genericSourceType string
		listType          string
		group             string
		expectedGroup     string
		processedFiles    []c.ProcessedFile
		expectedCount     int
		expectedValid     bool
	}{
		{
			name:              "empty processed files",
			genericSourceType: constants.SourceTypeDomain,
			listType:          constants.ListTypeBlocklist,
			group:             constants.GroupMini,
			entriesToIgnore:   u.NewStringSet([]string{}),
			processedFiles:    []c.ProcessedFile{},
			expectedCount:     0,
			expectedValid:     false,
			expectedGroup:     "",
		},
		{
			name:              "invalid generic source type",
			genericSourceType: "invalid_type",
			listType:          constants.ListTypeBlocklist,
			group:             constants.GroupMini,
			entriesToIgnore:   u.NewStringSet([]string{}),
			processedFiles:    []c.ProcessedFile{{Valid: true}},
			expectedCount:     0,
			expectedValid:     false,
			expectedGroup:     "",
		},
		{
			name:              "valid group consolidation",
			genericSourceType: constants.SourceTypeDomain,
			listType:          constants.ListTypeBlocklist,
			group:             constants.GroupMini,
			entriesToIgnore:   u.NewStringSet([]string{}),
			processedFiles: []c.ProcessedFile{
				{
					Valid:             true,
					GenericSourceType: constants.SourceTypeDomain,
					ListType:          constants.ListTypeBlocklist,
					Groups:            []string{constants.GroupMini},
					NumberOfEntries:   5,
					Filepath:          "/tmp/test1.txt",
				},
			},
			expectedCount: 0,
			expectedValid: false,
			expectedGroup: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test registry to isolate from global state
			testRegistry := con.NewConsolidatorRegistry()
			origRegistry := con.Consolidators
			con.Consolidators = testRegistry
			defer func() { con.Consolidators = origRegistry }()

			entries, summary := consolidateByGroup(
				logger,
				tt.genericSourceType,
				tt.listType,
				tt.group,
				tt.entriesToIgnore,
				tt.processedFiles,
			)

			assert.Equal(t, tt.expectedCount, len(entries.ToSlice()))
			assert.Equal(t, tt.expectedValid, summary.Valid)
			assert.Equal(t, tt.expectedGroup, summary.Group)
		})
	}
}

func TestGetFilesForGroup(t *testing.T) {
	tests := []struct {
		name           string
		group          string
		processedFiles []c.ProcessedFile
		expected       int
	}{
		{
			name:           "empty input",
			processedFiles: []c.ProcessedFile{},
			group:          constants.GroupMini,
			expected:       0,
		},
		{
			name: "no matching group",
			processedFiles: []c.ProcessedFile{
				{Groups: []string{constants.GroupBig}, Valid: true},
				{Groups: []string{constants.GroupLite}, Valid: true},
			},
			group:    constants.GroupMini,
			expected: 0,
		},
		{
			name: "matching group",
			processedFiles: []c.ProcessedFile{
				{Groups: []string{constants.GroupMini}, Valid: true},
				{Groups: []string{constants.GroupBig}, Valid: true},
				{Groups: []string{constants.GroupMini, constants.GroupLite}, Valid: true},
			},
			group:    constants.GroupMini,
			expected: 2,
		},
		{
			name: "invalid files filtered out",
			processedFiles: []c.ProcessedFile{
				{Groups: []string{constants.GroupMini}, Valid: true},
				{Groups: []string{constants.GroupMini}, Valid: false}, // Invalid file
				{Groups: []string{constants.GroupBig}, Valid: true},
			},
			group:    constants.GroupMini,
			expected: 1,
		},
		{
			name: "multiple groups per file",
			processedFiles: []c.ProcessedFile{
				{Groups: []string{constants.GroupMini, constants.GroupBig, constants.GroupLite}, Valid: true},
				{Groups: []string{constants.GroupBig, constants.GroupMini}, Valid: true},
				{Groups: []string{constants.GroupLite}, Valid: true},
			},
			group:    constants.GroupMini,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getFilesForGroup(tt.processedFiles, tt.group)
			assert.Equal(t, tt.expected, len(result))

			for _, file := range result {
				assert.True(t, file.Valid)
				assert.Contains(t, file.Groups, tt.group)
			}
		})
	}
}

func TestProcessGroupConsolidation(t *testing.T) {
	logger := multilog.NewLogger()

	con.InitForTesting()

	tests := []struct {
		name               string
		group              string
		processedFiles     []c.ProcessedFile
		genericSourceTypes []string
		expectedResults    int
	}{
		{
			name:               "empty input",
			group:              constants.GroupMini,
			processedFiles:     []c.ProcessedFile{},
			genericSourceTypes: []string{constants.SourceTypeDomain},
			expectedResults:    0,
		},
		{
			name:  "no matching group",
			group: constants.GroupMini,
			processedFiles: []c.ProcessedFile{
				{Groups: []string{constants.GroupBig}, Valid: true},
			},
			genericSourceTypes: []string{constants.SourceTypeDomain},
			expectedResults:    0,
		},
		{
			name:  "matching files",
			group: constants.GroupMini,
			processedFiles: []c.ProcessedFile{
				{
					Groups:            []string{constants.GroupMini},
					Valid:             true,
					GenericSourceType: constants.SourceTypeDomain,
					ListType:          constants.ListTypeBlocklist,
					Filepath:          createTempFileWithContentGroups(t, "example.com\ntest.com\n"),
				},
			},
			genericSourceTypes: []string{constants.SourceTypeDomain},
			expectedResults:    0, // TODO
		},
		{
			name:  "multiple source types",
			group: constants.GroupMini,
			processedFiles: []c.ProcessedFile{
				{
					Groups:            []string{constants.GroupMini},
					Valid:             true,
					GenericSourceType: constants.SourceTypeDomain,
					ListType:          constants.ListTypeBlocklist,
					Filepath:          createTempFileWithContentGroups(t, "example.com\ntest.com\n"),
				},
				{
					Groups:            []string{constants.GroupMini},
					Valid:             true,
					GenericSourceType: constants.SourceTypeIpv4,
					ListType:          constants.ListTypeBlocklist,
					Filepath:          createTempFileWithContentGroups(t, "192.168.1.1\n10.0.0.1\n"),
				},
			},
			genericSourceTypes: []string{constants.SourceTypeDomain, constants.SourceTypeIpv4},
			expectedResults:    0, // TODO
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := processGroupConsolidation(
				logger,
				tt.group,
				tt.processedFiles,
				tt.genericSourceTypes,
			)

			assert.Equal(t, tt.expectedResults, len(result))

			// Verify structure of results
			for sourceType, summaries := range result {
				assert.Contains(t, tt.genericSourceTypes, sourceType)
				assert.IsType(t, []c.ConsolidatedSummary{}, summaries)
			}
		})
	}
}
