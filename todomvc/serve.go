package main

import (
	"log"
	"net/http"
)

func main() {
	port := ":8000"
	log.Println("Serving on http://localhost" + port)
	http.ListenAndServe(port, http.FileServer(http.Dir(".")))
}
