package models

import (
	"github.com/dchest/uniuri"
	"github.com/go-humble/locstor"
)

// store is a datastore backed by localStorage.
var store = locstor.NewDataStore(locstor.JSONEncoding)

// TodoList is a model representing a list of todos.
type TodoList struct {
	todos           []*Todo
	changeListeners []func(*TodoList)
}

// OnChange can be used to register change listeners. Any functions passed to
// OnChange will be called when the todo list changes.
func (list *TodoList) OnChange(f func(*TodoList)) {
	list.changeListeners = append(list.changeListeners, f)
}

// changed is used to notify the todo list and its change listeners of a change.
// Whenever the list is changed, it must be explicitly called.
func (list *TodoList) changed() {
	for _, f := range list.changeListeners {
		f(list)
	}
}

// Load loads the list of todos from the datastore.
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

// Save saves the list of todos to the datastore.
func (list TodoList) Save() error {
	if err := store.Save("todos", list.todos); err != nil {
		return err
	}
	return nil
}

// AddTodo appends a new todo to the list.
func (list *TodoList) AddTodo(title string) {
	list.todos = append(list.todos, &Todo{
		id:    uniuri.New(),
		title: title,
		list:  list,
	})
	list.changed()
}

// ClearCompleted removes all the todos from the list that have been completed.
func (list *TodoList) ClearCompleted() {
	list.todos = list.Remaining()
	list.changed()
}

// ToggleAll toggles all the todos in the list.
func (list *TodoList) ToggleAll() {
	for _, todo := range list.todos {
		todo.completed = !todo.completed
	}
	list.changed()
}

// DeleteById removes the todo with the given id from the list.
func (list *TodoList) DeleteById(id string) {
	list.todos = list.Filter(todoNotById(id))
	list.changed()
}
