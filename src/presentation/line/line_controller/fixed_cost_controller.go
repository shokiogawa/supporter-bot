package line_controller

import (
	"household.api/src/usecase/command"
	"strconv"
	"strings"
)

type FixedCostController struct {
	saveFixedCostUseCase *command.SaveFixedCostUseCase
}

func NewFixedCostUseCase(saveFixedCostUseCase *command.SaveFixedCostUseCase) *FixedCostController {
	con := new(FixedCostController)
	con.saveFixedCostUseCase = saveFixedCostUseCase
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
