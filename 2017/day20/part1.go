// http://adventofcode.com/2017/day/20
// part 1
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

var swarm = make([]particle, 0)

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
	return r
}

// tick updates each particle in swarm according to its p, v, a
// runs for cycle count c
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
	}
}

// manhattan returns the distance of particle p from origin
func manhattan(p particle) (r int) {
	for _, i := range p.p {
		r += abs(i)
	}
	return
}

// closest returns the particle in swarm with the lowest manhattan distance
func closest() (minparticle int) {
	minparticle = -1
	mindistance := 1<<63 - 1
	for i, p := range swarm {
		d := manhattan(p)
		if d < mindistance {
			mindistance = d
			minparticle = i
		}
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// read all particles
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ", ")
		p := atoi(strings.Split(strings.Trim(line[0], "p=<>"), ","))
		v := atoi(strings.Split(strings.Trim(line[1], "v=<>"), ","))
		a := atoi(strings.Split(strings.Trim(line[2], "a=<>"), ","))

		swarm = append(swarm, particle{p, v, a})
	}

	// arbitrary tick count
	tick(50 * 1000)

	fmt.Println(closest())
}
