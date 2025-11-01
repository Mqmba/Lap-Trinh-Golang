package mouse

import (
	"errors"
	"strings"
)

type Mouse struct {
	Name string
}

// New creates a new Mouse instance with the given name.
func New(name string) (*Mouse, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return nil, errors.New("mouse name cannot be empty")
	}

	if len(name) > 50 {
		return nil, errors.New("mouse name cannot exceed 50 characters")
	}

	return &Mouse{Name: name}, nil
}

func (m *Mouse) GetName() string {
	return m.Name
}

func (m *Mouse) Speak() string {
	return "Squeak! My name is " + m.Name
}

func (m *Mouse) Eat() string {
	return m.Name + " is eating cheese."
}

func (m *Mouse) Run() string {
	return m.Name + " is running swiftly."
}
