package cmd

import (
	"fmt"
	"path/filepath"
	"strings"
	"sync"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	r "github.com/phani-kb/dns-toolkit/internal/processors"

	"github.com/phani-kb/dns-toolkit/internal/config"
	u "github.com/phani-kb/dns-toolkit/internal/utils"

	"github.com/phani-kb/multilog"

	"github.com/spf13/cobra"
)

// initializeSourceTypeCounts creates and returns a new map with all source types initialized to 0
func initializeSourceTypeCounts() map[string]int {
	return map[string]int{
		constants.SourceTypeIpv4:                     0,
		constants.SourceTypeIpv4RangeExpand:          0,
		constants.SourceTypeIpv6:                     0,
		constants.SourceTypeCidrIpv4:                 0,
		constants.SourceTypeIpv4CidrExpand:           0,
		constants.SourceTypeDomain:                   0,
		constants.SourceTypeDomainFinder:             0,
		constants.SourceTypeAdguard:                  0,
		constants.SourceTypeIpv4Hostname:             0,
		constants.SourceTypeHostname:                 0,
		constants.SourceTypeUnknown:                  0,
		constants.SourceTypeDomainAdguard:            0,
		constants.SourceTypeDomainCsvHttpUrlFind:     0,
		constants.SourceTypeDomainCustomCsvBlackbook: 0,
		constants.SourceTypeDomainCustomCsvMaltrail:  0,
		constants.SourceTypeDomainCustomHtmlCcam:     0,
		constants.SourceTypeDomainHttpUrl:            0,
		constants.SourceTypeDomainUrl:                0,
		constants.SourceTypeDomainWithCommentSuffix:  0,
		constants.SourceTypeIpv4CustomHtmlCcam:       0,
		constants.SourceTypeIpv4Find:                 0,
		constants.SourceTypeIpv4HttpUrl:              0,
		constants.SourceTypeIpv4Url:                  0,
		constants.SourceTypeIpv6Find:                 0,
		constants.SourceTypeIpv6Htaccess:             0,
	}
}

var sourceTypesSummaryCmd = &cobra.Command{
	Use:   "sts",
	Short: "Prints the source types summary",
	Run: func(cmd *cobra.Command, args []string) {
		printSourceTypeSummary(Logger, SourcesConfigs)
	},
}

func printSourceTypeSummary(logger *multilog.Logger, sourcesConfigs []config.SourcesConfig) {
	counts := initializeSourceTypeCounts()

	// Initialize a slice to keep track of mismatches
	var mismatches []string
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, sourcesConfig := range sourcesConfigs {
		for _, source := range sourcesConfig.GetEnabledSources(AppConfig.DNSToolkit.SourceFilters) {
			wg.Add(1)
			go func(source config.Source) {
				defer wg.Done()
				processSource(logger, source, counts, &mismatches, &mu)
			}(source)
		}
	}

	wg.Wait()

	// Print the summary
	logger.Infof("Processed %d source configs. SUMMARY:", len(sourcesConfigs))
	for sourceType, count := range counts {
		logger.Infof("%s: %d", sourceType, count)
	}

	if len(mismatches) > 0 {
		logger.Warnf("MISMATCHES:")
		for _, mismatch := range mismatches {
			logger.Warnf(mismatch)
		}
	}

	total := 0
	for _, count := range counts {
		total += count
	}
	logger.Infof("Total: %d", total)
}

func processSource(
	logger *multilog.Logger,
	source config.Source,
	counts map[string]int,
	mismatches *[]string,
	mu *sync.Mutex,
) {
	downloadFile, err := source.GetDownloadFile(logger, constants.DownloadDir)
	if err != nil {
		logger.Errorf("Getting download file error: %v", err)
		return
	}
	for _, target := range downloadFile.Targets {
		downloadedFile := filepath.Join(target.TargetFolder, target.TargetFile)
		logger.Debugf("Processing %s", downloadedFile)
		lines, err := u.PickRandomLines(downloadedFile, constants.MaxSampleLinesToCategorize)
		if err != nil {
			logger.Errorf("Error reading file %s: %v", downloadedFile, err)
			return
		}

		for _, sourceTypeObj := range source.Types {
			sourceType := sourceTypeObj.Name
			actualType := constants.SourceTypeUnknown

			if u.StringInSlice(sourceType, u.GetMapKeys(constants.SourceTypeRegexMap)) {
				if source.TypeCount == 1 {
					actualType = categorizeFileContent(logger, lines)
					logger.Infof("%s -> %s (expected: %s)", source.Name, actualType, sourceType)
					mu.Lock()
					counts[actualType]++
					mu.Unlock()
				} else {
					actualType = constants.SourceTypeMixed
					logger.Infof("%s -> %s", source.Name, sourceType)
					mu.Lock()
					counts[sourceType]++
					mu.Unlock()
				}
			} else {
				for _, listTypeObj := range sourceTypeObj.GetListTypes() {
					listType := listTypeObj.Name
					_, exists := r.Processors.GetProcessor(sourceType, listType)
					if exists {
						actualType = sourceType
						logger.Infof("%s -> %s", source.Name, sourceType)
						mu.Lock()
						counts[sourceType]++
						mu.Unlock()
					} else {
						logger.Warnf("%s: unknown source type %s", source.Name, sourceType)
						mu.Lock()
						counts[constants.SourceTypeUnknown]++
						mu.Unlock()
					}
				}
			}

			// Check for mismatches
			if actualType != sourceType && actualType != constants.SourceTypeMixed ||
				actualType == constants.SourceTypeUnknown {
				mu.Lock()
				*mismatches = append(
					*mismatches,
					fmt.Sprintf("%s: expected %s, got %s", source.Name, sourceType, actualType),
				)
				mu.Unlock()
			}
		}
	}
}

func categorizeFileContent(logger *multilog.Logger, lines []string) string {
	regexCounts := initializeSourceTypeCounts()

	for _, line := range lines {
		// NOTE: change in the order of regex checks may affect the results
		if constants.SourceTypeRegexMap[constants.SourceTypeIpv4Hostname].MatchString(line) {
			regexCounts[constants.SourceTypeIpv4Hostname]++
		} else if constants.SourceTypeRegexMap[constants.SourceTypeIpv6].MatchString(line) {
			regexCounts[constants.SourceTypeIpv6]++
		} else if constants.SourceTypeRegexMap[constants.SourceTypeCidrIpv4].MatchString(line) {
			regexCounts[constants.SourceTypeCidrIpv4]++
		} else if constants.SourceTypeRegexMap[constants.SourceTypeDomain].MatchString(line) {
			regexCounts[constants.SourceTypeDomain]++
		} else if constants.SourceTypeRegexMap[constants.SourceTypeIpv4].MatchString(line) {
			regexCounts[constants.SourceTypeIpv4]++
		}
	}

	// print each regex count
	sb := strings.Builder{}
	for sourceType, count := range regexCounts {
		if count > 0 {
			sb.WriteString(fmt.Sprintf("[%s: %d] ", sourceType, count))
		}
	}
	logger.Debugf("Regex counts:\n%s", sb.String())

	maxCount := 0
	detectedType := constants.SourceTypeUnknown
	for sourceType, count := range regexCounts {
		if count > maxCount {
			maxCount = count
			detectedType = sourceType
		}
	}

	if maxCount == 0 {
		detectedType = constants.SourceTypeUnknown
	}

	return detectedType
}
