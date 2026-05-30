package db

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
)

type ProcessedRepo struct {
	db *DB
}

func NewProcessedRepo(db *DB) *ProcessedRepo {
	return &ProcessedRepo{db: db}
}

type processedAggregateRow struct {
	Name              string
	GenericSourceType string
	ActualSourceType  string
	ListType          string
	SourceID          int64
	EntryCount        int
	Valid             bool
	MustConsider      bool
}

// ListProcessedSummaries reconstructs processed summaries from persisted entries metadata.
func (r *ProcessedRepo) ListProcessedSummaries(processedDir string) ([]c.ProcessedSummary, error) {
	rows, err := r.db.conn.Query(`
		SELECT s.id, s.name,
			e.generic_source_type,
			e.actual_source_type,
			e.list_type,
			e.valid,
			MAX(e.must_consider) as must_consider,
			COUNT(*) as entry_count
		FROM ` + constants.TableEntries + ` e
		INNER JOIN ` + constants.TableSources + ` s ON s.id = e.source_id
		GROUP BY s.id, s.name, e.generic_source_type, e.actual_source_type, e.list_type, e.valid
		ORDER BY s.name, e.actual_source_type, e.list_type, e.valid DESC`)
	if err != nil {
		return nil, fmt.Errorf("querying processed summaries: %w", err)
	}
	defer func() { _ = rows.Close() }() // nolint: errcheck

	summaryBySourceID := make(map[int64]*c.ProcessedSummary)
	orderedSourceIDs := make([]int64, 0)

	for rows.Next() {
		agg, scanErr := scanProcessedAggregateRow(rows)
		if scanErr != nil {
			return nil, scanErr
		}

		summary, exists := summaryBySourceID[agg.SourceID]
		if !exists {
			types, typesErr := r.getSourceTypes(agg.SourceID)
			if typesErr != nil {
				return nil, typesErr
			}
			summary = &c.ProcessedSummary{
				Name:  agg.Name,
				Types: types,
			}
			summaryBySourceID[agg.SourceID] = summary
			orderedSourceIDs = append(orderedSourceIDs, agg.SourceID)
		}

		groups, groupsErr := r.getEntryGroups(agg.SourceID, agg.ActualSourceType, agg.ListType)
		if groupsErr != nil {
			return nil, groupsErr
		}
		categories, categoriesErr := r.getEntryCategories(agg.SourceID, agg.ActualSourceType, agg.ListType)
		if categoriesErr != nil {
			return nil, categoriesErr
		}

		file := c.ProcessedFile{
			Name:              agg.Name,
			GenericSourceType: agg.GenericSourceType,
			ActualSourceType:  agg.ActualSourceType,
			ListType:          agg.ListType,
			Filepath: buildProcessedFilePath(
				processedDir,
				agg.Name,
				agg.ActualSourceType,
				agg.ListType,
				agg.Valid,
			),
			NumberOfEntries: agg.EntryCount,
			MustConsider:    agg.MustConsider,
			Valid:           agg.Valid,
			Groups:          groups,
			Categories:      categories,
		}

		if agg.Valid {
			summary.ValidFiles = append(summary.ValidFiles, file)
		} else {
			summary.InvalidFiles = append(summary.InvalidFiles, file)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating processed summary rows: %w", err)
	}

	summaries := make([]c.ProcessedSummary, 0, len(summaryBySourceID))
	for _, sourceID := range orderedSourceIDs {
		summary := summaryBySourceID[sourceID]
		sort.Slice(summary.ValidFiles, func(i, j int) bool {
			return strings.ToLower(summary.ValidFiles[i].Filepath) < strings.ToLower(summary.ValidFiles[j].Filepath)
		})
		sort.Slice(summary.InvalidFiles, func(i, j int) bool {
			return strings.ToLower(summary.InvalidFiles[i].Filepath) < strings.ToLower(summary.InvalidFiles[j].Filepath)
		})
		summaries = append(summaries, *summary)
	}

	sort.Slice(summaries, func(i, j int) bool {
		return strings.ToLower(summaries[i].Name) < strings.ToLower(summaries[j].Name)
	})

	return summaries, nil
}

func scanProcessedAggregateRow(scanner interface{ Scan(dest ...any) error }) (processedAggregateRow, error) {
	row := processedAggregateRow{}
	var validInt, mustConsiderInt int
	err := scanner.Scan(
		&row.SourceID,
		&row.Name,
		&row.GenericSourceType,
		&row.ActualSourceType,
		&row.ListType,
		&validInt,
		&mustConsiderInt,
		&row.EntryCount,
	)
	if err != nil {
		return processedAggregateRow{}, fmt.Errorf("scanning processed aggregate row: %w", err)
	}

	row.Valid = validInt == 1
	row.MustConsider = mustConsiderInt == 1

	return row, nil
}

func (r *ProcessedRepo) getSourceTypes(sourceID int64) ([]c.SourceType, error) {
	rows, err := r.db.conn.Query(`
		SELECT st.id, tn.name, COALESCE(st.notes, ''), st.disabled
		FROM `+constants.TableSourceTypes+` st
		INNER JOIN `+constants.TableTypeNames+` tn ON tn.id = st.type_name_id
		WHERE st.source_id = ?
		ORDER BY st.id`, sourceID)
	if err != nil {
		return nil, fmt.Errorf("querying source types for %d: %w", sourceID, err)
	}
	defer func() { _ = rows.Close() }() // nolint: errcheck

	types := make([]c.SourceType, 0)
	for rows.Next() {
		var sourceType c.SourceType
		var sourceTypeID int64
		var disabled int
		if err := rows.Scan(&sourceTypeID, &sourceType.Name, &sourceType.Notes, &disabled); err != nil {
			return nil, fmt.Errorf("scanning source type for %d: %w", sourceID, err)
		}
		sourceType.Disabled = disabled == 1

		listTypes, listTypeErr := r.getSourceListTypes(sourceTypeID)
		if listTypeErr != nil {
			return nil, listTypeErr
		}
		sourceType.ListTypes = listTypes
		types = append(types, sourceType)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating source types for %d: %w", sourceID, err)
	}

	return types, nil
}

func (r *ProcessedRepo) getSourceListTypes(sourceTypeID int64) ([]c.ListType, error) {
	rows, err := r.db.conn.Query(`
		SELECT slt.id, ltn.name, COALESCE(sltn.notes, ''), slt.disabled, slt.must_consider
		FROM `+constants.TableSourceListTypes+` slt
		INNER JOIN `+constants.TableListTypeNames+` ltn ON ltn.id = slt.list_type_name_id
		LEFT JOIN `+constants.TableSourceListTypeNotes+` sltn ON sltn.source_list_type_id = slt.id
		WHERE slt.source_type_id = ?
		ORDER BY slt.id`, sourceTypeID)
	if err != nil {
		return nil, fmt.Errorf("querying source list types for %d: %w", sourceTypeID, err)
	}
	defer func() { _ = rows.Close() }() // nolint: errcheck

	listTypes := make([]c.ListType, 0)
	for rows.Next() {
		var listType c.ListType
		var listTypeID int64
		var disabled, mustConsider int
		if err := rows.Scan(&listTypeID, &listType.Name, &listType.Notes, &disabled, &mustConsider); err != nil {
			return nil, fmt.Errorf("scanning source list type for %d: %w", sourceTypeID, err)
		}
		listType.Disabled = disabled == 1
		listType.MustConsider = mustConsider == 1

		groups, groupsErr := r.getSourceListTypeGroups(listTypeID)
		if groupsErr != nil {
			return nil, groupsErr
		}
		listType.Groups = groups
		listTypes = append(listTypes, listType)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating source list types for %d: %w", sourceTypeID, err)
	}

	return listTypes, nil
}

func (r *ProcessedRepo) getSourceListTypeGroups(sourceListTypeID int64) ([]string, error) {
	rows, err := r.db.conn.Query(`
		SELECT gn.name
		FROM `+constants.TableSourceListTypeGroups+` sltg
		INNER JOIN `+constants.TableGroupNames+` gn ON gn.id = sltg.group_name_id
		WHERE sltg.source_list_type_id = ?
		ORDER BY gn.name`, sourceListTypeID)
	if err != nil {
		return nil, fmt.Errorf("querying source list type groups for %d: %w", sourceListTypeID, err)
	}
	defer rows.Close() // nolint: errcheck

	groups := make([]string, 0)
	for rows.Next() {
		var group string
		if err := rows.Scan(&group); err != nil {
			return nil, fmt.Errorf("scanning source list type group for %d: %w", sourceListTypeID, err)
		}
		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating source list type groups for %d: %w", sourceListTypeID, err)
	}

	return groups, nil
}

func (r *ProcessedRepo) getEntryGroups(sourceID int64, sourceType, listType string) ([]string, error) {
	query := "SELECT group_name FROM " + constants.TableEntryGroups +
		" WHERE source_id = ? AND source_type = ? AND list_type = ? ORDER BY group_name"
	rows, err := r.db.conn.Query(query, sourceID, sourceType, listType)
	if err != nil {
		return nil, fmt.Errorf("querying entry groups for source %d: %w", sourceID, err)
	}
	defer rows.Close() // nolint: errcheck

	groups := make([]string, 0)
	for rows.Next() {
		var group string
		if err := rows.Scan(&group); err != nil {
			return nil, fmt.Errorf("scanning entry group for source %d: %w", sourceID, err)
		}
		groups = append(groups, group)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating entry groups for source %d: %w", sourceID, err)
	}

	return groups, nil
}

func (r *ProcessedRepo) getEntryCategories(sourceID int64, sourceType, listType string) ([]string, error) {
	query := "SELECT category FROM " + constants.TableEntryCategories +
		" WHERE source_id = ? AND source_type = ? AND list_type = ? ORDER BY category"
	rows, err := r.db.conn.Query(query, sourceID, sourceType, listType)
	if err != nil {
		return nil, fmt.Errorf("querying entry categories for source %d: %w", sourceID, err)
	}
	defer rows.Close() // nolint: errcheck

	categories := make([]string, 0)
	for rows.Next() {
		var category string
		if err := rows.Scan(&category); err != nil {
			return nil, fmt.Errorf("scanning entry category for source %d: %w", sourceID, err)
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating entry categories for source %d: %w", sourceID, err)
	}

	return categories, nil
}

func buildProcessedFilePath(processedDir, sourceName, sourceType, listType string, valid bool) string {
	entryType := "valid"
	if !valid {
		entryType = "invalid"
	}

	hash := md5.Sum([]byte(sourceName + sourceType + listType + entryType))
	listTypeShort := constants.ListTypeMap[strings.ToLower(listType)]
	if listTypeShort == "" {
		listTypeShort = listType
	}

	fileName := fmt.Sprintf(
		"%s_%s_%s_%s_%s.txt",
		sourceName,
		sourceType,
		listTypeShort,
		entryType,
		hex.EncodeToString(hash[:]),
	)

	return filepath.Join(processedDir, fileName)
}
