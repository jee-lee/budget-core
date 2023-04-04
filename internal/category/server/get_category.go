package server

import (
	"context"
	pb "github.com/jee-lee/budget-core/rpc/category"
)

func (s *Server) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.Category, error) {
	return nil, nil
}
