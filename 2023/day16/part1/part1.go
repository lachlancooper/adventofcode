package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(solve(scanner))
}

type layout struct {
	grid  [][]tile
	beams []laser
}

type tile struct {
	object rune
	seen   []direction
}

type laser struct {
	row, col int
	heading  direction
}

type direction int

const (
	north direction = iota
	east
	south
	west
)

func solve(scanner *bufio.Scanner) int {
	// read grid objects into layout
	layout := layout{}
	for row := 0; scanner.Scan(); row++ {
		layout.grid = append(layout.grid, []tile{})
		for _, obj := range scanner.Text() {
			layout.grid[row] = append(layout.grid[row], tile{object: obj})
		}
	}

	// inject laser beam at top left, facing right
	layout.beams = []laser{{row: 0, col: -1, heading: east}}

	// run steps until all beams have reached their end
	for layout.step() {
		// fmt.Printf("%s\n", myLayout)
	}

	// count number of energised tiles
	return layout.energised()
}

// step advances each beam on the grid by one step.
// Returns false if the grid has no more beams.
func (l *layout) step() bool {
	if len(l.beams) == 0 {
		return false
	}

	nextBeams := []laser{}
	for _, beam := range l.beams {
		switch beam.heading {
		case north:
			beam.row--
		case east:
			beam.col++
		case south:
			beam.row++
		case west:
			beam.col--
		}

		// next tile is beyond grid edge, stop tracking it
		if beam.row < 0 ||
			beam.row > len(l.grid)-1 ||
			beam.col < 0 ||
			beam.col > len(l.grid[0])-1 {
			continue
		}

		tile := l.grid[beam.row][beam.col]

		// next tile has already seen a beam in this direction, stop tracking it
		if slices.Contains(tile.seen, beam.heading) {
			continue
		}

		// mark tile as seen, in this direction
		l.grid[beam.row][beam.col].seen = append(tile.seen, beam.heading)

		// figure out which object the beam interacts with
		switch tile.object {
		case '/':
			switch beam.heading {
			case north:
				beam.heading = east
			case east:
				beam.heading = north
			case south:
				beam.heading = west
			case west:
				beam.heading = south
			}
		case '\\':
			switch beam.heading {
			case north:
				beam.heading = west
			case east:
				beam.heading = south
			case south:
				beam.heading = east
			case west:
				beam.heading = north
			}
		case '|':
			switch beam.heading {
			case east, west:
				// redirect this beam
				beam.heading = south
				// add another beam facing the other way
				split := beam
				split.heading = north
				nextBeams = append(nextBeams, split)
			}
		case '-':
			switch beam.heading {
			case north, south:
				// redirect this beam
				beam.heading = east
				// add another beam facing the other way
				split := beam
				split.heading = west
				nextBeams = append(nextBeams, split)
			}
		case '.':
			// no interaction
		}

		nextBeams = append(nextBeams, beam)
	}

	l.beams = nextBeams

	return true
}

// energised counts the number of energised tiles on a grid.
func (l layout) energised() int {
	result := 0

	for _, row := range l.grid {
		for _, obj := range row {
			if len(obj.seen) > 0 {
				result++
			}
		}
	}

	return result
}

func (l layout) String() string {
	result := ""
	for _, row := range l.grid {
		for _, tile := range row {
			o := tile.object
			if o == '.' {
				switch len(tile.seen) {
				case 0:
					// print '.'
				case 1:
					switch tile.seen[0] {
					case north:
						o = '^'
					case east:
						o = '>'
					case south:
						o = 'v'
					case west:
						o = '<'
					}
				default:
					// print the number of distinct directions seen
					o = rune(strconv.Itoa(len(tile.seen))[0])
				}
			}
			result += string(o)
		}
		result += "\n"
	}
	return result
}

func (d direction) String() string {
	switch d {
	case north:
		return "^"
	case east:
		return ">"
	case south:
		return "v"
	case west:
		return "<"
	}
	return ""
}
