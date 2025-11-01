package cat

import "errors"

type Cat struct {
	Name string
}

// New creates a new Cat instance with the given name.
func New(name string) (*Cat, error) {
	if name == "" {
		return nil, errors.New("dog name cannot be empty")
	}

	if len(name) > 50 {
		return nil, errors.New("dog name cannot exceed 50 characters")
	}

	return &Cat{Name: name}, nil
}

func (c *Cat) GetName() string {
	return c.Name
}

func (c *Cat) Speak() string {
	return "Meow! My name is " + c.Name
}

func (c *Cat) Eat() string {
	return c.Name + " is eating fish."
}

func (c *Cat) Run() string {
	return c.Name + " is running gracefully."
}
