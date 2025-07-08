package processors

import (
	"strings"

	"github.com/phani-kb/multilog"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/utils"
)

const sourceTypeIpv4Find = "ipv4_find"

type Ipv4FindProcessor struct {
	BaseProcessor
}

func NewIpv4FindProcessor(sourceType, listType string) *Ipv4FindProcessor {
	return &Ipv4FindProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

func (p *Ipv4FindProcessor) Process(_ *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if utils.IsComment(line) {
			continue
		}

		if constants.SourceTypeRegexMap[constants.SourceTypeIpv4].MatchString(line) {
			ip := constants.SourceTypeRegexMap[constants.SourceTypeIpv4].FindString(line)
			validEntries = append(validEntries, ip)
		} else {
			invalidEntries = append(invalidEntries, line)
		}
	}

	return validEntries, invalidEntries
}

func init() {
	RegisterProcessor(sourceTypeIpv4Find, func(st string, lt string) Processor {
		return NewIpv4FindProcessor(st, lt)
	})
}
