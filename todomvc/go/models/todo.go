package models

import "encoding/json"

type Todo struct {
	id        string
	completed bool
	title     string
	list      *TodoList
}

func (t *Todo) Toggle() error {
	t.completed = !t.completed
	return t.list.Save()
}

func (t *Todo) Completed() bool {
	return t.completed
}

func (t *Todo) Remaining() bool {
	return !t.completed
}

func (t *Todo) SetCompleted(completed bool) error {
	t.completed = completed
	return t.list.Save()
}

func (t *Todo) Title() string {
	return t.title
}

func (t *Todo) SetTitle(title string) error {
	t.title = title
	return t.list.Save()
}

func (t *Todo) Id() string {
	return t.id
}

type jsonTodo struct {
	Id        string
	Completed bool
	Title     string
}

func (todo Todo) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonTodo{
		Id:        todo.id,
		Completed: todo.completed,
		Title:     todo.title,
	})
}

func (todo *Todo) UnmarshalJSON(data []byte) error {
	jt := &jsonTodo{}
	if err := json.Unmarshal(data, jt); err != nil {
		return err
	}
	todo.id = jt.Id
	todo.completed = jt.Completed
	todo.title = jt.Title
	return nil
}
