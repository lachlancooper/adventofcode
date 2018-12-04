// http://adventofcode.com/2017/day/22
// part 2
package main

import (
	"bufio"
	"fmt"
	"os"
)

var virus = struct {
	// 0 = up
	// 1 = right
	// 2 = down
	// 3 = left
	facing int
	x, y   int
}{}

// 0 = clean
// W = weakened
// # = infected
// F = flagged
var grid = make(map[string]rune)

// print the current state of the grid, including virus
// only shows a region surrounding virus
func printgrid() {
	x, y := virus.x, virus.y

	for i := x - 10; i <= x+10; i++ {
		for j := y - 10; j <= y+10; j++ {
			if x == i && y == j {
				fmt.Printf("[")
			} else {
				fmt.Printf(" ")
			}

			pos := fmt.Sprintf("%v %v", i, j)
			if grid[pos] == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%c", grid[pos])
			}

			if x == i && y == j {
				fmt.Printf("]")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// burst wakes up the virus and activates it
// turns right if node is infected, left if clean,
// backwards if flagged, no turn if weakened
// infects or cleans current node
// advances in the new direction
// repeats for c cycles
// returns the total number of infections
func burst(c int) (infections int) {
	for i := 0; i < c; i++ {
		viruspos := fmt.Sprintf("%v %v", virus.x, virus.y)

		// turn and infect or clean
		switch grid[viruspos] {
		case 0:
			virus.facing += 3

			grid[viruspos] = 'W'
		case '#':
			virus.facing += 1

			grid[viruspos] = 'F'
		case 'W':
			virus.facing += 0

			grid[viruspos] = '#'
			infections++
		case 'F':
			virus.facing += 2

			delete(grid, viruspos)
		}
		virus.facing %= 4

		// move
		switch virus.facing {
		case 0:
			virus.x--
		case 1:
			virus.y++
		case 2:
			virus.x++
		case 3:
			virus.y--
		}
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// read in starting grid
	var i int
	for i = 0; scanner.Scan(); i++ {
		for j, c := range scanner.Text() {
			if c == '#' {
				pos := fmt.Sprintf("%v %v", i, j)
				grid[pos] = c
			}
		}
	}

	// virus starts in the middle
	virus.x = i / 2
	virus.y = virus.x

	// run simulation
	infections := burst(10 * 1000 * 1000)
	fmt.Println(infections)
}
