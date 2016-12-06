package main

import (
	"fmt"
	"os"

	"github.com/rvcevans/adventofcode/getinput"
)

func main() {
	inputs := getinput.MustGet(2016, 2, os.Getenv("ADVENT_SESSION"))

	keypads := map[int]*keypad{
		1: {
			buttons: map[vector]rune{
				vector{-1, 1}:  '1',
				vector{-1, 0}:  '4',
				vector{-1, -1}: '7',
				vector{0, 1}:   '2',
				vector{0, 0}:   '5',
				vector{0, -1}:  '8',
				vector{1, 1}:   '3',
				vector{1, 0}:   '6',
				vector{1, -1}:  '9',
			},
		},
		2: {
			buttons: map[vector]rune{
				vector{0, 0}:  '5',
				vector{1, 1}:  '2',
				vector{1, 0}:  '6',
				vector{1, -1}: 'A',
				vector{2, 2}:  '1',
				vector{2, 1}:  '3',
				vector{2, 0}:  '7',
				vector{2, -1}: 'B',
				vector{2, -2}: 'D',
				vector{3, 1}:  '4',
				vector{3, 0}:  '8',
				vector{3, -1}: 'C',
				vector{4, 0}:  '9',
			},
		},
	}

	for i, k := range keypads {
		solution := ""
		for _, i := range inputs {
			for _, m := range i {
				k.move(m)
			}
			solution += string(k.value())
		}
		fmt.Printf("Solution %v: %v\n", i, solution)
	}
}

type vector struct {
	x int
	y int
}

type keypad struct {
	buttons  map[vector]rune
	location vector
}

func (k *keypad) value() rune {
	return k.buttons[k.location]
}

func (k *keypad) move(r rune) {
	next := map[rune]func(vector) vector{
		'U': up,
		'R': right,
		'D': down,
		'L': left,
	}[r](k.location)
	if _, ok := k.buttons[next]; ok {
		k.location = next
	}
}

func up(v vector) vector {
	return vector{x: v.x, y: v.y + 1}
}

func down(v vector) vector {
	return vector{x: v.x, y: v.y - 1}
}

func right(v vector) vector {
	return vector{x: v.x + 1, y: v.y}
}

func left(v vector) vector {
	return vector{x: v.x - 1, y: v.y}
}
