// http://adventofcode.com/2017/day/14
// part 2

// this puzzle is so hard! ideas for a solution

// maintain two maps:
// map[coordinate]int - what group is this coordinate in?
// map[int]coordinate - what coordinates are in this group?

// (int is unique group id, coordinate is position in map)

// 1. consider one row, sort into horizontal groups. assign each grouped coordinate a group number, and each group its coordinates
// 2. consider the next row, ditto
// 3. then compare the rows, looking for vertically adjacent groups. if found, "flood fill" by picking one of the groups as the new id and moving all coordinates from the other group over, then updating map[coordinate]int accordingly
// 4. repeat 2-3.

// final answer is size of map[int]coordinate

package main

import (
	"bufio"
	"fmt"
	"os"
)

type cell struct {
	x int
	y int
}

var suffix = []int{17, 31, 73, 47, 23}
var grid2group = make(map[cell]int)
var group2grid = make(map[int][]cell)

// grid is 128x128
var grid = make([][]bool, 128)

func init() {
	for i := range grid {
		grid[i] = make([]bool, 128)
	}
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

// asciidecode converts a string to a slice of ints
// according to each ascii value, e.g.
// 1,2,3 -> 49,44,50,44,51
// appends "-<r>" before converting
func ascii(s string, row int) []int {
	s += fmt.Sprintf("-%v", row)
	// fmt.Println(s)
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

// replacegroup sets all cells currently in s to d, and removes s entirely
func replacegroup(s, d int) {
	for _, c := range group2grid[s] {
		assigngroup(c, d)
	}
	delete(group2grid, s)
}

// assigngroup puts c in g
func assigngroup(c cell, g int) {
	grid2group[c] = g
	group2grid[g] = append(group2grid[g], c)
}

// copygroup checks whether c already has a group
// if not, it assigns the group
// otherwise it replaces the group
func copygroup(c cell, g int) {
	cg := grid2group[c]
	if cg == 0 {
		assigngroup(c, g)
	} else {
		replacegroup(cg, g)
	}
}

// countgroups traverses grid, assigning each cell a group based on its adjacency
// if cell contains true
// first consider the cell to the left. if it is true, copy its group
// then consider the cell above. if it is true and we have a different group,
// update all members of our group to its group. otherwise, just copy its group
func paintgroups() {
	// g keeps track of unique group number
	g := 1

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// skip free cells
			if !grid[i][j] {
				continue
			}

			c := cell{i, j}
			// if we have a left neighbour that's active, adopt their group
			if j > 0 && grid[i][j-1] {
				n := cell{i, j - 1}
				assigngroup(c, grid2group[n])
			}
			// if we have a top neighbour that's active and different, *copy* their group
			if n := (cell{i - 1, j}); i > 0 && grid[i-1][j] && grid2group[n] != grid2group[c] {
				copygroup(c, grid2group[n])
			}
			// if we still don't have a group, assign a new one
			if grid2group[c] == 0 {
				assigngroup(c, g)
				g++
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// calculate 128 knot hashes
		for i := 0; i < 128; i++ {
			// initialise list
			list := create(256)

			// get key
			lengths := ascii(scanner.Text(), i)

			// append arbitrary suffix
			lengths = append(lengths, suffix...)

			// perform sparse hashing
			list = sparsehash(list, lengths)

			// perform dense hashing
			list = densehash(list)

			var s string
			// generate binary representation
			for _, v := range list {
				s += fmt.Sprintf("%08b", v)
				for j, c := range s {
					if c == '1' {
						grid[i][j] = true
					}
				}
			}
		}

		paintgroups()

		fmt.Println(len(group2grid))
	}
}
