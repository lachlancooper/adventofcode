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

	fmt.Println(sumCards(scanner))
}

func sumCards(scanner *bufio.Scanner) int {
	var game, total int
	copies := make(map[int]int)

	// loop over all games
	for scanner.Scan() {
		var score int

		game++
		copies[game]++

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

		// collect wins on this card
		for _, num := range strings.Split(parts[1], " ") {
			if num == "" {
				continue
			}
			if slices.Contains(winners, num) {
				score++
			}
		}

		// add copies of future cards based on score of this game
		for j := 1; j <= score; j++ {
			copies[game+j] += copies[game]
		}

		total += copies[game]
	}

	return total
}
