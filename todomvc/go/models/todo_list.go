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

func (list *TodoList) DeleteById(id string) {
	list.todos = list.Filter(todoNotById(id))
	list.changed()
}
