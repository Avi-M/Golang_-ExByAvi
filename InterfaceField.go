package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Rectangle struct {
	l, b float64
}

func (r Rectangle) Area() float64 {
	return r.l * r.b

}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Interface types can also be used as fields
type MyDrawing struct {
	shapes  []Shape
	bgColor string
	fgColor string
}

func (drawing MyDrawing) Area() float64 {
	ar := 0.0
	for _, sp := range drawing.shapes {
		ar += sp.Area()

	}
	return ar

}
func main() {
	drawing := MyDrawing{
		shapes: []Shape{
			Circle{2},
			Rectangle{3, 5},
			Rectangle{4, 7},
		},
		bgColor: "red",
		fgColor: "white",
	}

	fmt.Println("Drawing", drawing)
	fmt.Println("Drawing Area = ", drawing.Area())
}
