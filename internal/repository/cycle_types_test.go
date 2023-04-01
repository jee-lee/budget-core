package repository_test

import (
	"context"
	"github.com/BudjeeApp/budget-core/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCycleTypes(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	t.Run("should have the expected states", func(t *testing.T) {
		expected := []string{"weekly", "monthly", "quarterly", "semiannually", "annually"}
		assert.Equal(t, expected, repository.CycleTypes())
	})
}

func TestGetCycleTypes(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	t.Run("should return the correct CycleType", func(t *testing.T) {
		cycleType, err := testDB.GetCycleType(context.Background(), "weekly")
		assert.NoError(t, err)
		assert.Equal(t, "weekly", cycleType.Name)
	})

	t.Run("should return an error if cycle type is not found", func(t *testing.T) {
		cycleType, err := testDB.GetCycleType(context.Background(), "NOT_FOUND")
		assert.Error(t, err)
		assert.Nil(t, cycleType)
	})
}

func TestGetDefaultCycleType(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	t.Run("should return monthly", func(t *testing.T) {
		cycleType, err := testDB.GetDefaultCycleType(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, "monthly", cycleType.Name)
	})
}
