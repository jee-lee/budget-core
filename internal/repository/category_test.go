package repository_test

import (
	"context"
	"github.com/BudjeeApp/budget-core/internal/helpers"
	"github.com/BudjeeApp/budget-core/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCategory(t *testing.T) {
	assert := assert.New(t)
	setupTest(t)
	defer teardownTest(t)

	t.Run("should create a category with minimum fields", func(t *testing.T) {
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
		t.Skip("Pending")
	})

	t.Run("should require an existing CategoryID for the ParentCategoryID", func(t *testing.T) {
		t.Skip("Pending")
	})

	t.Run("should require an existing UserID for the JointUserID", func(t *testing.T) {
		t.Skip("Pending")
	})

	t.Run("should default to \"monthly\" cycle_type if CycleTypeID is not populated", func(t *testing.T) {
		t.Skip("Pending")
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
