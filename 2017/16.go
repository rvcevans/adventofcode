package main

import (
	"fmt"
	"github.com/rvcevans/adventofcode/getinput"
	"log"
	"os"
	"strconv"
	"strings"
)

const totalDances = 1000000000

func main() {
	input := getinput.MustGet(2017, 16, os.Getenv("ADVENT_SESSION"))[0]

	start := make([]rune, 16)
	for i, r := 0, 'a'; i < 16; i, r = i+1, r+1 {
		start[i] = r
	}

	var moves []move
	for _, move := range strings.Split(input, ",") {
		switch move[0] {
		case 's':
			moves = append(moves, spin(mustInt(move[1:])))
		case 'x':
			s := strings.Split(move[1:], "/")
			moves = append(moves, exchange(mustInt(s[0]), mustInt(s[1])))
		case 'p':
			moves = append(moves, partner(rune(move[1]), rune(move[3])))
		default:
			log.Fatalf("Unrecognized dance move '%v'", move[0])
		}
	}

	dance := func(s []rune, repeat int) []rune {
		p := make([]rune, 16)
		copy(p, s)
		for i := 0; i < repeat; i ++ {
			for _, m := range moves {
				p = m(p)
			}
		}
		return p
	}

	cycle := 0
	current := start
	for cycle == 0 || !sliceEqual(current, start) {
		current = dance(current, 1)
		cycle ++
	}

	fmt.Println(str(dance(start, 1)))
	fmt.Println(str(dance(start, totalDances % cycle)))
}

type move = func([]rune) []rune

func spin(size int) move {
	return func(s []rune) []rune {
		return append(s[len(s)-size:], s[:len(s)-size]...)
	}
}

func exchange(i, j int) move {
	return func(s []rune) []rune {
		s[i], s[j] = s[j], s[i]
		return s
	}
}

func partner(a, b rune) move {
	return func(s []rune) []rune {
		return exchange(index(s, a), index(s, b))(s)
	}
}

func index(s []rune, r rune) int {
	for i, p := range s {
		if p == r {
			return i
		}
	}
	log.Fatalf("Failed to index %s in slice %v", r, s)
	return -1
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Failed to convert %v to int: %v", s, err)
	}
	return i
}

func str(s []rune) string {
	str := ""
	for _, r := range s {
		str += string(r)
	}
	return str
}

func sliceEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}