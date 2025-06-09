package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

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
