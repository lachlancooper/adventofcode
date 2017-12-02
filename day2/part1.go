// http://adventofcode.com/2017/day/2
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

// rowdiff returns the difference between the largest
// and smallest integer values on the given row.
// values are separated by whitespace
func rowdiff(s string) int {
	d := atoi(strings.Fields(s))
	min, max := d[0], d[0]

	for _, v := range d {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return max - min
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	checksum := 0

	for scanner.Scan() {
		checksum += rowdiff(scanner.Text())
	}

	fmt.Println(checksum)
}
