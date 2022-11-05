package front

import "household.api/src/presentation/front/front_controller"

type Handler struct {
	CostController *front_controller.CostController
}

func NewHandler(costController *front_controller.CostController) *Handler {
	handler := new(Handler)
	handler.CostController = costController
	return handler
}
