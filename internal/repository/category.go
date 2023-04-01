package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgtype"
	pgxdecimal "github.com/jackc/pgtype/ext/shopspring-numeric"
	"github.com/jee-lee/budget-core/internal/helpers"
	"time"
)

type Category struct {
	ID               pgtype.UUID         `db:"id"`
	UserID           pgtype.UUID         `db:"user_id"`
	Name             string              `db:"name"`
	ParentCategoryID *pgtype.UUID        `db:"parent_category_id"`
	Maximum          *pgxdecimal.Numeric `db:"maximum"`
	CycleTypeID      *int                `db:"cycle_type_id"`
	Rollover         bool                `db:"rollover"`
	JointUserID      *pgtype.UUID        `db:"joint_user_id"`
	CreatedAt        time.Time           `db:"created_at"`
	UpdatedAt        time.Time           `db:"updated_at"`
}

type CategoryCreateRequest struct {
	UserID           pgtype.UUID         `db:"user_id"`
	Name             string              `db:"name"`
	ParentCategoryID *pgtype.UUID        `db:"parent_category_id"`
	Maximum          *pgxdecimal.Numeric `db:"maximum"`
	CycleTypeID      *int                `db:"cycle_type_id"`
	Rollover         bool                `db:"rollover"`
	JointUserID      *pgtype.UUID        `db:"joint_user_id"`
}

type CategoryUpdateRequest struct {
	ID               pgtype.UUID         `db:"id"`
	Name             string              `db:"name"`
	ParentCategoryId *pgtype.UUID        `db:"parent_category_id"`
	Maximum          *pgxdecimal.Numeric `db:"maximum"`
	Rollover         bool                `db:"rollover"`
	JointUserId      *pgtype.UUID        `db:"joint_user_id"`
}

func (repo *Repository) GetCategory(ctx context.Context, id pgtype.UUID) (*Category, error) {
	result := &Category{}
	statement := `
		SELECT id, user_id, name, parent_category_id, maximum, cycle_type_id, rollover, joint_user_id, created_at, updated_at
		FROM categories
		WHERE id = $1;
	`
	err := repo.Pool.GetContext(ctx, result, statement, helpers.UUIDToString(id))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TODO: Finish this
// func (repo *Repository) GetCategories() {}

func (repo *Repository) CreateCategory(ctx context.Context, category *CategoryCreateRequest) (*Category, error) {
	if category.CycleTypeID == nil {
		defaultCycleType, err := repo.GetDefaultCycleType(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get default cycle_type: %w", err)
		}
		category.CycleTypeID = &defaultCycleType.ID
	}
	result := &Category{}
	query := `
		INSERT INTO categories
			(user_id, name, parent_category_id, maximum, cycle_type_id, rollover, joint_user_id)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
		RETURNING
			id, user_id, name, parent_category_id, maximum, cycle_type_id, rollover, joint_user_id, created_at, updated_at
	`
	err := repo.Pool.QueryRowxContext(
		ctx,
		query,
		category.UserID, category.Name, category.ParentCategoryID, category.Maximum, category.CycleTypeID, category.Rollover, category.JointUserID,
	).StructScan(result)
	if err != nil {
		return nil, fmt.Errorf("failed to create category for user %s: %w", helpers.UUIDToString(category.UserID), err)
	}
	return result, nil
}

// TODO: Implement Update Category
// func (repo *Repository) UpdateCategory(ctx context.Context, category Category) (Category, error) {}
