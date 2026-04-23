package db

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"
)

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
	const q = "SELECT 1 FROM sqlite_master WHERE type='table' AND name = ? LIMIT 1"

	var exists int
	err := db.conn.QueryRow(q, tableName).Scan(&exists)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
