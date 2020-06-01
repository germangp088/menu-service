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

func (m *Menu) Update(meal meal.Meal) {
	for i := 0; i < len(m.Meals); i++ {
		if m.Meals[i].ID == meal.ID {
			m.Meals[i] = meal
		}
	}
}
