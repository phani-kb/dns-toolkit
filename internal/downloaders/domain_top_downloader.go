package downloaders

import (
	"bufio"
	"os"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	cfg "github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/multilog"
)

const domainTopDownloaderName = "transco"

type DomainTopDownloader struct {
	DefaultDownloader
}

// NewDomainTopDownloaderWithRetries creates a new DomainTopDownloader with custom retry settings
func NewDomainTopDownloaderWithRetries(maxRetries int) Downloader {
	return &DomainTopDownloader{
		DefaultDownloader: *NewDefaultDownloaderWithRetries(maxRetries),
	}
}

func (d *DomainTopDownloader) Download(
	logger *multilog.Logger,
	file c.DownloadFile,
	skipCertVerify bool,
	skipCertHosts []string,
	applicationConfig cfg.ApplicationConfig,
) (string, bool, error) {
	return d.DefaultDownloader.Download(logger, file, skipCertVerify, skipCertHosts, applicationConfig)
}

func (d *DomainTopDownloader) PostDownloadProcess(logger *multilog.Logger, filePath string, count int) error {
	file, ferr := os.Open(filePath)
	if ferr != nil {
		return ferr
	}

	scanner := bufio.NewScanner(file)
	domains := make([]string, 0, count)
	for scanner.Scan() {
		if len(domains) >= count {
			break
		}
		domains = append(domains, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if err := file.Close(); err != nil {
		logger.Errorf("Error closing file: %v", err)
		return err
	}

	outputFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := outputFile.Close(); cerr != nil {
			logger.Errorf("Error closing output file: %v", cerr)
			if err == nil {
				err = cerr
			}
		}
	}()

	for _, domain := range domains {
		if _, err := outputFile.WriteString(domain + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (d *DomainTopDownloader) Name() string {
	return domainTopDownloaderName
}

func init() {}
