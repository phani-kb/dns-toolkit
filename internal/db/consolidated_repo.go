package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/phani-kb/dns-toolkit/internal/constants"
)

// ConsolidatedRepo provides operations for consolidation results.
type ConsolidatedRepo struct {
	db *DB
}

func NewConsolidatedRepo(db *DB) *ConsolidatedRepo {
	return &ConsolidatedRepo{db: db}
}

func (r *ConsolidatedRepo) ConsolidateGeneral(genericSourceType, listType string, valid bool) (int64, error) {
	result, err := r.db.writeConn.Exec(`
		INSERT INTO `+constants.TableConsolidatedEntries+` (entry, generic_source_type, list_type,
			consolidation_type, valid, source_count)
		SELECT e.entry, e.generic_source_type, e.list_type, 'general', ?,
			COUNT(DISTINCT e.source_id)
		FROM `+constants.TableEntries+` e
		JOIN `+constants.TableSources+` s ON e.source_id = s.id
		WHERE e.generic_source_type = ?
			AND e.list_type = ?
			AND e.valid = ?
			AND s.disabled = 0
			AND s.skip_general_consolidation = 0
		GROUP BY e.entry`,
		boolToInt(valid), genericSourceType, listType, boolToInt(valid))
	if err != nil {
		return 0, fmt.Errorf("consolidating general %s/%s: %w", genericSourceType, listType, err)
	}
	return result.RowsAffected()
}

func (r *ConsolidatedRepo) ConsolidateByGroup(
	genericSourceType, listType, groupName string,
	valid bool,
) (int64, error) {
	result, err := r.db.writeConn.Exec(`
		INSERT INTO `+constants.TableConsolidatedEntries+` (entry, generic_source_type, list_type,
			consolidation_type, group_name, valid, source_count)
		SELECT e.entry, e.generic_source_type, e.list_type, 'group', ?, ?,
			COUNT(DISTINCT e.source_id)
		FROM `+constants.TableEntries+` e
		JOIN `+constants.TableSources+` s ON e.source_id = s.id
		JOIN `+constants.TableEntryGroups+` eg ON eg.source_id = e.source_id
			AND eg.source_type = e.generic_source_type
			AND eg.list_type = e.list_type
		WHERE e.generic_source_type = ?
			AND e.list_type = ?
			AND e.valid = ?
			AND s.disabled = 0
			AND s.skip_groups_consolidation = 0
			AND eg.group_name = ?
		GROUP BY e.entry`,
		groupName, boolToInt(valid), genericSourceType, listType, boolToInt(valid), groupName)
	if err != nil {
		return 0, fmt.Errorf("consolidating group %s %s/%s: %w", groupName, genericSourceType, listType, err)
	}
	return result.RowsAffected()
}

func (r *ConsolidatedRepo) ConsolidateByCategory(
	genericSourceType, listType, category string,
	valid bool,
) (int64, error) {
	result, err := r.db.writeConn.Exec(`
		INSERT INTO `+constants.TableConsolidatedEntries+` (entry, generic_source_type, list_type,
			consolidation_type, category, valid, source_count)
		SELECT e.entry, e.generic_source_type, e.list_type, 'category', ?, ?,
			COUNT(DISTINCT e.source_id)
		FROM `+constants.TableEntries+` e
		JOIN `+constants.TableSources+` s ON e.source_id = s.id
		JOIN `+constants.TableEntryCategories+` ec ON ec.source_id = e.source_id
			AND ec.source_type = e.generic_source_type
			AND ec.list_type = e.list_type
		WHERE e.generic_source_type = ?
			AND e.list_type = ?
			AND e.valid = ?
			AND s.disabled = 0
			AND s.skip_categories_consolidation = 0
			AND ec.category = ?
		GROUP BY e.entry`,
		category, boolToInt(valid), genericSourceType, listType, boolToInt(valid), category)
	if err != nil {
		return 0, fmt.Errorf("consolidating category %s %s/%s: %w", category, genericSourceType, listType, err)
	}
	return result.RowsAffected()
}

func (r *ConsolidatedRepo) ConsolidateByCountry(
	genericSourceType, listType, countryCode string,
	valid bool,
) (int64, error) {
	result, err := r.db.writeConn.Exec(`
		INSERT INTO `+constants.TableConsolidatedEntries+` (entry, generic_source_type, list_type,
			consolidation_type, category, valid, source_count)
		SELECT e.entry, e.generic_source_type, e.list_type, 'category', ?, ?,
			COUNT(DISTINCT e.source_id)
		FROM `+constants.TableEntries+` e
		JOIN `+constants.TableSources+` s ON e.source_id = s.id
		JOIN `+constants.TableSourceCountries+` sc ON sc.source_id = s.id
		WHERE e.generic_source_type = ?
			AND e.list_type = ?
			AND e.valid = ?
			AND s.disabled = 0
			AND sc.country_code = ?
		GROUP BY e.entry`,
		countryCode, boolToInt(valid), genericSourceType, listType, boolToInt(valid), countryCode)
	if err != nil {
		return 0, fmt.Errorf("consolidating country %s %s/%s: %w", countryCode, genericSourceType, listType, err)
	}
	return result.RowsAffected()
}

// FilterConsolidatedEntries removes entries from blocklist that appear in the allowlist.
func (r *ConsolidatedRepo) FilterConsolidatedEntries(genericSourceType, consolidationType string,
	groupName, category string,
) (int64, error) {
	var whereClause string
	var args []interface{}

	args = append(args, genericSourceType, consolidationType)

	if groupName != "" {
		whereClause = " AND ce.group_name = ?"
		args = append(args, groupName)
	}
	if category != "" {
		whereClause = " AND ce.category = ?"
		args = append(args, category)
	}

	result, err := r.db.writeConn.Exec(fmt.Sprintf(`
		DELETE FROM `+constants.TableConsolidatedEntries+`
		WHERE id IN (
			SELECT ce.id FROM `+constants.TableConsolidatedEntries+` ce
			WHERE ce.generic_source_type = ?
				AND ce.consolidation_type = ?
				AND ce.list_type = 'blocklist'
				%s
				AND ce.entry IN (
					SELECT ce2.entry FROM `+constants.TableConsolidatedEntries+` ce2
					WHERE ce2.generic_source_type = ce.generic_source_type
						AND ce2.consolidation_type = ce.consolidation_type
						AND ce2.list_type = 'allowlist'
						AND ce2.valid = 1
				)
		)`, whereClause), args...)
	if err != nil {
		return 0, fmt.Errorf("filtering consolidated entries: %w", err)
	}
	return result.RowsAffected()
}

func (r *ConsolidatedRepo) GetConsolidatedEntries(genericSourceType, listType,
	consolidationType string, valid bool,
) ([]string, error) {
	rows, err := r.db.readConn.Query(`
		SELECT entry FROM `+constants.TableConsolidatedEntries+`
		WHERE generic_source_type = ? AND list_type = ? AND consolidation_type = ? AND valid = ?
		ORDER BY entry`,
		genericSourceType, listType, consolidationType, boolToInt(valid))
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }() // nolint: errcheck

	var entries []string
	for rows.Next() {
		var entry string
		if err := rows.Scan(&entry); err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, rows.Err()
}

func (r *ConsolidatedRepo) GetConsolidatedCount(genericSourceType, listType,
	consolidationType string, valid bool,
) (int64, error) {
	var count int64
	err := r.db.readConn.QueryRow(`
		SELECT COUNT(*) FROM `+constants.TableConsolidatedEntries+`
		WHERE generic_source_type = ? AND list_type = ? AND consolidation_type = ? AND valid = ?`,
		genericSourceType, listType, consolidationType, boolToInt(valid)).Scan(&count)
	return count, err
}

func (r *ConsolidatedRepo) ClearConsolidated(consolidationType string) error {
	if _, err := r.db.writeConn.Exec("drop index if exists idx_consolidated_lookup"); err != nil {
		return fmt.Errorf("dropping idx_consolidated_lookup: %w", err)
	}
	if _, err := r.db.writeConn.Exec("drop index if exists idx_consolidated_type"); err != nil {
		return fmt.Errorf("dropping idx_consolidated_type: %w", err)
	}

	_, err := r.db.writeConn.Exec("DELETE FROM "+constants.TableConsolidatedEntries+
		" WHERE consolidation_type = ?", consolidationType)
	if err != nil {
		return err
	}

	if _, err := r.db.writeConn.Exec(
		`CREATE INDEX IF NOT EXISTS idx_consolidated_lookup ON ` +
			constants.TableConsolidatedEntries +
			` (entry, generic_source_type, list_type, consolidation_type)`,
	); err != nil {
		return fmt.Errorf("recreating idx_consolidated_lookup: %w", err)
	}
	if _, err := r.db.writeConn.Exec(
		`CREATE INDEX IF NOT EXISTS idx_consolidated_type ON ` +
			constants.TableConsolidatedEntries +
			` (consolidation_type, generic_source_type, list_type)`,
	); err != nil {
		return fmt.Errorf("recreating idx_consolidated_type: %w", err)
	}

	return nil
}

func (r *ConsolidatedRepo) ClearAllConsolidated() error {
	_, err := r.db.writeConn.Exec("DELETE FROM " + constants.TableConsolidatedEntries)
	return err
}

type ConsolidatedGroup struct {
	GenericSourceType string
	ListType          string
	ConsolidationType string
	GroupName         string
	Category          string
	Valid             bool
	Count             int
}

// ListConsolidatedGroups returns all distinct groups and categories for a given consolidation type.
func (r *ConsolidatedRepo) ListConsolidatedGroups(consolidationType string) ([]ConsolidatedGroup, error) {
	query := `
		SELECT generic_source_type, list_type, consolidation_type,
			COALESCE(group_name, '') AS group_name,
			COALESCE(category, '') AS category,
			valid, COUNT(*) AS count
		FROM ` + constants.TableConsolidatedEntries + `
		WHERE consolidation_type = ?
		GROUP BY generic_source_type, list_type, consolidation_type, group_name, category, valid
		ORDER BY generic_source_type, list_type, group_name, category
	`

	rows, err := r.db.readConn.Query(query, consolidationType)
	if err != nil {
		return nil, fmt.Errorf("querying consolidated groups: %w", err)
	}
	defer rows.Close() // nolint: errcheck

	var results []ConsolidatedGroup
	for rows.Next() {
		var g ConsolidatedGroup
		var validInt int
		if err := rows.Scan(
			&g.GenericSourceType, &g.ListType, &g.ConsolidationType,
			&g.GroupName, &g.Category, &validInt, &g.Count,
		); err != nil {
			return nil, fmt.Errorf("scanning consolidated group: %w", err)
		}
		g.Valid = validInt == 1
		results = append(results, g)
	}
	return results, rows.Err()
}

// GetConsolidatedEntriesByGroup returns entries filtered by group or category.
func (r *ConsolidatedRepo) GetConsolidatedEntriesByGroup(
	genericSourceType, listType, consolidationType, groupName, category string, valid bool,
) ([]string, error) {
	query := `SELECT entry FROM ` + constants.TableConsolidatedEntries + `
		WHERE generic_source_type = ? AND list_type = ? AND consolidation_type = ? AND valid = ?`
	args := []any{genericSourceType, listType, consolidationType, boolToInt(valid)}

	if groupName != "" {
		query += " AND group_name = ?"
		args = append(args, groupName)
	}
	if category != "" {
		query += " AND category = ?"
		args = append(args, category)
	}

	query += " ORDER BY entry"

	rows, err := r.db.readConn.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("querying consolidated entries by group: %w", err)
	}
	defer rows.Close() // nolint: errcheck

	var entries []string
	for rows.Next() {
		var entry string
		if err := rows.Scan(&entry); err != nil {
			return nil, fmt.Errorf("scanning entry: %w", err)
		}
		entries = append(entries, entry)
	}
	return entries, rows.Err()
}

type ConsolidatedEntryRow struct {
	Entry             string
	GenericSourceType string
	ListType          string
	ConsolidationType string // "general", "group", "category"
	GroupName         string // non-empty when ConsolidationType == "group"
	Category          string // non-empty when ConsolidationType == "category"
	Valid             bool
	SourceCount       int
}

func (r *ConsolidatedRepo) BulkInsertEntries(ctx context.Context, rows []ConsolidatedEntryRow) (int64, error) {
	if len(rows) == 0 {
		return 0, nil
	}

	const batchSize = constants.BulkInsertBatchSize
	var inserted int64

	err := r.db.InTransaction(ctx, func(tx *sql.Tx) error {
		for i := 0; i < len(rows); i += batchSize {
			end := i + batchSize
			if end > len(rows) {
				end = len(rows)
			}
			batch := rows[i:end]

			query := `INSERT INTO ` + constants.TableConsolidatedEntries +
				` (entry, generic_source_type, list_type, consolidation_type, group_name, category, valid, source_count) VALUES `
			args := make([]interface{}, 0, len(batch)*8)
			for j, row := range batch {
				if j > 0 {
					query += ","
				}
				query += "(?,?,?,?,?,?,?,?)"

				var groupName, category *string
				if row.GroupName != "" {
					groupName = &row.GroupName
				}
				if row.Category != "" {
					category = &row.Category
				}
				sourceCount := row.SourceCount
				if sourceCount <= 0 {
					sourceCount = 1
				}
				args = append(args, row.Entry, row.GenericSourceType, row.ListType,
					row.ConsolidationType, groupName, category, boolToInt(row.Valid), sourceCount)
			}

			result, execErr := tx.ExecContext(ctx, query, args...)
			if execErr != nil {
				return fmt.Errorf("batch insert consolidated entries: %w", execErr)
			}
			affected, _ := result.RowsAffected() // nolint: errcheck
			inserted += affected
		}
		return nil
	})
	return inserted, err
}
