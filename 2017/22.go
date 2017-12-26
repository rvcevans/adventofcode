package main

import (
	"fmt"
	"github.com/rvcevans/adventofcode/getinput"
	"os"
)

const (
	clean    = 0
	weakened = 1
	infected = 2
	flagged  = 3
)

func main() {
	input := getinput.MustGet(2017, 22, os.Getenv("ADVENT_SESSION"))
	//input = []string{"..#","#..","..."}

	l := len(input)
	grid := make(grid)
	for y, row := range input {
		for x, r := range row {
			n := vector{x, y}
			if r == '#' {
				grid[n] = infected
			} else {
				grid[n] = clean
			}
		}
	}

	v := virus{
		grid:      grid,
		position:  vector{l / 2, l / 2},
		direction: vector{0, -1},
	}

	for i := 0; i < 10000000; i++ {
		v.Next()
	}
	fmt.Println(v.infected)
}

func next(c int) int {
	return map[int]int{
		clean:    weakened,
		weakened: infected,
		infected: flagged,
		flagged:  clean,
	}[c]
}

type vector struct {
	x, y int
}

func (v vector) Left() vector  { return vector{v.y, -v.x} }
func (v vector) Right() vector { return vector{-v.y, v.x} }
func (a vector) Add(b vector) vector {
	return vector{a.x + b.x, a.y + b.y}
}

type grid map[vector]int

//func (g grid) String() string {
//	minX, maxX, minY, maxY := 0, 0, 0, 0
//	for v := range g {
//		minX = minInt(minX, v.x)
//		minY = minInt(minY, v.y)
//		maxX = maxInt(maxX, v.x)
//		maxY = maxInt(maxY, v.y)
//	}
//
//	var s string
//	for y := minY; y <= maxY; y++ {
//		for x := minX; x <= maxX; x++ {
//			if g[vector{x, y}] {
//				s += "#"
//			} else {
//				s += "."
//			}
//		}
//		s += "\n"
//	}
//	return s
//}

type virus struct {
	grid
	position, direction vector
	infected            int
}

func (v *virus) Next() {
	switch v.grid[v.position] {
	case clean:
		v.direction = v.direction.Left()
	case weakened:
	case infected:
		v.direction = v.direction.Right()
	case flagged:
		v.direction = v.direction.Right().Right()
	}

	v.grid[v.position] = next(v.grid[v.position])
	if v.grid[v.position] == infected {
		v.infected++
	}
	v.position = v.position.Add(v.direction)
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}
