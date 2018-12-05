package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fighterlyt/permutation"
	"github.com/rvcevans/adventofcode/getinput"
)

func calcTotal(cities []string, distances map[string]map[string]int) int {
	total := 0
	for i := 0; i < len(cities)-1; i++ {
		cityA := cities[i]
		cityB := cities[i+1]
		total += distances[cityA][cityB]
	}
	return total
}

func main() {
	distances := make(map[string]map[string]int)
	for _, line := range getinput.MustGet(2015, 9) {
		tokens := strings.Split(line, " ")
		cityA := tokens[0]
		cityB := tokens[2]
		distance, _ := strconv.Atoi(tokens[4])
		if _, ok := distances[cityA]; !ok {
			distances[cityA] = make(map[string]int)
		}
		distances[cityA][cityB] = distance
		if _, ok := distances[cityB]; !ok {
			distances[cityB] = make(map[string]int)
		}
		distances[cityB][cityA] = distance
	}

	var cities []string
	for k := range distances {
		cities = append(cities, k)
	}

	p, err := permutation.NewPerm(cities, nil)
	if err != nil {
		panic(err)
	}

	min := calcTotal(cities, distances)
	max := min
	for i, err := p.Next(); err == nil; i, err = p.Next() {
		citiesPerm := i.([]string)
		dist := calcTotal(citiesPerm, distances)
		if dist < min {
			min = dist
		} else if dist > max {
			max = dist
		}

	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
