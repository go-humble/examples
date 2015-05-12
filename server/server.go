package main

import (
	"github.com/albrow/temple-example/server/controllers"
	"github.com/albrow/temple-example/shared/models"
	_ "github.com/albrow/temple-example/shared/templates"
	"github.com/albrow/zoom"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
)

func init() {
	if err := models.RegisterAll(); err != nil {
		log.Fatal(err)
	}
	if err := zoom.Init(nil); err != nil {
		log.Fatal(err)
	}
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
