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
	Todos     *models.TodoList
	tmpl      *temple.Template
	predicate models.Predicate
	view.DefaultView
}

func (v *App) UseFilter(predicate models.Predicate) {
	v.predicate = predicate
}

func NewApp(todos *models.TodoList) *App {
	v := &App{
		Todos: todos,
		tmpl:  appTmpl,
	}
	v.SetElement(document.QuerySelector(".todoapp"))
	return v
}

func (v *App) tmplData() map[string]interface{} {
	return map[string]interface{}{
		"Todos": v.Todos,
		"Path":  dom.GetWindow().Location().Hash,
	}
}

func (v *App) Render() error {
	if err := v.tmpl.ExecuteEl(v.Element(), v.tmplData()); err != nil {
		return err
	}
	listEl := v.Element().QuerySelector(".todo-list")
	for _, todo := range v.Todos.Filter(v.predicate) {
		todoView := NewTodo(todo)
		view.AppendToEl(listEl, todoView)
		if err := todoView.Render(); err != nil {
			return err
		}
	}
	v.DelegateEvents()
	return nil
}

func (v *App) DelegateEvents() {
	view.AddEventListener(v, "keypress", ".new-todo",
		triggerOnKeyCode(enterKey, v.CreateTodo))
	view.AddEventListener(v, "click", ".clear-completed", v.ClearCompleted)
	view.AddEventListener(v, "click", ".toggle-all", v.ToggleAll)
}

func (v *App) CreateTodo(ev dom.Event) {
	input, ok := ev.Target().(*dom.HTMLInputElement)
	if !ok {
		panic("Could not convert event target to dom.HTMLInputElement")
	}
	v.Todos.AddTodo(input.Value)
	document.QuerySelector(".new-todo").(dom.HTMLElement).Focus()
}

func (v *App) ClearCompleted(ev dom.Event) {
	v.Todos.ClearCompleted()
}

func (v *App) ToggleAll(ev dom.Event) {
	input, ok := ev.Target().(*dom.HTMLInputElement)
	if !ok {
		panic("Could not convert event target to dom.HTMLInputElement")
	}
	for _, todo := range v.Todos.All() {
		todo.SetCompleted(input.Checked)
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
