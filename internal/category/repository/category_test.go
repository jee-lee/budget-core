package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jee-lee/budget-core/internal/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_CreateCategory(t *testing.T) {
	assert := assert.New(t)

	t.Run("should create a category with minimum fields", func(t *testing.T) {
		setupTest(t)
		defer teardownTest(t)
		req := &CategoryCreateRequest{
			UserID:      uuid.New(),
			Name:        "testName",
			CycleTypeID: 1,
		}
		category, err := testDB.CreateCategory(context.Background(), req)
		assert.NoError(err)
		assert.Equal(req.Name, category.Name)
		assert.Equal(req.UserID, category.UserID)
	})

	t.Run("should create a category with all fields", func(t *testing.T) {
		setupTest(t)
		defer teardownTest(t)
		parentCategoryRequest := &CategoryCreateRequest{
			UserID:      uuid.New(),
			Name:        "testName",
			CycleTypeID: 1,
		}
		parentCategory, err := testDB.CreateCategory(context.Background(), parentCategoryRequest)
		assert.NoError(err)

		subCategoryRequest := &CategoryCreateRequest{
			UserID:           uuid.New(),
			Name:             "Sub Category",
			ParentCategoryID: &parentCategory.ID,
			Maximum:          helpers.Pointer(100.32),
			CycleTypeID:      1,
			Rollover:         true,
			JointUserID:      helpers.Pointer(uuid.New()),
		}
		subCategory, err := testDB.CreateCategory(context.Background(), subCategoryRequest)
		assert.NoError(err)
		assert.Equal(subCategoryRequest.Name, subCategory.Name)
		assert.Equal(subCategoryRequest.UserID, subCategory.UserID)
		assert.Equal(subCategoryRequest.ParentCategoryID, subCategory.ParentCategoryID)
		assert.Equal(subCategoryRequest.Maximum, subCategory.Maximum)
		assert.Equal(subCategoryRequest.CycleTypeID, subCategory.CycleTypeID)
		assert.Equal(subCategoryRequest.Rollover, subCategory.Rollover)
		assert.Equal(subCategoryRequest.JointUserID, subCategory.JointUserID)
	})

	t.Run("should require an existing CategoryID for the ParentCategoryID", func(t *testing.T) {
		setupTest(t)
		defer teardownTest(t)
		nonexistentCategoryId := uuid.New()
		req := &CategoryCreateRequest{
			UserID:           uuid.New(),
			Name:             "Nonexistent Parent Category",
			ParentCategoryID: &nonexistentCategoryId,
		}
		category, err := testDB.CreateCategory(context.Background(), req)
		assert.Error(err)
		assert.Nil(category)
	})
}

func TestRepository_GetCategory(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	t.Run("should retrieve the correct category", func(t *testing.T) {
		req := &CategoryCreateRequest{
			UserID:      uuid.New(),
			Name:        "testName",
			CycleTypeID: 1,
		}
		createdCategory, err := testDB.CreateCategory(context.Background(), req)
		assert.NoError(t, err)

		retrievedCategory, err := testDB.GetCategory(context.Background(), &createdCategory.ID)
		assert.NoError(t, err)
		assert.Equal(t, createdCategory.ID, retrievedCategory.ID)
	})

	t.Run("should return an error if category is not found", func(t *testing.T) {
		retrievedCategory, err := testDB.GetCategory(context.Background(), helpers.Pointer(uuid.New()))
		assert.Nil(t, retrievedCategory)
		assert.Error(t, err)
	})
}
