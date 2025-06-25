package cmd

import (
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/spf13/cobra"
)

func TestVersionCommand(t *testing.T) {
	if AppConfig == nil {
		AppConfig = &config.AppConfig{}
		AppConfig.Application.Name = "DNS-Toolkit"
		AppConfig.Application.Version = "test-version"
	}

	runFunc := versionCmd.Run

	runFunc(versionCmd, []string{})
}

func TestVersionCommandNoArgs(t *testing.T) {
	cmd := &cobra.Command{}

	runFunc := versionCmd.Run

	runFunc(cmd, []string{})
}
