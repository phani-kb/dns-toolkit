package cmd

import (
	"context"
	"os"
	"path/filepath"
	"sync"
	"testing"

	con "github.com/phani-kb/dns-toolkit/internal/consolidators"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	idb "github.com/phani-kb/dns-toolkit/internal/db"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createTempFileWithContent(t *testing.T, content string) string {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	err := os.WriteFile(tmpFile, []byte(content), 0o644)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	return tmpFile
}

func TestFilterEntriesWithAllowlist(t *testing.T) {
	tests := []struct {
		name            string
		blockEntries    []string
		allowEntries    []string
		expectedKept    int
		expectedIgnored int
	}{
		{
			name:            "no overlap",
			blockEntries:    []string{"a.com", "b.com", "c.com"},
			allowEntries:    []string{"x.com", "y.com"},
			expectedKept:    3,
			expectedIgnored: 0,
		},
		{
			name:            "complete overlap",
			blockEntries:    []string{"a.com", "b.com"},
			allowEntries:    []string{"a.com", "b.com"},
			expectedKept:    0,
			expectedIgnored: 2,
		},
		{
			name:            "partial overlap",
			blockEntries:    []string{"a.com", "b.com", "c.com"},
			allowEntries:    []string{"b.com"},
			expectedKept:    2,
			expectedIgnored: 1,
		},
		{
			name:            "empty blocklist",
			blockEntries:    []string{},
			allowEntries:    []string{"a.com"},
			expectedKept:    0,
			expectedIgnored: 0,
		},
		{
			name:            "empty allowlist",
			blockEntries:    []string{"a.com", "b.com"},
			allowEntries:    []string{},
			expectedKept:    2,
			expectedIgnored: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blockSet := u.NewStringSet(tt.blockEntries)
			allowSet := u.NewStringSet(tt.allowEntries)

			filtered, ignored := filterEntriesWithAllowlist(blockSet, allowSet)

			assert.Equal(t, tt.expectedKept, filtered.Size())
			assert.Equal(t, tt.expectedIgnored, ignored.Size())
		})
	}
}

func TestFilterEntriesWithAllowlist_MustConsider(t *testing.T) {
	blockSet := u.NewStringSet([]string{})
	blockSet.AddWithConsider("important.com", true) // must_consider = true
	blockSet.AddWithConsider("normal.com", false)

	allowSet := u.NewStringSet([]string{})
	allowSet.AddWithConsider("important.com", false) // must_consider = false
	allowSet.AddWithConsider("normal.com", false)

	filtered, ignored := filterEntriesWithAllowlist(blockSet, allowSet)

	// important.com should be kept because block has must_consider=true, allow has must_consider=false
	assert.True(t, filtered.Contains("important.com"))
	// normal.com should be ignored because allow wins when both are false
	assert.True(t, ignored.Contains("normal.com"))
}

func TestConsolidateCategoriesCommand(t *testing.T) {
	assert.NotNil(t, consolidateCategoriesCmd)
	assert.Equal(t, "categories", consolidateCategoriesCmd.Use)
	assert.Contains(t, consolidateCategoriesCmd.Short, "category-based")
	assert.NotNil(t, consolidateCategoriesCmd.Run)
}

func TestProcessCategoryConsolidation_EmptyDatabase(t *testing.T) {
	logger := multilog.NewLogger()
	con.InitForTesting()

	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test.db")
	database, err := idb.Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	defer database.Close() // nolint: errcheck

	entriesRepo := idb.NewEntriesRepo(database)
	consolidatedRepo := idb.NewConsolidatedRepo(database)
	var persistMu sync.Mutex

	processCategoryConsolidation(
		ctx,
		logger,
		"ads",
		[]string{constants.SourceTypeDomain},
		entriesRepo,
		consolidatedRepo,
		&persistMu,
	)

	count, err := consolidatedRepo.GetConsolidatedCount(
		constants.SourceTypeDomain,
		constants.ListTypeBlocklist,
		"category",
		true,
	)
	require.NoError(t, err)
	assert.Equal(t, int64(0), count)
}
