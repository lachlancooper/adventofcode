// http://adventofcode.com/2017/day/24
// part 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// fixed list of available ports
var components = make([][]int, 0)

// list of valid bridges to test
var bridges = make([][]int, 0)

// atoi converts a slice of strings into a slice of ints
func atoi(s []string) (r []int) {
	for _, c := range s {
		i, _ := strconv.Atoi(c)
		r = append(r, i)
	}
	return r
}

// copybridge copies a slice of ints
func copybridge(oldbridge []int) []int {
	newbridge := make([]int, len(oldbridge))
	copy(newbridge, oldbridge)
	return newbridge
}

// copyparts copies a specific type of map
func copyparts(oldmap map[int]bool) map[int]bool {
	newmap := make(map[int]bool, len(oldmap))

	for k, v := range oldmap {
		newmap[k] = v
	}
	return newmap
}

// return strength of given bridge
func strength(bridge []int) (s int) {
	for _, c := range bridge {
		for _, p := range components[c] {
			s += p
		}
	}
	return
}

// return strength of strongest bridge
func strongest() (s int) {
	for _, b := range bridges {
		str := strength(b)
		if s < str {
			s = str
		}
	}
	return
}

// match checks whether component c matches port
// returns the value of the other port on the given component
// also returns true if matched, false otherwise
func match(c, port int) (int, bool) {
	for i, p := range components[c] {
		if port == p {
			otherport := components[c][1-i]
			return otherport, true
		}
	}
	return -1, false
}

// return all the components that could connect to port p
func findcomponents(port int) []int {
	comp := make([]int, 0)

	for i, c := range components {
		for _, p := range c {
			if port == p {
				comp = append(comp, i)
				break
			}
		}
	}
	return comp
}

// build all valid bridges from given set of parts, starting from given bridge and port
func buildbridges(parts map[int]bool, bridge []int, port int) {
	// check all parts
	for c := range parts {
		// does this part match?
		if newport, found := match(c, port); found {

			// yes, add this to a new bridge
			newbridge := copybridge(bridge)
			newbridge = append(newbridge, c)

			// save this bridge
			bridges = append(bridges, newbridge)

			// consume component and continue building from this bridge
			newparts := copyparts(parts)
			delete(newparts, c)
			buildbridges(newparts, newbridge, newport)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// read component list
	for scanner.Scan() {
		components = append(components, atoi(strings.Split(scanner.Text(), "/")))
	}

	// generate all valid bridges
	// start with all components, empty bridge, port 0
	parts := make(map[int]bool, len(components))
	for c := range components {
		parts[c] = true
	}
	buildbridges(parts, []int{}, 0)

	fmt.Println(strongest())
}
