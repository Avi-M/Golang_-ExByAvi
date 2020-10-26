package main

import (
	"fmt"
	"math"
)

type shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	l, b float64
}

func (r Rectangle) Area() float64 {
	return r.l * r.b

}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.l + r.b)

}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	var a shape
	a = Rectangle{3, 5}
	fmt.Println(a.Area())
	fmt.Println(a.Perimeter())
	fmt.Printf("type:%T\n", a)
	a = Circle{3.5}
	fmt.Println(a.Perimeter())
	fmt.Println(a.Area())
	fmt.Printf("type:%T\n", a)
}
