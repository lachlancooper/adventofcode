// http://adventofcode.com/2017/day/10
// part 2
package main

import (
	"bufio"
	"fmt"
	"os"
)

var suffix = []int{17, 31, 73, 47, 23}

// create initialises a list of size s
// list values are natural numbers
func create(s int) []int {
	list := make([]int, s)
	for i := range list {
		list[i] = i
	}
	return list
}

// asciidecode converts a string to a slice of ints
// according to each ascii value, e.g.
// 1,2,3 -> 49,44,50,44,51
func ascii(s string) []int {
	r := make([]int, len(s))
	for i, c := range s {
		r[i] = int(c)
	}
	return r
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

// xor performs bitwise numeric xor on ints e
func xor(e ...int) (x int) {
	for _, v := range e {
		x ^= v
	}
	return x
}

// densehash calculates the dense hash of list by
// xoring consecutive blocks of 16 numbers
func densehash(list []int) []int {
	blocksize := 16
	dense := make([]int, len(list)/blocksize)

	for i := range dense {
		start, end := i*blocksize, (i+1)*blocksize
		dense[i] = xor(list[start:end]...)
	}
	return dense
}

// sparsehash performs hashing over list based on lengths
// 64 rounds with currentpos and skipsize saved between rounds
func sparsehash(list, lengths []int) []int {
	rounds := 64
	currentpos := 0
	skipsize := 0

	for i := 0; i < rounds; i++ {
		for _, length := range lengths {
			reverse(list, currentpos, length)
			currentpos = (currentpos + skipsize + length) % len(list)
			skipsize++
		}
	}
	return list
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// initialise list
		list := create(256)

		// get lengths
		lengths := ascii(scanner.Text())

		// append arbitrary suffix
		lengths = append(lengths, suffix...)

		// perform sparse hashing
		list = sparsehash(list, lengths)

		// perform dense hashing
		list = densehash(list)

		// print hex representation
		for _, v := range list {
			fmt.Printf("%02x", v)
		}
		fmt.Println()
	}
}
