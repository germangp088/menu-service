package order

import (
	menu "menu-service/api/types/menu"
)

type Order struct {
	IDs []string `json:"ids"`
}

func (o Order) Calculate(m menu.Menu) float32 {
	var total float32
	for _, item := range o.IDs {
		meal := m.Find(item)
		total += meal.Price
	}
	return total
}
