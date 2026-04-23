// Package db provides SQLite database access for dns-toolkit.
package db

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite" // Pure-Go SQLite driver
)

type DB struct {
	conn            *sql.DB
	path            string
	schemaRecreated bool
}

// Open creates or opens a SQLite database at the given path.
// If forceRecreate is true, the schema is dropped and rebuilt unconditionally.
// The schema is also rebuilt automatically whenever schema.sql changes.
func Open(dbPath string, forceRecreate ...bool) (*DB, error) {
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, fmt.Errorf("creating database directory: %w", err)
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

	force := len(forceRecreate) > 0 && forceRecreate[0]
	recreated, err := db.EnsureSchema(force)
	if err != nil {
		return nil, closeOnError(conn, "ensuring schema", err)
	}
	db.schemaRecreated = recreated

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
func (db *DB) InTransaction(fn func(tx *sql.Tx) error) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return fmt.Errorf("beginning transaction: %w", err)
	}

	if err := fn(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("rollback failed: %w (original error: %v)", rbErr, err)
		}
		return err
	}

	return tx.Commit()
}

// closeOnError closes c and returns a combined error wrapping both the original
// operation error and any error from closing.
func closeOnError(c io.Closer, op string, opErr error) error {
	if closeErr := c.Close(); closeErr != nil {
		return fmt.Errorf("%s: %w; closing connection: %v", op, opErr, closeErr)
	}
	return fmt.Errorf("%s: %w", op, opErr)
}
