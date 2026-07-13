package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	cfg "github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/db"
	d "github.com/phani-kb/dns-toolkit/internal/downloaders"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"golang.org/x/time/rate"

	"github.com/spf13/cobra"
)

const defaultMaxRetries = constants.DefaultMaxRetries

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download enabled sources",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		database := openDB(ctx)
		defer database.CloseLogError(Logger)

		forceFlag, getBoolErr := cmd.Flags().GetBool("force")
		if getBoolErr != nil {
			Logger.Warnf("Failed to parse --force flag (defaulting to false): %v", getBoolErr)
			forceFlag = false
		}
		forceEnv := os.Getenv("DNS_TOOLKIT_FORCE_DOWNLOAD") == "true" || os.Getenv("DNS_TOOLKIT_FORCE_DOWNLOAD") == "1"
		forceDownload := forceFlag || forceEnv

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
		domainTopErr := d.RegisterDownloader(domainTopDownloader)
		if domainTopErr != nil {
			Logger.Warnf("Failed to register domain top downloader: %v", domainTopErr)
		}

		sourcesRepo := db.NewSourcesRepo(database)
		downloadsRepo := db.NewDownloadsRepo(database)
		if err := syncSourcesToDB(ctx, Logger, sourcesRepo, SourcesConfigs); err != nil {
			Logger.Errorf("Failed to sync source definitions to database: %v", err)
			os.Exit(1)
		}

		maxWorkers := runtime.GOMAXPROCS(0)
		if AppConfig != nil {
			if AppConfig.DNSToolkit.MaxWorkers > 0 {
				maxWorkers = AppConfig.DNSToolkit.MaxWorkers
			}
		}
		maxWorkers = max(maxWorkers, 1)
		Logger.Infof("Using worker pool with %d worker(s) for downloads", maxWorkers)
		limiter := createDownloadRateLimiter(maxWorkers)
		workerPool := c.NewDTWorkerPoolWithLimiter(maxWorkers, limiter)

		// Stats to track a download process
		var totalSources, successCount, failCount, downloadedCount int
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
					sourceID, sourceIDErr := sourcesRepo.GetSourceIDByName(source.Name)
					if sourceIDErr != nil {
						Logger.Warnf("Failed to get source ID for %s: %v", source.Name, sourceIDErr)
					}

					persistDownloadSummary := func(summary c.DownloadSummary, persistedPath string) {
						if sourceID <= 0 {
							return
						}

						downloadRow := db.DownloadRow{
							SourceID:                    sourceID,
							TypeCount:                   summary.TypeCount,
							CountToConsider:             summary.CountToConsider,
							SkipGeneralConsolidation:    summary.SkipGeneralConsolidation,
							SkipGroupsConsolidation:     summary.SkipGroupsConsolidation,
							SkipCategoriesConsolidation: summary.SkipCategoriesConsolidation,
							URL:                         summary.URL,
							Filepath:                    persistedPath,
							Frequency:                   summary.Frequency,
							Checksum:                    summary.Checksum,
							Error:                       summary.Error,
							LastDownloadTimestamp:       summary.LastDownloadTimestamp,
							LastCheckedTimestamp:        summary.LastCheckedTimestamp,
						}
						if err := downloadsRepo.UpsertDownload(downloadRow); err != nil {
							Logger.Warnf("Failed to upsert download record for %s: %v", source.Name, err)
						}
					}

					if forceDownload {
						Logger.Debugf("Force downloading source: %s", source.Name)
					}

					downloadFile, err := source.GetDownloadFile(Logger, constants.DownloadDir)
					if err != nil {
						Logger.Errorf("Getting download file error: %v", err)
						statsMutex.Lock()
						failCount++
						statsMutex.Unlock()
						summary := c.DownloadSummary{
							Name:                        source.Name,
							URL:                         source.URL,
							Frequency:                   source.Frequency,
							TypeCount:                   source.TypeCount,
							Types:                       source.Types,
							CountToConsider:             source.CountToConsider,
							Categories:                  source.Categories,
							SkipGeneralConsolidation:    source.SkipGeneralConsolidation,
							SkipGroupsConsolidation:     source.SkipGroupsConsolidation,
							SkipCategoriesConsolidation: source.SkipCategoriesConsolidation,
							Error:                       err.Error(),
							LastCheckedTimestamp:        u.GetTimestamp(),
						}
						persistDownloadSummary(summary, "")

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
							Name:                        source.Name,
							URL:                         downloadFile.URL,
							TypeCount:                   source.TypeCount,
							Types:                       source.Types,
							Filepath:                    targetFilePath,
							Frequency:                   source.Frequency,
							CountToConsider:             source.CountToConsider,
							Categories:                  source.Categories,
							SkipGeneralConsolidation:    source.SkipGeneralConsolidation,
							SkipGroupsConsolidation:     source.SkipGroupsConsolidation,
							SkipCategoriesConsolidation: source.SkipCategoriesConsolidation,
						}

						if err != nil {
							summary.LastCheckedTimestamp = u.GetTimestamp()

							switch e := err.(type) { // wrapped errors handling
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
							if !fetchSkipped {
								downloadedCount++
							}
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

							shouldReprocess := !fetchSkipped

							if fetchSkipped {
								if prevSummary, err := loadPreviousDownloadSummary(
									Logger,
									downloadsRepo,
									source.Name,
									targetFilePath,
								); err == nil && prevSummary != nil {
									if prevSummary.CountToConsider != summary.CountToConsider {
										Logger.Infof("Count to consider changed for %s: %d -> %d, re-processing...",
											source.Name, prevSummary.CountToConsider, summary.CountToConsider)
										shouldReprocess = true

										if downloadFile.IsArchive {
											Logger.Debugf("Re-extracting archive and copying target file for %s", source.Name)
											if err = u.ForceCopySourceToTarget(Logger, target); err != nil {
												Logger.Errorf("Failed to force re-copy target file for %s: %v", source.Name, err)
												summary.Error = fmt.Sprintf("Force re-copy target file error: %v", err)
												shouldReprocess = false
											}
										}
									}
								} else if err != nil {
									Logger.Debugf("Could not load previous summary for %s: %v", source.Name, err)
									shouldReprocess = true
								} else {
									shouldReprocess = true
								}
							}

							if shouldReprocess {
								if err := downloader.PostDownloadProcess(Logger, targetFilePath, summary.CountToConsider); err != nil {
									Logger.Errorf("Post download process error for %s: %v", source.Name, err)
									summary.Error = fmt.Sprintf("Post-download processing error: %v", err)
								}
							}

							if AppConfig != nil && AppConfig.DNSToolkit.FilesChecksum.Enabled {
								checksum := u.CalculateChecksum(Logger, filePath, AppConfig.DNSToolkit.FilesChecksum.Algorithm)
								summary.Checksum = checksum
							}
						}

						persistedPath := summary.Filepath
						if downloadFile.IsArchive {
							persistedPath = filePath
						}
						persistDownloadSummary(summary, persistedPath)
					}
				})
			}
		}

		workerPool.Wait()

		Logger.Infof("Download complete: %d sources processed, %d successful (%d downloaded, %d skipped), %d failed",
			totalSources, successCount, downloadedCount, successCount-downloadedCount, failCount)
	},
}

func init() {
	downloadCmd.Flags().Bool("force", false, "Force re-download of all sources (ignores existing summaries)")
}

func createDownloadRateLimiter(maxWorkers int) *rate.Limiter {
	maxWorkers = max(maxWorkers, 1)

	interval := constants.DownloadInterval
	perRequestInterval := interval
	if maxWorkers > 1 {
		perRequestInterval = interval / time.Duration(maxWorkers)
		if perRequestInterval <= 0 {
			perRequestInterval = time.Millisecond
		}
	}

	limit := rate.Every(perRequestInterval)
	b := max(maxWorkers, 1)

	return rate.NewLimiter(limit, b)
}
