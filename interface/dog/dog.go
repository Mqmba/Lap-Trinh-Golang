package dog

import (
	"errors"
	"strings"
)

type Dog struct {
	Name string
}

// New creates a new Dog instance with the given name.
func New(name string) (*Dog, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return nil, errors.New("dog name cannot be empty")
	}

	if len(name) > 50 {
		return nil, errors.New("dog name cannot exceed 50 characters")
	}

	return &Dog{Name: name}, nil
}

func (d *Dog) GetName() string {
	return d.Name
}

func (d *Dog) Speak() string {
	return "Woof! My name is " + d.Name
}

func (d *Dog) Eat() string {
	return d.Name + " is eating dog food."
}

func (d *Dog) Run() string {
	return d.Name + " is running on all fours."
}
