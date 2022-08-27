package query_service_interface

import "household.api/src/domain/entity"

type FetchRestaurantQueryService interface {
	Invoke(lat string, lang string) (restaurants []*entity.Restaurant, err error)
}
