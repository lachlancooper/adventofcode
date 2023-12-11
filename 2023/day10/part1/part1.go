package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(countSteps(scanner))
}

// Algorithm
// - read input into grid / array, saving 'S' as start
// - pick a valid direction away from start
// - follow directions, counting squares until return to start
// - answer is squares / 2 + 1

// Principles
// - define valid directions for each char in one place

type tile struct {
	y, x int
}

func (t tile) String() string {
	return fmt.Sprintf("(%d,%d)", t.y, t.x)
}

type direction int

const (
	invalid direction = iota
	north
	east
	south
	west
)

func (d direction) String() string {
	switch d {
	case north:
		return "north"
	case east:
		return "east"
	case south:
		return "south"
	case west:
		return "west"
	default:
		return "invalid"
	}
}

// ends returns the ends of a given pipe.
func ends(char byte) (direction, direction) {
	switch char {
	case '|':
		return north, south
	case '-':
		return east, west
	case 'L':
		return north, east
	case 'J':
		return north, west
	case '7':
		return south, west
	case 'F':
		return south, east
	default:
		return invalid, invalid
	}
}

// opposite of the given direction.
func (d direction) opposite() direction {
	switch d {
	case north:
		return south
	case south:
		return north
	case east:
		return west
	case west:
		return east
	}

	return invalid
}

func countSteps(scanner *bufio.Scanner) int {
	var start tile

	// read grid
	grid := [][]byte{}
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Bytes()
		b1 := make([]byte, len(line))
		copy(b1, line)
		grid = append(grid, b1)
		if pos := strings.Index(string(line), "S"); pos >= 0 {
			start = tile{i + 1, pos + 1}
		}
	}

	// pick any valid direction away from start
	heading := pickStartHeading(grid, start)

	// until return to start, just navigate
	steps := 0
	for current := start; ; steps++ {
		switch heading {
		case north:
			current.y--
		case south:
			current.y++
		case east:
			current.x++
		case west:
			current.x--
		}

		heading = navigate(grid, current, heading)

		if current == start {
			break
		}
	}

	return steps/2 + 1
}

// navigate decides the next direction after entering the specified grid tile with the given heading.
func navigate(grid [][]byte, current tile, heading direction) direction {
	// don't try to navigate beyond the grid limits
	if current.y-1 < 0 || current.y-1 >= len(grid) || current.x-1 < 0 || current.x-1 >= len(grid[current.y-1]) {
		return invalid
	}

	// get pipe at this tile
	p := grid[current.y-1][current.x-1]
	dir1, dir2 := ends(p)

	if dir1 != heading.opposite() && dir2 != heading.opposite() {
		// if this pipe doesn't have an exit leading back the way we came, we've done something wrong
		return invalid
	} else if dir1 == heading.opposite() {
		// don't just exit via the way we entered
		return dir2
	} else {
		return dir1
	}
}

// pickStartHeading determines a valid heading from the start location.
func pickStartHeading(grid [][]byte, start tile) direction {
	// test each direction in turn
	for _, heading := range []direction{north, east, south, west} {
		// simulate a move to the next grid tile
		next := start

		switch heading {
		case north:
			next.y--
		case south:
			next.y++
		case east:
			next.x++
		case west:
			next.x--
		}

		// the move was invalid
		if navigate(grid, next, heading) == invalid {
			continue
		}

		return heading
	}

	panic("Could not find a valid start direction")
}
