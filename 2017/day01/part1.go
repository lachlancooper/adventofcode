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

// atoi converts a slice of strings to ints
func atoi(s []string) (r []int) {
	for _, c := range s {
		i, _ := strconv.Atoi(c)
		r = append(r, i)
	}
	return
}

// uncaptcha returns the sum of all digits that match the next digit in
// the circular list (loops from end to start)
func uncaptcha(s string) (sum int) {
	d := atoi(strings.Split(s, ""))
	l := len(d)
	offset := 1

	for i := range d {
		cur := d[i]
		next := d[(i+offset)%l]

		if cur == next {
			sum += cur
		}
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(uncaptcha(scanner.Text()))
	}
}
