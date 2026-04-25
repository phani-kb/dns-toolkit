package db

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
)

type DownloadsRepo struct {
	db *DB
}

func NewDownloadsRepo(db *DB) *DownloadsRepo {
	return &DownloadsRepo{db: db}
}

type DownloadRow struct {
	Checksum                    string
	URL                         string
	Filepath                    string
	Frequency                   string
	Error                       string
	LastDownloadTimestamp       string
	LastCheckedTimestamp        string
	TypeCount                   int
	CountToConsider             int
	SourceID                    int64
	SkipGeneralConsolidation    bool
	SkipGroupsConsolidation     bool
	SkipCategoriesConsolidation bool
}

type downloadSummaryRow struct {
	Frequency                   string
	Checksum                    string
	LastCheckedTimestamp        string
	URL                         string
	Error                       string
	Name                        string
	Filepath                    string
	LastDownloadTimestamp       string
	TypeCount                   int
	SourceID                    int64
	CountToConsider             int
	SkipCategoriesConsolidation bool
	SkipGroupsConsolidation     bool
	SkipGeneralConsolidation    bool
}

// UpsertDownload inserts or updates a download record.
func (r *DownloadsRepo) UpsertDownload(d DownloadRow) error {
	lastDownloadTimestamp := normalizeDownloadTimestampValue(d.LastDownloadTimestamp)
	lastCheckedTimestamp := normalizeDownloadTimestampValue(d.LastCheckedTimestamp)

	_, err := r.db.conn.Exec(`
		INSERT INTO `+constants.TableDownloads+` (source_id, url, filepath, frequency, checksum, error,
			last_download_timestamp, last_checked_timestamp, type_count, count_to_consider,
			skip_general_consolidation, skip_groups_consolidation, skip_categories_consolidation)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(source_id) DO UPDATE SET
			url = excluded.url,
			filepath = excluded.filepath,
			frequency = excluded.frequency,
			checksum = excluded.checksum,
			error = excluded.error,
			last_download_timestamp = excluded.last_download_timestamp,
			last_checked_timestamp = excluded.last_checked_timestamp,
			type_count = excluded.type_count,
			count_to_consider = excluded.count_to_consider,
			skip_general_consolidation = excluded.skip_general_consolidation,
			skip_groups_consolidation = excluded.skip_groups_consolidation,
			skip_categories_consolidation = excluded.skip_categories_consolidation`,
		d.SourceID, d.URL, d.Filepath, d.Frequency, d.Checksum, d.Error,
		lastDownloadTimestamp, lastCheckedTimestamp, d.TypeCount, d.CountToConsider,
		boolToInt(d.SkipGeneralConsolidation),
		boolToInt(d.SkipGroupsConsolidation),
		boolToInt(d.SkipCategoriesConsolidation))
	if err != nil {
		return fmt.Errorf("upserting download for source %d: %w", d.SourceID, err)
	}
	return nil
}

// GetDownloadBySourceID retrieves a download record by source ID.
func (r *DownloadsRepo) GetDownloadBySourceID(sourceID int64) (*DownloadRow, error) {
	row := r.db.conn.QueryRow(`
		SELECT source_id, url, filepath, frequency, checksum, error,
			last_download_timestamp, last_checked_timestamp, type_count, count_to_consider,
			skip_general_consolidation, skip_groups_consolidation, skip_categories_consolidation
		FROM `+constants.TableDownloads+` WHERE source_id = ?`, sourceID)

	d := &DownloadRow{}
	var skipGen, skipGrp, skipCat int
	err := row.Scan(&d.SourceID, &d.URL, &d.Filepath, &d.Frequency,
		&d.Checksum, &d.Error, &d.LastDownloadTimestamp, &d.LastCheckedTimestamp,
		&d.TypeCount, &d.CountToConsider, &skipGen, &skipGrp, &skipCat)
	if err != nil {
		return nil, err
	}
	d.SkipGeneralConsolidation = skipGen == 1
	d.SkipGroupsConsolidation = skipGrp == 1
	d.SkipCategoriesConsolidation = skipCat == 1
	return d, nil
}

// GetDownloadChecksum returns the stored checksum for a source, or empty if not found.
func (r *DownloadsRepo) GetDownloadChecksum(sourceID int64) string {
	var checksum string
	err := r.db.conn.QueryRow(
		"SELECT COALESCE(checksum, '') FROM "+constants.TableDownloads+" WHERE source_id = ?",
		sourceID).Scan(&checksum)
	if err != nil {
		return "" // Not found, return empty
	}
	return checksum
}

// IsDownloadUnchanged checks if a download's checksum matches the stored value.
// Returns true if the source hasn't changed (can reuse persisted entries).
func (r *DownloadsRepo) IsDownloadUnchanged(sourceID int64, newChecksum string) bool {
	stored := r.GetDownloadChecksum(sourceID)
	return stored != "" && stored == newChecksum
}

// GetDownloadSummaryBySourceName returns all persisted download summaries for a source name.
// Archive sources can produce multiple target summaries (one per source file in the archive).
func (r *DownloadsRepo) GetDownloadSummaryBySourceName(sourceName, downloadDir string) ([]c.DownloadSummary, error) {
	row, err := r.getDownloadSummaryRowByName(sourceName)
	if err != nil {
		if err == sql.ErrNoRows {
			return []c.DownloadSummary{}, nil
		}
		return nil, err
	}

	summaries, err := r.buildDownloadSummaries(row, downloadDir)
	if err != nil {
		return nil, err
	}

	return summaries, nil
}

// ListDownloadSummaries reconstructs download summaries from persisted source and download metadata.
func (r *DownloadsRepo) ListDownloadSummaries(downloadDir string) ([]c.DownloadSummary, error) {
	rows, err := r.db.conn.Query(`
		SELECT s.id, s.name,
			COALESCE(NULLIF(d.url, ''), s.url),
			COALESCE(d.filepath, ''),
			COALESCE(NULLIF(d.frequency, ''), s.frequency),
			COALESCE(d.checksum, ''),
			COALESCE(d.error, ''),
			COALESCE(d.last_download_timestamp, ''),
			COALESCE(d.last_checked_timestamp, ''),
			d.type_count,
			d.count_to_consider,
			d.skip_general_consolidation,
			d.skip_groups_consolidation,
			d.skip_categories_consolidation
		FROM ` + constants.TableDownloads + ` d
		INNER JOIN ` + constants.TableSources + ` s ON s.id = d.source_id
		ORDER BY s.name`)
	if err != nil {
		return nil, fmt.Errorf("querying download summaries: %w", err)
	}
	defer func() { _ = rows.Close() }() // nolint: errcheck

	summaries := make([]c.DownloadSummary, 0)
	for rows.Next() {
		row, scanErr := scanDownloadSummaryRow(rows)
		if scanErr != nil {
			return nil, scanErr
		}

		reconstructed, buildErr := r.buildDownloadSummaries(row, downloadDir)
		if buildErr != nil {
			return nil, buildErr
		}
		summaries = append(summaries, reconstructed...)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating download summaries: %w", err)
	}

	return summaries, nil
}

func scanDownloadSummaryRow(scanner interface{ Scan(dest ...any) error }) (downloadSummaryRow, error) {
	row := downloadSummaryRow{}
	var skipGen, skipGrp, skipCat int
	err := scanner.Scan(
		&row.SourceID,
		&row.Name,
		&row.URL,
		&row.Filepath,
		&row.Frequency,
		&row.Checksum,
		&row.Error,
		&row.LastDownloadTimestamp,
		&row.LastCheckedTimestamp,
		&row.TypeCount,
		&row.CountToConsider,
		&skipGen,
		&skipGrp,
		&skipCat,
	)
	if err != nil {
		return downloadSummaryRow{}, fmt.Errorf("scanning download summary row: %w", err)
	}

	row.SkipGeneralConsolidation = skipGen == 1
	row.SkipGroupsConsolidation = skipGrp == 1
	row.SkipCategoriesConsolidation = skipCat == 1

	return row, nil
}

func (r *DownloadsRepo) getDownloadSummaryRowByName(sourceName string) (downloadSummaryRow, error) {
	row := r.db.conn.QueryRow(`
		SELECT s.id, s.name,
			COALESCE(NULLIF(d.url, ''), s.url),
			COALESCE(d.filepath, ''),
			COALESCE(NULLIF(d.frequency, ''), s.frequency),
			COALESCE(d.checksum, ''),
			COALESCE(d.error, ''),
			COALESCE(d.last_download_timestamp, ''),
			COALESCE(d.last_checked_timestamp, ''),
			d.type_count,
			d.count_to_consider,
			d.skip_general_consolidation,
			d.skip_groups_consolidation,
			d.skip_categories_consolidation
		FROM `+constants.TableDownloads+` d
		INNER JOIN `+constants.TableSources+` s ON s.id = d.source_id
		WHERE s.name = ?
		ORDER BY s.id
		LIMIT 1`, sourceName)

	return scanDownloadSummaryRow(row)
}

func (r *DownloadsRepo) buildDownloadSummaries(
	row downloadSummaryRow,
	downloadDir string,
) ([]c.DownloadSummary, error) {
	types, err := r.getSourceTypes(row.SourceID)
	if err != nil {
		return nil, err
	}

	categories, err := r.getSourceCategories(row.SourceID)
	if err != nil {
		return nil, err
	}

	files, err := r.getSourceFiles(row.SourceID)
	if err != nil {
		return nil, err
	}

	base := c.DownloadSummary{
		Name:                        row.Name,
		URL:                         row.URL,
		Filepath:                    row.Filepath,
		Frequency:                   row.Frequency,
		Checksum:                    row.Checksum,
		Error:                       row.Error,
		LastDownloadTimestamp:       row.LastDownloadTimestamp,
		LastCheckedTimestamp:        row.LastCheckedTimestamp,
		Types:                       types,
		Categories:                  categories,
		TypeCount:                   row.TypeCount,
		CountToConsider:             row.CountToConsider,
		SkipGeneralConsolidation:    row.SkipGeneralConsolidation,
		SkipGroupsConsolidation:     row.SkipGroupsConsolidation,
		SkipCategoriesConsolidation: row.SkipCategoriesConsolidation,
	}

	if len(files) == 0 {
		if base.Filepath == "" {
			base.Filepath = filepath.Join(downloadDir, row.Name+".txt")
		}
		return []c.DownloadSummary{base}, nil
	}

	summaries := make([]c.DownloadSummary, 0, len(files))
	for _, file := range files {
		summary := base
		summary.Filepath = filepath.Join(
			downloadDir,
			row.Name+"-"+strings.ReplaceAll(file, "/", "_"),
		)
		summaries = append(summaries, summary)
	}

	return summaries, nil
}

func (r *DownloadsRepo) getSourceCategories(sourceID int64) ([]string, error) {
	rows, err := r.db.conn.Query(`
		SELECT cn.name
		FROM `+constants.TableSourceCategories+` sc
		INNER JOIN `+constants.TableCategoryNames+` cn ON cn.id = sc.category_name_id
		WHERE sc.source_id = ?
		ORDER BY cn.name`,
		sourceID,
	)
	if err != nil {
		return nil, fmt.Errorf("querying source categories for %d: %w", sourceID, err)
	}
	defer func() { _ = rows.Close() }() // nolint: errcheck

	categories := make([]string, 0)
	for rows.Next() {
		var category string
		if err := rows.Scan(&category); err != nil {
			return nil, fmt.Errorf("scanning source category for %d: %w", sourceID, err)
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating source categories for %d: %w", sourceID, err)
	}

	return categories, nil
}

func (r *DownloadsRepo) getSourceFiles(sourceID int64) ([]string, error) {
	rows, err := r.db.conn.Query(
		"SELECT filename FROM "+constants.TableSourceFiles+" WHERE source_id = ? ORDER BY filename",
		sourceID,
	)
	if err != nil {
		return nil, fmt.Errorf("querying source files for %d: %w", sourceID, err)
	}
	defer func() { _ = rows.Close() }() // nolint: errcheck

	files := make([]string, 0)
	for rows.Next() {
		var file string
		if err := rows.Scan(&file); err != nil {
			return nil, fmt.Errorf("scanning source file for %d: %w", sourceID, err)
		}
		files = append(files, file)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating source files for %d: %w", sourceID, err)
	}

	return files, nil
}

func (r *DownloadsRepo) getSourceTypes(sourceID int64) ([]c.SourceType, error) {
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
		var typeID int64
		var sourceType c.SourceType
		var disabled int
		if err := rows.Scan(&typeID, &sourceType.Name, &sourceType.Notes, &disabled); err != nil {
			return nil, fmt.Errorf("scanning source type for %d: %w", sourceID, err)
		}
		sourceType.Disabled = disabled == 1

		listTypes, err := r.getSourceListTypes(typeID)
		if err != nil {
			return nil, err
		}
		sourceType.ListTypes = listTypes
		types = append(types, sourceType)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating source types for %d: %w", sourceID, err)
	}

	return types, nil
}

func (r *DownloadsRepo) getSourceListTypes(sourceTypeID int64) ([]c.ListType, error) {
	rows, err := r.db.conn.Query(`
		SELECT slt.id, ltn.name, COALESCE(sltn.notes, ''), slt.disabled, slt.must_consider
		FROM `+constants.TableSourceListTypes+` slt
		INNER JOIN `+constants.TableListTypeNames+` ltn ON ltn.id = slt.list_type_name_id
		LEFT JOIN `+constants.TableSourceListTypeNotes+` sltn ON sltn.source_list_type_id = slt.id
		WHERE slt.source_type_id = ?
		ORDER BY slt.id`, sourceTypeID)
	if err != nil {
		return nil, fmt.Errorf("querying list types for source type %d: %w", sourceTypeID, err)
	}
	defer func() { _ = rows.Close() }() // nolint: errcheck

	listTypes := make([]c.ListType, 0)
	for rows.Next() {
		var listTypeID int64
		var listType c.ListType
		var disabled, mustConsider int
		if err := rows.Scan(&listTypeID, &listType.Name, &listType.Notes, &disabled, &mustConsider); err != nil {
			return nil, fmt.Errorf("scanning list type for source type %d: %w", sourceTypeID, err)
		}
		listType.Disabled = disabled == 1
		listType.MustConsider = mustConsider == 1

		groups, err := r.getSourceListTypeGroups(listTypeID)
		if err != nil {
			return nil, err
		}
		listType.Groups = groups
		listTypes = append(listTypes, listType)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating list types for source type %d: %w", sourceTypeID, err)
	}

	return listTypes, nil
}

func (r *DownloadsRepo) getSourceListTypeGroups(sourceListTypeID int64) ([]string, error) {
	rows, err := r.db.conn.Query(`
		SELECT gn.name
		FROM `+constants.TableSourceListTypeGroups+` sltg
		INNER JOIN `+constants.TableGroupNames+` gn ON gn.id = sltg.group_name_id
		WHERE sltg.source_list_type_id = ?
		ORDER BY gn.name`,
		sourceListTypeID,
	)
	if err != nil {
		return nil, fmt.Errorf("querying list type groups for %d: %w", sourceListTypeID, err)
	}
	defer func() { _ = rows.Close() }() // nolint: errcheck

	groups := make([]string, 0)
	for rows.Next() {
		var group string
		if err := rows.Scan(&group); err != nil {
			return nil, fmt.Errorf("scanning list type group for %d: %w", sourceListTypeID, err)
		}
		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating list type groups for %d: %w", sourceListTypeID, err)
	}

	return groups, nil
}

func normalizeDownloadTimestampValue(ts string) any {
	normalized := normalizeDownloadTimestamp(ts)
	if normalized == "" {
		return nil
	}

	return normalized
}

func normalizeDownloadTimestamp(ts string) string {
	if ts == "" {
		return ""
	}

	if parsed, err := time.Parse(constants.TimestampFormat, ts); err == nil {
		return parsed.Format("2006-01-02 15:04:05")
	}

	for _, layout := range []string{time.RFC3339, time.RFC3339Nano, "2006-01-02 15:04:05"} {
		if parsed, err := time.Parse(layout, ts); err == nil {
			return parsed.Format("2006-01-02 15:04:05")
		}
	}

	return ""
}
