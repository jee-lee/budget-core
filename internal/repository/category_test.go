package repository_test

import (
	"context"
	pgxdecimal "github.com/jackc/pgtype/ext/shopspring-numeric"
	"github.com/jee-lee/budget-core/internal/helpers"
	"github.com/jee-lee/budget-core/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCategory(t *testing.T) {
	assert := assert.New(t)

	t.Run("should create a category with minimum fields", func(t *testing.T) {
		setupTest(t)
		defer teardownTest(t)
		req := &repository.CategoryCreateRequest{
			UserID: helpers.StringToUUID("3fdfb7d8-2f30-48dc-812a-73335363cf9a"),
			Name:   "testName",
		}
		category, err := testDB.CreateCategory(context.Background(), req)
		assert.NoError(err)
		assert.Equal(req.Name, category.Name)
		assert.Equal(req.UserID, category.UserID)
	})

	t.Run("should create a category with all fields", func(t *testing.T) {
		setupTest(t)
		defer teardownTest(t)
		parentCategoryRequest := &repository.CategoryCreateRequest{
			UserID: helpers.StringToUUID("3fdfb7d8-2f30-48dc-812a-73335363cf9a"),
			Name:   "testName",
		}
		parentCategory, err := testDB.CreateCategory(context.Background(), parentCategoryRequest)
		assert.NoError(err)

		amount := &pgxdecimal.Numeric{}
		amount.Set(100.32)
		jointUserId := helpers.StringToUUID("9dd34e9a-449c-45be-9e9e-7ea7d99da8f4")
		cycleTypeId := 1
		subCategoryRequest := &repository.CategoryCreateRequest{
			UserID:           helpers.StringToUUID("3fdfb7d8-2f30-48dc-812a-73335363cf9a"),
			Name:             "Sub Category",
			ParentCategoryID: &parentCategory.ID,
			Maximum:          amount,
			CycleTypeID:      &cycleTypeId,
			Rollover:         true,
			JointUserID:      &jointUserId,
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
		nonexistentCategoryId := helpers.StringToUUID("fccc016e-4ded-4d23-b604-ac604a5d7d48")
		req := &repository.CategoryCreateRequest{
			UserID:           helpers.StringToUUID("503f2c89-1c40-4388-8852-9f712aa83066"),
			Name:             "Nonexistent Parent Category",
			ParentCategoryID: &nonexistentCategoryId,
		}
		category, err := testDB.CreateCategory(context.Background(), req)
		assert.Error(err)
		assert.Nil(category)
	})

	t.Run("should default to the default cycle_type if CycleTypeID is not populated", func(t *testing.T) {
		setupTest(t)
		defer teardownTest(t)
		defaultCycleType, err := testDB.GetDefaultCycleType(context.Background())
		if err != nil {
			t.Fatalf("test setup failure. error occurred when trying to get default cycle type: %s", err.Error())
		}
		req := &repository.CategoryCreateRequest{
			UserID: helpers.StringToUUID("3fdfb7d8-2f30-48dc-812a-73335363cf9a"),
			Name:   "testName",
		}
		category, err := testDB.CreateCategory(context.Background(), req)
		assert.NoError(err)
		assert.Equal(&defaultCycleType.ID, category.CycleTypeID)
	})

}

func TestGetCategory(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	t.Run("should retrieve the correct category", func(t *testing.T) {
		req := &repository.CategoryCreateRequest{
			UserID: helpers.StringToUUID("3fdfb7d8-2f30-48dc-812a-733353123f9a"),
			Name:   "testName",
		}
		createdCategory, err := testDB.CreateCategory(context.Background(), req)
		assert.NoError(t, err)

		retrievedCategory, err := testDB.GetCategory(context.Background(), createdCategory.ID)
		assert.NoError(t, err)
		assert.Equal(t, createdCategory.ID, retrievedCategory.ID)
	})

	t.Run("should return an error if category is not found", func(t *testing.T) {
		retrievedCategory, err := testDB.GetCategory(context.Background(), helpers.StringToUUID("3fdfb7d8-2f30-48dc-812a-733353123f9a"))
		assert.Nil(t, retrievedCategory)
		assert.Error(t, err)
	})
}
