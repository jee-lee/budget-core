package repository

import (
	"context"
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
	removeCategoriesAndCycleTypes(t)
	if err = testDB.CreateCycleTypes(context.TODO()); err != nil {
		t.Fatal("could not seed cycle types in database " + err.Error())
	}
}

func teardownTest(t *testing.T) {
	removeCategoriesAndCycleTypes(t)
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
				testDB = NewRepository(pgdb)
			},
		)
		if err != nil {
			return nil, err
		}
	}
	return testDB, nil
}

func removeCategoriesAndCycleTypes(t testing.TB) {
	t.Helper()
	if err := pgdb.Get(&struct{}{}, "DELETE from categories"); err != nil && err != sql.ErrNoRows {
		t.Fatal("Failed to remove old categories from database: " + err.Error())
	}
	var result []CycleType
	statement := `DELETE FROM cycle_types WHERE name in ($1, $2, $3, $4, $5)`
	types := CycleTypes()
	err := pgdb.SelectContext(context.Background(), &result, statement, types[0], types[1], types[2], types[3], types[4])
	if err != nil {
		t.Fatal("failed to remove cycle types from database: " + err.Error())
	}
	// Reset the sequence
	_, err = pgdb.ExecContext(context.Background(), "ALTER SEQUENCE public.cycle_types_id_seq RESTART WITH 1")
	if err != nil {
		t.Fatal("failed to reset cycle_types_id_seq: " + err.Error())
	}
}
