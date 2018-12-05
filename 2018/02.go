package main

import (
	"github.com/rvcevans/adventofcode/getinput"
	"log"
)

func main() {
	ids := getinput.MustGet(2018, 2)

	repeatCount := map[int]int{}
	for _, id := range ids {
		letterCount := map[rune]int{}
		for _, r := range id {
			letterCount[r] ++
		}
		repeats := map[int]struct{}{}
		for _, v := range letterCount {
			repeats[v] = struct{}{}
		}
		for r := range repeats {
			repeatCount[r] ++
		}
	}

	log.Printf("Solution 1: %d", repeatCount[2]*repeatCount[3])
	
	for i := 0; i < len(ids[0]); i ++ {
		seen := map[string]struct{}{}
		for _, id := range ids {
			subid := id[:i] + id[i+1:]
			if _, ok := seen[subid]; ok {
				log.Printf("Solution 2: %s", subid)
				return
			}
			seen[subid] = struct{}{}
		}
	}
}
