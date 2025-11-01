package main

import "fmt"

type Employee struct {
	ID   int
	Name string
	Age  int
}

func Array() {

	employees := [...]Employee{
		{ID: 1, Name: "John Doe", Age: 30},
		{ID: 2, Name: "Jane Smith", Age: 25},
		{ID: 3, Name: "Alice Johnson", Age: 28},
		{ID: 4, Name: "Bob Brown", Age: 35},
	}

	for index, employee := range employees {
		fmt.Printf("Employee %d: ID=%d, Name=%s, Age=%d\n", index+1, employee.ID, employee.Name, employee.Age)
	}
	// var number = [...]int{10, 20, 30, 40, 50}
	// for index, value := range number {
	// 	fmt.Printf("Index: %d, Value: %d\n", index, value)
	// }

	// fmt.Println("-----")

	// for i := 0; i < len(number); i++ {
	// 	fmt.Printf("Index: %d, Value: %d\n", i, number[i])
	// }

	// fmt.Println(number[3])

	// var matrix = [3][4]int{
	// 	{10, 20, 30, 40},
	// 	{50, 60, 70, 80},
	// 	{90, 100, 110, 120},
	// }
	// for key, value := range matrix {
	// 	fmt.Printf("Array number[%d]: %v\n", key, value)
	// }

}
