package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileStat, err := os.Stat("test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileStat.Name())
	fmt.Println(fileStat.Size())
	fmt.Println("Permissions:", fileStat.Mode())
	fmt.Println("last modified:", fileStat.ModTime())
	fmt.Println(fileStat.IsDir())
}
