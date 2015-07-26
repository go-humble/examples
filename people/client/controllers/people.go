package controllers

import (
	"log"

	"github.com/go-humble/examples/people/client/views"
	"github.com/go-humble/examples/people/shared/models"
	"github.com/go-humble/rest"
	"github.com/go-humble/router"
)

type People struct{}

var client = rest.NewClient()

func (p People) New(context *router.Context) {
	if context.InitialLoad {
		return
	}
	if err := views.NewNewPerson(nil).Render(); err != nil {
		log.Fatal(err)
	}
}

func (p People) Index(context *router.Context) {
	if context.InitialLoad {
		return
	}
	people := []*models.Person{}
	if err := client.ReadAll(&people); err != nil {
		log.Fatal(err)
	}
	if err := views.NewIndexPeople(people).Render(); err != nil {
		log.Fatal(err)
	}
}

func (p People) Show(context *router.Context) {
	if context.InitialLoad {
		return
	}
	person := &models.Person{}
	if err := client.Read(context.Params["id"], person); err != nil {
		log.Fatal(err)
	}
	if err := views.NewShowPerson(person).Render(); err != nil {
		log.Fatal(err)
	}
}
