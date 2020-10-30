package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	f, e := os.Open("buffertest.txt")
	CheckError(e)

	br := bufio.NewReader(f)

	bbuf := make([]byte, 10)
	bbuf, e = br.Peek(6)
	CheckError(e)

	fmt.Println(string(bbuf))

	nr, e := br.Read(bbuf)
	CheckError(e)

	fmt.Println("Num bytes read", nr)

	singleByte, e := br.ReadByte()
	CheckError(e)

	fmt.Println("Single byte is", string(singleByte))

	br.Reset(f)
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
