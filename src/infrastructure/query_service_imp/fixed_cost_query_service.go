package query_service_imp

import (
	"household.api/src/domain/entity"
	"household.api/src/infrastructure"
)

type FixedCostQueryService struct {
	database *infrastructure.Database
}

func NewFixedCostQueryService(database *infrastructure.Database) *FixedCostQueryService {
	qs := new(FixedCostQueryService)
	qs.database = database
	return qs
}

type ReceiveFixedCostList struct {
	Id                int    `db:"id"`
	PublicFixedCostId string `db:"public_fixed_cost_id"`
	Name              string `db:"name"`
	OutCome           int    `db:"outcome"`
}

func (qs *FixedCostQueryService) GetFixedCostList(publicUserId string) (fixedCostList []*entity.FixedCost, err error) {
	db, err := qs.database.Connect()
	if err != nil {
		return
	}
	query := `SELECT id FROM users WHERE public_user_id = ?`
	var receiveVar ReceiveUserId
	err = db.Get(&receiveVar, query, publicUserId)
	if err != nil {
		return
	}
	query = `SELECT id, public_fixed_cost_id, name, outcome FROM fixed_costs WHERE user_id = ?`
	var fixedCostListVar []ReceiveFixedCostList
	err = db.Select(&fixedCostListVar, query, receiveVar.Id)
	if err != nil {
		return
	}

	for _, fixedCost := range fixedCostListVar {
		fixedCostList = append(fixedCostList, &entity.FixedCost{
			Id:      uint32(fixedCost.Id),
			Name:    fixedCost.Name,
			OutCome: fixedCost.OutCome,
		})
	}
	return
}
