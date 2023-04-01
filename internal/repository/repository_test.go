package repository_test

import (
	"context"
	"github.com/jee-lee/budget-core/internal/repository"
	"github.com/jee-lee/budget-core/internal/test"
	"testing"
)

var (
	testDB *repository.Repository
)

func setupTest(t *testing.T) {
	t.Helper()
	var err error
	testDB, err = test.GetDBInstance()
	if err != nil {
		t.Fatal(err.Error())
	}
	test.RemoveCategoriesAndCycleTypes(t, testDB)
	if err = testDB.CreateCycleTypes(context.TODO()); err != nil {
		t.Fatal("could not seed cycle types in database " + err.Error())
	}
}

func teardownTest(t *testing.T) {
	test.RemoveCategoriesAndCycleTypes(t, testDB)
}
