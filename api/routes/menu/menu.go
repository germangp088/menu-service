package menu

import (
	"encoding/json"
	meal "menu-service/api/types/meal"
	menu "menu-service/api/types/menu"
	"net/http"
)

var _menu menu.Menu

func Init() {
	Sections()
	Meals()
}

func Sections() {
	_menu.Sections = append(_menu.Sections, "SALADS")
	_menu.Sections = append(_menu.Sections, "PORK")
	_menu.Sections = append(_menu.Sections, "DRINK")
}

func Meals() {
	_menu.Meals = append(_menu.Meals, meal.Meal{Name: "Coleslaw", Section: _menu.Sections[0], Price: 9.05})
	_menu.Meals = append(_menu.Meals, meal.Meal{Name: "Caesar Saldad", Section: _menu.Sections[0], Price: 11.25})
	_menu.Meals = append(_menu.Meals, meal.Meal{Name: "Crunchy Pork", Section: _menu.Sections[1], Price: 15.10})
	_menu.Meals = append(_menu.Meals, meal.Meal{Name: "Pork BBQ", Section: _menu.Sections[1], Price: 20.50})
	_menu.Meals = append(_menu.Meals, meal.Meal{Name: "Lechon Kawali", Section: _menu.Sections[1], Price: 55.60})
	_menu.Meals = append(_menu.Meals, meal.Meal{Name: "Iced Tea", Section: _menu.Sections[2], Price: 2.99})
	_menu.Meals = append(_menu.Meals, meal.Meal{Name: "Coke", Section: _menu.Sections[2], Price: 6.00})
	_menu.Meals = append(_menu.Meals, meal.Meal{Name: "Beer", Section: _menu.Sections[2], Price: 9.99})
}

func GetMenu(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(_menu)
}
