package common

import (
	"encoding/json"
	"fmt"
	"strings"

	consts "github.com/phani-kb/dns-toolkit/internal/constants"
)

type Named interface {
	GetName() string
}

type SourceType struct {
	Name      string     `json:"name"`
	Disabled  bool       `json:"disabled,omitempty"`
	ListTypes []ListType `json:"list_types,omitempty"`
	Notes     string     `json:"notes,omitempty"`
}

func (st *SourceType) Validate() error {
	if st.Name == "" {
		return fmt.Errorf("name is required")
	}
	if !consts.ValidTypes[st.Name] {
		return fmt.Errorf("invalid type: %s", st.Name)
	}
	for _, lt := range st.ListTypes {
		if err := lt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (st *SourceType) GetListTypes() []ListType {
	listTypes := make([]ListType, 0)
	for _, lt := range st.ListTypes {
		if !lt.Disabled {
			listTypes = append(listTypes, lt)
		}
	}
	return listTypes
}

type ListType struct {
	Name         string   `json:"name"`
	Disabled     bool     `json:"disabled,omitempty"`
	Groups       []string `json:"groups,omitempty"`
	MustConsider bool     `json:"must_consider,omitempty"`
	Notes        string   `json:"notes,omitempty"`
}

func (lt *ListType) Validate() error {
	if lt.Name == "" {
		return fmt.Errorf("name is required")
	}
	if !consts.ValidListTypes[lt.Name] {
		return fmt.Errorf("invalid list type: %s", lt.Name)
	}
	for _, group := range lt.Groups {
		if !consts.ValidGroups[group] {
			return fmt.Errorf("invalid group: %s", group)
		}
	}
	return nil
}

// DownloadSummary represents information about a downloaded DNS blocklist file.
// It contains metadata about the source, content types, and download status.
type DownloadSummary struct {
	Name                  string       `json:"name"`                        // Name of the source
	URL                   string       `json:"url"`                         // URL where the list was downloaded from
	TypeCount             int          `json:"type_count"`                  // Number of entry types in the source
	Types                 []SourceType `json:"types"`                       // Array of entry types (ipv4, domain, etc.)
	Filepath              string       `json:"filepath"`                    // Path to the downloaded file
	CountToConsider       int          `json:"count_to_consider,omitempty"` // Number of entries to consider
	Frequency             string       `json:"frequency"`                   // Frequency of updates
	Categories            []string     `json:"categories,omitempty"`        // Categories this source belongs to
	Checksum              string       `json:"checksum"`                    // Checksum of the file content
	Error                 string       `json:"error"`                       // Error message if download failed
	LastDownloadTimestamp string       `json:"last_download_timestamp"`     // Timestamp of the last successful download
	LastCheckedTimestamp  string       `json:"last_checked_timestamp"`      // Timestamp when last checked for updates
}

// GetSourceTypes returns a list of source types that are not disabled.
func (ds *DownloadSummary) GetSourceTypes() []SourceType {
	sourceTypes := make([]SourceType, 0)
	for _, st := range ds.Types {
		if !st.Disabled {
			sourceTypes = append(sourceTypes, st)
		}
	}
	return sourceTypes
}

func (ds *DownloadSummary) GetName() string {
	return ds.Name
}

// UnmarshalJSON implements custom JSON unmarshalling for DownloadSummary to handle
// different formats of the Types field (could be a string or an array of strings).
func (ds *DownloadSummary) UnmarshalJSON(data []byte) error {
	type Alias DownloadSummary
	aux := &struct {
		Types      interface{} `json:"types"`
		Categories interface{} `json:"categories"`
		*Alias
	}{
		Alias: (*Alias)(ds),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Initialize Types as an empty slice if it's nil
	if ds.Types == nil {
		ds.Types = []SourceType{}
	}

	if ds.Categories == nil {
		ds.Categories = []string{}
	}

	if v, ok := aux.Categories.([]interface{}); ok {
		for _, item := range v {
			if category, ok := item.(string); ok {
				ds.Categories = append(ds.Categories, category)
			}
		}
	}

	if v, ok := aux.Types.([]interface{}); ok {
		for _, item := range v {
			if typeItem, ok := item.(map[string]interface{}); ok {
				sourceType := SourceType{}

				// Extract name
				if name, ok := typeItem["name"].(string); ok {
					sourceType.Name = name
				}

				// Extract disabled flag
				if disabled, ok := typeItem["disabled"].(bool); ok {
					sourceType.Disabled = disabled
				}

				// Extract list_types if present
				if listTypesRaw, ok := typeItem["list_types"].([]interface{}); ok {
					for _, ltRaw := range listTypesRaw {
						if ltMap, ok := ltRaw.(map[string]interface{}); ok {
							listType := ListType{}

							if name, ok := ltMap["name"].(string); ok {
								listType.Name = name
							}

							if disabled, ok := ltMap["disabled"].(bool); ok {
								listType.Disabled = disabled
							}

							if mustConsider, ok := ltMap["must_consider"].(bool); ok {
								listType.MustConsider = mustConsider
							}

							// Extract groups if present
							if groupsRaw, ok := ltMap["groups"].([]interface{}); ok {
								for _, groupRaw := range groupsRaw {
									if group, ok := groupRaw.(string); ok {
										listType.Groups = append(listType.Groups, group)
									}
								}
							}

							sourceType.ListTypes = append(sourceType.ListTypes, listType)
						}
					}
				}

				ds.Types = append(ds.Types, sourceType)
			}
		}
	} else if aux.Types != nil {
		return fmt.Errorf("types field must be an array of objects")
	}

	return nil
}

type DownloadFile struct {
	Name      string           `json:"name"`
	Folder    string           `json:"folder"`
	Filename  string           `json:"filename"`
	IsArchive bool             `json:"is_archive"`
	URL       string           `json:"url"`
	Frequency string           `json:"frequency"`
	Targets   []DownloadTarget `json:"targets"`
}

type DownloadTarget struct {
	SourceFolder string `json:"source_folder"`
	SourceFile   string `json:"source_file"`
	TargetFolder string `json:"target_folder"`
	TargetFile   string `json:"target_file"`
}

// ProcessedFile contains information about a file after it has been processed.
// It includes metadata about the content type and validation results.
type ProcessedFile struct {
	Name              string   `json:"name"`                    // Name of the source
	GenericSourceType string   `json:"generic_source_type"`     // Generic categorization of the content (domain, ipv4, etc.)
	ActualSourceType  string   `json:"actual_source_type"`      // Specific source type detected
	ListType          string   `json:"list_type"`               // Type of list (blocklist or allowlist)
	Filepath          string   `json:"filepath"`                // Path of the processed file
	NumberOfEntries   int      `json:"number_of_entries"`       // Count of entries in the file
	Checksum          string   `json:"checksum"`                // Checksum of the file content
	MustConsider      bool     `json:"must_consider,omitempty"` // Whether the file must be considered
	Valid             bool     `json:"valid"`                   // Whether the file contains valid entries
	Groups            []string `json:"groups,omitempty"`        // Size groups this file belongs to (mini, lite, normal, big)
	Categories        []string `json:"categories,omitempty"`    // Categories this file belongs to
}

// UnmarshalJSON implements custom JSON unmarshalling for ProcessedFile to handle the group array
func (pf *ProcessedFile) UnmarshalJSON(data []byte) error {
	type Alias ProcessedFile
	aux := &struct {
		Groups     interface{} `json:"groups"`
		Categories interface{} `json:"categories"`
		*Alias
	}{
		Alias: (*Alias)(pf),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Initialize Groups as an empty slice if it's nil
	if pf.Groups == nil {
		pf.Groups = []string{}
	}

	if pf.Categories == nil {
		pf.Categories = []string{}
	}

	// Handle groups if they exist
	if v, ok := aux.Groups.([]interface{}); ok {
		for _, item := range v {
			if group, ok := item.(string); ok {
				pf.Groups = append(pf.Groups, group)
			}
		}
	}

	// Handle categories if they exist
	if v, ok := aux.Categories.([]interface{}); ok {
		for _, item := range v {
			if category, ok := item.(string); ok {
				pf.Categories = append(pf.Categories, category)
			}
		}
	}

	return nil
}

// ProcessedSummary summarizes the results of processing a source file.
// It tracks both valid and invalid entries found during processing.
type ProcessedSummary struct {
	Name                   string          `json:"name"`                     // Source name
	Types                  []SourceType    `json:"types"`                    // Content types in the source
	ValidFiles             []ProcessedFile `json:"valid_files,omitempty"`    // Map of valid files by type
	InvalidFiles           []ProcessedFile `json:"invalid_files,omitempty"`  // Map of invalid files by type
	LastProcessedTimestamp string          `json:"last_processed_timestamp"` // When processing was completed
}

func (ps *ProcessedSummary) GetSourceTypes() []SourceType {
	sourceTypes := make([]SourceType, 0)
	for _, st := range ps.Types {
		if !st.Disabled {
			sourceTypes = append(sourceTypes, st)
		}
	}
	return sourceTypes
}

func (ps *ProcessedSummary) GetName() string {
	return ps.Name
}

// ConsolidatedSummary represents information about consolidated files that combine
// multiple source files of the same type into a single deduplicated file.
type ConsolidatedSummary struct {
	Type                      string   `json:"type"`                            // Type of entries (domain, ipv4, etc.)
	Filepath                  string   `json:"filepath"`                        // Path to the consolidated file
	ListType                  string   `json:"list_type"`                       // Type of list (blocklist or allowlist)
	Checksum                  string   `json:"checksum"`                        // Checksum of the consolidated file
	FilesCount                int      `json:"files_count"`                     // Number of source files consolidated
	Files                     []string `json:"files"`                           // List of source files that were consolidated
	Valid                     bool     `json:"valid"`                           // Whether this contains valid entries
	Count                     int      `json:"count"`                           // Number of entries in the file
	IgnoredEntriesCount       int      `json:"ignored_entries_count,omitempty"` // Number of entries ignored during consolidation
	IgnoredFilepath           string   `json:"ignored_filepath,omitempty"`      // Path to the ignored entries file
	LastConsolidatedTimestamp string   `json:"last_consolidated_timestamp"`     // When consolidation completed
	Group                     string   `json:"group,omitempty"`                 // Size group (mini, lite, normal, big)
	Category                  string   `json:"category,omitempty"`              // Category (advertising, malware, privacy, etc.)
}

// GetFilename generates the filename for a consolidated file based on its properties.
func (cs *ConsolidatedSummary) GetFilename() string {
	return cs.Type + "_" + cs.ListType + cs.getValidString() + ".txt"
}

// GetIgnoredFilename generates the filename for an ignored entries file based on its properties.
func (cs *ConsolidatedSummary) GetIgnoredFilename() string {
	return cs.Type + "_" + cs.ListType + cs.getValidString() + "_ignored.txt"
}

// getValidString returns "valid" or "invalid" string based on the Valid field.
func (cs *ConsolidatedSummary) getValidString() string {
	if cs.Valid {
		return ""
	}
	return "_invalid"
}

func (cs *ConsolidatedSummary) GetName() string {
	return cs.Type
}

// ConsolidatedGroupsSummary represents a group of consolidated summaries organized by size.
type ConsolidatedGroupsSummary struct {
	Group                     string                `json:"group"`                       // Size group (mini, lite, normal, big)
	ConsolidatedSummaries     []ConsolidatedSummary `json:"consolidated_summaries"`      // Consolidated summaries for this group
	LastConsolidatedTimestamp string                `json:"last_consolidated_timestamp"` // When consolidation completed
}

func (css *ConsolidatedGroupsSummary) GetName() string {
	return css.Group
}

// ConsolidatedCategoriesSummary represents consolidated summaries grouped by category.
type ConsolidatedCategoriesSummary struct {
	Category                  string                `json:"category"`                    // Category name (advertising, malware, privacy, etc)
	ConsolidatedSummaries     []ConsolidatedSummary `json:"consolidated_summaries"`      // Consolidated summaries for this category
	LastConsolidatedTimestamp string                `json:"last_consolidated_timestamp"` // When consolidation was completed
}

// ConsolidatedCategoriesSummaryLessFunc provides sorting logic for ConsolidatedCategoriesSummary
// objects by category name.
func ConsolidatedCategoriesSummaryLessFunc(i, j ConsolidatedCategoriesSummary) bool {
	return i.Category < j.Category
}

// OverlapDetailedSummary contains detailed information about the overlap analysis
// between different source files grouped by source type.
type OverlapDetailedSummary struct {
	SourceTypes      []OverlapSourceType `json:"source_types"`       // List of source types with overlap data
	SourceTypesCount int                 `json:"source_types_count"` // Number of source types analyzed
}

// OverlapSourceType represents overlap information for a specific source type.
type OverlapSourceType struct {
	Type       string        `json:"source_type"` // Source type (domain, ipv4, etc.)
	PairsCount int           `json:"pairs_count"` // Number of file pairs analyzed
	Pairs      []OverlapPair `json:"pairs"`       // List of file pairs with overlap data
}

// OverlapPair represents the overlap between two specific files.
type OverlapPair struct {
	Source   OverlapFileInfo `json:"source"`   // Source file information
	Target   OverlapFileInfo `json:"target"`   // Target file information
	Overlap  int             `json:"overlap"`  // Number of overlapping entries
	Filepath string          `json:"filepath"` // Path to the file containing the overlap entries
}

// OverlapFileInfo contains information about a file in an overlap analysis.
type OverlapFileInfo struct {
	Filename string `json:"filename"`  // Name of the file
	Name     string `json:"name"`      // Source name
	Type     string `json:"type"`      // Type of entries (domain, ipv4, etc.)
	ListType string `json:"list_type"` // Type of list (blocklist or allowlist)
	Count    int    `json:"count"`     // Number of entries in the file
	Percent  string `json:"percent"`   // Percentage of overlap
}

// OverlapSummary represents a compact summary of overlap information for a source.
type OverlapSummary struct {
	Source       string                  `json:"source"`        // Name of the source file
	ListType     string                  `json:"list_type"`     // Type of list (blocklist or allowlist)
	Type         string                  `json:"source_type"`   // Source type (domain, ipv4, etc.)
	Count        int                     `json:"count"`         // Number of entries in the source
	Unique       int                     `json:"unique"`        // Number of unique entries
	TargetsCount int                     `json:"targets_count"` // Number of target files with overlap
	TargetsList  []string                `json:"targets"`       // List of targets as strings
	Targets      []OverlapTargetFileInfo `json:"-"`             // Detailed target information (not serialized)
}

func (os *OverlapSummary) GetName() string {
	return os.Source
}

// OverlapTargetFileInfo contains information about a target file in a compact overlap summary.
type OverlapTargetFileInfo struct {
	Name     string `json:"name"`      // Name of the target
	ListType string `json:"list_type"` // Type of list (blocklist or allowlist)
	Type     string `json:"type"`      // Source type (domain, ipv4, etc.)
	Count    int    `json:"count"`     // Number of entries in the target
	Overlap  int    `json:"overlap"`   // Number of overlapping entries
	Percent  string `json:"percent"`   // Percentage of overlap
}

// GetString returns a formatted string representation of the overlap target file info.
func (ot *OverlapTargetFileInfo) GetString() string {
	return ot.Name + ", lt: " + ot.ListType + ", type: " + ot.Type + ", count: " + fmt.Sprint(
		ot.Count,
	) + ", overlap: " + fmt.Sprint(
		ot.Overlap,
	) + ", percent: " + ot.Percent
}

// FileInfo contains basic information about a file.
type FileInfo struct {
	Name         string `json:"name"`                    // Name of the file source
	Filepath     string `json:"filepath"`                // Path to the file
	MustConsider bool   `json:"must_consider,omitempty"` // Whether the file must be considered
}

// GetString returns a formatted string representation of file info.
func (fi *FileInfo) GetString() string {
	return fmt.Sprintf("%s [%s]%s", fi.Name, fi.Filepath,
		func() string {
			if fi.MustConsider {
				return " [must consider]"
			}
			return ""
		}())
}

// TemplateData contains data for dynamic templates used in generating output files.
type TemplateData struct {
	AppName        string
	AppVersion     string
	AppDescription string
	FileName       string
	LastUpdated    string
	Format         string
	Description    string
	Count          int
}

// TopSummary contains information about the top entries found across multiple sources.
type TopSummary struct {
	GenericSourceType string           `json:"generic_source_type"` // Type of entries (domain, ipv4, etc.)
	MinSources        int              `json:"min_sources"`         // Minimum number of sources to include entry
	TopEntries        []EntryCountPair `json:"-"`                   // List of top entries (not serialized)
	Count             int              `json:"count"`               // Number of entries
	Filepath          string           `json:"filepath"`            // Path to the file containing top entries
	ListType          string           `json:"list_type"`           // Type of list (blocklist or allowlist)
}

func (ts *TopSummary) GetName() string {
	return ts.GenericSourceType
}

// MarshalJSON implements custom JSON marshaling for TopSummary to include
// the count of TopEntries in the JSON output.
func (ts *TopSummary) MarshalJSON() ([]byte, error) {
	type Alias TopSummary
	return json.Marshal(&struct {
		Alias
		Count int `json:"count"`
	}{
		Alias: (Alias)(*ts),
		Count: len(ts.TopEntries),
	})
}

// EntryCountPair represents an entry and how many times it appears across sources.
type EntryCountPair struct {
	Entry string `json:"entry"` // The entry text (domain, IP, etc.)
	Count int    `json:"count"` // Number of sources containing this entry
}

// EntryHeap is a min-heap of EntryCountPair elements
type EntryHeap []EntryCountPair

func (h *EntryHeap) Len() int { return len(*h) }

func (h *EntryHeap) Less(i, j int) bool { return (*h)[i].Count < (*h)[j].Count } // Min-heap based on Count

func (h *EntryHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *EntryHeap) Push(x interface{}) {
	*h = append(*h, x.(EntryCountPair))
}

func (h *EntryHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// lessFuncByName is a generic helper function for sorting objects by name.
// It takes a function to extract the name from the object.
func lessFuncByName[T any](i, j T, getName func(T) string) bool {
	return strings.ToLower(getName(i)) < strings.ToLower(getName(j))
}

// DownloadSummaryLessFunc provides sorting logic for DownloadSummary objects by name.
func DownloadSummaryLessFunc(i, j DownloadSummary) bool {
	return lessFuncByName(i, j, func(ds DownloadSummary) string { return ds.Name })
}

// ProcessedSummaryLessFunc provides sorting logic for ProcessedSummary objects by name.
func ProcessedSummaryLessFunc(i, j ProcessedSummary) bool {
	return lessFuncByName(i, j, func(ps ProcessedSummary) string { return ps.Name })
}

// ConsolidatedSummaryLessFunc provides sorting logic for ConsolidatedSummary objects by type.
func ConsolidatedSummaryLessFunc(i, j ConsolidatedSummary) bool {
	return lessFuncByName(i, j, func(cs ConsolidatedSummary) string { return cs.Type })
}

// ConsolidatedGroupsSummaryLessFunc provides sorting logic for ConsolidatedGroupsSummary objects by group.
func ConsolidatedGroupsSummaryLessFunc(i, j ConsolidatedGroupsSummary) bool {
	// Define a map for ordering groups
	groupOrder := map[string]int{
		"mini":   0,
		"lite":   1,
		"normal": 2,
		"big":    3,
	}

	// If both groups are in the map, use the map order
	if orderI, okI := groupOrder[i.Group]; okI {
		if orderJ, okJ := groupOrder[j.Group]; okJ {
			return orderI < orderJ
		}
	}

	// Default to alphabetical ordering if not in map
	return i.Group < j.Group
}

// OverlapDetailedSummaryLessFunc provides sorting logic for OverlapDetailedSummary
// objects by the number of source types (descending).
func OverlapDetailedSummaryLessFunc(i, j OverlapDetailedSummary) bool {
	return len(i.SourceTypes) > len(j.SourceTypes)
}

// OverlapSummaryLessFunc provides sorting logic for OverlapSummary objects,
// first by source type and then by source name.
func OverlapSummaryLessFunc(i, j OverlapSummary) bool {
	// sort by source_type in ascending order
	if i.Type != j.Type {
		return i.Type < j.Type
	}

	// then by Source in ascending order
	if i.Source != j.Source {
		return i.Source < j.Source
	}

	return false
}

// TopSummaryLessFunc provides sorting logic for TopSummary objects,
// first by generic source type and then by min sources (descending).
func TopSummaryLessFunc(i, j TopSummary) bool {
	if i.GenericSourceType == j.GenericSourceType {
		return i.MinSources > j.MinSources
	}
	return lessFuncByName(i, j, func(ts TopSummary) string { return ts.GenericSourceType })
}

type ArchiveSummary struct {
	Folders      []ArchiveFolder      `json:"folders,omitempty"` // List of folders in the archive
	SummaryFiles []ArchiveSummaryFile `json:"files,omitempty"`   // List of files in the archive
}

func (as *ArchiveSummary) GetName() string {
	return "archive"
}

type ArchiveFolder struct {
	Name      string        `json:"name"`      // Name of the folder
	Count     int           `json:"count"`     // Number of sources containing this entry
	Files     []ArchiveFile `json:"files"`     // List of files in the folder
	Timestamp string        `json:"timestamp"` // Timestamp of the folder
}

type ArchiveFile struct {
	Name      string `json:"name"`      // Name of the file
	Filepath  string `json:"filepath"`  // Path to the file
	Count     int    `json:"count"`     // Number of lines in the file
	Checksum  string `json:"checksum"`  // Checksum of the file content
	Size      int64  `json:"size"`      // Size of the file in bytes
	Timestamp string `json:"timestamp"` // Timestamp of the file
}

type ArchiveSummaryFile struct {
	Name        string `json:"name"`         // Name of the file
	Filepath    string `json:"filepath"`     // Path to the file
	SummaryType string `json:"summary_type"` // Type of summary (download, processed, consolidated, etc.)
	Count       int    `json:"count"`        // Number of entries in the file of the specific summary type
	Checksum    string `json:"checksum"`     // Checksum of the file content
	Size        int64  `json:"size"`         // Size of the file in bytes
	Valid       bool   `json:"valid"`        // Whether the file contains valid entries
	Timestamp   string `json:"timestamp"`    // Timestamp of the file
}

func (asf *ArchiveSummaryFile) GetName() string {
	return asf.Name
}
