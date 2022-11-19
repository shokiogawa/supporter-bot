package front_controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"household.api/src/usecase/command"
	"household.api/src/usecase/query/query_service_interface"
	"household.api/src/view_model"
	"log"
)

type CostController struct {
	saveCostUseCase  command.SaveCostUseCase
	costQueryService query_service_interface.CostQueryService
}

func NewCostController(saveCostUseCase command.SaveCostUseCase, costQueryService query_service_interface.CostQueryService) *CostController {
	con := new(CostController)
	con.saveCostUseCase = saveCostUseCase
	con.costQueryService = costQueryService
	return con
}

func (con *CostController) CostPerDay(e echo.Context) (err error) {
	publicUserId := e.QueryParam("publicUserId")
	costs, err := con.costQueryService.FetchPerDay(publicUserId)
	if err != nil {
		log.Fatal(err)
		e.Error(err)
	}

	costPerDayViewModel := make([]view_model.CostPerDayViewModel, 0)
	for _, cost := range costs {
		costPerDayViewModel = append(costPerDayViewModel, view_model.CostPerDayViewModel{
			Title:   cost.Title,
			OutCome: cost.OutCome,
		})
	}
	costPerDayListViewModel := &view_model.CostPerDayListViewModel{
		CostPerDay: costPerDayViewModel,
	}
	//jsonにパースし、writeに書き込みフロントに返す。
	err = json.NewEncoder(e.Response().Writer).Encode(costPerDayListViewModel)
	if err != nil {
		e.Error(err)
		return
	}
	return
}
