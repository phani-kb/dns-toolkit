package cmd

import (
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	con "github.com/phani-kb/dns-toolkit/internal/consolidators"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"

	"github.com/stretchr/testify/assert"
)

func TestConsolidateGroupsCommand(t *testing.T) {

	assert.NotNil(t, consolidateGroupsCmd)
	assert.Equal(t, "groups", consolidateGroupsCmd.Use)
	assert.Contains(t, consolidateGroupsCmd.Short, "sized consolidated lists")
	assert.NotNil(t, consolidateGroupsCmd.Run)
}

func TestConsolidateByGroup(t *testing.T) {

	logger := multilog.NewLogger()

	tests := []struct {
		name              string
		genericSourceType string
		listType          string
		group             string
		entriesToIgnore   u.StringSet
		processedFiles    []c.ProcessedFile
		expectedCount     int
		expectedValid     bool
		expectedGroup     string
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
