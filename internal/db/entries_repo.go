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
	defer rows.Close() // nolint: errcheck

	var results []SearchResult
	for rows.Next() {
		var res SearchResult
		var validInt int
		if err := rows.Scan(&res.SourceName, &res.GenericSourceType, &res.ActualSourceType,
			&res.ListType, &validInt); err != nil {
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
	defer rows.Close() // nolint: errcheck

	var results []SearchResult
	for rows.Next() {
		var res SearchResult
		var validInt int
		if err := rows.Scan(&res.SourceName, &res.GenericSourceType, &res.ActualSourceType,
			&res.ListType, &validInt); err != nil {
			return nil, fmt.Errorf("scanning consolidated row: %w", err)
		}
		res.Valid = validInt == 1
		results = append(results, res)
	}
	return results, rows.Err()
}

type ConflictCountRow struct {
	Entry             string
	GenericSourceType string
	BlockSources      string // comma-separated source names
	AllowSources      string // comma-separated source names
	BlockCount        int
	AllowCount        int
}

type AllEntryCountRow struct {
	Entry             string
	GenericSourceType string
	BlockCount        int
	AllowCount        int
}

func (r *EntriesRepo) GetConflictCounts(_ context.Context) ([]ConflictCountRow, error) {
	query := `
		SELECT
			e.entry,
			e.generic_source_type,
			COUNT(DISTINCT CASE WHEN e.list_type = 'blocklist' THEN s.name END) AS block_count,
			COUNT(DISTINCT CASE WHEN e.list_type = 'allowlist' THEN s.name END) AS allow_count,
			GROUP_CONCAT(DISTINCT CASE WHEN e.list_type = 'blocklist' THEN s.name END) AS block_sources,
			GROUP_CONCAT(DISTINCT CASE WHEN e.list_type = 'allowlist' THEN s.name END) AS allow_sources
		FROM ` + constants.TableEntries + ` e
		JOIN ` + constants.TableSources + ` s ON s.id = e.source_id
		WHERE e.valid = 1
			AND s.disabled = 0
			AND s.skip_general_consolidation = 0
		GROUP BY e.entry, e.generic_source_type
		HAVING block_count > 0 AND allow_count > 0
		ORDER BY e.entry
	`

	rows, err := r.db.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("querying conflict counts: %w", err)
	}
	defer rows.Close() // nolint: errcheck

	var results []ConflictCountRow
	for rows.Next() {
		var row ConflictCountRow
		if err := rows.Scan(
			&row.Entry,
			&row.GenericSourceType,
			&row.BlockCount,
			&row.AllowCount,
			&row.BlockSources,
			&row.AllowSources,
		); err != nil {
			return nil, fmt.Errorf("scanning conflict count row: %w", err)
		}
		results = append(results, row)
	}
	return results, rows.Err()
}

func (r *EntriesRepo) GetAllEntryCounts(_ context.Context) ([]AllEntryCountRow, error) {
	query := `
		SELECT
			e.entry,
			e.generic_source_type,
			COUNT(DISTINCT CASE WHEN e.list_type = 'blocklist' THEN s.name END) AS block_count,
			COUNT(DISTINCT CASE WHEN e.list_type = 'allowlist' THEN s.name END) AS allow_count
		FROM ` + constants.TableEntries + ` e
		JOIN ` + constants.TableSources + ` s ON s.id = e.source_id
		WHERE e.valid = 1
			AND s.disabled = 0
			AND s.skip_general_consolidation = 0
		GROUP BY e.entry, e.generic_source_type
		ORDER BY e.entry
	`

	rows, err := r.db.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("querying all entry counts: %w", err)
	}
	defer rows.Close() // nolint: errcheck

	var results []AllEntryCountRow
	for rows.Next() {
		var row AllEntryCountRow
		if err := rows.Scan(
			&row.Entry,
			&row.GenericSourceType,
			&row.BlockCount,
			&row.AllowCount,
		); err != nil {
			return nil, fmt.Errorf("scanning entry count row: %w", err)
		}
		results = append(results, row)
	}
	return results, rows.Err()
}

func (r *EntriesRepo) GetSourcesForEntry(
	_ context.Context,
	entry, genericSourceType string,
) (blockSources, allowSources []string, err error) {
	query := `
		SELECT s.name, e.list_type
		FROM ` + constants.TableEntries + ` e
		JOIN ` + constants.TableSources + ` s ON s.id = e.source_id
		WHERE e.entry = ?
			AND e.generic_source_type = ?
			AND e.valid = 1
			AND s.disabled = 0
			AND s.skip_general_consolidation = 0
		ORDER BY s.name
	`

	rows, err2 := r.db.conn.Query(query, entry, genericSourceType)
	if err2 != nil {
		return nil, nil, fmt.Errorf("querying sources for entry: %w", err2)
	}
	defer rows.Close() // nolint: errcheck

	for rows.Next() {
		var name, listType string
		if err2 := rows.Scan(&name, &listType); err2 != nil {
			return nil, nil, fmt.Errorf("scanning source row: %w", err2)
		}
		switch listType {
		case "blocklist":
			blockSources = append(blockSources, name)
		case "allowlist":
			allowSources = append(allowSources, name)
		}
	}
	return blockSources, allowSources, rows.Err()
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

// ConsolidationEntry represents an entry for consolidation with its metadata.
type ConsolidationEntry struct {
	Entry             string
	GenericSourceType string
	ActualSourceType  string
	ListType          string
	SourceName        string
	MustConsider      bool
}

// GetEntriesByCategory returns all valid entries for sources that have the given category.
func (r *EntriesRepo) GetEntriesByCategory(
	_ context.Context,
	category string,
	genericSourceType string,
	listType string,
) ([]ConsolidationEntry, error) {
	query := `
		SELECT DISTINCT
			e.entry,
			e.generic_source_type,
			e.actual_source_type,
			e.list_type,
			e.must_consider,
			s.name
		FROM ` + constants.TableEntries + ` e
		JOIN ` + constants.TableSources + ` s ON s.id = e.source_id
		JOIN ` + constants.TableEntryCategories + ` c ON c.source_id = e.source_id
			AND c.source_type = e.generic_source_type
			AND c.list_type = e.list_type
		WHERE c.category = ?
			AND e.generic_source_type = ?
			AND e.list_type = ?
			AND e.valid = 1
			AND s.disabled = 0
			AND s.skip_categories_consolidation = 0
		ORDER BY e.entry
	`

	rows, err := r.db.conn.Query(query, category, genericSourceType, listType)
	if err != nil {
		return nil, fmt.Errorf("querying entries by category: %w", err)
	}
	defer rows.Close() // nolint: errcheck

	var results []ConsolidationEntry
	for rows.Next() {
		var entry ConsolidationEntry
		var mustConsiderInt int
		if err := rows.Scan(
			&entry.Entry,
			&entry.GenericSourceType,
			&entry.ActualSourceType,
			&entry.ListType,
			&mustConsiderInt,
			&entry.SourceName,
		); err != nil {
			return nil, fmt.Errorf("scanning entry row: %w", err)
		}
		entry.MustConsider = mustConsiderInt == 1
		results = append(results, entry)
	}
	return results, rows.Err()
}

// GetEntriesByGroup returns all valid entries for sources that have the given group.
func (r *EntriesRepo) GetEntriesByGroup(
	_ context.Context,
	groupName string,
	genericSourceType string,
	listType string,
) ([]ConsolidationEntry, error) {
	query := `
		SELECT DISTINCT
			e.entry,
			e.generic_source_type,
			e.actual_source_type,
			e.list_type,
			e.must_consider,
			s.name
		FROM ` + constants.TableEntries + ` e
		JOIN ` + constants.TableSources + ` s ON s.id = e.source_id
		JOIN ` + constants.TableEntryGroups + ` g ON g.source_id = e.source_id
			AND g.source_type = e.generic_source_type
			AND g.list_type = e.list_type
		WHERE g.group_name = ?
			AND e.generic_source_type = ?
			AND e.list_type = ?
			AND e.valid = 1
			AND s.disabled = 0
			AND s.skip_groups_consolidation = 0
		ORDER BY e.entry
	`

	rows, err := r.db.conn.Query(query, groupName, genericSourceType, listType)
	if err != nil {
		return nil, fmt.Errorf("querying entries by group: %w", err)
	}
	defer rows.Close() // nolint: errcheck

	var results []ConsolidationEntry
	for rows.Next() {
		var entry ConsolidationEntry
		var mustConsiderInt int
		if err := rows.Scan(
			&entry.Entry,
			&entry.GenericSourceType,
			&entry.ActualSourceType,
			&entry.ListType,
			&mustConsiderInt,
			&entry.SourceName,
		); err != nil {
			return nil, fmt.Errorf("scanning entry row: %w", err)
		}
		entry.MustConsider = mustConsiderInt == 1
		results = append(results, entry)
	}
	return results, rows.Err()
}

// GetEntriesForGeneralConsolidation returns all valid entries for general consolidation
// (not filtered by group or category).
func (r *EntriesRepo) GetEntriesForGeneralConsolidation(
	_ context.Context,
	genericSourceType string,
	listType string,
) ([]ConsolidationEntry, error) {
	query := `
		SELECT DISTINCT
			e.entry,
			e.generic_source_type,
			e.actual_source_type,
			e.list_type,
			e.must_consider,
			s.name
		FROM ` + constants.TableEntries + ` e
		JOIN ` + constants.TableSources + ` s ON s.id = e.source_id
		WHERE e.generic_source_type = ?
			AND e.list_type = ?
			AND e.valid = 1
			AND s.disabled = 0
			AND s.skip_general_consolidation = 0
		ORDER BY e.entry
	`

	rows, err := r.db.conn.Query(query, genericSourceType, listType)
	if err != nil {
		return nil, fmt.Errorf("querying entries for general consolidation: %w", err)
	}
	defer rows.Close() // nolint: errcheck

	var results []ConsolidationEntry
	for rows.Next() {
		var entry ConsolidationEntry
		var mustConsiderInt int
		if err := rows.Scan(
			&entry.Entry,
			&entry.GenericSourceType,
			&entry.ActualSourceType,
			&entry.ListType,
			&mustConsiderInt,
			&entry.SourceName,
		); err != nil {
			return nil, fmt.Errorf("scanning entry row: %w", err)
		}
		entry.MustConsider = mustConsiderInt == 1
		results = append(results, entry)
	}
	return results, rows.Err()
}

// GetInvalidEntriesForGeneralConsolidation returns all invalid entries for general consolidation.
func (r *EntriesRepo) GetInvalidEntriesForGeneralConsolidation(
	_ context.Context,
	genericSourceType string,
	listType string,
) ([]ConsolidationEntry, error) {
	query := `
		SELECT DISTINCT
			e.entry,
			e.generic_source_type,
			e.actual_source_type,
			e.list_type,
			e.must_consider,
			s.name
		FROM ` + constants.TableEntries + ` e
		JOIN ` + constants.TableSources + ` s ON s.id = e.source_id
		WHERE e.generic_source_type = ?
			AND e.list_type = ?
			AND e.valid = 0
			AND s.disabled = 0
			AND s.skip_general_consolidation = 0
		ORDER BY e.entry
	`

	rows, err := r.db.conn.Query(query, genericSourceType, listType)
	if err != nil {
		return nil, fmt.Errorf("querying invalid entries for general consolidation: %w", err)
	}
	defer rows.Close() // nolint: errcheck

	var results []ConsolidationEntry
	for rows.Next() {
		var entry ConsolidationEntry
		var mustConsiderInt int
		if err := rows.Scan(
			&entry.Entry,
			&entry.GenericSourceType,
			&entry.ActualSourceType,
			&entry.ListType,
			&mustConsiderInt,
			&entry.SourceName,
		); err != nil {
			return nil, fmt.Errorf("scanning invalid entry row: %w", err)
		}
		entry.MustConsider = mustConsiderInt == 1
		results = append(results, entry)
	}
	return results, rows.Err()
}

// GetUniqueCategoriesFromDB returns unique categories from entry_categories table
// for enabled sources not skipping categories consolidation.
func (r *EntriesRepo) GetUniqueCategoriesFromDB(_ context.Context) ([]string, error) {
	query := `
		SELECT DISTINCT c.category
		FROM ` + constants.TableEntryCategories + ` c
		JOIN ` + constants.TableSources + ` s ON s.id = c.source_id
		WHERE s.disabled = 0
			AND s.skip_categories_consolidation = 0
			AND c.category != ''
		ORDER BY c.category COLLATE NOCASE
	`

	rows, err := r.db.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("querying unique categories: %w", err)
	}
	defer rows.Close() // nolint: errcheck

	var categories []string
	for rows.Next() {
		var cat string
		if err := rows.Scan(&cat); err != nil {
			return nil, fmt.Errorf("scanning category: %w", err)
		}
		categories = append(categories, cat)
	}
	return categories, rows.Err()
}

// GetUniqueGroupsFromDB returns unique groups from entry_groups table
// for enabled sources not skipping groups consolidation.
func (r *EntriesRepo) GetUniqueGroupsFromDB(_ context.Context) ([]string, error) {
	query := `
		SELECT DISTINCT g.group_name
		FROM ` + constants.TableEntryGroups + ` g
		JOIN ` + constants.TableSources + ` s ON s.id = g.source_id
		WHERE s.disabled = 0
			AND s.skip_groups_consolidation = 0
			AND g.group_name != ''
		ORDER BY g.group_name COLLATE NOCASE
	`

	rows, err := r.db.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("querying unique groups: %w", err)
	}
	defer rows.Close() // nolint: errcheck

	var groups []string
	for rows.Next() {
		var group string
		if err := rows.Scan(&group); err != nil {
			return nil, fmt.Errorf("scanning group: %w", err)
		}
		groups = append(groups, group)
	}
	return groups, rows.Err()
}

// GetGenericSourceTypesFromDB returns unique generic source types from entries.
func (r *EntriesRepo) GetGenericSourceTypesFromDB(_ context.Context) ([]string, error) {
	query := `
		SELECT DISTINCT e.generic_source_type
		FROM ` + constants.TableEntries + ` e
		JOIN ` + constants.TableSources + ` s ON s.id = e.source_id
		WHERE s.disabled = 0
			AND e.valid = 1
		ORDER BY e.generic_source_type COLLATE NOCASE
	`

	rows, err := r.db.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("querying generic source types: %w", err)
	}
	defer rows.Close() // nolint: errcheck

	var types []string
	for rows.Next() {
		var gst string
		if err := rows.Scan(&gst); err != nil {
			return nil, fmt.Errorf("scanning source type: %w", err)
		}
		types = append(types, gst)
	}
	return types, rows.Err()
}
