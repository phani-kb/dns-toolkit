package processors

import (
	"slices"
	"strings"

	"github.com/phani-kb/multilog"

	"github.com/phani-kb/dns-toolkit/internal/utils"
)

const sourceTypeAdguardCsvHttpUrlFind = "adguard_csv_http_url_find"

// AdguardCsvHttpUrlFindProcessor extracts http/https URLs from CSV lines and converts
// them into AdGuard block rules
type AdguardCsvHttpUrlFindProcessor struct {
	BaseProcessor
}

func NewAdguardCsvHttpUrlFindProcessor(sourceType, listType string) *AdguardCsvHttpUrlFindProcessor {
	return &AdguardCsvHttpUrlFindProcessor{BaseProcessor: NewBaseProcessor(sourceType, listType)}
}

func (p *AdguardCsvHttpUrlFindProcessor) Process(_ *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")

	unique := make(map[string]struct{})

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if utils.IsComment(line) {
			continue
		}

		fields := strings.Split(line, ",")
		found := false

		for _, field := range fields {
			field = strings.TrimSpace(field)
			if field == "" {
				continue
			}
			// Only consider fields that look like URLs
			if strings.Contains(field, "://") {
				if entry, ok := ParseToAdguardEntry(field); ok {
					unique[entry] = struct{}{}
					found = true
					break // one URL per CSV line
				}
			}
		}

		if !found {
			invalidEntries = append(invalidEntries, line)
		}
	}

	for e := range unique {
		validEntries = append(validEntries, e)
	}
	slices.Sort(validEntries)
	return validEntries, invalidEntries
}

func init() {
	RegisterProcessor(sourceTypeAdguardCsvHttpUrlFind, func(st string, lt string) Processor {
		return NewAdguardCsvHttpUrlFindProcessor(st, lt)
	})
}
