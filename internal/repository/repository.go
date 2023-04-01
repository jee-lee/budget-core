package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type DB interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
}

type Repository struct {
	Pool DB
}

func NewRepository(db DB) *Repository {
	return &Repository{Pool: db}
}
