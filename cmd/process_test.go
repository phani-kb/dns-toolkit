package cmd

import (
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestMergeSummaries(t *testing.T) {
	t.Parallel()

	existing := &c.ProcessedSummary{
		Types: []c.SourceType{
			{
				Name: "domain",
			},
		},
		ValidFiles: []c.ProcessedFile{
			{
				Name:              "test1",
				GenericSourceType: "domain",
				ListType:          "blocklist",
				Filepath:          "/path/to/test1.txt",
				NumberOfEntries:   5,
				Valid:             true,
			},
		},
		InvalidFiles: []c.ProcessedFile{
			{
				Name:              "test1",
				GenericSourceType: "domain",
				ListType:          "blocklist",
				Filepath:          "/path/to/invalid1.txt",
				NumberOfEntries:   2,
				Valid:             false,
			},
		},
		LastProcessedTimestamp: "2023-01-01T00:00:00Z",
	}

	new := &c.ProcessedSummary{
		Types: []c.SourceType{
			{
				Name: "ipv4",
			},
		},
		ValidFiles: []c.ProcessedFile{
			{
				Name:              "test2",
				GenericSourceType: "ipv4",
				ListType:          "allowlist",
				Filepath:          "/path/to/test2.txt",
				NumberOfEntries:   8,
				Valid:             true,
			},
		},
		InvalidFiles: []c.ProcessedFile{
			{
				Name:              "test2",
				GenericSourceType: "ipv4",
				ListType:          "allowlist",
				Filepath:          "/path/to/invalid2.txt",
				NumberOfEntries:   3,
				Valid:             false,
			},
		},
		LastProcessedTimestamp: "2023-01-02T00:00:00Z",
	}

	mergeSummaries(existing, new)

	assert.Len(t, existing.Types, 2, "Should have 2 source types after merge")
	assert.Len(t, existing.ValidFiles, 2, "Should have 2 valid files after merge")
	assert.Len(t, existing.InvalidFiles, 2, "Should have 2 invalid files after merge")
	assert.Equal(t, "2023-01-02T00:00:00Z", existing.LastProcessedTimestamp, "Timestamp should be updated to latest")
}

func TestMergeSourceTypes(t *testing.T) {
	t.Parallel()

	existing := []c.SourceType{
		{
			Name: "domain",
		},
	}

	new := []c.SourceType{
		{
			Name: "ipv4",
		},
		{
			Name: "domain",
		},
	}

	result := mergeSourceTypes(existing, new)

	assert.Len(t, result, 2, "Should have 2 source types after merge")

	typeNames := make([]string, len(result))
	for i, st := range result {
		typeNames[i] = st.Name
	}
	assert.Contains(t, typeNames, "domain")
	assert.Contains(t, typeNames, "ipv4")
}

func TestMergeProcessedFiles(t *testing.T) {
	t.Parallel()

	existing := []c.ProcessedFile{
		{
			Name:              "test1",
			GenericSourceType: "domain",
			ListType:          "blocklist",
			Filepath:          "/path/to/test1.txt",
			NumberOfEntries:   5,
			Valid:             true,
		},
	}

	new := []c.ProcessedFile{
		{
			Name:              "test2",
			GenericSourceType: "ipv4",
			ListType:          "allowlist",
			Filepath:          "/path/to/test2.txt",
			NumberOfEntries:   8,
			Valid:             true,
		},
		{
			Name:              "test1",
			GenericSourceType: "domain",
			ListType:          "blocklist",
			Filepath:          "/path/to/test1_updated.txt",
			NumberOfEntries:   3,
			Valid:             false,
		},
	}

	result := mergeProcessedFiles(existing, new)

	assert.Len(t, result, 3, "Should have 3 processed files after merge")

	filepaths := make([]string, len(result))
	for i, pf := range result {
		filepaths[i] = pf.Filepath
	}
	assert.Contains(t, filepaths, "/path/to/test1.txt")
	assert.Contains(t, filepaths, "/path/to/test2.txt")
	assert.Contains(t, filepaths, "/path/to/test1_updated.txt")
}
