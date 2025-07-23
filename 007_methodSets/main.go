package main

import (
	"fmt"
)

// type Person struct {
// 	name     string
// 	age      int
// 	location string
// }

// func (p Person) Speak() {
// 	fmt.Println("I am speaking")
// }

// func (p *Person) Work() {
// 	fmt.Println("I am working")
// }

// type Speaker interface {
// 	Speak()
// }

// type Worker interface {
// 	Work()
// }

func main() {

	// p := Person{name: "Alice", age: 30, location: "Wonderland"}
	// fmt.Println("Person:", p)

	// // Using method sets
	// var s Speaker = p
	// s.Speak()

	// var w Worker = &p
	// w.Work()
	// if wp, ok := w.(*Person); ok {
	// 	wp.Speak()
	// // Demonstrating method sets
	// fmt.Println("Method sets:")

	fmt.Println("Speaker interface can call Speak method")

}
