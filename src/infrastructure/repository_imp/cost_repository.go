package repository_imp

import (
	"errors"
	"fmt"
	"household.api/src/domain/entity"
	"household.api/src/infrastructure"
	"log"
)

type CostRepository struct {
	database *infrastructure.Database
}

func NewCostRepository(database *infrastructure.Database) *CostRepository {
	cost := new(CostRepository)
	cost.database = database
	return cost
}

type ReceiveUserId struct {
	Id int `db:"id"`
}

func (repo *CostRepository) Save(cost *entity.Cost) (costId uint32, err error) {
	db, err := repo.database.Connect()
	if err != nil {
		return
	}
	//public_user_idより、user_idを取得
	query := `SELECT id FROM users WHERE public_user_id = ? LIMIT 1`
	var receiveVar ReceiveUserId
	err = db.Get(&receiveVar, query, cost.PublicUserId)
	if err != nil {
		return
	}
	userId := uint32(receiveVar.Id)

	query = `INSERT INTO costs (public_cost_id, user_id, title, outcome) VALUE (?,?,?,?)`
	result := db.MustExec(query, cost.PublicCostId, userId, cost.Title, cost.OutCome)
	resultNum, err := result.LastInsertId()
	if resultNum == 0 {
		fmt.Println("nothing affected")
		return
	}
	costId = uint32(resultNum)
	return
}
func (repo *CostRepository) Delete(costId int) (err error) {
	db, err := repo.database.Connect()
	if err != nil {
		return
	}
	query := `DELETE FROM costs WHERE id = ?`
	result := db.MustExec(query, costId)
	resultNum, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
		if resultNum == 0 {
			err = errors.New("データ削除に失敗しました。")
			return
		}
		return
	}
	return
}
func (repo *CostRepository) Update() {}
