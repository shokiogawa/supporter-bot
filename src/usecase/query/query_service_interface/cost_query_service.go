package query_service_interface

import (
	"household.api/src/domain/entity"
	"household.api/src/usecase/query/read_model"
)

type CostQueryService interface {
	FetchPerDay(lineUserId string) (listCost []entity.Cost, err error)
	FetchPerMonth(lineUserId string) (readModel []read_model.CostSumReamModel, err error)
	FetchPerMonthList() (readModel []read_model.CostSumListReadModel, err error)
}
