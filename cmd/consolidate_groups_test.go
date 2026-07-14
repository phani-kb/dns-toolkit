package cmd

import (
	"context"
	"path/filepath"
	"sync"
	"testing"

	con "github.com/phani-kb/dns-toolkit/internal/consolidators"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	idb "github.com/phani-kb/dns-toolkit/internal/db"
	"github.com/phani-kb/multilog"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProcessGroupConsolidationFromDB_EmptyDatabase(t *testing.T) {
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

	processGroupConsolidation(
		ctx,
		logger,
		constants.GroupMini,
		[]string{constants.SourceTypeDomain},
		entriesRepo,
		consolidatedRepo,
		&persistMu,
	)

	count, err := consolidatedRepo.GetConsolidatedCount(
		constants.SourceTypeDomain,
		constants.ListTypeBlocklist,
		"group",
		true,
	)
	require.NoError(t, err)
	assert.Equal(t, int64(0), count)
}
