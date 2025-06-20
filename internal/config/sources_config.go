package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
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

	sourceNameTracker := make(map[string]bool)
	for _, source := range sc.Sources {
		if sourceNameTracker[source.Name] {
			return fmt.Errorf("duplicate source name: %s", source.Name)
		}
		err := source.ValidateWithConfig(appConfig)
		if err != nil {
			return fmt.Errorf("source validation error: %w", err)
		}
		sourceNameTracker[source.Name] = true
	}
	return nil
}

type Source struct {
	Name       string         `json:"name"`
	URL        string         `json:"url"`
	Files      []string       `json:"files,omitempty"`
	TypeCount  int            `json:"type_count"`
	Types      []c.SourceType `json:"types"`
	Frequency  string         `json:"frequency,omitempty"`
	Categories []string       `json:"categories,omitempty"`
	Countries  []string       `json:"countries,omitempty"`
	License    string         `json:"license,omitempty"`
	Disabled   bool           `json:"disabled,omitempty"`
	Website    string         `json:"website,omitempty"`
}

func (s *Source) Validate() error {
	return s.ValidateWithConfig(nil)
}

func (s *Source) ValidateWithConfig(appConfig *AppConfig) error {
	if s.Name == "" {
		return fmt.Errorf("name is required")
	}

	if s.URL == "" {
		return fmt.Errorf("url is required")
	}

	if len(s.Types) == 0 {
		return fmt.Errorf("at least one type is required")
	}
	typesTracker := make(map[string]bool)
	for _, t := range s.Types {
		if typesTracker[t.Name] {
			return fmt.Errorf("duplicate type: %s", t.Name)
		}
		err := t.Validate()
		if err != nil {
			return fmt.Errorf("type validation error: %w", err)
		}
		typesTracker[t.Name] = true
	}
	if s.Frequency != "" && !constants.ValidFrequencies[s.Frequency] {
		return fmt.Errorf("invalid frequency: %s", s.Frequency)
	}

	countryTracker := make(map[string]bool)
	for _, country := range s.Countries {
		if countryTracker[country] {
			return fmt.Errorf("duplicate country: %s", country)
		}
		if len(country) != 2 {
			return fmt.Errorf("invalid country code: %s", country)
		}
		countryTracker[country] = true
	}

	return nil
}

var defaultValues = map[string]string{
	"Frequency": constants.FrequencyDaily,
	"Group":     constants.GroupBig,
	"ListType":  constants.ListTypeBlocklist,
	"Category":  constants.CategoryOthers,
}

// UnmarshalJSON custom unmarshalling to handle nested list_types within types
func (s *Source) UnmarshalJSON(data []byte) error {
	type Alias Source
	aux := &struct {
		Types []struct {
			Name      string `json:"name"`
			ListTypes []struct {
				Name         string `json:"name"`
				Groups       string `json:"groups"`
				MustConsider bool   `json:"must_consider"`
				Disabled     bool   `json:"disabled"`
				Notes        string `json:"notes"`
			} `json:"list_types,omitempty"`
		} `json:"types"`
		Categories string `json:"categories"`
		Countries  string `json:"countries"`
		Files      string `json:"files"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Handle Types and ListTypes
	s.Types = make([]c.SourceType, len(aux.Types))
	for i, t := range aux.Types {
		listTypes := make([]c.ListType, len(t.ListTypes))
		for j, lt := range t.ListTypes {
			var groups []string
			if lt.Groups != "" {
				groups = strings.Split(lt.Groups, ",")
			}
			listType := c.ListType{
				Name:         lt.Name,
				Groups:       groups,
				Disabled:     lt.Disabled,
				MustConsider: lt.MustConsider,
				Notes:        lt.Notes,
			}

			listTypes[j] = listType
		}
		s.Types[i] = c.SourceType{
			Name:      t.Name,
			ListTypes: listTypes,
		}
	}
	s.TypeCount = len(s.Types)

	// Handle Categories
	if aux.Categories != "" {
		s.Categories = strings.Split(aux.Categories, ",")
		for i := range s.Categories {
			s.Categories[i] = strings.TrimSpace(s.Categories[i])
		}
	} else {
		s.Categories = []string{}
	}

	// Handle Countries
	if aux.Countries != "" {
		s.Countries = strings.Split(aux.Countries, ",")
		for i := range s.Countries {
			s.Countries[i] = strings.TrimSpace(s.Countries[i])
		}
	} else {
		s.Countries = []string{}
	}

	// Handle Files
	if aux.Files != "" {
		s.Files = strings.Split(aux.Files, ",")
		for i := range s.Files {
			s.Files[i] = strings.TrimSpace(s.Files[i])
		}
	} else {
		s.Files = []string{}
	}

	return nil
}

func LoadSourcesConfig(logger *multilog.Logger, filePath string) (SourcesConfig, error) {
	logger.Debugf("Loading sources config from %s", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return SourcesConfig{}, fmt.Errorf("error opening sources file: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			logger.Errorf("error closing file: %v", err)
		}
	}()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return SourcesConfig{}, fmt.Errorf("error reading sources file: %w", err)
	}

	var config SourcesConfig
	if err := json.Unmarshal(byteValue, &config); err != nil {
		return SourcesConfig{}, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	// Set default values for missing fields
	for i, source := range config.Sources {
		if source.Frequency == "" {
			config.Sources[i].Frequency = defaultValues["Frequency"]
		}
		for j, t := range source.Types {
			if len(t.ListTypes) == 0 {
				listType := c.ListType{
					Name:         defaultValues["ListType"],
					Groups:       []string{defaultValues["Group"]},
					MustConsider: false,
				}
				source.Types[j].ListTypes = append(t.ListTypes, listType)
			} else {
				for k, lt := range t.ListTypes {
					if len(lt.Groups) == 0 {
						source.Types[j].ListTypes[k].Groups = []string{defaultValues["Group"]}
					} else {
						// Convert group names to IDs and find the lowest (most restrictive) group ID
						lowestGroupId := len(constants.SizeGroups)
						for _, group := range lt.Groups {
							group = strings.TrimSpace(group)
							if groupId, ok := constants.GroupIdMap[group]; ok {
								if groupId < lowestGroupId {
									lowestGroupId = groupId
								}
							} else {
								return SourcesConfig{}, fmt.Errorf("invalid group: %s", group)
							}
						}

						// Include all groups with ID >= lowestGroupId (all less restrictive groups)
						outputGroups := make([]string, 0, len(constants.SizeGroups)-lowestGroupId)
						for i := lowestGroupId; i < len(constants.SizeGroups); i++ {
							outputGroups = append(outputGroups, constants.SizeGroups[i])
						}

						source.Types[j].ListTypes[k].Groups = outputGroups
					}
				}
			}
		}
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

// IsEnabled returns true if the source is enabled.
func (sc *SourcesConfig) IsEnabled(name string) bool {
	source, found := sc.GetSourceByName(name)
	if !found {
		return false
	}
	return source.IsEnabled()
}

// GetSourceByField returns a slice of sources with the specified field and value.
func (sc *SourcesConfig) GetSourceByField(field, value string) []Source {
	var sources []Source
	for _, source := range sc.Sources {
		switch field {
		case "listType":
			for _, t := range source.Types {
				for _, lt := range t.GetListTypes() {
					if lt.Name == value {
						sources = append(sources, source)
						break
					}
				}
			}
		case "type":
			for _, t := range source.Types {
				if t.Name == value {
					sources = append(sources, source)
					break
				}
			}
		case "frequency":
			if source.Frequency == value {
				sources = append(sources, source)
			}
		case "category":
			for _, category := range source.Categories {
				if category == value {
					sources = append(sources, source)
					break
				}
			}
		case "country":
			for _, country := range source.Countries {
				if country == value {
					sources = append(sources, source)
					break
				}
			}
		}
	}
	return sources
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
