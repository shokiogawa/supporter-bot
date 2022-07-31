package entity

import (
	"errors"
	"github.com/google/uuid"
)

type Cost struct{
	Id uint32
	PublicFixedCostId uuid.UUID
	Title string
	OutCome int
	UserId string
}

func NewCost(title string , outCome int, userId string) (cost *Cost, err error){
	if title == "" || outCome == 0{
		err = errors.New("タイトルもしくは金額が入力されていません。")
		return
	}
	cost = new(Cost)
	cost.PublicFixedCostId = uuid.New()
	cost.Title = title
	cost.OutCome = outCome
	cost.UserId = userId
	return
}
