package main

import (
	"fmt"
	"github.com/rvcevans/adventofcode/getinput"
	"strings"
)

func main() {
	input := getinput.MustGet(2017, 18)

	var instructions [][]string
	for _, i := range input {
		instructions = append(instructions, strings.Split(i, " "))
	}

	c01, c10 := make(chan int, 1000), make(chan int, 1000)
	assemblies := map[int]*assembly{
		0: newAssembly(map[string]int{"p": 0}, instructions, c10, c01),
		1: newAssembly(map[string]int{"p": 1}, instructions, c01, c10),
	}
	for {
		progress := false
		for _, a := range assemblies {
			progress = progress || a.Run()
		}
		if !progress { // No programs have progressed
			fmt.Println(assemblies[1].sent)
			return
		}
	}
}
