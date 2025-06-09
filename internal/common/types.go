package common

import (
	"fmt"
)

type Named interface {
	GetName() string
}

type SourceType struct {
	Name      string     `json:"name"`
	ListTypes []ListType `json:"list_types,omitempty"`
}

func (st *SourceType) Validate() error {
	if st.Name == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}

func (st *SourceType) GetListTypes() []ListType {
	listTypes := make([]ListType, 0)
	return listTypes
}

type ListType struct {
	Name string `json:"name"`
}

func (lt *ListType) Validate() error {
	if lt.Name == "" {
		return fmt.Errorf("name is required")
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

func (ds *DownloadSummary) GetName() string {
	return ds.Name
}

type DownloadFile struct {
	Name      string           `json:"name"`
	Folder    string           `json:"folder"`
	Filename  string           `json:"filename"`
	IsArchive bool             `json:"is_archive"`
	URL       string           `json:"url"`
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
	Name              string `json:"name"`                // Name of the source
	GenericSourceType string `json:"generic_source_type"` // Generic categorization of the content (domain, ipv4, etc.)
	ActualSourceType  string `json:"actual_source_type"`  // Specific source type detected
	ListType          string `json:"list_type"`           // Type of list (blocklist or allowlist)
	Filepath          string `json:"filepath"`            // Path of the processed file
	NumberOfEntries   int    `json:"number_of_entries"`   // Count of entries in the file
	Valid             bool   `json:"valid"`               // Whether the file contains valid entries
}
