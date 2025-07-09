package processors

import (
	"regexp"
	"strings"

	"github.com/phani-kb/multilog"

	"github.com/phani-kb/dns-toolkit/internal/utils"
)

const sourceTypeIpv4Cidr = "ipv4_cidr_expand"

type Ipv4CidrProcessor struct {
	BaseProcessor
}

func NewIpv4CidrProcessor(sourceType, listType string) *Ipv4CidrProcessor {
	return &Ipv4CidrProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

func (p *Ipv4CidrProcessor) Process(logger *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")

	// Regex to match IPv4 CIDR
	ipv4CidrRegex := regexp.MustCompile(`\b\d{1,3}(\.\d{1,3}){3}/\d{1,2}\b`)
	if ipv4CidrRegex == nil {
		logger.Errorf("IPv4 CIDR regex not found")
		return nil, nil
	}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if utils.IsComment(line) {
			continue
		}

		cidr := ipv4CidrRegex.FindString(line)
		if cidr == "" {
			invalidEntries = append(invalidEntries, line)
			continue
		}

		ipList, err := utils.ExpandIpv4Cidr(logger, cidr)
		if err != nil || len(ipList) == 0 {
			logger.Errorf("Failed to expand CIDR %s: %v", cidr, err)
			invalidEntries = append(invalidEntries, line)
			continue
		}

		validEntries = append(validEntries, ipList...)
	}

	return validEntries, invalidEntries
}

func init() {
	RegisterProcessor(sourceTypeIpv4Cidr, func(st string, lt string) Processor {
		return NewIpv4CidrProcessor(st, lt)
	})
}
