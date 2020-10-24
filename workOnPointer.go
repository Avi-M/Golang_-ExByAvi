package main

import "fmt"

func main() {
	i, j := 42, 4000
	var p *int
	p = &i
	fmt.Println(*p)
	*p = 43
	fmt.Println(i)
	fmt.Println(p)
	p = &j
	*p = *p / 5
	fmt.Println(j)
	fmt.Println(p)
}
