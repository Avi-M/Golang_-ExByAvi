package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Stat("Emp")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("Emp", 4880)
		if errDir != nil {
			log.Fatal(err)
		}

	}
}
