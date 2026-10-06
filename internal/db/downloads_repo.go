package db

import (
	"context"
	"database/sql"
	"errors"
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
	LastProcessedTimestamp      string
	TypeCount                   int
	CountToConsider             int
	SourceID                    int64
	SkipGeneralConsolidation    bool
	SkipGroupsConsolidation     bool
	SkipCategoriesConsolidation bool
}

type downloadScan struct {
	Frequency              string `db:"frequency"`
	Checksum               string `db:"checksum"`
	LastCheckedTimestamp   string `db:"last_checked_timestamp"`
	LastProcessedTimestamp string `db:"last_processed_timestamp"`
	URL                    string `db:"url"`
	Error                  string `db:"error"`
	Name                   string `db:"name"`
	Filepath               string `db:"filepath"`
	LastDownloadTimestamp  string `db:"last_download_timestamp"`
	SourceID               int64  `db:"source_id"`
	TypeCount              int    `db:"type_count"`
	CountToConsider        int    `db:"count_to_consider"`
	SkipGeneral            int    `db:"skip_general_consolidation"`
	SkipGroups             int    `db:"skip_groups_consolidation"`
	SkipCategories         int    `db:"skip_categories_consolidation"`
}

// UpsertDownload inserts or updates a download record.
func (r *DownloadsRepo) UpsertDownload(d DownloadRow) error {
	lastDownloadTimestamp := normalizeDownloadTimestampValue(d.LastDownloadTimestamp)
	lastCheckedTimestamp := normalizeDownloadTimestampValue(d.LastCheckedTimestamp)
	lastProcessedTimestamp := normalizeDownloadTimestampValue(d.LastProcessedTimestamp)

	table := constants.TableDownloads
	_, err := r.db.writeConn.Exec(`
		INSERT INTO `+table+` (source_id, url, filepath, frequency, checksum, error,
			last_download_timestamp, last_checked_timestamp, last_processed_timestamp,
			type_count, count_to_consider,
			skip_general_consolidation, skip_groups_consolidation, skip_categories_consolidation)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(source_id) DO UPDATE SET
			url = excluded.url,
			filepath = COALESCE(NULLIF(excluded.filepath, ''), `+table+`.filepath),
			frequency = excluded.frequency,
			checksum = COALESCE(NULLIF(excluded.checksum, ''), `+table+`.checksum),
			error = excluded.error,
			last_download_timestamp = COALESCE(excluded.last_download_timestamp, `+table+`.last_download_timestamp),
			last_checked_timestamp = COALESCE(excluded.last_checked_timestamp, `+table+`.last_checked_timestamp),
			last_processed_timestamp = COALESCE(excluded.last_processed_timestamp, `+table+`.last_processed_timestamp),
			type_count = excluded.type_count,
			count_to_consider = excluded.count_to_consider,
			skip_general_consolidation = excluded.skip_general_consolidation,
			skip_groups_consolidation = excluded.skip_groups_consolidation,
			skip_categories_consolidation = excluded.skip_categories_consolidation`,
		d.SourceID, d.URL, d.Filepath, d.Frequency, d.Checksum, d.Error,
		lastDownloadTimestamp, lastCheckedTimestamp, lastProcessedTimestamp,
		d.TypeCount, d.CountToConsider,
		boolToInt(d.SkipGeneralConsolidation),
		boolToInt(d.SkipGroupsConsolidation),
		boolToInt(d.SkipCategoriesConsolidation))
	if err != nil {
		return fmt.Errorf("upserting download for source %d: %w", d.SourceID, err)
	}
	return nil
}

// GetLatestDownloadSummary returns the persisted download fields used by the download skip decision
func (r *DownloadsRepo) GetLatestDownloadSummary(sourceName string) (*c.DownloadSummary, error) {
	var row struct {
		Frequency             string `db:"frequency"`
		LastDownloadTimestamp string `db:"last_download_timestamp"`
		URL                   string `db:"url"`
		Error                 string `db:"error"`
	}
	err := r.db.getRead(context.Background(), &row, `
		SELECT COALESCE(NULLIF(d.frequency, ''), s.frequency) AS frequency,
			COALESCE(d.last_download_timestamp, '') AS last_download_timestamp,
			COALESCE(NULLIF(d.url, ''), s.url) AS url,
			COALESCE(d.error, '') AS error
		FROM `+constants.TableDownloads+` d
		JOIN `+constants.TableSources+` s ON s.id = d.source_id
		WHERE s.name = ? LIMIT 1`, sourceName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("querying latest download summary for %s: %w", sourceName, err)
	}
	return &c.DownloadSummary{
		Frequency:             row.Frequency,
		LastDownloadTimestamp: row.LastDownloadTimestamp,
		URL:                   row.URL,
		Error:                 row.Error,
	}, nil
}

// GetLastProcessedChecksum returns the checksum that was last successfully processed for a source.
func (r *DownloadsRepo) GetLastProcessedChecksum(sourceID int64) string {
	var checksum string
	err := r.db.getRead(context.Background(), &checksum,
		"SELECT COALESCE(last_processed_checksum, '') FROM "+constants.TableDownloads+" WHERE source_id = ?",
		sourceID)
	if err != nil {
		return ""
	}
	return checksum
}

// SetLastProcessedChecksum updates the last_processed_checksum after successful processing.
func (r *DownloadsRepo) SetLastProcessedChecksum(sourceID int64, checksum string) error {
	_, err := r.db.writeConn.Exec(
		"UPDATE "+constants.TableDownloads+" SET last_processed_checksum = ? WHERE source_id = ?",
		checksum, sourceID)
	if err != nil {
		return fmt.Errorf("setting last_processed_checksum for source %d: %w", sourceID, err)
	}
	return nil
}

// SetLastProcessedTimestamp updates the last_processed_timestamp after successful processing.
func (r *DownloadsRepo) SetLastProcessedTimestamp(sourceID int64, ts string) error {
	normalized := normalizeDownloadTimestampValue(ts)
	_, err := r.db.writeConn.Exec(
		"UPDATE "+constants.TableDownloads+" SET last_processed_timestamp = ? WHERE source_id = ?",
		normalized, sourceID)
	if err != nil {
		return fmt.Errorf("setting last_processed_timestamp for source %d: %w", sourceID, err)
	}
	return nil
}

// nullableString maps the empty string to SQL NULL.
func nullableString(s string) any {
	if s == "" {
		return nil
	}
	return s
}

// GetDownloadSummaryBySourceName returns all persisted download summaries for a source name.
// Archive sources can produce multiple target summaries (one per source file in the archive).
func (r *DownloadsRepo) GetDownloadSummaryBySourceName(sourceName, downloadDir string) ([]c.DownloadSummary, error) {
	row, err := r.getDownloadSummaryRowByName(sourceName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []c.DownloadSummary{}, nil
		}
		return nil, err
	}

	return r.buildDownloadSummariesBatch([]downloadScan{row}, downloadDir)
}

var downloadSummaryBaseQuery = `
	SELECT s.id AS source_id, s.name AS name,
		COALESCE(NULLIF(d.url, ''), s.url) AS url,
		COALESCE(d.filepath, '') AS filepath,
		COALESCE(NULLIF(d.frequency, ''), s.frequency) AS frequency,
		COALESCE(d.checksum, '') AS checksum,
		COALESCE(d.error, '') AS error,
		COALESCE(d.last_download_timestamp, '') AS last_download_timestamp,
		COALESCE(d.last_checked_timestamp, '') AS last_checked_timestamp,
		COALESCE(d.last_processed_timestamp, '') AS last_processed_timestamp,
		d.type_count AS type_count,
		d.count_to_consider AS count_to_consider,
		d.skip_general_consolidation AS skip_general_consolidation,
		d.skip_groups_consolidation AS skip_groups_consolidation,
		d.skip_categories_consolidation AS skip_categories_consolidation
	FROM ` + constants.TableDownloads + ` d
	INNER JOIN ` + constants.TableSources + ` s ON s.id = d.source_id`

// ListDownloadSummaries returns all persisted download summaries.
func (r *DownloadsRepo) ListDownloadSummaries(downloadDir string) ([]c.DownloadSummary, error) {
	var scanned []downloadScan
	if err := r.db.selectRead(context.Background(), &scanned,
		downloadSummaryBaseQuery+" ORDER BY s.name"); err != nil {
		return nil, fmt.Errorf("querying download summaries: %w", err)
	}
	return r.buildDownloadSummariesBatch(scanned, downloadDir)
}

func (r *DownloadsRepo) getDownloadSummaryRowByName(sourceName string) (downloadScan, error) {
	var scanned downloadScan
	if err := r.db.getRead(context.Background(), &scanned,
		downloadSummaryBaseQuery+" WHERE s.name = ? ORDER BY s.id LIMIT 1",
		sourceName); err != nil {
		return downloadScan{}, err
	}
	return scanned, nil
}

func (r *DownloadsRepo) buildDownloadSummariesBatch(
	scanned []downloadScan,
	downloadDir string,
) ([]c.DownloadSummary, error) {
	if len(scanned) == 0 {
		return []c.DownloadSummary{}, nil
	}
	loader := newSourceMetadataLoader(r.db)
	sourceIDs := make([]int64, 0, len(scanned))
	for _, row := range scanned {
		sourceIDs = append(sourceIDs, row.SourceID)
	}
	typesBySource, err := loader.BatchSourceTypes(sourceIDs)
	if err != nil {
		return nil, err
	}
	categoriesBySource, err := loader.BatchSourceCategories(sourceIDs)
	if err != nil {
		return nil, err
	}
	filesBySource, err := loader.BatchSourceFiles(sourceIDs)
	if err != nil {
		return nil, err
	}

	summaries := make([]c.DownloadSummary, 0, len(scanned))
	for _, row := range scanned {
		types := typesBySource[row.SourceID]
		categories := categoriesBySource[row.SourceID]
		files := filesBySource[row.SourceID]
		base := c.DownloadSummary{
			Name:                        row.Name,
			URL:                         row.URL,
			Filepath:                    row.Filepath,
			Frequency:                   row.Frequency,
			Checksum:                    row.Checksum,
			Error:                       row.Error,
			LastDownloadTimestamp:       row.LastDownloadTimestamp,
			LastCheckedTimestamp:        row.LastCheckedTimestamp,
			LastProcessedTimestamp:      row.LastProcessedTimestamp,
			Types:                       types,
			Categories:                  categories,
			TypeCount:                   row.TypeCount,
			CountToConsider:             row.CountToConsider,
			SkipGeneralConsolidation:    row.SkipGeneral == 1,
			SkipGroupsConsolidation:     row.SkipGroups == 1,
			SkipCategoriesConsolidation: row.SkipCategories == 1,
		}

		if len(files) == 0 {
			if base.Filepath == "" {
				base.Filepath = filepath.Join(downloadDir, row.Name+".txt")
			}
			summaries = append(summaries, base)
			continue
		}

		for _, file := range files {
			summary := base
			summary.Filepath = filepath.Join(
				downloadDir,
				row.Name+"-"+strings.ReplaceAll(file, "/", "_"),
			)
			summaries = append(summaries, summary)
		}
	}

	return summaries, nil
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

	for _, layout := range constants.DownloadTimestampLayouts {
		if parsed, err := time.Parse(layout, ts); err == nil {
			return parsed.Format("2006-01-02 15:04:05")
		}
	}

	return ""
}
