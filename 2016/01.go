package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/rvcevans/adventofcode/getinput"
)

func main() {
	input := getinput.MustGet(2016, 1, os.Getenv("ADVENT_SESSION"))[0]
	moves := strings.Split(input, ", ")

	t := newTaxi()
	visited := make(map[vector]struct{})
	var secondSolution vector
	secondFound := false
	for _, m := range moves {
		direction := string(m[0])
		blocks, err := strconv.Atoi(string(m[1:]))
		if err != nil {
			log.Fatalf("Failed to parse distance: %v", err)
		}

		t.turn(direction)
		for i := 0; i < blocks; i++ {
			t.move(1)
			if _, ok := visited[t.Location]; ok && !secondFound {
				secondFound = true
				secondSolution = t.Location
			}
			visited[t.Location] = struct{}{}
		}

	}

	fmt.Printf("1st solution: %v blocks away\n", t.Location.away())
	fmt.Printf("2nd solution: %v blocks away\n", secondSolution.away())
}

type vector struct {
	x int
	y int
}

type taxi struct {
	Location  vector
	direction vector
}

func newTaxi() *taxi {
	return &taxi{direction: vector{0, 1}}
}

func (t *taxi) move(blocks int) {
	t.Location.x += t.direction.x * blocks
	t.Location.y += t.direction.y * blocks
}

func (t *taxi) turn(direction string) {
	if direction == "R" {
		t.right()
	} else {
		t.left()
	}
}

func (t *taxi) right() {
	t.direction.x, t.direction.y = t.direction.y, -t.direction.x
}

func (t *taxi) left() {
	t.direction.x, t.direction.y = -t.direction.y, t.direction.x
}

func (l *vector) away() int {
	return abs(l.x) + abs(l.y)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
