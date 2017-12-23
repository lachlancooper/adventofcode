// http://adventofcode.com/2017/day/22
// part 1
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

var grid = make(map[string]bool)

// print the current state of the grid, including virus
// only shows the cells surrounding virus
func printgrid() {
	x, y := virus.x, virus.y

	for i := x - 15; i <= x+15; i++ {
		for j := y - 15; j <= y+15; j++ {
			if x == i && y == j {
				fmt.Printf("[")
			} else {
				fmt.Printf(" ")
			}

			if grid[fmt.Sprintf("%v %v", i, j)] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
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

// burst wakes up the virus, which turns right if current
// node is infected or left otherwise, advances in the new
// direction, and flips the state of the current node
// repeats for c cycles
// returns the total number of infections
func burst(c int) (infections int) {
	for i := 0; i < c; i++ {
		viruspos := fmt.Sprintf("%v %v", virus.x, virus.y)

		// turn and infect or clean
		if grid[viruspos] {
			virus.facing = (virus.facing + 1) % 4

			delete(grid, viruspos)
		} else {
			virus.facing = (virus.facing + 3) % 4

			grid[viruspos] = true
			infections++
		}

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
				grid[pos] = true
			}
		}
	}

	// virus starts in the middle
	virus.x = i / 2
	virus.y = virus.x

	// run simulation
	infections := burst(10 * 1000)
	fmt.Println(infections)
}
