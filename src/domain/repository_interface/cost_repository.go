package repository_interface

import "household.api/src/domain/entity"

type CostRepository interface {
	Save(cost *entity.Cost) (costId uint32, err error)
	Update()
	Delete()
}
