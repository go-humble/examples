package controllers

import (
	"github.com/go-humble/examples/people/client/views"
	"github.com/go-humble/examples/people/shared/models"
	"github.com/go-humble/rest"
	"github.com/go-humble/router"
	"log"
)

type People struct{}

func (p People) Index(context *router.Context) {
	people := []*models.Person{}
	if err := rest.ReadAll(&people); err != nil {
		log.Fatal(err)
	}
	if err := views.NewIndexPeople(people).Render(); err != nil {
		log.Fatal(err)
	}
}

func (p People) Show(context *router.Context) {
	person := &models.Person{}
	if err := rest.Read(context.Params["id"], person); err != nil {
		log.Fatal(err)
	}
	if err := views.NewShowPerson(person).Render(); err != nil {
		log.Fatal(err)
	}
}
