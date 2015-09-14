package models

// Filter is a function which operates on a *TodoList and returns a slice of
// todos that match certain criteria.
type Filter func(*TodoList) []*Todo

// Filters is a data structure with commonly used Filters.
var Filters = struct {
	All       Filter
	Completed Filter
	Remaining Filter
}{
	All:       (*TodoList).All,
	Completed: (*TodoList).Completed,
	Remaining: (*TodoList).Remaining,
}

// All is a filter that returns all the todos in the list.
func (list TodoList) All() []*Todo {
	return list.todos
}

// Completed is a filter that returns only the completed todos in the list.
func (list TodoList) Completed() []*Todo {
	return list.filter((*Todo).Completed)
}

// Remaining is a filter that returns only the remaining (or active) todos in
// the list.
func (list TodoList) Remaining() []*Todo {
	return list.filter((*Todo).Remaining)
}

// Filter calls the given function for each todo in the list, and returns a
// slice todos for which the function f returns true.
func (list TodoList) filter(f func(*Todo) bool) []*Todo {
	results := []*Todo{}
	for _, todo := range list.todos {
		if f(todo) {
			results = append(results, todo)
		}
	}
	return results
}

// Invert inverts a function that operates on a todo and returns the inverted
// function. Where f would return true, the inverted function would return false
// and where f would return false, the inverted function would return true.
func invert(f func(*Todo) bool) func(*Todo) bool {
	return func(todo *Todo) bool {
		return !f(todo)
	}
}

// todoById returns a function which returns true iff todo.id equals the given
// id.
func todoById(id string) func(*Todo) bool {
	return func(t *Todo) bool {
		return t.id == id
	}
}
