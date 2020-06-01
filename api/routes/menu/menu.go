package menu

import (
	"encoding/json"
	meal "menu-service/api/types/meal"
	menu "menu-service/api/types/menu"
	"net/http"

	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
)

var _menu menu.Menu
var _cash int

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
	_menu.Meals = append(_menu.Meals, meal.Meal{ID: guuid.New().String(), Name: "Coleslaw", Section: _menu.Sections[0], Price: 9.05})
	_menu.Meals = append(_menu.Meals, meal.Meal{ID: guuid.New().String(), Name: "Caesar Saldad", Section: _menu.Sections[0], Price: 11.25})
	_menu.Meals = append(_menu.Meals, meal.Meal{ID: guuid.New().String(), Name: "Crunchy Pork", Section: _menu.Sections[1], Price: 15.10})
	_menu.Meals = append(_menu.Meals, meal.Meal{ID: guuid.New().String(), Name: "Pork BBQ", Section: _menu.Sections[1], Price: 20.50})
	_menu.Meals = append(_menu.Meals, meal.Meal{ID: guuid.New().String(), Name: "Lechon Kawali", Section: _menu.Sections[1], Price: 55.60})
	_menu.Meals = append(_menu.Meals, meal.Meal{ID: guuid.New().String(), Name: "Iced Tea", Section: _menu.Sections[2], Price: 2.99})
	_menu.Meals = append(_menu.Meals, meal.Meal{ID: guuid.New().String(), Name: "Coke", Section: _menu.Sections[2], Price: 6.00})
	_menu.Meals = append(_menu.Meals, meal.Meal{ID: guuid.New().String(), Name: "Beer", Section: _menu.Sections[2], Price: 9.99})
}

func GetMenu(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(_menu)
}

func GetMeal(w http.ResponseWriter, req *http.Request) {
	query := mux.Vars(req)
	id := query["id"]
	_menu.Find(id)
	meal := _menu.Find(id)
	if meal == nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("Not found")
	} else {
		json.NewEncoder(w).Encode(meal)
	}
}

func GetCash(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(_cash)
}

func AddMeal(w http.ResponseWriter, req *http.Request) {
	var meal meal.Meal
	_ = json.NewDecoder(req.Body).Decode(&meal)

	meal.ID = guuid.New().String()

	_menu.Meals = append(_menu.Meals, meal)
	json.NewEncoder(w).Encode(_menu)
}
