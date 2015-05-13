package models

import (
	"github.com/albrow/zoom"
	"github.com/soroushjp/humble/detect"
	"log"
)

var People *zoom.ModelType

func init() {
	if detect.IsServer() {
		if err := RegisterAll(); err != nil {
			log.Fatal(err)
		}
		if err := zoom.Init(nil); err != nil {
			log.Fatal(err)
		}
	}
}

func RegisterAll() error {
	var err error
	People, err = zoom.Register(&Person{})
	return err
}

type Person struct {
	Name string
	Age  int
	zoom.RandomId
}

func (p Person) GetId() string {
	return p.Id
}

func (p Person) RootURL() string {
	return "http://localhost:3000/people"
}
