package postgres

import (
	"context"
	"database/sql"
)

//go:generate sqlc generate -f sqlc/sqlc.yaml

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func toNullString(str *string) sql.NullString {
	if str == nil {
		return sql.NullString{}
	}
	return sql.NullString{
		String: *str,
		Valid:  true,
	}
}
