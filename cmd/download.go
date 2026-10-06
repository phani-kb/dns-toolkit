package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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

type downloadStats struct {
	mu              sync.Mutex
	successCount    int
	failCount       int
	downloadedCount int
}

func (s *downloadStats) recordSuccess(downloaded bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.successCount++
	if downloaded {
		s.downloadedCount++
	}
}

func (s *downloadStats) recordFailure() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.failCount++
}

type downloadJob struct {
	sourcesRepo   *db.SourcesRepo
	downloadsRepo *db.DownloadsRepo
	appConfig     *cfg.AppConfig
	stats         *downloadStats
	source        cfg.Source
	maxRetries    int
}

// run performs the full download lifecycle for a single source
func (j *downloadJob) run() {
	source := j.source

	sourceID, sourceIDErr := j.sourcesRepo.GetSourceIDByName(source.Name)
	if sourceIDErr != nil {
		Logger.Warnf("Failed to get source ID for %s: %v", source.Name, sourceIDErr)
	}

	persist := func(summary c.DownloadSummary, persistedPath string) {
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
		if err := j.downloadsRepo.UpsertDownload(downloadRow); err != nil {
			Logger.Warnf("Failed to upsert download record for %s: %v", source.Name, err)
		}
	}

	downloadFile, err := source.GetDownloadFile(Logger, constants.DownloadDir)
	if err != nil {
		Logger.Errorf("Getting download file error: %v", err)
		j.stats.recordFailure()
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
		persist(summary, "")
		return
	}

	downloader := j.selectDownloader()

	var skipCertVerification bool
	var skipCertVerificationHosts []string
	var applicationConfig cfg.ApplicationConfig
	if j.appConfig != nil {
		skipCertVerification = j.appConfig.DNSToolkit.SkipCertVerification
		skipCertVerificationHosts = j.appConfig.DNSToolkit.SkipCertVerificationHosts
		applicationConfig = j.appConfig.Application
	}

	filePath, fetchSkipped, err := downloader.Download(
		Logger,
		downloadFile,
		skipCertVerification,
		skipCertVerificationHosts,
		applicationConfig,
	)

	summary := c.DownloadSummary{
		Name:                        source.Name,
		URL:                         downloadFile.URL,
		TypeCount:                   source.TypeCount,
		Types:                       source.Types,
		Filepath:                    filePath,
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

		j.stats.recordFailure()
		persist(summary, filePath)
		return
	}

	j.stats.recordSuccess(!fetchSkipped)

	if fetchSkipped {
		summary.LastCheckedTimestamp = u.GetTimestamp()
		if info, statErr := os.Stat(filePath); statErr == nil {
			summary.LastDownloadTimestamp = info.ModTime().Format(constants.TimestampFormat)
		} else {
			Logger.Errorf("Getting file info error: %v", statErr)
		}
	} else {
		summary.LastDownloadTimestamp = time.Now().Format(constants.TimestampFormat)
	}

	j.reprocessTargets(downloader, downloadFile, filePath, fetchSkipped, &summary)

	if j.appConfig != nil && j.appConfig.DNSToolkit.FilesChecksum.Enabled {
		summary.Checksum = u.CalculateChecksum(Logger, filePath, j.appConfig.DNSToolkit.FilesChecksum.Algorithm)
	}

	persist(summary, filePath)
}

func (j *downloadJob) reprocessTargets(
	downloader d.Downloader,
	downloadFile c.DownloadFile,
	filePath string,
	fetchSkipped bool,
	summary *c.DownloadSummary,
) {
	for _, target := range downloadFile.Targets {
		targetFilePath := filepath.Join(target.TargetFolder, target.TargetFile)

		shouldReprocess := !fetchSkipped

		if fetchSkipped {
			if prevSummary, err := loadPreviousDownloadSummary(
				Logger,
				j.downloadsRepo,
				summary.Name,
				targetFilePath,
			); err == nil && prevSummary != nil {
				if prevSummary.CountToConsider != summary.CountToConsider {
					Logger.Infof("Count to consider changed for %s: %d -> %d, re-processing...",
						summary.Name, prevSummary.CountToConsider, summary.CountToConsider)
					shouldReprocess = true

					if downloadFile.IsArchive {
						Logger.Debugf("Re-extracting archive and copying target file for %s", summary.Name)
						if copyErr := u.ForceCopySourceToTarget(Logger, target); copyErr != nil {
							Logger.Errorf("Failed to force re-copy target file for %s: %v", summary.Name, copyErr)
							summary.Error = fmt.Sprintf("Force re-copy target file error: %v", copyErr)
							shouldReprocess = false
						}
					}
				}
			} else if err != nil {
				Logger.Debugf("Could not load previous summary for %s: %v", summary.Name, err)
				shouldReprocess = true
			} else {
				shouldReprocess = true
			}
		}

		if shouldReprocess {
			if err := downloader.PostDownloadProcess(Logger, targetFilePath, summary.CountToConsider); err != nil {
				Logger.Errorf("Post download process error for %s: %v", summary.Name, err)
				summary.Error = fmt.Sprintf("Post-download processing error: %v", err)
			}
		}
	}
}

func (j *downloadJob) selectDownloader() d.Downloader {
	name := j.source.Downloader
	if name != "" {
		if downloader, exists := d.GetDownloader(name); exists {
			Logger.Debugf("Using registered downloader %q for %s", name, j.source.Name)
			return downloader
		}
	}

	downloader, _ := d.GetDownloader(d.DefaultDownloaderName())
	Logger.Debugf("Using default downloader with %d retries for %s", j.maxRetries, j.source.Name)
	return downloader
}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download enabled sources",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		database := openDB(ctx)
		defer database.CloseLogError(Logger)

		forceDownload := forceDownloadRequested(cmd)

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
			defaultDownloader.SetForceDownload(forceDownload)
			if initErr := d.RegisterDownloader(defaultDownloader); initErr != nil {
				Logger.Warnf("Failed to register default downloader with retry settings: %v", initErr)
			}
		}

		domainTopDownloader := d.NewDomainTopDownloaderWithRetries(maxRetries)
		domainTopDownloader.SetForceDownload(forceDownload)
		if domainTopErr := d.RegisterDownloader(domainTopDownloader); domainTopErr != nil {
			Logger.Warnf("Failed to register domain top downloader: %v", domainTopErr)
		}

		sourcesRepo := db.NewSourcesRepo(database)
		downloadsRepo := db.NewDownloadsRepo(database)
		defaultDownloader.SetSummaryProvider(downloadsRepo)
		domainTopDownloader.SetSummaryProvider(downloadsRepo)
		if err := syncSourcesToDB(ctx, Logger, sourcesRepo, SourcesConfigs); err != nil {
			Logger.Errorf("Failed to sync source definitions to database: %v", err)
			os.Exit(1)
		}

		maxWorkers := runtime.GOMAXPROCS(0)
		if AppConfig != nil && AppConfig.DNSToolkit.MaxWorkers > 0 {
			maxWorkers = AppConfig.DNSToolkit.MaxWorkers
		}
		maxWorkers = max(maxWorkers, 1)
		Logger.Infof("Using worker pool with %d worker(s) for downloads", maxWorkers)
		defaultInterval := getDefaultDownloadInterval()
		defaultBurst := getDefaultDownloadBurst(maxWorkers)
		defaultLimiter := createDownloadRateLimiter(maxWorkers, defaultInterval, defaultBurst)
		workerPool := c.NewDTWorkerPool(maxWorkers)

		stats := &downloadStats{}
		var totalSources int

		for _, sourcesConfig := range SourcesConfigs {
			var sourceFilters cfg.SourceFilters
			if AppConfig != nil {
				sourceFilters = AppConfig.DNSToolkit.SourceFilters
			}
			for _, source := range sourcesConfig.GetEnabledSources(sourceFilters) {
				totalSources++
				source := source // local copy for goroutine
				job := &downloadJob{
					source:        source,
					sourcesRepo:   sourcesRepo,
					downloadsRepo: downloadsRepo,
					appConfig:     AppConfig,
					maxRetries:    maxRetries,
					stats:         stats,
				}
				workerPool.Submit(func() {
					if !strings.HasPrefix(strings.TrimSpace(source.URL), "file://") {
						if defaultLimiter != nil {
							if err := defaultLimiter.Wait(context.Background()); err != nil {
								Logger.Warnf("Rate limiter wait failed for source %s: %v", source.Name, err)
							}
						}
					}
					job.run()
				})
			}
		}

		workerPool.Wait()

		stats.mu.Lock()
		successCount, downloadedCount, failCount := stats.successCount, stats.downloadedCount, stats.failCount
		stats.mu.Unlock()

		Logger.Infof("Download complete: %d sources processed, %d successful (%d downloaded, %d skipped), %d failed",
			totalSources, successCount, downloadedCount, successCount-downloadedCount, failCount)
	},
}

func init() {
	downloadCmd.Flags().Bool("force", false, "Force re-download of all sources (ignores existing summaries)")
}

func forceDownloadRequested(cmd *cobra.Command) bool {
	forceFlag, err := cmd.Flags().GetBool("force")
	if err != nil {
		Logger.Warnf("Failed to parse --force flag (defaulting to false): %v", err)
		forceFlag = false
	}
	forceEnv := os.Getenv("DNS_TOOLKIT_FORCE_DOWNLOAD") == "true" || os.Getenv("DNS_TOOLKIT_FORCE_DOWNLOAD") == "1"
	return forceFlag || forceEnv
}

func createDownloadRateLimiter(maxWorkers int, interval time.Duration, burst int) *rate.Limiter {
	maxWorkers = max(maxWorkers, 1)

	perRequestInterval := interval
	if maxWorkers > 1 {
		perRequestInterval = interval / time.Duration(maxWorkers)
		if perRequestInterval <= 0 {
			perRequestInterval = time.Millisecond
		}
	}

	limit := rate.Every(perRequestInterval)
	b := max(maxWorkers, 1)
	if burst > 0 {
		b = burst
	}

	return rate.NewLimiter(limit, b)
}

func getDefaultDownloadInterval() time.Duration {
	if AppConfig != nil && AppConfig.DNSToolkit.DownloadRateLimit.IntervalMs > 0 {
		return time.Duration(AppConfig.DNSToolkit.DownloadRateLimit.IntervalMs) * time.Millisecond
	}

	return constants.DownloadInterval
}

func getDefaultDownloadBurst(maxWorkers int) int {
	if AppConfig != nil && AppConfig.DNSToolkit.DownloadRateLimit.Burst > 0 {
		return AppConfig.DNSToolkit.DownloadRateLimit.Burst
	}

	return max(maxWorkers, 1)
}
