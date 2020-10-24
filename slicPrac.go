package main

import "fmt"

func printSlice(s []int) {
	fmt.Println(len(s), cap(s), s)
}

func slicePrint1(s string, x []int) {

	fmt.Println(s, len(x), cap(x), x)
}
func main() {
	var slice []int
	printSlice(slice)
	slice = append(slice, 0)
	printSlice(slice)
	slice = append(slice, 1)
	printSlice(slice)
	slice = append(slice, 1, 2, 3, 4)
	printSlice(slice)

	a := make([]int, 5)
	slicePrint1("a", a)
	b := make([]int, 0, 5) //len(b)=0,cap(b)=5
	slicePrint1("b", b)
	c := b[:2]
	slicePrint1("c", c)
	d := b[2:5]
	slicePrint1("d", d)
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s[:0])
	printSlice(s[3:])
	printSlice(s[2:5])
}
