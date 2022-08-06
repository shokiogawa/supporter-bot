package command

import (
	"fmt"
	"household.api/src/domain/entity"
	"household.api/src/domain/repository_interface"
)

type SaveCostUseCase struct {
	costRepository repository_interface.CostRepository
}

func NewSaveCostUseCase(costRepository repository_interface.CostRepository) *SaveCostUseCase {
	usecase := new(SaveCostUseCase)
	usecase.costRepository = costRepository
	return usecase
}

func (usecase *SaveCostUseCase) Invoke(title string, outcome int, userId string) (err error) {
	fmt.Println(userId)
	cost, err := entity.NewCost(title, outcome, userId)
	if err != nil {
		return
	}
	//データを保存する。
	fmt.Println(cost)
	return err
}
