package main

import "fmt"

func lyThuyetPointer() {
	name := "Hello world"
	fmt.Println("-=-=-=-=-Infomation name variable-=-=-=-=-")
	fmt.Printf("Type: %T\n", name)
	fmt.Printf("Value: %v\n", name)
	fmt.Printf("Address: %p\n", &name)
	fmt.Println()

	// Create a pointer to the name variable
	pointerToName := &name
	fmt.Println("-=-=-=-=-Infomation pointerToName variable-=-=-=-=-")
	fmt.Printf("Type: %T\n", pointerToName)
	fmt.Printf("Value: %v\n", pointerToName)
	fmt.Printf("Address: %p\n", &pointerToName)
	fmt.Printf("Find value name from pointerToName %v\n", *pointerToName)
	fmt.Println()

	pointerToName2 := &pointerToName
	fmt.Println("-=-=-=-=-Infomation pointerToName2 variable-=-=-=-=-")
	fmt.Printf("Type: %T\n", pointerToName2)
	fmt.Printf("Value: %v\n", pointerToName2)
	fmt.Printf("Address: %p\n", &pointerToName2)
	fmt.Printf("Find value name from pointerToName2 %v\n", **pointerToName2)
	fmt.Println()
}

func updateName(name *string) {
	// fmt.Println("Show: ", name)
	*name = "Nguyen Van B"
	fmt.Println("-=-=-=-=-Infomation name variable inside updateName-=-=-=-=-")
	fmt.Printf("Type: %T\n", name)
	fmt.Printf("Value: %v\n", *name)
	fmt.Printf("Address: %v\n", name)
	fmt.Println()
}

func main() {
	name := "Nguyen Van A"
	fmt.Println("-=-=-=-=-Infomation name variable-=-=-=-=-")
	fmt.Printf("Type: %T\n", name)
	fmt.Printf("Value: %v\n", name)
	fmt.Printf("Address: %v\n", &name)
	fmt.Println()

	name = "Tony Stark"
	fmt.Println("-=-=-=-=-Infomation name variable after update value-=-=-=-=-")
	fmt.Printf("Type: %T\n", name)
	fmt.Printf("Value: %v\n", name)
	fmt.Printf("Address: %v\n", &name)
	fmt.Println()

	updateName(&name)

	fmt.Println("-=-=-=-=-Infomation name variable after call updateName function-=-=-=-=-")
	fmt.Printf("Type: %T\n", name)
	fmt.Printf("Value: %v\n", name)
	fmt.Printf("Address: %v\n", &name)
	fmt.Println()
}
