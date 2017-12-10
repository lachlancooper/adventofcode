// http://adventofcode.com/2017/day/2
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

// rowdiv finds the only two integer values that
// divide evenly on the given row and returns
// the result of this division.
// values are separated by whitespace
func rowdiv(s string) int {
	d := atoi(strings.Fields(s))

	for i := range d {
		for j := i + 1; j < len(d); j++ {
			if d[i]%d[j] == 0 {
				return d[i] / d[j]
			} else if d[j]%d[i] == 0 {
				return d[j] / d[i]
			}
		}
	}

	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	checksum := 0

	for scanner.Scan() {
		checksum += rowdiv(scanner.Text())
	}

	fmt.Println(checksum)
}
