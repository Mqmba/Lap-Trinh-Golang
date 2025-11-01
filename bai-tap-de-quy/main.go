package main

import (
	"fmt"
	"strconv"
)

// Recursive function to calculate the sum from 1 to n
func add(n int) int {
	if n == 1 {
		return 1
	}
	return n + add(n-1)
}

// Recursive function to calculate the nth Fibonacci number
func fibonacci(n int) int {
	if n < 0 {
		return -1 // Invalid input
	} else if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Recursive function to calculate the total of Fibonacci numbers up to the nth number
func total_fibonacci(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n) + total_fibonacci(n-1)
}

func main() {
	for {
		fmt.Println("======================================")
		fmt.Println("Select an option:")
		fmt.Println("[1] Sum from 1 to n")
		fmt.Println("[2] Fibonacci sequence up to n")
		fmt.Println("[3] Total of Fibonacci numbers up to n")
		fmt.Println("[0] Exit")
		fmt.Println("======================================")

		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Please enter a valid number")
			continue
		}

		switch choice {
		case 1:
			var n int
			fmt.Print("Enter a number: ")
			_, err := fmt.Scan(&n)
			if err != nil || n < 0 {
				fmt.Println("Invalid input")
				continue
			}
			fmt.Printf("Total from 1 to %d: %d\n", n, add(n))

		case 2:
			str := ""
			var n int
			fmt.Print("Enter a number: ")
			_, err := fmt.Scan(&n)
			if err != nil || n < 0 {
				fmt.Println("Invalid input")
				continue
			}
			for i := 0; i < n; i++ {
				str += strconv.Itoa(fibonacci(i)) + " "
			}
			fmt.Println("Fibonacci sequence:", str)

		case 3:
			var n int
			fmt.Print("Enter a number: ")
			_, err := fmt.Scan(&n)
			if err != nil || n < 0 {
				fmt.Println("Invalid input")
				continue
			}
			fmt.Printf("Total of Fibonacci numbers up to %d: %d\n", n, total_fibonacci(n-1))

		case 0:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
			continue
		}
	}
}
