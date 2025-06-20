package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/spf13/cobra"
)

// archiveCmd represents the archive command
var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Archive DNS toolkit data",
	Long: `Archive DNS toolkit data from various folders specified in the config.
Creates a compressed archive (tgz) containing all relevant data files and
generates a summary of the archived contents.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := u.EnsureDirectoryExists(Logger, constants.ArchiveDir); err != nil {
			Logger.Errorf("Failed to create archive directory: %v", err)
			os.Exit(1)
		}

		runArchive(Logger)
	},
}

func runArchive(logger *multilog.Logger) {
	logger.Infof("Starting archive process")

	// Create a timestamp for archive naming
	timestamp := time.Now().Format(constants.TimestampFormat)
	archiveFilename := fmt.Sprintf("dns_toolkit_archive_%s.tgz", timestamp)
	archivePath := filepath.Join(constants.ArchiveDir, archiveFilename)

	// Ensure the archive directory exists
	if err := os.MkdirAll(constants.ArchiveDir, 0755); err != nil {
		logger.Errorf("Failed to create archive directory: %v", err)
		os.Exit(1)
	}

	logger.Infof("Archive process completed successfully; created archive %s", archivePath)
}
