package repository_interface

import "household.api/src/domain/entity"

type FixedCostRepository interface {
	Save(fixedCost *entity.FixedCost) (err error)
}
