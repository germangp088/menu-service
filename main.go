package main

import (
	"log"
	"net/http"

	menu "menu-service/api/routes/menu"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	Initialize()
}

func Initialize() {
	menu.Init()
	router := mux.NewRouter()

	router.HandleFunc("/menu", menu.GetMenu).Methods("GET")
	router.HandleFunc("/menu", menu.AddMeal).Methods("POST")
	router.HandleFunc("/meal/{id}", menu.GetMeal).Methods("GET")
	router.HandleFunc("/cash", menu.GetCash).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
