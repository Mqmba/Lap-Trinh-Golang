package service

// Animal defines a common interface for all animals.
type Animal interface {
	Speak() string
	GetName() string
}

type AnimalPlus interface {
	Animal
	Eat() string
	Run() string
}
