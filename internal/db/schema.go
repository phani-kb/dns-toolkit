package db

import (
	"context"
	"crypto/sha256"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/multilog"
)

//go:embed schema.sql
var schemaSQL string

// SchemaChecksum returns the SHA-256 hash of the embedded schema.sql.
func SchemaChecksum() string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(schemaSQL)))
}

// EnsureSchema checks if the DB schema matches the embedded schema.sql.
// If the schema has changed (or forceRecreate is true), it drops all tables
// and recreates them from schema.sql. Returns true if the schema was recreated.
func (db *DB) EnsureSchema(ctx context.Context, logger *multilog.Logger, forceRecreate bool) (bool, error) {
	currentChecksum := SchemaChecksum()

	if !forceRecreate {
		storedChecksum, err := db.storedSchemaChecksum()
		switch {
		case err == nil && storedChecksum == currentChecksum:
			return false, nil
		case err == nil:
			// checksum mismatch, recreate
		case errors.Is(err, sql.ErrNoRows):
			// metadata missing, recreate
		default:
			return false, fmt.Errorf("reading stored schema checksum: %w", err)
		}
	}

	if err := db.recreateSchema(ctx, logger, currentChecksum); err != nil {
		return false, fmt.Errorf("recreating schema: %w", err)
	}
	return true, nil
}

func (db *DB) storedSchemaChecksum() (string, error) {
	exists, err := db.tableExists(constants.SchemaMetadataTable)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", sql.ErrNoRows
	}

	var checksum string
	q := fmt.Sprintf("SELECT checksum FROM %s LIMIT 1", constants.SchemaMetadataTable)
	if err := db.conn.QueryRow(q).Scan(&checksum); err != nil {
		return "", err
	}
	return checksum, nil
}

// recreateSchema drops all user tables and rebuilds from schema.sql.
func (db *DB) recreateSchema(ctx context.Context, logger *multilog.Logger, checksum string) error {
	// to avoid cascade issues
	if _, err := db.conn.Exec("pragma foreign_keys=off"); err != nil {
		return fmt.Errorf("disabling foreign keys: %w", err)
	}

	tables, err := db.listUserTables(logger)
	if err != nil {
		return fmt.Errorf("listing tables: %w", err)
	}

	tx, err := db.conn.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("beginning drop transaction: %w", err)
	}

	for _, t := range tables {
		q := "DROP TABLE IF EXISTS " + t
		if _, err := tx.Exec(q); err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return fmt.Errorf("dropping table %s: %w; rollback failed: %v", t, err, rbErr)
			}
			return fmt.Errorf("dropping table %s: %w", t, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing drops: %w", err)
	}

	if _, err := db.conn.Exec("pragma foreign_keys=on"); err != nil {
		return fmt.Errorf("re-enabling foreign keys: %w", err)
	}

	if _, err := db.conn.Exec(schemaSQL); err != nil {
		return fmt.Errorf("executing schema.sql: %w", err)
	}

	q := fmt.Sprintf(`
	  CREATE TABLE IF NOT EXISTS %s (
	   checksum TEXT NOT NULL,
	   applied_at TEXT NOT NULL DEFAULT (datetime('now'))
	  );
	  DELETE FROM %s;
	  INSERT INTO %s (checksum) VALUES (?);
	 `, constants.SchemaMetadataTable, constants.SchemaMetadataTable, constants.SchemaMetadataTable)

	if _, err := db.conn.Exec(q, checksum); err != nil {
		return fmt.Errorf("storing schema checksum: %w", err)
	}

	return nil
}

// listUserTables returns all non-internal table names in the database.
func (db *DB) listUserTables(logger *multilog.Logger) ([]string, error) {
	q := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='table' "+
		"AND (name like '%s%%' OR name like '_%s%%') ORDER BY name", constants.TablePrefix, constants.TablePrefix)
	rows, err := db.conn.Query(q)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			logger.Warnf("Error closing rows: %v", err)
		}
	}(rows)

	var tables []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		tables = append(tables, name)
	}
	return tables, rows.Err()
}

// SchemaRecreated returns whether the schema was rebuilt on the last call.
func (db *DB) SchemaRecreated() bool {
	return db.schemaRecreated
}

// StoredChecksum returns the schema checksum stored in the database, or empty if none.
func (db *DB) StoredChecksum(logger *multilog.Logger) string {
	cs, err := db.storedSchemaChecksum()
	if err != nil {
		logger.Warnf("Error getting schema checksum: %v", err)
		return ""
	}
	return cs
}

// TableRowCounts returns a map of table name to row count for all user tables.
func (db *DB) TableRowCounts(logger *multilog.Logger) (map[string]int64, error) {
	tables, err := db.listUserTables(logger)
	if err != nil {
		return nil, err
	}

	counts := make(map[string]int64, len(tables))
	for _, t := range tables {
		var count int64
		q := "SELECT COUNT(*) FROM " + t
		err := db.conn.QueryRow(q).Scan(&count)
		if err != nil {
			counts[t] = -1
			continue
		}
		counts[t] = count
	}
	return counts, nil
}
