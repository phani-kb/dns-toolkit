package db

import (
	"context"

	"github.com/jmoiron/sqlx"
	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
)

// SourceMetadataLoader is the single shared loader for source type / list
// type / group metadata
type SourceMetadataLoader struct {
	readConn *sqlx.DB
}

func newSourceMetadataLoader(db *DB) SourceMetadataLoader {
	return SourceMetadataLoader{readConn: db.readConn}
}

// BatchSourceTypes loads source types + list types + groups for all sourceIDs
func (l SourceMetadataLoader) BatchSourceTypes(sourceIDs []int64) (map[int64][]c.SourceType, error) {
	out := make(map[int64][]c.SourceType, len(sourceIDs))
	if len(sourceIDs) == 0 {
		return out, nil
	}
	type typeRow struct {
		Name     string `db:"name"`
		Notes    string `db:"notes"`
		SourceID int64  `db:"source_id"`
		ID       int64  `db:"id"`
		Disabled int    `db:"disabled"`
	}
	allTypes, err := selectChunked[typeRow](context.Background(), l.readConn, sourceIDs, func(n int) string {
		return `
			select st.source_id as source_id, st.id as id, tn.name as name,
				coalesce(st.notes, '') as notes, st.disabled as disabled
			from ` + constants.TableSourceTypes + ` st
			INNER JOIN ` + constants.TableTypeNames + ` tn ON tn.id = st.type_name_id
			WHERE st.source_id IN (` + placeholders(n) + `) AND st.disabled = 0
			ORDER BY st.source_id, st.id`
	}, "batch querying source types")
	if err != nil {
		return nil, err
	}
	if len(allTypes) == 0 {
		return out, nil
	}
	typeIDs := make([]int64, 0, len(allTypes))
	for _, t := range allTypes {
		typeIDs = append(typeIDs, t.ID)
	}
	listTypesByType, err := l.BatchSourceListTypes(typeIDs)
	if err != nil {
		return nil, err
	}
	for _, t := range allTypes {
		lts := listTypesByType[t.ID]
		if len(lts) == 0 {
			continue
		}
		out[t.SourceID] = append(out[t.SourceID], c.SourceType{
			Name:      t.Name,
			Notes:     t.Notes,
			Disabled:  t.Disabled == 1,
			ListTypes: lts,
		})
	}
	return out, nil
}

// BatchSourceListTypes loads list types + groups for all source type IDs
func (l SourceMetadataLoader) BatchSourceListTypes(typeIDs []int64) (map[int64][]c.ListType, error) {
	out := make(map[int64][]c.ListType)
	if len(typeIDs) == 0 {
		return out, nil
	}
	type ltRow struct {
		Name         string `db:"name"`
		Notes        string `db:"notes"`
		TypeID       int64  `db:"source_type_id"`
		ID           int64  `db:"id"`
		Disabled     int    `db:"disabled"`
		MustConsider int    `db:"must_consider"`
	}
	all, err := selectChunked[ltRow](context.Background(), l.readConn, typeIDs, func(n int) string {
		return `
			select slt.source_type_id as source_type_id, slt.id as id, ltn.name as name,
				coalesce(sltn.notes, '') as notes, slt.disabled as disabled, slt.must_consider as must_consider
			from ` + constants.TableSourceListTypes + ` slt
			INNER JOIN ` + constants.TableListTypeNames + ` ltn ON ltn.id = slt.list_type_name_id
			LEFT JOIN ` + constants.TableSourceListTypeNotes + ` sltn ON sltn.source_list_type_id = slt.id
			WHERE slt.source_type_id IN (` + placeholders(n) + `) AND slt.disabled = 0
			ORDER BY slt.source_type_id, slt.id`
	}, "batch querying list types")
	if err != nil {
		return nil, err
	}
	ltIDs := make([]int64, 0, len(all))
	for _, row := range all {
		ltIDs = append(ltIDs, row.ID)
	}
	groupsByLT, err := l.BatchListTypeGroups(ltIDs)
	if err != nil {
		return nil, err
	}
	for _, row := range all {
		out[row.TypeID] = append(out[row.TypeID], c.ListType{
			Name:         row.Name,
			Notes:        row.Notes,
			Disabled:     row.Disabled == 1,
			MustConsider: row.MustConsider == 1,
			Groups:       groupsByLT[row.ID],
		})
	}
	return out, nil
}

// BatchSourceCategories loads category names for all source IDs in bulk
func (l SourceMetadataLoader) BatchSourceCategories(sourceIDs []int64) (map[int64][]string, error) {
	out := make(map[int64][]string, len(sourceIDs))
	if len(sourceIDs) == 0 {
		return out, nil
	}
	type categoryRow struct {
		Name     string `db:"name"`
		SourceID int64  `db:"source_id"`
	}
	rows, err := selectChunked[categoryRow](context.Background(), l.readConn, sourceIDs, func(n int) string {
		return `
			select sc.source_id as source_id, cn.name as name
			from ` + constants.TableSourceCategories + ` sc
			INNER JOIN ` + constants.TableCategoryNames + ` cn ON cn.id = sc.category_name_id
			WHERE sc.source_id IN (` + placeholders(n) + `)
			ORDER BY sc.source_id, cn.name`
	}, "batch querying source categories")
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		out[row.SourceID] = append(out[row.SourceID], row.Name)
	}
	return out, nil
}

// BatchSourceFiles loads archive filenames for all source IDs in bulk.
func (l SourceMetadataLoader) BatchSourceFiles(sourceIDs []int64) (map[int64][]string, error) {
	out := make(map[int64][]string, len(sourceIDs))
	if len(sourceIDs) == 0 {
		return out, nil
	}
	type fileRow struct {
		Filename string `db:"filename"`
		SourceID int64  `db:"source_id"`
	}
	rows, err := selectChunked[fileRow](context.Background(), l.readConn, sourceIDs, func(n int) string {
		return "select source_id, filename from " + constants.TableSourceFiles +
			" WHERE source_id IN (" + placeholders(n) + ") ORDER BY source_id, filename"
	}, "batch querying source files")
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		out[row.SourceID] = append(out[row.SourceID], row.Filename)
	}
	return out, nil
}

// BatchListTypeGroups loads group names for all source list type IDs
func (l SourceMetadataLoader) BatchListTypeGroups(ltIDs []int64) (map[int64][]string, error) {
	out := make(map[int64][]string)
	if len(ltIDs) == 0 {
		return out, nil
	}
	type listTypeGroupRow struct {
		Name       string `db:"name"`
		ListTypeID int64  `db:"source_list_type_id"`
	}
	rows, err := selectChunked[listTypeGroupRow](context.Background(), l.readConn, ltIDs, func(n int) string {
		return `
			select sltg.source_list_type_id as source_list_type_id, gn.name as name
			from ` + constants.TableSourceListTypeGroups + ` sltg
			INNER JOIN ` + constants.TableGroupNames + ` gn ON gn.id = sltg.group_name_id
			WHERE sltg.source_list_type_id IN (` + placeholders(n) + `)
			ORDER BY sltg.source_list_type_id, gn.name`
	}, "batch querying list type groups")
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		out[row.ListTypeID] = append(out[row.ListTypeID], row.Name)
	}
	return out, nil
}
