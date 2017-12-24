package main

import (
	"fmt"
	"github.com/rvcevans/adventofcode/getinput"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := getinput.MustGet(2017, 18, os.Getenv("ADVENT_SESSION"))

	var instructions [][]string
	for _, i := range input {
		instructions = append(instructions, strings.Split(i, " "))
	}

	c01, c10 := make(chan int, 1000), make(chan int, 1000)
	assemblies := map[int]*assembly{0: newAssembly(0, instructions, c10, c01), 1: newAssembly(1, instructions, c01, c10)}
	for {
		progress := false
		for _, a := range assemblies {
			progress = progress || a.Continue()
		}
		if !progress { // No programs have progressed
			fmt.Println(assemblies[1].sent)
			return
		}
	}
}

type assembly struct {
	instructions [][]string
	sent         int
	position     int
	in, out      chan int
	registry     map[string]int
}

func newAssembly(p int, instructions [][]string, in, out chan int) *assembly {
	return &assembly{
		instructions: instructions,
		in:           in,
		out:          out,
		registry:     map[string]int{"p": p},
	}
}

// Returns false if no progress is made
func (a *assembly) Continue() bool {
	progressed := false
	for {
		i := a.instructions[a.position]
		cmd := i[0]
		switch cmd {
		case "snd":
			a.out <- a.get(i[1])
			a.sent++
		case "set", "add", "mul", "mod":
			a.cmd(cmd)(i[1], i[2])
		case "rcv":
			select {
			case a.registry[i[1]] = <-a.in:
			default:
				return progressed
			}
		case "jgz":
			if a.get(i[1]) > 0 {
				a.position += a.get(i[2]) - 1
			}
		}
		progressed = true
		a.position++
	}

}

func (a *assembly) cmd(s string) func(x, y string) {
	return map[string]func(a, y string){
		"set": a.set,
		"add": a.add,
		"mul": a.mul,
		"mod": a.mod,
	}[s]
}
func (a *assembly) set(x, y string) { a.registry[x] = a.get(y) }
func (a *assembly) add(x, y string) { a.registry[x] += a.get(y) }
func (a *assembly) mul(x, y string) { a.registry[x] *= a.get(y) }
func (a *assembly) mod(x, y string) {
	m := a.get(y)
	a.registry[x] %= m
	if a.registry[x] < 0 {
		a.registry[x] += m
	}
}

func (a *assembly) get(y string) int {
	i, err := strconv.Atoi(y)
	if err != nil {
		return a.registry[y]
	}
	return i
}
