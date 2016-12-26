package main

import (
	"fmt"
	"strconv"
)

func main() {
	j := newJourney(1352, vector{1, 1})
	end := vector{31, 39}

	for !j.Been(end) {
		j.Next()
		if j.steps == 50 {
			fmt.Printf("Solution 2: %v\n", len(j.visited))
		}
	}

	fmt.Printf("Solution 1: %v\n", j.steps)
}

type journey struct {
	maze    *maze
	visited map[vector]struct{}
	steps   int
}

func newJourney(favorite int, start vector) *journey {
	return &journey{
		maze: newMaze(favorite),
		visited: map[vector]struct{}{
			start: struct{}{},
		},
		steps: 0,
	}
}

func (j *journey) Next() {
	visited := make(map[vector]struct{})
	for v := range j.visited {
		visited[v] = struct{}{}
		for n := range j.maze.Neighbours(v) {
			visited[n] = struct{}{}
		}
	}
	j.visited = visited
	j.steps++
}

func (j *journey) Been(v vector) bool {
	_, ok := j.visited[v]
	return ok
}

type vector struct {
	x, y int
}

func neighbours(v vector) map[vector]struct{} {
	return map[vector]struct{}{
		vector{x: v.x - 1, y: v.y}: struct{}{},
		vector{x: v.x + 1, y: v.y}: struct{}{},
		vector{x: v.x, y: v.y - 1}: struct{}{},
		vector{x: v.x, y: v.y + 1}: struct{}{},
	}
}

type maze struct {
	favorite int
	space    [][]bool
	size     *vector
}

func newMaze(favorite int) *maze {
	space := make([][]bool, 0)
	return &maze{
		favorite: favorite,
		space:    space,
		size:     &vector{x: -1, y: -1},
	}
}

// returns any neighbouring spaces to the location
func (m *maze) Neighbours(v vector) map[vector]struct{} {
	if !m.in(v) {
		return nil
	}

	out := make(map[vector]struct{})
	for n := range neighbours(v) {
		if m.in(n) {
			out[n] = struct{}{}
		}
	}
	return out
}

func (m *maze) in(v vector) bool {
	if v.x < 0 || v.y < 0 {
		return false
	}
	for v.x > m.size.x {
		m.addColumn()
	}
	for v.y > m.size.y {
		m.addRow()
	}
	return m.space[v.y][v.x]
}

func (m *maze) addColumn() {
	m.size.x++
	for y := 0; y <= m.size.y; y++ {
		m.space[y] = append(m.space[y], m.open(vector{m.size.x, y}))
	}
}

func (m *maze) addRow() {
	m.size.y++
	row := make([]bool, m.size.x+1)
	for x := 0; x <= m.size.x; x++ {
		row[x] = m.open(vector{x, m.size.y})
	}
	m.space = append(m.space, row)
}

func (m *maze) open(v vector) bool {
	return binarySum(v.x*v.x+3*v.x+2*v.x*v.y+v.y+v.y*v.y+m.favorite)%2 == 0
}

func binarySum(i int) int {
	bin := strconv.FormatInt(int64(i), 2)
	sum := 0
	for _, b := range bin {
		if b == '1' {
			sum++
		}
	}
	return sum
}
