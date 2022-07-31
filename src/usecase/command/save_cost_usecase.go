package command

import (
	"fmt"
	"household.api/src/domain/entity"
)

type SaveCostUseCase struct {

}

func NewSaveCostUseCase ()*SaveCostUseCase{
	usecase := new (SaveCostUseCase)
	return usecase
}

func (usecase *SaveCostUseCase) Invoke (title string, outcome int, userId string )(err error){
	cost, err := entity.NewCost(title, outcome, userId)
	if err != nil{
		return
	}
	//データを保存する。
	fmt.Println(cost)
	return err
}
