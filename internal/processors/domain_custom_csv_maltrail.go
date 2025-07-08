package processors

import (
	"strings"

	"github.com/phani-kb/multilog"

	"github.com/phani-kb/dns-toolkit/internal/utils"
)

const sourceTypeDomainCustomCsvMaltrail = "domain_custom_csv_maltrail"

type DomainCustomCsvMaltrailProcessor struct {
	BaseProcessor
}

func NewDomainCustomCsvMaltrailProcessor(sourceType, listType string) *DomainCustomCsvMaltrailProcessor {
	return &DomainCustomCsvMaltrailProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

func (p *DomainCustomCsvMaltrailProcessor) Process(_ *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")

	// Skip empty content
	if len(lines) == 0 {
		return validEntries, invalidEntries
	}

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Skip empty lines, comments, and lines starting with /
		if utils.IsComment(line) || strings.HasPrefix(line, "/") {
			continue
		}

		// Split the CSV line and get the first column
		fields := strings.Split(line, ",")
		if len(fields) > 0 {
			entry := strings.TrimSpace(fields[0])

			// Skip empty entries
			if entry == "" {
				continue
			}

			// Check if the entry contains a port number (e.g., :8000, :8889)
			if strings.Contains(entry, ":") {
				invalidEntries = append(invalidEntries, entry)
				continue
			}

			if utils.IsDomain(entry) {
				validEntries = append(validEntries, entry)
			} else {
				invalidEntries = append(invalidEntries, entry)
			}
		}
	}

	return validEntries, invalidEntries
}

func init() {
	RegisterProcessor(sourceTypeDomainCustomCsvMaltrail, func(st string, lt string) Processor {
		return NewDomainCustomCsvMaltrailProcessor(st, lt)
	})
}
