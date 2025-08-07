package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
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
	Name                        string         `json:"name"`
	URL                         string         `json:"url"`
	Frequency                   string         `json:"frequency,omitempty"`
	License                     string         `json:"license,omitempty"`
	Website                     string         `json:"website,omitempty"`
	Notes                       string         `json:"notes,omitempty"`
	Types                       []c.SourceType `json:"types"`
	Files                       []string       `json:"files,omitempty"`
	Categories                  []string       `json:"categories,omitempty"`
	Countries                   []string       `json:"countries,omitempty"`
	TypeCount                   int            `json:"type_count"`
	CountToConsider             int            `json:"count_to_consider,omitempty"`
	Disabled                    bool           `json:"disabled,omitempty"`
	SkipGeneralConsolidation    bool           `json:"skip_general_consolidation,omitempty"`
	SkipGroupsConsolidation     bool           `json:"skip_groups_consolidation,omitempty"`
	SkipCategoriesConsolidation bool           `json:"skip_categories_consolidation,omitempty"`
}

func (s *Source) Validate() error {
	return s.ValidateWithConfig(nil)
}

func (s *Source) ValidateWithConfig(appConfig *AppConfig) error {
	if s.Name == "" {
		return fmt.Errorf("name is required")
	}

	// Only perform an alphanumeric check if skip_name_special_chars_check is false
	skipCheck := appConfig != nil && appConfig.DNSToolkit.SkipNameSpecialCharsCheck
	if !skipCheck && !u.IsAlphanumericWithUnderscoresAndDashes(s.Name) {
		return fmt.Errorf("name can only contain alphanumeric characters, underscores, and dashes: %s", s.Name)
	}

	if s.URL == "" {
		return fmt.Errorf("url is required")
	} else {
		for _, ext := range constants.ArchiveExtensions {
			if strings.HasSuffix(s.URL, ext) {
				if len(s.Files) == 0 {
					return fmt.Errorf("files property is required for url ending with %s", ext)
				}
				break
			}
		}
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
	for _, category := range s.Categories {
		if !constants.ValidCategories[category] {
			return fmt.Errorf("invalid category: %s", category)
		}
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
	aux := &struct { // nolint: govet
		Types []struct { // nolint: govet
			ListTypes []struct {
				Name         string `json:"name"`
				Groups       string `json:"groups"`
				Notes        string `json:"notes"`
				MustConsider bool   `json:"must_consider"`
				Disabled     bool   `json:"disabled"`
			} `json:"list_types,omitempty"`
			Name string `json:"name"`
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

// resolveFilePath resolves file paths, handling relative paths in test mode
func resolveFilePath(filePath string) string {
	// If a path is absolute, return as-is
	if filepath.IsAbs(filePath) {
		return filePath
	}

	if os.Getenv("DNS_TOOLKIT_TEST_MODE") == "true" {
		wd, err := os.Getwd()
		if err != nil {
			return filePath
		}

		projectRoot, err := u.FindProjectRoot(wd)
		if err == nil && projectRoot != "" {
			resolvedPath := filepath.Join(projectRoot, filePath)
			if _, err := os.Stat(resolvedPath); err == nil {
				return resolvedPath
			}
		}
	}

	return filePath
}

func LoadSourcesConfig(logger *multilog.Logger, filePath string) (SourcesConfig, error) {
	logger.Debugf("Loading sources config from %s", filePath)

	resolvedPath := resolveFilePath(filePath)

	file, err := os.Open(resolvedPath)
	if err != nil {
		return SourcesConfig{}, fmt.Errorf("error opening sources file: %w", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			logger.Errorf("error closing file: %v", closeErr)
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
		if len(source.Categories) == 0 {
			config.Sources[i].Categories = []string{defaultValues["Category"]}
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
		if source.IsEnabled() && matchesFilters(source, filters) {
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

// GetSourcesForGeneralConsolidation returns sources that should be included in general consolidation.
// This excludes sources with SkipGeneralConsolidation=true.
func (sc *SourcesConfig) GetSourcesForGeneralConsolidation(filters SourceFilters) []Source {
	uniqueSources := make(map[string]Source)
	for _, source := range sc.Sources {
		if source.ShouldIncludeInGeneralConsolidation() && matchesFilters(source, filters) {
			key := fmt.Sprintf("%s_%s", source.Name, source.Types[0].Name)
			uniqueSources[key] = source
		}
	}
	var sources []Source
	for _, source := range uniqueSources {
		sources = append(sources, source)
	}
	sort.Slice(sources, func(i, j int) bool {
		return strings.ToLower(sources[i].Name) < strings.ToLower(sources[j].Name)
	})
	return sources
}

// GetSourcesForGroupsConsolidation returns sources that should be included in groups consolidation.
// This includes all enabled sources, regardless of SkipGeneralConsolidation setting.
func (sc *SourcesConfig) GetSourcesForGroupsConsolidation(filters SourceFilters) []Source {
	uniqueSources := make(map[string]Source)
	for _, source := range sc.Sources {
		if source.ShouldIncludeInGroupsConsolidation() && matchesFilters(source, filters) {
			key := fmt.Sprintf("%s_%s", source.Name, source.Types[0].Name)
			uniqueSources[key] = source
		}
	}
	var sources []Source
	for _, source := range uniqueSources {
		sources = append(sources, source)
	}
	sort.Slice(sources, func(i, j int) bool {
		return strings.ToLower(sources[i].Name) < strings.ToLower(sources[j].Name)
	})
	return sources
}

// GetSourcesForCategoriesConsolidation returns sources that should be included in categories consolidation.
func (sc *SourcesConfig) GetSourcesForCategoriesConsolidation(filters SourceFilters) []Source {
	uniqueSources := make(map[string]Source)
	for _, source := range sc.Sources {
		if source.ShouldIncludeInCategoriesConsolidation() && matchesFilters(source, filters) {
			key := fmt.Sprintf("%s_%s", source.Name, source.Types[0].Name)
			uniqueSources[key] = source
		}
	}
	var sources []Source
	for _, source := range uniqueSources {
		sources = append(sources, source)
	}
	sort.Slice(sources, func(i, j int) bool {
		return strings.ToLower(sources[i].Name) < strings.ToLower(sources[j].Name)
	})
	return sources
}

func matchesFilters(source Source, filters SourceFilters) bool {
	// Check Name filter
	if len(filters.Name.Contains) > 0 {
		matches := false
		for _, substr := range filters.Name.Contains {
			if strings.Contains(source.Name, substr) {
				matches = true
				break
			}
		}
		if !matches {
			return false
		}
	}
	if len(filters.Name.NotContains) > 0 {
		for _, substr := range filters.Name.NotContains {
			if strings.Contains(source.Name, substr) {
				return false
			}
		}
	}

	// Check Type filter
	if filters.Type != "" {
		typeMatches := false
		for _, t := range source.Types {
			if t.Name == filters.Type {
				typeMatches = true
				break
			}
		}
		if !typeMatches {
			return false
		}
	}

	// Check Frequency filter
	if filters.Frequency != "" && source.Frequency != filters.Frequency {
		return false
	}

	// Check Category filter
	if len(filters.Category.Contains) > 0 {
		matches := false
		for _, substr := range filters.Category.Contains {
			for _, category := range source.Categories {
				if strings.Contains(category, substr) {
					matches = true
					break
				}
			}
			if matches {
				break
			}
		}
		if !matches {
			return false
		}
	}
	if len(filters.Category.NotContains) > 0 {
		for _, substr := range filters.Category.NotContains {
			for _, category := range source.Categories {
				if strings.Contains(category, substr) {
					return false
				}
			}
		}
	}

	// Check ListType filter
	if filters.ListType != "" {
		listTypeMatches := false
		for _, t := range source.Types {
			for _, lt := range t.GetListTypes() {
				if lt.Name == filters.ListType {
					listTypeMatches = true
					break
				}
			}
			if listTypeMatches {
				break
			}
		}
		if !listTypeMatches {
			return false
		}
	}

	// Check Countries filter
	if len(filters.Countries.Contains) > 0 {
		matches := false
		for _, substr := range filters.Countries.Contains {
			for _, country := range source.Countries {
				if strings.Contains(country, substr) {
					matches = true
					break
				}
			}
			if matches {
				break
			}
		}
		if !matches {
			return false
		}
	}
	if len(filters.Countries.NotContains) > 0 {
		for _, substr := range filters.Countries.NotContains {
			for _, country := range source.Countries {
				if strings.Contains(country, substr) {
					return false
				}
			}
		}
	}

	// Check Group filter
	if filters.Group != "" {
		// groupMatches := false
		for _, t := range source.Types {
			for _, lt := range t.GetListTypes() {
				for _, group := range lt.Groups {
					if group == filters.Group {
						return true
					}
				}
			}
		}
		return false
	}

	return true
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
		Name:      s.Name,
		URL:       s.URL,
		Folder:    downloadDir,
		Frequency: s.Frequency,
		Targets:   make([]c.DownloadTarget, 0, len(s.Files)), // Pre-allocate capacity
	}

	if u.IsArchive(s.URL) {
		ext := u.GetArchiveExtension(s.URL)
		if ext == "" {
			return downloadFile, fmt.Errorf("unknown archive extension for %s", s.URL)
		}

		if len(s.Files) == 0 {
			return downloadFile, fmt.Errorf("files are required for url ending with %s", ext)
		}

		downloadFile.Filename = s.Name + ext
		downloadFile.IsArchive = true

		// Process all files in the archive
		for _, file := range s.Files {
			filename := fmt.Sprintf("%s-%s", s.Name, strings.ReplaceAll(file, "/", "_"))
			downloadFile.Targets = append(downloadFile.Targets, c.DownloadTarget{
				SourceFolder: downloadFile.Folder,
				SourceFile:   file,
				TargetFile:   filename,
				TargetFolder: downloadFile.Folder,
			})
		}
	} else {
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
	}

	return downloadFile, nil
}

// IsEnabled returns true if the source is enabled.
func (s *Source) IsEnabled() bool {
	return !s.Disabled
}

// ShouldIncludeInGeneralConsolidation returns true if the source should be included in general consolidation.
// Sources with SkipGeneralConsolidation=true will be excluded from general consolidation
// but still included in groups and categories consolidation.
func (s *Source) ShouldIncludeInGeneralConsolidation() bool {
	return s.IsEnabled() && !s.SkipGeneralConsolidation
}

// ShouldIncludeInGroupsConsolidation returns true if the source should be included in groups consolidation.
func (s *Source) ShouldIncludeInGroupsConsolidation() bool {
	return s.IsEnabled() && !s.SkipGroupsConsolidation
}

// ShouldIncludeInCategoriesConsolidation returns true if the source should be included in categories consolidation.
func (s *Source) ShouldIncludeInCategoriesConsolidation() bool {
	return s.IsEnabled() && !s.SkipCategoriesConsolidation
}

// GetUserAgent returns a user agent string using the util function.
// This is a wrapper to maintain API compatibility while avoiding import cycles.
func GetUserAgent(logger *multilog.Logger, applicationConfig ApplicationConfig) string {
	appName := applicationConfig.Name
	appVersion := applicationConfig.Version
	appDesc := applicationConfig.Description
	return u.GetUserAgent(logger, appName, appVersion, appDesc)
}
