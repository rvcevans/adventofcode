package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	log.Printf("Solution 1: %d", highestScore(424, 71144))
}

func highestScore(players, lastMarble int) int {
	g := newGame(players)
	for g.marble <= lastMarble {
		g.next()
	}
	return g.highest()
}

type game struct {
	circle   []int
	position int
	players  int
	player   int
	marble   int
	scores   map[int]int
}

func (g *game) String() string {
	var ss []string
	for p, c := range g.circle {
		s := strconv.Itoa(c)
		if p == g.position {
			s = "(" + s + ")"
		}
		ss = append(ss, s)
	}
	return fmt.Sprintf("[%d] %s", g.player, strings.Join(ss, " "))
}

func (g *game) highest() int {
	h := 0
	for _, s := range g.scores {
		if s > h {
			h = s
		}
	}
	return h
}

func (g *game) next() {
	if g.marble% 23 == 0 {
		g.scores[g.player] += g.marble
		g.move(-7)
		g.scores[g.player] += g.circle[g.position]
		fmt.Println(g.circle[g.position])
		g.circle = append(g.circle[:g.position], g.circle[g.position+1:]...)
	} else {
		g.move(2)
		g.circle = append(g.circle[:g.position], append([]int{g.marble}, g.circle[g.position:]...)...)
	}
	g.player ++
	if g.player > g.players {
		g.player -= g.players
	}
	g.marble ++
}

func (g *game) move(i int) {
	g.position += i
	for g.position > len(g.circle) {
		g.position -= len(g.circle)
	}
	for g.position < 0 {
		g.position += len(g.circle)
	}
}

func newGame(players int) *game {
	return &game{
		[]int{0},
		0,
		players,
		1,
		1,
		map[int]int{},
	}
}
