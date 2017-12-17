package main

import (
	"fmt"
	"github.com/rvcevans/adventofcode/getinput"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	mulA        = 16807
	mulB        = 48271
	div         = 1<<31 - 1
	pairCompare = 5000000
)

func main() {
	input := getinput.MustGet(2017, 15, os.Getenv("ADVENT_SESSION"))
	a, b := generator(input[0]), generator(input[1])

	achan, bchan := make(chan int, pairCompare), make(chan int, pairCompare)
	maybeSend := func(c chan int, v int) {
		select {
		case c <- v:
		default:
		}
	}

	pairs := 0
	total1, total2 := 0, 0
	go func() {
		for {
			if bitEqual(<-achan, <-bchan) {
				total2++
			}
			pairs++
		}
	}()

	i := 0
	for pairs < pairCompare {
		i++
		a = a * mulA % div
		b = b * mulB % div
		if i < 40000000 && bitEqual(a, b) {
			total1++
		}
		if a%4 == 0 {
			maybeSend(achan, a)
		}
		if b%8 == 0 {
			maybeSend(bchan, b)
		}
	}

	fmt.Println(total1)
	fmt.Println(total2)
}

func generator(row string) int {
	fields := strings.Fields(row)
	i, err := strconv.Atoi(fields[len(fields)-1])
	if err != nil {
		log.Fatalf("Failed to parse %v to int: %v", i, err)
	}
	return i
}

func bitEqual(a, b int) bool {
	return a&0xFFFF == b&0xFFFF
}
