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

func TestOverlapCommand(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, overlapCmd)
	assert.Equal(t, "overlap", overlapCmd.Use)
	assert.Contains(t, overlapCmd.Short, "overlap")
	assert.NotNil(t, overlapCmd.Run)
}

func TestComputeOverlapForType_NoSources(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_overlap.db")
	database, err := idb.Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	defer database.Close() // nolint: errcheck

	overlapRepo := idb.NewOverlapRepo(database)
	count := computeOverlapForType(ctx, logger, overlapRepo, "domain", "blocklist")
	assert.Equal(t, 0, count)
}

func TestComputeOverlapForType_WithOverlap(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_overlap2.db")
	database, err := idb.Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	defer database.Close() // nolint: errcheck

	// 2 sources
	_, err = database.Conn().Exec(`
		INSERT INTO dnstk_sources (name, disabled, skip_general_consolidation, skip_groups_consolidation, skip_categories_consolidation)
		VALUES ('source-a', 0, 0, 0, 0), ('source-b', 0, 0, 0, 0)`)
	require.NoError(t, err)

	_, err = database.Conn().Exec(`
		INSERT INTO dnstk_entries (source_id, entry, generic_source_type, actual_source_type, list_type, valid, must_consider)
		VALUES
			(1, 'shared.com', 'domain', 'domain', 'blocklist', 1, 0),
			(1, 'only-a.com', 'domain', 'domain', 'blocklist', 1, 0),
			(2, 'shared.com', 'domain', 'domain', 'blocklist', 1, 0),
			(2, 'only-b.com', 'domain', 'domain', 'blocklist', 1, 0),
			(2, 'also-b.com', 'domain', 'domain', 'blocklist', 1, 0)`)
	require.NoError(t, err)

	overlapRepo := idb.NewOverlapRepo(database)
	count := computeOverlapForType(ctx, logger, overlapRepo, "domain", "blocklist")
	assert.Equal(t, 1, count) // 1 pair with overlap

	results, err := overlapRepo.ListOverlapResults(ctx)
	require.NoError(t, err)
	assert.Len(t, results, 2) // both ways A to B and B to A

	var abResult *idb.OverlapResultRow
	for i := range results {
		if results[i].SourceName == "source-a" && results[i].TargetName == "source-b" {
			abResult = &results[i]
			break
		}
	}
	require.NotNil(t, abResult)
	assert.Equal(t, 1, abResult.OverlapCount)             // "shared.com"
	assert.Equal(t, 2, abResult.SourceCount)              // source-a has 2 entries
	assert.Equal(t, 3, abResult.TargetCount)              // source-b has 3 entries
	assert.InDelta(t, 50.0, abResult.OverlapPercent, 0.1) // 1/2 = 50%
}

func TestComputeOverlapForType_SingleSource(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_overlap3.db")
	database, err := idb.Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	defer database.Close() // nolint: errcheck

	// 1 source
	_, err = database.Conn().Exec(`
		INSERT INTO dnstk_sources (name, disabled, skip_general_consolidation, skip_groups_consolidation, skip_categories_consolidation)
		VALUES ('only-source', 0, 0, 0, 0)`)
	require.NoError(t, err)

	_, err = database.Conn().Exec(`
		INSERT INTO dnstk_entries (source_id, entry, generic_source_type, actual_source_type, list_type, valid, must_consider)
		VALUES (1, 'test.com', 'domain', 'domain', 'blocklist', 1, 0)`)
	require.NoError(t, err)

	overlapRepo := idb.NewOverlapRepo(database)
	count := computeOverlapForType(ctx, logger, overlapRepo, "domain", "blocklist")
	assert.Equal(t, 0, count)
}

func TestOverlapRepo_ClearOverlapResults(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_overlap_clear.db")
	database, err := idb.Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	defer database.Close() // nolint: errcheck

	repo := idb.NewOverlapRepo(database)

	err = repo.PersistOverlapResults(ctx, []idb.OverlapResultRow{{
		SourceName:        "a",
		TargetName:        "b",
		GenericSourceType: "domain",
		SourceListType:    "blocklist",
		TargetListType:    "blocklist",
		OverlapCount:      10,
		SourceCount:       100,
		TargetCount:       200,
		OverlapPercent:    10.0,
	}})
	require.NoError(t, err)

	err = repo.ClearOverlapResults()
	require.NoError(t, err)

	results, err := repo.ListOverlapResults(ctx)
	require.NoError(t, err)
	assert.Empty(t, results)
}
