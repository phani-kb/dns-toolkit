package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	cfg "github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	d "github.com/phani-kb/dns-toolkit/internal/downloaders"
	u "github.com/phani-kb/dns-toolkit/internal/utils"

	"github.com/spf13/cobra"
)

const defaultMaxRetries = constants.DefaultMaxRetries

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download enabled sources",
	Run: func(cmd *cobra.Command, args []string) {
		if err := u.EnsureDirectoryExists(Logger, constants.DownloadDir); err != nil {
			Logger.Errorf("Failed to create download directory: %v", err)
			os.Exit(1)
		}
		if err := u.EnsureDirectoryExists(Logger, constants.SummaryDir); err != nil {
			Logger.Errorf("Failed to create summary directory: %v", err)
			os.Exit(1)
		}

		maxRetries := defaultMaxRetries
		if AppConfig != nil && AppConfig.DNSToolkit.MaxRetries > 0 {
			maxRetries = AppConfig.DNSToolkit.MaxRetries
		}

		defaultDownloader := d.NewDefaultDownloaderWithRetries(maxRetries)
		if defaultDownloader == nil {
			Logger.Warnf("Failed to create default downloader with retry settings")
		} else {
			initErr := d.RegisterDownloader(defaultDownloader)
			if initErr != nil {
				Logger.Warnf("Failed to register default downloader with retry settings: %v", initErr)
			}
		}

		domainTopDownloader := d.NewDomainTopDownloaderWithRetries(maxRetries)
		if domainTopDownloader == nil {
			Logger.Warnf("Failed to create domain top downloader with retry settings")
		} else {
			domainTopErr := d.RegisterDownloader(domainTopDownloader)
			if domainTopErr != nil {
				Logger.Warnf("Failed to register domain top downloader: %v", domainTopErr)
			}
		}

		summaries := make([]c.DownloadSummary, 0)
		var mu sync.Mutex

		maxWorkers := runtime.GOMAXPROCS(0)
		if AppConfig != nil {
			if AppConfig.DNSToolkit.MaxWorkers > 0 {
				maxWorkers = AppConfig.DNSToolkit.MaxWorkers
			}
		}
		Logger.Infof("Using worker pool with %d worker(s) for downloads", maxWorkers)
		workerPool := c.NewDTWorkerPool(maxWorkers)

		ticker := time.NewTicker(constants.DownloadInterval)
		defer ticker.Stop()

		// Stats to track a download process
		var totalSources, successCount, failCount int
		var statsMutex sync.Mutex

		for _, sourcesConfig := range SourcesConfigs {
			var sourceFilters cfg.SourceFilters
			if AppConfig != nil {
				sourceFilters = AppConfig.DNSToolkit.SourceFilters
			}
			for _, source := range sourcesConfig.GetEnabledSources(sourceFilters) {
				totalSources++
				source := source // local copy for goroutine
				workerPool.Submit(func() {

					<-ticker.C

					downloadFile, err := source.GetDownloadFile(Logger, constants.DownloadDir)
					if err != nil {
						Logger.Errorf("Getting download file error: %v", err)
						statsMutex.Lock()
						failCount++
						statsMutex.Unlock()

						mu.Lock()
						summary := c.DownloadSummary{
							Name:                 source.Name,
							Error:                err.Error(),
							LastCheckedTimestamp: u.GetTimestamp(),
						}
						summaries = append(summaries, summary)
						mu.Unlock()

						return
					}

					var downloader d.Downloader
					if specificDownloader, exists := d.GetDownloader(source.Name); exists {
						downloader = specificDownloader
						Logger.Debugf("Using registered downloader for %s", source.Name)
					} else {
						downloader, _ = d.GetDownloader(d.DefaultDownloaderName())
						Logger.Debugf("Using default downloader with %d retries for %s", maxRetries, source.Name)
					}

					var skipCertVerification bool
					var skipCertVerificationHosts []string
					var applicationConfig cfg.ApplicationConfig

					if AppConfig != nil {
						skipCertVerification = AppConfig.DNSToolkit.SkipCertVerification
						skipCertVerificationHosts = AppConfig.DNSToolkit.SkipCertVerificationHosts
						applicationConfig = AppConfig.Application
					}

					filePath, fetchSkipped, err := downloader.Download(
						Logger,
						downloadFile,
						skipCertVerification,
						skipCertVerificationHosts,
						applicationConfig,
					)

					for _, target := range downloadFile.Targets {
						targetFilePath := filepath.Join(target.TargetFolder, target.TargetFile)
						summary := c.DownloadSummary{
							Name:            source.Name,
							URL:             downloadFile.URL,
							TypeCount:       source.TypeCount,
							Types:           source.Types,
							Filepath:        targetFilePath,
							Frequency:       source.Frequency,
							CountToConsider: source.CountToConsider,
							Categories:      source.Categories,
						}

						if err != nil {

							switch e := err.(type) {
							case *d.HTTPStatusError:
								Logger.Errorf("Downloading source %s error: HTTP status %d for %s", source.Name, e.StatusCode, e.URL)
								summary.Error = e.Error()
							case *d.CertVerificationError:
								Logger.Errorf("Downloading source %s error: Certificate verification failed for %s", source.Name, e.Host)
								summary.Error = e.Error()
							default:
								Logger.Errorf("Downloading source %s error: %v", source.Name, err)
								summary.Error = err.Error()
							}

							statsMutex.Lock()
							failCount++
							statsMutex.Unlock()
						} else {
							statsMutex.Lock()
							successCount++
							statsMutex.Unlock()

							if fetchSkipped {
								summary.LastCheckedTimestamp = u.GetTimestamp()
								if info, err := os.Stat(filePath); err == nil {
									summary.LastDownloadTimestamp = info.ModTime().Format(constants.TimestampFormat)
								} else {
									Logger.Errorf("Getting file info error: %v", err)
								}
							} else {
								summary.LastDownloadTimestamp = time.Now().Format(constants.TimestampFormat)
							}

							if err := downloader.PostDownloadProcess(Logger, targetFilePath, summary.CountToConsider); err != nil {
								Logger.Errorf("Post download process error for %s: %v", source.Name, err)
								summary.Error = fmt.Sprintf("Post-download processing error: %v", err)
							} else if AppConfig != nil && AppConfig.DNSToolkit.FilesChecksum.Enabled {
								checksum := u.CalculateChecksum(Logger, filePath, AppConfig.DNSToolkit.FilesChecksum.Algorithm)
								summary.Checksum = checksum
							}
						}

						mu.Lock()
						summaries = append(
							summaries,
							summary,
						)
						mu.Unlock()
					}
				})
			}
		}

		workerPool.Wait()

		Logger.Infof("Download complete: %d sources processed, %d successful, %d failed",
			totalSources, successCount, failCount)

		summaryFile := filepath.Join(
			constants.SummaryDir,
			constants.DefaultSummaryFiles["download"],
		)
		_, err := u.SaveSummaries(Logger, summaries, summaryFile, c.DownloadSummaryLessFunc)
		if err != nil {
			Logger.Errorf("Saving summaries error: %v", err)
		}
	},
}
