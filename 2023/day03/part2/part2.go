package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(sumGears(scanner))
}

// number represents a number found in the grid.
// It has a value and coordinates.
type number struct {
	value    int
	row      int
	colStart int
	colEnd   int
}

// approach to part 2
// for each '*'
//   find adjacent numbers
//   if != 2, continue
//   if 2, multiply numbers and add to total

// scratch that
// new approach
// loop over grid finding distinct numbers, storing their value and bounding box
// loop over grid finding '*', and check for adjacent numbers

// sumGears calculates the sum of gear ratios in a schematic.
func sumGears(scanner *bufio.Scanner) int {
	var total int
	var numbers []number

	grid := readGrid(scanner)

	// find numbers
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

				// record a number, in the previous char
				val, _ := strconv.Atoi(num)
				numbers = append(numbers, number{value: val, row: x, colStart: y - len(num), colEnd: y - 1})
				num = ""
			}
		}

		// we reached the end of a row and still have a number, so record it
		if num != "" {
			val, _ := strconv.Atoi(num)
			numbers = append(numbers, number{value: val, row: x, colStart: len(grid[x]) - len(num), colEnd: len(grid[x]) - 1})
		}
	}

	// find gears
	for x, row := range grid {
		// loop over each char in the row
		for y, char := range row {
			if char == '*' {
				// we found a possible gear, count adjacent numbers
				total += adjacentRatio(numbers, x, y)
			}
		}
	}

	return total
}

// adjacentRatio returns the ratio of the gear at the specified coordinates.
// If the gear is not adjacent to exactly two numbers, the ratio is 0.
// Otherwise, the ratio is the result of multiplying the two adjacent numbers.
func adjacentRatio(numbers []number, x, y int) int {
	var ratio int
	var tally int

	// check each number in the grid
	for _, num := range numbers {
		// if the number is not adjacent to the gear, skip it
		if !adjacent(num, x, y) {
			continue
		}

		// the number is adjacent, so bump up our tally for this gear
		tally++

		switch tally {
		// this is the first adjacent number
		case 1:
			ratio = num.value
		// this is the second adjacent number, multiple it
		case 2:
			ratio *= num.value
		// we have too many adjacent numbers, bail
		case 3:
			return 0
		}
	}

	if tally != 2 {
		return 0
	}

	return ratio
}

// adjacent determines whether a number is adjacent to the specified coordinates
func adjacent(num number, x, y int) bool {
	// num is more than one row up
	if x-1 > num.row {
		return false
	}

	// num is more than one row down
	if x+1 < num.row {
		return false
	}

	// num is more than one col left
	if y-1 > num.colEnd {
		return false
	}

	// num is more than one col right
	if y+1 < num.colStart {
		return false
	}

	return true
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
