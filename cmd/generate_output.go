package cmd

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/db"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/spf13/cobra"
)

var includeIgnored bool

// prepareDirectories creates necessary output directories
func prepareDirectories() error {
	// Create an output directory
	if err := os.MkdirAll(constants.OutputDir, 0o755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create subdirectories for different consolidated types
	for summaryType, dirName := range constants.SummaryTypesOutputDirMap {
		if constants.SummaryTypesOutputToSkipMap[summaryType] {
			continue
		}

		dir := filepath.Join(dirName)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("failed to create directory for summary type %s: %w", summaryType, err)
		}
	}

	// Create an ignored directory if the flag is set
	if includeIgnored {
		if err := os.MkdirAll(constants.OutputIgnoredDir, 0o755); err != nil {
			return fmt.Errorf("failed to create ignored directory: %w", err)
		}
	}

	// Create a summaries directory
	if err := os.MkdirAll(constants.OutputSummariesDir, 0o755); err != nil {
		return fmt.Errorf("failed to create summaries directory: %w", err)
	}

	// Create an archive directory
	if err := os.MkdirAll(constants.ArchiveDir, 0o755); err != nil {
		return fmt.Errorf("failed to create archive directory: %w", err)
	}

	return nil
}

// loadTemplates loads and parses template files
func loadTemplates() (*template.Template, []byte, error) {
	configsDir := "configs"

	if os.Getenv("DNS_TOOLKIT_TEST_MODE") == constants.BooleanTrue {
		if projectRoot, err := u.FindProjectRoot(""); err == nil {
			configsDir = filepath.Join(projectRoot, "configs")
		}
	}

	staticTemplatePath := filepath.Join(configsDir, "templates", "static_template.txt")
	dynamicTemplatePath := filepath.Join(configsDir, "templates", "dynamic_template.txt")

	staticTemplate, err := os.ReadFile(staticTemplatePath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read static template: %w", err)
	}

	dynamicTemplateTxt, err := os.ReadFile(dynamicTemplatePath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read dynamic template: %w", err)
	}

	funcMap := template.FuncMap{
		"subtract": func(a, b int) int { return a - b },
	}

	tmpl, err := template.New("dynamic").Funcs(funcMap).Parse(string(dynamicTemplateTxt))
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
	Long:  "Generate output files with static and dynamic templates prefixed",
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getenv("DNS_TOOLKIT_TEST_MODE") == "true" {
			return
		}

		Logger.Info("Starting generate output command...")
		ctx := context.Background()

		if err := u.EnsureDirectoryExists(Logger, constants.OutputDir); err != nil {
			Logger.Errorf("Failed to create output directory: %v", err)
			os.Exit(1)
		}

		if err := prepareDirectories(); err != nil {
			Logger.Error("Failed to prepare directories", "error", err)
			return
		}

		tmpl, staticTemplate, err := loadTemplates()
		if err != nil {
			Logger.Error("Failed to load templates", "error", err)
			return
		}

		database := openDB(ctx)
		defer database.CloseLogError(Logger)

		consolidatedRepo := db.NewConsolidatedRepo(database)
		topRepo := db.NewTopEntriesRepo(database)

		totalFiles := 0

		consolidationTypes := []struct {
			name      string
			outputDir string
		}{
			{"general", constants.OutputDir},
			{"group", constants.OutputGroupsDir},
			{"category", constants.OutputCategoriesDir},
		}

		for _, ct := range consolidationTypes {
			groups, gErr := consolidatedRepo.ListConsolidatedGroups(ct.name)
			if gErr != nil {
				Logger.Errorf("Failed to list consolidated groups for %s: %v", ct.name, gErr)
				continue
			}

			if err := u.EnsureDirectoryExists(Logger, ct.outputDir); err != nil {
				Logger.Errorf("Failed to create output directory %s: %v", ct.outputDir, err)
				continue
			}

			for _, group := range groups {
				if !group.Valid {
					continue // skip invalid entries for output
				}

				entries, eErr := consolidatedRepo.GetConsolidatedEntriesByGroup(
					group.GenericSourceType, group.ListType, ct.name,
					group.GroupName, group.Category, group.Valid,
				)
				if eErr != nil {
					Logger.Errorf("Failed to get entries for %s/%s: %v",
						group.GenericSourceType, group.ListType, eErr)
					continue
				}

				if len(entries) == 0 {
					continue
				}

				fileName := buildOutputFileName(ct.name, group)
				outputPath := filepath.Join(ct.outputDir, fileName)
				description := buildOutputDescription(ct.name, group)

				content := strings.Join(entries, "\n") + "\n"

				if err := writeOutputFile(
					tmpl, staticTemplate, fileName, description,
					group.Count, content, outputPath,
				); err != nil {
					Logger.Errorf("Failed to write output file %s: %v", outputPath, err)
					continue
				}

				totalFiles++
			}
		}

		topGroups, tErr := topRepo.ListTopEntryGroups(ctx)
		if tErr != nil {
			Logger.Errorf("Failed to list top entry groups: %v", tErr)
		} else if len(topGroups) > 0 {
			if err := u.EnsureDirectoryExists(Logger, constants.OutputTopDir); err != nil {
				Logger.Errorf("Failed to create top output directory: %v", err)
			} else {
				for _, tg := range topGroups {
					entries, eErr := topRepo.GetTopEntriesList(ctx, tg.GenericSourceType, tg.ListType, tg.MinSources)
					if eErr != nil {
						Logger.Errorf("Failed to get top entries for %s/%s min%d: %v",
							tg.GenericSourceType, tg.ListType, tg.MinSources, eErr)
						continue
					}
					if len(entries) == 0 {
						continue
					}

					fileName := fmt.Sprintf("top_%s_%s_min%d.txt", tg.GenericSourceType, tg.ListType, tg.MinSources)
					outputPath := filepath.Join(constants.OutputTopDir, fileName)
					description := fmt.Sprintf("Top %s %s (min %d sources)", tg.GenericSourceType, tg.ListType, tg.MinSources)

					content := strings.Join(entries, "\n") + "\n"

					if err := writeOutputFile(
						tmpl, staticTemplate, fileName, description,
						tg.Count, content, outputPath,
					); err != nil {
						Logger.Errorf("Failed to write top output file %s: %v", outputPath, err)
						continue
					}

					totalFiles++
				}
			}
		}

		Logger.Infof("Generated %d output files", totalFiles)
	},
}

// buildOutputFileName generates the output filename for a consolidated group.
func buildOutputFileName(consolidationType string, group db.ConsolidatedGroup) string {
	base := group.GenericSourceType + "_" + group.ListType
	switch consolidationType {
	case "group":
		if group.GroupName != "" {
			base = group.GroupName + "_" + base
		}
	case "category":
		if group.Category != "" {
			base = group.Category + "_" + base
		}
	}
	return base + ".txt"
}

// buildOutputDescription generates a human-readable description for the output file.
func buildOutputDescription(consolidationType string, group db.ConsolidatedGroup) string {
	desc := fmt.Sprintf("%s %s", group.GenericSourceType, group.ListType)
	switch consolidationType {
	case "group":
		if group.GroupName != "" {
			desc = fmt.Sprintf("%s [%s]", desc, group.GroupName)
		}
	case "category":
		if group.Category != "" {
			desc = fmt.Sprintf("%s [%s]", desc, group.Category)
		}
	}
	return desc
}

// writeOutputFile writes an output file with template headers prepended to the content.
func writeOutputFile(
	tmpl *template.Template,
	staticTemplate []byte,
	fileName, description string,
	count int,
	content string,
	outputPath string,
) error {
	var dynamicOutput bytes.Buffer
	err := tmpl.Execute(&dynamicOutput, common.TemplateData{
		AppName:        AppConfig.Application.Name,
		AppVersion:     AppConfig.Application.Version,
		AppDescription: AppConfig.Application.Description,
		FileName:       fileName,
		LastUpdated:    time.Now().Format(constants.TimestampFormat),
		Description:    description,
		Count:          count,
	})
	if err != nil {
		return fmt.Errorf("executing dynamic template: %w", err)
	}

	output := fmt.Sprintf("%s\n%s\n%s\n%s",
		dynamicOutput.String(),
		string(staticTemplate),
		constants.ContentSeparator,
		content,
	)

	if err := os.WriteFile(outputPath, []byte(output), 0o644); err != nil {
		return fmt.Errorf("writing output file: %w", err)
	}

	return nil
}

func init() {
	generateCmd.AddCommand(generateOutputCmd)
}
