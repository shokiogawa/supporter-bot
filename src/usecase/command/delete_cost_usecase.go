package command

import "household.api/src/domain/repository_interface"

type DeleteCostUseCase struct {
	costRepository repository_interface.CostRepository
}

func NewDeleteCostUseCase(costRepository repository_interface.CostRepository) *DeleteCostUseCase {
	usecase := new(DeleteCostUseCase)
	usecase.costRepository = costRepository
	return usecase
}

func (usecase *DeleteCostUseCase) Invoke(costId int) (err error) {
	err = usecase.costRepository.Delete(costId)
	if err != nil {
		return
	}
	return
}
