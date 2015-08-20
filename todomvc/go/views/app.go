package views

import (
	"fmt"
	"strings"

	"github.com/go-humble/examples/todomvc/go/models"
	"github.com/go-humble/examples/todomvc/go/templates"
	"github.com/go-humble/temple/temple"
	"github.com/go-humble/view"
	"honnef.co/go/js/dom"
)

const (
	enterKey  = 13
	escapeKey = 27
)

var (
	appTmpl  = templates.MustGetTemplate("app")
	document = dom.GetWindow().Document()
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
	view.AddEventListener(v, "keypress", ".new-todo",
		triggerOnKeyCode(enterKey, v.CreateTodo))
	view.AddEventListener(v, "click", ".clear-completed", v.ClearCompleted)
	view.AddEventListener(v, "click", ".toggle-all", v.ToggleAll)
	view.AddEventListener(v, "click", "li .toggle", v.ToggleTodo)
	view.AddEventListener(v, "click", "li .destroy", v.RemoveTodo)
	view.AddEventListener(v, "dblclick", "li label", v.EditTodo)
	view.AddEventListener(v, "blur", "li .edit", v.CommitEditTodo)
	view.AddEventListener(v, "keydown", "li .edit",
		triggerOnKeyCode(escapeKey, v.CancelEditTodo))
	view.AddEventListener(v, "keypress", "li .edit",
		triggerOnKeyCode(enterKey, v.CommitEditTodo))
}

func (v *App) CreateTodo(ev dom.Event) {
	input, ok := ev.Target().(*dom.HTMLInputElement)
	if !ok {
		panic("Could not convert event target to dom.HTMLInputElement")
	}
	v.Todos.AddTodo(input.Value)
	if err := v.Render(); err != nil {
		panic(err)
	}
	// When we call Render, a large portion of teh DOM is replaced, so we need
	// to select the new input element and call focus on it.
	document.QuerySelector(".new-todo").(dom.HTMLElement).Focus()
}

func (v *App) ClearCompleted(ev dom.Event) {
	v.Todos.ClearCompleted()
	if err := v.Render(); err != nil {
		panic(err)
	}
}

func (v *App) ToggleAll(ev dom.Event) {
	input, ok := ev.Target().(*dom.HTMLInputElement)
	if !ok {
		panic("Could not convert event target to dom.HTMLInputElement")
	}
	for _, todo := range v.Todos.All() {
		todo.Completed = input.Checked
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

func (v *App) EditTodo(ev dom.Event) {
	li := ev.Target().ParentElement().ParentElement()
	addClass(li, "editing")
	input, ok := li.QuerySelector(".edit").(*dom.HTMLInputElement)
	if !ok {
		panic("Could not convert to dom.HTMLInputElement")
	}
	input.Focus()
	// Move the cursor to the end of the input.
	input.SelectionStart = input.SelectionEnd + len(input.Value)
}

func (v *App) CommitEditTodo(ev dom.Event) {
	li := ev.Target().ParentElement()
	input, ok := li.QuerySelector(".edit").(*dom.HTMLInputElement)
	if !ok {
		panic("Could not convert to dom.HTMLInputElement")
	}
	todo, found := v.Todos.FindById(li.QuerySelector(".view").GetAttribute("data-id"))
	if !found {
		panic("Could not find todo")
	}
	todo.Title = input.Value
	if err := v.Render(); err != nil {
		panic(err)
	}
}

func (v *App) CancelEditTodo(ev dom.Event) {
	li := ev.Target().ParentElement()
	removeClass(li, "editing")
	input, ok := li.QuerySelector(".edit").(*dom.HTMLInputElement)
	if !ok {
		panic("Could not convert to dom.HTMLInputElement")
	}
	todo, found := v.Todos.FindById(li.QuerySelector(".view").GetAttribute("data-id"))
	if !found {
		panic("Could not find todo")
	}
	input.Value = todo.Title
	input.Blur()
}

func (v *App) RemoveTodo(ev dom.Event) {
	id := ev.CurrentTarget().ParentElement().GetAttribute("data-id")
	v.Todos.DeleteById(id)
	if err := v.Render(); err != nil {
		panic(err)
	}
}

func triggerOnKeyCode(keyCode int, listener func(dom.Event)) func(dom.Event) {
	return func(ev dom.Event) {
		keyEvent, ok := ev.(*dom.KeyboardEvent)
		if ok && keyEvent.KeyCode == keyCode {
			listener(ev)
		}
	}
}

func addClass(el dom.Element, value string) {
	newClasses := value
	if oldClasses := el.GetAttribute("class"); oldClasses != "" {
		newClasses = oldClasses + " " + value
	}
	el.SetAttribute("class", newClasses)
}

func removeClass(el dom.Element, value string) {
	oldClasses := el.GetAttribute("class")
	if oldClasses == value {
		fmt.Println("Only classes were the value itself")
		// The only class present was the one we want to remove. Remove the class
		// attribute entirely.
		el.RemoveAttribute("class")
	}
	classList := strings.Split(oldClasses, " ")
	for i, class := range classList {
		if class == value {
			newClassList := append(classList[:i], classList[i+1:]...)
			el.SetAttribute("class", strings.Join(newClassList, " "))
		}
	}
}
