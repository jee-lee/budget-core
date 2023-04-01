package category

import (
	"context"
	"github.com/jee-lee/budget-core/internal/config"
	pb "github.com/jee-lee/budget-core/rpc/category"
	"go.uber.org/zap"
)

type Server struct {
	Repository Repository
	Logger     Logger
}

// TODO: Move this outside of category server
type Logger interface {
	Info(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
}

func NewServer(r Repository) *Server {
	return &Server{
		r,
		config.Logger,
	}
}

func (s *Server) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.Category, error) {
	return nil, nil
}

func (s *Server) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.Category, error) {
	return nil, nil
}
