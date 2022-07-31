package query_service_interface

import "household.api/src/domain/entity"

type FetchWeatherQueryService interface {
	Invoke() (weather *entity.Weather, err error)
}
