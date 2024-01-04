// Package postgres provides
package postgres

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	db          *sqlx.DB
	stmtBuilder squirrel.StatementBuilderType
}

// NewRepository .
func NewRepository(conn *sqlx.DB) *Postgres {
	return &Postgres{
		db:          conn,
		stmtBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
