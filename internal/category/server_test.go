package category_test

import (
	"context"
	"database/sql"
	"github.com/golang/mock/gomock"
	"github.com/jee-lee/budget-core/internal/category"
	"github.com/jee-lee/budget-core/internal/helpers"
	"github.com/jee-lee/budget-core/internal/mocks"
	"github.com/jee-lee/budget-core/internal/repository"
	pb "github.com/jee-lee/budget-core/rpc/category"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var successfulCategory = &repository.Category{
	ID:     helpers.StringToUUID("13a6682f-795c-49c1-bfbb-f94f4b770eef"),
	UserID: helpers.StringToUUID("2b807819-078c-4d0d-b2b3-6204ff95f967"),
	Name:   "Successful Category",
}

func TestServer_CreateCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepository(ctrl)
	defer ctrl.Finish()

	server := category.NewServer(mockRepo)
	twirpHandler := pb.NewCategoryServiceServer(server)
	testServer := httptest.NewServer(twirpHandler)
	defer testServer.Close()

	client := pb.NewCategoryServiceProtobufClient(testServer.URL, http.DefaultClient)

	t.Run("should create the category for the user who requested it", func(t *testing.T) {
		t.Skip("Pending. May need to move this test")
	})

	t.Run("should respond with the generated category uuid", func(t *testing.T) {
		mockRepo.
			EXPECT().
			CreateCategory(gomock.Any(), gomock.Any()).
			Return(successfulCategory, nil).
			Times(1)
		req := &pb.CreateCategoryRequest{
			Name:             "Some Test Name",
			ParentCategoryId: "dd684402-9638-4576-9fdb-823688f44ff9",
			Maximum:          600.00,
		}
		resp, err := client.CreateCategory(context.Background(), req)
		assert.NoError(t, err)
		assert.Equal(t, successfulCategory.ID, resp.Id)

	})

	t.Run("should be able to accept only the category name in the request", func(t *testing.T) {
		mockRepo.
			EXPECT().
			CreateCategory(gomock.Any(), gomock.Any()).
			Return(successfulCategory, nil).
			Times(1)
		req := &pb.CreateCategoryRequest{
			Name: "Successful Category",
		}
		_, err := client.CreateCategory(context.Background(), req)
		assert.NoError(t, err)
	})

	t.Run("should respond with a 500 if the repository returns an error", func(t *testing.T) {
		mockRepo.
			EXPECT().
			CreateCategory(gomock.Any(), gomock.Any()).
			Return(nil, sql.ErrConnDone).
			Times(1)
		req := &pb.CreateCategoryRequest{
			Name:             "Some Test Name",
			ParentCategoryId: "dd684402-9638-4576-9fdb-823688f44ff9",
			Maximum:          600.00,
		}
		resp, err := client.CreateCategory(context.Background(), req)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("should respond with 400 if there is no category name in the request", func(t *testing.T) {
		mockRepo.
			EXPECT().
			CreateCategory(gomock.Any(), gomock.Any()).
			Times(0)
		req := &pb.CreateCategoryRequest{
			Maximum: 550.00,
		}
		resp, err := client.CreateCategory(context.Background(), req)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}
