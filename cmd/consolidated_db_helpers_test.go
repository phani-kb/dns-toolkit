package cmd

import (
	"context"
	"path/filepath"
	"sync"
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	idb "github.com/phani-kb/dns-toolkit/internal/db"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPersistConsolidatedEntries_Empty(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()

	err := persistConsolidatedEntries(
		ctx, logger, nil, nil,
		u.NewStringSet([]string{}),
		constants.SourceTypeDomain, constants.ListTypeBlocklist,
		"general", "", "", true,
	)
	assert.NoError(t, err)

	err = persistConsolidatedEntries(
		ctx, logger, nil, nil,
		nil,
		constants.SourceTypeDomain, constants.ListTypeBlocklist,
		"general", "", "", true,
	)
	assert.NoError(t, err)
}

func TestPersistConsolidatedEntries_WithEntries(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_persist.db")

	database, err := idb.Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	defer database.Close() // nolint: errcheck

	repo := idb.NewConsolidatedRepo(database)
	var mu sync.Mutex

	entries := u.NewStringSet([]string{"test1.com", "test2.com", "test3.com"})

	err = persistConsolidatedEntries(
		ctx, logger, repo, &mu,
		entries,
		constants.SourceTypeDomain, constants.ListTypeBlocklist,
		"general", "", "", true,
	)
	assert.NoError(t, err)
}

func TestPersistConsolidatedEntries_WithGroupAndCategory(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_persist_group.db")

	database, err := idb.Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	defer database.Close() // nolint: errcheck

	repo := idb.NewConsolidatedRepo(database)

	entries := u.NewStringSet([]string{"entry1.com"})

	// group
	err = persistConsolidatedEntries(
		ctx, logger, repo, nil,
		entries,
		constants.SourceTypeDomain, constants.ListTypeBlocklist,
		"group", "mini", "", true,
	)
	assert.NoError(t, err)

	// category
	err = persistConsolidatedEntries(
		ctx, logger, repo, nil,
		entries,
		constants.SourceTypeDomain, constants.ListTypeAllowlist,
		"category", "", "ads", true,
	)
	assert.NoError(t, err)
}

func TestOpenConsolidatedRepo(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	tmpDir := t.TempDir()

	oldAppConfig := AppConfig
	AppConfig = &config.AppConfig{
		DNSToolkit: config.DNSToolkitConfig{
			Database: config.DatabaseConfig{
				Path: filepath.Join(tmpDir, "test_open.db"),
			},
		},
	}
	defer func() { AppConfig = oldAppConfig }()

	database, repo, err := openConsolidatedRepo(ctx, logger, "general")
	require.NoError(t, err)
	require.NotNil(t, database)
	require.NotNil(t, repo)
	defer database.Close() // nolint: errcheck
}

func TestNewConsolidationManager(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_mgr.db")

	database, err := idb.Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	defer database.Close() // nolint: errcheck

	manager := NewConsolidationManager(logger, database)
	assert.NotNil(t, manager)
	assert.Equal(t, logger, manager.logger)
	assert.Equal(t, database, manager.database)
}

func TestGenerateConflictReport(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_conflict_report.db")
	cleanup, _ := setupTestEnvironmentForCmdTests(t)
	defer cleanup()

	database, err := idb.Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	defer database.Close() // nolint: errcheck

	resolutionCacheMu.Lock()
	resolutionCacheKey = ""
	resolutionCachedAllow = nil
	resolutionCacheMu.Unlock()

	manager := NewConsolidationManager(logger, database)

	// empty processed files
	processedFiles := []c.ProcessedFile{}
	err = manager.GenerateConflictReport(processedFiles)
	assert.NoError(t, err)
}
