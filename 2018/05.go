package main

import (
	"github.com/rvcevans/adventofcode/getinput"
	"unicode"
	"fmt"
	"strings"
)

func main() {
	polymer := getinput.MustGet(2018, 5)[0]
	fmt.Printf("Solution 1: %d\n", reactSize(polymer))

	smallest := len(polymer)
	for r := 'a'; unicode.IsLower(r); r ++ {
		if size := reactSize(stripLetter(polymer, r)); size < smallest {
			smallest = size
		}
	}
	fmt.Printf("Solution 2: %d\n", smallest)
}

func reactSize(polymer string) int {
	return len(reactPolymer([]rune(polymer)))
}

func reactPolymer(polymer []rune) []rune {
	i := 0
	reacted := false
	var next []rune
	for i < len(polymer) {
		if i == len(polymer) - 1 || !react(polymer[i], polymer[i+1]) {
			next = append(next, polymer[i])
			i ++
		} else {
			reacted = true
			i += 2
		}
	}
	if !reacted {
		return polymer
	}
	return reactPolymer(next)
}

func react(a, b rune) bool {
	return a != b && unicode.ToLower(a) == unicode.ToLower(b)
}

func stripLetter(s string, lower rune) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case lower, unicode.ToUpper(lower):
			return -1
		default:
			return r
		}
	}, s)
}