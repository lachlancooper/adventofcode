// http://adventofcode.com/2017/day/11
// part 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// compass is an ordered ring of named directions
var directions = []string{"n", "ne", "se", "s", "sw", "nw"}
var axes = len(directions) / 2

// compass maps directions to natural numbers
var compass map[string]int

func init() {
	compass = make(map[string]int)
	for i, v := range directions {
		compass[v] = i
	}
}

// opposite returns the opposing direction of s
func opposite(s string) string {
	opp := (compass[s] + axes) % len(directions)
	return directions[opp]
}

// min returns the smallest opposed entry in t
func min(t map[string]int) (minkey string, minval int) {
	minval = 1 << 62

	for k, v := range t {
		_, opp := t[opposite(k)]
		if opp && v < minval {
			minkey = k
			minval = v
		}
	}
	return
}

// distance returns number of steps from start
func distance(t map[string]int) (d int) {
	for _, v := range t {
		d += v
	}
	return
}

// normalise cancels out redundant directions in t
func normalise(t map[string]int) map[string]int {
	for {
		k, v := min(t)
		if k == "" {
			break
		}

		// normalise
		o := opposite(k)
		t[o] -= v
		if t[o] == 0 {
			delete(t, o)
		}
		delete(t, k)

		if len(t) <= axes {
			break
		}
	}
	return t
}

// traverse p and tally directions
func traverse(p []string) map[string]int {
	t := make(map[string]int)

	// tally
	for _, i := range p {
		t[i]++
		t = normalise(t)
	}
	return normalise(t)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		path := strings.Split(scanner.Text(), ",")

		end := traverse(path)
		fmt.Println(end)
	}
}
