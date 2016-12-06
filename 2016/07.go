package main

import (
	"fmt"
	"github.com/rvcevans/adventofcode/getinput"
	"os"
)

func main() {
	input := getinput.MustGet(2016, 7, os.Getenv("ADVENT_SESSION"))

	count1 := 0
	count2 := 0
	for _, i := range input {
		ip := newIp(i)
		if ip.Tls() {
			count1 ++
		}
		if ip.Ssl() {
			count2 ++
		}
	}

	fmt.Printf("Solution1: %v\n", count1)
	fmt.Printf("Solution2: %v\n", count2)
}

type ip struct {
	// bool represents inside
	sequences map[bool][]string
}

func newIp(s string) *ip {
	sequences := make(map[bool][]string)
	current := ""
	in := false
	for _, r := range s {
		if r == '[' || r == ']' {
			sequences[in] = append(sequences[in], current)
			current = ""
			in = !in
		} else {
			current += string(r)
		}
	}
	sequences[in] = append(sequences[in], current)

	return &ip{sequences}
}

func (ip *ip) Tls() bool {
	for _, i := range ip.sequences[true] {
		if abba(i) {
			return false
		}
	}
	for _, o := range ip.sequences[false] {
		if abba(o) {
			return true
		}
	}
	return false
}

func (ip *ip) Ssl() bool {
	abas := make(map[string]struct{})
	for _, o := range ip.sequences[false] {
		for a := range abaList(o) {
			abas[a] = struct {}{}
		}

	}
	for _, i := range ip.sequences[true] {
		for b := range abaList(i){
			if _, ok := abas[bab(b)]; ok {
				return true
			}
		}
	}
	return false
}

func abba(s string) bool {
	for i := 0; i <= len(s)-4; i = i + 1 {
		if s[i] == s[i+3] && s[i+1] == s[i+2] && s[i] != s[i+1] {
			return true
		}

	}
	return false
}

func abaList(s string) map[string]struct{} {
	abas := make(map[string]struct{})
	for i := 0; i <= len(s)-3; i = i + 1 {
		if s[i] == s[i+2] && s[i] != s[i+1] {
			abas[s[i:i+3]] = struct {}{}
		}

	}
	return abas
}

func bab(aba string) string {
	return string([]byte{aba[1], aba[0], aba[1]})
}