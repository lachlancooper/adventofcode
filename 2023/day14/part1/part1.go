package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(solve(scanner))
}

// platform represents a platform, with dimensions and a number of objects.
type platform struct {
	rows, cols     int
	rounded, cubic []object
}

// object represents a fixed or movable object, with a position.
type object struct {
	row, col int
}

type direction int

const (
	north direction = iota
	east
	south
	west
)

func solve(scanner *bufio.Scanner) int {
	input := platform{}

	// scan all objects on the platform
	row := 0
	for ; scanner.Scan(); row++ {
		for i, val := range scanner.Text() {
			input.cols = i + 1
			switch val {
			case '.':
				continue
			case 'O':
				input.rounded = append(input.rounded, object{row, i})
			case '#':
				input.cubic = append(input.cubic, object{row, i})
			}
		}
	}
	// set final dimensions of the platform
	input.rows = row

	// tilt the lever to the north
	input = input.tilt(north)

	// calculate total load on the platform
	return input.getLoad()
}

// getLoad calculates the overall load for a given arrangement of rounded rocks.
//
// For example:
// OOOO.#.O.. 10
// OO..#....#  9
// OO..O##..O  8
// O..#.OO...  7
// ........#.  6
// ..#....#.#  5
// ..O..#.O.O  4
// ..O.......  3
// #....###..  2
// #....#....  1
//
// Total 136
func (p platform) getLoad() int {
	total := 0

	for _, rock := range p.rounded {
		total += (p.rows - rock.row)
	}

	return total
}

// tilt tilts the platform in the given direction.
// All rounded rocks will roll in that direction.
// Repeat tilt until platform stops changing.
func (p platform) tilt(dir direction) platform {
	stepOne := p.step(dir)

	// repeat until rocks stop rolling
	stepTwo := stepOne.step(dir)
	if !slices.Equal(stepOne.rounded, stepTwo.rounded) {
		return stepTwo.tilt(dir)
	}

	return stepOne
}

// step moves rounded rocks one step in the specified direction.
func (p platform) step(dir direction) platform {
	result := p.Copy()

	if dir != north {
		panic("direction not implemented")
	}

	for i, rock := range p.rounded {
		// rocks roll to the top of the platform
		if rock.row == 0 {
			continue
		}

		// rocks are blocked by other objects
		if result.hasObjectAt(rock.row-1, rock.col) {
			continue
		}

		// otherwise, rocks roll north
		result.rounded[i].row -= 1
	}

	return result
}

// hasObjectAt detects whether the platform has an object at the specified coordinates.
func (p platform) hasObjectAt(row, col int) bool {
	return p.getObjectAt(row, col) != '.'
}

func (p platform) getObjectAt(row, col int) rune {
	for _, rock := range p.cubic {
		if rock.row == row && rock.col == col {
			return '#'
		}
	}

	for _, rock := range p.rounded {
		if rock.row == row && rock.col == col {
			return 'O'
		}
	}

	return '.'
}

func (p platform) String() string {
	output := ""

	for row := 0; row < p.rows; row++ {
		for col := 0; col < p.cols; col++ {
			output += string(p.getObjectAt(row, col))
		}
		output += "\n"
	}

	return output
}

func (p platform) Copy() platform {
	result := platform{}

	result.rows, result.cols = p.rows, p.cols

	result.cubic = make([]object, len(p.cubic))
	result.rounded = make([]object, len(p.rounded))

	copy(result.cubic, p.cubic)
	copy(result.rounded, p.rounded)

	return result
}
