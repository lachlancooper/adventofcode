// http://adventofcode.com/2017/day/12
// part 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// programs are stored as strings
var programs = make(map[string][]string)

// count returns the number of programs connected to o
func count(o string) int {
	// keep a list of programs to check
	tosee := []string{o}
	seen := make(map[string]bool)

	for {
		// if it's empty, we're done
		if len(tosee) == 0 {
			break
		}

		src := tosee[0]
		tosee = tosee[1:]
		seen[src] = true
		for _, dst := range programs[src] {
			if !seen[dst] {
				tosee = append(tosee, dst)
			}
		}
	}

	return len(seen)
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

	fmt.Println(count("0"))
}
