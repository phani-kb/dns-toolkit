// Package db provides SQLite database access for dns-toolkit.
package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/phani-kb/multilog"
	_ "modernc.org/sqlite"
)

type DB struct {
	conn            *sql.DB
	path            string
	schemaRecreated bool
}

// Open creates or opens a SQLite database at the given path.
// If forceRecreate is true, the schema is dropped and rebuilt unconditionally.
// The schema is also rebuilt automatically whenever schema.sql changes.
func Open(ctx context.Context, logger *multilog.Logger, dbPath string, forceRecreate bool) (*DB, error) {
	if forceRecreate {
		if err := removeDBFiles(dbPath); err != nil {
			return nil, fmt.Errorf("resetting database files: %w", err)
		}
	}

	db, err := openConn(dbPath)
	if err != nil {
		return nil, err
	}

	recreated, err := db.EnsureSchema(ctx, logger, forceRecreate)
	if err != nil {
		return nil, closeOnError(db, "ensuring schema", err)
	}
	db.schemaRecreated = recreated

	return db, nil
}

// OpenInspect opens the database for read-only inspection without running
// EnsureSchema, so it never modifies or recreates the schema.
func OpenInspect(dbPath string) (*DB, error) {
	return openConn(dbPath)
}

func openConn(dbPath string) (*DB, error) {
	if err := os.MkdirAll(filepath.Dir(dbPath), 0o755); err != nil {
		return nil, fmt.Errorf("creating db dir: %w", err)
	}

	conn, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}

	if err = conn.Ping(); err != nil {
		return nil, closeOnError(conn, "pinging database", err)
	}

	db := &DB{conn: conn, path: dbPath}

	if err = db.applyPragmas(); err != nil {
		return nil, closeOnError(conn, "applying pragmas", err)
	}

	return db, nil
}

// Close closes the database connection.
func (db *DB) Close() error {
	if db.conn != nil {
		return db.conn.Close()
	}
	return nil
}

// Conn returns the underlying sql.DB connection for advanced queries.
func (db *DB) Conn() *sql.DB {
	return db.conn
}

// Path returns the database file path.
func (db *DB) Path() string {
	return db.path
}

// applyPragmas sets SQLite performance tuning parameters.
func (db *DB) applyPragmas() error {
	pragmas := []string{
		"PRAGMA journal_mode=WAL",
		"PRAGMA synchronous=NORMAL",
		"PRAGMA cache_size=-64000",
		"PRAGMA foreign_keys=ON",
		"PRAGMA busy_timeout=5000",
		"PRAGMA temp_store=MEMORY",
	}
	for _, pragma := range pragmas {
		if _, err := db.conn.Exec(pragma); err != nil {
			return fmt.Errorf("executing %s: %w", pragma, err)
		}
	}
	return nil
}

// Vacuum reclaims unused space in the database file.
func (db *DB) Vacuum() error {
	_, err := db.conn.Exec("VACUUM")
	return err
}

// InTransaction executes fn within a database transaction.
func (db *DB) InTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := db.conn.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("beginning transaction: %w", err)
	}

	if err := fn(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("rollback failed: %w", errors.Join(err, rbErr))
		}
		return err
	}

	return tx.Commit()
}

// CloseLogError closes the DB and logs any error.
func (db *DB) CloseLogError(logger *multilog.Logger) {
	if db.conn != nil {
		if err := db.conn.Close(); err != nil {
			logger.Warnf("Error closing database: %v", err)
		}
	}
}
