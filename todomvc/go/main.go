package main

import (
	"log"

	"github.com/go-humble/examples/todomvc/go/models"
	"github.com/go-humble/examples/todomvc/go/views"
)

//go:generate temple build templates/templates templates/templates.go --partials templates/partials
//go:generate gopherjs build main.go -o ../js/app.js -m

func main() {
	log.Println("Starting")
	todos := &models.TodoList{}
	if err := todos.Load(); err != nil {
		panic(err)
	}
	appView := views.NewApp(todos)
	if err := appView.Render(); err != nil {
		panic(err)
	}
}
