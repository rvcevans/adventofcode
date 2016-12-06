package main

import (
	"fmt"
	"os"

	"github.com/rvcevans/adventofcode/getinput"
)

func main() {
	input := getinput.MustGet(2016, 6, os.Getenv("ADVENT_SESSION"))


	columns := make([]map[rune]int, 0)
	for j, row := range input {
		if j == 0 {
			for range row {
				columns = append(columns, make(map[rune]int))
			}
		}
		for i, r := range row {
			columns[i][r] ++
		}
	}

	solution1, solution2 := "", ""
	for _, counts := range columns {
		min, max := len(input), 0
		var letter1, letter2 rune
		for r, count := range counts {
			if count > max {
				max = count
				letter1 = r
			}
			if count < min {
				min = count
				letter2 = r
			}
		}
		solution1 += string(letter1)
		solution2 += string(letter2)
	}

	fmt.Printf("Solution 1: %v\n", solution1)
	fmt.Printf("Solution 2: %v\n", solution2)
}
