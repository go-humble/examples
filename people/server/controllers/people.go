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

var (
	peopleIndexTmpl = templates.MustGetTemplate("people/index")
	peopleShowTmpl  = templates.MustGetTemplate("people/show")
)

type People struct{}

func (p People) Index(res http.ResponseWriter, req *http.Request) {
	people := []*models.Person{}
	if err := models.People.FindAll(&people); err != nil {
		log.Fatal(err)
	}
	switch req.Header.Get("Accept") {
	case contentTypeJSON:
		r := render.New()
		r.JSON(res, http.StatusOK, people)
	default:
		if err := peopleIndexTmpl.Execute(res, people); err != nil {
			log.Fatal(err)
		}
	}
}

func (p People) Show(res http.ResponseWriter, req *http.Request) {
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
	default:
		if err := peopleShowTmpl.Execute(res, person); err != nil {
			log.Fatal(err)
		}
	}
}
