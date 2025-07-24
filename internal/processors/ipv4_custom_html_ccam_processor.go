package processors

import (
	"bytes"
	"regexp"
	"slices"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

// sourceTypeIpv4CustomHtmlCcam defines the type identifier for the IPv4 Custom HTML CCAM processor
const sourceTypeIpv4CustomHtmlCcam = "ipv4_custom_html_ccam"

// Ipv4CustomHtmlCcamProcessor implements a processor for extracting IPv4 addresses
// from HTML content, specifically targeting CCAM-formatted HTML tables.
type Ipv4CustomHtmlCcamProcessor struct {
	BaseProcessor
}

// NewIpv4CustomHtmlCcamProcessor creates a new instance of Ipv4CustomHtmlCcamProcessor.
//
// Returns:
//   - A configured Ipv4CustomHtmlCcamProcessor instance
func NewIpv4CustomHtmlCcamProcessor(sourceType, listType string) *Ipv4CustomHtmlCcamProcessor {
	return &Ipv4CustomHtmlCcamProcessor{
		BaseProcessor: NewBaseProcessor(sourceType, listType),
	}
}

// Process extracts valid IPv4 addresses from HTML table cells.
// It parses the HTML content, finds all table cells, and extracts IPv4 addresses from each cell.
// Duplicate addresses are automatically removed.
//
// Parameters:
//   - logger: Logger for recording operations and errors
//   - content: The HTML content to process
//
// Returns:
//   - A slice of valid IPv4 addresses found in the content, sorted alphabetically
//   - A slice of invalid entries that couldn't be parsed as IPv4 addresses
func (p *Ipv4CustomHtmlCcamProcessor) Process(logger *multilog.Logger, content string) ([]string, []string) {
	return extractUniqueFromHtmlTableCells(
		logger,
		content,
		constants.SourceTypeRegexMap[constants.SourceTypeIpv4],
		utils.IsIP,
	)
}

// extractUniqueFromHtmlTableCells is a shared helper.
func extractUniqueFromHtmlTableCells(
	logger *multilog.Logger,
	content string,
	regex *regexp.Regexp,
	isValid func(string) bool,
) (validEntries, invalidEntries []string) {
	reader := bytes.NewReader([]byte(content))
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		logger.Errorf("Error parsing HTML content: %v", err)
		return nil, nil
	}
	unique := make(map[string]struct{})
	doc.Find("table tr td").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if regex.MatchString(text) {
			matches := regex.FindAllString(text, -1)
			for _, val := range matches {
				if isValid(val) {
					unique[val] = struct{}{}
				}
			}
		} else if text != "" {
			invalidEntries = append(invalidEntries, text)
		}
	})
	validEntries = make([]string, 0, len(unique))
	for val := range unique {
		validEntries = append(validEntries, val)
	}
	slices.Sort(validEntries)
	return validEntries, invalidEntries
}

func init() {
	RegisterProcessor(sourceTypeIpv4CustomHtmlCcam, func(st string, lt string) Processor {
		return NewIpv4CustomHtmlCcamProcessor(st, lt)
	})
}
