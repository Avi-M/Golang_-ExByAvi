package main

import "fmt"

func main() {
	var n, k int
	fmt.Println("enter no of row")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		k = 0
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k != 2*i-1 {
			fmt.Print("*")
			k++
		}
		fmt.Println()

	}

}
