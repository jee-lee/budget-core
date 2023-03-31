package repository_test

import (
	"context"
	"github.com/BudjeeApp/budget-core/internal/repository"
	"github.com/BudjeeApp/budget-core/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCycleTypes(t *testing.T) {
	db, err := test.GetDBInstance()
	if err != nil {
		t.Fatal(err.Error())
	}
	test.RemoveCategoriesAndCycleTypes(t, db)
	defer test.RemoveCategoriesAndCycleTypes(t, db)
	if err = db.CreateCycleTypes(context.TODO()); err != nil {
		t.Fatal("could not seed cycle types in database " + err.Error())
	}

	t.Run("CycleTypes should have the expected states", func(t *testing.T) {
		expected := []string{"weekly", "monthly", "quarterly", "semiannually", "annually"}
		assert.Equal(t, expected, repository.CycleTypes())
	})

	t.Run("GetCycleTypes should return the correct CycleType", func(t *testing.T) {
		result, err := db.GetCycleType(context.Background(), "weekly")
		assert.NoError(t, err)
		assert.Equal(t, "weekly", result.Name)
	})

	t.Run("GetCycleTypes should return an error if cycle type is not found", func(t *testing.T) {
		_, err := db.GetCycleType(context.Background(), "NOT_FOUND")
		assert.Error(t, err)
	})

	t.Run("GetDefaultCycleType should return monthly", func(t *testing.T) {
		result, err := db.GetDefaultCycleType(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, "monthly", result.Name)
	})
}
