package main

import (
	"fmt"
	"github.com/rvcevans/adventofcode/getinput"
	"os"
	"strings"
)

func main() {
	input := getinput.MustGet(2017, 24)

	var components []*component
	for _, c := range input {
		s := strings.Split(c, "/")
		components = append(components, &component{mustInt(s[0]), mustInt(s[1])})
	}

	fmt.Println("Part1:", maxBridge(&bridge{end: 0}, components, false))
	fmt.Println("Part2:", maxBridge(&bridge{end: 0}, components, true))
}

type component struct {
	a, b int
}

type bridge struct {
	end        int
	components map[*component]struct{}
}

func (b *bridge) Add(c *component) (*bridge, bool) {
	if _, ok := b.components[c]; ok || (c.a != b.end && c.b != b.end) {
		return nil, false
	}
	next := c.a
	if c.a == b.end {
		next = c.b
	}

	m := map[*component]struct{}{c: struct{}{}}
	for c := range b.components {
		m[c] = struct{}{}
	}

	return &bridge{
		end:        next,
		components: m,
	}, true
}

func (b *bridge) Strength() int {
	total := 0
	for c := range b.components {
		total += c.a + c.b
	}
	return total
}

func (b *bridge) Len() int {
	return len(b.components)
}

func maxBridge(b *bridge, components []*component, longest bool) (int, int) {
	max, length := b.Strength(), b.Len()
	for _, c := range components {
		if d, ok := b.Add(c); ok {
			m, l := maxBridge(d, components, longest)
			if (longest && (l > length || (l == length && m > max))) || (!longest && m > max) {
				max, length = m, l
			}
		}
	}
	return max, length
}
