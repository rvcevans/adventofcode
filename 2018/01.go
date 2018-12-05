package main

import (
	"github.com/rvcevans/adventofcode/getinput"
	"log"
	"strconv"
)

func main() {
	lines := getinput.MustGet(2018, 1)

	var numbers []int
	endFrequency := 0
	for _, l := range lines {
		v, err := strconv.Atoi(l)
		if err != nil {
			log.Fatalf("failed to convert line to number: %v", err.Error())
		}
		numbers = append(numbers, v)
		endFrequency += v
	}

	log.Printf("Solution 1: %d", endFrequency)

	frequencies := map[int]struct{}{}
	frequency, i, l := numbers[0], 0, len(numbers)
	for seen := false; !seen; _, seen = frequencies[frequency] {
		frequencies[frequency] = struct{}{}

		i++
		if i == l {
			i -= l
		}
		frequency += numbers[i]
	}

	log.Printf("Solution 2: %d", frequency)
}
