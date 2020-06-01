package menu

import meal "menu-service/api/types/meal"

type Menu struct {
	Sections []string    `json:"sections"`
	Meals    []meal.Meal `json:"meals"`
}
