package cmd

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	idb "github.com/phani-kb/dns-toolkit/internal/db"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProcessAllSourcesUsesDatabaseDownloadSummaries(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	tempDir := t.TempDir()
	downloadDir := filepath.Join(tempDir, "download")
	processedDir := filepath.Join(tempDir, "processed")
	summaryDir := filepath.Join(tempDir, "summary")
	dbPath := filepath.Join(tempDir, "dns-toolkit.db")

	require.NoError(t, os.MkdirAll(downloadDir, 0o755))
	require.NoError(t, os.MkdirAll(processedDir, 0o755))
	require.NoError(t, os.MkdirAll(summaryDir, 0o755))

	database, err := idb.Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	t.Cleanup(func() { _ = database.Close() })

	source := config.Source{
		Name:      "db-source",
		URL:       "https://example.com/list.txt",
		Frequency: "daily",
		Types: []c.SourceType{
			{
				Name: "domain",
				ListTypes: []c.ListType{
					{Name: "blocklist", MustConsider: true},
				},
			},
		},
		Categories: []string{"malware"},
	}

	sourcesRepo := idb.NewSourcesRepo(database)
	_, _, err = sourcesRepo.ImportSourcesFromConfig(
		ctx,
		logger,
		config.SourcesConfig{Sources: []config.Source{source}},
		runtimeSourceFile,
	)
	require.NoError(t, err)

	sourceID, err := sourcesRepo.GetSourceIDByName(source.Name)
	require.NoError(t, err)
	require.NotZero(t, sourceID)

	downloadFilePath := filepath.Join(downloadDir, "db-source.txt")
	require.NoError(t, os.WriteFile(downloadFilePath, []byte("example.com\ninvalid_entry\n"), 0o644))

	downloadsRepo := idb.NewDownloadsRepo(database)
	err = downloadsRepo.UpsertDownload(idb.DownloadRow{
		SourceID:              sourceID,
		TypeCount:             1,
		URL:                   source.URL,
		Filepath:              downloadFilePath,
		Frequency:             source.Frequency,
		LastDownloadTimestamp: time.Now().Format(constants.TimestampFormat),
	})
	require.NoError(t, err)

	originalAppConfig := AppConfig
	originalSourcesConfigs := SourcesConfigs
	originalDownloadDir := constants.DownloadDir
	originalProcessedDir := constants.ProcessedDir
	originalSummaryDir := constants.SummaryDir
	defer func() {
		AppConfig = originalAppConfig
		SourcesConfigs = originalSourcesConfigs
		constants.DownloadDir = originalDownloadDir
		constants.ProcessedDir = originalProcessedDir
		constants.SummaryDir = originalSummaryDir
	}()

	AppConfig = &config.AppConfig{
		DNSToolkit: config.DNSToolkitConfig{
			Database:   config.DatabaseConfig{Path: dbPath},
			MaxWorkers: 1,
		},
	}
	SourcesConfigs = []config.SourcesConfig{{Sources: []config.Source{source}}}
	constants.DownloadDir = downloadDir
	constants.ProcessedDir = processedDir
	constants.SummaryDir = summaryDir

	processAllSources(ctx, logger, processedDir, false, true)

	processedRepo := idb.NewProcessedRepo(database)
	processedSummaries, err := processedRepo.ListProcessedSummaries(processedDir)
	require.NoError(t, err)
	require.Len(t, processedSummaries, 1)
	assert.Equal(t, "db-source", processedSummaries[0].Name)
	assert.NotEmpty(t, processedSummaries[0].ValidFiles)

	var entryCount, validCount, invalidCount int
	require.NoError(t, database.ReadConn().QueryRow(
		"SELECT COUNT(*) FROM "+constants.TableEntries+" WHERE source_id = ?",
		sourceID,
	).Scan(&entryCount))
	require.NoError(t, database.ReadConn().QueryRow(
		"SELECT COUNT(*) FROM "+constants.TableEntries+" WHERE source_id = ? AND valid = 1",
		sourceID,
	).Scan(&validCount))
	require.NoError(t, database.ReadConn().QueryRow(
		"SELECT COUNT(*) FROM "+constants.TableEntries+" WHERE source_id = ? AND valid = 0",
		sourceID,
	).Scan(&invalidCount))

	assert.Equal(t, 2, entryCount)
	assert.Equal(t, 1, validCount)
	assert.Equal(t, 1, invalidCount)

	var categoryCount int
	require.NoError(t, database.ReadConn().QueryRow(
		"SELECT COUNT(*) FROM "+constants.TableEntryCategories+" WHERE source_id = ?",
		sourceID,
	).Scan(&categoryCount))
	assert.Equal(t, 1, categoryCount)
}
