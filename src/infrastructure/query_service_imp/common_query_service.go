package query_service_imp

import (
	"household.api/src/infrastructure"
)

type CommonQueryService struct {
	database *infrastructure.Database
}

func NewCommonQueryService(database *infrastructure.Database) *CommonQueryService {
	qs := new(CommonQueryService)
	qs.database = database
	return qs
}

type ReceivePublicUserId struct {
	PublicUserId string `db:"public_user_id"`
}

func (qs *CommonQueryService) GetPublicUserId(lineUserId string) (publicUserId string, err error) {
	db, err := qs.database.Connect()
	if err != nil {
		return
	}
	var receiveVar ReceivePublicUserId
	query := `SELECT public_user_id FROM users WHERE line_user_id = ?`
	err = db.Get(&receiveVar, query, lineUserId)
	if err != nil {
		return
	}
	publicUserId = receiveVar.PublicUserId
	return
}
