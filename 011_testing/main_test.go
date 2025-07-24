package main

import (
	"testing"
)

func TestFactLoop(t *testing.T) {
	n := 5
	expected := 120
	result := FactLoop(n)
	if result != expected {
		t.Errorf("factLoop(%d) = %d; want %d", n, result, expected)
	}
}

func TestFactRec(t *testing.T) {
	n := 5
	expected := 120
	result := FactRec(n)
	if result != expected {
		t.Errorf("factRec(%d) = %d; want %d", n, result, expected)
	}
}

func BenchmarkFactLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactLoop(20)
	}
}

func BenchmarkFactRec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactRec(20)
	}
}

// /implement a table test for the factorial functions
func TestFactorialFunctions(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{"Factorial of 0", 0, 1},
		{"Factorial of 1", 1, 1},
		{"Factorial of 2", 2, 2},
		{"Factorial of 3", 3, 6},
		{"Factorial of 4", 4, 24},
		{"Factorial of 5", 5, 120},
	}

	for _, tt := range tests {
		t.Run(tt.name+" using loop", func(t *testing.T) {
			result := FactLoop(tt.n)
			if result != tt.expected {
				t.Errorf("factLoop(%d) = %d; want %d", tt.n, result, tt.expected)
			}
		})

		t.Run(tt.name+" using recursion", func(t *testing.T) {
			result := FactRec(tt.n)
			if result != tt.expected {
				t.Errorf("factRec(%d) = %d; want %d", tt.n, result, tt.expected)
			}
		})
	}
}

// func ExampleFactLoop() {
// 	n := 5
// 	result := FactLoop(n)
// 	fmt.Println(result)
// 	// Output: 120
// }

// func ExampleFactRec() {
// 	n := 5
// 	result := FactRec(n)
// 	fmt.Println(result)
// 	// Output: 120
// }
