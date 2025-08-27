package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

type ConflictInfo struct {
	Entry        string
	BlockSources []string
	AllowSources []string
	BlockCount   int
	AllowCount   int
}

// GenerateConflictReport scans processed summaries/consolidated files to find entries
// that appear in both allowlists and blocklists
func GenerateConflictReport(logger *multilog.Logger, processedFiles []c.ProcessedFile) (string, error) {
	logger.Infof("Generating conflicts report...")

	blockMap := make(map[string]map[string]struct{})
	allowMap := make(map[string]map[string]struct{})

	add := func(m map[string]map[string]struct{}, entry, source string) {
		if _, ok := m[entry]; !ok {
			m[entry] = make(map[string]struct{})
		}
		m[entry][source] = struct{}{}
	}

	for _, pf := range processedFiles {
		if pf.Filepath == "" || !pf.Valid { // only valid files
			continue
		}

		entries, _, err := u.ReadEntriesFromFile(logger, pf.Filepath)
		if err != nil {
			logger.Debugf("Skipping file %s due to read error: %v", pf.Filepath, err)
			continue
		}

		sourceLabel := pf.Name // just name
		for _, e := range entries {
			switch pf.ListType {
			case constants.ListTypeBlocklist:
				add(blockMap, e, sourceLabel)
			case constants.ListTypeAllowlist:
				add(allowMap, e, sourceLabel)
			}
		}
	}

	conflicts := make([]ConflictInfo, 0)
	for entry, bsrcs := range blockMap {
		if asrcs, ok := allowMap[entry]; ok {
			ci := ConflictInfo{Entry: entry}
			for s := range bsrcs {
				ci.BlockSources = append(ci.BlockSources, s)
			}
			for s := range asrcs {
				ci.AllowSources = append(ci.AllowSources, s)
			}
			ci.BlockCount = len(ci.BlockSources)
			ci.AllowCount = len(ci.AllowSources)
			conflicts = append(conflicts, ci)
		}
	}

	if len(conflicts) == 0 {
		logger.Infof("No conflicts found")
		return "", nil
	}

	sort.Slice(conflicts, func(i, j int) bool {
		if conflicts[i].BlockCount != conflicts[j].BlockCount {
			return conflicts[i].BlockCount > conflicts[j].BlockCount
		}
		return strings.ToLower(conflicts[i].Entry) < strings.ToLower(conflicts[j].Entry)
	})

	outDir := constants.OutputDir
	if err := u.EnsureDirectoryExists(logger, outDir); err != nil {
		return "", fmt.Errorf("failed to ensure output dir: %w", err)
	}

	fileName := "conflicts.md"
	filePath := filepath.Join(outDir, fileName)
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
	if err := write(fmt.Sprintf("Total conflicts: %d\n\n", len(conflicts))); err != nil {
		return "", fmt.Errorf("failed to write summary: %w", err)
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

	if err := bw.Flush(); err != nil {
		return "", fmt.Errorf("failed to flush writer: %w", err)
	}

	logger.Debugf("Conflicts report written to %s", filePath)
	return filePath, nil
}
