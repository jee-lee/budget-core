package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"sync"
	"testing"
)

var (
	oneTestDbPool sync.Once
	pgdb          *sqlx.DB
	testDB        Repository
)

const testDatabaseUrl = "postgres://postgres:postgres@127.0.0.1:9995/budget_test?sslmode=disable"

func setupTest(t *testing.T) {
	t.Helper()
	var err error
	testDB, err = getDBInstance()
	if err != nil {
		t.Fatal(err.Error())
	}
	removeCategories(t)
}

func teardownTest(t *testing.T) {
	removeCategories(t)
}

func getDBInstance() (Repository, error) {
	if pgdb == nil {
		var err error
		oneTestDbPool.Do(
			func() {
				var openErr error
				pgdb, openErr = sqlx.Open("postgres", testDatabaseUrl)
				if openErr != nil {
					err = fmt.Errorf("could not create database handle: %s", openErr.Error())
					return
				}
				if pingErr := pgdb.Ping(); pingErr != nil {
					err = fmt.Errorf("could not connect to database: %s", pingErr.Error())
					return
				}
				testDB = NewRepository(*pgdb)
			},
		)
		if err != nil {
			return nil, err
		}
	}
	return testDB, nil
}

func removeCategories(t testing.TB) {
	t.Helper()
	if err := pgdb.Get(&struct{}{}, "DELETE FROM categories"); err != nil && err != sql.ErrNoRows {
		t.Fatal("Failed to remove old categories from database: " + err.Error())
	}
}
