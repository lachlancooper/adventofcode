package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(solve(scanner))
}

func solve(scanner *bufio.Scanner) int {
	total := 0

	// scan each line separately
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		// find the spring map, repeated five times
		springs := strings.Join([]string{fields[0], fields[0], fields[0], fields[0], fields[0]}, "?")

		// find the expectedVals values, repeated five times
		expectedVals := []int{}
		vals := strings.Split(fields[1], ",")
		for _, v := range vals {
			count, _ := strconv.Atoi(v)
			expectedVals = append(expectedVals, count)
		}

		expected := append(expectedVals, expectedVals...)
		expected = append(expected, expectedVals...)
		expected = append(expected, expectedVals...)
		expected = append(expected, expectedVals...)

		clear(seenSprings)
		findArrangements(springs, expected)
		total += countArrangments()
	}

	return total
}

var seenSprings = make(map[string]bool)

func countArrangments() int {
	total := 0
	for _, valid := range seenSprings {
		if valid {
			total++
		}
	}
	return total
}

// findArrangements counts the number of valid findArrangements of springs in the given input.
func findArrangements(input string, expected []int) {
	if seenSprings[input] {
		return
	}

	need := sum(expected)
	remaining := strings.Count(input, "#")
	possible := strings.Count(input, "?")

	// we don't want more springs, so we have a valid arrangement
	if need == 0 && remaining == 0 {
		fmt.Printf("%s IS VALID!\n", input)
		seenSprings[input] = true
		return
	}

	// we need more springs than is possible in our remaining space, bail out
	if need > remaining+possible {
		fmt.Printf("%s lacks springs, bailing.\n", input)
		seenSprings[input] = false
		return
	}

	// we need less springs than is possible, bail out
	if need < remaining {
		fmt.Printf("%s has too many springs, bailing.\n", input)
		seenSprings[input] = false
		return
	}

	// fmt.Printf("%s %v\n", input, expected)
	// try to place first value
	fmt.Printf("%s - place %d ", input, expected[0])
	result, ok := place(input, expected[0])
	if ok {
		fmt.Printf("succeeded\n")
		// if we placed it, continue down this new path
		findArrangements(result, expected[1:])
	} else {
		fmt.Printf("failed, moving to next variant.\n")
	}

	// either way, also continue down the other path; flip the next '?'
	if possible > 0 {
		variant := strings.Replace(input, "?", ".", 1)
		findArrangements(variant, expected)
	}

	return
}

// place tries to place a group of springs of size val as far left as possible in the input spring.
// Converts placed springs to '*' characters.
// Placement is greedy, so the left-most ? will be used for a #.
// If placement is not possible at the left-most ?, returns false.
func place(input string, expect int) (string, bool) {
	var result string
	placed := 0
	i := 0

	for _, char := range input {
		i++

		switch char {
		case '.':
			// we had already started a group, now it's ended too early
			if placed > 0 {
				return "", false
			}
			result += "."
		case '*':
			// we had already started a group, now it's ended too early
			if placed > 0 {
				return "", false
			}
			// skip over already-placed springs
			result += "*"
		case '#', '?':
			// start or continue a group
			placed++
			result += "*"
		}

		// we placed the expected group, stop looking
		if placed == expect {
			break
		}
	}

	// we reached the end without placing the expected group
	if placed != expect {
		return "", false
	}

	// check for or add gap after this group, unless we're at the end of the input
	if i < len(input) {
		switch input[i] {
		case '.':
			// we already have a gap, do nothing
		case '#', '*':
			// we hit a spring when we needed a gap, we failed
			return "", false
		case '?':
			// convert to a gap and advance the input
			result += "."
			i++
		}
	}

	// add the rest of the string
	result += input[i:]

	return result, true
}

func sum(x []int) int {
	total := 0
	for _, i := range x {
		total += i
	}
	return total
}
