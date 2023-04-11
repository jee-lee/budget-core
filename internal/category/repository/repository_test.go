package repository_test

import (
	. "github.com/jee-lee/budget-core/internal/category/repository"
	"github.com/jee-lee/budget-core/internal/test"
	_ "github.com/lib/pq"
	"testing"
)

var (
	repo Repository
)

func setupTest(t *testing.T) {
	t.Helper()
	testRepo, err := test.GetDBInstance()
	if err != nil {
		t.Fatal(err.Error())
	}
	repo = testRepo.Category
	test.RemoveCategories(t)
	test.RemoveUsers(t)
}

func teardownTest(t *testing.T) {
	test.RemoveCategories(t)
	test.RemoveUsers(t)
}
