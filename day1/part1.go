// Package adventofcodeday1 solves
// http://adventofcode.com/2017/day/1
// part 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// uncaptcha returns the sum of all digits that match the next digit in the list
// the circular list (loops from end to start)
func uncaptcha(s string) int {
	sum := 0
	d := strings.Split(s, "")
	l := len(d)

	for i := range d {
		cur, _ := strconv.Atoi(d[i])
		next, _ := strconv.Atoi(d[(i+1)%l])

		if cur == next {
			sum += cur
		}
	}

	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(uncaptcha(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
