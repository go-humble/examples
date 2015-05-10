package main

import (
	"fmt"
	"github.com/albrow/temple"
	"github.com/albrow/temple-example/models"
	_ "github.com/albrow/temple-example/templates"
	"github.com/albrow/zoom"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
	router := mux.NewRouter()
	router.HandleFunc("/people/{id}", ShowPerson).Methods("GET")
	router.HandleFunc("/people", IndexPeople).Methods("GET")
	n := negroni.Classic()
	n.Use(negroni.NewStatic(http.Dir("../client")))
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

func IndexPeople(res http.ResponseWriter, req *http.Request) {
	indexView, found := temple.Templates["people/index"]
	if !found {
		log.Fatalf("Could not find template named %s", "people/index")
	}
	fmt.Println(indexView.Template.Tree)
	people := []*models.Person{}
	if err := models.People.FindAll(&people); err != nil {
		log.Fatal(err)
	}
	if err := indexView.Execute(res, people); err != nil {
		log.Fatal(err)
	}
}

func ShowPerson(res http.ResponseWriter, req *http.Request) {
	showView, found := temple.Templates["people/show"]
	if !found {
		log.Fatalf("Could not find template named %s", "people/show")
	}
	vars := mux.Vars(req)
	id := vars["id"]
	person := &models.Person{}
	if err := models.People.Find(id, person); err != nil {
		log.Fatal(err)
	}
	if err := showView.Execute(res, person); err != nil {
		log.Fatal(err)
	}
}
