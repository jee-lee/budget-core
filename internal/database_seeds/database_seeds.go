package database_seeds

import (
	"context"
	"errors"
	"fmt"
	"github.com/jee-lee/budget-core/internal/category/repository"
	"github.com/jee-lee/budget-core/internal/config"
	"os"
)

// Refresh truncates, then loads all tables with data
func Refresh() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv != "DEVELOPMENT" {
		config.Logger.Fatal(
			errors.New(
				"seedDatabase is available only in DEVELOPMENT. " +
					"Make sure you have APP_ENV set to DEVELOPMENT",
			).Error(),
		)
	}

	db := config.GetDB()
	repo := repository.NewRepository(db)

	tables := []string{
		"transactions",
		"categories",
		"accounts",
		"transaction_types",
		"account_types",
		"cycle_types",
	}
	for _, t := range tables {
		statement := fmt.Sprintf("TRUNCATE %v CASCADE", t)
		_, err := db.Exec(statement)
		if err != nil {
			config.Logger.Fatal(err.Error())
		}
	}

	err := repo.CreateCycleTypes(context.Background())
	if err != nil {
		config.Logger.Fatal(err.Error())
	}
}
