package entity

import "unicode/utf8"

type Restaurant struct {
	Name    string
	Address string
	Photo   string
	URL     string
}

func (restaurant *Restaurant) AdjustAddress() {
	if 60 < utf8.RuneCountInString(restaurant.Address) {
		restaurant.Address = string([]rune(restaurant.Address)[:60])
	}
}
