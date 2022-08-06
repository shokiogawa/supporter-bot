package repository_imp

import (
	"fmt"
	"household.api/src/domain/entity"
	"household.api/src/infrastructure"
)

type CostRepository struct {
	database *infrastructure.Database
}

func NewCostRepository(database *infrastructure.Database) *CostRepository {
	cost := new(CostRepository)
	cost.database = database
	return cost
}

func (repo *CostRepository) Save(cost *entity.Cost) (err error) {
	db, err := repo.database.Connect()
	if err != nil {
		return
	}
	query := `INSERT INTO costs (public_cost_id, user_id, title, outcome)`
	result := db.MustExec(query, cost.PublicFixedCostId, cost.UserId, cost.Title, cost.OutCome)
	resultNum, err := result.RowsAffected()
	if resultNum == 0 {
		fmt.Println("nothing affected")
		return
	}
	return

}
func (repo *CostRepository) Delete() {}
func (repo *CostRepository) Update() {}
