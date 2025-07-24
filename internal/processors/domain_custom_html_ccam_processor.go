package processors

import (
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

// sourceTypeDomainCustomHtmlCcam defines the type identifier for the Domain Custom HTML CCAM processor
const sourceTypeDomainCustomHtmlCcam = "domain_custom_html_ccam"

// DomainCustomHtmlCcamProcessor implements a processor for extracting domain names
// from HTML content, specifically focussing on CCAM-formatted HTML tables.
type DomainCustomHtmlCcamProcessor struct {
	BaseProcessor
}

// NewDomainCustomHtmlCcamProcessor creates a new instance of DomainCustomHtmlCcamProcessor.
//
// Returns:
//   - A configured DomainCustomHtmlCcamProcessor instance
func NewDomainCustomHtmlCcamProcessor(sourceType, listType string) *DomainCustomHtmlCcamProcessor {
	return &DomainCustomHtmlCcamProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

// Process extracts valid domain names from HTML table cells.
// It parses the HTML content, finds all table cells, and extracts domains from each cell.
// Duplicate domains are automatically removed.
//
// Parameters:
//   - logger: Logger for recording operations and errors
//   - content: The HTML content to process
//
// Returns:
//   - A slice of valid domain names found in the content, sorted alphabetically
//   - A slice of invalid entries that couldn't be parsed as domains
func (p *DomainCustomHtmlCcamProcessor) Process(logger *multilog.Logger, content string) ([]string, []string) {
	return extractUniqueFromHtmlTableCells(
		logger,
		content,
		constants.SourceTypeRegexMap[constants.SourceTypeDomain],
		utils.IsDomain,
	)
}

func init() {
	RegisterProcessor(sourceTypeDomainCustomHtmlCcam, func(st string, lt string) Processor {
		return NewDomainCustomHtmlCcamProcessor(st, lt)
	})
}
