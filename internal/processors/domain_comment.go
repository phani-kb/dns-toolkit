package processors

import (
	"regexp"
	"strings"

	"github.com/phani-kb/multilog"

	"github.com/phani-kb/dns-toolkit/internal/utils"
)

const sourceTypeCommentDomain = "domain_comment"

type DomainCommentProcessor struct {
	BaseProcessor
}

func NewDomainCommentProcessor(sourceType, listType string) *DomainCommentProcessor {
	return &DomainCommentProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

func (p *DomainCommentProcessor) Process(_ *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")

	domainRegex := `#\s*([a-zA-Z0-9.-]+)`

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if utils.IsComment(line) {
			continue
		}

		matches := regexp.MustCompile(domainRegex).FindStringSubmatch(line)
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
	RegisterProcessor(sourceTypeCommentDomain, func(st string, lt string) Processor {
		return NewDomainCommentProcessor(st, lt)
	})
}
