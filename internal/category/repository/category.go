package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Category struct {
	ID               uuid.UUID  `db:"id"`
	UserID           uuid.UUID  `db:"user_id"`
	Name             string     `db:"name"`
	ParentCategoryID *uuid.UUID `db:"parent_category_id"`
	Maximum          *float64   `db:"maximum"`
	CycleTypeID      int        `db:"cycle_type_id"`
	Rollover         bool       `db:"rollover"`
	JointUserID      *uuid.UUID `db:"joint_user_id"`
	CreatedAt        time.Time  `db:"created_at"`
	UpdatedAt        time.Time  `db:"updated_at"`
}

type CategoryCreateRequest struct {
	UserID           uuid.UUID  `db:"user_id"`
	Name             string     `db:"name"`
	ParentCategoryID *uuid.UUID `db:"parent_category_id"`
	Maximum          *float64   `db:"maximum"`
	CycleTypeID      int        `db:"cycle_type_id"`
	Rollover         bool       `db:"rollover"`
	JointUserID      *uuid.UUID `db:"joint_user_id"`
}

func (repo *repository) GetCategory(ctx context.Context, id *uuid.UUID) (*Category, error) {
	result := &Category{}
	statement := `
		SELECT id, user_id, name, parent_category_id, maximum, cycle_type_id, rollover, joint_user_id, created_at, updated_at
		FROM categories
		WHERE id = $1;
	`
	err := repo.Pool.GetContext(ctx, result, statement, id.String())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *repository) CreateCategory(ctx context.Context, category *CategoryCreateRequest) (*Category, error) {
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
		return nil, fmt.Errorf("failed to create category for user %s: %w", category.UserID.String(), err)
	}
	return result, nil
}
