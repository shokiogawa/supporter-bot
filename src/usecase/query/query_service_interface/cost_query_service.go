package query_service_interface

import "household.api/src/domain/entity"

type CostQueryService interface {
	FetchPerDay(lineUserId string) (listCost []entity.Cost, err error)
	FetchPerMonth(lineUserId string) (costSum int, err error)
}
