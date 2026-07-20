package db

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"fmt"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/multilog"
)

type SourcesRepo struct {
	db *DB
}

type importLookupCache struct {
	typeNameIDs     map[string]int64
	listTypeNameIDs map[string]int64
	groupNameIDs    map[string]int64
	categoryNameIDs map[string]int64
}

func newImportLookupCache() *importLookupCache {
	return &importLookupCache{
		typeNameIDs:     make(map[string]int64),
		listTypeNameIDs: make(map[string]int64),
		groupNameIDs:    make(map[string]int64),
		categoryNameIDs: make(map[string]int64),
	}
}

func NewSourcesRepo(db *DB) *SourcesRepo {
	return &SourcesRepo{db: db}
}

type SourceRow struct {
	License                     string
	Website                     string
	DefinitionChecksum          string
	SourceFile                  string
	Notes                       string
	URLPerGroup                 string
	Frequency                   string
	Name                        string
	URL                         string
	URLPerCategory              string
	TypeCount                   int
	ID                          int64
	CountToConsider             int
	SkipCategoriesConsolidation bool
	SkipGroupsConsolidation     bool
	SkipGeneralConsolidation    bool
	Disabled                    bool
}

func (r *SourcesRepo) GetSourceByName(name string) (*SourceRow, error) {
	row := r.db.readConn.QueryRow(`
		SELECT id, name, url, url_per_category, url_per_group, frequency, license, website, notes,
			type_count, count_to_consider, disabled, skip_general_consolidation, skip_groups_consolidation,
			skip_categories_consolidation, source_file, definition_checksum
		FROM `+constants.TableSources+` WHERE name = ?`, name)

	s := &SourceRow{}
	var disabled, skipGen, skipGrp, skipCat int
	err := row.Scan(&s.ID, &s.Name, &s.URL, &s.URLPerCategory, &s.URLPerGroup, &s.Frequency,
		&s.License, &s.Website, &s.Notes, &s.TypeCount, &s.CountToConsider,
		&disabled, &skipGen, &skipGrp, &skipCat, &s.SourceFile, &s.DefinitionChecksum)
	if err != nil {
		return nil, err
	}
	s.Disabled = disabled == 1
	s.SkipGeneralConsolidation = skipGen == 1
	s.SkipGroupsConsolidation = skipGrp == 1
	s.SkipCategoriesConsolidation = skipCat == 1
	return s, nil
}

// UpsertSource inserts a source or updates it if it exists by name.
// Returns the source ID.
func (r *SourcesRepo) UpsertSource(s *SourceRow) (int64, error) {
	result, err := r.db.writeConn.Exec(`
		INSERT INTO `+constants.TableSources+` (name, url, url_per_category, url_per_group, frequency,
			license, website, notes, type_count, count_to_consider, disabled,
			skip_general_consolidation, skip_groups_consolidation, skip_categories_consolidation,
			source_file, definition_checksum)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(name, source_file) DO UPDATE SET
			url = excluded.url,
			url_per_category = excluded.url_per_category,
			url_per_group = excluded.url_per_group,
			frequency = excluded.frequency,
			license = excluded.license,
			website = excluded.website,
			notes = excluded.notes,
			type_count = excluded.type_count,
			count_to_consider = excluded.count_to_consider,
			disabled = excluded.disabled,
			skip_general_consolidation = excluded.skip_general_consolidation,
			skip_groups_consolidation = excluded.skip_groups_consolidation,
			skip_categories_consolidation = excluded.skip_categories_consolidation,
			definition_checksum = excluded.definition_checksum`,
		s.Name, s.URL, s.URLPerCategory, s.URLPerGroup, s.Frequency,
		s.License, s.Website, s.Notes, s.TypeCount, s.CountToConsider,
		boolToInt(s.Disabled),
		boolToInt(s.SkipGeneralConsolidation),
		boolToInt(s.SkipGroupsConsolidation),
		boolToInt(s.SkipCategoriesConsolidation),
		s.SourceFile, s.DefinitionChecksum)
	if err != nil {
		return 0, fmt.Errorf("upserting source %s: %w", s.Name, err)
	}

	sourceID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("getting last insert ID for source %s: %w", s.Name, err)
	}
	return sourceID, nil
}

// ImportSource inserts or updates a source and all its related data.
// Returns true if the source was updated (changed), false if unchanged.
func (r *SourcesRepo) ImportSource(
	logger *multilog.Logger,
	tx *sql.Tx,
	source config.Source,
	lookupCache *importLookupCache,
	sourceFile string,
) (bool, error) {
	checksum := computeSourceChecksum(source)

	// Check if source already exists with same checksum
	var existingID int64
	var existingChecksum string
	err := tx.QueryRow(
		"SELECT id, definition_checksum FROM "+constants.TableSources+" WHERE name = ? AND source_file = ?",
		source.Name, sourceFile,
	).Scan(&existingID, &existingChecksum)
	if err != nil && err != sql.ErrNoRows {
		return false, fmt.Errorf("querying existing source %s: %w", source.Name, err)
	}

	if err == nil && existingChecksum == checksum {
		return false, nil
	}

	// Delete existing source and cascade to related tables
	if err == nil {
		deleteResult, deleteErr := tx.Exec(
			"DELETE FROM "+constants.TableSources+" WHERE id = ?",
			existingID)
		if deleteErr != nil {
			return false, fmt.Errorf("deleting existing source %s: %w", source.Name, deleteErr)
		}
		rowCount, _ := deleteResult.RowsAffected() // nolint: errcheck
		if rowCount > 0 {
			logger.Infof("Deleted %d rows for existing source %s with ID %d", rowCount, source.Name, existingID)
		}
	}

	result, err := tx.Exec(`
		INSERT INTO `+constants.TableSources+` (name, url, url_per_category, url_per_group, frequency, license,
			website, notes, type_count, count_to_consider, disabled,
			skip_general_consolidation, skip_groups_consolidation,
			skip_categories_consolidation, source_file, definition_checksum)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		source.Name, source.URL, source.URLPerCategory, source.URLPerGroup,
		defaultString(source.Frequency, "daily"),
		source.License, source.Website, source.Notes,
		source.TypeCount, source.CountToConsider, boolToInt(source.Disabled),
		boolToInt(source.SkipGeneralConsolidation),
		boolToInt(source.SkipGroupsConsolidation),
		boolToInt(source.SkipCategoriesConsolidation),
		sourceFile, checksum,
	)
	if err != nil {
		return false, fmt.Errorf("inserting source %s: %w", source.Name, err)
	}

	sourceID, err := result.LastInsertId()
	if err != nil {
		return false, fmt.Errorf("getting source ID for %s: %w", source.Name, err)
	}

	if err = r.insertSourceTypes(tx, sourceID, source.Types, lookupCache); err != nil {
		return false, err
	}

	upsertCategoryStmt, err := tx.Prepare(
		"INSERT INTO " + constants.TableCategoryNames + " (name) VALUES (?) ON CONFLICT(name) DO NOTHING",
	)
	if err != nil {
		return false, fmt.Errorf("preparing category upsert statement: %w", err)
	}
	defer func() { _ = upsertCategoryStmt.Close() }() // nolint: errcheck

	selectCategoryIDStmt, err := tx.Prepare(
		"SELECT id FROM " + constants.TableCategoryNames + " WHERE name = ?",
	)
	if err != nil {
		return false, fmt.Errorf("preparing category id lookup statement: %w", err)
	}
	defer func() { _ = selectCategoryIDStmt.Close() }() // nolint: errcheck

	insertSourceCategoryStmt, err := tx.Prepare(
		"INSERT INTO " + constants.TableSourceCategories + " (source_id, category_name_id) VALUES (?, ?)",
	)
	if err != nil {
		return false, fmt.Errorf("preparing source category insert statement: %w", err)
	}
	defer func() { _ = insertSourceCategoryStmt.Close() }() // nolint: errcheck

	for _, cat := range source.Categories {
		catNameID, idErr := getOrCreateLookupID(
			cat,
			lookupCache.categoryNameIDs,
			upsertCategoryStmt,
			selectCategoryIDStmt,
		)
		if idErr != nil {
			return false, fmt.Errorf("getting category name ID for %s: %w", cat, idErr)
		}

		if _, err := insertSourceCategoryStmt.Exec(sourceID, catNameID); err != nil {
			return false, fmt.Errorf("upserting category name %s for %s: %w", cat, source.Name, err)
		}
	}

	for _, country := range source.Countries {
		if _, err := tx.Exec(
			"INSERT INTO "+constants.TableSourceCountries+" (source_id, country_code) VALUES (?, ?)",
			sourceID, country,
		); err != nil {
			return false, fmt.Errorf("inserting country %s for %s: %w", country, source.Name, err)
		}
	}

	if err := r.insertContent(tx, sourceID, "content", source.Content); err != nil {
		return false, err
	}
	if err := r.insertContent(tx, sourceID, "content_per_category", source.ContentPerCategory); err != nil {
		return false, err
	}
	if err := r.insertContent(tx, sourceID, "content_per_group", source.ContentPerGroup); err != nil {
		return false, err
	}

	for _, f := range source.Files {
		if _, err := tx.Exec(
			"INSERT INTO "+constants.TableSourceFiles+" (source_id, filename) VALUES (?, ?)",
			sourceID, f,
		); err != nil {
			return false, fmt.Errorf("inserting file %s for %s: %w", f, source.Name, err)
		}
	}

	return true, nil
}

func (r *SourcesRepo) insertSourceTypes(
	tx *sql.Tx,
	sourceID int64,
	types []c.SourceType,
	lookupCache *importLookupCache,
) error {
	upsertTypeStmt, err := tx.Prepare(
		"INSERT INTO " + constants.TableTypeNames + " (name) VALUES (?) ON CONFLICT(name) DO NOTHING",
	)
	if err != nil {
		return fmt.Errorf("preparing type upsert statement: %w", err)
	}
	defer func() { _ = upsertTypeStmt.Close() }() // nolint: errcheck

	selectTypeIDStmt, err := tx.Prepare("SELECT id FROM " + constants.TableTypeNames + " WHERE name = ?")
	if err != nil {
		return fmt.Errorf("preparing type id lookup statement: %w", err)
	}
	defer func() { _ = selectTypeIDStmt.Close() }() // nolint: errcheck

	insertSourceTypeStmt, err := tx.Prepare(
		"INSERT INTO " + constants.TableSourceTypes +
			" (source_id, type_name_id, notes, disabled) VALUES (?, ?, ?, ?)",
	)
	if err != nil {
		return fmt.Errorf("preparing source type insert statement: %w", err)
	}
	defer func() { _ = insertSourceTypeStmt.Close() }() // nolint: errcheck

	upsertListTypeStmt, err := tx.Prepare(
		"INSERT INTO " + constants.TableListTypeNames + " (name) VALUES (?) ON CONFLICT(name) DO NOTHING",
	)
	if err != nil {
		return fmt.Errorf("preparing list type upsert statement: %w", err)
	}
	defer func() { _ = upsertListTypeStmt.Close() }() // nolint: errcheck

	selectListTypeIDStmt, err := tx.Prepare(
		"SELECT id FROM " + constants.TableListTypeNames + " WHERE name = ?",
	)
	if err != nil {
		return fmt.Errorf("preparing list type id lookup statement: %w", err)
	}
	defer func() { _ = selectListTypeIDStmt.Close() }() // nolint: errcheck

	insertSourceListTypeStmt, err := tx.Prepare(
		"INSERT INTO " + constants.TableSourceListTypes +
			" (source_type_id, list_type_name_id, disabled, must_consider) VALUES (?, ?, ?, ?)",
	)
	if err != nil {
		return fmt.Errorf("preparing source list type insert statement: %w", err)
	}
	defer func() { _ = insertSourceListTypeStmt.Close() }() // nolint: errcheck

	insertSourceListTypeNotesStmt, err := tx.Prepare(
		"INSERT INTO " + constants.TableSourceListTypeNotes + " (source_list_type_id, notes) VALUES (?, ?)",
	)
	if err != nil {
		return fmt.Errorf("preparing source list type notes insert statement: %w", err)
	}
	defer func() { _ = insertSourceListTypeNotesStmt.Close() }() // nolint: errcheck

	upsertGroupStmt, err := tx.Prepare(
		"INSERT INTO " + constants.TableGroupNames + " (name) VALUES (?) ON CONFLICT(name) DO NOTHING",
	)
	if err != nil {
		return fmt.Errorf("preparing group upsert statement: %w", err)
	}
	defer func() { _ = upsertGroupStmt.Close() }() // nolint: errcheck

	selectGroupIDStmt, err := tx.Prepare("SELECT id FROM " + constants.TableGroupNames + " WHERE name = ?")
	if err != nil {
		return fmt.Errorf("preparing group id lookup statement: %w", err)
	}
	defer func() { _ = selectGroupIDStmt.Close() }() // nolint: errcheck

	insertSourceListTypeGroupStmt, err := tx.Prepare(
		"INSERT INTO " + constants.TableSourceListTypeGroups +
			" (source_list_type_id, group_name_id) VALUES (?, ?)",
	)
	if err != nil {
		return fmt.Errorf("preparing source list type group insert statement: %w", err)
	}
	defer func() { _ = insertSourceListTypeGroupStmt.Close() }() // nolint: errcheck

	for _, st := range types {
		typeNameID, idErr := getOrCreateLookupID(
			st.Name,
			lookupCache.typeNameIDs,
			upsertTypeStmt,
			selectTypeIDStmt,
		)
		if idErr != nil {
			return fmt.Errorf("getting type name ID for %s: %w", st.Name, idErr)
		}

		// Insert the per-source source_type row
		result, err := insertSourceTypeStmt.Exec(sourceID, typeNameID, st.Notes, boolToInt(st.Disabled))
		if err != nil {
			return fmt.Errorf("inserting source type %s: %w", st.Name, err)
		}

		stID, err := result.LastInsertId()
		if err != nil {
			return fmt.Errorf("getting source type ID for %s: %w", st.Name, err)
		}

		for _, lt := range st.ListTypes {
			listTypeNameID, idErr := getOrCreateLookupID(
				lt.Name,
				lookupCache.listTypeNameIDs,
				upsertListTypeStmt,
				selectListTypeIDStmt,
			)
			if idErr != nil {
				return fmt.Errorf("getting list type name ID for %s: %w", lt.Name, idErr)
			}

			// Insert the per-source-type source_list_type row (no notes column)
			ltResult, err := insertSourceListTypeStmt.Exec(
				stID,
				listTypeNameID,
				boolToInt(lt.Disabled),
				boolToInt(lt.MustConsider),
			)
			if err != nil {
				return fmt.Errorf("inserting list type %s: %w", lt.Name, err)
			}

			ltID, err := ltResult.LastInsertId()
			if err != nil {
				return fmt.Errorf("getting list type ID for %s: %w", lt.Name, err)
			}

			// Store notes in separate table only when non-empty
			if lt.Notes != "" {
				if _, err := insertSourceListTypeNotesStmt.Exec(ltID, lt.Notes); err != nil {
					return fmt.Errorf("inserting list type notes for %s: %w", lt.Name, err)
				}
			}

			for _, g := range lt.Groups {
				groupNameID, idErr := getOrCreateLookupID(
					g,
					lookupCache.groupNameIDs,
					upsertGroupStmt,
					selectGroupIDStmt,
				)
				if idErr != nil {
					return fmt.Errorf("getting group name ID for %s: %w", g, idErr)
				}

				if _, err := insertSourceListTypeGroupStmt.Exec(ltID, groupNameID); err != nil {
					return fmt.Errorf("inserting group %s: %w", g, err)
				}
			}
		}
	}
	return nil
}

func getOrCreateLookupID(
	name string,
	cache map[string]int64,
	insertStmt *sql.Stmt,
	selectStmt *sql.Stmt,
) (int64, error) {
	if id, ok := cache[name]; ok {
		return id, nil
	}

	if _, err := insertStmt.Exec(name); err != nil {
		return 0, err
	}

	var id int64
	if err := selectStmt.QueryRow(name).Scan(&id); err != nil {
		return 0, err
	}

	cache[name] = id
	return id, nil
}

func (r *SourcesRepo) insertContent(tx *sql.Tx, sourceID int64, contentType string, entries []string) error {
	if len(entries) == 0 {
		return nil
	}
	stmt, err := tx.Prepare(
		"INSERT INTO " + constants.TableSourceContent + " (source_id, content_type, entry) VALUES (?, ?, ?)",
	)
	if err != nil {
		return fmt.Errorf("preparing content insert: %w", err)
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	for _, entry := range entries {
		if _, err := stmt.Exec(sourceID, contentType, entry); err != nil {
			return fmt.Errorf("inserting content entry: %w", err)
		}
	}
	return nil
}

// ImportSourcesFromConfig imports all sources from a SourcesConfig.
// Returns the number of sources imported, the number of sources skipped, and an error if any.
func (r *SourcesRepo) ImportSourcesFromConfig(
	ctx context.Context,
	logger *multilog.Logger,
	sourcesConfig config.SourcesConfig,
	sourceFile string,
) (int, int, error) {
	imported := 0
	skipped := 0

	err := r.db.InTransaction(ctx, func(tx *sql.Tx) error {
		lookupCache := newImportLookupCache()
		for _, source := range sourcesConfig.Sources {
			changed, err := r.ImportSource(logger, tx, source, lookupCache, sourceFile)
			if err != nil {
				return err
			}
			if changed {
				imported++
			} else {
				skipped++
			}
		}
		return nil
	})

	return imported, skipped, err
}

// GetSourceIDByName returns the source ID for a given name, or 0 if not found.
func (r *SourcesRepo) GetSourceIDByName(name string) (int64, error) {
	var id int64
	err := r.db.readConn.QueryRow("SELECT id FROM "+constants.TableSources+" WHERE name = ?", name).Scan(&id)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return id, err
}

// GetEnabledSources returns all enabled sources.
func (r *SourcesRepo) GetEnabledSources() ([]SourceRow, error) {
	rows, err := r.db.readConn.Query(`
		SELECT id, name, url, frequency, disabled, source_file,
			skip_general_consolidation, skip_groups_consolidation, skip_categories_consolidation
		FROM ` + constants.TableSources + ` WHERE disabled = 0
		ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	var sources []SourceRow
	for rows.Next() {
		s := SourceRow{}
		var disabled, skipGen, skipGrp, skipCat int
		if err := rows.Scan(&s.ID, &s.Name, &s.URL, &s.Frequency, &disabled, &s.SourceFile,
			&skipGen, &skipGrp, &skipCat); err != nil {
			return nil, err
		}
		s.Disabled = disabled == 1
		s.SkipGeneralConsolidation = skipGen == 1
		s.SkipGroupsConsolidation = skipGrp == 1
		s.SkipCategoriesConsolidation = skipCat == 1
		sources = append(sources, s)
	}
	return sources, rows.Err()
}

func (r *SourcesRepo) GetSourceCount() (int, error) {
	var count int
	err := r.db.readConn.QueryRow("SELECT COUNT(*) FROM " + constants.TableSources).Scan(&count)
	return count, err
}

func (r *SourcesRepo) ClearAllSources() error {
	_, err := r.db.writeConn.Exec("DELETE FROM " + constants.TableSources)
	return err
}

func computeSourceChecksum(source config.Source) string {
	data, err := json.Marshal(source)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", md5.Sum(data))
}

func defaultString(s, def string) string {
	if s == "" {
		return def
	}
	return s
}
