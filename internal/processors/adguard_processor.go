package processors

import (
	"strings"

	"github.com/phani-kb/multilog"

	u "github.com/phani-kb/dns-toolkit/internal/utils"
)

const (
	sourceTypeAdguard       = "adguard"
	adguardSourceTypeDomain = "domain_adguard"
	adguardExceptionPrefix  = "@@"
	adguardBlockPrefix      = "||"
	adguardBlockSuffix      = "^"
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

		if !strings.HasPrefix(line, adguardExceptionPrefix) {
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

		if strings.HasPrefix(line, adguardExceptionPrefix) {
			validEntries = append(validEntries, line)
		} else {
			// Non-matching format for allowlist
			invalidEntries = append(invalidEntries, line)
		}
	}
	return validEntries, invalidEntries
}

func init() {
}
