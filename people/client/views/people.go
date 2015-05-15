package views

import (
	"github.com/go-humble/examples/people/shared/models"
	"github.com/go-humble/examples/people/shared/templates"
	"github.com/go-humble/temple/temple"
	"github.com/go-humble/view"
	"honnef.co/go/js/dom"
)

var (
	showPersonTmpl  temple.Partial
	indexPeopleTmpl temple.Partial
	mainEl          = dom.GetWindow().Document().QuerySelector("#main")
)

func init() {
	var found bool
	showPersonTmpl, found = templates.Partials["person"]
	if !found {
		panic("Could not find tepmlate called person")
	}
	indexPeopleTmpl, found = templates.Partials["people"]
	if !found {
		panic("Could not find tepmlate called people")
	}
}

type ShowPerson struct {
	Person *models.Person
	view.DefaultView
}

func NewShowPerson(person *models.Person) *ShowPerson {
	v := &ShowPerson{
		Person: person,
	}
	v.SetElement(mainEl)
	return v
}

func (v ShowPerson) Render() error {
	return showPersonTmpl.ExecuteEl(v.Element(), v.Person)
}

type IndexPeople struct {
	People []*models.Person
	view.DefaultView
}

func NewIndexPeople(people []*models.Person) *IndexPeople {
	v := &IndexPeople{
		People: people,
	}
	v.SetElement(mainEl)
	return v
}

func (v IndexPeople) Render() error {
	return indexPeopleTmpl.ExecuteEl(v.Element(), v.People)
}
