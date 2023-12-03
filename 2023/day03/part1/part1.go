package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(sumParts(scanner))
}

// Approach 1:
// - read every line into a 2D slice
// - loop over lines looking for distinct numbers
// - for each number, look at all surrounding squares for adjacent symbols
// - if number is adjacent, add it to the total

// Approach 2:
// - find all numbers and save as objects with coordinates and values
// - find all symbols and save as objects with coordinates and values
// - for each number, check all symbols for adjacent ones
// - if number is adjacent, add it to the total

// Approach 3:
// - examine sets of 3 consecutive lines at a time

// sumParts calculates the sum of part numbers in a schematic.
func sumParts(scanner *bufio.Scanner) int {
	var total int

	grid := readGrid(scanner)

	// loop over each row in the grid
	for x, row := range grid {
		// num collects the number on a row
		var num string

		// loop over each char in the row
		for y, char := range row {
			switch char {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				// we started or continued a number
				num += string(char)
			default:
				// we don't have a number yet
				if num == "" {
					continue
				}

				// we reached the end of a number last square, so check it for adjacency
				if adjacent(grid, num, x, y-1) {
					val, _ := strconv.Atoi(num)
					total += val
				}
				num = ""
			}
		}

		// we reached the end of a row and still have a number, so check it for adjacency
		if num != "" {
			if adjacent(grid, num, x, len(grid[x])-1) {
				val, _ := strconv.Atoi(num)
				total += val
			}
		}
	}

	return total
}

// adjacent returns whether the number ending at the specified coordinates is adjacent to a symbol.
func adjacent(grid [][]rune, num string, x, y int) bool {
	return symbol(grid, len(num), x-1, y-len(num)) || // row above the number
		symbol(grid, len(num), x, y-len(num)) || // row of the number
		symbol(grid, len(num), x+1, y-len(num)) // row below the number
}

// symbol returns whether the string starting at the specified grid coordinates is a symbol.
func symbol(grid [][]rune, length, x, y int) bool {
	// can't check before the first row or after the last row
	if x < 0 || x >= len(grid) {
		return false
	}

	ymin := y
	ymax := y + length + 2

	// cap min and max of y
	if ymin < 0 {
		ymin = 0
	}

	if ymax >= len(grid[x]) {
		ymax = len(grid[x]) - 1
	}

	return strings.ContainsAny(string(grid[x][ymin:ymax]), "*#+$&=%-/@")
}

// readGrid loads the input into a slice of runes
func readGrid(scanner *bufio.Scanner) [][]rune {
	grid := [][]rune{}

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		row := make([]rune, len(line))

		for j, char := range line {
			row[j] = char
		}

		grid = append(grid, row)
	}

	return grid
}
