package processors

import (
	"github.com/phani-kb/multilog"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/utils"
)

const sourceTypeIpv4FromDomain = "ipv4_from_domain"

var supportedIpv4FromDomainListTypes = []string{
	constants.ListTypeBlocklist,
	constants.ListTypeAllowlist,
}

type Ipv4FromDomainProcessor struct {
	BaseProcessor
}

func NewIpv4FromDomainProcessor(sourceType, listType string) *Ipv4FromDomainProcessor {
	return &Ipv4FromDomainProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

func (p *Ipv4FromDomainProcessor) Process(logger *multilog.Logger, content string) ([]string, []string) {
	validEntries, invalidEntries := utils.ExtractEntriesWithRegex(
		content,
		constants.SourceTypeRegexMap[constants.SourceTypeDomain],
	)

	// resolve IP addresses from domains
	ipAddresses, failedDomains := utils.ResolveDomainsToIPv4(logger, validEntries)
	if len(failedDomains) > 0 {
		logger.Warnf("Failed to resolve %v domains", len(failedDomains))
		invalidEntries = append(invalidEntries, failedDomains...)
	}

	return ipAddresses, invalidEntries
}

func init() {
	RegisterProcessorTypes(
		sourceTypeIpv4FromDomain,
		supportedIpv4FromDomainListTypes,
		func(st string, lt string) Processor {
			return NewIpv4FromDomainProcessor(st, lt)
		},
	)
}
