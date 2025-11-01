package main

import (
	"fmt"

	"mamba.com/interface/cat"
	"mamba.com/interface/dog"
	"mamba.com/interface/mouse"
	"mamba.com/interface/service"
)

// func makeSound(a service.Animal) {
// 	fmt.Println("Animal Name:", a.GetName())
// 	fmt.Println(a.Speak())
// }

func makeSoundPlus(p service.AnimalPlus) {
	fmt.Println("Animal Name:", p.GetName())
	fmt.Println(p.Speak())
	fmt.Println(p.Eat())
	fmt.Println(p.Run())
}

func printValue(v interface{}) {
	// str, ok := v.(string)
	// if ok {
	// 	fmt.Println("String value:", str)
	// } else {
	// 	fmt.Println("Not a string")
	// }

	// num, ok := v.(int)
	// if ok {
	// 	fmt.Println("Integer value:", num)
	// } else {
	// 	fmt.Println("Not an integer")
	// }

	// Using type switch
	switch val := v.(type) {
	case string:
		fmt.Println("String value:", val)
	case int:
		fmt.Println("Integer value:", val)
	case bool:
		fmt.Println("Boolean value:", val)
	case float64:
		fmt.Println("Float value:", val)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	myDog, err := dog.New("Buddy")
	if err != nil {
		panic(err)
	}
	// makeSound(myDog)
	makeSoundPlus(myDog)

	fmt.Println("----------")

	myCat, err := cat.New("Whiskers")
	if err != nil {
		panic(err)
	}
	// makeSound(myCat)
	makeSoundPlus(myCat)

	fmt.Println("----------")

	myMouse, err := mouse.New("Jerry")
	if err != nil {
		panic(err)
	}
	// makeSound(myMouse)
	makeSoundPlus(myMouse)

	fmt.Println("----------")

	printValue("Hello, World!")
	printValue(42)
	printValue(3.14)
	printValue(false)
}
