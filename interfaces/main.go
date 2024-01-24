package main

import (
	"fmt"
	"math"
)

// Shape interface declered
type Shape interface {
	Area() float64
}

// Rectangle struct declered
type Rectangle struct {
	Width, Height float64
}

// Circle struct declered
type Circle struct {
	Radius float64
}

// declered method for Rectangle struct
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// declered method for Area struct
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// getArea function
func getArea(shape Shape) {

	fmt.Println(shape.Area())
}

func main() {
	// create structs
	r := Rectangle{Width: 7, Height: 8}
	c := Circle{Radius: 5}
	// run the getArea() function
	getArea(r)
	getArea(c)
}
