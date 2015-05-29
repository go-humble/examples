package views

import (
	"github.com/go-humble/examples/people/shared/models"
	"github.com/go-humble/examples/people/shared/templates"
	"github.com/go-humble/view"
	"honnef.co/go/js/dom"
)

var (
	showPersonTmpl  = templates.MustGetPartial("person")
	indexPeopleTmpl = templates.MustGetPartial("people")
	mainEl          = dom.GetWindow().Document().QuerySelector("#main")
)

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
