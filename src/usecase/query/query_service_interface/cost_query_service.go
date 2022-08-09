package query_service_interface

import "household.api/src/domain/entity"

type CostQueryService interface {
	FetchList(lineUserId string) (listCost []entity.Cost, err error)
}