package main

import (
	"log"
	"net/http"

	"github.com/albrow/zoom"
	"github.com/codegangsta/negroni"
	"github.com/go-humble/examples/people/server/controllers"
	"github.com/go-humble/examples/people/shared/models"
	_ "github.com/go-humble/examples/people/shared/templates"
	"github.com/gorilla/mux"
)

//go:generate temple build ../shared/templates/templates ../shared/templates/templates.go --partials=../shared/templates/partials --layouts=../shared/templates/layouts
//go:generate gopherjs build ../client/main.go -o public/js/app.js

func init() {
	if err := CreateInitialPeople(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer func() {
		if err := zoom.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	people := controllers.People{}
	router := mux.NewRouter()
	router.HandleFunc("/people/{id}", people.Show).Methods("GET")
	router.HandleFunc("/people", people.Index).Methods("GET")
	n := negroni.Classic()
	n.UseHandler(router)
	s := negroni.NewStatic(http.Dir("."))
	s.Prefix = "/public"
	n.Run(":3000")
}

func CreateInitialPeople() error {
	if count, err := models.People.Count(); err != nil {
		return err
	} else if count == 0 {
		log.Println("Creating people...")
		t := zoom.NewTransaction()
		for i, name := range []string{"Foo", "Bar", "Baz"} {
			person := &models.Person{
				Age:  i + 20,
				Name: name,
			}
			t.Save(models.People, person)
		}
		if err := t.Exec(); err != nil {
			return err
		}
	}
	return nil
}
