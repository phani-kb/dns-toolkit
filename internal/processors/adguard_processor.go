package processors

import (
	"slices"
	"strings"

	"github.com/phani-kb/multilog"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
)

const (
	sourceTypeAdguard        = "adguard"
	adguardSourceTypeDomain  = "domain_adguard"
	sourceTypeAdguardDomain  = "adguard_domain"
	sourceTypeAdguardHttpUrl = "adguard_http_url"
)

// AdGuardBlocklistProcessor handles AdGuard blocklist entries
type AdGuardBlocklistProcessor struct {
	BaseProcessor
}

// NewAdGuardBlocklistProcessor creates a new AdGuard blocklist processor
func NewAdGuardBlocklistProcessor(sourceType, listType string) *AdGuardBlocklistProcessor {
	return &AdGuardBlocklistProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

// Process processes AdGuard blocklist entries
func (p *AdGuardBlocklistProcessor) Process(_ *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if u.IsComment(line) {
			continue
		}

		if !strings.HasPrefix(line, AdguardExceptionPrefix) {
			validEntries = append(validEntries, line)
		} else {
			// Matching exception format
			invalidEntries = append(invalidEntries, line)
		}
	}
	return validEntries, invalidEntries
}

// AdGuardAllowlistProcessor handles AdGuard allowlist entries (different format)
type AdGuardAllowlistProcessor struct {
	BaseProcessor
}

// NewAdGuardAllowlistProcessor creates a new AdGuard allowlist processor
func NewAdGuardAllowlistProcessor(sourceType, listType string) *AdGuardAllowlistProcessor {
	return &AdGuardAllowlistProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

// Process processes AdGuard allowlist entries
func (p *AdGuardAllowlistProcessor) Process(_ *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if u.IsComment(line) {
			continue
		}

		if strings.HasPrefix(line, AdguardExceptionPrefix) {
			validEntries = append(validEntries, line)
		} else {
			// Non-matching format for allowlist
			invalidEntries = append(invalidEntries, line)
		}
	}
	return validEntries, invalidEntries
}

func ExtractAllowlistDomains(logger *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if u.IsComment(line) {
			continue
		}

		if domain, ok := strings.CutPrefix(line, AdguardExceptionPrefix); ok {
			// extract domain between @@|| and ^
			if domain, ok := strings.CutPrefix(domain, AdguardBlockPrefix); ok {
				if domain, ok := strings.CutSuffix(domain, AdguardBlockSuffix); ok {
					if u.IsDomain(domain) {
						validEntries = append(validEntries, domain)
					} else {
						invalidEntries = append(invalidEntries, line)
					}
				} else {
					invalidEntries = append(invalidEntries, line)
				}
			} else {
				invalidEntries = append(invalidEntries, line)
			}
		} else {
			// Non-matching format for allowlist
			invalidEntries = append(invalidEntries, line)
		}
	}
	return validEntries, invalidEntries
}

func ExtractBlocklistDomains(logger *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if u.IsComment(line) {
			continue
		}

		// extract domain between || and ^
		if domain, ok := strings.CutPrefix(line, AdguardBlockPrefix); ok {
			if domain, ok := strings.CutSuffix(domain, AdguardBlockSuffix); ok {
				if u.IsDomain(domain) {
					validEntries = append(validEntries, domain)
				} else {
					invalidEntries = append(invalidEntries, line)
				}
			} else {
				invalidEntries = append(invalidEntries, line)
			}
		} else {
			// Non-matching format for blocklist
			invalidEntries = append(invalidEntries, line)
		}
	}
	return validEntries, invalidEntries
}

// AdGuardHttpUrlProcessor converts http/https URLs into AdGuard blocklist rules (||host/path$)
type AdGuardHttpUrlProcessor struct {
	BaseProcessor
}

// NewAdGuardHttpUrlProcessor creates a new processor for adguard_httpurl
func NewAdGuardHttpUrlProcessor(sourceType, listType string) *AdGuardHttpUrlProcessor {
	return &AdGuardHttpUrlProcessor{BaseProcessor: NewBaseProcessor(sourceType, listType)}
}

// Process converts each valid http/https URL into an AdGuard block rule
func (p *AdGuardHttpUrlProcessor) Process(_ *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")
	unique := make(map[string]struct{})

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || u.IsComment(line) {
			continue
		}

		if strings.HasPrefix(line, AdguardBlockPrefix) || strings.HasPrefix(line, AdguardExceptionPrefix) {
			invalidEntries = append(invalidEntries, line)
			continue
		}

		if entry, ok := ParseToAdguardEntry(line); ok {
			unique[entry] = struct{}{}
		} else {
			invalidEntries = append(invalidEntries, line)
		}
	}

	for e := range unique {
		validEntries = append(validEntries, e)
	}
	slices.Sort(validEntries)

	return validEntries, invalidEntries
}

// AdGuardDomainBlocklistProcessor converts plain domains into AdGuard blocklist rules (||domain^)
type AdGuardDomainBlocklistProcessor struct {
	BaseProcessor
}

// NewAdGuardDomainBlocklistProcessor creates a new processor for adguard_domain blocklist
func NewAdGuardDomainBlocklistProcessor(sourceType, listType string) *AdGuardDomainBlocklistProcessor {
	return &AdGuardDomainBlocklistProcessor{BaseProcessor: NewBaseProcessor(sourceType, listType)}
}

// Process converts each valid domain into an AdGuard block rule
func (p *AdGuardDomainBlocklistProcessor) Process(_ *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || u.IsComment(line) {
			continue
		}

		if strings.HasPrefix(line, AdguardBlockPrefix) || strings.HasPrefix(line, AdguardExceptionPrefix) {
			invalidEntries = append(invalidEntries, line)
			continue
		}
		if u.IsDomain(line) {
			validEntries = append(validEntries, AdguardBlockPrefix+line+AdguardBlockSuffix)
		} else {
			invalidEntries = append(invalidEntries, line)
		}
	}
	return validEntries, invalidEntries
}

// AdGuardDomainAllowlistProcessor converts plain domains into AdGuard allowlist rules (@@||domain^)
type AdGuardDomainAllowlistProcessor struct {
	BaseProcessor
}

// NewAdGuardDomainAllowlistProcessor creates a new processor for adguard_domain allowlist
func NewAdGuardDomainAllowlistProcessor(sourceType, listType string) *AdGuardDomainAllowlistProcessor {
	return &AdGuardDomainAllowlistProcessor{BaseProcessor: NewBaseProcessor(sourceType, listType)}
}

// Process converts each valid domain into an AdGuard allow rule
func (p *AdGuardDomainAllowlistProcessor) Process(_ *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || u.IsComment(line) {
			continue
		}

		if strings.HasPrefix(line, AdguardBlockPrefix) || strings.HasPrefix(line, AdguardExceptionPrefix) {
			invalidEntries = append(invalidEntries, line)
			continue
		}
		if u.IsDomain(line) {
			validEntries = append(validEntries, AdguardExceptionPrefix+AdguardBlockPrefix+line+AdguardBlockSuffix)
		} else {
			invalidEntries = append(invalidEntries, line)
		}
	}
	return validEntries, invalidEntries
}

func init() {
	SpecialProcessors.Register(sourceTypeAdguard, constants.ListTypeBlocklist,
		func(st string, lt string) Processor {
			return NewAdGuardBlocklistProcessor(st, lt)
		},
	)

	SpecialProcessors.Register(sourceTypeAdguard, constants.ListTypeAllowlist,
		func(st string, lt string) Processor {
			return NewAdGuardAllowlistProcessor(st, lt)
		},
	)

	RegisterGenericProcessor(adguardSourceTypeDomain, ExtractAllowlistDomains, false, true)
	RegisterGenericProcessor(adguardSourceTypeDomain, ExtractBlocklistDomains, true, false)

	SpecialProcessors.Register(sourceTypeAdguardDomain, constants.ListTypeBlocklist, func(st, lt string) Processor {
		return NewAdGuardDomainBlocklistProcessor(st, lt)
	})
	SpecialProcessors.Register(sourceTypeAdguardDomain, constants.ListTypeAllowlist, func(st, lt string) Processor {
		return NewAdGuardDomainAllowlistProcessor(st, lt)
	})

	SpecialProcessors.Register(sourceTypeAdguardHttpUrl, constants.ListTypeBlocklist, func(st, lt string) Processor {
		return NewAdGuardHttpUrlProcessor(st, lt)
	})
}
