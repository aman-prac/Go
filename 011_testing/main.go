// This package demonstrates how to calculate factorial using both recursion and iteration in Go.
// It includes functions for both methods and a main function to demonstrate their usage.
package main

import (
	"fmt"
)

// Factorial function using loop
// This function calculates the factorial of a number using a loop.
func FactLoop(n int) int {
	fact := 1
	for i := 2; i <= n; i++ {
		fact *= i
	}
	return fact
}

// Factorial function using recursion
// This function calculates the factorial of a number using recursion.
func FactRec(n int) int {
	if n == 0 {
		return 1
	}
	return n * FactRec(n-1)
}

// // Test cases for factorial functions-- loop
// func ExampleFactLoop() {
// 	n := 5
// 	result := factLoop(n)
// 	fmt.Println(result)
// 	// Output: 120
// }

// // Test cases for factorial functions-- recursion
// func ExampleFactRec() {
// 	n := 5
// 	result := factRec(n)
// 	fmt.Println(result)
// 	// Output: 120
// }

func main() {
	n := 11
	fmt.Println("Factorial of", n, "using recursion:", FactRec(n))
	fmt.Println("Factorial of", n, "using loop:", FactLoop(n))
}
