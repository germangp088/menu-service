package menu

import (
	meal "menu-service/api/types/meal"
)

type Menu struct {
	Sections []string    `json:"sections"`
	Meals    []meal.Meal `json:"meals"`
}

func (m Menu) Find(id string) *meal.Meal {
	for _, meal := range m.Meals {
		if meal.ID == id {
			return &meal
		}
	}
	return nil
}
