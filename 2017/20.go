package main

import (
	"fmt"
	"github.com/rvcevans/adventofcode/getinput"
	"log"
	"os"
	"strconv"
	"strings"
	"math"
	"sort"
)

func main() {
	input := getinput.MustGet(2017, 20)

	var particles []*particle
	minAcceleration := -1
	for i, p := range input {
		s := strings.Split(p, ", ")
		p := &particle{
			number: i,
			p:      newVector3(s[0][2:]),
			v:      newVector3(s[1][2:]),
			a:      newVector3(s[2][2:]),
		}
		particles = append(particles, p)
		if acceleration := p.a.Distance(); minAcceleration == -1 || acceleration < minAcceleration {
			minAcceleration = acceleration
		}
	}

	type collision struct {
		a, b int
	}
	collisions := make(map[int][]collision)
	for i, a := range particles[:len(particles)-1] {
		for _, b := range particles[i+1:] {
			if t := a.CollisionTime(b); t >= 0 {
				collisions[t] = append(collisions[t], collision{a.number, b.number})
			}
		}
	}
	var collisionTimes []int
	for t := range collisions {
		collisionTimes = append(collisionTimes, t)
	}
	sort.Ints(collisionTimes)

	collided := make(map[int]struct{})
	for _, t := range collisionTimes {
		nextCollisions := make(map[int]struct{}) // We create a new map to handle simultaneous collisions
		for _, c := range collisions[t] {
			_, aOk := collided[c.a]
			_, bOk := collided[c.b]
			if !aOk && !bOk {
				nextCollisions[c.a] = struct{}{}
				nextCollisions[c.b] = struct{}{}
			}
		}
		for n := range nextCollisions {
			collided[n] = struct{}{}
		}
	}

	var minParticles []*particle
	for _, p := range particles {
		if p.a.Distance() == minAcceleration {
			minParticles = append(minParticles, p)
		}
	}

	dif := true
	for dif {
		dif = false
		for _, p := range minParticles {
			p.Next()
			dif = dif || !p.a.SameDirection(p.v)
		}
	}

	minDistance := -1
	var minParticle *particle
	for _, p := range minParticles {
		if d := p.p.Distance(); minDistance == -1 || d < minDistance {
			minDistance = d
			minParticle = p
		}
	}

	fmt.Println(minParticle.number)
	fmt.Println(len(particles)-len(collided))
}

type particle struct {
	number  int
	p, v, a vector3
}

func (p *particle) Next() {
	p.v = p.v.Add(p.a)
	p.p = p.p.Add(p.p)
}

// 1/2 * at(t + 1) + vt + p = 0
// at(t + 1) * 2vt + 2p = 0
// at^2 + (a + 2v)t + 2p = 0
func (a *particle) CollisionTime(b *particle) int {
	min, max := solveQuadraticPos(a.a.x - b.a.x, a.a.x + 2*a.v.x - b.a.x - 2*b.v.x, 2*a.p.x - 2*b.p.x)
	for _, v := range []int{min, max}{
		if v < 0 {
			continue
		}
		if quadraticSolution(a.a.y - b.a.y, a.a.y + 2*a.v.y - b.a.y - 2*b.v.y, 2*a.p.y - 2*b.p.y, v) &&
			quadraticSolution(a.a.z - b.a.z, a.a.z + 2*a.v.z - b.a.z - 2*b.v.z, 2*a.p.z - 2*b.p.z, v) {
			return v
		}
	}
	return -1
	
}

func newVector3(v string) vector3 {
	s := strings.Split(strings.Trim(v, "<>"), ",")
	return vector3{mustInt(s[0]), mustInt(s[1]), mustInt(s[2])}
}

type vector3 struct {
	x, y, z int
}

func (a vector3) Add(b vector3) vector3 {
	return vector3{a.x + b.x, a.y + b.y, a.z + b.z}
}

func (v vector3) Distance() int {
	return absInt(v.x) + absInt(v.y) + absInt(v.z)
}

func (a vector3) SameDirection(b vector3) bool {
	return sameSignInt(a.x, b.x) && sameSignInt(a.y, b.y) && sameSignInt(a.z, b.z)
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sameSignInt(a, b int) bool {
	return signInt(a) == 0 || signInt(b) == 0 || signInt(a) == signInt(b)
}

func signInt(a int) int {
	if a < 0 {
		return -1
	}
	if a > 0 {
		return 1
	}
	return 0
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Failed to convert %v to int: %v", s, err)
	}
	return i
}

func mustBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatalf("Failed to convert %v to bool: %v", s, err)
	}
	return b
}

// Returns false if no integer solutions exist
func solveQuadraticPos(a, b, c int) (int, int) {
	if a == 0 {
		if b == 0 {
			return -1, -1
		}
		if c % b != 0 {
			return -1, -1
		}
		return -c/b, -1
	}

	d := b * b - 4 * a * c
	if d < 0 {
		return -1, -1
	}

	s := math.Sqrt(float64(d))
	i := int(s)
	if i * i < d {
		i++
	}
	if i * i != d {
		return -1, -1
	}

	neg, pos := -b - i, -b +i
	f := func(p int, div int) int {
		if p % div != 0 {
			return -1
		}
		return p / div
	}
	return f(neg, 2 * a), f(pos, 2 *a)
}

func quadraticSolution(a, b, c, v int) bool {
	return a*v*v + b*v + c == 0
}