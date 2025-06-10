package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/multilog"
)

type SourcesConfig struct {
	Sources []Source `json:"sources"`
}

func (sc *SourcesConfig) Validate() error {
	return sc.ValidateWithConfig(nil)
}

func (sc *SourcesConfig) ValidateWithConfig(appConfig *AppConfig) error {
	if len(sc.Sources) == 0 {
		return fmt.Errorf("at least one source is required")
	}

	return nil
}

type Source struct {
	Name      string         `json:"name"`
	URL       string         `json:"url"`
	Files     []string       `json:"files,omitempty"`
	TypeCount int            `json:"type_count"`
	Types     []c.SourceType `json:"types"`
	License   string         `json:"license,omitempty"`
	Disabled  bool           `json:"disabled,omitempty"`
	Website   string         `json:"website,omitempty"`
}

func (s *Source) Validate() error {
	return s.ValidateWithConfig(nil)
}

func (s *Source) ValidateWithConfig(appConfig *AppConfig) error {
	if s.Name == "" {
		return fmt.Errorf("name is required")
	}

	return nil
}

func LoadSourcesConfig(logger *multilog.Logger, filePath string) (SourcesConfig, error) {
	logger.Debugf("Loading sources config from %s", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return SourcesConfig{}, fmt.Errorf("error opening sources file: %w", err)
	}

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return SourcesConfig{}, fmt.Errorf("error reading sources file: %w", err)
	}

	var config SourcesConfig
	if err := json.Unmarshal(byteValue, &config); err != nil {
		return SourcesConfig{}, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return config, nil
}

// GetEnabledSources returns a slice of enabled sources.
func (sc *SourcesConfig) GetEnabledSources(filters SourceFilters) []Source {
	uniqueSources := make(map[string]Source)
	for _, source := range sc.Sources {
		if source.IsEnabled() {
			key := fmt.Sprintf("%s_%s", source.Name, source.Types[0].Name)
			uniqueSources[key] = source
		}
	}
	var enabledSources []Source
	for _, source := range uniqueSources {
		enabledSources = append(enabledSources, source)
	}
	sort.Slice(enabledSources, func(i, j int) bool {
		return strings.ToLower(enabledSources[i].Name) < strings.ToLower(enabledSources[j].Name)
	})
	return enabledSources
}

// GetSourceByName returns a source with the specified name.
func (sc *SourcesConfig) GetSourceByName(name string) (Source, bool) {
	for _, source := range sc.Sources {
		if source.Name == name {
			return source, true
		}
	}
	return Source{}, false
}

// GetDownloadFile returns a DownloadFile struct for the source.
func (s *Source) GetDownloadFile(_ *multilog.Logger, downloadDir string) (c.DownloadFile, error) {
	downloadFile := c.DownloadFile{
		Name:    s.Name,
		URL:     s.URL,
		Folder:  downloadDir,
		Targets: make([]c.DownloadTarget, 0, len(s.Files)), // Pre-allocate capacity
	}

	// Handle non-archive files
	if len(s.Files) != 0 {
		return downloadFile, fmt.Errorf("files are not required for url %s", s.URL)
	}

	downloadFile.Filename = s.Name + ".txt"
	downloadFile.Targets = append(downloadFile.Targets, c.DownloadTarget{
		SourceFolder: downloadFile.Folder,
		SourceFile:   downloadFile.Filename,
		TargetFile:   downloadFile.Filename,
		TargetFolder: downloadFile.Folder,
	})

	return downloadFile, nil
}

// IsEnabled returns true if the source is enabled.
func (s *Source) IsEnabled() bool {
	return !s.Disabled
}
