package processors

import (
	"regexp"
	"strings"

	"github.com/phani-kb/multilog"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/utils"
)

const sourceTypeDomainWithCommentSuffix = "domain_with_comment_suffix"

var supportedDomainWithCommentSuffixListTypes = []string{
	constants.ListTypeAllowlist,
}

type DomainWithCommentSuffixProcessor struct {
	BaseProcessor
}

func NewDomainWithCommentSuffixProcessor(sourceType, listType string) *DomainWithCommentSuffixProcessor {
	return &DomainWithCommentSuffixProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

func (p *DomainWithCommentSuffixProcessor) Process(_ *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")

	domainWithCommentSuffixRegex := `([a-zA-Z0-9.-]+)\s*#.*`

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if utils.IsComment(line) {
			continue
		}

		matches := regexp.MustCompile(domainWithCommentSuffixRegex).FindStringSubmatch(line)
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
	RegisterProcessorTypes(
		sourceTypeDomainWithCommentSuffix,
		supportedDomainWithCommentSuffixListTypes,
		func(st string, lt string) Processor {
			return NewDomainWithCommentSuffixProcessor(st, lt)
		},
	)
}
