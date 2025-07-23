package myModules

import "fmt"

type Solution struct {
	qName        string
	qNumber      int
	NoOfStudents int
}

func PrintAman() {
	println("Aman Girdhar")
	println("Aman Girdhar is a software engineer")
	println("Aman Girdhar is a Go developer")
	println("Aman Girdhar is a full stack developer")
	println("Aman Girdhar is a cloud engineer\n \n")
}

func PrintAmanWithAge(age int) {
	println("Aman Girdhar is", age, "years old")
	println("Aman Girdhar is a software engineer")
	println("Aman Girdhar is a Go developer")
	println("Aman Girdhar is a full stack developer")
	println("Aman Girdhar is a cloud engineer\n \n")
}

func PrintAmanWithAgeAndLocation(age int, location string) {
	println("Aman Girdhar is", age, "years old")
	println("Aman Girdhar lives in", location)
	println("Aman Girdhar is a software engineer")
	println("Aman Girdhar is a Go developer")
	println("Aman Girdhar is a full stack developer")
	println("Aman Girdhar is a cloud engineer")
	println("Aman Girdhar is a DevOps engineer")
	println("Aman Girdhar is a data engineer")
	println("Aman Girdhar is a machine learning engineer")
	println("Aman Girdhar is a data scientist\n \n")
}

func SliceArray() {
	Aman := make(map[string][]string)
	mp := map[string]int{
		"aman":      21,
		"girdhar":   22,
		"gopher":    23,
		"golang":    24,
		"developer": 25,
		"engineer":  26,
		"fullstack": 27,
		"cloud":     28,
	}
	fmt.Println("The age of aman is", mp["aman"])
	fmt.Printf("%T \n", mp)
	Aman["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	Aman["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	Aman["no_dr"] = []string{"cats", "ice cream", "sunsets"}
	fmt.Println("Aman:", Aman)

}

// `bond_james`: `shaken, not stirred`, `martinis`, `fast cars`,
// `moneypenny_jenny`: `intelligence`, `literature`, `computer science`,
// `no_dr`:`cats`, `ice cream`, `sunsets`,
