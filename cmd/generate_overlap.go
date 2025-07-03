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

var generateOverlapCmd = &cobra.Command{
	Use:   "overlap",
	Short: "Generate detailed overlap analysis markdown file",
	Long:  "Generate a detailed overlap.md file containing overlap analysis between different sources including overlap percentages, targets, and detailed statistics",
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getenv("DNS_TOOLKIT_TEST_MODE") == "true" {
			Logger.Debug("Skipping generate overlap command in test mode")
			return
		}

		Logger.Info("Generating detailed overlap analysis...")

		if err := u.EnsureDirectoryExists(Logger, constants.OutputDir); err != nil {
			Logger.Errorf("Failed to create output directory: %v", err)
			os.Exit(1)
		}

		overlapMd, err := generateDetailedOverlapAnalysis()
		if err != nil {
			Logger.Errorf("Failed to generate overlap analysis: %v", err)
			os.Exit(1)
		}

		overlapPath := filepath.Join(constants.OutputDir, "overlap.md")
		if err := os.WriteFile(overlapPath, []byte(overlapMd), 0644); err != nil {
			Logger.Errorf("Failed to write overlap file: %v", err)
			os.Exit(1)
		}

		Logger.Infof("Successfully generated overlap.md at %s", overlapPath)
	},
}

type OverlapDetail struct {
	Source     string
	ListType   string
	SourceType string
	Count      int
	Targets    []TargetDetail
	LastUpdate string
}

type TargetDetail struct {
	Name           string
	ListType       string
	SourceType     string
	Count          int
	OverlapCount   int
	OverlapPercent float64
}

func generateDetailedOverlapAnalysis() (string, error) {
	summaryFile := filepath.Join(constants.SummaryDir, constants.DefaultSummaryFiles["overlap"])
	if _, err := os.Stat(summaryFile); os.IsNotExist(err) {
		return "", fmt.Errorf("overlap summary file not found")
	}

	content, err := os.ReadFile(summaryFile)
	if err != nil {
		return "", err
	}

	var overlapSummaries []c.OverlapSummary
	if err := json.Unmarshal(content, &overlapSummaries); err != nil {
		return "", err
	}

	if len(overlapSummaries) == 0 {
		return "", fmt.Errorf("no overlap data found")
	}

	var sb strings.Builder

	sb.WriteString("# DNS Toolkit - Detailed Overlap Analysis\n\n")
	sb.WriteString("This document provides comprehensive overlap analysis between different DNS sources, ")
	sb.WriteString("showing how entries are shared across blocklists and allowlists.\n\n")
	sb.WriteString(fmt.Sprintf("**Last Updated:** %s\n\n", time.Now().Format("2006-01-02 15:04:05 UTC")))

	// Summary overview
	sb.WriteString("## Overview\n\n")
	sb.WriteString("| Metric | Value |\n")
	sb.WriteString("|--------|-------|\n")
	sb.WriteString(fmt.Sprintf("| Total Sources Analyzed | %d |\n", len(overlapSummaries)))

	// Calculate totals
	totalEntries := 0
	listTypeCount := make(map[string]int)
	sourceTypeCount := make(map[string]int)

	for _, summary := range overlapSummaries {
		totalEntries += summary.Count
		listTypeCount[summary.ListType]++
		sourceTypeCount[summary.Type]++
	}

	sb.WriteString(fmt.Sprintf("| Total Entries Analyzed | %s |\n", formatNumber(totalEntries)))
	sb.WriteString("\n")

	if len(listTypeCount) > 0 {
		sb.WriteString("**Sources by List Type:**\n\n")
		sb.WriteString("| List Type | Count |\n")
		sb.WriteString("|-----------|-------|\n")
		for listType, count := range listTypeCount {
			sb.WriteString(fmt.Sprintf("| %s | %d |\n", listType, count))
		}
		sb.WriteString("\n")
	}

	if len(sourceTypeCount) > 0 {
		sb.WriteString("**Sources by Type:**\n\n")
		sb.WriteString("| Source Type | Count |\n")
		sb.WriteString("|-------------|-------|\n")
		for sourceType, count := range sourceTypeCount {
			sb.WriteString(fmt.Sprintf("| %s | %d |\n", sourceType, count))
		}
		sb.WriteString("\n")
	}

	// Detailed analysis per source
	sb.WriteString("## Detailed Source Analysis\n\n")

	sort.Slice(overlapSummaries, func(i, j int) bool {
		return overlapSummaries[i].Source < overlapSummaries[j].Source
	})

	for _, summary := range overlapSummaries {
		sb.WriteString(fmt.Sprintf("### %s\n\n", summary.Source))

		// Source details
		sb.WriteString("| Property | Value |\n")
		sb.WriteString("|----------|-------|\n")
		sb.WriteString(fmt.Sprintf("| List Type | %s |\n", summary.ListType))
		sb.WriteString(fmt.Sprintf("| Source Type | %s |\n", summary.Type))
		sb.WriteString(fmt.Sprintf("| Total Entries | %s |\n", formatNumber(summary.Count)))
		sb.WriteString(fmt.Sprintf("| Unique Entries | %s |\n", formatNumber(summary.Unique)))
		sb.WriteString(fmt.Sprintf("| Target Sources | %d |\n", summary.TargetsCount))
		sb.WriteString("\n")

		// Targets analysis
		if len(summary.TargetsList) > 0 {
			sb.WriteString("**Overlap with Other Sources:**\n\n")
			sb.WriteString("| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |\n")
			sb.WriteString("|---------------|-----------|-------------|--------------|---------------|----------|\n")

			for _, targetStr := range summary.TargetsList {
				target := parseTargetString(targetStr)
				if target != nil {
					sb.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %.1f%% |\n",
						target.Name,
						target.ListType,
						target.SourceType,
						formatNumber(target.Count),
						formatNumber(target.OverlapCount),
						target.OverlapPercent))
				}
			}
			sb.WriteString("\n")
		} else {
			sb.WriteString("*No overlaps found with other sources.*\n\n")
		}

		sb.WriteString("---\n\n")
	}

	sb.WriteString("## About\n\n")
	sb.WriteString(
		"This overlap analysis is automatically generated by the [DNS Toolkit](https://github.com/phani-kb/dns-toolkit) ",
	)
	sb.WriteString("to help understand relationships between different DNS sources. ")
	sb.WriteString(
		"High overlap percentages may indicate redundant sources, while low overlap percentages suggest unique content.\n\n",
	)
	sb.WriteString("**Note:** Overlap percentages are calculated as: (overlap_count / source_total_count) × 100\n\n")

	return sb.String(), nil
}

// parseTargetString parses target strings like "abpvn_hosts, lt: blocklist, type: adguard, count: 1051, overlap: 10, percent: 1.0"
func parseTargetString(target string) *TargetDetail {
	parts := strings.Split(target, ", ")
	if len(parts) < 6 {
		return nil
	}

	detail := &TargetDetail{
		Name: strings.TrimSpace(parts[0]),
	}

	for _, part := range parts[1:] {
		kv := strings.SplitN(strings.TrimSpace(part), ": ", 2)
		if len(kv) != 2 {
			continue
		}

		key := strings.TrimSpace(kv[0])
		value := strings.TrimSpace(kv[1])

		switch key {
		case "lt":
			detail.ListType = value
		case "type":
			detail.SourceType = value
		case "count":
			if count, err := parseIntFromString(value); err == nil {
				detail.Count = count
			}
		case "overlap":
			if overlap, err := parseIntFromString(value); err == nil {
				detail.OverlapCount = overlap
			}
		case "percent":
			if percent, err := parseFloatFromString(value); err == nil {
				detail.OverlapPercent = percent
			}
		}
	}

	return detail
}

func parseIntFromString(s string) (int, error) {
	var result int
	_, err := fmt.Sscanf(s, "%d", &result)
	return result, err
}

func parseFloatFromString(s string) (float64, error) {
	var result float64
	_, err := fmt.Sscanf(s, "%f", &result)
	return result, err
}

func init() {
	generateCmd.AddCommand(generateOverlapCmd)
}
