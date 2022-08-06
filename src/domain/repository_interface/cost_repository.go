package repository_interface

import "household.api/src/domain/entity"

type CostRepository interface {
	Save(cost *entity.Cost) (err error)
	Update()
	Delete()
}
