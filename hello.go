package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	a := 5
	b := 6
	c := "Hello"
	d := float32(3.14)
	e := 123.123455667889
	fmt.Printf("%T %T %T %T %T /n", a, b, c, d, e)
	var aman int = 12
	fmt.Printf("%T", aman)
}
