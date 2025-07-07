package processors

import (
	"regexp"
	"strings"

	"github.com/phani-kb/multilog"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/utils"
)

const sourceTypeTopDomain = "domain_top"

var supportedTopDomainListTypes = []string{
	constants.ListTypeAllowlist,
}

type DomainTopProcessor struct {
	BaseProcessor
}

func NewDomainTopProcessor(sourceType, listType string) *DomainTopProcessor {
	return &DomainTopProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

func (p *DomainTopProcessor) Process(_ *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")

	domainRegex := regexp.MustCompile(`^\d+,\s*(([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,})$`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if utils.IsComment(line) {
			continue
		}

		if match := domainRegex.FindStringSubmatch(line); len(match) > 0 {
			validEntries = append(validEntries, match[1])
		} else {
			invalidEntries = append(invalidEntries, line)
		}
	}

	return validEntries, invalidEntries
}

func init() {
	RegisterProcessorTypes(sourceTypeTopDomain, supportedTopDomainListTypes, func(st string, lt string) Processor {
		return NewDomainTopProcessor(st, lt)
	})
}
