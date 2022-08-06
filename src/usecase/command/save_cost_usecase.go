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

func (usecase *SaveCostUseCase) Invoke(title string, outcome int, userLineId string) (err error) {
	cost, err := entity.NewCost(title, outcome, userLineId)
	if err != nil {
		return
	}
	err = usecase.costRepository.Save(cost)
	if err != nil {
		fmt.Println(err)
		return
	}
	//データを保存する。
	return err
}
