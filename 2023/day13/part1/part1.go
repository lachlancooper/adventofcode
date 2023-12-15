package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(solve(scanner))
}

type pattern []string

func solve(scanner *bufio.Scanner) int {
	total := 0

	// scan each pattern separately
	var pat pattern
	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			pat = append(pat, line)
			continue
		}

		// blank line, check the previous pattern
		total += pat.findReflection()
		pat = pattern{}
	}

	// check the last pattern (no trailing blank line)
	total += pat.findReflection()

	return total
}

// findReflection finds the reflection point in the given pattern
func (p pattern) findReflection() int {
	// check for vertical line of reflection
	for col := 1; col < len(p[0]); col++ {
		if p.reflectVert(col) {
			return col
		}
	}

	// check horizontal line of reflection
	for row := 1; row < len(p); row++ {
		if p.reflectHori(row) {
			return 100 * row
		}
	}

	panic("no reflection found in pattern")
}

// reflectVert checks for reflection around a vertical line at col.
func (p pattern) reflectVert(col int) bool {
	// for each line in pattern
	for _, line := range p {
		// compare char by char, on either side of col
		for i := 0; i < col && i < len(line)-col; i++ {
			if line[col-i-1] != line[col+i] {
				return false
			}
		}
	}

	return true
}

// reflectHori checks for reflection around a horizontal line at row.
func (p pattern) reflectHori(row int) bool {
	// compare line by line, on either side of row
	for i := 0; i < row && i < len(p)-row; i++ {
		if p[row-i-1] != p[row+i] {
			return false
		}
	}

	return true
}

func (p pattern) String() string {
	result := ""
	for _, line := range p {
		result += line + "\n"
	}

	return result
}
