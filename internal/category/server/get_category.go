package server

import (
	"context"
	"database/sql"
	"github.com/jee-lee/budget-core/internal/helpers"
	pb "github.com/jee-lee/budget-core/rpc/category"
	"github.com/twitchtv/twirp"
)

func (s *Server) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.Category, error) {
	categoryId, err := helpers.GetUUID(req.GetCategoryId())
	if categoryId == nil {
		return nil, twirp.RequiredArgumentError("category_id")
	}
	if err != nil {
		return nil, twirp.InvalidArgumentError("category_id", "is an invalid uuid")
	}
	category, err := s.Repo.GetCategory(ctx, categoryId)
	if err == sql.ErrNoRows {
		return nil, twirp.NotFoundError("category not found")
	} else if err != nil {
		return nil, twirp.InternalError(InternalError)
	}
	categoryResponse, err := s.makeCategoryResponse(ctx, category)
	if err != nil {
		return nil, twirp.InternalError(InternalError)
	}
	return categoryResponse, nil
}
