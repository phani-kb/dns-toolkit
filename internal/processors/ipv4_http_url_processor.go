package processors

import (
	"regexp"
	"strings"

	"github.com/phani-kb/multilog"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/utils"
)

const sourceTypeIpv4HttpUrl = "ipv4_http_url"

type Ipv4HttpUrlProcessor struct {
	BaseProcessor
}

func NewIpv4HttpUrlProcessor(sourceType, listType string) *Ipv4HttpUrlProcessor {
	return &Ipv4HttpUrlProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

func isValidIPv4(ip string) bool {
	return utils.IsIP(ip)
}

func (p *Ipv4HttpUrlProcessor) Process(_ *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")

	urlRegex := regexp.MustCompile(`https?://\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}(:\d{1,5})?(?:/\S*)?`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if utils.IsComment(line) {
			continue
		}

		if match := urlRegex.FindString(line); match != "" {
			if ipv4 := constants.SourceTypeRegexMap[constants.SourceTypeIpv4].FindString(match); ipv4 != "" && isValidIPv4(ipv4) {
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
	RegisterProcessor(sourceTypeIpv4HttpUrl, func(st string, lt string) Processor {
		return NewIpv4HttpUrlProcessor(st, lt)
	})
}
