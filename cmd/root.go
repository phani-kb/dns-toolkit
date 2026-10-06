package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/spf13/cobra"
)

var (
	startTime time.Time
	endTime   time.Time
)

// create a map of cobra commands to be ignored for completion
var ignoreCmds = map[string]bool{
	"help":    true,
	"sts":     true,
	"version": true,
}

// dirMapping=config+default+constant
type dirMapping struct {
	configVal *string
	target    *string
	fallback  string
}

func validateAndSetDirs() {
	if AppConfig == nil {
		return
	}

	projectRoot := resolveProjectRoot()
	testMode := isTestContext()

	folders := &AppConfig.DNSToolkit.Folders
	mappings := []dirMapping{
		{&folders.Download, &constants.DownloadDir, filepath.Join("data", "download")},
		{&folders.Processed, &constants.ProcessedDir, filepath.Join("data", "processed")},
		{&folders.Consolidated, &constants.ConsolidatedDir, filepath.Join("data", "consolidated")},
		{&folders.ConsolidatedGroups, &constants.ConsolidatedGroupsDir, filepath.Join("data", "consolidated_groups")},
		{
			&folders.ConsolidatedCategories,
			&constants.ConsolidatedCategoriesDir,
			filepath.Join("data", "consolidated_categories"),
		},
		{&folders.Summary, &constants.SummaryDir, "data"},
		{&folders.Overlap, &constants.OverlapDir, filepath.Join("data", "overlap")},
		{&folders.Top, &constants.TopDir, filepath.Join("data", "top")},
		{&folders.Archive, &constants.ArchiveDir, filepath.Join("data", "archive")},
		{&folders.Backup, &constants.BackupDir, filepath.Join("data", "backup")},
		{&folders.Output, &constants.OutputDir, filepath.Join("data", "output")},
		{&folders.Profiles, &constants.ProfilesDir, filepath.Join("data", "profiles")},
	}

	for _, m := range mappings {
		dir := *m.configVal
		if dir == "" {
			dir = m.fallback
		}
		if testMode {
			dir = remapDataDirForTests(dir)
		}
		if !filepath.IsAbs(dir) && projectRoot != "" {
			dir = filepath.Join(projectRoot, dir)
		}
		*m.target = dir
	}

	// Update computed output subdirectories after OutputDir is set
	constants.OutputGroupsDir = filepath.Join(constants.OutputDir, "groups")
	constants.OutputCategoriesDir = filepath.Join(constants.OutputDir, "categories")
	constants.OutputIgnoredDir = filepath.Join(constants.OutputDir, "ignored")
	constants.OutputTopDir = filepath.Join(constants.OutputDir, "top")
	constants.OutputSummariesDir = filepath.Join(constants.OutputDir, "summaries")
	constants.RefreshDerivedPaths()
}

func resolveProjectRoot() string {
	if isTestContext() {
		if configPath := os.Getenv("DNS_TOOLKIT_TEST_CONFIG_PATH"); configPath != "" {
			if !filepath.IsAbs(configPath) {
				if absPath, err := filepath.Abs(configPath); err == nil {
					configPath = absPath
				}
			}

			configDir := filepath.Dir(configPath)
			if filepath.Base(configDir) == "testdata" {
				return filepath.Dir(configDir)
			}
			return configDir
		}
	}

	projectRoot, err := utils.FindProjectRoot("")
	if err != nil {
		return ""
	}

	return projectRoot
}

func remapDataDirForTests(dir string) string {
	clean := filepath.Clean(dir)
	dataRoot := "data"
	dataPrefix := dataRoot + string(filepath.Separator)

	if clean == dataRoot {
		return "testdata"
	}

	if after, ok := strings.CutPrefix(clean, dataPrefix); ok {
		trimmed := after
		return filepath.Join("testdata", trimmed)
	}

	return dir
}

// InitForTesting initializes directories for testing when cobra.OnInitialize is not called
func InitForTesting() {
	validateAndSetDirs()
}

var rootCmd = &cobra.Command{
	Use:   "dns-toolkit",
	Short: "DNS Toolkit",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if ignoreCmds[cmd.Name()] {
			return
		}
		startTime = time.Now()
		Logger.Infof("Command %s started", cmd.Name())
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if ignoreCmds[cmd.Name()] {
			return
		}
		endTime = time.Now()
		duration := endTime.Sub(startTime)
		Logger.Infof("Command %s completed in %s", cmd.Name(), duration)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func isTestMode() bool {
	return os.Getenv("DNS_TOOLKIT_TEST_MODE") == constants.BooleanTrue
}

func isTestContext() bool {
	return isTestMode() || os.Getenv("DNS_TOOLKIT_TEST_CONFIG_PATH") != ""
}

func shouldValidateSourcesSchemaOnInit() bool {
	if len(os.Args) <= 1 {
		return false
	}

	switch os.Args[1] {
	case validateSourcesCommandName:
		return true
	default:
		return false
	}
}

func init() {
	if isTestMode() && os.Getenv("DNS_TOOLKIT_TEST_CONFIG_PATH") == "" {
		Logger = utils.NewTestLogger()
		AppConfig = &config.AppConfig{
			Application: config.ApplicationConfig{
				Name:        "dns-toolkit-test",
				Version:     "0.0.0-test",
				Description: "Test configuration",
			},
		}
		cobra.OnInitialize(validateAndSetDirs)
		registerCommands()
		return
	}

	configPath, err := GetConfigPath()
	if err != nil {
		slog.Error("Failed to get config path", "error", err)
		os.Exit(1)
	}

	Logger = common.InitLogger(configPath)

	// Skip validation for help commands
	if len(os.Args) <= 1 || (os.Args[1] != "help" && os.Args[1] != "--help" && os.Args[1] != "-h") {
		if err := validateConfigWithSchema(configPath, shouldValidateSourcesSchemaOnInit()); err != nil {
			slog.Error("Config validation failed", "error", err)
			os.Exit(1)
		}
	}

	cobra.OnInitialize(validateAndSetDirs)
	registerCommands()
}

func registerCommands() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(processCmd)
	rootCmd.AddCommand(consolidateCmd)
	rootCmd.AddCommand(sourceTypesSummaryCmd)
	rootCmd.AddCommand(validateSourcesCmd)
	rootCmd.AddCommand(overlapCmd)
	rootCmd.AddCommand(topEntriesCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(archiveCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(schemaCmd)
}
