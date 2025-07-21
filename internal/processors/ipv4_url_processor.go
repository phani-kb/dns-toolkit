package processors

import (
	"regexp"
	"strings"

	"github.com/phani-kb/multilog"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/utils"
)

const sourceTypeIpv4Url = "ipv4_url"

type Ipv4UrlProcessor struct {
	BaseProcessor
}

func NewIpv4UrlProcessor(sourceType, listType string) *Ipv4UrlProcessor {
	return &Ipv4UrlProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

func (p *Ipv4UrlProcessor) Process(_ *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")

	urlRegex := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}(:\d{1,5})?/\S*`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if utils.IsComment(line) {
			continue
		}

		if match := urlRegex.FindString(line); match != "" {
			if ipv4 := constants.SourceTypeRegexMap[constants.SourceTypeIpv4].FindString(match); ipv4 != "" &&
				utils.IsIP(ipv4) {
				validEntries = append(validEntries, ipv4)
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
	RegisterProcessor(sourceTypeIpv4Url, func(st string, lt string) Processor {
		return NewIpv4UrlProcessor(st, lt)
	})
}
