package models

import (
	"github.com/albrow/zoom"
	"github.com/go-humble/detect"
	"log"
)

var People *zoom.ModelType

func init() {
	if detect.IsServer() {
		if err := RegisterAll(); err != nil {
			log.Fatal(err)
		}
		zoom.Init(nil)
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

func (p Person) RootURL() string {
	return "/people"
}
