package views

import (
	"fmt"

	"github.com/go-humble/form"

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
	errTmpl         = templates.MustGetPartial("errors")
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

type Errors struct {
	Errors []error
	view.DefaultView
}

func (v *Errors) Render() error {
	if err := errTmpl.ExecuteEl(v.Element(), v.Errors); err != nil {
		return err
	}
	return nil
}

func NewErrors(errors []error) *Errors {
	return &Errors{
		Errors: errors,
	}
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
	view.AddEventListener(v, "submit", "#person-form", v.CreatePerson)
}

var errorsView *Errors

func (v *NewPerson) CreatePerson(ev dom.Event) {
	ev.PreventDefault()
	formEl, ok := ev.CurrentTarget().(*dom.HTMLFormElement)
	if !ok {
		panic("Could not cast target to dom.HTMLFormElement: " + fmt.Sprintf("%T", ev.CurrentTarget()))
	}
	f, err := form.Parse(formEl)
	if err != nil {
		panic(err)
	}
	f.Validate("name").Required()
	f.Validate("age").Required().IsInt().Greater(0)
	if f.HasErrors() {
		if errorsView == nil {
			errorsView = NewErrors(f.Errors)
			view.InsertBefore(errorsView, v)
		}
		errorsView.Errors = f.Errors
		if err := errorsView.Render(); err != nil {
			panic(err)
		}
		return
	}
	person := &models.Person{}
	if err := f.Bind(person); err != nil {
		panic(err)
	}
	restClient := rest.NewClient()
	restClient.ContentType = rest.ContentJSON
	go func() {
		if err := restClient.Create(person); err != nil {
			if httpErr, ok := err.(rest.HTTPError); ok {
				fmt.Println(string(httpErr.Body))
			}
			panic(err)
		}
		v.Router.Navigate("/people")
	}()
}
