// http://adventofcode.com/2017/day/20
// part 2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type particle struct {
	p []int
	v []int
	a []int
}

var swarm = make(map[int]particle)

// abs returns the absolute value of i
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// atoi converts a slice of strings into a slice of ints
func atoi(s []string) (r []int) {
	for _, c := range s {
		i, _ := strconv.Atoi(c)
		r = append(r, i)
	}
	return
}

// collide destroys any particles occupying the same position
func collide() {
	// create map of unique positions this tick
	positions := make(map[string][]int)
	for i, p := range swarm {
		pos := fmt.Sprint(p.p)
		positions[pos] = append(positions[pos], i)
	}

	// destroy all particles occupying the same position
	for _, p := range positions {
		if len(p) > 1 {
			for _, i := range p {
				delete(swarm, i)
			}
		}
	}
}

// tick updates each particle in swarm according to its p, v, a
// runs for cycle count c
// colliding particles are destroyed after each tick
func tick(c int) {
	for i := 0; i < c; i++ {
		for _, p := range swarm {
			// update velocity
			for j := range p.v {
				p.v[j] += p.a[j]
			}
			// update position
			for j := range p.p {
				p.p[j] += p.v[j]
			}
		}

		collide()
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// read all particles
	for i := 0; scanner.Scan(); i++ {
		line := strings.Split(scanner.Text(), ", ")
		p := atoi(strings.Split(strings.Trim(line[0], "p=<>"), ","))
		v := atoi(strings.Split(strings.Trim(line[1], "v=<>"), ","))
		a := atoi(strings.Split(strings.Trim(line[2], "a=<>"), ","))

		swarm[i] = particle{p, v, a}
	}

	// arbitrary tick count
	tick(5 * 1000)

	fmt.Println(len(swarm))
}
