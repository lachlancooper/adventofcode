// http://adventofcode.com/2017/day/6
// part 2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// atoi converts a slice of strings to ints
func atoi(s []string) (r []int) {
	for _, c := range s {
		i, _ := strconv.Atoi(c)
		r = append(r, i)
	}
	return
}

// redistribute performs a single redistribution
// cycle over banks b
func redistribute(b []int) {
	var biggestbank, mostblocks int

	// find bank with most blocks
	for bank, blocks := range b {
		if blocks > mostblocks {
			mostblocks = blocks
			biggestbank = bank
		}
	}

	// remove all blocks from biggest bank
	b[biggestbank] = 0

	// redistribute
	for i := (biggestbank + 1) % len(b); mostblocks > 0; i = (i + 1) % len(b) {
		b[i]++
		mostblocks--
	}
}

// cycles returns the number of cycles in the infinite
// loop that arises from redistributing blocks over banks b
func looplength(b []int) int {
	var state string
	seen := make(map[string]int)

	for cycles := 0; ; {
		state = fmt.Sprint(b)

		switch seen[state] {
		case 1:
			cycles++
		case 2:
			return cycles
		}

		seen[state]++

		redistribute(b)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		bank := atoi(strings.Fields(scanner.Text()))
		fmt.Println(looplength(bank))
	}
}
