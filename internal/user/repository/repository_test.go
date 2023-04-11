package repository_test

import (
	"github.com/jee-lee/budget-core/internal/test"
	. "github.com/jee-lee/budget-core/internal/user/repository"
	_ "github.com/lib/pq"
	"testing"
)

var (
	repo Repository
)

func setupTest(t *testing.T) {
	t.Helper()
	testRepo, err := test.GetDBInstance()
	repo = testRepo.User
	if err != nil {
		t.Fatal(err.Error())
	}
	test.RemoveUsers(t)
}

func teardownTest(t *testing.T) {
	test.RemoveUsers(t)
}
