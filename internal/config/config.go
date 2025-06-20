package config

import (
	"errors"
	"fmt"
	"os"
	"runtime"

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
	Enabled   bool   `yaml:"enabled"`
	Algorithm string `yaml:"algorithm"`
}

type DNSToolkitConfig struct {
	Folders                   FoldersConfig       `yaml:"folders"`
	SourceFiles               []string            `yaml:"source_files"`
	SourceFilters             SourceFilters       `yaml:"source_filters"`
	FilesChecksum             FilesChecksumConfig `yaml:"files_checksum"`
	MaxWorkers                int                 `yaml:"max_workers"`
	MaxRetries                int                 `yaml:"max_retries"`
	SkipUnchangedDownloads    bool                `yaml:"skip_unchanged_downloads"`
	SkipCertVerification      bool                `yaml:"skip_cert_verification,omitempty"`
	SkipCertVerificationHosts []string            `yaml:"skip_cert_verification_hosts,omitempty"`
	SkipNameSpecialCharsCheck bool                `yaml:"skip_name_special_chars_check,omitempty"`
}

func (dc *DNSToolkitConfig) Validate() error {
	if len(dc.SourceFiles) == 0 {
		return errors.New("at least one source file is required")
	}
	for _, sourceFile := range dc.SourceFiles {
		if _, err := os.Stat(sourceFile); err != nil {
			return fmt.Errorf("source file %s not found: %w", sourceFile, err)
		}

	}

	if dc.MaxWorkers > runtime.GOMAXPROCS(0) {
		dc.MaxWorkers = runtime.GOMAXPROCS(0)
	}

	return nil
}

type FoldersConfig struct {
	Download           string `yaml:"download"`
	Processed          string `yaml:"processed"`
	Consolidated       string `yaml:"consolidated"`
	Summary            string `yaml:"summary"`
	Overlap            string `yaml:"overlap"`
	Top                string `yaml:"top"`
	ConsolidatedGroups string `yaml:"consolidated_groups"`
}

type AppConfig struct {
	Application ApplicationConfig `yaml:"application"`
	DNSToolkit  DNSToolkitConfig  `yaml:"dns_toolkit"`
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
	defer file.Close()

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
