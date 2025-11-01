package main

import "fmt"

func calculator(a int, b int) (int, int, int, float32) {
	sum := a + b
	diff := a - b
	prod := a * b
	quot := float32(a) / float32(b)
	return sum, diff, prod, quot
}

func countdown(n int) {
	if n > 0 {
		fmt.Println(n)
		countdown(n - 1)
	}
}

func main() {
	a, b := 10, 5
	sum, diff, prod, quot := calculator(a, b)
	fmt.Printf("Sum: %d, Difference: %d, Product: %d, Quotient: %.2f\n", sum, diff, prod, quot)
	countdown(50)
}
