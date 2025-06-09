package config

import (
	"fmt"
)

func ValidateAppConfig(config AppConfig) error {
	if err := config.Validate(); err != nil {
		return fmt.Errorf("application validation error: %w", err)
	}
	return nil
}
