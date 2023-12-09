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

	fmt.Println(countTotalsReverse(scanner))
}

func countTotalsReverse(scanner *bufio.Scanner) int {
	total := 0

	for scanner.Scan() {
		history := []int{}
		for _, num := range strings.Split(scanner.Text(), " ") {
			val, _ := strconv.Atoi(num)
			history = append(history, val)
		}
		history = reverse(history)
		total += findNext(history)
	}

	return total
}

func findNext(history []int) int {
	sequences := [][]int{}
	layers := 0

	// history is the first layer
	sequences = append(sequences, history)

	// build layers
	for {
		nextLayer := []int{}
		for i := 1; i < len(sequences[layers]); i++ {
			diff := sequences[layers][i] - sequences[layers][i-1]
			nextLayer = append(nextLayer, diff)
		}

		sequences = append(sequences, nextLayer)
		layers++

		if allZero(sequences[layers]) {
			break
		}
	}

	// go back up through layers
	for ; layers > 0; layers-- {
		thisLayer := sequences[layers]
		upperLayer := sequences[layers-1]
		nextVal := thisLayer[len(thisLayer)-1] + upperLayer[len(upperLayer)-1]
		sequences[layers-1] = append(sequences[layers-1], nextVal)
	}

	firstLayer := sequences[0]
	return firstLayer[len(firstLayer)-1]
}

func allZero(history []int) bool {
	for _, val := range history {
		if val != 0 {
			return false
		}
	}

	return true
}

func reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
