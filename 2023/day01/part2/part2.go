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
// Digits spelled out in english letters also count as valid "digits".
func calibrateLine(s string) int {
	var first, last int

	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(scanDigits)

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
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

// scanDigits is a split function for a Scanner that returns each byte as a digit.
func scanDigits(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	switch {
	// For letter digits we advance the input one less than the length of the word.
	// This allows overlaps in words (though this is never possible for "four", "six", or "seven").
	// e.g. "six9jhnloneightf" has calibration value 68, not 61.
	case strings.HasPrefix(string(data), "one"):
		return 2, []byte("1"), nil
	case strings.HasPrefix(string(data), "two"):
		return 2, []byte("2"), nil
	case strings.HasPrefix(string(data), "three"):
		return 4, []byte("3"), nil
	case strings.HasPrefix(string(data), "four"):
		return 3, []byte("4"), nil
	case strings.HasPrefix(string(data), "five"):
		return 3, []byte("5"), nil
	case strings.HasPrefix(string(data), "six"):
		return 2, []byte("6"), nil
	case strings.HasPrefix(string(data), "seven"):
		return 4, []byte("7"), nil
	case strings.HasPrefix(string(data), "eight"):
		return 4, []byte("8"), nil
	case strings.HasPrefix(string(data), "nine"):
		return 3, []byte("9"), nil
	default:
		return 1, data[0:1], nil
	}
}
