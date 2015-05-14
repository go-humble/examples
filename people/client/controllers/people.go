package controllers

import (
	"github.com/go-humble/examples/people/shared/models"
	_ "github.com/go-humble/examples/people/shared/templates"
	"github.com/go-humble/humble/rest"
	"github.com/go-humble/temple"
	"honnef.co/go/js/dom"
	"log"
)

var (
	mainEl     = dom.GetWindow().Document().QuerySelector("#main")
	peopleTmpl temple.Partial
	personTmpl temple.Partial
)

func init() {
	var found bool
	peopleTmpl, found = temple.Partials["people"]
	if !found {
		log.Fatal("Could not find people partial")
	}
	personTmpl, found = temple.Partials["person"]
	if !found {
		log.Fatal("Could not find person partial")
	}
}

type People struct{}

func (p People) Index(params map[string]string) {
	people := []models.Person{}
	if err := rest.ReadAll(&people); err != nil {
		panic(err)
	}
	if err := temple.ExecuteToEl(peopleTmpl, mainEl, people); err != nil {
		panic(err)
	}
}

func (p People) Show(params map[string]string) {
	person := &models.Person{}
	if err := rest.Read(params["id"], person); err != nil {
		panic(err)
	}
	if err := temple.ExecuteToEl(personTmpl, mainEl, person); err != nil {
		panic(err)
	}
}
