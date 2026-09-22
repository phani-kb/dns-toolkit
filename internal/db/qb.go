package db

import (
	sq "github.com/Masterminds/squirrel"
)

// qb returns a squirrel StatementBuilder with Question placeholders (?)
func qb() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Question)
}
