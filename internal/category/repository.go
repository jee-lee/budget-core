package category

import (
	"context"
	"github.com/google/uuid"
	"github.com/jee-lee/budget-core/internal/repository"
)

type Repository interface {
	CreateCategory(ctx context.Context, request repository.CategoryCreateRequest) (*repository.Category, error)
	GetCategory(ctx context.Context, id uuid.UUID) (*repository.Category, error)
}
