package common

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSourceType_Validate(t *testing.T) {
	tests := []struct {
		name        string
		sourceType  SourceType
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid source type",
			sourceType: SourceType{
				Name: "domain",
				ListTypes: []ListType{
					{Name: "blocklist", Groups: []string{"mini"}},
				},
			},
			expectError: false,
		},
		{
			name: "empty name",
			sourceType: SourceType{
				Name: "",
			},
			expectError: true,
			errorMsg:    "name is required",
		},
		{
			name: "invalid type",
			sourceType: SourceType{
				Name: "invalid_type",
			},
			expectError: true,
			errorMsg:    "invalid type: invalid_type",
		},
		{
			name: "invalid list type",
			sourceType: SourceType{
				Name: "domain",
				ListTypes: []ListType{
					{Name: "invalid_list_type"},
				},
			},
			expectError: true,
			errorMsg:    "invalid list type: invalid_list_type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.sourceType.Validate()
			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSourceType_GetListTypes(t *testing.T) {
	sourceType := SourceType{
		Name: "domain",
		ListTypes: []ListType{
			{Name: "blocklist", Disabled: false},
			{Name: "allowlist", Disabled: true},
			{Name: "blocklist", Disabled: false},
		},
	}

	listTypes := sourceType.GetListTypes()
	assert.Len(t, listTypes, 2)
	for _, lt := range listTypes {
		assert.False(t, lt.Disabled)
	}
}

func TestListType_Validate(t *testing.T) {
	tests := []struct {
		name        string
		listType    ListType
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid list type",
			listType: ListType{
				Name:   "blocklist",
				Groups: []string{"mini", "lite"},
			},
			expectError: false,
		},
		{
			name: "empty name",
			listType: ListType{
				Name: "",
			},
			expectError: true,
			errorMsg:    "name is required",
		},
		{
			name: "invalid list type name",
			listType: ListType{
				Name: "invalid_list",
			},
			expectError: true,
			errorMsg:    "invalid list type: invalid_list",
		},
		{
			name: "invalid group",
			listType: ListType{
				Name:   "blocklist",
				Groups: []string{"invalid_group"},
			},
			expectError: true,
			errorMsg:    "invalid group: invalid_group",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.listType.Validate()
			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDownloadSummary_GetSourceTypes(t *testing.T) {
	ds := DownloadSummary{
		Types: []SourceType{
			{Name: "domain", Disabled: false},
			{Name: "ipv4", Disabled: true},
			{Name: "ipv6", Disabled: false},
		},
	}

	sourceTypes := ds.GetSourceTypes()
	assert.Len(t, sourceTypes, 2)
	for _, st := range sourceTypes {
		assert.False(t, st.Disabled)
	}
}

func TestDownloadSummary_GetName(t *testing.T) {
	ds := DownloadSummary{Name: "test_source"}
	assert.Equal(t, "test_source", ds.GetName())
}

func TestDownloadSummary_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		expected DownloadSummary
		hasError bool
	}{
		{
			name: "valid JSON with types array",
			jsonData: `{
				"name": "test",
				"url": "http://example.com",
				"types": [
					{
						"name": "domain",
						"disabled": false,
						"list_types": [
							{
								"name": "blocklist",
								"disabled": false,
								"groups": ["mini", "lite"]
							}
						]
					}
				],
				"categories": ["malware", "ads"]
			}`,
			expected: DownloadSummary{
				Name: "test",
				URL:  "http://example.com",
				Types: []SourceType{
					{
						Name:     "domain",
						Disabled: false,
						ListTypes: []ListType{
							{
								Name:     "blocklist",
								Disabled: false,
								Groups:   []string{"mini", "lite"},
							},
						},
					},
				},
				Categories: []string{"malware", "ads"},
			},
		},
		{
			name: "empty types",
			jsonData: `{
				"name": "test",
				"url": "http://example.com",
				"types": []
			}`,
			expected: DownloadSummary{
				Name:       "test",
				URL:        "http://example.com",
				Types:      []SourceType{},
				Categories: []string{},
			},
		},
		{
			name: "invalid types field",
			jsonData: `{
				"name": "test",
				"types": "invalid"
			}`,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ds DownloadSummary
			err := json.Unmarshal([]byte(tt.jsonData), &ds)

			if tt.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.Name, ds.Name)
				assert.Equal(t, tt.expected.URL, ds.URL)
				assert.Equal(t, len(tt.expected.Types), len(ds.Types))
				if len(tt.expected.Types) > 0 {
					assert.Equal(t, tt.expected.Types[0].Name, ds.Types[0].Name)
				}
				assert.Equal(t, tt.expected.Categories, ds.Categories)
			}
		})
	}
}

func TestProcessedFile_UnmarshalJSON(t *testing.T) {
	jsonData := `{
		"name": "test",
		"generic_source_type": "domain",
		"groups": ["mini", "lite"],
		"categories": ["malware", "ads"]
	}`

	var pf ProcessedFile
	err := json.Unmarshal([]byte(jsonData), &pf)

	require.NoError(t, err)
	assert.Equal(t, "test", pf.Name)
	assert.Equal(t, "domain", pf.GenericSourceType)
	assert.Equal(t, []string{"mini", "lite"}, pf.Groups)
	assert.Equal(t, []string{"malware", "ads"}, pf.Categories)
}

func TestProcessedSummary_GetSourceTypes(t *testing.T) {
	ps := ProcessedSummary{
		Types: []SourceType{
			{Name: "domain", Disabled: false},
			{Name: "ipv4", Disabled: true},
		},
	}

	sourceTypes := ps.GetSourceTypes()
	assert.Len(t, sourceTypes, 1)
	assert.Equal(t, "domain", sourceTypes[0].Name)
}

func TestProcessedSummary_GetName(t *testing.T) {
	ps := ProcessedSummary{Name: "test_source"}
	assert.Equal(t, "test_source", ps.GetName())
}

func TestConsolidatedSummary_GetFilename(t *testing.T) {
	tests := []struct {
		name     string
		cs       ConsolidatedSummary
		expected string
	}{
		{
			name: "valid consolidated summary",
			cs: ConsolidatedSummary{
				Type:     "domain",
				ListType: "blocklist",
				Valid:    true,
			},
			expected: "domain_blocklist.txt",
		},
		{
			name: "invalid consolidated summary",
			cs: ConsolidatedSummary{
				Type:     "ipv4",
				ListType: "allowlist",
				Valid:    false,
			},
			expected: "ipv4_allowlist_invalid.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.cs.GetFilename())
		})
	}
}

func TestConsolidatedSummary_GetIgnoredFilename(t *testing.T) {
	cs := ConsolidatedSummary{
		Type:     "domain",
		ListType: "blocklist",
		Valid:    true,
	}

	expected := "domain_blocklist_ignored.txt"
	assert.Equal(t, expected, cs.GetIgnoredFilename())
}

func TestConsolidatedSummary_GetName(t *testing.T) {
	cs := ConsolidatedSummary{Type: "domain"}
	assert.Equal(t, "domain", cs.GetName())
}

func TestConsolidatedGroupsSummary_GetName(t *testing.T) {
	cgs := ConsolidatedGroupsSummary{Group: "mini"}
	assert.Equal(t, "mini", cgs.GetName())
}

func TestConsolidatedCategoriesSummaryLessFunc(t *testing.T) {
	cs1 := ConsolidatedCategoriesSummary{Category: "ads"}
	cs2 := ConsolidatedCategoriesSummary{Category: "malware"}

	assert.True(t, ConsolidatedCategoriesSummaryLessFunc(cs1, cs2))
	assert.False(t, ConsolidatedCategoriesSummaryLessFunc(cs2, cs1))
}

func TestOverlapSummary_GetName(t *testing.T) {
	os := OverlapSummary{Source: "test_source"}
	assert.Equal(t, "test_source", os.GetName())
}

func TestOverlapTargetFileInfo_GetString(t *testing.T) {
	ot := OverlapTargetFileInfo{
		Name:     "target",
		ListType: "blocklist",
		Type:     "domain",
		Count:    100,
		Overlap:  50,
		Percent:  "50%",
	}

	expected := "target, lt: blocklist, type: domain, count: 100, overlap: 50, percent: 50%"
	assert.Equal(t, expected, ot.GetString())
}

func TestFileInfo_GetString(t *testing.T) {
	tests := []struct {
		name     string
		fi       FileInfo
		expected string
	}{
		{
			name: "without must consider",
			fi: FileInfo{
				Name:         "test",
				Filepath:     "/path/to/file",
				MustConsider: false,
			},
			expected: "test [/path/to/file]",
		},
		{
			name: "with must consider",
			fi: FileInfo{
				Name:         "test",
				Filepath:     "/path/to/file",
				MustConsider: true,
			},
			expected: "test [/path/to/file] [must consider]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.fi.GetString())
		})
	}
}

func TestTopSummary_GetName(t *testing.T) {
	ts := TopSummary{GenericSourceType: "domain"}
	assert.Equal(t, "domain", ts.GetName())
}

func TestTopSummary_MarshalJSON(t *testing.T) {
	ts := TopSummary{
		GenericSourceType: "domain",
		TopEntries: []EntryCountPair{
			{Entry: "example.com", Count: 5},
			{Entry: "test.com", Count: 3},
		},
		Count:    100,
		Filepath: "/path/to/file",
		ListType: "blocklist",
	}

	data, err := json.Marshal(ts)
	require.NoError(t, err)

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	require.NoError(t, err)

	assert.Equal(t, "domain", result["generic_source_type"])
	assert.Equal(t, "/path/to/file", result["filepath"])
	assert.Equal(t, "blocklist", result["list_type"])

	countValue, exists := result["count"]
	assert.True(t, exists, "count field should exist")

	switch v := countValue.(type) {
	case float64:
		assert.True(t, v >= 0, "count should be non-negative")
	case int:
		assert.True(t, v >= 0, "count should be non-negative")
	default:
		t.Errorf("count field should be numeric, got %T", v)
	}

	_, topEntriesExists := result["TopEntries"]
	assert.False(t, topEntriesExists, "TopEntries should not be in JSON output")
}

func TestEntryHeap(t *testing.T) {
	h := &EntryHeap{}

	h.Push(EntryCountPair{Entry: "test1", Count: 5})
	h.Push(EntryCountPair{Entry: "test2", Count: 3})
	h.Push(EntryCountPair{Entry: "test3", Count: 7})

	assert.Equal(t, 3, h.Len())

	if h.Len() >= 2 {
		entry1 := (*h)[0]
		entry2 := (*h)[1]
		result := h.Less(0, 1)
		expected := entry1.Count < entry2.Count
		assert.Equal(t, expected, result)
	}

	if h.Len() >= 2 {
		elem0 := (*h)[0]
		elem1 := (*h)[1]
		h.Swap(0, 1)
		assert.Equal(t, elem0, (*h)[1])
		assert.Equal(t, elem1, (*h)[0])
		h.Swap(0, 1)
	}

	initialLen := h.Len()
	popped := h.Pop().(EntryCountPair)
	assert.Equal(t, initialLen-1, h.Len())
	assert.NotEmpty(t, popped.Entry)
	assert.Greater(t, popped.Count, 0)
}

func TestSortingFunctions(t *testing.T) {
	t.Run("DownloadSummaryLessFunc", func(t *testing.T) {
		ds1 := DownloadSummary{Name: "abc"}
		ds2 := DownloadSummary{Name: "def"}

		assert.True(t, DownloadSummaryLessFunc(ds1, ds2))
		assert.False(t, DownloadSummaryLessFunc(ds2, ds1))
	})

	t.Run("ProcessedSummaryLessFunc", func(t *testing.T) {
		ps1 := ProcessedSummary{Name: "abc"}
		ps2 := ProcessedSummary{Name: "def"}

		assert.True(t, ProcessedSummaryLessFunc(ps1, ps2))
		assert.False(t, ProcessedSummaryLessFunc(ps2, ps1))
	})

	t.Run("ConsolidatedSummaryLessFunc", func(t *testing.T) {
		cs1 := ConsolidatedSummary{Type: "domain"}
		cs2 := ConsolidatedSummary{Type: "ipv4"}

		assert.True(t, ConsolidatedSummaryLessFunc(cs1, cs2))
		assert.False(t, ConsolidatedSummaryLessFunc(cs2, cs1))
	})

	t.Run("ConsolidatedGroupsSummaryLessFunc", func(t *testing.T) {
		cgs1 := ConsolidatedGroupsSummary{Group: "mini"}
		cgs2 := ConsolidatedGroupsSummary{Group: "lite"}
		cgs3 := ConsolidatedGroupsSummary{Group: "unknown"}

		assert.True(t, ConsolidatedGroupsSummaryLessFunc(cgs1, cgs2))
		assert.False(t, ConsolidatedGroupsSummaryLessFunc(cgs2, cgs1))
		assert.True(t, ConsolidatedGroupsSummaryLessFunc(cgs1, cgs3))
	})

	t.Run("OverlapDetailedSummaryLessFunc", func(t *testing.T) {
		ods1 := OverlapDetailedSummary{SourceTypes: make([]OverlapSourceType, 5)}
		ods2 := OverlapDetailedSummary{SourceTypes: make([]OverlapSourceType, 3)}

		assert.True(t, OverlapDetailedSummaryLessFunc(ods1, ods2))
		assert.False(t, OverlapDetailedSummaryLessFunc(ods2, ods1))
	})

	t.Run("OverlapSummaryLessFunc", func(t *testing.T) {
		os1 := OverlapSummary{Type: "domain", Source: "source1"}
		os2 := OverlapSummary{Type: "domain", Source: "source2"}
		os3 := OverlapSummary{Type: "ipv4", Source: "source1"}

		assert.True(t, OverlapSummaryLessFunc(os1, os2))
		assert.True(t, OverlapSummaryLessFunc(os1, os3))
		assert.False(t, OverlapSummaryLessFunc(os3, os1))
	})

	t.Run("TopSummaryLessFunc", func(t *testing.T) {
		ts1 := TopSummary{GenericSourceType: "domain", MinSources: 5}
		ts2 := TopSummary{GenericSourceType: "domain", MinSources: 3}
		ts3 := TopSummary{GenericSourceType: "ipv4", MinSources: 5}

		assert.True(t, TopSummaryLessFunc(ts1, ts2))
		assert.True(t, TopSummaryLessFunc(ts1, ts3))
	})
}

func TestArchiveSummary_GetName(t *testing.T) {
	as := ArchiveSummary{}
	assert.Equal(t, "archive", as.GetName())
}

func TestArchiveSummaryFile_GetName(t *testing.T) {
	asf := ArchiveSummaryFile{Name: "test_file"}
	assert.Equal(t, "test_file", asf.GetName())
}

func TestNamed_Interface(t *testing.T) {
	var named Named

	ds := DownloadSummary{Name: "test"}
	named = &ds
	assert.Equal(t, "test", named.GetName())

	ps := ProcessedSummary{Name: "test"}
	named = &ps
	assert.Equal(t, "test", named.GetName())

	cs := ConsolidatedSummary{Type: "test"}
	named = &cs
	assert.Equal(t, "test", named.GetName())

	cgs := ConsolidatedGroupsSummary{Group: "test"}
	named = &cgs
	assert.Equal(t, "test", named.GetName())

	os := OverlapSummary{Source: "test"}
	named = &os
	assert.Equal(t, "test", named.GetName())

	ts := TopSummary{GenericSourceType: "test"}
	named = &ts
	assert.Equal(t, "test", named.GetName())

	as := ArchiveSummary{}
	named = &as
	assert.Equal(t, "archive", named.GetName())

	asf := ArchiveSummaryFile{Name: "test"}
	named = &asf
	assert.Equal(t, "test", named.GetName())
}

func TestComplexScenarios(t *testing.T) {
	t.Run("SourceType with multiple ListTypes", func(t *testing.T) {
		st := SourceType{
			Name: "domain",
			ListTypes: []ListType{
				{Name: "blocklist", Groups: []string{"mini", "lite"}, Disabled: false},
				{Name: "allowlist", Groups: []string{"normal"}, Disabled: true},
				{Name: "blocklist", Groups: []string{"big"}, Disabled: false},
			},
		}

		err := st.Validate()
		assert.NoError(t, err)

		enabledTypes := st.GetListTypes()
		assert.Len(t, enabledTypes, 2)
	})

	t.Run("DownloadSummary JSON edge cases", func(t *testing.T) {
		jsonData := `{"name": "test", "types": null}`
		var ds DownloadSummary
		err := json.Unmarshal([]byte(jsonData), &ds)
		assert.NoError(t, err)
		assert.Empty(t, ds.Types)

		jsonData = `{"name": "test", "categories": null}`
		err = json.Unmarshal([]byte(jsonData), &ds)
		assert.NoError(t, err)
		assert.Empty(t, ds.Categories)
	})

	t.Run("ProcessedFile JSON edge cases", func(t *testing.T) {
		jsonData := `{"name": "test", "groups": null, "categories": null}`
		var pf ProcessedFile
		err := json.Unmarshal([]byte(jsonData), &pf)
		assert.NoError(t, err)
		assert.Empty(t, pf.Groups)
		assert.Empty(t, pf.Categories)
	})
}

func TestDownloadSummaryToJSON(t *testing.T) {
	tests := []struct {
		name     string
		summary  DownloadSummary
		expected string
	}{
		{
			name: "valid summary",
			summary: DownloadSummary{
				Name: "test-source",
				URL:  "https://example.com",
				Types: []SourceType{
					{Name: "domain"},
				},
			},
			expected: `{"name":"test-source","url":"https://example.com","filepath":"","frequency":"","checksum":"","error":"","last_download_timestamp":"","last_checked_timestamp":"","types":[{"name":"domain"}],"type_count":0}`,
		},
		{
			name:     "empty summary",
			summary:  DownloadSummary{},
			expected: `{"name":"","url":"","filepath":"","frequency":"","checksum":"","error":"","last_download_timestamp":"","last_checked_timestamp":"","types":null,"type_count":0}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.summary.ToJSON()
			assert.JSONEq(t, tt.expected, result)
		})
	}
}

func TestTopSummaryMarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		summary TopSummary
		wantErr bool
	}{
		{
			name: "valid top summary",
			summary: TopSummary{
				GenericSourceType: "domain",
				ListType:          "blocklist",
				MinSources:        3,
				Count:             1000,
				Filepath:          "/test/path.txt",
				TopEntries:        []EntryCountPair{{Entry: "example.com", Count: 5}},
			},
			wantErr: false,
		},
		{
			name:    "empty top summary",
			summary: TopSummary{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.summary.MarshalJSON()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, data)

				var unmarshaled TopSummary
				err = json.Unmarshal(data, &unmarshaled)
				assert.NoError(t, err)
				assert.Equal(t, tt.summary.GenericSourceType, unmarshaled.GenericSourceType)
				assert.Equal(t, tt.summary.ListType, unmarshaled.ListType)
				assert.Equal(t, tt.summary.MinSources, unmarshaled.MinSources)
				assert.Equal(t, len(tt.summary.TopEntries), unmarshaled.Count)
				assert.Equal(t, tt.summary.Filepath, unmarshaled.Filepath)
			}
		})
	}
}
