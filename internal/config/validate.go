package config

// ValidateDNSToolkitAppConfig is another validation helper
func ValidateDNSToolkitAppConfig(config AppConfig) error {
	return config.Validate()
}
