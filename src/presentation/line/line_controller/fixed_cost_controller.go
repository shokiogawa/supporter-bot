package line_controller

import (
	"household.api/src/usecase/command"
	"household.api/src/usecase/query/query_service_interface"
	"strconv"
	"strings"
)

type FixedCostController struct {
	saveFixedCostUseCase  *command.SaveFixedCostUseCase
	fixedCostQueryService query_service_interface.FixedCostQueryService
}

func NewFixedCostUseCase(
	saveFixedCostUseCase *command.SaveFixedCostUseCase,
	fixedCostQueryService query_service_interface.FixedCostQueryService,
) *FixedCostController {
	con := new(FixedCostController)
	con.saveFixedCostUseCase = saveFixedCostUseCase
	con.fixedCostQueryService = fixedCostQueryService
	return con
}

func (con *FixedCostController) SaveFixedCost(message string, publicUserId string) (replyMessage string, err error) {
	content := strings.Split(message, ":")
	title := content[1]
	outcome, err := strconv.Atoi(content[2])

	err = con.saveFixedCostUseCase.Invoke(title, outcome, publicUserId)
	if err != nil {
		return
	}
	replyMessage = title + "を固定費に追加しました。"
	return
}

func (con *FixedCostController) GetFixedCostList(publicUserId string) (replyMessage string, err error) {
	fixedCostList, err := con.fixedCostQueryService.GetFixedCostList(publicUserId)
	if err != nil {
		return
	}
	var totalFixedCost int
	replyMessage = "現在の固定費" + "\n"
	for _, fixedCost := range fixedCostList {
		totalFixedCost += fixedCost.OutCome
		replyMessage = replyMessage + fixedCost.Name + ":" + strconv.Itoa(fixedCost.OutCome) + "円" + "\n"
	}
	replyMessage += "-------------------" + "\n"
	replyMessage += "合計:" + strconv.Itoa(totalFixedCost)
	return
}
