package cmd

import (
	"context"
	"path/filepath"
	"testing"

	idb "github.com/phani-kb/dns-toolkit/internal/db"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadPreviousDownloadSummary_NilRepo(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)

	summary, err := loadPreviousDownloadSummary(logger, nil, "source", "path")
	assert.NoError(t, err)
	assert.Nil(t, summary)
}

func TestLoadPreviousDownloadSummary_NoSummaries(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_prev_download.db")

	database, err := idb.Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	defer database.Close() // nolint: errcheck

	repo := idb.NewDownloadsRepo(database)

	summary, err := loadPreviousDownloadSummary(logger, repo, "nonexistent", "")
	assert.NoError(t, err)
	assert.Nil(t, summary)

	summary, err = loadPreviousDownloadSummary(logger, repo, "test", "/tmp/test.txt")
	assert.NoError(t, err)
	assert.Nil(t, summary)
}

func TestLoadPreviousDownloadSummary_WithFilepathFilter(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_prev_download2.db")

	database, err := idb.Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	defer database.Close() // nolint: errcheck

	_, err = database.Conn().Exec(
		`insert into dnstk_sources (name, disabled, skip_general_consolidation, skip_groups_consolidation, skip_categories_consolidation) values ('test-source', 0, 0, 0, 0)`,
	)
	require.NoError(t, err)

	_, err = database.Conn().Exec(
		`insert into dnstk_downloads (source_id, url, filepath, checksum) values (1, 'https://example.com/list.txt', '/tmp/download/test.txt', 'abc123')`,
	)
	require.NoError(t, err)

	repo := idb.NewDownloadsRepo(database)

	firstSummary, err := loadPreviousDownloadSummary(logger, repo, "test-source", "/tmp/download/test.txt")
	assert.NoError(t, err)
	assert.NotNil(t, firstSummary)

	secSummary, err := loadPreviousDownloadSummary(logger, repo, "test-source", "/tmp/other.txt")
	assert.NoError(t, err)
	assert.NotNil(t, secSummary)

	assert.Equal(t, firstSummary, secSummary)

	secSummary, err = loadPreviousDownloadSummary(logger, repo, "test-source", "")
	assert.NoError(t, err)
	assert.NotNil(t, secSummary)

	assert.Equal(t, secSummary, firstSummary)
}
