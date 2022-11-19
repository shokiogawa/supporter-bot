package query_service_interface

import "household.api/src/domain/entity"

type FixedCostQueryService interface {
	GetFixedCostList(publicUserId string) (fixedCostList []*entity.FixedCost, err error)
}
