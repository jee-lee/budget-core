package test

import (
	"fmt"
	"github.com/jee-lee/budget-core/internal/repository"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"sync"
)

var (
	oneTestDbPool sync.Once
	pgdb          *repository.Repository
)

const testDatabaseUrl = "postgres://postgres:postgres@127.0.0.1:9995/budget_test?sslmode=disable"

func GetDBInstance() (*repository.Repository, error) {
	if pgdb == nil {
		var err error
		oneTestDbPool.Do(
			func() {
				var testDbPool *sqlx.DB
				var openErr error
				testDbPool, openErr = sqlx.Open("postgres", testDatabaseUrl)
				if openErr != nil {
					err = fmt.Errorf("could not create database handle: %s", openErr.Error())
					return
				}

				if pingErr := testDbPool.Ping(); pingErr != nil {
					err = fmt.Errorf("could not connect to database: %s", pingErr.Error())
					return
				}

				pgdb = repository.NewRepository(testDbPool)
			},
		)
		if err != nil {
			return nil, err
		}
	}
	return pgdb, nil
}
