package processors

import (
	"strings"

	"github.com/phani-kb/multilog"

	"github.com/phani-kb/dns-toolkit/internal/constants"
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

		cleanLine := line
		minIndex := len(line)
		for _, prefix := range constants.CommentPrefixes {
			if prefix == "--" {
				// only match -- if it's preceded by whitespace
				for i := 1; i < len(line); i++ {
					if line[i-1] == ' ' || line[i-1] == '\t' {
						if strings.HasPrefix(line[i:], "--") {
							if i < minIndex {
								minIndex = i
							}
							break
						}
					}
				}
			} else {
				if idx := strings.Index(line, prefix); idx != -1 && idx < minIndex {
					minIndex = idx
				}
			}
		}
		if minIndex < len(line) {
			cleanLine = strings.TrimSpace(line[:minIndex])
		}

		// Split remaining line by whitespace
		fields := strings.Fields(cleanLine)
		if len(fields) >= 2 {
			// Skip IP address (fields[0]) and get domain (fields[1])
			domain := strings.TrimSpace(fields[1])

			// Validate and store domain
			if constants.HostnameRegex.MatchString(domain) &&
				len(domain) <= constants.MaxDomainLength {
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
