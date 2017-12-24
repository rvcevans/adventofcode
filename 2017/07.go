package main

import (
	"github.com/rvcevans/adventofcode/getinput"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"fmt"
)

var (
	regexp7single = regexp.MustCompile(`([a-z]*) \(([0-9]*)\)`)
	regexp7multi  = regexp.MustCompile(`([a-z]*) \(([0-9]*)\) -> (.*)`)
)

func main() {
	input := getinput.MustGet(2017, 7, os.Getenv("ADVENT_SESSION"))

	weights := make(map[string]int)
	above := make(map[string][]string)
	below := make(map[string]string)
	towers := make(map[string]int)
	for _, row := range input {
		node, weight, upperNodes := parseRow(row)
		weights[node] = weight
		above[node] = upperNodes
		if len(upperNodes) == 0 {
			towers[node] = weight
		}
		for _, up := range upperNodes {
			below[up] = node
		}
	}

	done := false
	for !done {
		loop:
		for n, ans := range above {
			total := weights[n]
			for _, an := range ans {
				v, ok := towers[an]
				if !ok {
					continue loop
				}
				total += v
			}
			towers[n] = total
		}
		done = true
		for n := range weights {
			if _, ok := towers[n]; !ok {
				done = false
			}
		}
	}

	balances := make(map[int]int)
	for _, t := range towers {
		balances[t] ++
	}
	minIncorrect := -1
	for b, count := range balances {
		if count == 1 && (minIncorrect == -1 || b < minIncorrect) {
			minIncorrect = b
		}
	}
	for n, t := range towers {
		if t == minIncorrect {
			for _, neighbour := range above[below[n]] {
				if neighbour != n {
					fmt.Println(weights[n]+ towers[neighbour]- towers[n])
					return
				}
			}
		}
	}
}

func parseRow(s string) (string, int, []string) {
	m := regexp7multi.FindStringSubmatch(s)
	var upperNodes []string
	if len(m) == 0 {
		m = regexp7single.FindStringSubmatch(s)
	} else {
		upperNodes = strings.Split(m[3], ", ")
	}

	weight, err := strconv.Atoi(m[2])
	if err != nil {
		log.Fatalf("failed to convert weight to number: %v", err.Error())
	}

	return m[1], weight, upperNodes
}
