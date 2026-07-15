package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/phani-kb/dns-toolkit/internal/constants"
)

type OverlapRepo struct {
	db *DB
}

func NewOverlapRepo(db *DB) *OverlapRepo {
	return &OverlapRepo{db: db}
}

type OverlapResultRow struct {
	SourceName        string
	TargetName        string
	GenericSourceType string
	SourceListType    string
	TargetListType    string
	OverlapCount      int
	SourceCount       int
	TargetCount       int
	OverlapPercent    float64
}

type SourceEntryCount struct {
	SourceName string
	SourceID   int64
	EntryCount int
}

// GetSourceEntryCounts returns all sources with their entry counts for a given
// generic source type and list type.
func (r *OverlapRepo) GetSourceEntryCounts(
	_ context.Context,
	genericSourceType string,
	listType string,
) ([]SourceEntryCount, error) {
	query := `
		SELECT s.id, s.name, COUNT(DISTINCT e.entry) AS entry_count
		FROM ` + constants.TableEntries + ` e
		JOIN ` + constants.TableSources + ` s ON s.id = e.source_id
		WHERE e.generic_source_type = ?
			AND e.list_type = ?
			AND e.valid = 1
			AND s.disabled = 0
		GROUP BY s.id, s.name
		HAVING entry_count > 0
		ORDER BY s.name
	`

	rows, err := r.db.conn.Query(query, genericSourceType, listType)
	if err != nil {
		return nil, fmt.Errorf("querying source entry counts: %w", err)
	}
	defer rows.Close() // nolint: errcheck

	var results []SourceEntryCount
	for rows.Next() {
		var row SourceEntryCount
		if err := rows.Scan(&row.SourceID, &row.SourceName, &row.EntryCount); err != nil {
			return nil, fmt.Errorf("scanning source entry count: %w", err)
		}
		results = append(results, row)
	}
	return results, rows.Err()
}

// ComputePairOverlap computes the overlap count between two sources.
func (r *OverlapRepo) ComputePairOverlap(
	_ context.Context,
	sourceID, targetID int64,
	genericSourceType string,
) (int, error) {
	query := `
		SELECT COUNT(*) FROM (
			SELECT DISTINCT a.entry
			FROM ` + constants.TableEntries + ` a
			JOIN ` + constants.TableEntries + ` b
				ON a.entry = b.entry
				AND a.generic_source_type = b.generic_source_type
			WHERE a.source_id = ?
				AND b.source_id = ?
				AND a.generic_source_type = ?
				AND a.valid = 1
				AND b.valid = 1
		)
	`

	var count int
	if err := r.db.conn.QueryRow(query, sourceID, targetID, genericSourceType).Scan(&count); err != nil {
		return 0, fmt.Errorf("computing pair overlap: %w", err)
	}
	return count, nil
}

// PersistOverlapResults stores computed overlap results in the database.
func (r *OverlapRepo) PersistOverlapResults(ctx context.Context, results []OverlapResultRow) error {
	if len(results) == 0 {
		return nil
	}

	return r.db.InTransaction(ctx, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare(`
			INSERT INTO ` + constants.TableOverlapResults + `
			(source_name, target_name, generic_source_type, source_list_type, target_list_type,
			 overlap_count, source_count, target_count, overlap_percent)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`)
		if err != nil {
			return fmt.Errorf("preparing overlap insert: %w", err)
		}
		defer func() { _ = stmt.Close() }() // nolint: errcheck

		for _, row := range results {
			if _, err := stmt.Exec(
				row.SourceName, row.TargetName, row.GenericSourceType,
				row.SourceListType, row.TargetListType,
				row.OverlapCount, row.SourceCount, row.TargetCount, row.OverlapPercent,
			); err != nil {
				return fmt.Errorf("inserting overlap result: %w", err)
			}
		}
		return nil
	})
}

// ClearOverlapResults removes all rows from the overlap results table.
func (r *OverlapRepo) ClearOverlapResults() error {
	_, err := r.db.conn.Exec("DELETE FROM " + constants.TableOverlapResults)
	return err
}

// ListOverlapResults returns all stored overlap results.
func (r *OverlapRepo) ListOverlapResults(_ context.Context) ([]OverlapResultRow, error) {
	query := `
		SELECT source_name, target_name, generic_source_type,
			source_list_type, target_list_type,
			overlap_count, source_count, target_count, overlap_percent
		FROM ` + constants.TableOverlapResults + `
		ORDER BY generic_source_type, source_name, target_name
	`

	rows, err := r.db.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("querying overlap results: %w", err)
	}
	defer rows.Close() // nolint: errcheck

	var results []OverlapResultRow
	for rows.Next() {
		var row OverlapResultRow
		if err := rows.Scan(
			&row.SourceName, &row.TargetName, &row.GenericSourceType,
			&row.SourceListType, &row.TargetListType,
			&row.OverlapCount, &row.SourceCount, &row.TargetCount, &row.OverlapPercent,
		); err != nil {
			return nil, fmt.Errorf("scanning overlap result: %w", err)
		}
		results = append(results, row)
	}
	return results, rows.Err()
}
