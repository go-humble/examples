package models

import (
	"github.com/albrow/zoom"
)

var People *zoom.ModelType

func RegisterAll() error {
	var err error
	People, err = zoom.Register(&Person{})
	return err
}

type Person struct {
	Name string
	Age  int
	zoom.DefaultData
}

func (p *Person) GetId() string {
	return p.Id()
}

func (p *Person) RootURL() string {
	return "localhost:3000/persons"
}
