package repository_test

import (
	"context"
	"github.com/BudjeeApp/budget-core/internal/helpers"
	"github.com/BudjeeApp/budget-core/internal/repository"
	"github.com/BudjeeApp/budget-core/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategory(t *testing.T) {
	assert := assert.New(t)
	db, err := test.GetDBInstance()
	if err != nil {
		t.Fatal(err.Error())
	}
	test.RemoveCategoriesAndCycleTypes(t, db)
	defer test.RemoveCategoriesAndCycleTypes(t, db)
	if err = db.CreateCycleTypes(context.TODO()); err != nil {
		t.Fatal("could not seed cycle types in database " + err.Error())
	}

	t.Run("CreateCategory should create a category with minimum fields", func(t *testing.T) {
		req := &repository.CategoryCreateRequest{
			UserId: helpers.StringToUUID("3fdfb7d8-2f30-48dc-812a-73335363cf9a"),
			Name:   "testName",
		}
		category, err := db.CreateCategory(context.Background(), req)
		assert.NoError(err)
		assert.Equal(req.Name, category.Name)
		assert.Equal(req.UserId, category.UserId)
	})

	t.Run("CreateCategory should create a category with all fields", func(t *testing.T) {

	})
}
