package cmd

import (
	"context"
	"path/filepath"
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	idb "github.com/phani-kb/dns-toolkit/internal/db"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDBWithEntries(t *testing.T, logger *multilog.Logger) (*idb.DB, map[string]int64) {
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_resolution.db")

	database, err := idb.Open(ctx, logger, dbPath, true)
	require.NoError(t, err)

	t.Cleanup(func() { _ = database.Close() })

	sourcesRepo := idb.NewSourcesRepo(database)
	sourceConfigs := config.SourcesConfig{
		Sources: []config.Source{
			{
				Name:      "BlockSource1",
				URL:       "https://example.com/block1.txt",
				Frequency: "daily",
				Types: []c.SourceType{{
					Name:      "domain",
					ListTypes: []c.ListType{{Name: "blocklist"}},
				}},
			},
			{
				Name:      "BlockSource2",
				URL:       "https://example.com/block2.txt",
				Frequency: "daily",
				Types: []c.SourceType{{
					Name:      "domain",
					ListTypes: []c.ListType{{Name: "blocklist"}},
				}},
			},
			{
				Name:      "AllowSource1",
				URL:       "https://example.com/allow1.txt",
				Frequency: "daily",
				Types: []c.SourceType{{
					Name:      "domain",
					ListTypes: []c.ListType{{Name: "allowlist"}},
				}},
			},
			{
				Name:      "AllowSource2",
				URL:       "https://example.com/allow2.txt",
				Frequency: "daily",
				Types: []c.SourceType{{
					Name:      "domain",
					ListTypes: []c.ListType{{Name: "allowlist"}},
				}},
			},
		},
	}

	_, _, err = sourcesRepo.ImportSourcesFromConfig(ctx, logger, sourceConfigs, "test.json")
	require.NoError(t, err)

	sourceIDs := make(map[string]int64)
	for _, src := range sourceConfigs.Sources {
		id, err := sourcesRepo.GetSourceIDByName(src.Name)
		require.NoError(t, err)
		sourceIDs[src.Name] = id
	}

	return database, sourceIDs
}

func insertTestEntries(t *testing.T, database *idb.DB, sourceIDs map[string]int64, entries map[string][]string) {
	ctx := context.Background()
	entriesRepo := idb.NewEntriesRepo(database)

	for sourceName, entryList := range entries {
		sourceID := sourceIDs[sourceName]
		require.NotZero(t, sourceID, "source %s not found", sourceName)

		listType := constants.ListTypeBlocklist
		if sourceName[:5] == "Allow" {
			listType = constants.ListTypeAllowlist
		}

		entryRows := make([]idb.EntryRow, 0, len(entryList))
		for _, entry := range entryList {
			entryRows = append(entryRows, idb.EntryRow{
				Entry:             entry,
				GenericSourceType: constants.SourceTypeDomain,
				ActualSourceType:  constants.SourceTypeDomain,
				ListType:          listType,
				Valid:             true,
				MustConsider:      false,
			})
		}

		err := entriesRepo.ReplaceSourceData(ctx, sourceID, entryRows, nil, nil)
		require.NoError(t, err)
	}
}

func TestGetResolutionSets(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()

	dbPath := filepath.Join(t.TempDir(), "test.db")
	database, err := idb.Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	defer database.Close() // nolint: errcheck

	allow1, block1, conflicts1, _, _, _, err := GetResolutionSets(logger, database)
	assert.NoError(t, err)
	assert.NotNil(t, allow1)
	assert.NotNil(t, block1)
	assert.NotNil(t, conflicts1)

	_, _, _, _, _, _, err = GetResolutionSets(logger, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "database is required")
}

func TestBuildResolutionSets_BlocklistWins(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	database, sourceIDs := setupTestDBWithEntries(t, logger)

	insertTestEntries(t, database, sourceIDs, map[string][]string{
		"BlockSource1": {"conflict.com", "block-only.com"},
		"BlockSource2": {"conflict.com"},
		"AllowSource1": {"conflict.com", "allow-only.com"},
	})

	allowByType, blockByType, conflicts, _, _, _, err := BuildResolutionSets(logger, database)
	require.NoError(t, err)

	assert.True(t, blockByType[constants.SourceTypeDomain].Contains("conflict.com"),
		"conflict.com should be in blocklist (2 block > 1 allow)")
	assert.False(t, allowByType[constants.SourceTypeDomain].Contains("conflict.com"),
		"conflict.com should NOT be in allowlist")

	assert.True(t, blockByType[constants.SourceTypeDomain].Contains("block-only.com"),
		"block-only.com should be in blocklist")

	assert.True(t, allowByType[constants.SourceTypeDomain].Contains("allow-only.com"),
		"allow-only.com should be in allowlist")

	assert.Empty(t, conflicts, "should have no conflicts when counts differ")
}

func TestBuildResolutionSets_AllowlistWins(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	database, sourceIDs := setupTestDBWithEntries(t, logger)

	insertTestEntries(t, database, sourceIDs, map[string][]string{
		"BlockSource1": {"conflict.com"},
		"AllowSource1": {"conflict.com"},
		"AllowSource2": {"conflict.com"},
	})

	allowByType, blockByType, conflicts, _, _, _, err := BuildResolutionSets(logger, database)
	require.NoError(t, err)

	assert.True(t, allowByType[constants.SourceTypeDomain].Contains("conflict.com"),
		"conflict.com should be in allowlist (2 allow > 1 block)")
	assert.False(t, blockByType[constants.SourceTypeDomain].Contains("conflict.com"),
		"conflict.com should NOT be in blocklist")
	assert.Empty(t, conflicts, "should have no conflicts when counts differ")
}

func TestBuildResolutionSets_EqualCountsConflict(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	database, sourceIDs := setupTestDBWithEntries(t, logger)

	// Insert test entries: "conflict.com" is in 1 blocklist and 1 allowlist (equal)
	insertTestEntries(t, database, sourceIDs, map[string][]string{
		"BlockSource1": {"conflict.com"},
		"AllowSource1": {"conflict.com"},
	})

	// Clear cache

	allowByType, blockByType, conflicts, _, _, _, err := BuildResolutionSets(logger, database)
	require.NoError(t, err)

	// "conflict.com" should be in conflicts (1 block == 1 allow)
	assert.False(t, allowByType[constants.SourceTypeDomain].Contains("conflict.com"),
		"conflict.com should NOT be in allowlist")
	assert.False(t, blockByType[constants.SourceTypeDomain].Contains("conflict.com"),
		"conflict.com should NOT be in blocklist")
	assert.Len(t, conflicts, 1, "should have 1 conflict")
	assert.Equal(t, "conflict.com", conflicts[0].Entry)
	assert.Equal(t, 1, conflicts[0].BlockCount)
	assert.Equal(t, 1, conflicts[0].AllowCount)
}

// TestBuildResolutionSets_ThresholdBehavior tests that thresholds are applied correctly.
func TestBuildResolutionSets_ThresholdBehavior(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	database, sourceIDs := setupTestDBWithEntries(t, logger)

	// Insert: "maybe.com" has 2 allow sources and 1 block source
	insertTestEntries(t, database, sourceIDs, map[string][]string{
		"BlockSource1": {"maybe.com"},
		"AllowSource1": {"maybe.com"},
		"AllowSource2": {"maybe.com"},
	})

	// Clear cache

	// Set threshold: require 3 allowlist sources to win
	oldAppConfig := AppConfig
	AppConfig = &config.AppConfig{}
	AppConfig.DNSToolkit.Override.Enabled = true
	AppConfig.DNSToolkit.Override.Thresholds = []config.ThresholdConfig{
		{Name: "allowlist", MinSources: 3},
		{Name: "blocklist", MinSources: 1},
	}
	defer func() { AppConfig = oldAppConfig }()

	allowByType, blockByType, conflicts, _, _, _, err := BuildResolutionSets(logger, database)
	require.NoError(t, err)

	// With threshold=3 for allow, 2 allow sources is not enough -> conflict
	assert.False(t, allowByType[constants.SourceTypeDomain].Contains("maybe.com"),
		"maybe.com should NOT be allowed (below threshold)")
	assert.False(t, blockByType[constants.SourceTypeDomain].Contains("maybe.com"),
		"maybe.com should NOT be blocked (allow count is higher)")
	assert.Len(t, conflicts, 1, "should have 1 conflict due to threshold")
}
