package main

import (
	"github.com/rvcevans/adventofcode/getinput"
	"os"
	"fmt"21
	"log"
	"strings"
)

func main() {
	book := map[int]map[int]square{2: make(map[int]square), 3: make(map[int]square)}

	for _, rule := range getinput.MustGet(2017, 21, os.Getenv("ADVENT_SESSION")) {
		s := strings.Split(rule, " => ")
		from, to := newSquare(s[0]), newSquare(s[1])

		for j := 0; j < 2; j++ {
			for i := 0; i < 4; i++ {
				book[from.Len()][from.Value()] = to
				from.Right()
			}
			from.Flip()
		}
	}

	on := func(iterations int) int {
		s := newSquare(".#./..#/###")
		for i := 0; i < iterations; i++ {
			var mapped []square
			for _, n := range s.Split() {
				mapped = append(mapped, book[n.Len()][n.Value()])
			}
			s = joinSquares(mapped)
		}
		return s.Sum()
	}

	fmt.Println("Part1:", on(5))
	fmt.Println("Part2:", on(18))
}

func newSquare(s string) square {
	var square [][]bool
	for _, l := range strings.Split(s, "/") {
		var row []bool
		for _, r := range l {
			if r == '.' {
				row = append(row, false)
			} else {
				row = append(row, true)
			}
		}
		square = append(square, row)
	}
	return square
}

type square [][]bool

func (s square) Len() int {
	return len(s)
}

func (s square) Value() int {
	v, i := 0, uint(0)
	for _, l := range s {
		for _, b := range l {
			if b {
				v += 1 << i
			}
			i++
		}
	}
	return v
}

func (s square) Sum() int {
	total := 0
	for _, l := range s {
		for _, r := range l {
			if r {
				total++
			}
		}
	}
	return total
}

func (s square) Right() {
	switch len(s) {
	case 2:
		s[0][0], s[0][1], s[1][1], s[1][0] = s[1][0], s[0][0], s[0][1], s[1][1]
	case 3:
		s[0][0], s[0][2], s[2][2], s[2][0] = s[2][0], s[0][0], s[0][2], s[2][2]
		s[0][1], s[1][2], s[2][1], s[1][0] = s[1][0], s[0][1], s[1][2], s[2][1]
	default:
		log.Fatalf("Unsupported square size %v", len(s))
	}
}

func (s square) Flip() {
	switch len(s) {
	case 2:
		s[0][0], s[0][1], s[1][0], s[1][1] = s[1][0], s[1][1], s[0][0], s[0][1]
	case 3:
		s[0][0], s[0][1], s[0][2], s[2][0], s[2][1], s[2][2] = s[2][0], s[2][1], s[2][2], s[0][0], s[0][1], s[0][2]
	default:
		log.Fatalf("Unsupported square size %v", len(s))
	}
}

func (s square) Split() []square {
	l := len(s)
	size := 2
	if l%size != 0 {
		size = 3
	}

	var squares []square
	for i := 0; i < l; i += size {
		for j := 0; j < l; j += size {
			var ss square
			for x := 0; x < size; x++ {
				var row []bool
				for y := 0; y < size; y++ {
					row = append(row, s[i+x][j+y])
				}
				ss = append(ss, row)
			}
			squares = append(squares, ss)
		}
	}
	return squares
}

func joinSquares(squares []square) square {
	size := len(squares[0])
	count := 1
	for count*count < len(squares) {
		count++
	}
	l := size * count

	s := make([][]bool, l)
	for i := 0; i < l; i++ {
		s[i] = make([]bool, l)
	}

	for i := 0; i < count; i += 1 {
		for j := 0; j < count; j += 1 {
			for x := 0; x < size; x++ {
				for y := 0; y < size; y++ {
					s[i*size+x][j*size+y] = squares[i*count+j][x][y]
				}
			}
		}
	}
	return s
}
