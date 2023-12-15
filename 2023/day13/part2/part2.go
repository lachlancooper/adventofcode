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
		total += pat.findSmudgeReflection()
		pat = pattern{}
	}

	// check the last pattern (no trailing blank line)
	total += pat.findSmudgeReflection()

	return total
}

// findSmudgeReflection finds the smudged reflection point in the given pattern.
// This is required to be different to the "original" reflection point.
func (p pattern) findSmudgeReflection() int {
	original, _ := p.findReflection(0)

	testPattern := make(pattern, len(p))
	newChar := ""
	// range over every char in p, smudging it
	for i, line := range p {
		for j, char := range line {
			copy(testPattern, p)
			// smudge this char
			if char == '#' {
				newChar = "."
			} else {
				newChar = "#"
			}
			testPattern[i] = line[:j] + newChar + line[j+1:]

			if val, ok := testPattern.findReflection(original); ok {
				return val
			}
		}
	}

	panic("no smudged reflection found in pattern")
}

// findReflection finds the reflection point in the given pattern, and whether or not it was found.
// The specified reflection point will be skipped.
func (p pattern) findReflection(skip int) (int, bool) {
	// check for vertical line of reflection
	for col := 1; col < len(p[0]); col++ {
		if p.reflectVert(col) && col != skip {
			return col, true
		}
	}

	// check horizontal line of reflection
	for row := 1; row < len(p); row++ {
		if p.reflectHori(row) && 100*row != skip {
			return 100 * row, true
		}
	}

	return 0, false
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
