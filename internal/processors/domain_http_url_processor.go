package processors

import (
	"regexp"
	"slices"
	"strings"

	"github.com/phani-kb/multilog"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/utils"
)

// sourceTypeDomainHttpUrl defines the type identifier for the Domain HTTP URL processor
const sourceTypeDomainHttpUrl = "domain_http_url"

// DomainHttpUrlProcessor implements a processor for extracting domain names
// from HTTP and HTTPS URLs.
type DomainHttpUrlProcessor struct {
	BaseProcessor
}

// NewDomainHttpUrlProcessor creates a new instance of DomainHttpUrlProcessor.
//
// Returns:
//   - A configured DomainHttpUrlProcessor instance
func NewDomainHttpUrlProcessor(sourceType, listType string) *DomainHttpUrlProcessor {
	return &DomainHttpUrlProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

// Process extracts valid domain names from HTTP/HTTPS URLs.
// It parses each line, extracts the domain from URLs, and validates them.
// Duplicate domains are automatically removed.
//
// Parameters:
//   - logger: Logger for recording operations and errors
//   - content: The content containing URLs to process
//
// Returns:
//   - A slice of valid domain names found in the content
//   - A slice of invalid entries that couldn't be parsed
func (p *DomainHttpUrlProcessor) Process(_ *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	lines := strings.Split(content, "\n")

	urlRegex := regexp.MustCompile(constants.DomainHttpUrlRegex)
	uniqueDomains := make(map[string]struct{})

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || utils.IsComment(line) {
			continue
		}

		matches := urlRegex.FindStringSubmatch(line)
		if len(matches) > 1 {
			domain := matches[1]
			// Remove any trailing dots
			domain = strings.TrimRight(domain, ".")
			if utils.IsDomain(domain) {
				uniqueDomains[domain] = struct{}{}
			} else {
				invalidEntries = append(invalidEntries, line)
			}
		} else {
			invalidEntries = append(invalidEntries, line)
		}
	}

	// Convert unique domains map to slice
	for domain := range uniqueDomains {
		validEntries = append(validEntries, domain)
	}

	// Sort the results for consistency
	slices.Sort(validEntries)
	return validEntries, invalidEntries
}

func init() {
	RegisterProcessor(sourceTypeDomainHttpUrl, func(st string, lt string) Processor {
		return NewDomainHttpUrlProcessor(st, lt)
	})
}
