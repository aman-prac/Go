// package main

// import "fmt"

// type circle struct {
// 	radius float64
// }

// type square struct {
// 	side float64
// }

// func (C *circle) area() float64 {
// 	return 3.14 * C.radius * C.radius
// }

// func (C circle) perimeter() float64 {
// 	return 2 * 3.14 * C.radius
// }

// func (S square) area() float64 {
// 	return S.side * S.side
// }

// func (S square) perimeter() float64 {
// 	return 4 * S.side
// }

// type areaCalculator interface {
// 	area()
// }

// type perimeterCalculator interface {
// 	perimeter()
// }

// func printArea(a areaCalculator) {
// 	fmt.Println("Area:", a.area())
// }

// func printPerimeter(p perimeterCalculator) {
// 	fmt.Println("Perimeter:", p.perimeter())
// }

// func main() {
// 	c1 := circle{radius: 5.0}
// 	s1 := square{side: 4.0}

// 	// ✅ square has area with value receiver
// 	printArea(s1)

// 	// ❌ circle doesn't satisfy areaCalculator (value)
// 	// printArea(c1) // error

// 	// ✅ *circle does satisfy areaCalculator
// 	printArea(&c1)

// 	// ✅ both circle and *circle have perimeter (value receiver)
// 	printPerimeter(c1)
// 	printPerimeter(&c1)

//		// ✅ square and *square also satisfy perimeter
//		printPerimeter(s1)
//		printPerimeter(&s1)
//	}
package main

import "fmt"

type circle struct {
	radius float64
}

type square struct {
	side float64
}

func (C *circle) area() float64 {
	return 3.14 * C.radius * C.radius
}

func (C circle) perimeter() float64 {
	return 2 * 3.14 * C.radius
}

func (S square) area() float64 {
	return S.side * S.side
}

func (S square) perimeter() float64 {
	return 4 * S.side
}

type areaCalculator interface {
	area() float64
}

type perimeterCalculator interface {
	perimeter() float64
}

func printArea(a areaCalculator) {
	fmt.Println("Area:", a.area())
}

func printPerimeter(p perimeterCalculator) {
	fmt.Println("Perimeter:", p.perimeter())
}

func main() {
	c1 := circle{radius: 5.0}
	s1 := square{side: 4.0}

	printArea(s1)

	// printArea(c1)

	printArea(&c1)

	printPerimeter(c1)
	printPerimeter(&c1)

	printPerimeter(s1)
	printPerimeter(&s1)
}
