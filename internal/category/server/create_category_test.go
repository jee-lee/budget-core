package server

import (
	"context"
	"database/sql"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/jee-lee/budget-core/internal/category/mocks"
	"github.com/jee-lee/budget-core/internal/category/repository"
	pb "github.com/jee-lee/budget-core/rpc/category"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var (
	successfulCategory = &repository.Category{
		ID:     uuid.MustParse("13a6682f-795c-49c1-bfbb-f94f4b770eef"),
		UserID: uuid.MustParse("2b807819-078c-4d0d-b2b3-6204ff95f967"),
		Name:   "Successful Category",
	}
	defaultCycleType = &repository.CycleType{
		ID:   2,
		Name: "monthly",
	}
)

func TestServer_CreateCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepository(ctrl)
	defer ctrl.Finish()

	server := NewServer(mockRepo)
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
			GetDefaultCycleType(gomock.Any()).
			Return(defaultCycleType, nil).
			Times(1)
		mockRepo.
			EXPECT().
			CreateCategory(gomock.Any(), gomock.Any()).
			Return(successfulCategory, nil).
			Times(1)
		mockRepo.
			EXPECT().
			GetCycleTypeByID(gomock.Any(), gomock.Any()).
			Return(defaultCycleType, nil).
			Times(1)
		req := &pb.CreateCategoryRequest{
			Name:             "Some Test Name",
			ParentCategoryId: "dd684402-9638-4576-9fdb-823688f44ff9",
			Maximum:          600.00,
		}
		resp, err := client.CreateCategory(context.Background(), req)
		assert.NoError(t, err)
		assert.Equal(t, successfulCategory.ID.String(), resp.Id)

	})

	t.Run("should respond with the correct cycle type if cycle type is given", func(t *testing.T) {
		var (
			expectedCycleType = &repository.CycleType{
				ID:   3,
				Name: "quarterly",
			}
			expectedCategory = &repository.Category{
				ID:               uuid.New(),
				UserID:           uuid.New(),
				Name:             "Some Test Name",
				ParentCategoryID: nil,
				Maximum:          nil,
				CycleTypeID:      3,
				Rollover:         false,
				JointUserID:      nil,
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			}
		)

		mockRepo.
			EXPECT().
			GetCycleTypeByName(gomock.Any(), "quarterly").
			Return(expectedCycleType, nil).
			Times(1)
		mockRepo.
			EXPECT().
			CreateCategory(gomock.Any(), gomock.Any()).
			Return(expectedCategory, nil).
			Times(1)
		mockRepo.
			EXPECT().
			GetCycleTypeByID(gomock.Any(), 3).
			Return(expectedCycleType, nil).
			Times(1)
		req := &pb.CreateCategoryRequest{
			Name:      "Some Test Name",
			CycleType: "quarterly",
		}
		resp, err := client.CreateCategory(context.Background(), req)
		assert.NoError(t, err)
		assert.Equal(t, "quarterly", resp.CycleType)
	})

	t.Run("should be able to accept only the category name in the request", func(t *testing.T) {
		mockRepo.
			EXPECT().
			GetDefaultCycleType(gomock.Any()).
			Return(defaultCycleType, nil).
			Times(1)
		mockRepo.
			EXPECT().
			CreateCategory(gomock.Any(), gomock.Any()).
			Return(successfulCategory, nil).
			Times(1)
		mockRepo.
			EXPECT().
			GetCycleTypeByID(gomock.Any(), gomock.Any()).
			Return(defaultCycleType, nil).
			Times(1)
		req := &pb.CreateCategoryRequest{
			Name: "Successful Category",
		}
		_, err := client.CreateCategory(context.Background(), req)
		assert.NoError(t, err)
	})

	t.Run("should use default cycleType when cycleType is empty in the request", func(t *testing.T) {
		mockRepo.
			EXPECT().
			GetDefaultCycleType(gomock.Any()).
			Return(defaultCycleType, nil).
			Times(1)
		mockRepo.
			EXPECT().
			CreateCategory(gomock.Any(), gomock.Any()).
			Return(successfulCategory, nil).
			Times(1)
		mockRepo.
			EXPECT().
			GetCycleTypeByID(gomock.Any(), gomock.Any()).
			Return(defaultCycleType, nil).
			Times(1)
		req := &pb.CreateCategoryRequest{
			Name:             "Successful Category",
			ParentCategoryId: "1b227819-078c-4d0d-b2b3-6204ff95f967",
			Maximum:          100.00,
			Rollover:         false,
			JointUserId:      "13a6682f-795c-49c1-bfbb-f94f4b220eef",
		}
		resp, err := client.CreateCategory(context.Background(), req)
		assert.NoError(t, err)
		assert.Equal(t, defaultCycleType.Name, resp.CycleType)
	})

	repositoryErrorTestCases := []struct {
		TestName string
		RepoFunc func(repo *mocks.MockRepository)
		Expected string
	}{
		{
			TestName: "should return an invalid_argument error when there is an invalid cycle_type in the request",
			RepoFunc: func(repo *mocks.MockRepository) {
				repo.EXPECT().GetCycleTypeByName(gomock.Any(), gomock.Any()).Return(nil, sql.ErrNoRows).Times(1)
			},
			Expected: "invalid_argument",
		},
		{
			TestName: "should return an internal error when there is an invalid cycle_type in the request",
			RepoFunc: func(repo *mocks.MockRepository) {
				repo.EXPECT().GetCycleTypeByName(gomock.Any(), gomock.Any()).Return(nil, sql.ErrConnDone).Times(1)
			},
			Expected: "internal",
		},
		{
			TestName: "should return an internal error when there is an invalid cycle_type in the request",
			RepoFunc: func(repo *mocks.MockRepository) {
				repo.EXPECT().GetCycleTypeByName(gomock.Any(), gomock.Any()).Return(&repository.CycleType{
					ID:   1,
					Name: "weekly",
				}, nil).Times(1)
				repo.EXPECT().CreateCategory(gomock.Any(), gomock.Any()).Return(nil, sql.ErrConnDone).Times(1)
			},
			Expected: "internal",
		},
	}
	for _, tc := range repositoryErrorTestCases {
		t.Run(tc.TestName, func(t *testing.T) {
			tc.RepoFunc(mockRepo)
			req := &pb.CreateCategoryRequest{
				Name:      "Some Category",
				CycleType: "Something",
			}
			resp, err := client.CreateCategory(context.Background(), req)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tc.Expected)
			assert.Nil(t, resp)

		})
	}

	invalidArgumentTestCases := []struct {
		TestName              string
		CreateCategoryRequest *pb.CreateCategoryRequest
	}{
		{
			TestName: "empty Name",
			CreateCategoryRequest: &pb.CreateCategoryRequest{
				Maximum: 550.00,
			},
		},
		{
			TestName: "invalid parentCategoryId",
			CreateCategoryRequest: &pb.CreateCategoryRequest{
				Name:             "Category",
				ParentCategoryId: "123",
			},
		},
		{
			TestName: "invalid jointUserId",
			CreateCategoryRequest: &pb.CreateCategoryRequest{
				Name:        "Category",
				JointUserId: "13a6682f-795c-49c1-bfbb-f94f40eef",
			},
		},
	}

	for _, tc := range invalidArgumentTestCases {
		t.Run("should return an invalid argument error for"+tc.TestName, func(t *testing.T) {
			mockRepo.
				EXPECT().
				CreateCategory(gomock.Any(), gomock.Any()).
				Times(0)
			req := tc.CreateCategoryRequest
			resp, err := client.CreateCategory(context.Background(), req)
			assert.Error(t, err, "expected an error")
			assert.Contains(t, err.Error(), "invalid_argument")
			assert.Nil(t, resp)
		})
	}
}
