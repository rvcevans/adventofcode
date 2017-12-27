package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type registry map[string]int

func (r registry) String() string {
	var s, keys []string
	for k := range r {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		s = append(s, fmt.Sprintf("%v:%v", k, r[k]))
	}
	return "{" + strings.Join(s, ", ") + "}"
}

type assembly struct {
	instructions         [][]string
	sent, received       int
	position, iterations int
	in, out              chan int
	registry             registry
	cmdCount             map[string]int
}

func newAssembly(registry map[string]int, instructions [][]string, in, out chan int) *assembly {
	return &assembly{
		instructions: instructions,
		in:           in,
		out:          out,
		registry:     registry,
		cmdCount:     make(map[string]int),
	}
}

// Returns false if no progress is made
func (a *assembly) Run() bool {
	progressed := false
	for {
		if a.position < 0 || a.position >= len(a.instructions) {
			return progressed
		}
		a.iterations++

		i := a.instructions[a.position]
		cmd := i[0]
		a.cmdCount[cmd]++
		switch cmd {
		case "snd":
			a.out <- a.get(i[1])
			a.sent++
		case "set", "add", "sub", "mul", "mod":
			a.cmd(cmd)(i[1], i[2])
		case "rcv":
			select {
			case a.registry[i[1]] = <-a.in:
				a.received++
			default:
				return progressed
			}
		case "jgz":
			if a.get(i[1]) > 0 {
				a.position += a.get(i[2]) - 1
			}
		case "jnz":
			if a.get(i[1]) != 0 {
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
		"sub": a.sub,
		"mul": a.mul,
		"mod": a.mod,
	}[s]
}
func (a *assembly) set(x, y string) { a.registry[x] = a.get(y) }
func (a *assembly) add(x, y string) { a.registry[x] += a.get(y) }
func (a *assembly) sub(x, y string) { a.registry[x] -= a.get(y) }
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
