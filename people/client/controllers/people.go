package controllers

import (
	"github.com/go-humble/examples/people/client/views"
	"github.com/go-humble/examples/people/shared/models"
	"github.com/go-humble/rest"
	"log"
)

type People struct{}

func (p People) Index(params map[string]string) {
	people := []*models.Person{}
	if err := rest.ReadAll(&people); err != nil {
		log.Fatal(err)
	}
	if err := views.NewIndexPeople(people).Render(); err != nil {
		log.Fatal(err)
	}
}

func (p People) Show(params map[string]string) {
	person := &models.Person{}
	if err := rest.Read(params["id"], person); err != nil {
		log.Fatal(err)
	}
	if err := views.NewShowPerson(person).Render(); err != nil {
		log.Fatal(err)
	}
}
