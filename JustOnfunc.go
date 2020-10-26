package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func VariadicEx(s ...string) {
	fmt.Println(s[2])
	fmt.Println(s[3])
}
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)

}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func CalculationEx(str string, y ...int) int {
	area := 1
	for _, val := range y {
		if str == "rectangle" {
			area *= val
		} else if str == "square" {
			area = val * val
		}
	}
	return area
}

func main() {
	fmt.Println(CalculationEx("square", 20))
	fmt.Println(CalculationEx("rectangle", 20, 30))

	VariadicEx("Ram", "Laksh", "Luv", "Kush")

	v := Vertex{4, 3}
	fmt.Println(Abs(v))
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
