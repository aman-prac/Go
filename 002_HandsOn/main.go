package main

import "math"

type rectangle struct {
	length float64
	width  float64
}

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func calcArea(s shape) float64 {
	return s.area()
}

func main() {
	r := rectangle{
		length: 5,
		width:  3,
	}
	s := square{
		side: 5,
	}
	c := circle{
		radius: 4,
	}

	println("Area of rectangle:", calcArea(r))
	println("Area of square:", calcArea(s))
	println("Area of circle:", calcArea(c))

}
