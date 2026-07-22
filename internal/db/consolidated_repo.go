package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
)

// ConsolidatedRepo provides operations for consolidation results.
type ConsolidatedRepo struct {
	db *DB
}

type ResolvedAllowEntry struct {
	GenericSourceType string
	Entry             string
	MustConsider      bool
}

type BlocklistConsolidationResult struct {
	OriginalCount int64 // before allow filtering
	FinalCount    int64 // rows remaining after allow filtering
	SourceCount   int64 // distinct sources
}

type ConsolidatedEntryRow struct {
	Entry             string
	GenericSourceType string
	ListType          string
	ConsolidationType string // general, group, category
	GroupName         string // non-empty when ConsolidationType == group
	Category          string // non-empty when ConsolidationType == category
	Valid             bool
	SourceCount       int
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

type scopedConsolidationParams struct {
	ConsolidationType string // group or category
	JoinTable         string // TableEntryGroups or TableEntryCategories
	ScopeColumn       string // group_name or category
	SkipColumn        string // skip_groups_consolidation or skip_categories_consolidation
	ScopeValue        string // the group or category value
	GenericSourceType string
}

type scopedConsolidationResult struct {
	ScopeValue  string
	SourceNames []string
	SourceCount int64
	FinalCount  int64
}

type scopeConfig struct {
	joinTable   string
	scopeColumn string
	skipColumn  string
}

func NewConsolidatedRepo(db *DB) *ConsolidatedRepo {
	return &ConsolidatedRepo{db: db}
}

// ComputeConsolidationFingerprint computes a hash of all processed checksums for enabled sources
// relevant to the given consolidation type
func (r *ConsolidatedRepo) ComputeConsolidationFingerprint(consolidationType string) (string, error) {
	skipCol := "skip_general_consolidation"
	switch consolidationType {
	case constants.ConsolidationTypeGroup:
		skipCol = "skip_groups_consolidation"
	case constants.ConsolidationTypeCategory:
		skipCol = "skip_categories_consolidation"
	}

	rows, err := r.db.readConn.Query(`
		SELECT COALESCE(d.last_processed_checksum, '')
		FROM ` + constants.TableDownloads + ` d
		JOIN ` + constants.TableSources + ` s ON s.id = d.source_id
		WHERE s.disabled = 0 AND s.` + skipCol + ` = 0
		ORDER BY s.id`,
	)
	if err != nil {
		return "", fmt.Errorf("querying source checksums for fingerprint: %w", err)
	}
	defer rows.Close() // nolint: errcheck

	var sb strings.Builder
	for rows.Next() {
		var checksum string
		if err := rows.Scan(&checksum); err != nil {
			return "", err
		}
		sb.WriteString(checksum)
		sb.WriteByte('\n')
	}
	if err := rows.Err(); err != nil {
		return "", err
	}

	hash := u.CalculateChecksumFromContent([]byte(sb.String()), "md5")
	return hash, nil
}

// GetStoredFingerprint returns the stored consolidation fingerprint for a given type.
func (r *ConsolidatedRepo) GetStoredFingerprint(consolidationType string) string {
	var fp string
	err := r.db.readConn.QueryRow(
		"select fingerprint from dnstk_consolidation_state where consolidation_type = ? and generic_source_type = ''",
		consolidationType).Scan(&fp)
	if err != nil {
		return ""
	}
	return fp
}

// GetStoredTypeFingerprint returns the stored fingerprint for a specific source type.
func (r *ConsolidatedRepo) GetStoredTypeFingerprint(consolidationType, genericSourceType string) string {
	var fp string
	err := r.db.readConn.QueryRow(
		"select fingerprint from dnstk_consolidation_state where consolidation_type = ? and generic_source_type = ?",
		consolidationType, genericSourceType).Scan(&fp)
	if err != nil {
		return ""
	}
	return fp
}

// SetStoredFingerprint saves the consolidation fingerprint after a successful run.
func (r *ConsolidatedRepo) SetStoredFingerprint(consolidationType, fingerprint string) error {
	_, err := r.db.writeConn.Exec(`
		insert into dnstk_consolidation_state (consolidation_type, generic_source_type, fingerprint, last_consolidated_at)
		values (?, '', ?, datetime('now'))
		on conflict(consolidation_type, generic_source_type) do update set
			fingerprint = excluded.fingerprint,
			last_consolidated_at = excluded.last_consolidated_at`,
		consolidationType, fingerprint)
	return err
}

// SetStoredTypeFingerprint saves the fingerprint for a specific source type.
func (r *ConsolidatedRepo) SetStoredTypeFingerprint(consolidationType, genericSourceType, fingerprint string) error {
	_, err := r.db.writeConn.Exec(`
		insert into dnstk_consolidation_state (consolidation_type, generic_source_type, fingerprint, last_consolidated_at)
		values (?, ?, ?, datetime('now'))
		on conflict(consolidation_type, generic_source_type) do update set
			fingerprint = excluded.fingerprint,
			last_consolidated_at = excluded.last_consolidated_at`,
		consolidationType, genericSourceType, fingerprint)
	return err
}

// HasConsolidatedData returns true if consolidated entries exist for this type.
func (r *ConsolidatedRepo) HasConsolidatedData(consolidationType string) bool {
	var count int
	err := r.db.readConn.QueryRow(
		"SELECT COUNT(1) FROM "+constants.TableConsolidatedEntries+" WHERE consolidation_type = ? LIMIT 1",
		consolidationType).Scan(&count)
	return err == nil && count > 0
}

// HasConsolidatedDataForType returns true if consolidated entries exist for this
// consolidation_type + generic_source_type combination.
func (r *ConsolidatedRepo) HasConsolidatedDataForType(consolidationType, genericSourceType string) bool {
	var count int
	err := r.db.readConn.QueryRow(
		"SELECT 1 FROM "+constants.TableConsolidatedEntries+
			" WHERE consolidation_type = ? AND generic_source_type = ? LIMIT 1",
		consolidationType, genericSourceType).Scan(&count)
	return err == nil && count > 0
}

// ComputeTypeFingerprint computes a fingerprint for a specific generic_source_type
// within a consolidation type.
func (r *ConsolidatedRepo) ComputeTypeFingerprint(consolidationType, genericSourceType string) (string, error) {
	skipCol := "skip_general_consolidation"
	switch consolidationType {
	case constants.ConsolidationTypeGroup:
		skipCol = "skip_groups_consolidation"
	case constants.ConsolidationTypeCategory:
		skipCol = "skip_categories_consolidation"
	}

	rows, err := r.db.readConn.Query(`
		SELECT COALESCE(d.last_processed_checksum, '')
		FROM `+constants.TableDownloads+` d
		JOIN `+constants.TableSources+` s ON s.id = d.source_id
		JOIN `+constants.TableEntries+` e ON e.source_id = s.id
		WHERE s.disabled = 0 AND s.`+skipCol+` = 0
			AND e.generic_source_type = ?
		GROUP BY s.id
		ORDER BY s.id`, genericSourceType,
	)
	if err != nil {
		return "", fmt.Errorf("querying type fingerprint for %s/%s: %w", consolidationType, genericSourceType, err)
	}
	defer rows.Close() // nolint: errcheck

	var sb strings.Builder
	sb.WriteString(genericSourceType)
	sb.WriteByte('\n')
	for rows.Next() {
		var checksum string
		if err := rows.Scan(&checksum); err != nil {
			return "", err
		}
		sb.WriteString(checksum)
		sb.WriteByte('\n')
	}
	if err := rows.Err(); err != nil {
		return "", err
	}

	hash := u.CalculateChecksumFromContent([]byte(sb.String()), "md5")
	return hash, nil
}

// ClearConsolidatedRowsForType deletes consolidated entries for a specific source type.
func (r *ConsolidatedRepo) ClearConsolidatedRowsForType(
	ctx context.Context, consolidationType, genericSourceType string,
) error {
	_, err := r.db.writeConn.ExecContext(ctx,
		"DELETE FROM "+constants.TableConsolidatedEntries+
			" WHERE consolidation_type = ? AND generic_source_type = ?",
		consolidationType, genericSourceType)
	return err
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
	var args []any

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
	if _, err := r.db.writeConn.Exec("drop index if exists idx_consolidated_type_valid_entry"); err != nil {
		return fmt.Errorf("dropping idx_consolidated_type_valid_entry: %w", err)
	}
	if _, err := r.db.writeConn.Exec("drop index if exists idx_consolidated_type_valid_group_entry"); err != nil {
		return fmt.Errorf("dropping idx_consolidated_type_valid_group_entry: %w", err)
	}
	if _, err := r.db.writeConn.Exec("drop index if exists idx_consolidated_type_valid_category_entry"); err != nil {
		return fmt.Errorf("dropping idx_consolidated_type_valid_category_entry: %w", err)
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
	if _, err := r.db.writeConn.Exec(
		`CREATE INDEX IF NOT EXISTS idx_consolidated_type_valid_entry ON ` +
			constants.TableConsolidatedEntries +
			` (consolidation_type, generic_source_type, list_type, valid, entry)`,
	); err != nil {
		return fmt.Errorf("recreating idx_consolidated_type_valid_entry: %w", err)
	}
	if _, err := r.db.writeConn.Exec(
		`CREATE INDEX IF NOT EXISTS idx_consolidated_type_valid_group_entry ON ` +
			constants.TableConsolidatedEntries +
			` (consolidation_type, generic_source_type, list_type, valid, group_name, entry)`,
	); err != nil {
		return fmt.Errorf("recreating idx_consolidated_type_valid_group_entry: %w", err)
	}
	if _, err := r.db.writeConn.Exec(
		`CREATE INDEX IF NOT EXISTS idx_consolidated_type_valid_category_entry ON ` +
			constants.TableConsolidatedEntries +
			` (consolidation_type, generic_source_type, list_type, valid, category, entry)`,
	); err != nil {
		return fmt.Errorf("recreating idx_consolidated_type_valid_category_entry: %w", err)
	}

	return nil
}

func (r *ConsolidatedRepo) ClearAllConsolidated(ctx context.Context) error {
	return r.db.DropAndRecreateTable(ctx, constants.TableConsolidatedEntries)
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

func (r *ConsolidatedRepo) BulkInsertEntries(ctx context.Context, rows []ConsolidatedEntryRow) (int64, error) {
	if len(rows) == 0 {
		return 0, nil
	}

	const batchSize = constants.BulkInsertBatchSize
	var inserted int64

	err := r.db.InBulkWriteTransaction(ctx, func(tx *sql.Tx) error {
		for i := 0; i < len(rows); i += batchSize {
			end := min(i+batchSize, len(rows))
			batch := rows[i:end]

			var query strings.Builder
			query.WriteString(`INSERT INTO ` + constants.TableConsolidatedEntries +
				` (entry, generic_source_type, list_type, consolidation_type, group_name, category, valid, source_count) VALUES `)
			args := make([]any, 0, len(batch)*8)
			for j, row := range batch {
				if j > 0 {
					query.WriteString(",")
				}
				query.WriteString("(?,?,?,?,?,?,?,?)")

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

			result, execErr := tx.ExecContext(ctx, query.String(), args...)
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

func (r *ConsolidatedRepo) ClearConsolidatedRows(ctx context.Context, consolidationType string) error {
	_, err := r.db.writeConn.ExecContext(ctx,
		"DELETE FROM "+constants.TableConsolidatedEntries+" WHERE consolidation_type = ?",
		consolidationType)
	return err
}

func (r *ConsolidatedRepo) DropConsolidatedIndexes(ctx context.Context) error {
	for _, n := range []string{
		"idx_consolidated_lookup",
		"idx_consolidated_type",
		"idx_consolidated_type_valid_entry",
		"idx_consolidated_type_valid_group_entry",
		"idx_consolidated_type_valid_category_entry",
	} {
		if _, err := r.db.writeConn.ExecContext(ctx, "DROP INDEX IF EXISTS "+n); err != nil {
			return fmt.Errorf("dropping %s: %w", n, err)
		}
	}
	return nil
}

func (r *ConsolidatedRepo) CreateConsolidatedIndexes(ctx context.Context) error {
	stmts := []string{
		`CREATE INDEX IF NOT EXISTS idx_consolidated_lookup ON ` + constants.TableConsolidatedEntries +
			` (entry, generic_source_type, list_type, consolidation_type)`,
		`CREATE INDEX IF NOT EXISTS idx_consolidated_type ON ` + constants.TableConsolidatedEntries +
			` (consolidation_type, generic_source_type, list_type)`,
		`CREATE INDEX IF NOT EXISTS idx_consolidated_type_valid_entry ON ` + constants.TableConsolidatedEntries +
			` (consolidation_type, generic_source_type, list_type, valid, entry)`,
		`CREATE INDEX IF NOT EXISTS idx_consolidated_type_valid_group_entry ON ` + constants.TableConsolidatedEntries +
			` (consolidation_type, generic_source_type, list_type, valid, group_name, entry)`,
		`CREATE INDEX IF NOT EXISTS idx_consolidated_type_valid_category_entry ON ` + constants.TableConsolidatedEntries +
			` (consolidation_type, generic_source_type, list_type, valid, category, entry)`,
	}
	for _, s := range stmts {
		if _, err := r.db.writeConn.ExecContext(ctx, s); err != nil {
			return fmt.Errorf("recreating consolidated index: %w", err)
		}
	}
	return nil
}

// LoadResolvedAllowSet clears the helper table and loads the resolved allow set.
func (r *ConsolidatedRepo) LoadResolvedAllowSet(ctx context.Context, entries []ResolvedAllowEntry) error {
	if _, err := r.db.writeConn.ExecContext(ctx, "DELETE FROM "+constants.TableResolvedAllow); err != nil {
		return fmt.Errorf("clearing %s: %w", constants.TableResolvedAllow, err)
	}
	if len(entries) == 0 {
		return nil
	}

	const cols = 3
	batchSize := constants.BulkInsertBatchSize

	return r.db.InTransaction(ctx, func(tx *sql.Tx) error {
		for i := 0; i < len(entries); i += batchSize {
			end := min(i+batchSize, len(entries))
			batch := entries[i:end]

			var q strings.Builder
			q.WriteString("INSERT OR IGNORE INTO " + constants.TableResolvedAllow +
				" (generic_source_type, entry, must_consider) VALUES ")
			args := make([]any, 0, len(batch)*cols)
			for j, e := range batch {
				if j > 0 {
					q.WriteByte(',')
				}
				q.WriteString("(?,?,?)")
				args = append(args, e.GenericSourceType, e.Entry, boolToInt(e.MustConsider))
			}
			if _, err := tx.ExecContext(ctx, q.String(), args...); err != nil {
				return fmt.Errorf("inserting resolved allow batch: %w", err)
			}
		}
		return nil
	})
}

func (r *ConsolidatedRepo) ConsolidateBlocklistGeneral(
	ctx context.Context,
	genericSourceType string,
	valid bool,
) (BlocklistConsolidationResult, error) {
	var res BlocklistConsolidationResult

	if err := r.db.readConn.QueryRowContext(ctx, `
		SELECT COUNT(DISTINCT e.entry), COUNT(DISTINCT e.source_id)
		FROM `+constants.TableEntries+` e
		JOIN `+constants.TableSources+` s ON s.id = e.source_id
		WHERE e.generic_source_type = ?
			AND e.list_type = 'blocklist'
			AND e.valid = ?
			AND s.disabled = 0
			AND s.skip_general_consolidation = 0`,
		genericSourceType, boolToInt(valid),
	).Scan(&res.OriginalCount, &res.SourceCount); err != nil {
		return res, fmt.Errorf("counting blocklist entries for %s: %w", genericSourceType, err)
	}
	if res.OriginalCount == 0 {
		return res, nil
	}

	err := r.db.InTransaction(ctx, func(tx *sql.Tx) error {
		result, execErr := tx.ExecContext(ctx, `
			INSERT INTO `+constants.TableConsolidatedEntries+`
				(entry, generic_source_type, list_type, consolidation_type,
				 group_name, category, valid, source_count)
			SELECT e.entry, e.generic_source_type, 'blocklist', 'general',
				NULL, NULL, ?, COUNT(DISTINCT e.source_id)
			FROM `+constants.TableEntries+` e
			JOIN `+constants.TableSources+` s ON s.id = e.source_id
			LEFT JOIN `+constants.TableResolvedAllow+` ra
				ON ra.generic_source_type = e.generic_source_type
				AND ra.entry = e.entry
			WHERE e.generic_source_type = ?
				AND e.list_type = 'blocklist'
				AND e.valid = ?
				AND s.disabled = 0
				AND s.skip_general_consolidation = 0
			GROUP BY e.entry
			HAVING MAX(CASE WHEN ra.entry IS NOT NULL THEN 1 ELSE 0 END) = 0
			    OR (MAX(e.must_consider) = 1 AND COALESCE(MAX(ra.must_consider), 0) = 0)`,
			boolToInt(valid), genericSourceType, boolToInt(valid))
		if execErr != nil {
			return fmt.Errorf("consolidating blocklist %s: %w", genericSourceType, execErr)
		}
		affected, _ := result.RowsAffected() // nolint: errcheck
		res.FinalCount = affected
		return nil
	})
	return res, err
}

func scopedTargetValues(p scopedConsolidationParams) (groupName, category *string) {
	v := p.ScopeValue
	switch p.ConsolidationType {
	case constants.ConsolidationTypeGroup:
		return &v, nil
	case constants.ConsolidationTypeCategory:
		return nil, &v
	}
	return nil, nil
}

// ConsolidateScopedAllowlist consolidates allowlist entries for one scope value (group or category).
func (r *ConsolidatedRepo) ConsolidateScopedAllowlist(
	ctx context.Context,
	p scopedConsolidationParams,
	valid bool,
) (int64, error) {
	groupNameVal, categoryVal := scopedTargetValues(p)

	query := `
		INSERT INTO ` + constants.TableConsolidatedEntries + `
			(entry, generic_source_type, list_type, consolidation_type,
			 group_name, category, valid, source_count)
		SELECT e.entry, e.generic_source_type, 'allowlist', ?, ?, ?, ?,
			COUNT(DISTINCT e.source_id)
		FROM ` + constants.TableEntries + ` e
		JOIN ` + constants.TableSources + ` s ON s.id = e.source_id
		JOIN ` + p.JoinTable + ` sc
			ON sc.source_id = e.source_id
			AND sc.source_type = e.generic_source_type
			AND sc.list_type = e.list_type
			AND sc.` + p.ScopeColumn + ` = ?
		WHERE e.generic_source_type = ?
			AND e.list_type = 'allowlist'
			AND e.valid = ?
			AND s.disabled = 0
			AND s.` + p.SkipColumn + ` = 0
		GROUP BY e.entry`

	var inserted int64
	err := r.db.InTransaction(ctx, func(tx *sql.Tx) error {
		result, execErr := tx.ExecContext(ctx, query,
			p.ConsolidationType, groupNameVal, categoryVal, boolToInt(valid),
			p.ScopeValue, p.GenericSourceType, boolToInt(valid))
		if execErr != nil {
			return fmt.Errorf("consolidating scoped allowlist %s/%s: %w",
				p.ScopeValue, p.GenericSourceType, execErr)
		}
		inserted, _ = result.RowsAffected() // nolint: errcheck
		return nil
	})
	return inserted, err
}

// ConsolidateScopedBlocklist consolidates blocklist entries for one scope value
// and generic source type, applying the allowlist filtering.
func (r *ConsolidatedRepo) ConsolidateScopedBlocklist(
	ctx context.Context,
	p scopedConsolidationParams,
	valid bool,
) (BlocklistConsolidationResult, error) {
	var res BlocklistConsolidationResult
	groupNameVal, categoryVal := scopedTargetValues(p)

	countQuery := `
		SELECT COUNT(DISTINCT e.entry), COUNT(DISTINCT e.source_id)
		FROM ` + constants.TableEntries + ` e
		JOIN ` + constants.TableSources + ` s ON s.id = e.source_id
		JOIN ` + p.JoinTable + ` sc
			ON sc.source_id = e.source_id
			AND sc.source_type = e.generic_source_type
			AND sc.list_type = e.list_type
			AND sc.` + p.ScopeColumn + ` = ?
		WHERE e.generic_source_type = ?
			AND e.list_type = 'blocklist'
			AND e.valid = ?
			AND s.disabled = 0
			AND s.` + p.SkipColumn + ` = 0`
	if err := r.db.readConn.QueryRowContext(ctx, countQuery,
		p.ScopeValue, p.GenericSourceType, boolToInt(valid),
	).Scan(&res.OriginalCount, &res.SourceCount); err != nil {
		return res, fmt.Errorf("counting scoped blocklist %s/%s: %w",
			p.ScopeValue, p.GenericSourceType, err)
	}
	if res.OriginalCount == 0 {
		return res, nil
	}

	insertQuery := `
		INSERT INTO ` + constants.TableConsolidatedEntries + `
			(entry, generic_source_type, list_type, consolidation_type,
			 group_name, category, valid, source_count)
		SELECT be.entry, be.generic_source_type, 'blocklist', ?, ?, ?, ?,
			COUNT(DISTINCT be.source_id)
		FROM ` + constants.TableEntries + ` be
		JOIN ` + constants.TableSources + ` bs ON bs.id = be.source_id
		JOIN ` + p.JoinTable + ` bsc
			ON bsc.source_id = be.source_id
			AND bsc.source_type = be.generic_source_type
			AND bsc.list_type = be.list_type
			AND bsc.` + p.ScopeColumn + ` = ?
		LEFT JOIN (
			SELECT ae.entry AS entry, MAX(ae.must_consider) AS allow_must
			FROM ` + constants.TableEntries + ` ae
			JOIN ` + constants.TableSources + ` asrc ON asrc.id = ae.source_id
			JOIN ` + p.JoinTable + ` ac
				ON ac.source_id = ae.source_id
				AND ac.source_type = ae.generic_source_type
				AND ac.list_type = ae.list_type
				AND ac.` + p.ScopeColumn + ` = ?
			WHERE ae.generic_source_type = ?
				AND ae.list_type = 'allowlist'
				AND ae.valid = 1
				AND asrc.disabled = 0
				AND asrc.` + p.SkipColumn + ` = 0
			GROUP BY ae.entry
		) al ON al.entry = be.entry
		WHERE be.generic_source_type = ?
			AND be.list_type = 'blocklist'
			AND be.valid = ?
			AND bs.disabled = 0
			AND bs.` + p.SkipColumn + ` = 0
		GROUP BY be.entry
		HAVING MAX(CASE WHEN al.entry IS NOT NULL THEN 1 ELSE 0 END) = 0
		    OR (MAX(be.must_consider) = 1 AND COALESCE(MAX(al.allow_must), 0) = 0)`

	err := r.db.InTransaction(ctx, func(tx *sql.Tx) error {
		result, execErr := tx.ExecContext(ctx, insertQuery,
			p.ConsolidationType, groupNameVal, categoryVal, boolToInt(valid),
			p.ScopeValue,
			p.ScopeValue,
			p.GenericSourceType,
			p.GenericSourceType,
			boolToInt(valid))
		if execErr != nil {
			return fmt.Errorf("consolidating scoped blocklist %s/%s: %w",
				p.ScopeValue, p.GenericSourceType, execErr)
		}
		res.FinalCount, _ = result.RowsAffected() // nolint: errcheck
		return nil
	})
	return res, err
}

// GetScopeSourceNames returns the distinct enabled source names that contribute
// entries for a given consolidation scope (general/group/category), generic
// source type, and list type.
func (r *ConsolidatedRepo) GetScopeSourceNames(
	ctx context.Context,
	consolidationType, gst, listType, groupName, category string,
) ([]string, error) {
	var sb strings.Builder
	sb.WriteString("SELECT DISTINCT s.name FROM " + constants.TableEntries + " e ")
	sb.WriteString("JOIN " + constants.TableSources + " s ON s.id = e.source_id ")

	var args []any
	switch consolidationType {
	case constants.ConsolidationTypeGroup:
		sb.WriteString("JOIN " + constants.TableEntryGroups + ` sc
			ON sc.source_id = e.source_id
			AND sc.source_type = e.generic_source_type
			AND sc.list_type = e.list_type
			AND sc.group_name = ? `)
		args = append(args, groupName)
	case constants.ConsolidationTypeCategory:
		sb.WriteString("JOIN " + constants.TableEntryCategories + ` sc
			ON sc.source_id = e.source_id
			AND sc.source_type = e.generic_source_type
			AND sc.list_type = e.list_type
			AND sc.category = ? `)
		args = append(args, category)
	}

	sb.WriteString("WHERE e.generic_source_type = ? AND e.list_type = ? AND e.valid = 1 AND s.disabled = 0")
	args = append(args, gst, listType)

	switch consolidationType {
	case constants.ConsolidationTypeGroup:
		sb.WriteString(" AND s.skip_groups_consolidation = 0")
	case constants.ConsolidationTypeCategory:
		sb.WriteString(" AND s.skip_categories_consolidation = 0")
	default:
		sb.WriteString(" AND s.skip_general_consolidation = 0")
	}
	sb.WriteString(" ORDER BY s.name")

	rows, err := r.db.readConn.QueryContext(ctx, sb.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("querying scope source names: %w", err)
	}
	defer rows.Close() // nolint: errcheck

	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("scanning source name: %w", err)
		}
		names = append(names, name)
	}
	return names, rows.Err()
}

// ScopeSourceKey identifies a unique scope for source name lookup.
type ScopeSourceKey struct {
	GenericSourceType string
	ListType          string
	GroupName         string
	Category          string
}

// GetAllScopeSourceNames returns all source names for a consolidation type.
func (r *ConsolidatedRepo) GetAllScopeSourceNames(
	ctx context.Context, consolidationType string,
) (map[ScopeSourceKey][]string, error) {
	sc, err := scopeConfigFor(consolidationType)
	if err != nil {
		return r.getAllGeneralSourceNames(ctx)
	}

	query := `SELECT sc.` + sc.scopeColumn + `, e.generic_source_type, e.list_type, s.name
		FROM ` + constants.TableEntries + ` e
		JOIN ` + constants.TableSources + ` s ON s.id = e.source_id
		JOIN ` + sc.joinTable + ` sc
			ON sc.source_id = e.source_id
			AND sc.source_type = e.generic_source_type
			AND sc.list_type = e.list_type
		WHERE e.valid = 1
			AND s.disabled = 0
			AND s.` + sc.skipColumn + ` = 0
		GROUP BY sc.` + sc.scopeColumn + `, e.generic_source_type, e.list_type, s.name
		ORDER BY sc.` + sc.scopeColumn + `, e.generic_source_type, e.list_type, s.name`

	rows, err := r.db.readConn.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("querying all scope source names (%s): %w", consolidationType, err)
	}
	defer rows.Close() // nolint: errcheck

	result := make(map[ScopeSourceKey][]string)
	for rows.Next() {
		var scopeValue, gst, listType, name string
		if err := rows.Scan(&scopeValue, &gst, &listType, &name); err != nil {
			return nil, fmt.Errorf("scanning scope source name: %w", err)
		}
		key := ScopeSourceKey{GenericSourceType: gst, ListType: listType}
		if consolidationType == constants.ConsolidationTypeCategory {
			key.Category = scopeValue
		} else {
			key.GroupName = scopeValue
		}
		result[key] = append(result[key], name)
	}
	return result, rows.Err()
}

func (r *ConsolidatedRepo) getAllGeneralSourceNames(ctx context.Context) (map[ScopeSourceKey][]string, error) {
	query := `SELECT e.generic_source_type, e.list_type, s.name
		FROM ` + constants.TableEntries + ` e
		JOIN ` + constants.TableSources + ` s ON s.id = e.source_id
		WHERE e.valid = 1
			AND s.disabled = 0
			AND s.skip_general_consolidation = 0
		GROUP BY e.generic_source_type, e.list_type, s.name
		ORDER BY e.generic_source_type, e.list_type, s.name`

	rows, err := r.db.readConn.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("querying all general source names: %w", err)
	}
	defer rows.Close() // nolint: errcheck

	result := make(map[ScopeSourceKey][]string)
	for rows.Next() {
		var gst, listType, name string
		if err := rows.Scan(&gst, &listType, &name); err != nil {
			return nil, fmt.Errorf("scanning general source name: %w", err)
		}
		key := ScopeSourceKey{GenericSourceType: gst, ListType: listType}
		result[key] = append(result[key], name)
	}
	return result, rows.Err()
}

func scopeConfigFor(consolidationType string) (scopeConfig, error) {
	switch consolidationType {
	case constants.ConsolidationTypeGroup:
		return scopeConfig{constants.TableEntryGroups, "group_name", "skip_groups_consolidation"}, nil
	case constants.ConsolidationTypeCategory:
		return scopeConfig{constants.TableEntryCategories, "category", "skip_categories_consolidation"}, nil
	default:
		return scopeConfig{}, fmt.Errorf("unsupported scoped consolidation type: %s", consolidationType)
	}
}

func (r *ConsolidatedRepo) loadScopedAllowSet(
	ctx context.Context, consolidationType string, sc scopeConfig, gst string,
) error {
	// clear prior rows for this consolidation type
	if _, err := r.db.writeConn.ExecContext(ctx,
		"delete from dnstk_scoped_allow where consolidation_type = ?",
		consolidationType); err != nil {
		return fmt.Errorf("clearing scoped_allow: %w", err)
	}

	_, err := r.db.writeConn.ExecContext(ctx, `
		INSERT OR IGNORE INTO dnstk_scoped_allow
			(consolidation_type, scope_value, entry, must_consider)
		SELECT ?, sc.`+sc.scopeColumn+`, ae.entry, MAX(ae.must_consider)
		FROM `+constants.TableEntries+` ae
		JOIN `+constants.TableSources+` s ON s.id = ae.source_id
		JOIN `+sc.joinTable+` sc
			ON sc.source_id = ae.source_id
			AND sc.source_type = ae.generic_source_type
			AND sc.list_type = ae.list_type
		WHERE ae.generic_source_type = ?
			AND ae.list_type = 'allowlist'
			AND ae.valid = 1
			AND s.disabled = 0
			AND s.`+sc.skipColumn+` = 0
		GROUP BY sc.`+sc.scopeColumn+`, ae.entry`,
		consolidationType, gst)
	if err != nil {
		return fmt.Errorf("loading scoped_allow for %s/%s: %w", consolidationType, gst, err)
	}
	return nil
}

// ConsolidateScopedBlocklistAll consolidates blocklist entries for all scope values (groups or categories)
// for a given generic source type
func (r *ConsolidatedRepo) ConsolidateScopedBlocklistAll(
	ctx context.Context, consolidationType, gst string, valid bool,
) ([]scopedConsolidationResult, error) {
	sc, err := scopeConfigFor(consolidationType)
	if err != nil {
		return nil, err
	}

	if valid {
		if err := r.loadScopedAllowSet(ctx, consolidationType, sc, gst); err != nil {
			return nil, err
		}
	}

	scopeTargetCol := "group_name"
	if consolidationType == "category" {
		scopeTargetCol = "category"
	}

	insertQuery := `
		INSERT INTO ` + constants.TableConsolidatedEntries + `
			(entry, generic_source_type, list_type, consolidation_type,
			 ` + scopeTargetCol + `, ` + otherTargetCol(consolidationType) + `,
			 valid, source_count)
		SELECT be.entry, be.generic_source_type, 'blocklist', ?,
			bsc.` + sc.scopeColumn + `, NULL,
			?, COUNT(DISTINCT be.source_id)
		FROM ` + constants.TableEntries + ` be
		JOIN ` + constants.TableSources + ` bs ON bs.id = be.source_id
		JOIN ` + sc.joinTable + ` bsc
			ON bsc.source_id = be.source_id
			AND bsc.source_type = be.generic_source_type
			AND bsc.list_type = be.list_type
		LEFT JOIN dnstk_scoped_allow sa
			ON sa.consolidation_type = ?
			AND sa.scope_value = bsc.` + sc.scopeColumn + `
			AND sa.entry = be.entry
		WHERE be.generic_source_type = ?
			AND be.list_type = 'blocklist'
			AND be.valid = ?
			AND bs.disabled = 0
			AND bs.` + sc.skipColumn + ` = 0
		GROUP BY bsc.` + sc.scopeColumn + `, be.entry
		HAVING MAX(CASE WHEN sa.entry IS NOT NULL THEN 1 ELSE 0 END) = 0
		    OR (MAX(be.must_consider) = 1 AND COALESCE(MAX(sa.must_consider), 0) = 0)`

	if err := r.db.InTransaction(ctx, func(tx *sql.Tx) error {
		_, execErr := tx.ExecContext(ctx, insertQuery,
			consolidationType, boolToInt(valid), // SELECT cols
			consolidationType,     // LEFT JOIN sa
			gst, boolToInt(valid)) // WHERE
		return execErr
	}); err != nil {
		return nil, fmt.Errorf("collapsed blocklist consolidation %s/%s: %w",
			consolidationType, gst, err)
	}

	return r.scopedSummaries(ctx, consolidationType, sc, gst, "blocklist", valid)
}

// ConsolidateScopedAllowlistAll consolidates allowlist entries for all scope values (groups or categories)
func (r *ConsolidatedRepo) ConsolidateScopedAllowlistAll(
	ctx context.Context,
	consolidationType, gst string,
	valid bool,
) ([]scopedConsolidationResult, error) {
	sc, err := scopeConfigFor(consolidationType)
	if err != nil {
		return nil, err
	}
	scopeTargetCol := "group_name"
	if consolidationType == constants.ConsolidationTypeCategory {
		scopeTargetCol = "category"
	}

	insertQuery := `
		INSERT INTO ` + constants.TableConsolidatedEntries + `
			(entry, generic_source_type, list_type, consolidation_type,
			 ` + scopeTargetCol + `, ` + otherTargetCol(consolidationType) + `,
			 valid, source_count)
		SELECT e.entry, e.generic_source_type, 'allowlist', ?,
			sc.` + sc.scopeColumn + `, NULL,
			?, COUNT(DISTINCT e.source_id)
		FROM ` + constants.TableEntries + ` e
		JOIN ` + constants.TableSources + ` s ON s.id = e.source_id
		JOIN ` + sc.joinTable + ` sc
			ON sc.source_id = e.source_id
			AND sc.source_type = e.generic_source_type
			AND sc.list_type = e.list_type
		WHERE e.generic_source_type = ?
			AND e.list_type = 'allowlist'
			AND e.valid = ?
			AND s.disabled = 0
			AND s.` + sc.skipColumn + ` = 0
		GROUP BY sc.` + sc.scopeColumn + `, e.entry`

	if err := r.db.InTransaction(ctx, func(tx *sql.Tx) error {
		_, execErr := tx.ExecContext(ctx, insertQuery,
			consolidationType, boolToInt(valid), gst, boolToInt(valid))
		return execErr
	}); err != nil {
		return nil, fmt.Errorf("collapsed allowlist consolidation %s/%s: %w",
			consolidationType, gst, err)
	}

	return r.scopedSummaries(ctx, consolidationType, sc, gst, "allowlist", valid)
}

func otherTargetCol(consolidationType string) string {
	if consolidationType == constants.ConsolidationTypeCategory {
		return "group_name"
	}
	return "category"
}

func (r *ConsolidatedRepo) scopedSummaries(
	ctx context.Context,
	consolidationType string, sc scopeConfig, gst, listType string, valid bool,
) ([]scopedConsolidationResult, error) {
	scopeTargetCol := "group_name"
	if consolidationType == constants.ConsolidationTypeCategory {
		scopeTargetCol = "category"
	}

	finalCounts := map[string]int64{}
	fcRows, err := r.db.readConn.QueryContext(ctx, `
		SELECT `+scopeTargetCol+`, COUNT(*)
		FROM `+constants.TableConsolidatedEntries+`
		WHERE consolidation_type = ? AND generic_source_type = ?
			AND list_type = ? AND valid = ?
		GROUP BY `+scopeTargetCol,
		consolidationType, gst, listType, boolToInt(valid))
	if err != nil {
		return nil, fmt.Errorf("reading scoped final counts: %w", err)
	}
	for fcRows.Next() {
		var scope string
		var cnt int64
		if scanErr := fcRows.Scan(&scope, &cnt); scanErr != nil {
			fcRows.Close() // nolint: errcheck
			return nil, scanErr
		}
		finalCounts[scope] = cnt
	}
	fcRows.Close() // nolint: errcheck
	if rowsErr := fcRows.Err(); rowsErr != nil {
		return nil, rowsErr
	}

	scRows, err := r.db.readConn.QueryContext(ctx, `
		SELECT sc.`+sc.scopeColumn+`,
			COUNT(DISTINCT s.id),
			GROUP_CONCAT(DISTINCT s.name)
		FROM `+constants.TableEntries+` e
		JOIN `+constants.TableSources+` s ON s.id = e.source_id
		JOIN `+sc.joinTable+` sc
			ON sc.source_id = e.source_id
			AND sc.source_type = e.generic_source_type
			AND sc.list_type = e.list_type
		WHERE e.generic_source_type = ? AND e.list_type = ? AND e.valid = ?
			AND s.disabled = 0 AND s.`+sc.skipColumn+` = 0
		GROUP BY sc.`+sc.scopeColumn,
		gst, listType, boolToInt(valid))
	if err != nil {
		return nil, fmt.Errorf("reading scoped source names: %w", err)
	}
	defer scRows.Close() // nolint: errcheck

	var results []scopedConsolidationResult
	for scRows.Next() {
		var scope, names string
		var srcCount int64
		if err := scRows.Scan(&scope, &srcCount, &names); err != nil {
			return nil, err
		}
		fc := finalCounts[scope]
		if fc == 0 {
			continue // scope had no entries for this list type
		}
		results = append(results, scopedConsolidationResult{
			ScopeValue:  scope,
			SourceCount: srcCount,
			FinalCount:  fc,
			SourceNames: splitConcat(names),
		})
	}
	return results, scRows.Err()
}

func splitConcat(s string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, ",")
}
