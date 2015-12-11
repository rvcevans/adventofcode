package main

import (
	"github.com/ifross89/adventofcode/getinput"

	"bytes"
	"fmt"
	"os"
)

func sayNumbers(in []byte) []byte{
	out := bytes.Buffer{}
	currentByte := in[0]
	byteCount := byte(0)

	for _, b := range in {
		if b == currentByte {
			byteCount++
		} else {
			out.WriteByte(byteCount + '0')
			out.WriteByte(currentByte)
			currentByte = b
			byteCount = 1
		}
	}
	out.WriteByte(byteCount + '0')
	out.WriteByte(currentByte)
	return out.Bytes()
}

func main() {

	input := []byte(getinput.MustGet(10, os.Getenv("ADVENT_SESSION"))[0])
	fmt.Printf("Input: %s\n", input)

	for i := 0; i < 50; i++ {
		input = sayNumbers(input)
		if i == 39 {
			fmt.Println("Length at 40:", len(input))
		}
	}

	fmt.Println("Length at 50:", len(input))
}
