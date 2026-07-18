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
	ListType   string
	SourceID   int64
	EntryCount int
}

// SourceListKey uniquely identifies a source + list type combination.
type SourceListKey struct {
	ListType string
	SourceID int64
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
		row.ListType = listType
		if err := rows.Scan(&row.SourceID, &row.SourceName, &row.EntryCount); err != nil {
			return nil, fmt.Errorf("scanning source entry count: %w", err)
		}
		results = append(results, row)
	}
	return results, rows.Err()
}

// GetSourceEntryCountsAllListTypes returns all sources with their entry counts.
func (r *OverlapRepo) GetSourceEntryCountsAllListTypes(
	_ context.Context,
	genericSourceType string,
) ([]SourceEntryCount, error) {
	query := `
		SELECT s.id, s.name, e.list_type, COUNT(DISTINCT e.entry) AS entry_count
		FROM ` + constants.TableEntries + ` e
		JOIN ` + constants.TableSources + ` s ON s.id = e.source_id
		WHERE e.generic_source_type = ?
			AND e.valid = 1
			AND s.disabled = 0
		GROUP BY s.id, s.name, e.list_type
		HAVING entry_count > 0
		ORDER BY s.name, e.list_type
	`

	rows, err := r.db.conn.Query(query, genericSourceType)
	if err != nil {
		return nil, fmt.Errorf("querying source entry counts all list types: %w", err)
	}
	defer rows.Close() // nolint: errcheck

	var results []SourceEntryCount
	for rows.Next() {
		var row SourceEntryCount
		if err := rows.Scan(&row.SourceID, &row.SourceName, &row.ListType, &row.EntryCount); err != nil {
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
	listType string,
) (int, error) {
	query := `
		SELECT COUNT(*) FROM (
			SELECT DISTINCT a.entry
			FROM ` + constants.TableEntries + ` a
			JOIN ` + constants.TableEntries + ` b
				ON a.entry = b.entry
				AND a.generic_source_type = b.generic_source_type
				AND a.list_type = b.list_type
			WHERE a.source_id = ?
				AND b.source_id = ?
				AND a.generic_source_type = ?
				AND a.list_type = ?
				AND a.valid = 1
				AND b.valid = 1
		)
	`

	var count int
	if err := r.db.conn.QueryRow(query, sourceID, targetID, genericSourceType, listType).Scan(&count); err != nil {
		return 0, fmt.Errorf("computing pair overlap: %w", err)
	}
	return count, nil
}

// PairOverlap holds the overlap count between two sources.
type PairOverlap struct {
	SourceListType string
	TargetListType string
	SourceID       int64
	TargetID       int64
	OverlapCount   int
}

// ComputeAllPairOverlaps loads all entries for a type/list.
func (r *OverlapRepo) ComputeAllPairOverlaps(
	_ context.Context,
	genericSourceType string,
	listType string,
) ([]PairOverlap, error) {
	query := `
		SELECT DISTINCT e.entry, e.source_id
		FROM ` + constants.TableEntries + ` e
		JOIN ` + constants.TableSources + ` s ON s.id = e.source_id
		WHERE e.generic_source_type = ?
			AND e.list_type = ?
			AND e.valid = 1
			AND s.disabled = 0
	`

	rows, err := r.db.conn.Query(query, genericSourceType, listType)
	if err != nil {
		return nil, fmt.Errorf("loading entries for overlap: %w", err)
	}
	defer rows.Close() // nolint: errcheck

	entrySources := make(map[string][]int64)
	for rows.Next() {
		var entry string
		var sourceID int64
		if err := rows.Scan(&entry, &sourceID); err != nil {
			return nil, fmt.Errorf("scanning entry for overlap: %w", err)
		}
		entrySources[entry] = append(entrySources[entry], sourceID)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	type pairKey struct{ a, b int64 }
	pairCounts := make(map[pairKey]int)

	for _, sourceIDs := range entrySources {
		if len(sourceIDs) < 2 {
			continue
		}
		for i := 0; i < len(sourceIDs); i++ {
			for j := i + 1; j < len(sourceIDs); j++ {
				a, b := sourceIDs[i], sourceIDs[j]
				if a > b {
					a, b = b, a
				}
				pairCounts[pairKey{a, b}]++
			}
		}
	}

	results := make([]PairOverlap, 0, len(pairCounts))
	for k, count := range pairCounts {
		results = append(results, PairOverlap{
			SourceID:       k.a,
			TargetID:       k.b,
			SourceListType: listType,
			TargetListType: listType,
			OverlapCount:   count,
		})
	}
	return results, nil
}

// ComputeAllPairOverlapsAcrossListTypes loads all entries for a generic source type
// across all list types.
func (r *OverlapRepo) ComputeAllPairOverlapsAcrossListTypes(
	_ context.Context,
	genericSourceType string,
) ([]PairOverlap, error) {
	query := `
		SELECT DISTINCT e.entry, e.source_id, e.list_type
		FROM ` + constants.TableEntries + ` e
		JOIN ` + constants.TableSources + ` s ON s.id = e.source_id
		WHERE e.generic_source_type = ?
			AND e.valid = 1
			AND s.disabled = 0
	`

	rows, err := r.db.conn.Query(query, genericSourceType)
	if err != nil {
		return nil, fmt.Errorf("loading entries for overlap: %w", err)
	}
	defer rows.Close() // nolint: errcheck

	// entry → list of (sourceID, listType)
	type sourceEntry struct {
		listType string
		sourceID int64
	}
	entrySources := make(map[string][]sourceEntry)
	for rows.Next() {
		var entry string
		var se sourceEntry
		if err := rows.Scan(&entry, &se.sourceID, &se.listType); err != nil {
			return nil, fmt.Errorf("scanning entry for overlap: %w", err)
		}
		entrySources[entry] = append(entrySources[entry], se)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	type pairKey struct {
		srcListType, tgtListType string
		srcID, tgtID             int64
	}
	pairCounts := make(map[pairKey]int)

	for _, sources := range entrySources {
		if len(sources) < 2 {
			continue
		}
		for i := 0; i < len(sources); i++ {
			for j := i + 1; j < len(sources); j++ {
				a, b := sources[i], sources[j]
				if a.sourceID > b.sourceID || (a.sourceID == b.sourceID && a.listType > b.listType) {
					a, b = b, a
				}
				pairCounts[pairKey{a.listType, b.listType, a.sourceID, b.sourceID}]++
			}
		}
	}

	results := make([]PairOverlap, 0, len(pairCounts))
	for k, count := range pairCounts {
		results = append(results, PairOverlap{
			SourceID:       k.srcID,
			TargetID:       k.tgtID,
			SourceListType: k.srcListType,
			TargetListType: k.tgtListType,
			OverlapCount:   count,
		})
	}
	return results, nil
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
