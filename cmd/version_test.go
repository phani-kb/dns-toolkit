package cmd

import (
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestVersionCommand(t *testing.T) {
	oldConfig := AppConfig
	defer func() { AppConfig = oldConfig }()

	AppConfig = &config.AppConfig{
		Application: config.ApplicationConfig{
			Name:    "DNS-Toolkit",
			Version: "test-version",
		},
	}

	assert.Equal(t, "version", versionCmd.Use)
	assert.Equal(t, "Print the version number of DNS Toolkit", versionCmd.Short)
	assert.NotNil(t, versionCmd.Run)

	runFunc := versionCmd.Run
	runFunc(versionCmd, []string{})
}

func TestVersionCommandStructure(t *testing.T) {
	assert.Equal(t, "version", versionCmd.Use)
	assert.Contains(t, versionCmd.Short, "version")
}

func TestVersionCommandNoArgs(t *testing.T) {
	cmd := &cobra.Command{}

	runFunc := versionCmd.Run

	runFunc(cmd, []string{})
}
