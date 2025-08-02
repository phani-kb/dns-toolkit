package processors

import (
	"regexp"
	"strings"

	"github.com/phani-kb/multilog"

	"github.com/phani-kb/dns-toolkit/internal/utils"
)

const sourceTypeIpv4CsvHttpUrlFind = "ipv4_csv_http_url_find"

type Ipv4CsvHttpUrlFindProcessor struct {
	BaseProcessor
}

func NewIpv4CsvHttpUrlFindProcessor(sourceType, listType string) *Ipv4CsvHttpUrlFindProcessor {
	return &Ipv4CsvHttpUrlFindProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

func (p *Ipv4CsvHttpUrlFindProcessor) Process(_ *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")

	urlRegex := regexp.MustCompile(`(?:https?:\/\/)?(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})(?::\d+)?(?:\/|$)`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if utils.IsComment(line) {
			continue
		}

		// Split the line into CSV fields
		fields := strings.Split(line, ",")
		foundValidIP := false

		// Check each field for URLs containing IPv4 addresses
		for _, field := range fields {
			field = strings.TrimSpace(field)
			// Only check fields that look like URLs (contain :// or start with an IP pattern)
			if strings.Contains(field, "://") {
				matches := urlRegex.FindStringSubmatch(field)
				if len(matches) > 1 {
					ip := matches[1]
					if utils.IsIP(ip) {
						validEntries = append(validEntries, ip)
						foundValidIP = true
						break // Found a valid IP in this line, move to next line
					}
				}
			}
		}

		if !foundValidIP {
			invalidEntries = append(invalidEntries, line)
		}
	}

	return validEntries, invalidEntries
}

func init() {
	RegisterProcessor(sourceTypeIpv4CsvHttpUrlFind, func(st string, lt string) Processor {
		return NewIpv4CsvHttpUrlFindProcessor(st, lt)
	})
}
