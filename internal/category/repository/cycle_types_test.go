package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_CycleTypes(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	t.Run("should have the expected states", func(t *testing.T) {
		expected := []string{"weekly", "monthly", "quarterly", "semiannually", "annually"}
		assert.Equal(t, expected, CycleTypes())
	})
}

func TestRepository_GetCycleTypeByName(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	t.Run("should return the correct CycleType", func(t *testing.T) {
		cycleType, err := testDB.GetCycleTypeByName(context.Background(), "weekly")
		assert.NoError(t, err)
		assert.Equal(t, "weekly", cycleType.Name)
	})

	t.Run("should return an error if cycle type is not found", func(t *testing.T) {
		cycleType, err := testDB.GetCycleTypeByName(context.Background(), "NOT_FOUND")
		assert.Error(t, err)
		assert.Nil(t, cycleType)
	})
}

func TestRepository_GetCycleTypeByID(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	t.Run("should return the correct CycleType", func(t *testing.T) {
		cycleType, err := testDB.GetCycleTypeByID(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, "weekly", cycleType.Name)
	})

	t.Run("should return an error if cycle type is not found", func(t *testing.T) {
		cycleType, err := testDB.GetCycleTypeByID(context.Background(), 9999)
		assert.Error(t, err)
		assert.Nil(t, cycleType)
	})
}

func TestRepository_GetDefaultCycleType(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	t.Run("should return monthly", func(t *testing.T) {
		cycleType, err := testDB.GetDefaultCycleType(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, "monthly", cycleType.Name)
	})
}
