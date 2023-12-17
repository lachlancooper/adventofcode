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

		total += arrangements(springs, expected)
	}

	return total
}

// new strategy:
// go through ??? finding and adding each expected group, in order
//
// ???.### 1,1,3
// start with first 1
// #.?.### place it at the first available ?, then next must be .
// move to next 1
// #.#.### place it at the first available ?
// we have no more ?, so test this out, it's valid
//
// ???? 1,1
// #.?? start with first 1, place it at first available ?, with following .
// #.#. move to next 1, place it at next available ?, with following .
// test, it's valid, +1
// #..# move to previous step and try other available ?, it's valid, +1
// .#.? move to first step, place at next available ?, with .
// .#.# place second 1, it's valid, +1
// ..#. move to previous step

// arrangements counts the number of valid arrangements of springs in the given input.
func arrangements(input string, expected []int) int {
	total := 0

	// fmt.Printf("%100s %v\n", input, expected)

	// we don't want more springs,
	if len(expected) == 0 && !strings.Contains(input, "#") {
		// fmt.Printf("%s was valid!\n", input)
		return 1
	}

	// advance input char by char
	for i := 0; i < len(input); i++ {
		variant := input[i:]

		need := sum(expected)
		have := strings.Count(variant, "#")
		// we need more springs than is possible in our remaining space, bail out
		if need > have+strings.Count(variant, "?") {
			break
		}

		// we need less springs than is possible, bail out
		if need < have {
			break
		}

		// skip over initially empty cells
		if variant[0] == '.' {
			continue
		}

		// try to place first value
		val := expected[0]
		result, ok := place(variant, val)
		if !ok {
			// we can't place the val here, this variant has failed, move to the next
			continue
		}

		// if it succeeds, recurse with next expected val and trimmed variant string
		// trim 'val' springs off the next variant
		newInput := result[val:]
		// remove 'val' from expected list
		newExpected := expected[1:]

		// recurse, with new variant and expected
		total += arrangements(newInput, newExpected)
	}

	return total
}

// place tries to place a group of springs of size val as far left as possible in the input spring.
// If this is not possible, returns false.
// Placement is greedy, so the left-most ? will be used for a #.
// If this fails, we'll move on to the next spring at a higher level anyway.
func place(input string, expect int) (string, bool) {
	var result string
	foundGroup := 0
	i := 0

	for _, char := range input {
		i++

		switch char {
		case '.':
			// we had already started a group, now it's ended too early
			if foundGroup > 0 {
				return "", false
			}
			result += "."
		case '#', '?':
			// start or continue a group
			foundGroup++
			result += "#"
		}

		// we placed the expected group, stop looking
		if foundGroup == expect {
			break
		}
	}

	// we reached the end without placing the expected group
	if foundGroup != expect {
		return "", false
	}

	// check for or add gap after this group, unless we're at the end of the input
	if i < len(input) {
		switch input[i] {
		case '.':
			// we already have a gap, do nothing
		case '#':
			// we hit a spring when we wanted a gap, we failed
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
