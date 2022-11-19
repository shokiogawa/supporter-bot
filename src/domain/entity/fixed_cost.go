package entity

import (
	"errors"
	"github.com/google/uuid"
)

type FixedCost struct {
	Id                uint32
	PublicFixedCostId uuid.UUID
	Name              string
	OutCome           int
	PublicUserId      string
}

func NewFixedCost(name string, outcome int, publicUserId string) (fixedCost *FixedCost, err error) {
	if name == "" {
		err = errors.New("固定費の名前がnullです")
		return
	}
	fixedCost = new(FixedCost)
	fixedCost.PublicFixedCostId = uuid.New()
	fixedCost.Name = name
	fixedCost.OutCome = outcome
	fixedCost.PublicUserId = publicUserId
	return
}
