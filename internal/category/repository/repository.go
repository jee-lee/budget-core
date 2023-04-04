package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateCategory(ctx context.Context, request *CategoryCreateRequest) (*Category, error)
	GetCategory(ctx context.Context, id uuid.UUID) (*Category, error)
	GetCycleTypeByName(ctx context.Context, name string) (*CycleType, error)
	GetCycleTypeByID(ctx context.Context, id int) (*CycleType, error)
	GetDefaultCycleType(ctx context.Context) (*CycleType, error)
	CreateCycleTypes(ctx context.Context) error
}

type repository struct {
	Pool                 *sqlx.DB
	CycleTypeByNameCache map[string]*CycleType
	CycleTypeByIdCache   map[int]*CycleType
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		Pool:                 db,
		CycleTypeByNameCache: make(map[string]*CycleType),
		CycleTypeByIdCache:   make(map[int]*CycleType),
	}
}
