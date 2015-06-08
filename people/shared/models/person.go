package models

import (
	"github.com/albrow/zoom"
	"github.com/go-humble/detect"
	"log"
)

var (
	People *zoom.ModelType
	pool   *zoom.Pool
)

func init() {
	if detect.IsServer() {
		pool = zoom.NewPool(nil)
		var err error
		People, err = pool.Register(&Person{})
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
	Name string
	Age  int
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
