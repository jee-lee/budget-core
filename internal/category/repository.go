package category

import (
	"context"
	"github.com/jackc/pgtype"
	"github.com/jee-lee/budget-core/internal/repository"
)

type Repository interface {
	CreateCategory(ctx context.Context, request repository.CategoryCreateRequest) (*repository.Category, error)
	GetCategory(ctx context.Context, id pgtype.UUID) (*repository.Category, error)
}
