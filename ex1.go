//Go program to create empty file
package main

import (
	"log"
	"os"
)

func main() {
	emptyfile, err := os.Create("empty.txt")
	if err != nil {
		log.Fatal(err)

	}
	log.Println(emptyfile)
	emptyfile.Close()
}
