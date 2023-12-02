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

	fmt.Println(sumPowers(scanner))
}

// sumPowers calculates the sum of powers of the minimum set of cubes across all games.
func sumPowers(scanner *bufio.Scanner) int {
	var total int

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		total += power(line[1])
	}

	return total
}

// power calculates the power of the minimum set of cubes needed for a game.
func power(game string) int {
	var red, green, blue int

	iterations := strings.Split(game, ";")

	for _, iteration := range iterations {
		cubes := strings.Split(iteration, ",")
		for _, cube := range cubes {
			parts := strings.Split(cube, " ")
			count, _ := strconv.Atoi(parts[1])

			switch parts[2] {
			case "red":
				if count > red {
					red = count
				}
			case "green":
				if count > green {
					green = count
				}
			case "blue":
				if count > blue {
					blue = count
				}
			default:
				panic("no colour detected")
			}
		}
	}

	return red * green * blue
}
