package models

import (
	"log"

	"github.com/albrow/zoom"
	"github.com/go-humble/detect"
)

var (
	People *zoom.Collection
	pool   *zoom.Pool
)

func init() {
	if detect.IsServer() {
		pool = zoom.NewPool("localhost:6379")
		var err error
		colOptions := zoom.DefaultCollectionOptions.WithIndex(true)
		People, err = pool.NewCollectionWithOptions(&Person{}, colOptions)
		if err != nil {
			log.Fatal(err)
		}
		if err := CreateInitialPeople(); err != nil {
			log.Fatal(err)
		}
	}
}

func ClosePool() error {
	return pool.Close()
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	zoom.RandomId
}

func (p Person) RootURL() string {
	return "/people"
}

func CreateInitialPeople() error {
	if count, err := People.Count(); err != nil {
		return err
	} else if count == 0 {
		log.Println("Creating people...")
		t := pool.NewTransaction()
		for i, name := range []string{"Foo", "Bar", "Baz"} {
			person := &Person{
				Age:  i + 20,
				Name: name,
			}
			t.Save(People, person)
		}
		if err := t.Exec(); err != nil {
			return err
		}
	}
	return nil
}
