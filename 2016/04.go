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
	input := getinput.MustGet(2016, 4)

	sectorSum := 0
	var sector2 int
	for _, i := range input {
		r := newRoom(i)
		if r.valid() {
			sectorSum += r.sector
		}

		if strings.Contains(r.decrypt(), "object") {
			sector2 = r.sector
		}
	}

	fmt.Printf("Solution 1: %v\n", sectorSum)
	fmt.Printf("Solution 2: %v\n", sector2)
}

type room struct {
	name     []string
	letters  map[rune]int
	sector   int
	checksum map[rune]struct{}
}

func newRoom(code string) *room {
	s := strings.SplitN(code, "[", 2)
	words := strings.Split(s[0], "-")
	name, sectorString := words[:len(words)-1], words[len(words)-1]

	sector, err := strconv.Atoi(sectorString)
	if err != nil {
		log.Fatalf("Unable to convert sector to integer: %v", err)
	}

	letters := make(map[rune]int)
	for _, w := range name {
		for _, r := range w {
			letters[r]++
		}
	}

	checksum := make(map[rune]struct{})
	for _, r := range strings.TrimSuffix(s[1], "]") {
		checksum[r] = struct{}{}
	}

	return &room{
		name:     name,
		letters:  letters,
		checksum: checksum,
		sector:   sector,
	}
}

func (r *room) valid() bool {
	if len(r.checksum) < 5 {
		return false
	}
	for l, count := range r.letters {
		// if the letter is not in the checksum we want to see if there is any letter in the checksum that
		// occurs less or equal times and alphabetically after
		if _, ok := r.checksum[l]; !ok {
			for c := range r.checksum {
				v, ok := r.letters[c]
				if !ok {
					return false
				}
				if v <= count && c > l {
					return false
				}
			}
		}
	}
	return true
}

func (r *room) decrypt() string {
	out := make([]string, 0)
	for _, w := range r.name {
		n := ""
		for _, l := range w {
			n += string(rotate(l, r.sector))
		}
		out = append(out, n)
	}
	return strings.Join(out, " ")
}

func rotate(r rune, n int) rune {
	offset := (int(r-'a') + n) % 26
	if offset < 0 {
		offset += 26
	}
	return rune('a' + offset)
}
