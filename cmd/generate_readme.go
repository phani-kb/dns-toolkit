package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/spf13/cobra"
)

var generateReadmeCmd = &cobra.Command{
	Use:   "output-readme",
	Short: "Generate README.md for output branch with daily workflow summary",
	Long:  "Generate a README.md for the output branch containing tabular summaries of daily workflow runs including download, processing, consolidation, and analysis statistics", // nolint:lll
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getenv("DNS_TOOLKIT_TEST_MODE") == "true" {
			Logger.Debug("Skipping generate readme command in test mode")
			return
		}

		Logger.Info("Generating README.md for output branch...")

		if err := u.EnsureDirectoryExists(Logger, constants.OutputDir); err != nil {
			Logger.Errorf("Failed to create output directory: %v", err)
			os.Exit(1)
		}

		readme := generateOutputBranchReadme()

		readmePath := filepath.Join(constants.OutputDir, "README.md")
		if err := os.WriteFile(readmePath, []byte(readme), 0644); err != nil {
			Logger.Errorf("Failed to write README file: %v", err)
			os.Exit(1)
		}

		Logger.Infof("Successfully generated README.md at %s", readmePath)
	},
}

// WorkflowSummary holds summary data for the entire workflow
type WorkflowSummary struct {
	LastRun     string
	Processing  ProcessingStats
	Groups      GroupsStats
	Categories  CategoriesStats
	Top         TopStats
	Consolidate ConsolidateStats
	Overlap     OverlapStats
	Download    DownloadStats
}

type DownloadStats struct {
	SourcesByType  map[string]int
	LastUpdateTime string
	ErrorSources   []string
	TotalSources   int
	SuccessCount   int
	FailedCount    int
}

type ProcessingStats struct {
	ValidFilesByType   map[string]int
	InvalidFilesByType map[string]int
	LastUpdateTime     string
	TotalSources       int
}

type ConsolidateStats struct {
	FilesByType    map[string]ConsolidateTypeStats
	LastUpdateTime string
	TotalFiles     int
}

type ConsolidateTypeStats struct {
	Blocklist ConsolidateListStats
	Allowlist ConsolidateListStats
}

type ConsolidateListStats struct {
	Count        int
	FilesCount   int
	IgnoredCount int
}

type GroupsStats struct {
	GroupSummary   map[string]int
	GroupListTypes map[string][]string
	LastUpdateTime string
	TotalGroups    int
}

type CategoriesStats struct {
	CategorySummary   map[string]int
	CategoryListTypes map[string][]string
	LastUpdateTime    string
	TotalCategories   int
}

type OverlapStats struct {
	LastUpdateTime string
	TotalAnalyzed  int
}

type TopStats struct {
	FilesByType    map[string]int
	FileDetails    map[string][]TopFileDetail
	LastUpdateTime string
	TotalFiles     int
}

type TopFileDetail struct {
	ListType   string
	MinSources int
	Count      int
}

func generateOutputBranchReadme() string {
	summary := collectWorkflowSummary()

	var sb strings.Builder

	sb.WriteString("# DNS Toolkit - Daily Processing Results\n\n")
	sb.WriteString("This branch contains the daily processed and consolidated DNS blocklists and allowlists.\n\n")
	sb.WriteString(fmt.Sprintf("**Last Updated:** %s\n\n", summary.LastRun))

	sb.WriteString("## Quick Start\n\n")
	sb.WriteString("Add any of these URLs to your DNS filtering solution:\n\n")
	sb.WriteString("```\n")

	topLevelFiles, err := getTopLevelTxtFiles()
	if err == nil && len(topLevelFiles) > 0 {
		sb.WriteString("# Consolidated Blocklists or Allowlists\n")
		for _, filename := range topLevelFiles {
			sb.WriteString(fmt.Sprintf("%s/%s\n", constants.GitHubRawURL, filename))
		}
		sb.WriteString("\n")
	}

	if summary.Groups.TotalGroups > 0 {
		sb.WriteString("# Size-based lists\n")
		var groups []string
		for group := range summary.Groups.GroupSummary {
			groups = append(groups, group)
		}
		sort.Strings(groups)

		for _, group := range groups {
			if listTypes, exists := summary.Groups.GroupListTypes[group]; exists {
				for _, typeListCombination := range listTypes {
					sb.WriteString(
						fmt.Sprintf("%s/groups/%s_%s.txt\n", constants.GitHubRawURL, group, typeListCombination),
					)
				}
			}
		}
		sb.WriteString("\n")
	}

	if summary.Categories.TotalCategories > 0 {
		sb.WriteString("# Category-based lists\n")
		var categories []string
		for category := range summary.Categories.CategorySummary {
			categories = append(categories, category)
		}
		sort.Strings(categories)

		for _, category := range categories {
			if listTypes, exists := summary.Categories.CategoryListTypes[category]; exists {
				for _, typeListCombination := range listTypes {
					sb.WriteString(
						fmt.Sprintf("%s/categories/%s_%s.txt\n", constants.GitHubRawURL, category, typeListCombination),
					)
				}
			}
		}
		sb.WriteString("\n")
	}

	if summary.Top.TotalFiles > 0 {
		sb.WriteString("# High-confidence lists (top entries by number of sources)\n")
		var types []string
		for sourceType := range summary.Top.FilesByType {
			types = append(types, sourceType)
		}
		sort.Strings(types)

		for _, sourceType := range types {
			if details, exists := summary.Top.FileDetails[sourceType]; exists && len(details) > 0 {
				for _, detail := range details {
					sb.WriteString(
						fmt.Sprintf(
							"%s/top/top_%s_%s_min%d.txt\n",
							constants.GitHubRawURL,
							sourceType,
							detail.ListType,
							detail.MinSources,
						),
					)
				}
			}
		}
		sb.WriteString("\n")
	}
	sb.WriteString("```\n\n")

	// Daily Workflow Summary
	sb.WriteString("## Daily Workflow Summary\n\n")

	// Download Summary
	sb.WriteString("### Download Statistics\n\n")
	sb.WriteString("| Metric | Count |\n")
	sb.WriteString("|--------|-------|\n")
	sb.WriteString(fmt.Sprintf("| Total Sources | %d |\n", summary.Download.TotalSources))
	sb.WriteString(fmt.Sprintf("| Successful Downloads | %d |\n", summary.Download.SuccessCount))
	sb.WriteString(fmt.Sprintf("| Failed Downloads | %d |\n", summary.Download.FailedCount))
	sb.WriteString(
		fmt.Sprintf(
			"| Success Rate | %.1f%% |\n",
			float64(summary.Download.SuccessCount)/float64(summary.Download.TotalSources)*100,
		),
	)
	sb.WriteString(fmt.Sprintf("| Last Update | %s |\n", summary.Download.LastUpdateTime))
	sb.WriteString("\n")

	// Sources by Type
	if len(summary.Download.SourcesByType) > 0 {
		sb.WriteString("**Sources by Type:**\n\n")
		sb.WriteString("| Source Type | Count |\n")
		sb.WriteString("|-------------|-------|\n")

		// Sort source types alphabetically
		var sourceTypes []string
		for sourceType := range summary.Download.SourcesByType {
			sourceTypes = append(sourceTypes, sourceType)
		}
		sort.Strings(sourceTypes)

		for _, sourceType := range sourceTypes {
			count := summary.Download.SourcesByType[sourceType]
			sb.WriteString(fmt.Sprintf("| %s | %d |\n", sourceType, count))
		}
		sb.WriteString("\n")
	}

	// Failed Sources
	if len(summary.Download.ErrorSources) > 0 {
		sb.WriteString("**Failed Sources:**\n")
		for _, source := range summary.Download.ErrorSources {
			sb.WriteString(fmt.Sprintf("- %s\n", source))
		}
		sb.WriteString("\n")
	}

	// Processing Summary
	sb.WriteString("### Processing Statistics\n\n")
	sb.WriteString("| Source Type | Valid Files | Invalid Files | Total |\n")
	sb.WriteString("|-------------|-------------|---------------|-------|\n")

	allTypes := make(map[string]bool)
	for t := range summary.Processing.ValidFilesByType {
		allTypes[t] = true
	}
	for t := range summary.Processing.InvalidFilesByType {
		allTypes[t] = true
	}

	var types []string
	for t := range allTypes {
		types = append(types, t)
	}
	sort.Strings(types)

	for _, sourceType := range types {
		valid := summary.Processing.ValidFilesByType[sourceType]
		invalid := summary.Processing.InvalidFilesByType[sourceType]
		total := valid + invalid
		sb.WriteString(fmt.Sprintf("| %s | %d | %d | %d |\n", sourceType, valid, invalid, total))
	}
	sb.WriteString(fmt.Sprintf("| **Last Update** | | | %s |\n", summary.Processing.LastUpdateTime))
	sb.WriteString("\n")

	// Consolidation Summary
	sb.WriteString("### Consolidation Statistics\n\n")
	sb.WriteString("| Type | Blocklist Entries | Allowlist Entries | Total Files |\n")
	sb.WriteString("|------|-------------------|-------------------|-------------|\n")

	for _, sourceType := range types {
		if stats, exists := summary.Consolidate.FilesByType[sourceType]; exists {
			sb.WriteString(fmt.Sprintf("| %s | %s | %s | %d |\n",
				sourceType,
				formatConsolidateCount(stats.Blocklist),
				formatConsolidateCount(stats.Allowlist),
				stats.Blocklist.FilesCount+stats.Allowlist.FilesCount,
			))
		}
	}
	sb.WriteString(fmt.Sprintf("| **Last Update** | | | %s |\n", summary.Consolidate.LastUpdateTime))
	sb.WriteString("\n")

	// Groups Summary
	if summary.Groups.TotalGroups > 0 {
		sb.WriteString("### Size Groups Summary\n\n")
		sb.WriteString("| Group | Total Entries |\n")
		sb.WriteString("|-------|---------------|\n")

		var groups []string
		for group := range summary.Groups.GroupSummary {
			groups = append(groups, group)
		}
		sort.Strings(groups)

		for _, group := range groups {
			count := summary.Groups.GroupSummary[group]
			sb.WriteString(fmt.Sprintf("| %s | %s |\n", group, formatNumber(count)))
		}
		sb.WriteString(fmt.Sprintf("| **Last Update** | %s |\n", summary.Groups.LastUpdateTime))
		sb.WriteString("\n")
	}

	// Categories Summary
	if summary.Categories.TotalCategories > 0 {
		sb.WriteString("### Categories Summary\n\n")
		sb.WriteString("| Category | Total Entries |\n")
		sb.WriteString("|----------|---------------|\n")

		var categories []string
		for category := range summary.Categories.CategorySummary {
			categories = append(categories, category)
		}
		sort.Strings(categories)

		for _, category := range categories {
			count := summary.Categories.CategorySummary[category]
			sb.WriteString(fmt.Sprintf("| %s | %s |\n", category, formatNumber(count)))
		}
		sb.WriteString(fmt.Sprintf("| **Last Update** | %s |\n", summary.Categories.LastUpdateTime))
		sb.WriteString("\n")
	}

	// Overlap Summary
	if summary.Overlap.TotalAnalyzed > 0 {
		sb.WriteString("### Overlap Analysis Summary\n\n")
		sb.WriteString("| Metric | Count |\n")
		sb.WriteString("|--------|-------|\n")
		sb.WriteString(fmt.Sprintf("| Total Sources Analyzed | %d |\n", summary.Overlap.TotalAnalyzed))
		sb.WriteString(fmt.Sprintf("| **Last Update** | %s |\n", summary.Overlap.LastUpdateTime))
		sb.WriteString("\n")
		sb.WriteString("**[View Detailed Overlap Analysis →](overlap.md)**\n\n")
	}

	// Top Entries Summary
	if summary.Top.TotalFiles > 0 {
		sb.WriteString("### Top Entries Summary\n\n")
		sb.WriteString("| Type | List Type | Min Sources | Entries Count | Files Generated |\n")
		sb.WriteString("|------|-----------|-------------|---------------|----------------|\n")

		// Create a sorted list of all entries
		type sortableTopEntry struct {
			sourceType string
			detail     TopFileDetail
		}

		var allEntries []sortableTopEntry
		for _, sourceType := range types {
			if details, exists := summary.Top.FileDetails[sourceType]; exists && len(details) > 0 {
				for _, detail := range details {
					allEntries = append(allEntries, sortableTopEntry{
						sourceType: sourceType,
						detail:     detail,
					})
				}
			}
		}

		// Sort entries: by source type, then by list type (allowlist first), then by min sources desc
		sort.Slice(allEntries, func(i, j int) bool {
			if allEntries[i].sourceType != allEntries[j].sourceType {
				return allEntries[i].sourceType < allEntries[j].sourceType
			}
			if allEntries[i].detail.ListType != allEntries[j].detail.ListType {
				return allEntries[i].detail.ListType == "allowlist"
			}
			return allEntries[i].detail.MinSources > allEntries[j].detail.MinSources
		})

		for _, entry := range allEntries {
			url := fmt.Sprintf("%s/top/top_%s_%s_min%d.txt",
				constants.GitHubRawURL,
				entry.sourceType,
				entry.detail.ListType,
				entry.detail.MinSources,
			)
			minSourcesLink := fmt.Sprintf("[%d](%s)", entry.detail.MinSources, url)
			sb.WriteString(fmt.Sprintf("| %s | %s | %s | %s | 1 |\n",
				entry.sourceType,
				entry.detail.ListType,
				minSourcesLink,
				formatNumber(entry.detail.Count)))
		}
		sb.WriteString(fmt.Sprintf("| **Last Update** | | | | %s |\n", summary.Top.LastUpdateTime))
		sb.WriteString("\n")
	}

	sb.WriteString("## About\n\n")
	sb.WriteString(fmt.Sprintf(
		"These lists are automatically generated daily by the [DNS Toolkit](%s) ",
		constants.GitHubRepoURL,
	))
	sb.WriteString("from multiple reputable sources.\n\n")

	return sb.String()
}

func collectWorkflowSummary() *WorkflowSummary {
	summary := &WorkflowSummary{
		LastRun: time.Now().Format("2006-01-02 15:04:05 UTC"),
	}

	if err := collectDownloadStats(&summary.Download); err != nil {
		Logger.Warnf("Failed to collect download stats: %v", err)
	}

	if err := collectProcessingStats(&summary.Processing); err != nil {
		Logger.Warnf("Failed to collect processing stats: %v", err)
	}

	if err := collectConsolidateStats(&summary.Consolidate); err != nil {
		Logger.Warnf("Failed to collect consolidation stats: %v", err)
	}

	if err := collectGroupsStats(&summary.Groups); err != nil {
		Logger.Warnf("Failed to collect groups stats: %v", err)
	}

	if err := collectCategoriesStats(&summary.Categories); err != nil {
		Logger.Warnf("Failed to collect categories stats: %v", err)
	}

	if err := collectOverlapStats(&summary.Overlap); err != nil {
		Logger.Warnf("Failed to collect overlap stats: %v", err)
	}

	if err := collectTopStats(&summary.Top); err != nil {
		Logger.Warnf("Failed to collect top stats: %v", err)
	}

	return summary
}

func collectDownloadStats(stats *DownloadStats) error {
	summaryFile := filepath.Join(constants.OutputSummariesDir, constants.DefaultSummaryFiles["download"])
	if _, err := os.Stat(summaryFile); os.IsNotExist(err) {
		return fmt.Errorf("download summary file not found")
	}

	content, err := os.ReadFile(summaryFile)
	if err != nil {
		return err
	}

	var downloadSummaries []c.DownloadSummary
	if err := json.Unmarshal(content, &downloadSummaries); err != nil {
		return err
	}

	stats.TotalSources = len(downloadSummaries)
	stats.SourcesByType = make(map[string]int)
	stats.ErrorSources = []string{}

	for _, summary := range downloadSummaries {
		if summary.Error != "" {
			stats.FailedCount++
			stats.ErrorSources = append(stats.ErrorSources, summary.Name)
		} else {
			stats.SuccessCount++
			if summary.LastDownloadTimestamp != "" {
				stats.LastUpdateTime = summary.LastDownloadTimestamp
			}
		}

		// Count by source type
		for _, sourceType := range summary.Types {
			stats.SourcesByType[sourceType.Name]++
		}
	}

	return nil
}

func collectProcessingStats(stats *ProcessingStats) error {
	summaryFile := filepath.Join(constants.OutputSummariesDir, constants.DefaultSummaryFiles["processed"])
	if _, err := os.Stat(summaryFile); os.IsNotExist(err) {
		return fmt.Errorf("processed summary file not found")
	}

	content, err := os.ReadFile(summaryFile)
	if err != nil {
		return err
	}

	var processedSummaries []c.ProcessedSummary
	if err := json.Unmarshal(content, &processedSummaries); err != nil {
		return err
	}

	stats.TotalSources = len(processedSummaries)
	stats.ValidFilesByType = make(map[string]int)
	stats.InvalidFilesByType = make(map[string]int)

	for _, summary := range processedSummaries {
		if summary.LastProcessedTimestamp != "" {
			stats.LastUpdateTime = summary.LastProcessedTimestamp
		}

		for _, validFile := range summary.ValidFiles {
			stats.ValidFilesByType[validFile.GenericSourceType]++
		}

		for _, invalidFile := range summary.InvalidFiles {
			stats.InvalidFilesByType[invalidFile.GenericSourceType]++
		}
	}

	return nil
}

func collectConsolidateStats(stats *ConsolidateStats) error {
	summaryFile := filepath.Join(constants.OutputSummariesDir, constants.DefaultSummaryFiles["consolidated"])
	if _, err := os.Stat(summaryFile); os.IsNotExist(err) {
		return fmt.Errorf("consolidated summary file not found")
	}

	content, err := os.ReadFile(summaryFile)
	if err != nil {
		return err
	}

	var consolidatedSummaries []c.ConsolidatedSummary
	if err := json.Unmarshal(content, &consolidatedSummaries); err != nil {
		return err
	}

	stats.TotalFiles = len(consolidatedSummaries)
	stats.FilesByType = make(map[string]ConsolidateTypeStats)

	for _, summary := range consolidatedSummaries {
		if summary.LastConsolidatedTimestamp != "" {
			stats.LastUpdateTime = summary.LastConsolidatedTimestamp
		}

		typeStats := stats.FilesByType[summary.Type]

		switch summary.ListType {
		case "blocklist":
			typeStats.Blocklist.Count = summary.Count
			typeStats.Blocklist.FilesCount = summary.FilesCount
			typeStats.Blocklist.IgnoredCount = summary.IgnoredEntriesCount
		case "allowlist":
			typeStats.Allowlist.Count = summary.Count
			typeStats.Allowlist.FilesCount = summary.FilesCount
			typeStats.Allowlist.IgnoredCount = summary.IgnoredEntriesCount
		}

		stats.FilesByType[summary.Type] = typeStats
	}

	return nil
}

// collectConsolidatedStats is a generic function to collect stats from consolidated summaries
func collectConsolidatedStats(summaryFileKey string, processFunc func([]byte) error) error {
	summaryFile := filepath.Join(constants.OutputSummariesDir, constants.DefaultSummaryFiles[summaryFileKey])
	if _, err := os.Stat(summaryFile); os.IsNotExist(err) {
		return fmt.Errorf("%s summary file not found", summaryFileKey)
	}

	content, err := os.ReadFile(summaryFile)
	if err != nil {
		return err
	}

	return processFunc(content)
}

// collectConsolidatedStatsGeneric is a generic function to collect stats from consolidated summaries
// It processes summaries and extracts identifier-based statistics (groups or categories)
func collectConsolidatedStatsGeneric(
	summaryFileKey string,
	getIdentifier func(c.ConsolidatedSummary) string,
	setStats func(map[string]int, map[string][]string, int, string),
) error {
	return collectConsolidatedStats(summaryFileKey, func(content []byte) error {
		var consolidatedSummaries []c.ConsolidatedSummary
		if err := json.Unmarshal(content, &consolidatedSummaries); err != nil {
			return err
		}

		summary := make(map[string]int)
		listTypes := make(map[string][]string)
		identifiersSet := make(map[string]bool)
		var lastUpdateTime string

		for _, consolidatedSummary := range consolidatedSummaries {
			if consolidatedSummary.LastConsolidatedTimestamp != "" {
				lastUpdateTime = consolidatedSummary.LastConsolidatedTimestamp
			}

			identifier := getIdentifier(consolidatedSummary)
			if identifier != "" {
				identifiersSet[identifier] = true
				summary[identifier] += consolidatedSummary.Count

				// Track list types for each identifier
				if _, exists := listTypes[identifier]; !exists {
					listTypes[identifier] = []string{}
				}

				// Add list type if not already present
				typeListCombination := fmt.Sprintf("%s_%s", consolidatedSummary.Type, consolidatedSummary.ListType)
				listTypeExists := false
				for _, existingType := range listTypes[identifier] {
					if existingType == typeListCombination {
						listTypeExists = true
						break
					}
				}
				if !listTypeExists {
					listTypes[identifier] = append(listTypes[identifier], typeListCombination)
				}
			}
		}

		// Sort list types for each identifier
		for identifier := range listTypes {
			sort.Strings(listTypes[identifier])
		}

		setStats(summary, listTypes, len(identifiersSet), lastUpdateTime)
		return nil
	})
}

func collectGroupsStats(stats *GroupsStats) error {
	return collectConsolidatedStatsGeneric(
		"consolidated_groups",
		func(summary c.ConsolidatedSummary) string {
			return summary.Group
		},
		func(summary map[string]int, listTypes map[string][]string, total int, lastUpdate string) {
			stats.GroupSummary = summary
			stats.GroupListTypes = listTypes
			stats.TotalGroups = total
			stats.LastUpdateTime = lastUpdate
		},
	)
}

func collectCategoriesStats(stats *CategoriesStats) error {
	return collectConsolidatedStatsGeneric(
		"consolidated_categories",
		func(summary c.ConsolidatedSummary) string {
			return summary.Category
		},
		func(summary map[string]int, listTypes map[string][]string, total int, lastUpdate string) {
			stats.CategorySummary = summary
			stats.CategoryListTypes = listTypes
			stats.TotalCategories = total
			stats.LastUpdateTime = lastUpdate
		},
	)
}

func collectOverlapStats(stats *OverlapStats) error {
	summaryFile := filepath.Join(constants.OutputSummariesDir, constants.DefaultSummaryFiles["overlap"])
	if fileInfo, err := os.Stat(summaryFile); os.IsNotExist(err) {
		return fmt.Errorf("overlap summary file not found")
	} else {
		stats.LastUpdateTime = fileInfo.ModTime().Format("2006-01-02 15:04:05")
	}

	content, err := os.ReadFile(summaryFile)
	if err != nil {
		return err
	}

	var overlapSummaries []c.OverlapSummary
	if err := json.Unmarshal(content, &overlapSummaries); err != nil {
		return err
	}

	stats.TotalAnalyzed = len(overlapSummaries)

	return nil
}

func getTopLevelTxtFiles() ([]string, error) {
	files, err := os.ReadDir(constants.OutputDir)
	if err != nil {
		return nil, err
	}

	var txtFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".txt") {
			txtFiles = append(txtFiles, file.Name())
		}
	}

	sort.Strings(txtFiles)
	return txtFiles, nil
}

func collectTopStats(stats *TopStats) error {
	summaryFile := filepath.Join(constants.OutputSummariesDir, constants.DefaultSummaryFiles["top"])
	if fileInfo, err := os.Stat(summaryFile); os.IsNotExist(err) {
		return fmt.Errorf("top summary file not found")
	} else {
		stats.LastUpdateTime = fileInfo.ModTime().Format("2006-01-02 15:04:05")
	}

	content, err := os.ReadFile(summaryFile)
	if err != nil {
		return err
	}

	var topSummaries []c.TopSummary
	if err := json.Unmarshal(content, &topSummaries); err != nil {
		return err
	}

	stats.TotalFiles = len(topSummaries)
	stats.FilesByType = make(map[string]int)
	stats.FileDetails = make(map[string][]TopFileDetail)

	for _, summary := range topSummaries {
		stats.FilesByType[summary.GenericSourceType]++

		detail := TopFileDetail{
			MinSources: summary.MinSources,
			Count:      summary.Count,
			ListType:   summary.ListType,
		}

		stats.FileDetails[summary.GenericSourceType] = append(
			stats.FileDetails[summary.GenericSourceType],
			detail,
		)
	}

	return nil
}

func formatConsolidateCount(stats ConsolidateListStats) string {
	if stats.Count == 0 {
		return "-"
	}
	result := formatNumber(stats.Count)
	if stats.IgnoredCount > 0 {
		result += fmt.Sprintf(" (-%s ignored)", formatNumber(stats.IgnoredCount))
	}
	return result
}

func formatNumber(n int) string {
	if n >= 1000000 {
		return fmt.Sprintf("%.1fM", float64(n)/1000000)
	} else if n >= 1000 {
		return fmt.Sprintf("%.1fK", float64(n)/1000)
	}
	return fmt.Sprintf("%d", n)
}

func init() {
	generateCmd.AddCommand(generateReadmeCmd)
}
