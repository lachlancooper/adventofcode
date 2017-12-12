// http://adventofcode.com/2017/day/12
// part 2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// programs are stored as strings
var programs = make(map[string][]string)
var groups = make(map[string][]string)

// group sorts all programs into buckets depending on interconnections
func group() {
	seen := make(map[string]bool)

	// check every program
	for c := range programs {
		// if we've seen it, move on
		if seen[c] {
			continue
		}

		// follow all connections for c
		tosee := []string{c}
		for {
			// nothing left for c
			if len(tosee) == 0 {
				break
			}

			src := tosee[0]
			tosee = tosee[1:]
			seen[src] = true
			groups[c] = append(groups[c], src)
			for _, dst := range programs[src] {
				if !seen[dst] {
					tosee = append(tosee, dst)
				}
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// read connections into programs map
	for scanner.Scan() {
		pipe := strings.Fields(scanner.Text())
		src := pipe[0]

		for _, dst := range pipe[2:] {
			dst = strings.TrimRight(dst, ",")
			programs[src] = append(programs[src], dst)
		}
	}

	group()

	fmt.Println(len(groups))
}
