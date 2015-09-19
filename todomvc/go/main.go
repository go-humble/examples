package main

import (
	"log"

	"github.com/go-humble/router"

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
	todos.OnChange(func(newTodos *models.TodoList) {
		appView.Todos = newTodos
		go func() {
			if err := appView.Todos.Save(); err != nil {
				panic(err)
			}
		}()
		if err := appView.Render(); err != nil {
			panic(err)
		}
	})

	r := router.New()
	r.ForceHashURL = true
	r.HandleFunc("/", func(_ *router.Context) {
		appView.UseFilter(models.Predicates.All)
		if err := appView.Render(); err != nil {
			panic(err)
		}
	})
	r.HandleFunc("/active", func(_ *router.Context) {
		appView.UseFilter(models.Predicates.Remaining)
		if err := appView.Render(); err != nil {
			panic(err)
		}
	})
	r.HandleFunc("/completed", func(_ *router.Context) {
		appView.UseFilter(models.Predicates.Completed)
		if err := appView.Render(); err != nil {
			panic(err)
		}
	})
	r.Start()
}
