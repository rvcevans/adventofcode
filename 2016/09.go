package main

import (
	"fmt"
	"github.com/rvcevans/adventofcode/getinput"
	"log"
	"strconv"
)

func main() {
	input := getinput.MustGet(2016, 9)[0]

	fmt.Printf("Solution 1: %v\n", decompress(input, false))
	fmt.Printf("Solution 2: %v\n", decompress(input, true))
}

func decompress(s string, recursive bool) int {
	out := 0
	for i := 0; i < len(s); {
		if s[i] != '(' {
			i++
			out++
			continue
		}
		j := i + 1
		for s[j] != ')' {
			j++
		}
		m, err := newMarker(s[i : j+1])
		if err != nil {
			log.Fatal(err)
		}
		if !recursive {
			out += m.repeat * m.len
		} else {
			out += m.repeat * decompress(s[j+1:j+1+m.len], recursive)
		}
		i = j + m.len + 1
	}
	return out
}

type marker struct {
	repeat, len int
}

func newMarker(s string) (*marker, error) {
	if s[0] != '(' {
		return nil, fmt.Errorf("Missing ( parenthesis: %v", s)
	}
	if s[len(s)-1] != ')' {
		return nil, fmt.Errorf("Missing ) parenthesis: %v", s)
	}
	x := 0
	for i := range s {
		if s[i] == 'x' {
			x = i
		}
	}
	l, err := strconv.Atoi(s[1:x])
	if err != nil {
		return nil, fmt.Errorf("Unable to parse repeat: %v, error: %s", s, err)
	}
	repeat, err := strconv.Atoi(s[x+1 : len(s)-1])
	if err != nil {
		return nil, fmt.Errorf("Unable to parse len: %v, error: %s", s, err)
	}
	return &marker{repeat: repeat, len: l}, nil
}
