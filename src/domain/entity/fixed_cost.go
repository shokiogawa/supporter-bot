package entity

import (
	"errors"
	"github.com/google/uuid"
)

type FixedCost struct {
	Id                uint32
	PublicFixedCostId uuid.UUID
	Name              string
	PublicUserId      string
}

func NewFixedCost(name string, publicUserId string) (fixedCost *FixedCost, err error) {
	if name == "" {
		err = errors.New("固定費の名前がnullです")
		return
	}
	fixedCost = new(FixedCost)
	fixedCost.PublicFixedCostId = uuid.New()
	fixedCost.Name = name
	fixedCost.PublicUserId = publicUserId
	return
}
