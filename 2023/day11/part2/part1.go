package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(sumPaths(scanner, 1_000_000))
}

// Algorithm
// - read input into grid
// - record which rows / columns are empty
// - record galaxy locations
// - for each pair of galaxies, find the distance between them

// Principles
// - first focus on getting the basic algorithm working, ignoring empty rows/columns

type galaxy struct {
	y, x int
}

func (g galaxy) String() string {
	return fmt.Sprintf("(%d,%d)", g.y, g.x)
}

func sumPaths(scanner *bufio.Scanner, multiple int) int {
	grid := [][]byte{}
	emptyRows := []int{}
	emptyCols := []int{}

	// read galaxies
	galaxies := []galaxy{}
	for i := 0; scanner.Scan(); i++ {
		emptyRow := true
		grid = append(grid, []byte(scanner.Text()))
		line := scanner.Text()
		for j, char := range line {
			if char == '#' {
				galaxies = append(galaxies, galaxy{i, j})
				emptyRow = false
			}
		}
		if emptyRow {
			emptyRows = append(emptyRows, i)
		}
	}

	for j := range grid[0] {
		emptyCol := true
		for i := range grid {
			if grid[i][j] == '#' {
				emptyCol = false
			}
		}
		if emptyCol {
			emptyCols = append(emptyCols, j)
		}
	}

	// calculate distances between galaxies
	length := 0
	for i, a := range galaxies {
		for _, b := range galaxies[i:] {
			if a == b {
				continue
			}
			length += distance(emptyRows, emptyCols, multiple, a, b)
		}
	}

	return length
}

func distance(rows, cols []int, multiple int, a, b galaxy) int {
	base := abs(a.y-b.y) + abs(a.x-b.x)

	rowBonus := 0
	for _, row := range rows {
		if row < a.y && row > b.y ||
			row < b.y && row > a.y {
			rowBonus += multiple - 1
		}
	}

	colBonus := 0
	for _, col := range cols {
		if col < a.x && col > b.x ||
			col < b.x && col > a.x {
			colBonus += multiple - 1
		}
	}

	total := base + rowBonus + colBonus
	// fmt.Printf("Distance between %s and %s: %d + %d + %d = %d\n", a, b, base, rowBonus, colBonus, total)
	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
