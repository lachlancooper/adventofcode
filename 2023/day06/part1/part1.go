package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(countWins(scanner))
}

func countWins(scanner *bufio.Scanner) int {
	return 0
}

// wins counts the number of ways to win, given an amount of time and a current distance record.
func wins(time, distance int) int {
	// maxDistance is the maximum distance achievable in the given time.
	maxDistance := time / 2 * (time/2 + time%2)

	difference := maxDistance - distance

	steps := countSteps(difference)

	// add one extra win for even times (there's only one combination that gives "best" distance)
	return steps*2 + (1 - time%2)
}

// countSteps counts how many multiples of two are required to reach the specified value.
func countSteps(val int) int {
	// v = n^2 - n
	// n = (1+(1-4v)^0.5) / 2
	return int(math.Round(math.Sqrt(float64(val))))
}

func wins2(time, distance int) int {
	return int(math.Round(math.Sqrt(float64(time/2*(time/2+time%2)-distance))))*2 + (1 - time%2)
}
