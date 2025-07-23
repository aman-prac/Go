package main

import (
	"fmt"
)

func main() {
	// for i := 0; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i, "is even")
	// 	} else {
	// 		fmt.Println(i, "is odd")
	// 	}
	// }

	// i := 5
	// for i < 50 {
	// 	if i == 25 {
	// 		i += 5
	// 		continue
	// 	}

	// 	fmt.Println("value is", i)
	// 	i += 5

	// }

	xi := []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100}

	for i, v := range xi {
		fmt.Printf("(%d,%d)\n", i, v)
	}
}
