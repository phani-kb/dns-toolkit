package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
)

const metadataChunkSize = 400

func removeDBFiles(dbPath string) error {
	paths := []string{dbPath, dbPath + "-wal", dbPath + "-shm"}
	for _, path := range paths {
		if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("removing %s: %w", path, err)
		}
	}
	return nil
}

func closeOnError(c io.Closer, op string, opErr error) error {
	if closeErr := c.Close(); closeErr != nil {
		return fmt.Errorf("%s: %w", op, errors.Join(opErr, closeErr))
	}
	return fmt.Errorf("%s: %w", op, opErr)
}

func (db *DB) tableExists(tableName string) (bool, error) {
	const q = "select 1 from sqlite_master where type='table' and name = ? limit 1"

	var exists int
	err := db.readConn.QueryRow(q, tableName).Scan(&exists)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// boolToInt converts a boolean to 0 or 1 for SQLite storage.
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func placeholders(n int) string {
	if n <= 0 {
		return ""
	}
	return strings.TrimRight(strings.Repeat("?,", n), ",")
}

func int64Args(ids []int64) []any {
	args := make([]any, len(ids))
	for i, id := range ids {
		args[i] = id
	}
	return args
}

// selectChunked runs query in chunk overs ids
func selectChunked[T any](
	ctx context.Context,
	db *sqlx.DB,
	ids []int64,
	query func(n int) string,
	errLabel string,
) ([]T, error) {
	var all []T
	for i := 0; i < len(ids); i += metadataChunkSize {
		chunk := ids[i:min(i+metadataChunkSize, len(ids))]
		var rows []T
		if err := db.SelectContext(ctx, &rows, query(len(chunk)), int64Args(chunk)...); err != nil {
			return nil, fmt.Errorf("%s: %w", errLabel, err)
		}
		all = append(all, rows...)
	}
	return all, nil
}
