package downloaders

import (
	"net/http"
	"time"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/multilog"
)

// HandleArchiveForTest is a test-friendly wrapper for handleArchiveFile
func (d *DefaultDownloader) HandleArchiveForTest(logger *multilog.Logger, file c.DownloadFile, filePath string) error {
	return d.handleArchiveFile(logger, file, filePath)
}

// CanSkipDownloadForTest is a test-friendly wrapper for canSkipDownload
func (d *DefaultDownloader) CanSkipDownloadForTest(
	logger *multilog.Logger,
	client *http.Client,
	userAgent string,
	file c.DownloadFile,
	localFileSize int64,
	localModTime time.Time,
) bool {
	return d.canSkipDownload(logger, client, userAgent, file, localFileSize, localModTime)
}
