package main

import (
	"Go/gitTest"
	"Go/myModules"
	"fmt"
)

func main() {
	fmt.Print("Hello, World!\n")
	fmt.Println(gitTest.SwapString(("Hello"), ("Gophers")))
	myModules.PrintAman()
	myModules.PrintAmanWithAge(21)
	myModules.PrintAmanWithAgeAndLocation(21, "India")
}
