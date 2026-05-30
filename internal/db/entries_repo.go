package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/phani-kb/dns-toolkit/internal/constants"
)

type EntriesRepo struct {
	db *DB
}

func NewEntriesRepo(db *DB) *EntriesRepo {
	return &EntriesRepo{db: db}
}

type EntryRow struct {
	Entry             string
	GenericSourceType string
	ActualSourceType  string
	ListType          string
	Valid             bool
	MustConsider      bool
}

type EntryGroupRow struct {
	SourceType string
	ListType   string
	GroupName  string
}

type EntryCategoryRow struct {
	SourceType string
	ListType   string
	Category   string
}

// ReplaceSourceData rewrites all processed-entry rows for a source in one transaction.
func (r *EntriesRepo) ReplaceSourceData(
	ctx context.Context,
	sourceID int64,
	entries []EntryRow,
	groups []EntryGroupRow,
	categories []EntryCategoryRow,
) error {
	if sourceID <= 0 {
		return fmt.Errorf("invalid source id: %d", sourceID)
	}

	return r.db.InTransaction(ctx, func(tx *sql.Tx) error {
		if _, err := tx.Exec("DELETE FROM "+constants.TableEntryCategories+" WHERE source_id = ?", sourceID); err != nil {
			return fmt.Errorf("clearing entry categories for source %d: %w", sourceID, err)
		}
		if _, err := tx.Exec("DELETE FROM "+constants.TableEntryGroups+" WHERE source_id = ?", sourceID); err != nil {
			return fmt.Errorf("clearing entry groups for source %d: %w", sourceID, err)
		}
		if _, err := tx.Exec("DELETE FROM "+constants.TableEntries+" WHERE source_id = ?", sourceID); err != nil {
			return fmt.Errorf("clearing entries for source %d: %w", sourceID, err)
		}

		entriesStmt, err := tx.Prepare(`
			INSERT OR IGNORE INTO ` + constants.TableEntries + `
			(source_id, entry, generic_source_type, actual_source_type, list_type, valid, must_consider)
			VALUES (?, ?, ?, ?, ?, ?, ?)`)
		if err != nil {
			return fmt.Errorf("preparing entries insert: %w", err)
		}
		defer func() { _ = entriesStmt.Close() }() // nolint: errcheck

		for _, row := range dedupeEntryRows(entries) {
			if row.Entry == "" || row.GenericSourceType == "" || row.ActualSourceType == "" || row.ListType == "" {
				continue
			}
			if _, rowErr := entriesStmt.Exec(
				sourceID,
				row.Entry,
				row.GenericSourceType,
				row.ActualSourceType,
				row.ListType,
				boolToInt(row.Valid),
				boolToInt(row.MustConsider),
			); rowErr != nil {
				return fmt.Errorf("inserting entry for source %d: %w", sourceID, rowErr)
			}
		}

		groupsStmt, err := tx.Prepare(`
			INSERT OR IGNORE INTO ` + constants.TableEntryGroups + `
			(source_id, source_type, list_type, group_name)
			VALUES (?, ?, ?, ?)`)
		if err != nil {
			return fmt.Errorf("preparing entry groups insert: %w", err)
		}
		defer func() { _ = groupsStmt.Close() }() // nolint: errcheck

		for _, row := range dedupeEntryGroupRows(groups) {
			if row.SourceType == "" || row.ListType == "" || row.GroupName == "" {
				continue
			}
			if _, rowErr := groupsStmt.Exec(sourceID, row.SourceType, row.ListType, row.GroupName); rowErr != nil {
				return fmt.Errorf("inserting entry group for source %d: %w", sourceID, rowErr)
			}
		}

		categoriesStmt, err := tx.Prepare(`
			INSERT OR IGNORE INTO ` + constants.TableEntryCategories + `
			(source_id, source_type, list_type, category)
			VALUES (?, ?, ?, ?)`)
		if err != nil {
			return fmt.Errorf("preparing entry categories insert: %w", err)
		}
		defer func() { _ = categoriesStmt.Close() }() // nolint: errcheck

		for _, row := range dedupeEntryCategoryRows(categories) {
			if row.SourceType == "" || row.ListType == "" || row.Category == "" {
				continue
			}
			if _, rowErr := categoriesStmt.Exec(sourceID, row.SourceType, row.ListType, row.Category); rowErr != nil {
				return fmt.Errorf("inserting entry category for source %d: %w", sourceID, rowErr)
			}
		}

		return nil
	})
}

// SearchResult holds a single hit returned from a DB-backed search.
type SearchResult struct {
	SourceName        string
	GenericSourceType string
	ActualSourceType  string
	ListType          string
	Valid             bool
}

// SearchEntries searches dnstk_entries for rows whose entry matches the
// query.  When exactMatch is true only exact equality is used; otherwise a
// case-insensitive substring match (LIKE) is applied.
// Only valid entries are returned unless includeInvalid is true.
func (r *EntriesRepo) SearchEntries(
	_ context.Context,
	query string,
	exactMatch bool,
	includeInvalid bool,
) ([]SearchResult, error) {
	var sb strings.Builder
	sb.WriteString(`
		SELECT s.name, e.generic_source_type, e.actual_source_type, e.list_type, e.valid
		FROM ` + constants.TableEntries + ` e
		JOIN ` + constants.TableSources + ` s ON s.id = e.source_id
		WHERE `)

	args := []any{}
	if exactMatch {
		sb.WriteString("e.entry = ?")
		args = append(args, query)
	} else {
		sb.WriteString("e.entry LIKE ? ESCAPE '\\'")
		args = append(args, "%"+escapeLike(query)+"%")
	}

	if !includeInvalid {
		sb.WriteString(" AND e.valid = 1")
	}

	sb.WriteString(" ORDER BY s.name, e.generic_source_type")

	rows, err := r.db.conn.Query(sb.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("searching entries: %w", err)
	}
	defer rows.Close()

	var results []SearchResult
	for rows.Next() {
		var res SearchResult
		var validInt int
		if err := rows.Scan(&res.SourceName, &res.GenericSourceType, &res.ActualSourceType, &res.ListType, &validInt); err != nil {
			return nil, fmt.Errorf("scanning entry row: %w", err)
		}
		res.Valid = validInt == 1
		results = append(results, res)
	}
	return results, rows.Err()
}

// SearchConsolidatedEntries searches dnstk_consolidated_entries for matching rows.
func (r *EntriesRepo) SearchConsolidatedEntries(
	_ context.Context,
	query string,
	exactMatch bool,
) ([]SearchResult, error) {
	var sb strings.Builder
	sb.WriteString(`
		SELECT consolidation_type || COALESCE('/' || group_name, '') || COALESCE('/' || category, '') AS source_name,
		       generic_source_type, generic_source_type, list_type, valid
		FROM ` + constants.TableConsolidatedEntries + `
		WHERE `)

	args := []any{}
	if exactMatch {
		sb.WriteString("entry = ?")
		args = append(args, query)
	} else {
		sb.WriteString("entry LIKE ? ESCAPE '\\'")
		args = append(args, "%"+escapeLike(query)+"%")
	}

	sb.WriteString(" ORDER BY consolidation_type, generic_source_type")

	rows, err := r.db.conn.Query(sb.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("searching consolidated entries: %w", err)
	}
	defer rows.Close()

	var results []SearchResult
	for rows.Next() {
		var res SearchResult
		var validInt int
		if err := rows.Scan(&res.SourceName, &res.GenericSourceType, &res.ActualSourceType, &res.ListType, &validInt); err != nil {
			return nil, fmt.Errorf("scanning consolidated row: %w", err)
		}
		res.Valid = validInt == 1
		results = append(results, res)
	}
	return results, rows.Err()
}

// escapeLike escapes special LIKE characters in s.
func escapeLike(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, "%", `\%`)
	s = strings.ReplaceAll(s, "_", `\_`)
	return s
}

func dedupeEntryRows(rows []EntryRow) []EntryRow {
	if len(rows) == 0 {
		return nil
	}

	result := make([]EntryRow, 0, len(rows))
	seen := make(map[string]struct{}, len(rows))
	for _, row := range rows {
		key := fmt.Sprintf("%s|%s|%s|%s", row.Entry, row.GenericSourceType, row.ActualSourceType, row.ListType)
		if _, exists := seen[key]; exists {
			continue
		}
		seen[key] = struct{}{}
		result = append(result, row)
	}

	return result
}

func dedupeEntryGroupRows(rows []EntryGroupRow) []EntryGroupRow {
	if len(rows) == 0 {
		return nil
	}

	result := make([]EntryGroupRow, 0, len(rows))
	seen := make(map[string]struct{}, len(rows))
	for _, row := range rows {
		key := fmt.Sprintf("%s|%s|%s", row.SourceType, row.ListType, row.GroupName)
		if _, exists := seen[key]; exists {
			continue
		}
		seen[key] = struct{}{}
		result = append(result, row)
	}

	return result
}

func dedupeEntryCategoryRows(rows []EntryCategoryRow) []EntryCategoryRow {
	if len(rows) == 0 {
		return nil
	}

	result := make([]EntryCategoryRow, 0, len(rows))
	seen := make(map[string]struct{}, len(rows))
	for _, row := range rows {
		key := fmt.Sprintf("%s|%s|%s", row.SourceType, row.ListType, row.Category)
		if _, exists := seen[key]; exists {
			continue
		}
		seen[key] = struct{}{}
		result = append(result, row)
	}

	return result
}
