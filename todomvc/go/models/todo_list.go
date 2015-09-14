package models

import (
	"github.com/dchest/uniuri"
	"github.com/go-humble/locstor"
)

var store = locstor.NewDataStore(locstor.JSONEncoding)

type TodoList struct {
	todos           []*Todo
	changeListeners []func(*TodoList)
}

func (list *TodoList) OnChange(f func(*TodoList)) {
	list.changeListeners = append(list.changeListeners, f)
}

func (list *TodoList) changed() {
	for _, f := range list.changeListeners {
		f(list)
	}
}

func (list *TodoList) Load() error {
	if err := store.Find("todos", &list.todos); err != nil {
		if _, ok := err.(locstor.ItemNotFoundError); ok {
			return list.Save()
		}
		return err
	}
	for i := range list.todos {
		list.todos[i].list = list
	}
	return nil
}

func (list TodoList) Save() error {
	if err := store.Save("todos", list.todos); err != nil {
		return err
	}
	return nil
}

type Filter func(*TodoList) []*Todo

var Filters = struct {
	All       Filter
	Completed Filter
	Remaining Filter
}{
	All:       (*TodoList).All,
	Completed: (*TodoList).Completed,
	Remaining: (*TodoList).Remaining,
}

func (list TodoList) All() []*Todo {
	return list.todos
}

func (list TodoList) Completed() []*Todo {
	return list.filter((*Todo).Completed)
}

func (list TodoList) Remaining() []*Todo {
	return list.filter((*Todo).Remaining)
}

func (list *TodoList) AddTodo(title string) {
	list.todos = append(list.todos, &Todo{
		id:    uniuri.New(),
		title: title,
		list:  list,
	})
	list.changed()
}

func (list *TodoList) ClearCompleted() {
	list.todos = list.Remaining()
	list.changed()
}

func (list *TodoList) FindById(id string) *Todo {
	if todos := list.filter(todoById(id)); len(todos) > 0 {
		return todos[0]
	}
	return nil
}

func (list *TodoList) DeleteById(id string) {
	list.todos = list.filter(invert(todoById(id)))
	list.changed()
}

func (list TodoList) filter(f func(*Todo) bool) []*Todo {
	results := []*Todo{}
	for _, todo := range list.todos {
		if f(todo) {
			results = append(results, todo)
		}
	}
	return results
}

func invert(f func(*Todo) bool) func(*Todo) bool {
	return func(todo *Todo) bool {
		return !f(todo)
	}
}

func todoById(id string) func(*Todo) bool {
	return func(t *Todo) bool {
		return t.id == id
	}
}
