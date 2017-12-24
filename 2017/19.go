package main

import (
	"fmt"
	"github.com/rvcevans/adventofcode/getinput"
	"os"
)

func main() {
	input := getinput.MustGet(2017, 19, os.Getenv("ADVENT_SESSION"))

	var start int
	for i, r := range input[0] {
		if r == '|' {
			start = i
		}
	}

	d := &diagram{current: vector{start, 0}, direction: vector{0, 1}, lines: input}
	steps := 1
	for d.Next() {
		steps++
	}
	fmt.Println(d.seen)
	fmt.Println(steps)
}

type vector struct {
	x, y int
}

func (v vector) Right() vector {
	return vector{v.y, -v.x}
}
func (v vector) Left() vector {
	return vector{-v.y, v.x}
}
func (v vector) Value() rune {
	if v.x != 0 {
		return '-'
	}
	return '|'
}
func (a vector) Add(b vector) vector {
	return vector{a.x + b.x, a.y + b.y}
}

type diagram struct {
	current, direction vector
	lines              []string
	seen               string
}

func (d *diagram) Next() bool {
	d.current = d.current.Add(d.direction)
	v := d.Value(d.current)
	switch v {
	case '+':
		if right := d.direction.Right(); d.Value(d.current.Add(right)) == right.Value() {
			d.direction = right
		} else {
			d.direction = d.direction.Left()
		}
	case '-', '|':
	case ' ':
		return false
	default:
		d.seen += string(v)
	}

	return true
}

func (d *diagram) Value(v vector) rune {
	if v.x < 0 || v.y < 0 || v.x >= len(d.lines[0]) || v.y >= len(d.lines) {
		return ' '
	}
	return rune(d.lines[v.y][v.x])
}
