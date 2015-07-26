package main

import (
	"log"

	"github.com/codegangsta/negroni"
	"github.com/go-humble/examples/people/server/controllers"
	"github.com/go-humble/examples/people/shared/models"
	"github.com/gorilla/mux"
)

//go:generate temple build ../shared/templates/templates ../shared/templates/templates.go --partials=../shared/templates/partials --layouts=../shared/templates/layouts
//go:generate gopherjs build ../client/main.go -o public/js/app.js

func main() {
	defer func() {
		if err := models.ClosePool(); err != nil {
			log.Fatal(err)
		}
	}()
	people := controllers.People{}
	router := mux.NewRouter()
	router.HandleFunc("/people/new", people.New).Methods("GET")
	router.HandleFunc("/people", people.Create).Methods("POST")
	router.HandleFunc("/people/{id}", people.Show).Methods("GET")
	router.HandleFunc("/people", people.Index).Methods("GET")
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}
