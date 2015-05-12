package main

import (
	"github.com/albrow/temple-example/public/go/controllers"
	"github.com/soroushjp/humble/router"
	"log"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	log.Println("Starting...")

	peopleCtrl := controllers.People{}

	r := router.New()
	r.HandleFunc("/people", peopleCtrl.Index)
	r.HandleFunc("/people/{id}", peopleCtrl.Show)
	r.InterceptLinks()
	r.Start()
}
