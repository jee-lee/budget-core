package server

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/jee-lee/budget-core/internal/category/repository"
	"github.com/jee-lee/budget-core/internal/helpers"
	pb "github.com/jee-lee/budget-core/rpc/category"
	"github.com/twitchtv/twirp"
)

func (s *Server) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.Category, error) {
	if req.Name == "" {
		return nil, twirp.RequiredArgumentError("name")
	}
	parentCategoryId, err := helpers.GetUUID(req.GetParentCategoryId())
	if err != nil {
		return nil, twirp.InvalidArgumentError("parent_category_id", "invalid uuid")
	}
	maximum := req.GetMaximum()
	jointUserId, err := helpers.GetUUID(req.GetJointUserId())
	if err != nil {
		return nil, twirp.InvalidArgumentError("joint_user_id", "invalid uuid")
	}

	var cycleType *repository.CycleType
	if req.CycleType == "" {
		cycleType, err = s.Repo.GetDefaultCycleType(ctx)
		if err != nil {
			return nil, twirp.InternalError(InternalError)
		}
	} else {
		cycleType, err = s.Repo.GetCycleTypeByName(ctx, req.CycleType)
		if err == sql.ErrNoRows {
			return nil, twirp.InvalidArgumentError("cycle_type", "invalid cycle_type")
		} else if err != nil {
			return nil, twirp.InternalError(InternalError)
		}
	}
	repoCategoryCreateRequest := &repository.CategoryCreateRequest{
		UserID:           uuid.New(),
		Name:             req.Name,
		ParentCategoryID: parentCategoryId,
		Maximum:          &maximum,
		CycleTypeID:      cycleType.ID,
		Rollover:         req.Rollover,
		JointUserID:      jointUserId,
	}
	createdCategory, err := s.Repo.CreateCategory(ctx, repoCategoryCreateRequest)
	if err != nil {
		return nil, twirp.InternalError(InternalError)
	}
	categoryResponse, err := s.makeCategoryResponse(ctx, createdCategory)
	if err != nil {
		return nil, twirp.InternalError(InternalError)
	}
	return categoryResponse, nil
}
