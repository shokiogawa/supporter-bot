package controller

import (
	"fmt"
	"household.api/src/usecase/command"
	"strconv"
	"strings"
)

type CostController struct {
	saveCostUseCase command.SaveCostUseCase
}

func NewCostController(saveCostUseCase command.SaveCostUseCase)*CostController {
	con := new(CostController)
	con.saveCostUseCase = saveCostUseCase
	return con
}

func (con *CostController) SaveCost(message string, userId string)(replyMessage string , err error){
	content := strings.Split(message, ":")
	title := content[0]
	outcome, err := strconv.Atoi(content[1])

	err = con.saveCostUseCase.Invoke(title,outcome,userId)
	if err != nil{
		fmt.Println(err)
		return
	}
	replyMessage = title + "を" + content[1] + "で登録しました。"
	return
}
