package main

import "fmt"

func main() {
	var n, k int
	fmt.Println("enter value of n")
	fmt.Scanln(&n)
	fmt.println("enter value of k")
	fmt.Println(&k)
	for n<k {
		f:=0
		for i := 2; i < n/2; i++ {
			if n%i==0 {
				f=1
				break
				
			}
			
		}
		if f==0 {
		fmt.Println(n)
		
		}
		n++	
			
		}
		

	}

}