package controllers

import (
	"github.com/go-humble/examples/people/shared/models"
	"github.com/go-humble/examples/people/shared/templates"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"log"
	"net/http"
)

const (
	contentTypeJSON = "application/json"
	contentTypeHTML = "text/html"
)

type People struct{}

func (p People) Index(res http.ResponseWriter, req *http.Request) {
	indexTmpl, found := templates.Templates["people/index"]
	if !found {
		log.Fatalf("Could not find template named %s", "people/index")
	}
	people := []*models.Person{}
	if err := models.People.FindAll(&people); err != nil {
		log.Fatal(err)
	}
	switch req.Header.Get("Accept") {
	case contentTypeJSON:
		r := render.New()
		r.JSON(res, http.StatusOK, people)
	case contentTypeHTML:
		if err := indexTmpl.Execute(res, people); err != nil {
			log.Fatal(err)
		}
	default:
		if err := indexTmpl.Execute(res, people); err != nil {
			log.Fatal(err)
		}
	}
}

func (p People) Show(res http.ResponseWriter, req *http.Request) {
	showTmpl, found := templates.Templates["people/show"]
	if !found {
		log.Fatalf("Could not find template named %s", "people/show")
	}
	vars := mux.Vars(req)
	id := vars["id"]
	person := &models.Person{}
	if err := models.People.Find(id, person); err != nil {
		log.Fatal(err)
	}
	switch req.Header.Get("Accept") {
	case contentTypeJSON:
		r := render.New()
		r.JSON(res, http.StatusOK, person)
	case contentTypeHTML:
		if err := showTmpl.Execute(res, person); err != nil {
			log.Fatal(err)
		}
	default:
		if err := showTmpl.Execute(res, person); err != nil {
			log.Fatal(err)
		}
	}
}
