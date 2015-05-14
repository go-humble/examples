package main

import (
	"github.com/go-humble/examples/people/client/controllers"
	"github.com/go-humble/router"
	"log"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	log.Println("Starting...")

	peopleCtrl := controllers.People{}

	r := router.New()
	r.HandleFunc("/people", peopleCtrl.Index)
	r.HandleFunc("/people/{id}", peopleCtrl.Show)
	r.ShouldInterceptLinks = true
	r.Start()
}
