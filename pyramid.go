package main

import "fmt"

func main() {
	var n int
	fmt.Println("enter no of row")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()

	}

}
