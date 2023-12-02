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

	fmt.Println(sumIDs(scanner))
}

// sumIDs calculates the sum of IDs of games that are possible.
func sumIDs(scanner *bufio.Scanner) int {
	var i, total int

	for scanner.Scan() {
		i++
		line := strings.Split(scanner.Text(), ":")
		if possible(line[1]) {
			total += i
		}
	}

	return total
}

// possible determines whether a game is possible given the cubes revealed.
// We assume the bag contains only 12 red, 13 green, and 14 blue cubes.
func possible(game string) bool {
	iterations := strings.Split(game, ";")

	for _, iteration := range iterations {
		cubes := strings.Split(iteration, ",")
		for _, cube := range cubes {
			parts := strings.Split(cube, " ")
			count, _ := strconv.Atoi(parts[1])

			switch parts[2] {
			case "red":
				if count > 12 {
					return false
				}
			case "green":
				if count > 13 {
					return false
				}
			case "blue":
				if count > 14 {
					return false
				}
			default:
				panic("no colour detected")
			}
		}
	}

	return true
}
