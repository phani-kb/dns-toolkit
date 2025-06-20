package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/spf13/cobra"
)

var includeIgnored bool
var deleteFolders bool

// prepareDirectories creates necessary output directories
func prepareDirectories() error {
	// Create an output directory
	if err := os.MkdirAll(constants.OutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	return nil
}

// loadTemplates loads and parses template files
func loadTemplates() (*template.Template, []byte, error) {
	staticTemplatePath := filepath.Join("configs", "templates", "static_template.txt")
	dynamicTemplatePath := filepath.Join("configs", "templates", "dynamic_template.txt")

	staticTemplate, err := os.ReadFile(staticTemplatePath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read static template: %w", err)
	}

	dynamicTemplateTxt, err := os.ReadFile(dynamicTemplatePath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read dynamic template: %w", err)
	}

	tmpl, err := template.New("dynamic").Parse(string(dynamicTemplateTxt))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse dynamic template: %w", err)
	}

	return tmpl, staticTemplate, nil
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate different types of outputs",
	Long:  "Generate various outputs from DNS toolkit data",
}

var generateOutputCmd = &cobra.Command{
	Use:   "output",
	Short: "Generate output files with templates prefixed to them",
	Long:  "Generate output files with static and dynamic templates prefixed to the summary types defined in SummaryTypesWithTemplateMap",
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Info("Starting generate prefixes command...")

		if err := u.EnsureDirectoryExists(Logger, constants.OutputDir); err != nil {
			Logger.Errorf("Failed to create output directory: %v", err)
			os.Exit(1)
		}

		// Prepare directories
		if err := prepareDirectories(); err != nil {
			Logger.Error("Failed to prepare directories", "error", err)
			return
		}

		// Load templates
		_, _, err := loadTemplates()
		if err != nil {
			Logger.Error("Failed to load templates", "error", err)
			return
		}

		Logger.Info("Finished generate prefixes command.")
	},
}

func init() {
	generateCmd.AddCommand(generateOutputCmd)

	// Add the includeIgnored flag
	generateOutputCmd.Flags().BoolVarP(&includeIgnored, "include-ignored", "i", false,
		"Include ignored files in the output by copying to the ignored subfolder")

	// Add the deleteFolders flag
	generateOutputCmd.Flags().BoolVarP(&deleteFolders, "delete-folders", "d", false,
		"Delete folders specified in FoldersToDeleteAfterOutputGeneration after output generation")
}
