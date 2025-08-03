package processors

import (
	"bytes"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

const sourceTypeDomainCustomHtmlPuppyScams = "domain_custom_html_puppyscams"

// DomainCustomHtmlPuppyScamsProcessor implements a processor for extracting domain names
// from HTML content, specifically focusing on PuppyScams.
type DomainCustomHtmlPuppyScamsProcessor struct {
	BaseProcessor
}

// NewDomainCustomHtmlPuppyScamsProcessor creates a new instance of DomainCustomHtmlPuppyScamsProcessor.
//
// Returns:
//   - A configured DomainCustomHtmlPuppyScamsProcessor instance
func NewDomainCustomHtmlPuppyScamsProcessor(sourceType, listType string) *DomainCustomHtmlPuppyScamsProcessor {
	return &DomainCustomHtmlPuppyScamsProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

// Process extracts valid domain names from HTML tags.
//
// Parameters:
//   - logger: Logger for recording operations and errors
//   - content: The HTML content to process
//
// Returns:
//   - A slice of valid domain names found in the content, sorted alphabetically
//   - A slice of invalid entries that couldn't be parsed as domains
func (p *DomainCustomHtmlPuppyScamsProcessor) Process(logger *multilog.Logger, content string) ([]string, []string) {
	var validEntries, invalidEntries []string
	uniqueDomains := make(map[string]struct{})

	// Parse HTML using goquery
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(content)))
	if err != nil {
		logger.Errorf("Error parsing HTML content: %v", err)
		return validEntries, invalidEntries
	}

	// Helper function to process domain text
	processDomain := func(text string) {
		domain := strings.ToLower(strings.TrimSpace(text))
		if domain == "" {
			return
		}
		if utils.IsDomain(domain) {
			if _, exists := uniqueDomains[domain]; !exists {
				uniqueDomains[domain] = struct{}{}
			}
		} else {
			invalidEntries = append(invalidEntries, domain)
		}
	}

	// Iterate over all <a> tags and process text
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		processDomain(s.Text())
	})

	// Convert map keys to sorted slice
	for domain := range uniqueDomains {
		validEntries = append(validEntries, domain)
	}
	sort.Strings(validEntries)

	return validEntries, invalidEntries
}

func init() {
	RegisterProcessor(sourceTypeDomainCustomHtmlPuppyScams, func(st string, lt string) Processor {
		return NewDomainCustomHtmlPuppyScamsProcessor(st, lt)
	})
}
