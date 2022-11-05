package line_controller

import (
	"household.api/src/domain/entity"
	"household.api/src/usecase/query/query_service_interface"
)

type RestaurantController struct {
	fetchRestaurant query_service_interface.FetchRestaurantQueryService
}

func NewRestaurantController(fetchRestaurant query_service_interface.FetchRestaurantQueryService) *RestaurantController {
	con := new(RestaurantController)
	con.fetchRestaurant = fetchRestaurant
	return con
}

func (con *RestaurantController) GetRestaurant(lat string, lang string) (restaurants []*entity.Restaurant, err error) {
	restaurants, err = con.fetchRestaurant.Invoke(lat, lang)
	for _, restaurant := range restaurants {
		restaurant.AdjustAddress()
	}
	if err != nil {
		return
	}
	return
}
