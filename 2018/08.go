package main

import (
	"github.com/rvcevans/adventofcode/getinput"
	"log"
	"strconv"
	"strings"
)

func main() {
	var is []int
	for _, s := range strings.Split(getinput.MustGet(2018, 8)[0], " ") {
		i, _ := strconv.Atoi(s)
		is = append(is, i)
	}

	node, _ := parseNode(is)
	log.Printf("Solution 1: %d", node.MetaTotal())
	log.Printf("Solution 2: %d", node.ValueTotal())
}

type node struct {
	children []*node
	meta     []int
}

func (n *node) MetaTotal() int {
	t := 0
	for _, c := range n.children {
		t += c.MetaTotal()
	}
	for _, m := range n.meta {
		t += m
	}
	return t
}

func (n *node) ValueTotal() int {
	if len(n.children) == 0 {
		return n.MetaTotal()
	} else {
		t := 0
		for _, m := range n.meta {
			if m <= len(n.children) {
				t += n.children[m-1].ValueTotal()
			}
		}
		return t
	}
}

func parseNode(s []int) (*node, []int) {
	n := node{
		children: make([]*node, s[0]),
		meta:     make([]int, s[1]),
	}
	s = s[2:]
	for j := range n.children {
		n.children[j], s = parseNode(s)
	}
	for j := range n.meta {
		n.meta[j] = s[j]
	}
	return &n, s[len(n.meta):]
}
