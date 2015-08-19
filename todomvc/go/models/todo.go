package models

type Todo struct {
	Id        string
	Completed bool
	Title     string
}

func (t *Todo) Toggle() {
	t.Completed = !t.Completed
}
