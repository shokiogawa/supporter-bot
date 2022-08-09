package controller

import (
	"fmt"
	"household.api/src/usecase/command"
	"household.api/src/usecase/query/query_service_interface"
	"strconv"
	"strings"
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

func (con *CostController) SaveCost(message string, userId string) (replyMessage string, err error) {
	content := strings.Split(message, ":")
	title := content[0]
	outcome, err := strconv.Atoi(content[1])

	err = con.saveCostUseCase.Invoke(title, outcome, userId)
	if err != nil {
		fmt.Println(err)
		return
	}
	replyMessage = title + "を" + content[1] + "で登録しました。"
	return
}

func (con *CostController) ListCost(lineUserId string) (replyMessage string, err error) {
	costs, err := con.costQueryService.FetchList(lineUserId)
	replyMessage = "今日のこれまでの支出"
	totalOutCome := 0
	for _, cost := range costs {
		costString := strconv.Itoa(cost.OutCome)
		replyMessage = replyMessage + "\n" + cost.Title + ":" + costString
		totalOutCome = totalOutCome + cost.OutCome
	}
	replyMessage = replyMessage + "\n" + "----------------------" + "\n" + "合計:" + strconv.Itoa(totalOutCome)
	return
}
