// http://adventofcode.com/2017/day/10
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
func atoi(s []string) []int {
	r := make([]int, len(s))
	for i, c := range s {
		j, _ := strconv.Atoi(c)
		r[i] = j
	}
	return r
}

// create initialises a list of size s
// list values are natural numbers
func create(s int) []int {
	list := make([]int, s)
	for i := range list {
		list[i] = i
	}
	return list
}

// reverse works inplace to reverse the elements
// of list given start and length
func reverse(list []int, start, length int) {
	end := (start + length - 1)
	l := len(list)

	for i, j := start, end; i < j; i, j = i+1, j-1 {
		list[i%l], list[j%l] = list[j%l], list[i%l]
	}
}

// hash performs hashing over list based on lengths
func hash(list, lengths []int) {
	currentpos := 0

	for skip, length := range lengths {
		reverse(list, currentpos, length)
		currentpos = (currentpos + skip + length) % len(list)
	}
}

// checkSum calculates the checksum of a list
// by multiplying the first two values
func checkSum(list []int) int {
	return list[0] * list[1]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		// initialise list
		size, _ := strconv.Atoi(line[0])
		list := create(size)

		// get lengths
		lengths := atoi(strings.Split(line[1], ","))

		// perform hashing
		hash(list, lengths)

		// print checksum
		fmt.Println(checkSum(list))
	}
}
