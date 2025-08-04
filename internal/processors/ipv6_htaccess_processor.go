package processors

import (
	"slices"
	"strings"

	"github.com/phani-kb/multilog"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
)

const sourceTypeIpv6Htaccess = "ipv6_htaccess"

// Ipv6HtaccessProcessor implements a processor for extracting IPv6 addresses
// from htaccess file content.
type Ipv6HtaccessProcessor struct {
	BaseProcessor
}

func NewIpv6HtaccessProcessor(sourceType, listType string) *Ipv6HtaccessProcessor {
	return &Ipv6HtaccessProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

func (p *Ipv6HtaccessProcessor) Process(_ *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")

	// Filter and process lines
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if u.IsComment(line) {
			continue
		}
		if constants.SourceTypeRegexMap[constants.SourceTypeIpv6].MatchString(line) {
			ip := constants.SourceTypeRegexMap[constants.SourceTypeIpv6].FindString(line)
			validEntries = append(validEntries, ip)
		} else {
			invalidEntries = append(invalidEntries, line)
		}
	}

	validEntries = u.RemoveDuplicates(validEntries)
	slices.Sort(validEntries)

	return validEntries, invalidEntries
}

func init() {
	RegisterProcessor(sourceTypeIpv6Htaccess, func(st string, lt string) Processor {
		return NewIpv6HtaccessProcessor(st, lt)
	})
}
