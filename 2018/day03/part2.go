// http://adventofcode.com/2018/day/3
// part 2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type claim struct {
	id     int
	origin []int
	size   []int
}

// atoi converts a slice of strings to ints
func atoi(s []string) (r []int) {
	for _, c := range s {
		i, _ := strconv.Atoi(c)
		r = append(r, i)
	}
	return
}

// do the given claims overlap?
func overlap(c1, c2 claim) bool {
	left := &c1
	right := &c2
	top := &c1
	bottom := &c2

	if c1.origin[0] > c2.origin[0] {
		left = &c2
		right = &c1
	}
	if c1.origin[1] > c2.origin[1] {
		top = &c2
		bottom = &c1
	}

	return (left.origin[0]+left.size[0] > right.origin[0]) && (top.origin[1]+top.size[1] > bottom.origin[1])
}

func main() {
	claims := make([]claim, 0)
	invalid := make(map[int]bool)
	scanner := bufio.NewScanner(os.Stdin)

	// read all claims
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		id, _ := strconv.Atoi(strings.Trim(line[0], "#"))
		origin := atoi(strings.Split(strings.Trim(line[2], ":"), ","))
		size := atoi(strings.Split(line[3], "x"))

		claims = append(claims, claim{id, origin, size})
	}

	// loop through claims
	for i, candidate := range claims {
		// check against all other claims
		for _, c := range claims[i+1:] {
			// if there's an overlap, eliminate both claims
			if overlap(candidate, c) {
				invalid[candidate.id] = true
				invalid[c.id] = true
			}
		}

		if !invalid[candidate.id] {
			// we discovered no overlaps, this must be the winner
			fmt.Println(candidate.id)
			return
		}
	}
}
