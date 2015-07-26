package controllers

import (
	"log"
	"net/http"

	"github.com/albrow/forms"
	"github.com/go-humble/examples/people/shared/models"
	"github.com/go-humble/examples/people/shared/templates"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/unrolled/render"
)

const (
	contentTypeJSON = "application/json"
	contentTypeHTML = "text/html"
)

var (
	peopleIndexTmpl = templates.MustGetTemplate("people/index")
	peopleShowTmpl  = templates.MustGetTemplate("people/show")
	peopleNewTmpl   = templates.MustGetTemplate("people/new")
)

var store = sessions.NewCookieStore([]byte("a-secret-string"))

type People struct{}

func (p People) New(res http.ResponseWriter, req *http.Request) {
	personData := map[string]interface{}{}
	tmplData := map[string]interface{}{}
	session, err := store.Get(req, "flash-session")
	if err != nil {
		panic(err)
	}
	if errors := session.Flashes("errors"); len(errors) > 0 {
		tmplData["Errors"] = errors
	}
	if names := session.Flashes("person.name"); len(names) > 0 {
		personData["Name"] = names[0]
	}
	if ages := session.Flashes("person.age"); len(ages) > 0 {
		personData["Age"] = ages[0]
	}
	session.Save(req, res)
	tmplData["Person"] = personData
	log.Println(tmplData)
	if err := peopleNewTmpl.Execute(res, tmplData); err != nil {
		log.Fatal(err)
	}
}

func (p People) Create(res http.ResponseWriter, req *http.Request) {
	personData, err := forms.Parse(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(personData)
	val := personData.Validator()
	val.Require("name")
	val.Require("age")
	val.TypeInt("age")
	val.Greater("age", 0)
	if val.HasErrors() {
		switch req.Header.Get("Accept") {
		case contentTypeJSON:
			r := render.New()
			r.JSON(res, 422, val.ErrorMap())
		default:
			session, err := store.Get(req, "flash-session")
			if err != nil {
				panic(err)
			}
			for _, errMsg := range val.Messages() {
				session.AddFlash(errMsg, "errors")
			}
			session.AddFlash(personData.Get("name"), "person.name")
			session.AddFlash(personData.Get("age"), "person.age")
			session.Save(req, res)
			http.Redirect(res, req, "/people/new", http.StatusFound)
			return
		}
	}
	person := &models.Person{
		Name: personData.Get("name"),
		Age:  personData.GetInt("age"),
	}
	if err := models.People.Save(person); err != nil {
		panic(err)
	}
	http.Redirect(res, req, "/people", http.StatusSeeOther)
}

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
		if err := peopleIndexTmpl.Execute(res, map[string]interface{}{
			"People": people,
		}); err != nil {
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
		if err := peopleShowTmpl.Execute(res, map[string]interface{}{
			"person": person,
		}); err != nil {
			log.Fatal(err)
		}
	}
}
