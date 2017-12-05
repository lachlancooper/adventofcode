// http://adventofcode.com/2017/day/4
// part 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// steps returns the number of steps taken to exit
// the maze s of jump offsets
func steps(s []int) int {
	steps := 0
	for pos := 0; pos >= 0 && pos < len(s); steps += 1 {
		newpos := pos + s[pos]
		s[pos] += 1
		pos = newpos
	}
	return steps
}

func main() {
	var offsets []int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		offsets = append(offsets, line)
	}

	fmt.Println(steps(offsets))
}
