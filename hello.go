package main

import (
	"fmt"
	"log"
)

type student struct {
	name    string
	age     int
	country string
}
type teacher struct {
	name    string
	age     int
	country string
}

func (s student) speak() {
	fmt.Println("My name is", s.name, "I am", s.age, "years old and I am from", s.country, "and I am a student")
}

func (t teacher) speak() {
	fmt.Println("My name is", t.name, "I am", t.age, "years old and I am from", t.country, "and I am a teacher")
}

type person interface {
	speak()
}

func saySomething(p person) {
	p.speak()
	fmt.Println("I am a person I can speak through the interface in Go")
}

func (s student) String() string {
	return fmt.Sprintf("My Name: %s, My Age: %d, My Country: %s and I am a student\n", s.name, s.age, s.country)
}
func (t teacher) String() string {
	return fmt.Sprintf("My Name: %s, My Age: %d, My Country: %s and I am a teacher\n", t.name, t.age, t.country)
}

type house int

func (h house) String() string {
	return fmt.Sprintf("House Number: %d\n", h)
}

func customPrint(s fmt.Stringer) {
	log.Printf("ERROR 128 %s", s.String())
}

func main() {
	// fmt.Print("Hello, World!\n")
	// fmt.Println(gitTest.SwapString(("Hello"), ("Gophers")))
	// myModules.PrintAman()
	// myModules.PrintAmanWithAge(21)
	// myModules.PrintAmanWithAgeAndLocation(21, "India")
	// myModules.SliceArray()
	aman := student{
		name:    "Aman",
		age:     21,
		country: "India",
	}
	// aman.speak()
	john := teacher{
		name:    "John",
		age:     30,
		country: "USA",
	}
	// john.speak()

	saySomething(aman)
	saySomething(john)
	fmt.Println("This is a demonstration of interfaces in Go")
	var h1 house = 100
	// fmt.Printf("___________________________________________________________\n")
	// fmt.Println(aman)
	// fmt.Println(john)
	// fmt.Println(h1)
	/*
		------------------------------------------------------------------------------------------------------------------
	*/
	// fmt.Printf("Default: %v\n", aman)
	// fmt.Printf("With fields: %+v\n", aman)
	// fmt.Printf("Go syntax: %#v\n", aman)
	// fmt.Printf("Type: %T\n", aman)
	customPrint(aman)
	customPrint(john)
	customPrint(h1)
	fmt.Printf("___________________________________________________________\n")
}
