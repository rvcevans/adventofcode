package main

import (
	"fmt"
	"regexp"
	"strconv"
	"github.com/rvcevans/adventofcode/getinput"
)

func main() {
	input := getinput.MustGet(2016, 8)

	s := newScreen()
	for _, i := range input {
		rect := newRect(i)
		if rect != nil {
			s.Rect(*rect)
			continue
		}
		rotate := newRotate(i)
		s.Rotate(*rotate)
	}

	fmt.Printf("Solution 1: %v\n", s.Voltage())
	s.Print()
}

func newScreen() *screen {
	pixels := make([][]bool, 6)
	for r := range pixels {
		pixels[r] = make([]bool, 50)
	}
	return &screen{pixels}
}

type screen struct {
	pixels [][]bool
}

func (s *screen) Print() {
	for _, row := range s.pixels {
		line := ""
		for _, v := range row {
			if v {
				line += "#"
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}
}

func (s *screen) Voltage() int {
	count := 0
	for _, row := range s.pixels {
		for _, v := range row {
			if v {
				count ++
			}
		}
	}
	return count
}

func (s *screen) Rect(r rect) {
	for x := 0; x < r.x; x++ {
		for y := 0; y < r.y; y++ {
			s.pixels[y][x] = true
		}
	}
}

func (s *screen) Rotate(r rotate) {
	if r.row {
		s.setRow(r.index, rotateArray(s.getRow(r.index), r.offset))
	} else {
		s.setCol(r.index, rotateArray(s.getCol(r.index), r.offset))
	}
}

func (s *screen) getRow(i int) []bool {
	return s.pixels[i]
}

func (s *screen) setRow(i int, row []bool) {
	s.pixels[i] = row
}

func (s *screen) getCol(i int) []bool {
	c := make([]bool, 0)
	for r := range s.pixels {
		c = append(c, s.pixels[r][i])
	}
	return c
}

func (s *screen) setCol(i int, col []bool) {
	for r := range s.pixels {
		s.pixels[r][i] = col[r]
	}
}

func rotateArray(array []bool, offset int) []bool {
	l := len(array)
	offsetMod := offset % l
	if offsetMod < 0 {
		offsetMod += l
	}
	return append(array[l- offsetMod:], array[:l-offsetMod]...)
}

type rect struct {
	x, y int
}

type rotate struct {
	row           bool
	index, offset int
}

func newRect(s string) *rect {
	r := regexp.MustCompile(`rect ([0-9]{1,2})x([0-9]{1,2})`)
	m := r.FindStringSubmatch(s)
	if len(m) != 3 {
		return nil
	}
	x, err := strconv.Atoi(m[1])
	if err != nil {
		return nil
	}
	y, err := strconv.Atoi(m[2])
	if err != nil {
		return nil
	}
	return &rect{x, y}
}

func newRotate(s string) *rotate {
	r := regexp.MustCompile(`rotate (.*) \w{1}=([0-9]{1,2}) by ([0-9]{1,2})`)
	m := r.FindStringSubmatch(s)
	if len(m) != 4 {
		return nil
	}
	index, err := strconv.Atoi(m[2])
	if err != nil {
		return nil
	}
	offset, err := strconv.Atoi(m[3])
	if err != nil {
		return nil
	}
	row := m[1] == "row"
	return &rotate{row, index, offset}
}
