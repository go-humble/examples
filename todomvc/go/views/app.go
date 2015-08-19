package views

import (
	"github.com/go-humble/examples/todomvc/go/models"
	"github.com/go-humble/examples/todomvc/go/templates"
	"github.com/go-humble/temple/temple"
	"github.com/go-humble/view"
	"honnef.co/go/js/dom"
)

var (
	appTmpl = templates.MustGetTemplate("app")
)

type App struct {
	Todos models.TodoList
	tmpl  *temple.Template
	view.DefaultView
}

func NewApp(todos models.TodoList) *App {
	v := &App{
		Todos: todos,
		tmpl:  appTmpl,
	}
	v.SetElement(document.QuerySelector(".todoapp"))
	return v
}

func (v *App) Render() error {
	if err := v.tmpl.ExecuteEl(v.Element(), v.Todos); err != nil {
		return err
	}
	v.DelegateEvents()
	return nil
}

func (v *App) DelegateEvents() {
	view.AddEventListener(v, "keypress", ".new-todo", func(ev dom.Event) {
		if isEnterPress(ev) {
			v.CreateTodo(ev)
		}
	})
	view.AddEventListener(v, "click", ".clear-completed", v.ClearCompleted)
	view.AddEventListener(v, "click", ".toggle-all", v.ToggleAll)
	view.AddEventListener(v, "click", "li .toggle", v.ToggleTodo)
	view.AddEventListener(v, "click", "li .destroy", v.RemoveTodo)
}

func (v *App) CreateTodo(ev dom.Event) {
	inputEl, ok := ev.Target().(*dom.HTMLInputElement)
	if !ok {
		panic("Could not convert event target to inputEl")
	}
	v.Todos.AddTodo(inputEl.Value)
	if err := v.Render(); err != nil {
		panic(err)
	}
}

func (v *App) ClearCompleted(ev dom.Event) {
	v.Todos.ClearCompleted()
	if err := v.Render(); err != nil {
		panic(err)
	}
}

func (v *App) ToggleAll(ev dom.Event) {
	inputEl, ok := ev.Target().(*dom.HTMLInputElement)
	if !ok {
		panic("Could not convert event target to inputEl")
	}
	for _, todo := range v.Todos.All() {
		todo.Completed = inputEl.Checked
	}
	if err := v.Render(); err != nil {
		panic(err)
	}
}

func (v *App) ToggleTodo(ev dom.Event) {
	id := ev.CurrentTarget().ParentElement().GetAttribute("data-id")
	todo, found := v.Todos.FindById(id)
	if !found {
		panic("Could not find todo with id: " + id)
	}
	todo.Toggle()
	if err := v.Render(); err != nil {
		panic(err)
	}
}

func (v *App) RemoveTodo(ev dom.Event) {
	id := ev.CurrentTarget().ParentElement().GetAttribute("data-id")
	v.Todos.DeleteById(id)
	if err := v.Render(); err != nil {
		panic(err)
	}
}

const enterKey = 13

func isEnterPress(ev dom.Event) bool {
	keyEvent, ok := ev.(*dom.KeyboardEvent)
	return ok && keyEvent.KeyCode == enterKey
}
