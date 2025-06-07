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
