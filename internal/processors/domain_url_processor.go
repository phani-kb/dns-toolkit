package processors

import (
	"regexp"
	"strings"

	"github.com/phani-kb/multilog"

	"github.com/phani-kb/dns-toolkit/internal/utils"
)

const sourceTypeDomainUrl = "domain_url"

type DomainUrlProcessor struct {
	BaseProcessor
}

func NewDomainUrlProcessor(sourceType, listType string) *DomainUrlProcessor {
	return &DomainUrlProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

func (p *DomainUrlProcessor) Process(_ *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")

	urlRegex := regexp.MustCompile(`^([a-zA-Z0-9.-]+\.[a-zA-Z]{2,})(?::\d+)?(?:/.*)?$`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if utils.IsComment(line) {
			continue
		}

		matches := urlRegex.FindStringSubmatch(line)
		if len(matches) > 1 {
			domain := matches[1]
			// Remove any trailing dots
			domain = strings.TrimRight(domain, ".")
			if utils.IsDomain(domain) {
				validEntries = append(validEntries, domain)
			} else {
				invalidEntries = append(invalidEntries, line)
			}
		} else {
			invalidEntries = append(invalidEntries, line)
		}
	}

	return validEntries, invalidEntries
}

func init() {
	RegisterProcessor(sourceTypeDomainUrl, func(st string, lt string) Processor {
		return NewDomainUrlProcessor(st, lt)
	})
}
