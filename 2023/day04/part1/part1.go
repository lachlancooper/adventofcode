package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(sumWins(scanner))
}

func sumWins(scanner *bufio.Scanner) int {
	var total int

	for scanner.Scan() {
		var score int
		line := strings.Split(scanner.Text(), ":")
		parts := strings.Split(line[1], "|")

		// collect winning numbers
		winners := []string{}
		for _, winner := range strings.Split(parts[0], " ") {
			if winner == "" {
				continue
			}
			winners = append(winners, winner)
		}

		for _, num := range strings.Split(parts[1], " ") {
			if num == "" {
				continue
			}
			if slices.Contains(winners, num) {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		total += score
	}

	return total
}
