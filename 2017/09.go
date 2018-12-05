package main

import (
	"fmt"
	"github.com/rvcevans/adventofcode/getinput"
)

var characters = map[rune]struct{}{
	'{': struct{}{},
	'}': struct{}{},
	'<': struct{}{},
	'>': struct{}{},
}

func main() {
	input := getinput.MustGet(2017, 9)

	cleaned, removed := removeCharacters(input[0])
	fmt.Println(score(cleaned))
	fmt.Println(removed)
}

func removeCharacters(s string) (string, int) {
	var out string
	var removed int
	var garbage, negate bool
	for _, r := range s {
		if garbage {
			if negate {
				negate = false
			} else {
				switch r {
				case '!':
					negate = true
				case '>':
					garbage = false
				default:
					removed++
				}
			}
		} else {
			switch r {
			case '<':
				garbage = true
			case '{', '}':
				out += string(r)
			default:
			}
		}

	}
	return out, removed
}

func score(s string) int {
	var total, value int
	for _, r := range s {
		switch r {
		case '{':
			value++
		case '}':
			total += value
			value--
		}
	}
	return total
}
