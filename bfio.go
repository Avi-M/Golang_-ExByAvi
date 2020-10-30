package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, e := os.OpenFile("test1.txt", os.O_WRONLY, 0666)
	if e != nil {
		fmt.Println(e)
	}
	bw := bufio.NewWriterSize(f, 4)
	bw.Write([]byte("H"))
	bw.Write([]byte("E"))
	bw.Write([]byte("L"))
	bw.Write([]byte("L"))
	bw.Write([]byte("o"))
	bw.Write([]byte(" "))
	bw.Write([]byte("W"))
	bw.Write([]byte("o"))
	bw.Write([]byte("r"))
	bw.Write([]byte("l"))
	bw.Write([]byte("d"))
	fmt.Println(bw.Buffered())
	fmt.Println(bw.Available())
}
