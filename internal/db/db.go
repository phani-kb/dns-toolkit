// Package db provides SQLite database access for dns-toolkit.
package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/phani-kb/multilog"
	_ "modernc.org/sqlite"
)

const dsnPragmas = "_pragma=journal_mode(WAL)" +
	"&_pragma=synchronous(NORMAL)" +
	"&_pragma=busy_timeout(5000)" +
	"&_pragma=cache_size(-64000)" +
	"&_pragma=foreign_keys(ON)" +
	"&_pragma=temp_store(MEMORY)"

type DB struct {
	readConn        *sql.DB
	writeConn       *sql.DB
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

func buildDataSource(dbPath string) string {
	sep := "?"
	if strings.Contains(dbPath, "?") {
		sep = "&"
	}
	return dbPath + sep + dsnPragmas
}

func openConn(dbPath string) (*DB, error) {
	if err := os.MkdirAll(filepath.Dir(dbPath), 0o755); err != nil {
		return nil, fmt.Errorf("creating db dir: %w", err)
	}

	dsn := buildDataSource(dbPath)

	writeConn, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("opening write connection: %w", err)
	}
	writeConn.SetMaxOpenConns(1)
	writeConn.SetMaxIdleConns(1)
	writeConn.SetConnMaxLifetime(0)

	if err = writeConn.Ping(); err != nil {
		return nil, closeOnError(writeConn, "pinging write database", err)
	}

	readConn, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, closeOnError(writeConn, "opening read connection", err)
	}
	readMax := max(3, runtime.GOMAXPROCS(0))
	readConn.SetMaxOpenConns(readMax)
	readConn.SetMaxIdleConns(readMax)
	readConn.SetConnMaxLifetime(0)

	if err = readConn.Ping(); err != nil {
		_ = writeConn.Close() // nolint: errcheck
		return nil, closeOnError(readConn, "pinging read database", err)
	}

	return &DB{readConn: readConn, writeConn: writeConn, path: dbPath}, nil
}

// Close closes the database connection.
func (db *DB) Close() error {
	var errs []error
	if db.readConn != nil {
		if err := db.readConn.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	if db.writeConn != nil {
		if err := db.writeConn.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

// ReadConn returns the multi-connection read pool.
func (db *DB) ReadConn() *sql.DB {
	return db.readConn
}

// WriteConn returns the single-connection write pool.
func (db *DB) WriteConn() *sql.DB {
	return db.writeConn
}

// Path returns the database file path.
func (db *DB) Path() string {
	return db.path
}

// Vacuum reclaims unused space in the database file.
func (db *DB) Vacuum() error {
	_, err := db.writeConn.Exec("vacuum")
	return err
}

// InTransaction executes fn within a write transaction on the single-writer pool.
func (db *DB) InTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := db.writeConn.BeginTx(ctx, nil)
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

// InReadTransaction executes fn within a read-only transaction on the read pool.
func (db *DB) InReadTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := db.readConn.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return fmt.Errorf("beginning read transaction: %w", err)
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback() // nolint: errcheck
		return err
	}

	return tx.Commit()
}

// CloseLogError closes the DB and logs any error.
func (db *DB) CloseLogError(logger *multilog.Logger) {
	if db.Close() != nil {
		if err := db.Close(); err != nil {
			logger.Warnf("Error closing database: %v", err)
		}
	}
}
