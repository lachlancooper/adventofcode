// Package adventofcodeday1 solves
// http://adventofcode.com/2017/day/1
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
func atoi(s []string) []int {
	r := make([]int, len(s))

	for i, c := range s {
		r[i], _ = strconv.Atoi(c)
	}
	return r
}

// uncaptcha returns the sum of all digits that match the opposite digit in
// the circular list (loops from end to start)
func uncaptcha(s string) int {
	sum := 0
	d := atoi(strings.Split(s, ""))
	l := len(d)
	offset := l / 2

	for i := range d {
		cur := d[i]
		next := d[(i+offset)%l]

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
}
