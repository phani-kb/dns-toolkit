package utils

import (
	"encoding/json"
	"os"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/multilog"
)

func GetLastSummary[T any](logger *multilog.Logger, summaryFile string, sourceName string) (T, error) {
	var zeroValue T
	logger.Debugf("Getting last summary from file: %s", summaryFile)
	if _, err := os.Stat(summaryFile); err != nil {
		if os.IsNotExist(err) {
			logger.Debugf("Summary file does not exist: %s", summaryFile)
			return zeroValue, nil
		}
		logger.Errorf("Error checking file: %v (file: %s)", err, summaryFile)
		return zeroValue, err
	}
	content, err := os.ReadFile(summaryFile)
	if err != nil {
		logger.Errorf("Reading file error: %v (file: %s)", err, summaryFile)
		return zeroValue, err
	}

	var summaries []T
	if err := json.Unmarshal(content, &summaries); err != nil {
		logger.Errorf("Unmarshalling JSON error: %v", err)
		return zeroValue, err
	}

	if len(summaries) == 0 {
		logger.Debugf("No summaries found")
		return zeroValue, nil
	}

	for _, summary := range summaries {
		switch v := any(summary).(type) {
		case c.DownloadSummary:
			if v.Name == sourceName {
				return summary, nil
			}
		case c.ProcessedSummary:
			if v.Name == sourceName {
				return summary, nil
			}
		default:
			logger.Errorf("Unknown summary type: %T", summary)
		}
	}

	logger.Debugf("No summary found for source: %s", sourceName)
	return zeroValue, nil
}

// GetSummaryFiles is a generic function that retrieves files from a summary file
func GetSummaryFiles[T any](
	logger *multilog.Logger,
	sourceType string,
	summaryFile string,
	extractFiles func(T, string) []string,
) ([]string, error) {
	content, err := os.ReadFile(summaryFile)
	if err != nil {
		logger.Errorf("Error reading %s file: %s", summaryFile, err)
		return nil, err
	}

	var summaries []T
	err = json.Unmarshal(content, &summaries)
	if err != nil {
		logger.Errorf("Parsing summary error: %v", err)
		return nil, err
	}

	files := make([]string, 0)
	for _, summary := range summaries {
		files = append(files, extractFiles(summary, sourceType)...)
	}
	return files, nil
}
