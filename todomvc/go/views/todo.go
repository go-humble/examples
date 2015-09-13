package views

import (
	"github.com/go-humble/examples/todomvc/go/models"
	"github.com/go-humble/examples/todomvc/go/templates"
	"github.com/go-humble/temple/temple"
	"github.com/go-humble/view"
	"honnef.co/go/js/dom"
)

var (
	todoTmpl = templates.MustGetPartial("todo")
)

type Todo struct {
	Model *models.Todo
	tmpl  *temple.Partial
	view.DefaultView
}

func NewTodo(todo *models.Todo) *Todo {
	return &Todo{
		Model: todo,
		tmpl:  todoTmpl,
	}
}

func (v *Todo) Render() error {
	if err := v.tmpl.ExecuteEl(v.Element(), v.Model); err != nil {
		return err
	}
	v.DelegateEvents()
	return nil
}

func (v *Todo) DelegateEvents() {
	view.AddEventListener(v, "click", ".toggle", v.Toggle)
	view.AddEventListener(v, "click", ".destroy", v.Remove)
	view.AddEventListener(v, "dblclick", "label", v.Edit)
	view.AddEventListener(v, "blur", ".edit", v.CommitEdit)
	view.AddEventListener(v, "keypress", ".edit",
		triggerOnKeyCode(enterKey, v.CommitEdit))
	view.AddEventListener(v, "keydown", ".edit",
		triggerOnKeyCode(escapeKey, v.CancelEdit))

}

func (v *Todo) Toggle(ev dom.Event) {
	v.Model.Toggle()
}

func (v *Todo) Remove(ev dom.Event) {
	v.Model.Remove()
}

func (v *Todo) Edit(ev dom.Event) {
	li := v.Element().QuerySelector("li")
	addClass(li, "editing")
	input, ok := li.QuerySelector(".edit").(*dom.HTMLInputElement)
	if !ok {
		panic("Could not convert to dom.HTMLInputElement")
	}
	input.Focus()
	// Move the cursor to the end of the input.
	input.SelectionStart = input.SelectionEnd + len(input.Value)
}

func (v *Todo) CommitEdit(ev dom.Event) {
	li := v.Element().QuerySelector("li")
	input, ok := li.QuerySelector(".edit").(*dom.HTMLInputElement)
	if !ok {
		panic("Could not convert to dom.HTMLInputElement")
	}
	v.Model.SetTitle(input.Value)
}

func (v *Todo) CancelEdit(ev dom.Event) {
	li := v.Element().QuerySelector("li")
	removeClass(li, "editing")
	input, ok := li.QuerySelector(".edit").(*dom.HTMLInputElement)
	if !ok {
		panic("Could not convert to dom.HTMLInputElement")
	}
	input.Value = v.Model.Title()
	input.Blur()
}
