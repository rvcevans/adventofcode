package main

import (
	"github.com/rvcevans/adventofcode/getinput"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := strconv.Atoi(getinput.MustGet(2017, 3, os.Getenv("ADVENT_SESSION"))[0])
	if err != nil {
		log.Fatalf("Failed to convert input to integer: %v", err.Error())
	}

	location := vector{0, 0}
	g := grid{location: 1}
	v := 0
	for v < input {
		location = location.Next()
		v = g.GetValue(location)
	}

	log.Println(v)
}

type vector struct {
	x, y int
}

func (v vector) Left() vector  { return vector{v.x - 1, v.y} }
func (v vector) Right() vector { return vector{v.x + 1, v.y} }
func (v vector) Up() vector    { return vector{v.x, v.y + 1} }
func (v vector) Down() vector  { return vector{v.x, v.y - 1} }

func (v vector) Next() vector {
	switch {
	case v.y <= -abs(v.x):
		return v.Right()
	case v.x <= -abs(v.y):
		return v.Down()
	case v.y >= abs(v.x):
		return v.Left()
	default:
		return v.Up()
	}
}

type grid map[vector]int

func (g grid) GetValue(v vector) int {
	total := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i != 0 || j != 0 {
				total += g[vector{v.x + i, v.y + j}]
			}
		}
	}
	g[v] = total
	return total
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}