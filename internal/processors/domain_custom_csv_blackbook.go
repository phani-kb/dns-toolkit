package processors

import (
	"strings"

	"github.com/phani-kb/multilog"

	"github.com/phani-kb/dns-toolkit/internal/utils"
)

const sourceTypeDomainCustomCsvBlackbook = "domain_custom_csv_blackbook"

type DomainCustomCsvBlackbookProcessor struct {
	BaseProcessor
}

func NewDomainCustomCsvBlackbookProcessor(sourceType, listType string) *DomainCustomCsvBlackbookProcessor {
	return &DomainCustomCsvBlackbookProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

func (p *DomainCustomCsvBlackbookProcessor) Process(_ *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")

	// Skip empty content
	if len(lines) == 0 {
		return validEntries, invalidEntries
	}

	// Skip header line
	for i, line := range lines {
		line = strings.TrimSpace(line)

		// Skip header, empty lines, and comments
		if i == 0 || utils.IsComment(line) {
			continue
		}

		// Split the CSV line and get the first column (domain)
		fields := strings.Split(line, ",")
		if len(fields) > 0 {
			domain := strings.TrimSpace(fields[0])
			if utils.IsDomain(domain) {
				validEntries = append(validEntries, domain)
			} else {
				invalidEntries = append(invalidEntries, domain)
			}
		}
	}

	return validEntries, invalidEntries
}

func init() {
	RegisterProcessor(sourceTypeDomainCustomCsvBlackbook, func(st string, lt string) Processor {
		return NewDomainCustomCsvBlackbookProcessor(st, lt)
	})
}
