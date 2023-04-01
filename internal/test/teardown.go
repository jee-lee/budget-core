package test

import (
	"context"
	"database/sql"
	"github.com/BudjeeApp/budget-core/internal/repository"
	"testing"
)

func RemoveCategoriesAndCycleTypes(t testing.TB, db *repository.Repository) {
	t.Helper()
	if err := db.Pool.Get(&struct{}{}, "DELETE from categories"); err != nil && err != sql.ErrNoRows {
		t.Fatal("Failed to remove old categories from database: " + err.Error())
	}
	var result []repository.CycleType
	statement := `DELETE FROM cycle_types WHERE name in ($1, $2, $3, $4, $5)`
	types := repository.CycleTypes()
	err := db.Pool.SelectContext(context.Background(), &result, statement, types[0], types[1], types[2], types[3], types[4])
	if err != nil {
		t.Fatal("failed to remove cycle types from database: " + err.Error())
	}
	// Reset the sequence
	_, err = db.Pool.ExecContext(context.Background(), "ALTER SEQUENCE public.cycle_types_id_seq RESTART WITH 1")
	if err != nil {
		t.Fatal("failed to reset cycle_types_id_seq: " + err.Error())
	}
}
