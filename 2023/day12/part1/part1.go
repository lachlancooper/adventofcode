package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

		// find the spring map
		springs := fields[0]

		// find the expected values
		expected := []int{}
		vals := strings.Split(fields[1], ",")
		for _, v := range vals {
			count, _ := strconv.Atoi(v)
			expected = append(expected, count)
		}

		total += arrangements(springs, expected)
	}

	return total
}

// arrangements counts the number of valid arrangements of springs in the given input.
func arrangements(input string, expected []int) int {
	total := 0

	for i := 0; i < pow(2, strings.Count(input, "?")); i++ {
		testInput := replace(input, i)
		if valid(testInput, expected) {
			total++
		}
	}

	return total
}

// replace replaces ? with # in a pattern matching the variant.
// This helps to generate every possible combination of springs, e.g.
//
//	"???.###"" has 2 ^ 3 = 8 combos
//	replace("???.###"", 0) == "###.###"
//	replace("???.###"", 1) == "##..###"
//	replace("???.###"", 2) == "#.#.###"
//	replace("???.###"", 3) == "#...###"
//	replace("???.###"", 4) == ".##.###"
//	replace("???.###"", 5) == ".#..###"
//	replace("???.###"", 6) == "..#.###"
//	replace("???.###"", 7) == "....###"
func replace(input string, variant int) string {
	result := input

	// iterate over all ? in input
	// find next ? in result
	// if corresponding binary digit of variant is 0, set it to #
	// otherwise, set it to ?
	for i := 0; i < strings.Count(input, "?"); i++ {
		nextOffset := strings.Index(result, "?")
		var replaceChar string
		if variant&pow(2, i) == 0 {
			replaceChar = "#"
		} else {
			replaceChar = "."
		}
		result = result[:nextOffset] + replaceChar + result[nextOffset+1:]
	}

	return result
}

// valid validates whether a given input satisfies the expected group count.
func valid(input string, expected []int) bool {
	groups := []int{}
	for _, v := range strings.Split(input, ".") {
		if v == "" {
			continue
		}
		groups = append(groups, len(v))
	}

	result := slices.Equal(groups, expected)

	// fmt.Printf("Checked %s: %v\n", input, result)
	return result
}

func pow(x, y int) int {
	if y == 0 {
		return 1
	}

	result := x
	for ; y > 1; y-- {
		result *= x
	}

	return result
}
