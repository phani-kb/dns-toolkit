package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/spf13/cobra"
)

var generateCreditsCmd = &cobra.Command{
	Use:   "credits",
	Short: "Generate and update source credits in README.md",
	Long:  "Update the CREDITS section in README.md with a collapsible table for each source file.",
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getenv("DNS_TOOLKIT_TEST_MODE") == "true" {
			Logger.Debugf("Skipping generate credits command in test mode")
			return
		}

		configPath, err := GetConfigPath()
		if err != nil {
			Logger.Errorf("Failed to get config path: %v", err)
			os.Exit(1)
		}

		appConfig, sourcesConfigs, err := config.LoadAppConfig(Logger, configPath)
		if err != nil {
			Logger.Errorf("Failed to load config: %v", err)
			os.Exit(1)
		}
		AppConfig = &appConfig
		SourcesConfigs = sourcesConfigs

		if err := updateReadmeWithCredits(Logger, "README.md"); err != nil {
			Logger.Errorf("Failed to update README with credits: %v", err)
			os.Exit(1)
		}

		Logger.Info("Successfully updated README.md with credits")
	},
}

func updateReadmeWithCredits(logger *multilog.Logger, readmePath string) error {
	content, err := os.ReadFile(readmePath)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", readmePath, err)
	}

	creditsContent := generateCreditsSection()

	startMarker := "<!-- CREDITS_START -->"
	endMarker := "<!-- CREDITS_END -->"

	// Check if credits section already exists
	startIndex := strings.Index(string(content), startMarker)
	endIndex := strings.Index(string(content), endMarker)

	var newContent string

	if startIndex != -1 && endIndex != -1 {
		// Replace existing section
		before := string(content)[:startIndex]
		after := string(content)[endIndex+len(endMarker):]
		newContent = before + creditsContent + after
	} else {
		logger.Warnf("Credits section not found in README.md")
	}

	if err := os.WriteFile(readmePath, []byte(newContent), 0o644); err != nil {
		return fmt.Errorf("failed to write %s: %w", readmePath, err)
	}

	return nil
}

func generateCreditsSection() string {
	var sb strings.Builder

	sb.WriteString("<!-- CREDITS_START -->\n")
	sb.WriteString("## Credits\n\n")
	sb.WriteString("This project is made possible by the following blocklist and allowlist sources:\n\n")
	sb.WriteString("Legend: S = Status, C/U/X = Count / Unique / Conflicts\n\n")

	sourcesByFile := make(map[string][]config.Source)

	if AppConfig != nil && len(AppConfig.DNSToolkit.SourceFiles) > 0 {
		for i, sourcesConfig := range SourcesConfigs {
			if i < len(AppConfig.DNSToolkit.SourceFiles) {
				filename := filepath.Base(AppConfig.DNSToolkit.SourceFiles[i])
				sourcesByFile[filename] = sourcesConfig.Sources
			}
		}
	} else {
		sourcesByFile["No sources configured"] = []config.Source{}
	}

	var filenames []string
	for filename := range sourcesByFile {
		filenames = append(filenames, filename)
	}
	utils.SortCaseInsensitiveStrings(filenames)

	overlapMap := make(map[string][3]int)
	overlapFile := filepath.Join(constants.SummaryDir, constants.DefaultSummaryFiles["overlap"])
	if data, err := os.ReadFile(overlapFile); err == nil {
		var overlaps []struct {
			Source    string `json:"source"`
			Count     int    `json:"count"`
			Unique    int    `json:"unique"`
			Conflicts int    `json:"conflicts"`
		}
		if err := json.Unmarshal(data, &overlaps); err == nil {
			for _, o := range overlaps {
				overlapMap[o.Source] = [3]int{o.Count, o.Unique, o.Conflicts}
			}
		}
	}

	// read processed_summary.json for sources not in overlap
	processedFile := filepath.Join(constants.SummaryDir, constants.DefaultSummaryFiles["processed"])
	if data, err := os.ReadFile(processedFile); err == nil {
		var processed []struct {
			Name       string `json:"name"`
			ValidFiles []struct {
				NumberOfEntries int `json:"number_of_entries"`
			} `json:"valid_files"`
			InvalidFiles []struct {
				NumberOfEntries int `json:"number_of_entries"`
			} `json:"invalid_files"`
		}
		if err := json.Unmarshal(data, &processed); err == nil {
			for _, p := range processed {
				// add if not already in overlapMap
				if _, exists := overlapMap[p.Name]; !exists {
					validCount := 0
					for _, vf := range p.ValidFiles {
						validCount += vf.NumberOfEntries
					}
					invalidCount := 0
					for _, inf := range p.InvalidFiles {
						invalidCount += inf.NumberOfEntries
					}
					totalCount := validCount + invalidCount

					if totalCount > 0 {
						// no overlap analysis, show count / 0 / 0
						overlapMap[p.Name] = [3]int{totalCount, 0, 0}
					}
				}
			}
		}
	}

	for _, filename := range filenames {
		sources := sourcesByFile[filename]
		if len(sources) == 0 {
			continue
		}

		sort.Slice(sources, func(i, j int) bool {
			return utils.CaseInsensitiveLess(sources[i].Name, sources[j].Name)
		})

		sb.WriteString("<details>\n")
		_, _ = fmt.Fprintf(&sb, "<summary><strong>📄 %s</strong> (%d sources)</summary>\n\n", filename, len(sources))

		// If overlapMap has entries, replace AL/BL column with Count/Unique/Conflicts column.
		hasOverlap := len(overlapMap) > 0
		if hasOverlap {
			sb.WriteString("| Name | S | Categories |         C / U / X        | Notes |\n")
			sb.WriteString("|------|---|------------|--------------------------|-------|\n")
		} else {
			sb.WriteString("| Name | S | Categories |          AL / BL         | Notes |\n")
			sb.WriteString("|------|---|------------|--------------------------|-------|\n")
		}

		for _, source := range sources {
			name := source.Name
			if source.URL != "" {
				name = source.Name
				// if strings.HasPrefix(source.URL, "file://") {
				// 	name = source.Name
				// } else {
				// 	name = fmt.Sprintf("[%s](%s)", source.Name, source.URL)
				// }
			}

			status := "✅"
			if source.Disabled {
				status = "❌"
			}

			categories := "-"
			if len(source.Categories) > 0 {
				categories = strings.Join(source.Categories, ", ")
			}

			notes := "-"
			if source.Notes != "" {
				notes = strings.ReplaceAll(source.Notes, "|", "\\|")
				notes = strings.ReplaceAll(notes, "\n", " ")
			}

			listTypes := getListTypes(source)
			if hasOverlap {
				if vals, ok := overlapMap[source.Name]; ok {
					_, _ = fmt.Fprintf(&sb, "| %s | %s | %s | %d / %d / %d | %s |\n",
						name, status, categories, vals[0], vals[1], vals[2], notes)
				} else {
					_, _ = fmt.Fprintf(&sb, "| %s | %s | %s | - | %s |\n", name, status, categories, notes)
				}
			} else {
				_, _ = fmt.Fprintf(&sb, "| %s | %s | %s | %s | %s |\n", name, status, categories, listTypes, notes)
			}
		}

		sb.WriteString("\n")
		sb.WriteString("</details>\n\n")
	}

	sb.WriteString("<!-- CREDITS_END -->")

	return sb.String()
}

func getListTypes(source config.Source) []string {
	listTypes := utils.NewStringSet([]string{})
	for _, t := range source.Types {
		for _, lt := range t.ListTypes {
			listTypes.Add(constants.ListTypeMap[lt.Name])
		}
	}
	return listTypes.ToSliceSorted()
}

func init() {
	generateCmd.AddCommand(generateCreditsCmd)
}
