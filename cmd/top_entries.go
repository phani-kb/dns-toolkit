package cmd

import (
	"os"
	"runtime"

	cfg "github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/top"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/spf13/cobra"
)

var minSources int
var maxEntries int
var cpuProfile bool
var memProfile bool
var goroutineProfile bool
var blockProfile bool
var profileDir string
var maxWorkers int

var topEntriesCmd = &cobra.Command{
	Use:   "top",
	Short: "Find top entry(s) in each generic source type",
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Infof("Profiling configuration: CPU=%v, Memory=%v, Goroutine=%v, Block=%v",
			cpuProfile, memProfile, goroutineProfile, blockProfile)

		maxWorkers = AppConfig.DNSToolkit.MaxWorkers
		if maxWorkers <= 0 {
			maxWorkers = runtime.GOMAXPROCS(0)
		}
		Logger.Infof("Using %d worker(s) for top entry(s) processing", maxWorkers)

		if profileDir != "" {
			if err := os.MkdirAll(profileDir, 0755); err != nil {
				Logger.Errorf("Failed to create profile directory %s: %v", profileDir, err)
				profileDir = ""
			} else {
				Logger.Infof("Using profile directory: %s", profileDir)
			}
		}
		if profileDir == "" {
			profileDir = AppConfig.DNSToolkit.Folders.Profiles
			if err := os.MkdirAll(profileDir, 0755); err != nil {
				Logger.Errorf(
					"Failed to create default profile directory %s: %v. Profiling might fail.",
					profileDir,
					err,
				)
			}
		}

		stopProfiling := u.StartProfiling(Logger, u.ProfileOptions{
			CPUProfile:       cpuProfile,
			MemProfile:       memProfile,
			GoroutineProfile: goroutineProfile,
			BlockProfile:     blockProfile,
			ProfileNameBase:  "top",
			OutputDir:        profileDir,
			BlockProfileRate: 1000,
		})

		profilingEnabled := cpuProfile || memProfile || goroutineProfile || blockProfile

		var mainErr error
		func() {
			defer stopProfiling()

			if maxEntries < 1 {
				Logger.Errorf("Error: maxEntries must be at least 1, got %d", maxEntries)
				return
			}

			// Create an instance of the top entries service
			topService := top.NewDefaultService(constants.TopDir, constants.SummaryDir)

			// Ensure directories exist
			if err := u.EnsureDirectoryExists(Logger, constants.TopDir); err != nil {
				Logger.Errorf("Failed to create top entry(s) directory %s: %v", constants.TopDir, err)
				mainErr = err
				return
			}
			if err := u.EnsureDirectoryExists(Logger, constants.SummaryDir); err != nil {
				Logger.Errorf("Failed to create summary directory %s: %v", constants.SummaryDir, err)
				mainErr = err
				return
			}

			// Get processed files and generic source types
			_, genericSourceTypes, processedFiles := cfg.GetProcessedSummaries(Logger, SourcesConfigs, *AppConfig)

			// Determine min sources range
			minSourcesRange := []int{minSources}
			if minSources == 0 {
				minSourcesRange = constants.DefaultMinSourcesRange
			}

			// Process top entries using the service
			_, err := topService.ProcessTopEntries(
				Logger,
				genericSourceTypes,
				processedFiles,
				minSourcesRange,
				maxEntries,
				maxWorkers,
			)

			if err != nil {
				Logger.Errorf("Error processing top entries: %v", err)
				mainErr = err
			}
		}()

		if mainErr != nil {
			Logger.Errorf("Top entry(s) processing finished with error: %v", mainErr)
		}

		if profilingEnabled {
			Logger.Infof("Analyzing collected profiles...")
			err := u.AnalyzeProfiles(Logger, u.ProfileOptions{
				ProfileNameBase: "top",
				OutputDir:       profileDir,
			})
			if err != nil {
				Logger.Errorf("Failed to analyze profiles: %v", err)
			}
		}
	},
}

func init() {
	topEntriesCmd.Flags().IntVarP(&minSources, "min-sources", "m", 0, "Minimum sources (3-12)")
	topEntriesCmd.Flags().IntVarP(&maxEntries, "max-entries", "x", int(^uint(0)>>1), "Max entries")

	AddProfilingFlags(topEntriesCmd, &cpuProfile, &memProfile, &goroutineProfile, &blockProfile, &profileDir)
}
