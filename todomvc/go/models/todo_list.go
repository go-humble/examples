package models

import "github.com/dchest/uniuri"

type TodoList struct {
	todos []*Todo
}

func (list TodoList) All() []*Todo {
	return list.todos
}

func (list TodoList) Completed() []*Todo {
	return list.filter(todoCompleted)
}

func (list TodoList) Remaining() []*Todo {
	return list.filter(todoRemaining)
}

func (list *TodoList) AddTodo(title string) {
	list.todos = append(list.todos, &Todo{
		Id:    uniuri.New(),
		Title: title,
	})
}

func (list *TodoList) ClearCompleted() {
	list.todos = list.filter(todoRemaining)
}

func (list *TodoList) FindById(id string) (*Todo, bool) {
	if todos := list.filter(todoById(id)); len(todos) > 0 {
		return todos[0], true
	}
	return &Todo{}, false
}

func (list *TodoList) DeleteById(id string) {
	for i, todo := range list.todos {
		if todo.Id == id {
			list.todos = append(list.todos[:i], list.todos[i+1:]...)
			break
		}
	}
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

func todoCompleted(t *Todo) bool {
	return t.Completed
}

func todoRemaining(t *Todo) bool {
	return !t.Completed
}

func todoById(id string) func(*Todo) bool {
	return func(t *Todo) bool {
		return t.Id == id
	}
}
