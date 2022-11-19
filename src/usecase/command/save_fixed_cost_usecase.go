package command

import (
	"household.api/src/domain/entity"
	"household.api/src/domain/repository_interface"
)

type SaveFixedCostUseCase struct {
	fixedCostRepository repository_interface.FixedCostRepository
}

func NewSaveFixedCostUseCase(fixedCostRepository repository_interface.FixedCostRepository) *SaveFixedCostUseCase {
	uc := new(SaveFixedCostUseCase)
	uc.fixedCostRepository = fixedCostRepository
	return uc
}

func (usecase *SaveFixedCostUseCase) Invoke(name string, outcome int, publicUserId string) (err error) {
	fixedCost, err := entity.NewFixedCost(name, outcome, publicUserId)
	if err != nil {
		return
	}
	err = usecase.fixedCostRepository.Save(fixedCost)
	if err != nil {
		return
	}
	return
}
