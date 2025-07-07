package processors

import (
	"strings"

	"github.com/phani-kb/multilog"

	"github.com/phani-kb/dns-toolkit/internal/utils"
)

const sourceTypeHostname = "hostname"

type HostnameProcessor struct {
	BaseProcessor
}

func NewHostnameProcessor(sourceType, listType string) *HostnameProcessor {
	return &HostnameProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

func (p *HostnameProcessor) Process(_ *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")

	// Skip empty content
	if len(lines) == 0 {
		return validEntries, invalidEntries
	}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if utils.IsComment(line) {
			continue
		}

		// Split by # to remove comments first
		parts := strings.SplitN(line, "#", 2)
		cleanLine := strings.TrimSpace(parts[0])

		// Split remaining line by whitespace
		fields := strings.Fields(cleanLine)
		if len(fields) >= 2 {
			// Skip IP address (fields[0]) and get domain (fields[1])
			domain := strings.TrimSpace(fields[1])

			// Validate and store domain
			if utils.IsDomain(domain) {
				validEntries = append(validEntries, domain)
			} else {
				invalidEntries = append(invalidEntries, line)
			}
		}
	}

	return validEntries, invalidEntries
}

func init() {
	RegisterProcessor(sourceTypeHostname, func(st string, lt string) Processor {
		return NewHostnameProcessor(st, lt)
	})
}
