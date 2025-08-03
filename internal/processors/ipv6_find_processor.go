package processors

import (
	"strings"

	"github.com/phani-kb/multilog"

	"github.com/phani-kb/dns-toolkit/internal/constants"
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

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if utils.IsComment(line) {
			continue
		}

		if constants.SourceTypeRegexMap[constants.SourceTypeIpv6].MatchString(line) {
			ip := constants.SourceTypeRegexMap[constants.SourceTypeIpv6].FindString(line)
			validEntries = append(validEntries, ip)
		} else {
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
