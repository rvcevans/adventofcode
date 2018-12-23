package main

import (
	"github.com/rvcevans/adventofcode/getinput"
	"log"
	"strconv"
	"strings"
)

const maxTotal = 10000

func main() {
	lines := getinput.MustGet(2018, 6)

	locations := map[vector]struct{}{}
	for _, l := range lines {
		var is []int
		for _, s := range strings.Split(l, ", ") {
			v, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("failed to convert line to number: %v", err.Error())
			}
			is = append(is, v)
		}
		locations[vector{is[0], is[1]}] = struct{}{}
	}

	var max vector
	for l := range locations {
		if max.x == 0 {
			max = l
		} else {
			if l.x > max.x {
				max.x = l.x
			}
			if l.y > max.y {
				max.y = l.y
			}
		}
	}

	area := map[vector]int{}
	border := map[vector]struct{}{}
	regionSize := 0
	for x := 0; x <= max.x; x++ {
		for y := 0; y <= max.y; y++ {
			v := vector{x, y}
			var closest vector
			min := -1
			totalDistance := 0
			for l := range locations {
				d := v.Distance(l)
				totalDistance += d
				if min == -1 || d < min {
					min = d
					closest = l
				} else if d == min {
					closest = vector{}
				}
			}

			if totalDistance < maxTotal {
				regionSize++
			}

			if closest.x == 0 {
				continue
			}

			area[closest]++
			// Check if we are on the border
			if x == 0 || y == 0 || x == max.x || y == max.y {
				border[closest] = struct{}{}
			}
		}
	}

	var maxArea int
	for l := range locations {
		if _, onBorder := border[l]; onBorder {
			continue
		}

		if a := area[l]; a > maxArea {
			maxArea = a
		}
	}

	log.Printf("Solution 1: %d", maxArea)
	log.Printf("Solution 2: %d", regionSize)
}

type vector struct {
	x, y int
}

func (v vector) Distance(u vector) int {
	return abs(v.x-u.x) + abs(v.y-u.y)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
