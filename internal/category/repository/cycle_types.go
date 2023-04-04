package repository

import "context"

type CycleType struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

// CycleTypes The allowable cycle types a budget category may have
func CycleTypes() []string {
	return []string{"weekly", "monthly", "quarterly", "semiannually", "annually"}
}

func (repo *repository) GetCycleTypeByName(ctx context.Context, name string) (*CycleType, error) {
	result := &CycleType{}
	if ct, ok := repo.CycleTypeByNameCache[name]; ok {
		return ct, nil
	}
	statement := `SELECT id, name FROM cycle_types WHERE name = $1;`
	err := repo.Pool.GetContext(ctx, result, statement, name)
	if err != nil {
		return nil, err
	}
	repo.cacheCycleType(result)
	return result, nil
}

func (repo *repository) GetCycleTypeByID(ctx context.Context, id int) (*CycleType, error) {
	result := &CycleType{}
	if ct, ok := repo.CycleTypeByIdCache[id]; ok {
		return ct, nil
	}
	statement := `SELECT id, name FROM cycle_types WHERE id = $1;`
	err := repo.Pool.GetContext(ctx, result, statement, id)
	if err != nil {
		return nil, err
	}
	repo.cacheCycleType(result)
	return result, nil
}

func (repo *repository) GetDefaultCycleType(ctx context.Context) (*CycleType, error) {
	return repo.GetCycleTypeByName(ctx, "monthly")
}

func (repo *repository) CreateCycleTypes(ctx context.Context) error {
	var result []CycleType
	statement := `INSERT INTO cycle_types (name) VALUES ($1), ($2), ($3), ($4), ($5)`
	types := CycleTypes()
	err := repo.Pool.SelectContext(ctx, &result, statement, types[0], types[1], types[2], types[3], types[4])
	return err
}

func (repo *repository) cacheCycleType(cycleType *CycleType) {
	repo.CycleTypeByIdCache[cycleType.ID] = cycleType
	repo.CycleTypeByNameCache[cycleType.Name] = cycleType
}
