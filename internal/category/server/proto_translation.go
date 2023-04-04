package server

import (
	"context"
	"github.com/jee-lee/budget-core/internal/category/repository"
	pb "github.com/jee-lee/budget-core/rpc/category"
	"time"
)

func (s *Server) makeCategoryResponse(ctx context.Context, c *repository.Category) (*pb.Category, error) {
	parentCategoryID := ""
	if c.ParentCategoryID != nil {
		parentCategoryID = c.ParentCategoryID.String()
	}
	maximum := 0.0
	if c.Maximum != nil {
		maximum = *c.Maximum
	}
	jointUserID := ""
	if c.JointUserID != nil {
		jointUserID = c.JointUserID.String()
	}
	cycleType, err := s.Repo.GetCycleTypeByID(ctx, c.CycleTypeID)
	if err != nil {
		return nil, err
	}
	return &pb.Category{
		Id:               c.ID.String(),
		UserId:           c.UserID.String(),
		Name:             c.Name,
		ParentCategoryId: parentCategoryID,
		Maximum:          maximum,
		CycleType:        cycleType.Name,
		Rollover:         c.Rollover,
		JointUserId:      jointUserID,
		CreatedAt:        c.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        c.UpdatedAt.Format(time.RFC3339),
	}, nil
}
