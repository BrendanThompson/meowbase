package cat

import (
	"fmt"

	"github.com/imdario/mergo"
)

type Cat struct {
	ID        int    // Unique identifier for the cat
	FirstName string // First name of the cat
	LastName  string // Last name of the cat
	Quote     string // Cats favourite quote
}

type Cats []Cat

var (
	cats   Cats
	nextID = 1
)

func GetCats() Cats {
	return cats
}

func (c *Cat) Add() error {
	c.ID = nextID
	nextID++

	cats = append(cats, *c)
	return nil
}

func (c *Cat) Get(id int) error {
	for _, cat := range cats {
		if cat.ID == id {
			c = &cat
			return nil
		}
	}
	return fmt.Errorf("unable to find cat with ID: %d", id)
}

func (c *Cat) Update(id int) error {
	for i, cat := range cats {
		if cat.ID == id {
			err := mergo.Merge(c, cats[i])
			if err != nil {
				panic(err)
			}
			cats[i] = *c
			return nil
		}
	}
	return fmt.Errorf("unable to update cat with ID: %d", id)
}

func (c *Cat) Delete(id int) error {
	for i, cat := range cats {
		if cat.ID == id {
			cats = cats[:i+copy(cats[i:], cats[i+1:])]
			return nil
		}
	}
	return fmt.Errorf("unable to delete cat with ID: %d", id)
}
