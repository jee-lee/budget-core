package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateCategory(ctx context.Context, request CategoryCreateRequest) (*Category, error)
	GetCategory(ctx context.Context, id string) (*Category, error)
}

type repository struct {
	Pool sqlx.DB
}

func NewRepository(db sqlx.DB) Repository {
	return repository{
		Pool: db,
	}
}
