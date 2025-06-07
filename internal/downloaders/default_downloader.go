package downloaders

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	cfg "github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

const defaultDownloaderName string = "default"
const retryDelay = time.Second * constants.DefaultRetryDelayInSeconds

// HTTPStatusError Custom error types for better handling
type HTTPStatusError struct {
	StatusCode int
	Status     string
	URL        string
}

func (e *HTTPStatusError) Error() string {
	return fmt.Sprintf("HTTP request returned non-success status: %d %s for %s", e.StatusCode, e.Status, e.URL)
}

type CertVerificationError struct {
	Host string
	Err  error
}

func (e *CertVerificationError) Error() string {
	return fmt.Sprintf("Certificate verification failed for host %s: %v", e.Host, e.Err)
}

type DefaultDownloader struct {
	rnd        *rand.Rand
	maxRetries int
}

// NewDefaultDownloaderWithRetries creates a new DefaultDownloader with custom retry settings
func NewDefaultDownloaderWithRetries(maxRetries int) *DefaultDownloader {
	return &DefaultDownloader{
		rnd:        rand.New(rand.NewSource(time.Now().UnixNano())),
		maxRetries: maxRetries,
	}
}

// DefaultDownloaderName returns the name of the default downloader
func DefaultDownloaderName() string {
	return defaultDownloaderName
}

func (d *DefaultDownloader) Name() string {
	return defaultDownloaderName
}

func (d *DefaultDownloader) Download(
	logger *multilog.Logger,
	file c.DownloadFile,
	skipCertVerify bool,
	skipCertHosts []string,
	applicationConfig cfg.ApplicationConfig,
) (string, bool, error) {
	fileUrl := file.URL
	if strings.HasPrefix(fileUrl, "file://") {
		return d.copyLocalFile(logger, fileUrl, file.Folder, file.Filename)
	}
	return d.downloadFile(logger, file, skipCertVerify, skipCertHosts, applicationConfig)
}

func (d *DefaultDownloader) PostDownloadProcess(_ *multilog.Logger, _ string, _ int) error {
	return nil
}

func (d *DefaultDownloader) copyLocalFile(
	logger *multilog.Logger,
	url, destFolder, fileName string,
) (string, bool, error) {
	localPath := strings.TrimPrefix(url, "file://")
	destPath := filepath.Join(destFolder, fileName)
	fileExists := false
	if _, err := os.Stat(destPath); err == nil {
		fileExists = true
	}
	input, err := os.Open(localPath)
	if err != nil {
		logger.Errorf("Opening local file error: %v", err)
		return "", false, err
	}
	defer u.CloseFile(logger, input)
	output, err := os.Create(destPath)
	if err != nil {
		logger.Errorf("Creating destination file error: %v", err)
		return "", false, err
	}
	defer u.CloseFile(logger, output)
	if _, err = io.Copy(output, input); err != nil {
		logger.Errorf("Copying file error: %v", err)
		return "", false, err
	}
	logger.Infof("Copied local file to %s", destPath)
	return destPath, fileExists, nil
}

func (d *DefaultDownloader) downloadFile(
	logger *multilog.Logger,
	file c.DownloadFile,
	skipCertVerify bool,
	skipCertHosts []string,
	applicationConfig cfg.ApplicationConfig,
) (string, bool, error) {
	filePath := filepath.Join(file.Folder, file.Filename)
	fileExists := false
	if _, err := os.Stat(filePath); err == nil {
		fileExists = true
	}
	parsedURL, err := url.Parse(file.URL)
	if err != nil {
		logger.Errorf("Parsing URL error: %v", err)
		return "", false, err
	}
	client := d.createHTTPClient(logger, skipCertVerify, skipCertHosts, parsedURL)
	userAgent := u.GetUserAgent(logger, applicationConfig)
	logger.Debugf("User-Agent: %s", userAgent)
	if fileExists {
		return filePath, true, nil
	}

	var resp *http.Response
	var lastErr error

	for attempt := 1; attempt <= d.maxRetries; attempt++ {
		req, err := http.NewRequest("GET", file.URL, nil)
		if err != nil {
			logger.Errorf("Creating request error: %v", err)
			lastErr = err
			continue
		}

		req.Header.Set("User-Agent", userAgent)
		resp, err = client.Do(req)

		if err != nil {
			var urlErr *url.Error
			isTLSErr := errors.As(err, &urlErr)
			if isTLSErr && strings.Contains(urlErr.Error(), "certificate") {
				certErr := &CertVerificationError{
					Host: parsedURL.Host,
					Err:  err,
				}
				logger.Warnf("Certificate validation error: %v", certErr)
				lastErr = certErr
			} else {
				lastErr = err
				logger.Warnf("Attempt %d: Failed to download file: %v", attempt, err)
			}

			if isTLSErr && strings.Contains(urlErr.Error(), "certificate") && attempt == 1 {
				logger.Warnf(
					"Certificate error detected, consider adding %s to skipCertVerificationHosts in config",
					parsedURL.Host,
				)
				if !skipCertVerify || !containsHost(skipCertHosts, parsedURL.Host) {
					return "", false, lastErr
				}
			}

			time.Sleep(retryDelay * time.Duration(1<<uint(attempt-1)))
			continue
		}

		// Check response status
		if resp.StatusCode == http.StatusTooManyRequests && attempt < d.maxRetries {
			u.CloseBody(logger, resp.Body)

			// For 429, use exponential backoff with jitter
			waitTime := retryDelay * time.Duration(1<<uint(attempt))
			// Add jitter to avoid synchronized retries
			jitter := time.Duration(d.rnd.Intn(1000)) * time.Millisecond
			backoffTime := waitTime + jitter

			logger.Warnf("Attempt %d: Rate limited (429) for %s. Retrying in %.1f seconds...",
				attempt, file.URL, backoffTime.Seconds())

			time.Sleep(backoffTime)
			continue
		}

		break
	}

	if resp == nil {
		return "", false, lastErr
	}

	defer u.CloseBody(logger, resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		statusErr := &HTTPStatusError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			URL:        file.URL,
		}

		switch resp.StatusCode {
		case http.StatusForbidden:
			logger.Errorf(
				"Access forbidden (403) - the server may require authentication or the source may be blocking automated downloads: %s",
				file.URL,
			)
		case http.StatusNotFound:
			logger.Errorf(
				"Resource not found (404) - the URL may be incorrect or the resource may have moved: %s",
				file.URL,
			)
		case http.StatusTooManyRequests:
			logger.Errorf(
				"Rate limited (429) - too many requests to this source, consider increasing the download interval: %s",
				file.URL,
			)
		case http.StatusServiceUnavailable:
			logger.Errorf(
				"Service unavailable (503) - the server is temporarily unavailable, try again later: %s",
				file.URL,
			)
		case http.StatusUnauthorized:
			logger.Errorf("Unauthorized access (401) - authentication required: %s", file.URL)
		default:
			logger.Errorf("HTTP error %d %s for %s", resp.StatusCode, resp.Status, file.URL)
		}

		return "", false, statusErr
	}

	filePath, err = u.SaveFile(logger, file.Folder, file.Filename, resp.Body)
	if err != nil {
		return "", false, err
	}
	logger.Infof("Downloaded %s", filePath)

	return filePath, false, nil
}

func (d *DefaultDownloader) createHTTPClient(
	logger *multilog.Logger,
	skipCertVerify bool,
	skipCertHosts []string,
	parsedURL *url.URL,
) *http.Client {
	client := &http.Client{}
	if skipCertVerify {
		for _, host := range skipCertHosts {
			if strings.Contains(parsedURL.Host, host) {
				logger.Debugf("Skipping certificate verification for host: %s", parsedURL.Host)
				client.Transport = &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				}
				break
			}
		}
	}
	return client
}

func (d *DefaultDownloader) ShouldDownload(
	logger *multilog.Logger,
	summaryFile string,
	file c.DownloadFile,
) bool {
	filePath := filepath.Join(file.Folder, file.Filename)
	if _, err := os.Stat(filePath); err == nil {
		logger.Debugf("File %s already exists", filePath)
		return false
	}
	return true
}

// Helper function to check if a host is in the skipCertHosts list
func containsHost(hosts []string, targetHost string) bool {
	for _, host := range hosts {
		if strings.Contains(targetHost, host) {
			return true
		}
	}
	return false
}

func init() {}
