package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/multilog"
	"gopkg.in/yaml.v2"
)

type NameFilter struct {
	Contains    []string `yaml:"contains,omitempty"`
	NotContains []string `yaml:"not_contains,omitempty"`
}

type SourceFilters struct {
	Name      NameFilter `yaml:"name,omitempty"`
	Type      string     `yaml:"type,omitempty"`
	Frequency string     `yaml:"frequency,omitempty"`
	Category  NameFilter `yaml:"category,omitempty"`
	Group     string     `yaml:"group,omitempty"`
	ListType  string     `yaml:"list_type,omitempty"`
	Countries NameFilter `yaml:"countries,omitempty"`
}

type FilesChecksumConfig struct {
	Algorithm string `yaml:"algorithm"`
	Enabled   bool   `yaml:"enabled"`
}

type DNSToolkitConfig struct {
	SourceFiles               []string            `yaml:"source_files"`
	SkipCertVerificationHosts []string            `yaml:"skip_cert_verification_hosts,omitempty"`
	Folders                   FoldersConfig       `yaml:"folders"`
	SourceFilters             SourceFilters       `yaml:"source_filters"`
	FilesChecksum             FilesChecksumConfig `yaml:"files_checksum"`
	MaxWorkers                int                 `yaml:"max_workers"`
	MaxRetries                int                 `yaml:"max_retries"`
	SkipUnchangedDownloads    bool                `yaml:"skip_unchanged_downloads"`
	SkipCertVerification      bool                `yaml:"skip_cert_verification,omitempty"`
	SkipNameSpecialCharsCheck bool                `yaml:"skip_name_special_chars_check,omitempty"`
}

func (dc *DNSToolkitConfig) Validate() error {
	if len(dc.SourceFiles) == 0 {
		return errors.New("at least one source file is required")
	}
	for _, sourceFile := range dc.SourceFiles {
		if err := validateSourceFile(sourceFile); err != nil {
			return fmt.Errorf("source file %s not found: %w", sourceFile, err)
		}
	}

	if dc.MaxWorkers > runtime.GOMAXPROCS(0) {
		dc.MaxWorkers = runtime.GOMAXPROCS(0)
	}

	return nil
}

// validateSourceFile checks if a source file exists, handling relative paths in test mode
func validateSourceFile(sourceFile string) error {
	resolvedPath := resolveFilePath(sourceFile)
	_, err := os.Stat(resolvedPath)
	return err
}

type FoldersConfig struct {
	Download               string `yaml:"download"`
	Processed              string `yaml:"processed"`
	Consolidated           string `yaml:"consolidated"`
	Summary                string `yaml:"summary"`
	Overlap                string `yaml:"overlap"`
	Top                    string `yaml:"top"`
	ConsolidatedGroups     string `yaml:"consolidated_groups"`
	ConsolidatedCategories string `yaml:"consolidated_categories"`
	Archive                string `yaml:"archive"`
	Output                 string `yaml:"output"`
	Summaries              string `yaml:"summaries"`
	Backup                 string `yaml:"backup"`
	Profiles               string `yaml:"profiles"`
}

type AppConfig struct { // nolint: govet
	DNSToolkit  DNSToolkitConfig  `yaml:"dns_toolkit"`
	Application ApplicationConfig `yaml:"application"`
	Multilog    interface{}       `yaml:"multilog"`
}

func (ac *AppConfig) Validate() error {
	if err := ac.Application.Validate(); err != nil {
		return errors.New("application validation failed, " + err.Error())
	}
	if err := ac.DNSToolkit.Validate(); err != nil {
		return errors.New("dns toolkit validation failed, " + err.Error())
	}
	if ac.Multilog == nil {
		return errors.New("multilog configuration is required")
	}

	return nil
}

type ApplicationConfig struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Description string `yaml:"description"`
}

func (ac *ApplicationConfig) Validate() error {
	if ac.Name == "" {
		return errors.New("application name is required")
	}
	if ac.Version == "" {
		return errors.New("application version is required")
	}
	if ac.Description == "" {
		return errors.New("application description is required")
	}
	return nil
}

func LoadAppConfig(logger *multilog.Logger, configPath string) (AppConfig, []SourcesConfig, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return AppConfig{}, nil, fmt.Errorf("opening config file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.Errorf("Error closing config file: %v", err)
		}
	}(file)

	var appConfig AppConfig
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&appConfig); err != nil {
		return AppConfig{}, nil, fmt.Errorf("decoding config file: %w", err)
	}

	if err := ValidateAppConfig(appConfig); err != nil {
		return AppConfig{}, nil, err
	}

	var sourcesConfigs []SourcesConfig
	for _, sourceFile := range appConfig.DNSToolkit.SourceFiles {
		sourcesConfig, err := LoadSourcesConfig(logger, sourceFile)
		if err != nil {
			return AppConfig{}, nil, fmt.Errorf("loading sources config: %w", err)
		}

		// Validate the sources config with the AppConfig
		if err := sourcesConfig.ValidateWithConfig(&appConfig); err != nil {
			return AppConfig{}, nil, fmt.Errorf("validating sources config: %w", err)
		}

		sourcesConfigs = append(sourcesConfigs, sourcesConfig)
	}

	return appConfig, sourcesConfigs, nil
}

func GetGenericSourceType(sourceType string) string {
	if alias, exists := constants.GenericSourceTypeAliases[sourceType]; exists {
		return alias
	}

	for _, genericType := range constants.GenericSourceTypes {
		if strings.HasPrefix(sourceType, genericType) {
			return genericType
		}
	}
	return sourceType
}

func IsEnabledSource(sourceName string, sourceConfigs []SourcesConfig, appConfig AppConfig) bool {
	return IsEnabledSourceForConsolidation(sourceName, sourceConfigs, appConfig, "general")
}

// IsEnabledSourceForConsolidation checks if a source is enabled for a specific consolidation type.
// consolidationType can be "general", "groups", or "categories"
func IsEnabledSourceForConsolidation(
	sourceName string,
	sourceConfigs []SourcesConfig,
	appConfig AppConfig,
	consolidationType string,
) bool {
	for _, sourcesConfig := range sourceConfigs {
		var sources []Source
		switch consolidationType {
		case "general":
			sources = sourcesConfig.GetSourcesForGeneralConsolidation(appConfig.DNSToolkit.SourceFilters)
		case "groups", "categories":
			sources = sourcesConfig.GetSourcesForGroupsAndCategories(appConfig.DNSToolkit.SourceFilters)
		default:
			// For unknown consolidation types, default to general consolidation behavior
			sources = sourcesConfig.GetSourcesForGeneralConsolidation(appConfig.DNSToolkit.SourceFilters)
		}

		for _, source := range sources {
			if source.Name == sourceName {
				return true
			}
		}
	}
	return false
}

// GetProcessedSummaries reads the processed summaries from the default summary file and returns them along with the
// generic source types.
func GetProcessedSummaries(
	logger *multilog.Logger,
	sourcesConfigs []SourcesConfig,
	appConfig AppConfig,
) ([]c.ProcessedSummary, []string, []c.ProcessedFile) {
	return GetProcessedSummariesForConsolidation(logger, sourcesConfigs, appConfig, "general")
}

// GetProcessedSummariesForConsolidation reads the processed summaries and filters them based on consolidation type.
// consolidationType can be "general", "groups", or "categories"
func GetProcessedSummariesForConsolidation(
	logger *multilog.Logger,
	sourcesConfigs []SourcesConfig,
	appConfig AppConfig,
	consolidationType string,
) ([]c.ProcessedSummary, []string, []c.ProcessedFile) {
	summaryFile := filepath.Join(constants.SummaryDir, constants.DefaultSummaryFiles["processed"])
	content, err := os.ReadFile(summaryFile)
	if err != nil {
		logger.Errorf("Reading file %s: %v", summaryFile, err)
		return nil, nil, nil
	}

	var summaries []c.ProcessedSummary
	if err := json.Unmarshal(content, &summaries); err != nil {
		logger.Errorf("Unmarshalling JSON: %v", err)
		return nil, nil, nil
	}

	enabledSummaries := filterEnabledSummariesForConsolidation(summaries, sourcesConfigs, appConfig, consolidationType)
	sort.Slice(enabledSummaries, func(i, j int) bool {
		return enabledSummaries[i].Name < enabledSummaries[j].Name
	})

	genericSourceTypes := extractGenericSourceTypes(enabledSummaries)
	processedFiles := GetAllProcessedFiles(summaries)
	logger.Infof(
		"Processed summaries count: %d, generic source types count: %d, files count: %d, consolidation type: %s",
		len(enabledSummaries),
		len(genericSourceTypes),
		len(processedFiles),
		consolidationType,
	)

	return enabledSummaries, genericSourceTypes, processedFiles
}

func GetAllProcessedFiles(
	processedSummaries []c.ProcessedSummary,
) []c.ProcessedFile {
	var processedFiles []c.ProcessedFile
	for _, summary := range processedSummaries {
		processedFiles = append(processedFiles, summary.ValidFiles...)
		processedFiles = append(processedFiles, summary.InvalidFiles...)
	}
	return processedFiles
}

func filterEnabledSummariesForConsolidation(
	summaries []c.ProcessedSummary,
	sourcesConfigs []SourcesConfig,
	appConfig AppConfig,
	consolidationType string,
) []c.ProcessedSummary {
	var enabledSummaries []c.ProcessedSummary
	for _, summary := range summaries {
		if IsEnabledSourceForConsolidation(summary.Name, sourcesConfigs, appConfig, consolidationType) {
			enabledSummaries = append(enabledSummaries, summary)
		}
	}
	return enabledSummaries
}

func extractGenericSourceTypes(summaries []c.ProcessedSummary) []string {
	sourceTypeMap := make(map[string]struct{})
	for _, summary := range summaries {
		for _, processedFile := range summary.ValidFiles {
			genericSourceType := processedFile.GenericSourceType
			sourceTypeMap[genericSourceType] = struct{}{}
		}
		for _, processedFile := range summary.InvalidFiles {
			genericSourceType := processedFile.GenericSourceType
			sourceTypeMap[genericSourceType] = struct{}{}
		}
	}

	var genericSourceTypes []string
	for sourceType := range sourceTypeMap {
		genericSourceTypes = append(genericSourceTypes, sourceType)
	}
	sort.Strings(genericSourceTypes)
	return genericSourceTypes
}
