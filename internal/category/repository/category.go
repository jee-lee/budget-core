package repository

import (
	"context"
	"database/sql"
	"fmt"
	pb "github.com/jee-lee/budget-core/rpc/category"
	"time"
)

type Category struct {
	ID               string         `db:"id"`
	UserID           string         `db:"user_id"`
	Name             string         `db:"name"`
	ParentCategoryID sql.NullString `db:"parent_category_id"`
	Allowance        int64          `db:"allowance"`
	CycleType        string         `db:"cycle_type"`
	Rollover         bool           `db:"rollover"`
	JointUserID      sql.NullString `db:"joint_user_id"`
	CreatedAt        time.Time      `db:"created_at"`
	UpdatedAt        time.Time      `db:"updated_at"`
}

type CategoryCreateRequest struct {
	UserID           string         `db:"user_id"`
	Name             string         `db:"name"`
	ParentCategoryID sql.NullString `db:"parent_category_id"`
	Allowance        int64          `db:"allowance"`
	CycleType        string         `db:"cycle_type"`
	Rollover         bool           `db:"rollover"`
	JointUserID      sql.NullString `db:"joint_user_id"`
}

func (c Category) ToProto() pb.Category {
	return pb.Category{
		Id:               c.ID,
		UserId:           c.UserID,
		Name:             c.Name,
		ParentCategoryId: c.ParentCategoryID.String,
		Allowance:        c.Allowance,
		CycleType:        cycleTypeToPB(c.CycleType),
		Rollover:         c.Rollover,
		JointUserId:      c.JointUserID.String,
		CreatedAt:        c.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        c.UpdatedAt.Format(time.RFC3339),
	}
}

func (repo repository) GetCategory(ctx context.Context, id string) (*Category, error) {
	result := &Category{}
	statement := `
		SELECT id, user_id, name, parent_category_id, allowance, cycle_type, rollover, joint_user_id, created_at, updated_at
		FROM categories
		WHERE id = $1;
	`
	err := repo.Pool.GetContext(ctx, result, statement, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo repository) CreateCategory(ctx context.Context, category CategoryCreateRequest) (*Category, error) {
	if category.ParentCategoryID.String == "" {
		category.ParentCategoryID.Valid = false
	}
	if category.JointUserID.String == "" {
		category.ParentCategoryID.Valid = false
	}
	result := &Category{}
	query := `
		INSERT INTO categories
			(user_id, name, parent_category_id, allowance, cycle_type, rollover, joint_user_id)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
		RETURNING
			id, user_id, name, parent_category_id, allowance, cycle_type, rollover, joint_user_id, created_at, updated_at
	`
	err := repo.Pool.QueryRowxContext(
		ctx,
		query,
		category.UserID, category.Name, category.ParentCategoryID, category.Allowance, category.CycleType, category.Rollover, category.JointUserID,
	).StructScan(result)
	if err != nil {
		return nil, fmt.Errorf("failed to create category for user %s: %w", category.UserID, err)
	}
	return result, nil
}

func cycleTypeToPB(cycleType string) pb.CycleType {
	switch cycleType {
	case "monthly":
		return pb.CycleType_monthly
	case "weekly":
		return pb.CycleType_weekly
	case "quarterly":
		return pb.CycleType_quarterly
	case "semiannually":
		return pb.CycleType_semiannually
	case "annually":
		return pb.CycleType_annually
	default:
		return pb.CycleType_monthly
	}
}
