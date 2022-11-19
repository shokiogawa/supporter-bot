package repository_imp

import (
	"errors"
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

func (repo *FixedCostRepository) Save(fixedCost *entity.FixedCost) (err error) {
	db, err := repo.database.Connect()
	if err != nil {
		return
	}

	//line_user_idより、user_idを取得
	query := `SELECT id FROM users WHERE public_user_id = ?`
	var receiveVar ReceiveUserId
	err = db.Get(&receiveVar, query, fixedCost.PublicUserId)
	if err != nil {
		return
	}
	userId := uint32(receiveVar.Id)

	query = `INSERT INTO fixed_costs (public_fixed_cost_id, user_id, name) VALUE (?,?,?)`
	result, err := db.MustExec(query, fixedCost.PublicFixedCostId, userId, fixedCost.Name).RowsAffected()
	if err != nil {
		return
	}
	if result == 0 {
		err = errors.New("データが保存されていません。")
		return
	}
	return
}
