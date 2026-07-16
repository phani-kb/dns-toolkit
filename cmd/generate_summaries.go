package cmd

import (
	"context"
	"os"
	"path/filepath"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/db"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/spf13/cobra"
)

var generateSummariesCmd = &cobra.Command{
	Use:   "summaries",
	Short: "Generate all summary files for publishing",
	Long: "Generates download_summary.json and processed_summary.json " +
		"and saves them to the output summaries directory for publishing " +
		"to summaries branch.",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		if err := u.EnsureDirectoryExists(Logger, constants.OutputSummariesDir); err != nil {
			Logger.Errorf("Failed to create output summaries directory: %v", err)
			os.Exit(1)
		}

		database := openDB(ctx)
		defer database.CloseLogError(Logger)

		summaryCount := 0

		downloadsRepo := db.NewDownloadsRepo(database)
		downloadSummaries, downloadErr := downloadsRepo.ListDownloadSummaries(constants.DownloadDir)
		if downloadErr != nil {
			Logger.Errorf("Failed to load download summaries: %v", downloadErr)
		} else {
			downloadSummaryFile := filepath.Join(constants.OutputSummariesDir,
				constants.DefaultSummaryFiles["download"])
			count, saveErr := u.SaveSummaries(Logger, downloadSummaries, downloadSummaryFile,
				c.DownloadSummaryLessFunc)
			if saveErr != nil {
				Logger.Errorf("Failed to save download summaries: %v", saveErr)
			} else {
				Logger.Infof("Generated download_summary.json with %d summaries", count)
				summaryCount += count
			}
		}

		processedRepo := db.NewProcessedRepo(database)
		processedSummaries, processedErr := processedRepo.ListProcessedSummaries(constants.ProcessedDir)
		if processedErr != nil {
			Logger.Errorf("Failed to load processed summaries: %v", processedErr)
		} else {
			processedSummaryFile := filepath.Join(constants.OutputSummariesDir,
				constants.DefaultSummaryFiles["processed"])
			count, saveErr := u.SaveSummaries(Logger, processedSummaries, processedSummaryFile,
				c.ProcessedSummaryLessFunc)
			if saveErr != nil {
				Logger.Errorf("Failed to save processed summaries: %v", saveErr)
			} else {
				Logger.Infof("Generated processed_summary.json with %d summaries", count)
				summaryCount += count
			}
		}

		Logger.Infof("Successfully generated summary files: %d total summaries", summaryCount)
	},
}

func init() {
	generateCmd.AddCommand(generateSummariesCmd)
}
