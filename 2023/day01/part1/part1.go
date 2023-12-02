package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(calibrate(scanner))
}

// calibrate calculates the overall calibration value for a document.
func calibrate(scanner *bufio.Scanner) int {
	var total int

	for scanner.Scan() {
		line := scanner.Text()
		total += calibrateLine(line)
	}

	return total
}

// calibrateLine combines the first and last digit found on a line to form a single two-digit number.
func calibrateLine(s string) int {
	var first, last int

	for _, c := range s {
		i, err := strconv.Atoi(string(c))
		if err != nil {
			continue
		}

		if first == 0 {
			first = i
		}
		last = i
	}

	return first*10 + last
}
