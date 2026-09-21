package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/phani-kb/dns-toolkit/internal/constants"
)

func insertRowsBatch(
	ctx context.Context, tx *sql.Tx, table string, columns []string, rows [][]any, orIgnore bool,
) (int64, error) {
	if len(rows) == 0 {
		return 0, nil
	}

	var inserted int64
	for i := 0; i < len(rows); i += constants.BulkInsertBatchSize {
		end := min(i+constants.BulkInsertBatchSize, len(rows))
		batch := rows[i:end]

		q := qb().Insert(table).Columns(columns...)
		if orIgnore {
			q = q.Options("OR IGNORE")
		}
		for _, row := range batch {
			q = q.Values(row...)
		}

		query, args, err := q.ToSql()
		if err != nil {
			return inserted, fmt.Errorf("building insert for %s: %w", table, err)
		}
		result, err := tx.ExecContext(ctx, query, args...)
		if err != nil {
			return inserted, fmt.Errorf("inserting batch into %s: %w", table, err)
		}
		if affected, affErr := result.RowsAffected(); affErr == nil {
			inserted += affected
		}
	}
	return inserted, nil
}
