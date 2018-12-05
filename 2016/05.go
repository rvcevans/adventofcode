package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/rvcevans/adventofcode/getinput"
)

func main() {
	input := getinput.MustGet(2016, 5)[0]

	i := 0
	solution1 := ""
	solution2 := []byte(strings.Repeat("_", 8))
	for strings.Contains(string(solution2), "_") {
		hash := hash(input + strconv.Itoa(i))
		if string(hash[:5]) == "00000" {
			if len(solution1) < 8 {
				solution1 += string(hash[5])
			}

			pos2, err := strconv.ParseInt(string(hash[5]), 16, 0)
			if err != nil {
				log.Fatalf("Cannot parse positions: %v", err)
			}
			if pos2 < 8 {
				if solution2[pos2] == '_' {
					solution2[pos2] = hash[6]
					fmt.Printf("Solution 2: %v\n", string(solution2))
				}
			}
		}

		i++
	}

	fmt.Printf("Solution 1: %v\n", string(solution1))
}

func hash(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
