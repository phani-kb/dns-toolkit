package cmd

import (
	"context"
	"path/filepath"
	"strings"
	"testing"

	idb "github.com/phani-kb/dns-toolkit/internal/db"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateFileName(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()

	tests := []struct {
		name       string
		fileName   string
		sourceType string
		listType   string
		entryType  string
		expected   string
	}{
		{
			name:       "standard filename generation",
			fileName:   "example",
			sourceType: "domain",
			listType:   "blocklist",
			entryType:  "valid",
			expected:   "example_domain_BL_valid_",
		},
		{
			name:       "filename with allowlist",
			fileName:   "test",
			sourceType: "ip",
			listType:   "allowlist",
			entryType:  "invalid",
			expected:   "test_ip_AL_invalid_",
		},
		{
			name:       "unknown list type",
			fileName:   "unknown",
			sourceType: "custom",
			listType:   "unknown_type",
			entryType:  "valid",
			expected:   "unknown_custom_unknown_type_valid_",
		},
		{
			name:       "empty inputs",
			fileName:   "",
			sourceType: "",
			listType:   "",
			entryType:  "",
			expected:   "_____",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateFileName(logger, tt.fileName, tt.sourceType, tt.listType, tt.entryType)
			// Check that the result starts with the expected prefix
			assert.True(t, len(result) > len(tt.expected), "Result should include hash suffix")
			assert.True(t, len(result) > 30, "Result should include MD5 hash (32 chars)")
			assert.Contains(t, result, tt.expected[:len(tt.expected)-1], "Result should contain expected prefix")
			assert.True(t, strings.HasSuffix(result, ".txt"), "Result should end with .txt")
		})
	}
}

func TestProcessAllowlists_EmptyDatabase(t *testing.T) {
	ctx := context.Background()
	logger := multilog.NewLogger()
	dbPath := filepath.Join(t.TempDir(), "test_consolidate.db")
	database, err := idb.Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	defer database.Close() // nolint: errcheck

	entriesRepo := idb.NewEntriesRepo(database)
	consolidatedRepo := idb.NewConsolidatedRepo(database)

	genericSourceTypes := []string{"domain"}
	allowlistEntriesByType := make(map[string]u.StringSet)
	resolvedBlockByType := make(map[string]u.StringSet)

	processAllowlists(
		ctx,
		logger,
		entriesRepo,
		consolidatedRepo,
		genericSourceTypes,
		resolvedBlockByType,
		allowlistEntriesByType,
	)

	assert.Contains(t, allowlistEntriesByType, "domain")
	assert.Equal(t, 0, allowlistEntriesByType["domain"].Size())
}

func TestProcessAllowlists_WithEntries(t *testing.T) {
	ctx := context.Background()
	logger := multilog.NewLogger()
	dbPath := filepath.Join(t.TempDir(), "test_consolidate.db")
	database, err := idb.Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	defer database.Close() // nolint: errcheck

	_, err = database.WriteConn().
		Exec(`insert into dnstk_sources (name, disabled, skip_general_consolidation, skip_groups_consolidation, skip_categories_consolidation) values ('test-source', 0, 0, 0, 0)`)
	require.NoError(t, err)

	_, err = database.WriteConn().
		Exec(`insert into dnstk_entries (source_id, entry, generic_source_type, actual_source_type, list_type, valid, must_consider)
		values (1, 'allow1.com', 'domain', 'domain', 'allowlist', 1, 0),
		       (1, 'allow2.com', 'domain', 'domain', 'allowlist', 1, 0),
		       (1, 'must-keep.com', 'domain', 'domain', 'allowlist', 1, 1)`)
	require.NoError(t, err)

	entriesRepo := idb.NewEntriesRepo(database)
	consolidatedRepo := idb.NewConsolidatedRepo(database)

	genericSourceTypes := []string{"domain"}
	allowlistEntriesByType := make(map[string]u.StringSet)

	resolvedBlockByType := map[string]u.StringSet{
		"domain": u.NewStringSet([]string{"allow1.com"}),
	}

	processAllowlists(
		ctx,
		logger,
		entriesRepo,
		consolidatedRepo,
		genericSourceTypes,
		resolvedBlockByType,
		allowlistEntriesByType,
	)

	domainEntries := allowlistEntriesByType["domain"]
	assert.Equal(t, 2, domainEntries.Size())
	assert.True(t, domainEntries.Contains("allow2.com"))
	assert.True(t, domainEntries.Contains("must-keep.com"))
	assert.False(t, domainEntries.Contains("allow1.com"))
}

func TestProcessAllowlists_MustConsiderOverridesResolution(t *testing.T) {
	ctx := context.Background()
	logger := multilog.NewLogger()
	dbPath := filepath.Join(t.TempDir(), "test_consolidate.db")
	database, err := idb.Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	defer database.Close() // nolint: errcheck

	_, err = database.WriteConn().
		Exec(`insert into dnstk_sources (name, disabled, skip_general_consolidation, skip_groups_consolidation, skip_categories_consolidation) values ('test-source', 0, 0, 0, 0)`)
	require.NoError(t, err)

	_, err = database.WriteConn().
		Exec(`insert into dnstk_entries (source_id, entry, generic_source_type, actual_source_type, list_type, valid, must_consider)
		values (1, 'important.com', 'domain', 'domain', 'allowlist', 1, 1)`)
	require.NoError(t, err)

	entriesRepo := idb.NewEntriesRepo(database)
	consolidatedRepo := idb.NewConsolidatedRepo(database)

	genericSourceTypes := []string{"domain"}
	allowlistEntriesByType := make(map[string]u.StringSet)

	resolvedBlockByType := map[string]u.StringSet{
		"domain": u.NewStringSet([]string{"important.com"}),
	}

	processAllowlists(
		ctx,
		logger,
		entriesRepo,
		consolidatedRepo,
		genericSourceTypes,
		resolvedBlockByType,
		allowlistEntriesByType,
	)

	// must_consider entries should be kept even if in resolved blocklist
	domainEntries := allowlistEntriesByType["domain"]
	assert.Equal(t, 1, domainEntries.Size())
	assert.True(t, domainEntries.Contains("important.com"))
}
