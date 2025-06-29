package cmd

import (
	"os"
	"runtime"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	cfg "github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/overlap"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/spf13/cobra"
)

// Profiling flags for overlap command
var overlapCPUProfile bool
var overlapMemProfile bool
var overlapGoroutineProfile bool
var overlapBlockProfile bool
var overlapProfileDir string
var overlapMaxWorkers int

var overlapCmd = &cobra.Command{
	Use:   "overlap",
	Short: "Find overlap between source files",
	Run: func(cmd *cobra.Command, args []string) {

		overlapMaxWorkers = AppConfig.DNSToolkit.MaxWorkers
		if overlapMaxWorkers <= 0 {
			overlapMaxWorkers = runtime.GOMAXPROCS(0)
		}
		Logger.Infof("Using %d worker(s) for overlap entry processing", overlapMaxWorkers)
		// Validate profile directory
		if overlapProfileDir != "" {
			if err := os.MkdirAll(overlapProfileDir, 0755); err != nil {
				Logger.Errorf("Failed to create profile directory %s: %v", overlapProfileDir, err)
				overlapProfileDir = ""
			} else {
				Logger.Infof("Using profile directory: %s", overlapProfileDir)
			}
		}

		// Set the default profile directory if not provided
		if overlapProfileDir == "" {
			overlapProfileDir = AppConfig.DNSToolkit.Folders.Profiles
		}

		// Setup profiling
		stopProfiling := u.StartProfiling(Logger, u.ProfileOptions{
			CPUProfile:       overlapCPUProfile,
			MemProfile:       overlapMemProfile,
			GoroutineProfile: overlapGoroutineProfile,
			BlockProfile:     overlapBlockProfile,
			ProfileNameBase:  "overlap",
			OutputDir:        overlapProfileDir,
			BlockProfileRate: constants.DefaultBlockProfileRate,
		})

		profilingEnabled := overlapCPUProfile || overlapMemProfile || overlapGoroutineProfile || overlapBlockProfile

		var processedFiles []c.ProcessedFile
		var genericSourceTypes []string
		var processedSummaries []c.ProcessedSummary
		func() {
			defer stopProfiling()

			if err := u.EnsureDirectoryExists(Logger, constants.OverlapDir); err != nil {
				Logger.Errorf("Failed to create overlap directory: %v", err)
				os.Exit(1)
			}
			if err := u.EnsureDirectoryExists(Logger, constants.SummaryDir); err != nil {
				Logger.Errorf("Failed to create summary directory: %v", err)
				os.Exit(1)
			}

			processedSummaries, genericSourceTypes, processedFiles = cfg.GetProcessedSummaries(
				Logger,
				SourcesConfigs,
				*AppConfig,
			)

			if len(processedSummaries) == 0 {
				Logger.Errorf("No processed files found")
				return
			}

			if len(genericSourceTypes) == 0 {
				Logger.Errorf("No generic source types found")
				return
			}

			// Create an instance of the overlap service
			overlapService := overlap.NewDefaultService(constants.OverlapDir, constants.SummaryDir)

			// Use the service to write compact overlap summaries
			_, err := overlapService.WriteCompactOverlapSummaries(
				Logger,
				processedFiles,
				genericSourceTypes,
				overlapMaxWorkers,
			)

			if err != nil {
				Logger.Errorf("Error writing overlap summaries: %v", err)
			}
		}()

		// Now analyze profiles after profiling has stopped
		if profilingEnabled {
			u.AnalyzeProfiles(Logger, u.ProfileOptions{
				ProfileNameBase: "overlap",
				OutputDir:       overlapProfileDir,
			})
		}
	},
}

func init() {
	// Add profiling flags to the overlap command
	AddProfilingFlags(
		overlapCmd,
		&overlapCPUProfile,
		&overlapMemProfile,
		&overlapGoroutineProfile,
		&overlapBlockProfile,
		&overlapProfileDir,
	)
}
