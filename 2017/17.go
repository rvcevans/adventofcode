package main

import (
	"fmt"
	"github.com/rvcevans/adventofcode/getinput"
	"os"
	"strconv"
	"log"
)

func main() {
	input := mustInt(getinput.MustGet(2017, 17, os.Getenv("ADVENT_SESSION"))[0])

	s := []int{0}
	pos := 0
	for i := 1; i <= 2017; i ++ {
		pos = (pos + input + 1) % len(s)
		next := make([]int, len(s)+1)
		for j := 0; j < len(s) + 1; j++ {
			switch {
			case j <= pos:
				next[j] = s[j]
			case j == pos+1:
				next[j] = i
			case j > pos:
				next[j] = s[j-1]
			}
		}

		s = next
	}

	fmt.Println("Part 1:", s[pos+2])

	pos, i, last := 0, 0, 0
	for i <= 50000000 {
		dif := i + 1 - pos
		if dif <= input + 1 {
			pos = (pos + input + 1) % (i + 1)
			i ++
		} else {
			m := dif / (input + 1)
			i += m
			pos += (input + 1) * m
		}
		if pos == 0 {
			last = i
		}
	}

	fmt.Println("Part 2:", last)
}
