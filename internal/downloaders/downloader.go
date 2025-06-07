package downloaders

import (
	c "github.com/phani-kb/dns-toolkit/internal/common"
	cfg "github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/multilog"
)

type Downloader interface {
	// Download is a method to implement the logic to download the file
	Download(
		logger *multilog.Logger,
		file c.DownloadFile,
		skipCertVerify bool,
		skipCertHosts []string,
		applicationConfig cfg.ApplicationConfig,
	) (string, bool, error)

	// PostDownloadProcess is a method to implement the logic to consider the given count of entries after the download
	PostDownloadProcess(logger *multilog.Logger, filePath string, count int) error

	// Name is a method to return the name of the downloader
	Name() string

	ShouldDownload(logger *multilog.Logger, summaryFile string, file c.DownloadFile) bool
}
