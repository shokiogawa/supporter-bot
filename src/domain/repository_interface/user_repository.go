package repository_interface

import "household.api/src/domain/entity"

type UserRepository interface {
	Save(user *entity.User) (err error)
}
