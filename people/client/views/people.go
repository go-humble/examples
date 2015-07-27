package views

import (
	"fmt"
	"strconv"

	"github.com/go-humble/examples/people/shared/models"
	"github.com/go-humble/examples/people/shared/templates"
	"github.com/go-humble/rest"
	"github.com/go-humble/router"
	"github.com/go-humble/view"
	"honnef.co/go/js/dom"
)

var (
	showPersonTmpl  = templates.MustGetPartial("people/show")
	indexPeopleTmpl = templates.MustGetPartial("people/index")
	newPersonTmpl   = templates.MustGetPartial("people/new")
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

func (v *ShowPerson) Render() error {
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

func (v *IndexPeople) Render() error {
	return indexPeopleTmpl.ExecuteEl(v.Element(), v.People)
}

type NewPerson struct {
	Person *models.Person
	Router *router.Router
	view.DefaultView
}

func NewNewPerson(person *models.Person, router *router.Router) *NewPerson {
	v := &NewPerson{
		Person: person,
		Router: router,
	}
	v.SetElement(mainEl)
	return v
}

func (v *NewPerson) Render() error {
	if err := newPersonTmpl.ExecuteEl(v.Element(), v.Person); err != nil {
		return err
	}
	v.DelegateEvents()
	return nil
}

func (v *NewPerson) DelegateEvents() {
	view.AddEventListener(v, "submit", "#person-form", NewCreatePersonListener(v.Router))
}

func NewCreatePersonListener(router *router.Router) func(dom.Event) {
	return func(ev dom.Event) {
		ev.PreventDefault()
		form, ok := ev.CurrentTarget().(*dom.HTMLFormElement)
		if !ok {
			panic("Could not cast target to dom.HTMLFormElement: " + fmt.Sprintf("%T", ev.CurrentTarget()))
		}
		person := &models.Person{}
		for _, el := range form.Elements() {
			input, ok := el.(*dom.HTMLInputElement)
			if !ok {
				continue
			}
			if input.Type == "submit" {
				continue
			}
			switch input.Name {
			case "age":
				ageInt, err := strconv.Atoi(input.Value)
				if err != nil {
					panic(err)
				}
				person.Age = ageInt
			case "name":
				person.Name = input.Value
			}
		}
		restClient := rest.NewClient()
		restClient.ContentType = rest.ContentJSON
		go func() {
			if err := restClient.Create(person); err != nil {
				panic(err)
			}
			router.Navigate("/people")
		}()
	}
}
