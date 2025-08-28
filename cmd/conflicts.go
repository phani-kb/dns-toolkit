package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

// GenerateConflictReport generates markdown from JSON summary
func GenerateConflictReport(logger *multilog.Logger, summaryFile string) (string, error) {
	logger.Infof("Generating conflicts report from JSON summary...")

	content, err := os.ReadFile(summaryFile)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Infof("No conflicts summary found at %s", summaryFile)
			return "", nil
		}
		return "", fmt.Errorf("failed to read conflicts json: %w", err)
	}

	type overrideRecord struct {
		Entry      string   `json:"entry"`
		Decision   string   `json:"decision"`
		Reason     string   `json:"reason"`
		BlockSrcs  []string `json:"block_sources"`
		AllowSrcs  []string `json:"allow_sources"`
		BlockCount int      `json:"block_count"`
		AllowCount int      `json:"allow_count"`
	}
	var overrides []overrideRecord
	if err = json.Unmarshal(content, &overrides); err != nil {
		logger.Infof("Summary is not overrides JSON; skipping report generation: %s", summaryFile)
		return "", nil
	}

	conflicts := make([]ConflictDetail, 0)
	autoResolved := make([]overrideRecord, 0)
	for _, o := range overrides {
		d := strings.ToLower(o.Decision)
		switch d {
		case "conflict":
			conflicts = append(conflicts, ConflictDetail{
				Entry:        o.Entry,
				BlockSources: o.BlockSrcs,
				AllowSources: o.AllowSrcs,
				BlockCount:   o.BlockCount,
				AllowCount:   o.AllowCount,
			})
		case "allow", "block":
			autoResolved = append(autoResolved, o)
		}
	}

	if len(conflicts) == 0 && len(autoResolved) == 0 {
		logger.Infof("No conflicts or auto-resolved entries in JSON file")
		return "", nil
	}

	sort.Slice(conflicts, func(i, j int) bool {
		if conflicts[i].BlockCount != conflicts[j].BlockCount {
			return conflicts[i].BlockCount > conflicts[j].BlockCount
		}
		return strings.ToLower(conflicts[i].Entry) < strings.ToLower(conflicts[j].Entry)
	})
	sort.Slice(autoResolved, func(i, j int) bool {
		if autoResolved[i].BlockCount != autoResolved[j].BlockCount {
			return autoResolved[i].BlockCount > autoResolved[j].BlockCount
		}
		if autoResolved[i].AllowCount != autoResolved[j].AllowCount {
			return autoResolved[i].AllowCount > autoResolved[j].AllowCount
		}
		return strings.ToLower(autoResolved[i].Entry) < strings.ToLower(autoResolved[j].Entry)
	})

	outDir := constants.OutputDir
	if err = u.EnsureDirectoryExists(logger, outDir); err != nil {
		return "", fmt.Errorf("failed to ensure output dir: %w", err)
	}

	filePath := filepath.Join(outDir, "conflicts.md")
	f, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create conflicts file: %w", err)
	}
	defer u.CloseFile(logger, f)

	bw := bufio.NewWriter(f)
	write := func(s string) error {
		_, err := bw.WriteString(s)
		return err
	}

	timestamp := u.GetTimestamp()
	if err := write("# Conflicts report\n\n"); err != nil {
		return "", fmt.Errorf("failed to write header: %w", err)
	}
	if err := write(fmt.Sprintf("Generated: %s\n\n", timestamp)); err != nil {
		return "", fmt.Errorf("failed to write meta: %w", err)
	}

	if len(conflicts) > 0 {
		// nolint:lll
		if err := write(fmt.Sprintf("<details>\n<summary><strong>Conflicts — total: %d</strong></summary>\n\n", len(conflicts))); err != nil {
			return "", fmt.Errorf("failed to write conflicts section header: %w", err)
		}
		if err := write("| Entry | Blocklist sources (count) | Allowlist sources (count) |\n"); err != nil {
			return "", fmt.Errorf("failed to write table header: %w", err)
		}
		if err := write("|---|---:|---:|\n"); err != nil {
			return "", fmt.Errorf("failed to write table separator: %w", err)
		}
		for _, cinfo := range conflicts {
			bs := strings.Join(cinfo.BlockSources, "<br />")
			as := strings.Join(cinfo.AllowSources, "<br />")
			// nolint:lll
			if err := write(fmt.Sprintf("| %s | %s (%d) | %s (%d) |\n", cinfo.Entry, bs, cinfo.BlockCount, as, cinfo.AllowCount)); err != nil {
				return "", fmt.Errorf("failed to write row: %w", err)
			}
		}
		if err := write("\n</details>\n\n"); err != nil {
			return "", fmt.Errorf("failed to close conflicts details: %w", err)
		}
	} else {
		if err := write("**Conflicts — total: 0**\n\n"); err != nil {
			return "", fmt.Errorf("failed to write no conflicts line: %w", err)
		}
	}

	if len(autoResolved) > 0 {
		// nolint:lll
		if err := write(fmt.Sprintf("<details>\n<summary><strong>Auto-resolved — total: %d</strong></summary>\n\n", len(autoResolved))); err != nil {
			return "", fmt.Errorf("failed to write auto-resolved section header: %w", err)
		}
		// nolint:lll
		if err := write("| Entry | Decision | Reason | Blocklist sources (count) | Allowlist sources (count) |\n"); err != nil {
			return "", fmt.Errorf("failed to write auto table header: %w", err)
		}
		if err := write("|---|---|---|---:|---:|\n"); err != nil {
			return "", fmt.Errorf("failed to write auto table separator: %w", err)
		}
		for _, a := range autoResolved {
			bs := strings.Join(a.BlockSrcs, "<br />")
			as := strings.Join(a.AllowSrcs, "<br />")
			// nolint:lll
			if err := write(fmt.Sprintf("| %s | %s | %s | %s (%d) | %s (%d) |\n", a.Entry, a.Decision, a.Reason, bs, a.BlockCount, as, a.AllowCount)); err != nil {
				return "", fmt.Errorf("failed to write auto row: %w", err)
			}
		}
		if err := write("\n</details>\n\n"); err != nil {
			return "", fmt.Errorf("failed to close auto-resolved details: %w", err)
		}
	} else {
		if err := write("**Auto-resolved — total: 0**\n\n"); err != nil {
			return "", fmt.Errorf("failed to write no auto-resolved line: %w", err)
		}
	}

	if err := bw.Flush(); err != nil {
		return "", fmt.Errorf("failed to flush writer: %w", err)
	}

	logger.Debugf("Conflicts report written to %s", filePath)
	return filePath, nil
}
