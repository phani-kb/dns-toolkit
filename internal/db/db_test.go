package db

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"testing"

	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpen(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_open.db")

	database, err := Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	require.NotNil(t, database)
	defer database.Close() // nolint: errcheck

	_, statErr := os.Stat(dbPath)
	assert.NoError(t, statErr)
}

func TestOpen_ForceRecreate(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_recreate.db")

	db1, err := Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	db1.Close() // nolint: errcheck

	db2, err := Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	require.NotNil(t, db2)
	defer db2.Close() // nolint: errcheck
}

func TestOpen_ExistingDB(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_existing.db")

	db1, err := Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	db1.Close() // nolint: errcheck

	db2, err := Open(ctx, logger, dbPath, false)
	require.NoError(t, err)
	require.NotNil(t, db2)
	defer db2.Close() // nolint: errcheck
}

func TestClose(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_close.db")

	database, err := Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	require.NotNil(t, database)

	err = database.Close()
	assert.NoError(t, err)

	database.conn = nil
	err = database.Close()
	assert.NoError(t, err)
}

func TestConn(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_conn.db")

	database, err := Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	defer database.Close() // nolint: errcheck

	conn := database.Conn()
	assert.NotNil(t, conn)

	err = conn.Ping()
	assert.NoError(t, err)
}

func TestPath(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_path.db")

	database, err := Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	defer database.Close() // nolint: errcheck

	assert.Equal(t, dbPath, database.Path())
}

func TestVacuum(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_vacuum.db")

	database, err := Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	defer database.Close() // nolint: errcheck

	_, err = database.conn.Exec("CREATE TABLE IF NOT EXISTS test_table (id INTEGER PRIMARY KEY, data TEXT)")
	require.NoError(t, err)

	for i := 0; i < 100; i++ {
		_, err = database.conn.Exec("INSERT INTO test_table (data) VALUES (?)", "test data")
		require.NoError(t, err)
	}

	_, err = database.conn.Exec("DELETE FROM test_table")
	require.NoError(t, err)

	err = database.Vacuum()
	assert.NoError(t, err)
}

func TestOpenInspect(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_inspect.db")

	db1, err := Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	db1.Close() // nolint: errcheck

	db2, err := OpenInspect(dbPath)
	require.NoError(t, err)
	require.NotNil(t, db2)
	defer db2.Close() // nolint: errcheck

	assert.NotNil(t, db2.Conn())
}

func TestInTransaction(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_transaction.db")

	database, err := Open(ctx, logger, dbPath, true)
	require.NoError(t, err)
	defer database.Close() // nolint: errcheck

	_, err = database.conn.Exec("CREATE TABLE IF NOT EXISTS tx_test (id INTEGER PRIMARY KEY, val TEXT)")
	require.NoError(t, err)

	err = database.InTransaction(ctx, func(tx *sql.Tx) error {
		_, err := tx.Exec("INSERT INTO tx_test (val) VALUES (?)", "test")
		return err
	})
	assert.NoError(t, err)

	var count int
	err = database.conn.QueryRow("SELECT COUNT(*) FROM tx_test").Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, 1, count)

	err = database.InTransaction(ctx, func(tx *sql.Tx) error {
		_, err := tx.Exec("INSERT INTO tx_test (val) VALUES (?)", "will rollback")
		if err != nil {
			return err
		}
		return assert.AnError // trigger rollback
	})
	assert.Error(t, err)

	err = database.conn.QueryRow("SELECT COUNT(*) FROM tx_test").Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, 1, count)
}

func TestCloseLogError(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ctx := context.Background()
	dbPath := filepath.Join(t.TempDir(), "test_close_log.db")

	database, err := Open(ctx, logger, dbPath, true)
	require.NoError(t, err)

	database.CloseLogError(logger)

	// 2nd call
	database.CloseLogError(logger)

	database.conn = nil
	database.CloseLogError(logger)
}
