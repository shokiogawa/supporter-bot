package controller

import (
	"fmt"
	"household.api/src/usecase/command"
	"household.api/src/usecase/query/query_service_interface"
	"strconv"
	"strings"
	"time"
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

func (con *CostController) CostPerDay(lineUserId string) (replyMessage string, err error) {
	costs, err := con.costQueryService.FetchPerDay(lineUserId)
	replyMessage = "今日のこれまでの支出"
	totalOutCome := 0
	for _, cost := range costs {
		costString := strconv.Itoa(cost.OutCome)
		replyMessage = replyMessage + "\n" + cost.Title + ":" + costString
		totalOutCome = totalOutCome + cost.OutCome
	}
	replyMessage = replyMessage + "\n" + "----------------------" + "\n" + "合計:" + strconv.Itoa(totalOutCome) + "円"
	return
}

func (con *CostController) CostPerMonth(lineUserId string) (replyMessage string, err error) {
	costLists, err := con.costQueryService.FetchPerMonth(lineUserId)
	today := time.Now()
	year := strconv.Itoa(today.Year())
	month := strconv.Itoa(int(today.Month()))
	replyMessage = year + "年" + month + "月" + "の合計支出は。\n"
	totalCost := 0
	for _, cost := range costLists {
		replyMessage = replyMessage + cost.Date + ":" + strconv.Itoa(cost.OutCome) + "円" + "\n"
		totalCost += cost.OutCome
	}
	replyMessage += "合計支出:" + strconv.Itoa(totalCost) + "円"
	//costSumString := strconv.Itoa(costSum)
	//replyMessage = replyMessage + "\n" + costSumString + "円" + "\n" + "です。"
	return
}
