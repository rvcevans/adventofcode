package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rvcevans/adventofcode/getinput"
)

func main() {
	inputs := getinput.MustGet(2016, 2, os.Getenv("ADVENT_SESSION"))
	keypads := map[int]*keypad{
		1: newKeypad("123,456,789", '5'),
		2: newKeypad("  1  , 234 ,56789, ABC ,  D  ", '5'),
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

func newKeypad(layout string, initialButton rune) *keypad {
	buttons := make(map[vector]rune)
	location := vector{}
	for j, row := range strings.Split(layout, ",") {
		for i, button := range row {
			if button == ' ' {
				continue
			}
			if button == initialButton {
				location = vector{x: i, y: -j}
			}
			buttons[vector{x: i, y: -j}] = button
		}
	}

	return &keypad{
		buttons:  buttons,
		location: location,
	}
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
