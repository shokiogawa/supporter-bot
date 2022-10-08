package repository_imp

import (
	"fmt"
	"household.api/src/domain/entity"
	"household.api/src/infrastructure"
)

type FixedCostRepository struct {
	database *infrastructure.Database
}

func NewFixedCostRepository(database *infrastructure.Database) *FixedCostRepository {
	repo := new(FixedCostRepository)
	repo.database = database
	return repo
}

func (repo *FixedCostRepository) Save(cost *entity.Cost) {
	db, err := repo.database.Connect()
	if err != nil {
		return
	}

	//line_user_idより、user_idを取得
	query := `SELECT id FROM users WHERE line_user_id = ?`
	var receiveVar ReceiveUserId
	err = db.Get(&receiveVar, query, cost.UserLineId)
	if err != nil {
		return
	}
	userId := uint32(receiveVar.Id)
	fmt.Print(userId)
}
