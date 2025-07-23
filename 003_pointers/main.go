package main

import (
	"fmt"
)

func main() {
	var a int = 11
	p := &a
	fmt.Println("The address of a is ", p, "and the value of a is", *p)
	fmt.Printf("The type of p is %T \n", p)
}
