package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/phani-kb/multilog"
)

// ProfileOptions defines the configuration for profiling
type ProfileOptions struct {
	ProfileNameBase  string
	OutputDir        string
	BlockProfileRate int
	CPUProfile       bool
	MemProfile       bool
	GoroutineProfile bool
	BlockProfile     bool
}

// createProfile creates a pprof profile of the specified type
func createProfile(logger *multilog.Logger, profileType, outputDir, profileNameBase string) {
	profilePath := filepath.Join(outputDir, fmt.Sprintf("%s_%s.prof", profileNameBase, profileType))
	logger.Infof("Creating %s profile at %s", profileType, profilePath)

	f, err := os.Create(profilePath)
	if err != nil {
		logger.Errorf("Could not create %s profile: %v", profileType, err)
		return
	}

	err = pprof.Lookup(profileType).WriteTo(f, 0)
	closeErr := f.Close()

	if err != nil {
		logger.Errorf("Could not write %s profile: %v", profileType, err)
	} else {
		logger.Infof("%s profile saved to %s", profileType, profilePath)
	}

	if closeErr != nil {
		logger.Errorf("Could not close %s profile file: %v", profileType, closeErr)
	}
}

// StartProfiling starts profiling based on provided options
// Returns a function that should be deferred to stop profiling and create profiles if enabled
func StartProfiling(logger *multilog.Logger, opts ProfileOptions) func() {
	startTime := time.Now()
	var cpuFile *os.File
	var cpuProfileActive bool

	// If no profiles are enabled, return a simple function that just logs the execution time
	if !opts.CPUProfile && !opts.MemProfile && !opts.GoroutineProfile && !opts.BlockProfile {
		return func() {
			elapsed := time.Since(startTime)
			logger.Debugf("Command completed in %s", elapsed)
		}
	}

	// Use command name if profile name base is not provided
	if opts.ProfileNameBase == "" {
		opts.ProfileNameBase = "command"
	}

	// Ensure output directory exists
	if opts.OutputDir != "" {
		err := os.MkdirAll(opts.OutputDir, 0755)
		if err != nil {
			logger.Errorf("Failed to create output directory for profiles: %v", err)
			opts.OutputDir = ""
		}
	}

	// Enable block profiling if requested
	if opts.BlockProfile {
		rate := opts.BlockProfileRate
		if rate <= 0 {
			rate = 1 // Default value if isn't specified or invalid
		}
		runtime.SetBlockProfileRate(rate)
		logger.Info("Block profiling enabled")
	}

	// Start CPU profiling if enabled
	if opts.CPUProfile {
		cpuProfilePath := filepath.Join(opts.OutputDir, fmt.Sprintf("%s_cpu.prof", opts.ProfileNameBase))
		logger.Infof("CPU profiling enabled, writing to %s", cpuProfilePath)

		var err error
		cpuFile, err = os.Create(cpuProfilePath)
		if err != nil {
			logger.Errorf("Could not create CPU profile: %v", err)
		} else {
			err = pprof.StartCPUProfile(cpuFile)
			if err != nil {
				logger.Errorf("Could not start CPU profile: %v", err)
				if err := cpuFile.Close(); err != nil {
					logger.Errorf("Could not close CPU profile file: %v", err)
				}
				cpuFile = nil
			} else {
				cpuProfileActive = true
				logger.Info("CPU profiling started successfully")
			}
		}
	}

	// Return a function that stops profiling and records memory profile if enabled
	return func() {
		// Stop CPU profiling if it was started
		if opts.CPUProfile && cpuFile != nil {
			logger.Info("Stopping CPU profiling...")
			pprof.StopCPUProfile()
			if err := cpuFile.Close(); err != nil {
				logger.Errorf("Could not close CPU profile file: %v", err)
			}

			// Verify that the profile file has content
			if cpuProfileActive {
				cpuProfilePath := filepath.Join(opts.OutputDir, fmt.Sprintf("%s_cpu.prof", opts.ProfileNameBase))
				info, err := os.Stat(cpuProfilePath)
				if err != nil {
					logger.Errorf("Error checking CPU profile file: %v", err)
				} else if info.Size() == 0 {
					logger.Warnf("CPU profile file is empty. No profiling data was collected.")
				} else {
					logger.Infof("CPU profiling completed successfully. Profile size: %d bytes", info.Size())
				}
			}
		}

		// Create a memory profile if enabled
		if opts.MemProfile {
			memProfilePath := filepath.Join(opts.OutputDir, fmt.Sprintf("%s_mem.prof", opts.ProfileNameBase))
			logger.Infof("Creating memory profile at %s", memProfilePath)

			// Force garbage collection to get a more accurate memory profile
			runtime.GC()

			f, err := os.Create(memProfilePath)
			if err != nil {
				logger.Errorf("Could not create memory profile: %v", err)
			} else {
				err = pprof.WriteHeapProfile(f)
				closeErr := f.Close()

				if err != nil {
					logger.Errorf("Could not write memory profile: %v", err)
				} else {
					logger.Infof("Memory profile saved to %s", memProfilePath)
				}

				if closeErr != nil {
					logger.Errorf("Could not close memory profile file: %v", closeErr)
				}
			}
		}

		// Create a goroutine profile if enabled
		if opts.GoroutineProfile {
			createProfile(logger, "goroutine", opts.OutputDir, opts.ProfileNameBase)
		}

		// Create a block profile if enabled
		if opts.BlockProfile {
			createProfile(logger, "block", opts.OutputDir, opts.ProfileNameBase)
		}

		LogMemStats(logger, "End")

		// Log execution time
		elapsed := time.Since(startTime)
		logger.Infof("Command completed in %s", elapsed)
	}
}

// AnalyzeProfiles runs pprof analysis on generated profiles and saves the result to text files
// This is useful for automated analysis of profiles without requiring manual pprof commands
func AnalyzeProfiles(logger *multilog.Logger, opts ProfileOptions) {
	if opts.ProfileNameBase == "" {
		opts.ProfileNameBase = "command"
	}

	if opts.OutputDir == "" {
		opts.OutputDir = "."
	}

	// Helper function to run pprof analysis
	runPprofAnalysis := func(profileType, profilePath string) error {
		outputPath := filepath.Join(
			opts.OutputDir,
			fmt.Sprintf("%s_%s_analysis.txt", opts.ProfileNameBase, profileType),
		)
		logger.Infof("Analyzing %s profile at %s, saving results to %s", profileType, profilePath, outputPath)

		// Create an output file for analysis
		outFile, err := os.Create(outputPath)
		if err != nil {
			return fmt.Errorf("could not create analysis output file: %v", err)
		}
		defer func(outFile *os.File) {
			err := outFile.Close()
			if err != nil {
				logger.Errorf("Error closing output file %s: %v", outFile.Name(), err)
			}
		}(outFile)

		// Setup and run command to analyze profile
		cmd := exec.Command("go", "tool", "pprof", "-text", profilePath)
		cmd.Stdout = outFile
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to analyze %s profile: %v", profileType, err)
		}

		logger.Infof("Profile analysis for %s completed and saved to %s", profileType, outputPath)
		return nil
	}

	// Process CPU profile if it exists
	cpuProfilePath := filepath.Join(opts.OutputDir, fmt.Sprintf("%s_cpu.prof", opts.ProfileNameBase))
	if info, err := os.Stat(cpuProfilePath); err == nil {
		if info.Size() == 0 {
			logger.Warnf("CPU profile file %s is empty, skipping analysis", cpuProfilePath)
		} else if analysisErr := runPprofAnalysis("cpu", cpuProfilePath); analysisErr != nil {
			logger.Errorf("CPU profile analysis failed: %v", analysisErr)
		}
	} else {
		logger.Debugf("CPU profile not found at %s: %v", cpuProfilePath, err)
	}

	// Process memory profile if it exists
	memProfilePath := filepath.Join(opts.OutputDir, fmt.Sprintf("%s_mem.prof", opts.ProfileNameBase))
	if info, err := os.Stat(memProfilePath); err == nil {
		if info.Size() == 0 {
			logger.Warnf("Memory profile file %s is empty, skipping analysis", memProfilePath)
		} else if err := runPprofAnalysis("mem", memProfilePath); err != nil {
			logger.Errorf("Memory profile analysis failed: %v", err)
		}
	} else {
		logger.Debugf("No memory profile found at %s, skipping analysis", memProfilePath)
	}

	// Process goroutine profile if it exists
	goroutineProfilePath := filepath.Join(opts.OutputDir, fmt.Sprintf("%s_goroutine.prof", opts.ProfileNameBase))
	if info, err := os.Stat(goroutineProfilePath); err == nil {
		if info.Size() == 0 {
			logger.Warnf("Goroutine profile file %s is empty, skipping analysis", goroutineProfilePath)
		} else if err := runPprofAnalysis("goroutine", goroutineProfilePath); err != nil {
			logger.Errorf("Goroutine profile analysis failed: %v", err)
		}
	} else {
		logger.Debugf("No goroutine profile found at %s, skipping analysis", goroutineProfilePath)
	}

	// Process block profile if it exists
	blockProfilePath := filepath.Join(opts.OutputDir, fmt.Sprintf("%s_block.prof", opts.ProfileNameBase))
	if info, err := os.Stat(blockProfilePath); err == nil {
		if info.Size() == 0 {
			logger.Warnf("Block profile file %s is empty, skipping analysis", blockProfilePath)
		} else if err := runPprofAnalysis("block", blockProfilePath); err != nil {
			logger.Errorf("Block profile analysis failed: %v", err)
		}
	} else {
		logger.Debugf("No block profile found at %s, skipping analysis", blockProfilePath)
	}

}
