package processors

import (
	"regexp"
	"strings"

	"github.com/phani-kb/multilog"

	"github.com/phani-kb/dns-toolkit/internal/utils"
)

const sourceTypeDomainCsvHttpUrlFind = "domain_csv_http_url_find"

type DomainCsvHttpUrlFindProcessor struct {
	BaseProcessor
}

func NewDomainCsvHttpUrlFindProcessor(sourceType, listType string) *DomainCsvHttpUrlFindProcessor {
	return &DomainCsvHttpUrlFindProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

func (p *DomainCsvHttpUrlFindProcessor) Process(_ *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")

	urlRegex := regexp.MustCompile(`https?:\/\/([^\/:\s]+)(?::\d+)?(?:\/.*)?`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if utils.IsComment(line) {
			continue
		}

		// Split the line into CSV fields
		fields := strings.Split(line, ",")
		foundValidDomain := false

		// Check each field for URLs
		for _, field := range fields {
			field = strings.TrimSpace(field)
			// Only check fields that look like URLs (contain ://)
			if strings.Contains(field, "://") {
				matches := urlRegex.FindStringSubmatch(field)
				if len(matches) > 1 {
					hostname := matches[1]
					hostname = strings.TrimRight(hostname, ".")
					// Only consider it valid if it's a domain (not an IP address)
					if utils.IsDomain(hostname) {
						validEntries = append(validEntries, hostname)
						foundValidDomain = true
						break // Found a valid domain in this line, move to next line
					}
				}
			}
		}

		if !foundValidDomain {
			invalidEntries = append(invalidEntries, line)
		}
	}

	return validEntries, invalidEntries
}

func init() {
	RegisterProcessor(sourceTypeDomainCsvHttpUrlFind, func(st string, lt string) Processor {
		return NewDomainCsvHttpUrlFindProcessor(st, lt)
	})
}
