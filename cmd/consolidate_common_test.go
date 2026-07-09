package cmd

import (
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
)

func TestConsolidationParams(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		params ConsolidationParams
	}{
		{
			name: "create consolidation params for group",
			params: ConsolidationParams{
				GenericSourceType: constants.SourceTypeDomain,
				ListType:          constants.ListTypeBlocklist,
				Identifier:        "test-group",
				OutputDir:         "/tmp/output",
				IdentifierField:   "Group",
			},
		},
		{
			name: "create consolidation params for category",
			params: ConsolidationParams{
				GenericSourceType: constants.SourceTypeIpv4,
				ListType:          constants.ListTypeAllowlist,
				Identifier:        "test-category",
				OutputDir:         "/tmp/output",
				IdentifierField:   "Category",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotEmpty(t, tt.params.GenericSourceType, "GenericSourceType should not be empty")
			assert.NotEmpty(t, tt.params.ListType, "ListType should not be empty")
			assert.NotEmpty(t, tt.params.Identifier, "Identifier should not be empty")
			assert.NotEmpty(t, tt.params.OutputDir, "OutputDir should not be empty")
			assert.NotEmpty(t, tt.params.IdentifierField, "IdentifierField should not be empty")
		})
	}
}

func TestProcessingConfig(t *testing.T) {
	t.Parallel()

	mockGetFilesFunc := func(files []c.ProcessedFile, identifier string) []c.ProcessedFile {
		return files
	}

	mockConsolidateFunc := func(
		logger *multilog.Logger,
		gst, listType, identifier string,
		entriesToIgnore u.StringSet,
		processedFiles []c.ProcessedFile,
	) (u.StringSet, c.ConsolidatedSummary) {
		return u.NewStringSet([]string{}), c.ConsolidatedSummary{}
	}

	cfg := ProcessingConfig{
		GetFilesFunc:       mockGetFilesFunc,
		ConsolidateFunc:    mockConsolidateFunc,
		ProcessedFiles:     []c.ProcessedFile{},
		GenericSourceTypes: []string{constants.SourceTypeDomain},
		Identifier:         "test",
		IdentifierField:    "Group",
	}

	assert.NotNil(t, cfg.GetFilesFunc, "GetFilesFunc should not be nil")
	assert.NotNil(t, cfg.ConsolidateFunc, "ConsolidateFunc should not be nil")
	assert.NotNil(t, cfg.ProcessedFiles, "ProcessedFiles should not be nil")
	assert.NotEmpty(t, cfg.GenericSourceTypes, "GenericSourceTypes should not be empty")
	assert.NotEmpty(t, cfg.Identifier, "Identifier should not be empty")
	assert.NotEmpty(t, cfg.IdentifierField, "IdentifierField should not be empty")
}
