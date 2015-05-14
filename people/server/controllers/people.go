package controllers

import (
	"github.com/go-humble/examples/people/shared/models"
	"github.com/go-humble/temple"
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
	indexView, found := temple.Templates["people/index"]
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
		if err := indexView.Execute(res, people); err != nil {
			log.Fatal(err)
		}
	default:
		if err := indexView.Execute(res, people); err != nil {
			log.Fatal(err)
		}
	}
}

func (p People) Show(res http.ResponseWriter, req *http.Request) {
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
	switch req.Header.Get("Accept") {
	case contentTypeJSON:
		r := render.New()
		r.JSON(res, http.StatusOK, person)
	case contentTypeHTML:
		if err := showView.Execute(res, person); err != nil {
			log.Fatal(err)
		}
	default:
		if err := showView.Execute(res, person); err != nil {
			log.Fatal(err)
		}
	}
}
