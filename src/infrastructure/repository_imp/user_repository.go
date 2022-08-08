package repository_imp

import (
	"fmt"
	"household.api/src/domain/entity"
	"household.api/src/infrastructure"
)

type UserRepository struct {
	database *infrastructure.Database
}

func NewUserRepository(database *infrastructure.Database) *UserRepository {
	repo := new(UserRepository)
	repo.database = database
	return repo
}

func (repo *UserRepository) Save(user *entity.User) (err error) {
	db, err := repo.database.Connect()
	if err != nil {
		return
	}
	query := `INSERT INTO users (public_user_id, line_user_id, image, name) VALUE (?,?,?,?)`
	result := db.MustExec(query, user.PublicUserId, user.LineUserId, user.Image, user.Name)
	resultNum, err := result.RowsAffected()
	if err != nil {
		return
	}
	if resultNum == 0 {
		fmt.Println("nothing affected")
		return
	}
	return
}

func (repo *UserRepository) Delete() {}
func (repo *UserRepository) Update() {}
