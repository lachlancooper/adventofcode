// http://adventofcode.com/2018/day/3
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

// fabric is 1000x1000
var fabric = make([][]rune, 1000)

// create strips of fabric
func init() {
	for i := range fabric {
		fabric[i] = make([]rune, 1000)
	}
}

// display fabric visually, for debugging
func printFabric() {
	for i := range fabric {
		for j := range fabric[i] {
			square := fabric[i][j]
			if square == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(square)
			}
		}
		fmt.Println()
	}
}

// count number of overlapping claim squares in fabric
func countOverlaps() (sum int) {
	for i := range fabric {
		for j := range fabric[i] {
			if fabric[i][j] > 1 {
				sum++
			}
		}
	}
	return sum
}

// stake the claim of given origin and size on fabric
func stakeClaim(origin []int, size []int) {
	for i := 0; i < size[0]; i++ {
		for j := 0; j < size[1]; j++ {
			fabric[origin[0]+i][origin[1]+j] += 1
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		origin := atoi(strings.Split(strings.Trim(line[2], ":"), ","))
		size := atoi(strings.Split(line[3], "x"))
		stakeClaim(origin, size)
	}

	// printFabric()
	fmt.Println(countOverlaps())
}
