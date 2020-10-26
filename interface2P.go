package main

import (
	"fmt"
	"math"
)

type shape interface {
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
func CalulateTotalArea(shapes ...shape) float64 {
	totalArea := 0.0
	for _, val := range shapes {
		totalArea += val.Area()
	}
	return totalArea
}
func main() {
	totalAr := CalulateTotalArea(Rectangle{3.5, 6}, Circle{7.5})
	fmt.Println(totalAr)

}
