package processors

import (
	"regexp"
	"strings"

	"github.com/phani-kb/multilog"

	"github.com/phani-kb/dns-toolkit/internal/utils"
)

const sourceTypeIpv6Find = "ipv6_find"

type Ipv6FindProcessor struct {
	BaseProcessor
}

func NewIpv6FindProcessor(sourceType, listType string) *Ipv6FindProcessor {
	return &Ipv6FindProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

func (p *Ipv6FindProcessor) Process(_ *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")

	// Regex to find potential ipv6 addresses in a line
	ipv6Regex := regexp.MustCompile(`[0-9a-fA-F:]+`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if utils.IsComment(line) {
			continue
		}

		// Find all ipv6 addresses in the line
		matches := ipv6Regex.FindAllString(line, -1)
		found := false

		for _, match := range matches {
			if utils.IsIPv6(match) {
				validEntries = append(validEntries, match)
				found = true
			}
		}

		if !found {
			invalidEntries = append(invalidEntries, line)
		}
	}

	return validEntries, invalidEntries
}

func init() {
	RegisterProcessor(sourceTypeIpv6Find, func(st string, lt string) Processor {
		return NewIpv6FindProcessor(st, lt)
	})
}
