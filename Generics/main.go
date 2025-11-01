package main

import (
	"cmp"
	"fmt"

	"golang.org/x/exp/constraints"
)

type Box[T any] struct {
	Content     T
	Description T
}

type Number interface {
	constraints.Integer | constraints.Float
}

func sum[T Number](a, b T) T {
	return a + b
}

func printValue(value string) {
	fmt.Println(value)
}

func isEqual[T comparable](a, b T) bool {
	return a == b
}

func isNotEqual[T comparable](a, b T) bool {
	return a != b
}

func max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	result := isEqual(53, 5)
	printValue(fmt.Sprintf("Are the numbers equal? %v", result))

	resultStr := isEqual("hello", "hello")
	printValue(fmt.Sprintf("Are the strings equal? %v", resultStr))

	resultFloat := isEqual(3.14, 2.71)
	printValue(fmt.Sprintf("Are the floats equal? %v", resultFloat))

	resultNotEqual := isNotEqual(10, 20)
	printValue(fmt.Sprintf("Are the numbers not equal? %v", resultNotEqual))

	maxValue := max(30.2, 20.5)
	printValue(fmt.Sprintf("The maximum value is: %v", maxValue))

	stringBox := Box[string]{Content: "Golang Generics", Description: "A box for strings"}
	intBox := Box[int]{Content: 100, Description: 100}
	printValue(fmt.Sprintf("Boxed string: %s, Description: %s", stringBox.Content, stringBox.Description))
	printValue(fmt.Sprintf("Boxed integer: %d, Description: %d", intBox.Content, intBox.Description))

	sumResult := sum(15.5, 25)
	printValue(fmt.Sprintf("The sum is: %v", sumResult))
}
