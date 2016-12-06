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
	input := getinput.MustGet(2016, 3, os.Getenv("ADVENT_SESSION"))

	rows := make([][]int, 0, len(input))
	for _, i := range input {
		fields := strings.Fields(i)
		row := make([]int, 0)
		for _, s := range fields {
			i, err := strconv.Atoi(strings.TrimSpace(s))
			if err != nil {
				log.Fatalf("Failed to parse side as integer: %v", err)
			}
			row = append(row, i)
		}
		rows = append(rows, row)
	}

	valid1, valid2 := 0, 0
	for i, row := range rows {
		c2 := i % 3
		r2 := i - c2

		t1 := triangle{a: row[0], b: row[1], c: row[2]}
		t2 := triangle{a: rows[r2][c2], b: rows[r2 + 1][c2], c: rows[r2 + 2][c2]}
		if t1.valid() {
			valid1++
		}
		if t2.valid() {
			valid2++
		}
	}

	fmt.Printf("Solution 1: %v\n", valid1)
	fmt.Printf("Solution 2: %v\n", valid2)
}

type triangle struct {
	a int
	b int
	c int
}

func (t *triangle) valid() bool {
	if t.a >= t.b+t.c || t.b >= t.a+t.c || t.c >= t.a+t.b {
		return false
	}
	return true
}
