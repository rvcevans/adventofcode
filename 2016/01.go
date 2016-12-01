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
	visited := make(map[location]struct{})
	var secondSolution location
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

type location struct {
	x int
	y int
}

type taxi struct {
	Location   location
	xDirection int
	yDirection int
}

func newTaxi() *taxi {
	return &taxi{xDirection: 0, yDirection: 1, Location: location{}}
}

func (t *taxi) move(blocks int) {
	t.Location.x += t.xDirection * blocks
	t.Location.y += t.yDirection * blocks
}

func (t *taxi) turn(direction string) {
	if direction == "R" {
		t.right()
	} else {
		t.left()
	}
}

func (t *taxi) right() {
	t.xDirection, t.yDirection = t.yDirection, -t.xDirection
}

func (t *taxi) left() {
	t.xDirection, t.yDirection = -t.yDirection, t.xDirection
}

func (l *location) away() int {
	return abs(l.x) + abs(l.y)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
