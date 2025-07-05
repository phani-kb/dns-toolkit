package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type SourceStats struct {
	LastUpdated      string
	Categories       []string
	SourceTypes      []string
	Countries        []string
	TotalSources     int
	EnabledSources   int
	DisabledSources  int
	BlocklistSources int
	AllowlistSources int
}

var generateStatsCmd = &cobra.Command{
	Use:   "stats-readme",
	Short: "Generate and update source statistics in README.md",
	Long:  "Analyze source configuration files and update the project README.md with comprehensive statistics including source counts, categories, types, and geographic coverage",
	Run: func(cmd *cobra.Command, args []string) {
		// Skip execution in test mode
		if os.Getenv("DNS_TOOLKIT_TEST_MODE") == "true" {
			Logger.Debugf("Skipping generate stats-readme command in test mode")
			return
		}

		stats := collectSourceStats()

		if err := updateReadmeWithStats(stats, "README.md"); err != nil {
			Logger.Errorf("Failed to update README with statistics: %v", err)
			os.Exit(1)
		}

		Logger.Infof("Successfully updated README.md with source statistics")
		Logger.Infof("Total sources: %d (%d enabled, %d disabled)",
			stats.TotalSources, stats.EnabledSources, stats.DisabledSources)
		Logger.Infof("Categories: %d, Source types: %d, Countries: %d",
			len(stats.Categories), len(stats.SourceTypes), len(stats.Countries))
	},
}

func collectSourceStats() *SourceStats {
	stats := &SourceStats{
		LastUpdated: time.Now().Format("2006-01-02 15:04:05 UTC"),
	}

	// Collect categories, source types, and countries from all sources
	categoriesMap := make(map[string]bool)
	sourceTypesMap := make(map[string]bool)
	countriesMap := make(map[string]bool)

	// Process all loaded source configurations
	for _, sourcesConfig := range SourcesConfigs {
		for _, source := range sourcesConfig.Sources {
			stats.TotalSources++

			// Check if source is disabled
			if source.Disabled {
				stats.DisabledSources++
			} else {
				stats.EnabledSources++
			} // Collect categories
			for _, cat := range source.Categories {
				cat = strings.TrimSpace(cat)
				if cat != "" {
					categoriesMap[cat] = true
				}
			}

			// Collect countries
			for _, country := range source.Countries {
				country = strings.TrimSpace(country)
				if country != "" {
					countriesMap[country] = true
				}
			}

			// Collect source types and count list types
			hasBlocklist := false
			hasAllowlist := false

			for _, sourceType := range source.Types {
				sourceTypesMap[sourceType.Name] = true

				for _, listType := range sourceType.ListTypes {
					switch listType.Name {
					case "blocklist":
						hasBlocklist = true
					case "allowlist":
						hasAllowlist = true
					}
				}
			}

			// Count sources that provide blocklists/allowlists
			if hasBlocklist {
				stats.BlocklistSources++
			}
			if hasAllowlist {
				stats.AllowlistSources++
			}
		}
	}

	// Convert maps to sorted slices
	for category := range categoriesMap {
		stats.Categories = append(stats.Categories, category)
	}
	sort.Strings(stats.Categories)

	for sourceType := range sourceTypesMap {
		stats.SourceTypes = append(stats.SourceTypes, sourceType)
	}
	sort.Strings(stats.SourceTypes)

	for country := range countriesMap {
		stats.Countries = append(stats.Countries, country)
	}
	sort.Strings(stats.Countries)

	return stats
}

func updateReadmeWithStats(stats *SourceStats, readmePath string) error {
	content, err := os.ReadFile(readmePath)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", readmePath, err)
	}

	readmeContent := string(content)

	statsSection := generateStatsSection(stats)

	startMarker := "<!-- STATS_START -->"
	endMarker := "<!-- STATS_END -->"

	// Check if statistics section already exists
	startIndex := strings.Index(readmeContent, startMarker)
	endIndex := strings.Index(readmeContent, endMarker)

	var newContent string

	if startIndex != -1 && endIndex != -1 {
		// Replace existing section
		before := readmeContent[:startIndex]
		after := readmeContent[endIndex+len(endMarker):]
		newContent = before + statsSection + after
	} else {
		// Add new section before "## Published Outputs"
		publishedOutputsIndex := strings.Index(readmeContent, "## Published Outputs")
		if publishedOutputsIndex == -1 {
			return fmt.Errorf("could not find '## Published Outputs' section in README.md")
		}

		before := readmeContent[:publishedOutputsIndex]
		after := readmeContent[publishedOutputsIndex:]
		newContent = before + statsSection + "\n" + after
	}

	// Write updated content back to README.md
	if err := os.WriteFile(readmePath, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("failed to write %s: %w", readmePath, err)
	}

	return nil
}

func generateStatsSection(stats *SourceStats) string {
	var sb strings.Builder

	sb.WriteString("<!-- STATS_START -->\n")
	sb.WriteString("## Source Statistics\n\n")
	sb.WriteString("*Automatically generated statistics from source configuration files*\n\n")
	sb.WriteString("| Metric | Count | Details |\n")
	sb.WriteString("|--------|-------|---------|\n")
	sb.WriteString(fmt.Sprintf("| **Total Sources** | %d | %d enabled, %d disabled |\n",
		stats.TotalSources, stats.EnabledSources, stats.DisabledSources))
	sb.WriteString(fmt.Sprintf("| **Blocklist Sources** | %d | Sources providing blocking rules |\n",
		stats.BlocklistSources))
	sb.WriteString(fmt.Sprintf("| **Allowlist Sources** | %d | Sources providing exception rules |\n",
		stats.AllowlistSources))
	sb.WriteString(fmt.Sprintf("| **Categories** | %d | %s |\n",
		len(stats.Categories), strings.Join(stats.Categories, ", ")))
	sb.WriteString(fmt.Sprintf("| **Source Types** | %d | %s |\n",
		len(stats.SourceTypes), strings.Join(stats.SourceTypes, ", ")))
	sb.WriteString(fmt.Sprintf("| **Geographic Coverage** | %d countries | %s |\n",
		len(stats.Countries), strings.Join(stats.Countries, ", ")))
	sb.WriteString(fmt.Sprintf("| **Last Updated** | %s | Statistics generation time |\n",
		stats.LastUpdated))
	sb.WriteString("\n")

	sb.WriteString("<!-- STATS_END -->")

	return sb.String()
}

func init() {
	generateCmd.AddCommand(generateStatsCmd)
}
