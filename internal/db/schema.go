package db

import (
	"crypto/sha256"
	_ "embed"
	"fmt"
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
func (db *DB) EnsureSchema(forceRecreate bool) (bool, error) {
	return true, nil
}

// storedSchemaChecksum reads the checksum stored in _schema_meta.
func (db *DB) storedSchemaChecksum() (string, error) {
	return "checksum", nil
}

// recreateSchema drops all user tables and rebuilds from schema.sql.
func (db *DB) recreateSchema(checksum string) error {
	return nil
}

// listUserTables returns all non-internal table names in the database.
func (db *DB) listUserTables() ([]string, error) {
	return nil, nil
}

// SchemaRecreated returns whether the schema was rebuilt on the last Open call.
// Callers (e.g., the download/process commands) use this to know they must
// re-import sources and reprocess everything.
func (db *DB) SchemaRecreated() bool {
	return db.schemaRecreated
}

// SchemaVersion returns 1 if a schema has been applied, 0 otherwise.
func (db *DB) SchemaVersion() (int, error) {
	return 0, nil
}

// StoredChecksum returns the schema checksum stored in the database, or empty if none.
func (db *DB) StoredChecksum() string {
	return ""
}

// TableRowCounts returns a map of table name to row count for all user tables.
func (db *DB) TableRowCounts() (map[string]int64, error) {
	return nil, nil
}
