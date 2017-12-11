// http://adventofcode.com/2017/day/11
// part 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var compass = map[string]int{
	"n":  0,
	"ne": 1,
	"se": 2,
	"s":  3,
	"sw": 4,
	"nw": 5,
}

// distance returns number of steps from start
// requires a squashed path p
func distance(p []int) (dist int) {
	for _, v := range p {
		dist += v
	}
	return
}

// squash eliminates all redundancies in p
func squash(p []int) []int {
	for this := range p {
		if p[this] == 0 {
			continue
		}

		// TODO: fix hard-coded assumption
		oppo := (this + 3) % len(p)
		switch {
		case p[oppo] == 0: //do nothing
		case p[oppo] >= p[this]:
			p[oppo] -= p[this]
			p[this] = 0
			continue
		case p[oppo] < p[this]:
			p[this] -= p[oppo]
			p[oppo] = 0
		}

		// TODO: fix hard-coded assumption
		diag := (this + 2) % len(p)
		next := (this + 1) % len(p)
		switch {
		case p[diag] == 0: //do nothing
		case p[diag] >= p[this]:
			p[diag] -= p[this]
			p[next] += p[this]
			p[this] = 0
		case p[diag] < p[this]:
			p[this] -= p[diag]
			p[next] += p[diag]
			p[diag] = 0
		}
	}
	return p
}

// traverse p and tally directions
func traverse(p []string) []int {
	t := make([]int, len(compass))

	for _, d := range p {
		t[compass[d]]++
		t = squash(t)
	}
	return t
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		path := strings.Split(scanner.Text(), ",")

		end := traverse(path)
		fmt.Println(distance(end))
	}
}
