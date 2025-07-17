package processors

import (
	"regexp"
	"strings"

	"github.com/phani-kb/multilog"

	"github.com/phani-kb/dns-toolkit/internal/utils"
)

const sourceTypeIpv4Range = "ipv4_range_expand"

type Ipv4RangeProcessor struct {
	BaseProcessor
}

func NewIpv4RangeProcessor(sourceType, listType string) *Ipv4RangeProcessor {
	return &Ipv4RangeProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

func (p *Ipv4RangeProcessor) Process(logger *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")

	// Regex to match IPv4 range
	ipv4RangeRegex := regexp.MustCompile(`\b\d{1,3}(\.\d{1,3}){3}-\d{1,3}(\.\d{1,3}){3}\b`)
	if ipv4RangeRegex == nil {
		logger.Errorf("IPv4 range regex not found")
		return nil, nil
	}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if utils.IsComment(line) {
			continue
		}

		ipRange := ipv4RangeRegex.FindString(line)
		if ipRange == "" {
			invalidEntries = append(invalidEntries, line)
			continue
		}

		ipList := utils.ExpandIpv4Range(logger, ipRange)
		if len(ipList) == 0 {
			invalidEntries = append(invalidEntries, line)
			continue
		}

		validEntries = append(validEntries, ipList...)
	}

	return validEntries, invalidEntries
}

func init() {
	RegisterProcessor(sourceTypeIpv4Range, func(st string, lt string) Processor {
		return NewIpv4RangeProcessor(st, lt)
	})
}
