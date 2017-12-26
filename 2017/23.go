package main

import (
	"fmt"
	"github.com/rvcevans/adventofcode/getinput"
	"math"
	"os"
	"strings"
)

func main() {
	input := getinput.MustGet(2017, 23, os.Getenv("ADVENT_SESSION"))

	var instructions [][]string
	for _, i := range input {
		instructions = append(instructions, strings.Split(i, " "))
	}

	a := newAssembly(make(map[string]int), instructions, nil, nil, 0)
	for a.Run() {
	}

	fmt.Println(a.cmdCount["mul"])

	h := 0
	for b := 107900; b <= 124900; b += 17 {
		if !isPrime(b) {
			h++
		}
	}

	fmt.Println(h)
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
