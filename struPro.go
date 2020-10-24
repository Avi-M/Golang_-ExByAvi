package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	fmt.Println(Point{1, 2})
	p := Point{5, 7}
	fmt.Println(p)
	p.X = 8
	fmt.Println(p)
	ptr := &p
	ptr.Y = 9
	fmt.Println(p)
	p2 := Point{X: 2}
	p3 := Point{}
	fmt.Println(p2, p3, ptr)

}
