package main

import (
	"fmt"
	"github.com/rvcevans/adventofcode/getinput"
	"os"
	"regexp"
	"strings"
)

func main() {
	start, diagnostic, blueprint := parseBlueprint(getinput.MustGet(2017, 25, os.Getenv("ADVENT_SESSION")))

	m := newMachine(start)
	for diagnostic > 0 {
		m.Next(blueprint)
		diagnostic--
	}

	fmt.Println(m.CheckSum())
}

func newMachine(start rune) machine {
	return machine{state: start, cursor: 0, tape: []bool{false}}
}

type machine struct {
	state  rune
	cursor int
	tape   []bool
}

func (m *machine) Next(b blueprint) {
	rule := b[m.state][m.tape[m.cursor]]
	m.tape[m.cursor] = rule.write
	if rule.right {
		m.cursor++
		if m.cursor == len(m.tape) {
			m.tape = append(m.tape, false)
		}
	} else {
		if m.cursor == 0 {
			m.tape = append([]bool{false}, m.tape...)
		} else {
			m.cursor--
		}
	}
	m.state = rule.state
}

func (m *machine) CheckSum() int {
	sum := 0
	for _, b := range m.tape {
		if b {
			sum++
		}
	}
	return sum
}

type rule struct {
	write, right bool
	state        rune
}

type state map[bool]rule

type blueprint map[rune]state

func parseBlueprint(s []string) (rune, int, blueprint) {
	stateRegex := regexp.MustCompile(".*state ([A-Z]).*")
	getRune := func(s string) rune {
		m := stateRegex.FindStringSubmatch(s)
		return []rune(m[1])[0]
	}
	getBool := func(s string) bool {
		return strings.Contains(s, "1")
	}
	getRight := func(s string) bool {
		return strings.Contains(s, "right")
	}

	getRule := func(s []string) rule {
		return rule{getBool(s[0]), getRight(s[1]), getRune(s[2])}
	}
	getState := func(s []string) state {
		return map[bool]rule{false: getRule(s[1:4]), true: getRule(s[5:8])}
	}

	blueprint := make(map[rune]state)
	for i, v := range s {
		if v == "" {
			blueprint[getRune(s[i+1])] = getState(s[i+2 : i+10])
		}
	}

	return getRune(s[0]), mustInt(regexp.MustCompile(".* ([0-9]*) steps.").FindStringSubmatch(s[1])[1]), blueprint
}
