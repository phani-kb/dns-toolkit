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

func TestTopEntriesCommand(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, topEntriesCmd)
	assert.Equal(t, "top", topEntriesCmd.Use)
	assert.Contains(t, topEntriesCmd.Short, "top entry")
	assert.NotNil(t, topEntriesCmd.Run)
}

func TestTopEntriesFlags(t *testing.T) {
	t.Parallel()

	flags := topEntriesCmd.Flags()

	minSourcesFlag := flags.Lookup("min-sources")
	assert.NotNil(t, minSourcesFlag)

	maxEntriesFlag := flags.Lookup("max-entries")
	assert.NotNil(t, maxEntriesFlag)
}

func TestTopEntriesRepo_GetTopEntries_Empty(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_top.db")

	database, err := idb.Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	defer database.Close() // nolint: errcheck

	repo := idb.NewTopEntriesRepo(database)

	entries, err := repo.GetTopEntries(ctx, "domain", "blocklist", 3, 100)
	assert.NoError(t, err)
	assert.Empty(t, entries)
}

func TestTopEntriesRepo_GetTopEntries_WithData(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_top2.db")

	database, err := idb.Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	defer database.Close() // nolint: errcheck

	// 3 sources
	for i := 1; i <= 3; i++ {
		_, err = database.WriteConn().Exec(
			`INSERT INTO dnstk_sources (name, disabled, skip_general_consolidation, skip_groups_consolidation, skip_categories_consolidation) VALUES (?, 0, 0, 0, 0)`,
			"source"+string(rune('0'+i)),
		)
		require.NoError(t, err)
	}

	// common.com in all 3, two.com in 2, one.com in 1
	_, err = database.WriteConn().
		Exec(`INSERT INTO dnstk_entries (source_id, entry, generic_source_type, actual_source_type, list_type, valid, must_consider)
		VALUES (1, 'common.com', 'domain', 'domain', 'blocklist', 1, 0),
		       (2, 'common.com', 'domain', 'domain', 'blocklist', 1, 0),
		       (3, 'common.com', 'domain', 'domain', 'blocklist', 1, 0),
		       (1, 'two.com', 'domain', 'domain', 'blocklist', 1, 0),
		       (2, 'two.com', 'domain', 'domain', 'blocklist', 1, 0),
		       (1, 'one.com', 'domain', 'domain', 'blocklist', 1, 0)`)
	require.NoError(t, err)

	repo := idb.NewTopEntriesRepo(database)

	entries, err := repo.GetTopEntries(ctx, "domain", "blocklist", 2, 100)
	assert.NoError(t, err)
	assert.Len(t, entries, 2)
	assert.Equal(t, "common.com", entries[0].Entry)
	assert.Equal(t, 3, entries[0].SourceCount)
	assert.Equal(t, "two.com", entries[1].Entry)
	assert.Equal(t, 2, entries[1].SourceCount)

	entries, err = repo.GetTopEntries(ctx, "domain", "blocklist", 3, 100)
	assert.NoError(t, err)
	assert.Len(t, entries, 1)
	assert.Equal(t, "common.com", entries[0].Entry)

	entries, err = repo.GetTopEntries(ctx, "domain", "blocklist", 2, 1)
	assert.NoError(t, err)
	assert.Len(t, entries, 1)
}

func TestTopEntriesRepo_PersistTopEntries(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_top3.db")

	database, err := idb.Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	defer database.Close() // nolint: errcheck

	repo := idb.NewTopEntriesRepo(database)

	entries := []idb.TopEntryRow{
		{Entry: "example.com", SourceCount: 5},
		{Entry: "test.org", SourceCount: 3},
	}

	err = repo.PersistTopEntries(ctx, "domain", "blocklist", 3, entries)
	assert.NoError(t, err)

	var count int
	err = database.ReadConn().QueryRow("SELECT COUNT(*) FROM dnstk_top_entries").Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, 2, count)

	err = repo.PersistTopEntries(ctx, "domain", "blocklist", 3, entries[:1])
	assert.NoError(t, err)

	err = database.ReadConn().QueryRow("SELECT COUNT(*) FROM dnstk_top_entries").Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, 1, count)
}
