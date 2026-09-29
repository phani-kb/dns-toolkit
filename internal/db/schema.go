package db

import (
	"bufio"
	"context"
	"crypto/sha256"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/multilog"
)

//go:embed schema.sql
var schemaSQL string

var (
	schemaChecksumOnce         sync.Once
	schemaChecksumValue        string
	schemaObjectsChecksumOnce  sync.Once
	schemaObjectsChecksumValue string
)

// SchemaChecksum returns the SHA-256 hash of the embedded schema.sql.
func SchemaChecksum() string {
	schemaChecksumOnce.Do(func() {
		schemaChecksumValue = fmt.Sprintf("%x", sha256.Sum256([]byte(schemaSQL)))
	})
	return schemaChecksumValue
}

// EmbeddedSchemaObjectsChecksum returns a checksum of the embedded schema objects
func EmbeddedSchemaObjectsChecksum() string {
	schemaObjectsChecksumOnce.Do(func() {
		keys := schemaTableKeysFromConstants()
		keys = append(keys, extractSchemaIndexKeys(schemaSQL)...)
		schemaObjectsChecksumValue = checksumForKeys(keys)
	})
	return schemaObjectsChecksumValue
}

// LiveSchemaObjectsChecksum returns a checksum of the live schema objects
func (db *DB) LiveSchemaObjectsChecksum() (string, error) {
	var objects []struct {
		Type string `db:"type"`
		Name string `db:"name"`
	}
	if err := db.selectRead(context.Background(), &objects, `
		select type, name
		from sqlite_master
		where type in ('table', 'view', 'index', 'trigger')
			and sql is not null
			and name not glob 'sqlite_*'
			and (
				(type in ('table', 'view') and (name glob ? or name glob ?))
				or (type = 'index' and (tbl_name glob ? or tbl_name glob ?))
				or (type = 'trigger' and (tbl_name glob ? or tbl_name glob ?))
			)
			and not (
				(type = 'table' and name = ?)
				or (type = 'index' and tbl_name = ?)
			)
		order by type, name`,
		constants.TablePrefix+"*",
		"_"+constants.TablePrefix+"*",
		constants.TablePrefix+"*",
		"_"+constants.TablePrefix+"*",
		constants.TablePrefix+"*",
		"_"+constants.TablePrefix+"*",
		constants.SchemaMetadataTable,
		constants.SchemaMetadataTable,
	); err != nil {
		return "", fmt.Errorf("querying live schema objects: %w", err)
	}

	keys := make([]string, 0, len(objects))
	for _, obj := range objects {
		keys = append(keys, strings.ToLower(obj.Type)+":"+strings.ToLower(obj.Name))
	}

	return checksumForKeys(keys), nil
}

func schemaTableNamesFromConstants() []string {
	return []string{
		constants.TableSources,
		constants.TableTypeNames,
		constants.TableListTypeNames,
		constants.TableGroupNames,
		constants.TableCategoryNames,
		constants.TableSourceTypes,
		constants.TableSourceListTypes,
		constants.TableSourceListTypeNotes,
		constants.TableSourceListTypeGroups,
		constants.TableSourceCategories,
		constants.TableSourceCountries,
		constants.TableSourceContent,
		constants.TableSourceFiles,
		constants.TableDownloads,
		constants.TableEntries,
		constants.TableEntryGroups,
		constants.TableEntryCategories,
		constants.TableSourceEntryCounts,
		constants.TableConsolidatedGeneral,
		constants.TableConsolidatedGroup,
		constants.TableConsolidatedCategory,
		constants.TableOverlapResults,
		constants.TableTopEntries,
		constants.TableResolvedAllow,
		constants.TableScopedAllow,
	}
}

func schemaTableKeysFromConstants() []string {
	tables := schemaTableNamesFromConstants()
	keys := make([]string, 0, len(tables))
	for _, t := range tables {
		keys = append(keys, "table:"+strings.ToLower(t))
	}
	return keys
}

func extractSchemaIndexKeys(sqlText string) []string {
	scanner := bufio.NewScanner(strings.NewReader(sqlText))
	seen := make(map[string]struct{})
	keys := make([]string, 0)

	for scanner.Scan() {
		line := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if line == "" || strings.HasPrefix(line, "--") {
			continue
		}

		for _, tok := range strings.Fields(line) {
			clean := strings.Trim(tok, "(),;\t\n\r\"'`")
			if !strings.HasPrefix(clean, "idx_") {
				continue
			}
			if _, ok := seen[clean]; ok {
				continue
			}
			seen[clean] = struct{}{}
			keys = append(keys, "index:"+clean)
		}
	}

	return keys
}

func checksumForKeys(keys []string) string {
	if len(keys) == 0 {
		return fmt.Sprintf("%x", sha256.Sum256(nil))
	}

	sorted := append([]string(nil), keys...)
	sort.Strings(sorted)

	return fmt.Sprintf("%x", sha256.Sum256([]byte(strings.Join(sorted, "\n"))))
}

// EnsureSchema checks if the DB schema matches the embedded schema.sql.
// If the schema has changed (or forceRecreate is true), it drops all tables
// and recreates them from schema.sql.
// Returns true if the schema was recreated.
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

	if err := db.recreateSchema(ctx, currentChecksum); err != nil {
		return false, fmt.Errorf("recreating schema: %w", err)
	}
	if logger != nil {
		logger.Infof("Recreated database schema from schema.sql (checksum %s)", currentChecksum)
	}
	return true, nil
}

func (db *DB) storeSchemaChecksum(checksum string) error {
	q := fmt.Sprintf(`
	  create table if not exists %s (
	   checksum text not null,
	   applied_at text not null default (datetime('now'))
	  );
	  delete from %s;
	  insert into %s (checksum) values (?);
	 `, constants.SchemaMetadataTable, constants.SchemaMetadataTable, constants.SchemaMetadataTable)

	if _, err := db.writeConn.Exec(q, checksum); err != nil {
		return fmt.Errorf("storing schema checksum: %w", err)
	}
	return nil
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
	q := fmt.Sprintf("select checksum from %s limit 1", constants.SchemaMetadataTable)
	if err := db.readConn.QueryRow(q).Scan(&checksum); err != nil {
		return "", err
	}
	return checksum, nil
}

func (db *DB) recreateSchema(ctx context.Context, checksum string) error {
	// to avoid cascade issues
	if _, err := db.writeConn.Exec("pragma foreign_keys=off"); err != nil {
		return fmt.Errorf("disabling foreign keys: %w", err)
	}

	tables, err := db.listUserTables()
	if err != nil {
		return fmt.Errorf("listing tables: %w", err)
	}
	views, err := db.listUserViews()
	if err != nil {
		return fmt.Errorf("listing views: %w", err)
	}

	tx, err := db.writeConn.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("beginning drop transaction: %w", err)
	}

	for _, v := range views {
		q := "drop view if exists " + v
		if _, err := tx.Exec(q); err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return fmt.Errorf("dropping view %s: %w; rollback failed: %v", v, err, rbErr)
			}
			return fmt.Errorf("dropping view %s: %w", v, err)
		}
	}
	for _, t := range tables {
		q := "drop table if exists " + t
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

	if _, err := db.writeConn.Exec("pragma foreign_keys=on"); err != nil {
		return fmt.Errorf("re-enabling foreign keys: %w", err)
	}

	if _, err := db.writeConn.Exec(schemaSQL); err != nil {
		return fmt.Errorf("executing schema.sql: %w", err)
	}

	if err := db.storeSchemaChecksum(checksum); err != nil {
		return err
	}

	return nil
}

func (db *DB) listUserTables() ([]string, error) {
	return db.listUserObjects("table")
}

func (db *DB) listUserViews() ([]string, error) {
	return db.listUserObjects("view")
}

func (db *DB) listUserObjects(objType string) ([]string, error) {
	q := fmt.Sprintf("select name from sqlite_master where type='%s' "+
		"AND (name like '%s%%' OR name like '_%s%%') ORDER BY name", objType, constants.TablePrefix, constants.TablePrefix)
	var names []string
	if err := db.readConn.SelectContext(context.Background(), &names, q); err != nil {
		return nil, err
	}
	return names, nil
}

func isValidTableName(name string) bool {
	for _, t := range schemaTableNamesFromConstants() {
		if name == t {
			return true
		}
	}
	return false
}

func splitSchemaStatements(schema string) []string {
	var stmts []string
	var sb strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(schema))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(strings.TrimSpace(line), "--") {
			continue
		}
		sb.WriteString(line)
		sb.WriteString("\n")
		if strings.Contains(line, ";") {
			stmt := strings.TrimSpace(sb.String())
			if stmt != "" {
				stmts = append(stmts, stmt)
			}
			sb.Reset()
		}
	}
	if rest := strings.TrimSpace(sb.String()); rest != "" {
		stmts = append(stmts, rest)
	}
	return stmts
}

func statementsForTable(schema, table string) ([]string, error) {
	if !isValidTableName(table) {
		return nil, fmt.Errorf("refusing to recreate unknown table %q", table)
	}
	lowerTable := strings.ToLower(table)
	var out []string
	for _, stmt := range splitSchemaStatements(schema) {
		lowered := strings.ToLower(strings.Join(strings.Fields(stmt), " "))
		if (strings.Contains(lowered, "create table") || strings.Contains(lowered, "create virtual table")) &&
			strings.Contains(lowered, lowerTable) {
			out = append(out, stmt)
			continue
		}
		if strings.Contains(lowered, "create view") && strings.Contains(lowered, lowerTable) {
			out = append(out, stmt)
			continue
		}
		if strings.Contains(lowered, "create index") && strings.Contains(lowered, " on "+lowerTable+" ") {
			out = append(out, stmt)
		}
	}
	return out, nil
}

func (db *DB) SchemaRecreated() bool {
	return db.schemaRecreated
}

func (db *DB) StoredChecksum(logger *multilog.Logger) string {
	cs, err := db.storedSchemaChecksum()
	if err != nil {
		logger.Warnf("Error getting schema checksum: %v", err)
		return ""
	}
	return cs
}

func (db *DB) TableRowCounts() (map[string]int64, error) {
	tables, err := db.listUserTables()
	if err != nil {
		return nil, err
	}

	counts := make(map[string]int64, len(tables))
	for _, t := range tables {
		var count int64
		q := "select count(*) from " + t
		err := db.readConn.QueryRow(q).Scan(&count)
		if err != nil {
			counts[t] = -1
			continue
		}
		counts[t] = count
	}
	return counts, nil
}
