// http://adventofcode.com/2017/day/13
// part 2
//
// caught iff you meet a scanner on its cycle, i.e.
// your arrival (layer + delay) % depth*2-2 == 0
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var firewall = make(map[int]int)

func sieve(n int) int {
	caught := make([]bool, n+1)
	// layer 0 always catches us at time 0
	caught[0] = true

	// take out all invalid starting times
	for layer, depth := range firewall {
		// each depth defines a period
		period := depth*2 - 2

		// first time eliminated is (period - layer)
		// but must be > 0
		starttime := period - layer
		if starttime < 0 {
			starttime += ((-(period - layer) / period) + 1) * period
		}

		// Update all multiples of period
		for i := starttime; i <= n; i += period {
			caught[i] = true
		}
	}

	// return the first valid starting time <= n
	for i, c := range caught {
		if !c {
			return i
		}
	}

	// there is no valid starting time <= n
	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ": ")
		layer, _ := strconv.Atoi(line[0])
		depth, _ := strconv.Atoi(line[1])
		firewall[layer] = depth
	}

	// adjust sieve size as required
	fmt.Println(sieve(1000 * 1000 * 10))
}
