package cmd

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

// Test the Execute function
func TestExecute(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"dns-toolkit", "version"}

	Execute()
}

func TestValidateAndSetDirs(t *testing.T) {
	oldConfig := AppConfig
	AppConfig = nil
	validateAndSetDirs()

	AppConfig = oldConfig
	validateAndSetDirs()

	InitForTesting()
}

// Test the root command's PersistentPreRun and PersistentPostRun
func TestRootCmdHooks(t *testing.T) {
	preRun := rootCmd.PersistentPreRunE
	if preRun == nil {
		preRun = func(cmd *cobra.Command, args []string) error {
			if rootCmd.PersistentPreRun != nil {
				rootCmd.PersistentPreRun(cmd, args)
			}
			return nil
		}
	}

	postRun := rootCmd.PersistentPostRunE
	if postRun == nil {
		postRun = func(cmd *cobra.Command, args []string) error {
			if rootCmd.PersistentPostRun != nil {
				rootCmd.PersistentPostRun(cmd, args)
			}
			return nil
		}
	}

	cmd := &cobra.Command{Use: "test-cmd"}
	var args []string

	err := preRun(cmd, args)
	assert.NoError(t, err)

	err = postRun(cmd, args)
	assert.NoError(t, err)

	cmd = &cobra.Command{Use: "help"}
	cmd.Name()

	err = preRun(cmd, args)
	assert.NoError(t, err)

	err = postRun(cmd, args)
	assert.NoError(t, err)
}
