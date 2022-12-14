package entity

import (
	"errors"
	"github.com/google/uuid"
)

type Cost struct {
	Id           uint32
	PublicCostId uuid.UUID
	Title        string
	OutCome      int
	PublicUserId string
}

func NewCost(title string, outCome int, publicUserId string) (cost *Cost, err error) {
	if title == "" || outCome == 0 {
		err = errors.New("タイトルもしくは金額が入力されていません。")
		return
	}
	cost = new(Cost)
	cost.PublicCostId = uuid.New()
	cost.Title = title
	cost.OutCome = outCome
	cost.PublicUserId = publicUserId
	return
}
