// http://adventofcode.com/2017/day/19
// part 2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var grid = make([]string, 0)

var packet = struct {
	x, y   int
	facing int
}{}

// return the next cell according to the current facing
func nextcell() byte {
	switch packet.facing {
	case 0:
		return grid[packet.y+1][packet.x]
	case 1:
		return grid[packet.y][packet.x-1]
	case 2:
		return grid[packet.y-1][packet.x]
	case 3:
		return grid[packet.y][packet.x+1]
	}
	return ' '
}

// advance packet to end of line
// return steps taken
func run() int {
	for i := 1; ; i++ {
		// take a step forward
		switch packet.facing {
		case 0:
			packet.y++
		case 1:
			packet.x--
		case 2:
			packet.y--
		case 3:
			packet.x++
		}

		// read and interpret cell at new position
		cell := grid[packet.y][packet.x]
		switch {
		case cell == '|' || cell == '-':
		case cell >= 'A' && cell <= 'Z':
		case cell == '+':
			// turn right
			packet.facing = (packet.facing + 1) % 4

			// if that was wrong, instead turn left
			if nextcell() == ' ' {
				packet.facing = (packet.facing + 2) % 4
			}
		case cell == ' ':
			return i
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	// find starting point
	packet.x = strings.IndexRune(grid[0], '|')

	fmt.Println(run())
}
