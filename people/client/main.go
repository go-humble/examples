package main

import (
	"log"

	"github.com/go-humble/examples/people/client/controllers"
	"github.com/go-humble/router"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	log.Println("Starting...")

	r := router.New()
	peopleCtrl := controllers.People{
		Router: r,
	}
	r.HandleFunc("/people/new", peopleCtrl.New)
	r.HandleFunc("/people", peopleCtrl.Index)
	r.HandleFunc("/people/{id}", peopleCtrl.Show)
	r.ShouldInterceptLinks = true
	r.Start()
}
