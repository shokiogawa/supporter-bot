package repository_imp

import "household.api/src/infrastructure"

type CostRepository struct {
	database *infrastructure.Database
}

func NewCostRepository(database *infrastructure.Database) *CostRepository {
	cost := new(CostRepository)
	cost.database = database
	return cost
}

func (cost *CostRepository) Save()   {}
func (cost *CostRepository) Delete() {}
func (cost *CostRepository) Update() {}
