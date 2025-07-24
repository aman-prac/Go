package main

import "fmt"

func doSomething() error {
	return fmt.Errorf("an error occurred")
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("No error occurred")
	}
}

func main() {
	if err := doSomething(); err != nil {
		handleError(err)
	}
}
