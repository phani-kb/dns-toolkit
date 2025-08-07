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

const (
	defaultDownloaderName string = "default"
	retryDelay                   = time.Second * constants.DefaultRetryDelayInSeconds
)

// HTTPStatusError Custom error types for better handling
type HTTPStatusError struct {
	Status     string
	URL        string
	StatusCode int
}

func (e *HTTPStatusError) Error() string {
	return fmt.Sprintf("HTTP request returned non-success status: %d %s for %s", e.StatusCode, e.Status, e.URL)
}

type CertVerificationError struct { // nolint: govet
	Host string
	Err  error
}

func (e *CertVerificationError) Error() string {
	return fmt.Sprintf("Certificate verification failed for host %s: %v", e.Host, e.Err)
}

type DefaultDownloader struct {
	rnd           *rand.Rand
	maxRetries    int
	retryDelay    time.Duration
	clientTimeout time.Duration
}

// NewDefaultDownloaderWithRetries creates a new DefaultDownloader with custom retry
func NewDefaultDownloaderWithRetries(maxRetries int) *DefaultDownloader {
	return &DefaultDownloader{
		rnd:           rand.New(rand.NewSource(time.Now().UnixNano())),
		maxRetries:    maxRetries,
		retryDelay:    retryDelay,
		clientTimeout: time.Second * constants.DefaultClientTimeoutInSeconds,
	}
}

// NewDefaultDownloaderForTesting creates a new DefaultDownloader with retry
func NewDefaultDownloaderForTesting(maxRetries int, testRetryDelay time.Duration) *DefaultDownloader {
	testTimeout := 500 * time.Millisecond
	return &DefaultDownloader{
		rnd:           rand.New(rand.NewSource(time.Now().UnixNano())),
		maxRetries:    maxRetries,
		retryDelay:    testRetryDelay,
		clientTimeout: testTimeout,
	}
}

// NewDefaultDownloaderWithOptions creates a new DefaultDownloader with fully customizable options
func NewDefaultDownloaderWithOptions(maxRetries int, retryDelay, clientTimeout time.Duration) *DefaultDownloader {
	return &DefaultDownloader{
		rnd:           rand.New(rand.NewSource(time.Now().UnixNano())),
		maxRetries:    maxRetries,
		retryDelay:    retryDelay,
		clientTimeout: clientTimeout,
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
	var localModTime time.Time
	var localFileSize int64
	if info, err := os.Stat(filePath); err == nil {
		fileExists = true
		localModTime = info.ModTime()
		localFileSize = info.Size()
	}
	parsedURL, err := url.Parse(strings.TrimSpace(file.URL))
	if err != nil {
		logger.Errorf("Parsing URL error: %v", err)
		return "", false, err
	}
	client := d.createHTTPClient(logger, skipCertVerify, skipCertHosts, parsedURL)
	userAgent := cfg.GetUserAgent(logger, applicationConfig)
	logger.Debugf("User-Agent: %s", userAgent)
	if fileExists && d.canSkipDownload(logger, client, userAgent, file, localFileSize, localModTime) {
		archiveErr := d.handleArchiveFile(logger, file, filePath)
		return filePath, true, archiveErr
	}

	var resp *http.Response
	var lastErr error

	for attempt := 1; attempt <= d.maxRetries; attempt++ {
		req, reqErr := http.NewRequest("GET", file.URL, nil)
		if reqErr != nil {
			logger.Errorf("Creating request error: %v", reqErr)
			lastErr = reqErr
			continue
		}

		req.Header.Set("User-Agent", userAgent)
		resp, err = client.Do(req)

		// If we get a response but encounter an error later, we should still close the body
		if err == nil && resp != nil && resp.Body != nil {
			defer func() {
				if closeErr := resp.Body.Close(); closeErr != nil {
					logger.Warnf("Failed to close response body: %v", closeErr)
				}
			}()
		}

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

			time.Sleep(d.retryDelay * time.Duration(1<<uint(attempt-1)))
			continue
		}

		// Check response status
		if resp.StatusCode == http.StatusTooManyRequests && attempt < d.maxRetries {
			u.CloseBody(logger, resp.Body)

			// For 429, use exponential backoff with jitter
			waitTime := d.retryDelay * time.Duration(1<<uint(attempt))
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

		// nolint:lll
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
	err = d.handleArchiveFile(logger, file, filePath)
	return filePath, false, err
}

func (d *DefaultDownloader) createHTTPClient(
	logger *multilog.Logger,
	skipCertVerify bool,
	skipCertHosts []string,
	parsedURL *url.URL,
) *http.Client {
	client := &http.Client{
		Timeout: d.clientTimeout,
	}

	if skipCertVerify {
		for _, host := range skipCertHosts {
			if strings.Contains(parsedURL.Host, host) {
				logger.Debugf("Skipping certificate verification for host: %s", parsedURL.Host)
				transport := &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				}
				client.Transport = transport
				break
			}
		}
	}
	return client
}

func (d *DefaultDownloader) canSkipDownload(
	logger *multilog.Logger,
	client *http.Client,
	userAgent string,
	file c.DownloadFile,
	localFileSize int64,
	localModTime time.Time,
) bool {
	summaryFile := filepath.Join(constants.SummaryDir, constants.DefaultSummaryFiles["download"])
	if !d.ShouldDownload(logger, summaryFile, file) {
		return true
	}

	headReq, err := http.NewRequest("HEAD", file.URL, nil)
	if err != nil {
		return false
	}
	headReq.Header.Set("User-Agent", userAgent)
	resp, err := client.Do(headReq)

	// Always ensure we close the response body if we got a response
	if resp != nil && resp.Body != nil {
		defer func() {
			if closeErr := resp.Body.Close(); closeErr != nil {
				logger.Warnf("Failed to close response body: %v", closeErr)
			}
		}()
	}

	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}

	filePath := filepath.Join(file.Folder, file.Filename)
	if contentLength := resp.ContentLength; contentLength != -1 && contentLength == localFileSize {
		if lastModStr := resp.Header.Get("Last-Modified"); lastModStr != "" {
			lastMod, err := time.Parse(http.TimeFormat, lastModStr)
			if err == nil && !lastMod.After(localModTime) {
				logger.Infof("Skipping download, local file is up-to-date: %s", filePath)
				return true
			}
		} else {
			logger.Infof("Skipping download, local file exists with same size: %s", filePath)
			return true
		}
	}
	return false
}

func (d *DefaultDownloader) handleArchiveFile(logger *multilog.Logger, file c.DownloadFile, filePath string) error {
	if file.IsArchive {
		if err := u.ExtractArchive(logger, filePath, file.Folder); err != nil {
			logger.Errorf("Failed to extract archive: %v", err)
			return err
		}
		logger.Infof("Extracted archive %s", filePath)
		for _, target := range file.Targets {
			if err := u.CopySourceToTarget(logger, target); err != nil {
				logger.Errorf("Failed to copy target file: %v", err)
				return err
			}
		}
		logger.Infof("Copied target %d file(s)", len(file.Targets))
		return nil
	}
	return nil
}

func (d *DefaultDownloader) ShouldDownload(
	logger *multilog.Logger,
	summaryFile string,
	file c.DownloadFile,
) bool {
	filePath := filepath.Join(file.Folder, file.Filename)
	if _, err := os.Stat(filePath); err == nil {
		logger.Debugf("File %s already exists", filePath)
		return u.ShouldDownloadSource(logger, summaryFile, file.Name)
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
