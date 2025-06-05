package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/phani-kb/multilog"
	"gopkg.in/yaml.v2"
)

type NameFilter struct {
	Contains    []string `yaml:"contains,omitempty"`
	NotContains []string `yaml:"not_contains,omitempty"`
}

type SourceFilters struct {
	Name NameFilter `yaml:"name,omitempty"`
	Type string     `yaml:"type,omitempty"`
}

type FilesChecksumConfig struct {
	Enabled   bool   `yaml:"enabled"`
	Algorithm string `yaml:"algorithm"`
}

type DNSToolkitConfig struct {
	Folders     FoldersConfig `yaml:"folders"`
	SourceFiles []string      `yaml:"source_files"`
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

	return nil
}

type FoldersConfig struct {
	Download string `yaml:"download"`
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

	var appConfig AppConfig
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&appConfig); err != nil {
		return AppConfig{}, nil, fmt.Errorf("decoding config file: %w", err)
	}

	var sourcesConfigs []SourcesConfig
	for _, sourceFile := range appConfig.DNSToolkit.SourceFiles {
		sourcesConfig, err := LoadSourcesConfig(logger, sourceFile)
		if err != nil {
			return AppConfig{}, nil, fmt.Errorf("loading sources config: %w", err)
		}

		sourcesConfigs = append(sourcesConfigs, sourcesConfig)
	}

	return appConfig, sourcesConfigs, nil
}
