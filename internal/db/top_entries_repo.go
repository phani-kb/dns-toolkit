package db

import (
	"context"
	"database/sql"
	"fmt"
	"math"

	"github.com/phani-kb/dns-toolkit/internal/constants"
)

type TopEntriesRepo struct {
	db *DB
}

func NewTopEntriesRepo(db *DB) *TopEntriesRepo {
	return &TopEntriesRepo{db: db}
}

type TopEntryRow struct {
	Entry       string
	SourceCount int
}

// GetTopEntries returns entries that appear in at least minSources distinct sources
// for the given generic source type and list type.
func (r *TopEntriesRepo) GetTopEntries(
	_ context.Context,
	genericSourceType string,
	listType string,
	minSources int,
	maxEntries int,
) ([]TopEntryRow, error) {
	query := `
		SELECT e.entry, COUNT(DISTINCT e.source_id) AS source_count
		FROM ` + constants.TableEntries + ` e
		JOIN ` + constants.TableSources + ` s ON s.id = e.source_id
		WHERE e.generic_source_type = ?
			AND e.list_type = ?
			AND e.valid = 1
			AND s.disabled = 0
		GROUP BY e.entry
		HAVING source_count >= ?
		ORDER BY source_count DESC, e.entry ASC
	`

	args := []any{genericSourceType, listType, minSources}

	if maxEntries > 0 && maxEntries < math.MaxInt {
		query += " LIMIT ?"
		args = append(args, maxEntries)
	}

	rows, err := r.db.readConn.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("querying top entries: %w", err)
	}
	defer rows.Close() // nolint: errcheck

	var results []TopEntryRow
	for rows.Next() {
		var row TopEntryRow
		if err := rows.Scan(&row.Entry, &row.SourceCount); err != nil {
			return nil, fmt.Errorf("scanning top entry row: %w", err)
		}
		results = append(results, row)
	}
	return results, rows.Err()
}

type TopEntryGroup struct {
	GenericSourceType string
	ListType          string
	MinSources        int
	Count             int
}

// ListTopEntryGroups returns all distinct groups with counts.
func (r *TopEntriesRepo) ListTopEntryGroups(_ context.Context) ([]TopEntryGroup, error) {
	query := `
		SELECT generic_source_type, list_type, min_sources, COUNT(*) AS count
		FROM ` + constants.TableTopEntries + `
		GROUP BY generic_source_type, list_type, min_sources
		ORDER BY generic_source_type, list_type, min_sources
	`

	rows, err := r.db.readConn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("querying top entry groups: %w", err)
	}
	defer rows.Close() // nolint: errcheck

	var results []TopEntryGroup
	for rows.Next() {
		var g TopEntryGroup
		if err := rows.Scan(&g.GenericSourceType, &g.ListType, &g.MinSources, &g.Count); err != nil {
			return nil, fmt.Errorf("scanning top entry group: %w", err)
		}
		results = append(results, g)
	}
	return results, rows.Err()
}

// GetTopEntriesList returns just the entry strings for a given group.
func (r *TopEntriesRepo) GetTopEntriesList(
	_ context.Context,
	genericSourceType, listType string,
	minSources int,
) ([]string, error) {
	query := `
		SELECT entry FROM ` + constants.TableTopEntries + `
		WHERE generic_source_type = ? AND list_type = ? AND min_sources = ?
		ORDER BY source_count DESC, entry ASC
	`

	rows, err := r.db.readConn.Query(query, genericSourceType, listType, minSources)
	if err != nil {
		return nil, fmt.Errorf("querying top entries list: %w", err)
	}
	defer rows.Close() // nolint: errcheck

	var entries []string
	for rows.Next() {
		var entry string
		if err := rows.Scan(&entry); err != nil {
			return nil, fmt.Errorf("scanning top entry: %w", err)
		}
		entries = append(entries, entry)
	}
	return entries, rows.Err()
}

// PersistTopEntries stores the computed top entries in dnstk_top_entries.
func (r *TopEntriesRepo) PersistTopEntries(
	ctx context.Context,
	genericSourceType string,
	listType string,
	minSources int,
	entries []TopEntryRow,
) error {
	return r.db.InTransaction(ctx, func(tx *sql.Tx) error {
		_, err := tx.Exec(
			"DELETE FROM "+constants.TableTopEntries+" WHERE generic_source_type = ? AND list_type = ? AND min_sources = ?",
			genericSourceType,
			listType,
			minSources,
		)
		if err != nil {
			return fmt.Errorf("clearing top entries: %w", err)
		}

		if len(entries) == 0 {
			return nil
		}

		stmt, err := tx.Prepare(`
			INSERT OR IGNORE INTO ` + constants.TableTopEntries + `
			(entry, generic_source_type, list_type, source_count, min_sources)
			VALUES (?, ?, ?, ?, ?)`)
		if err != nil {
			return fmt.Errorf("preparing top entries insert: %w", err)
		}
		defer func() { _ = stmt.Close() }() // nolint: errcheck

		for _, row := range entries {
			if _, err := stmt.Exec(row.Entry, genericSourceType, listType, row.SourceCount, minSources); err != nil {
				return fmt.Errorf("inserting top entry: %w", err)
			}
		}
		return nil
	})
}

// ClearAllTopEntries removes all rows from the top entries table.
func (r *TopEntriesRepo) ClearAllTopEntries() error {
	_, err := r.db.writeConn.Exec("DELETE FROM " + constants.TableTopEntries)
	return err
}
