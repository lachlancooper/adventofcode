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

type box struct {
	lenses []lens
}

type lens struct {
	label string
	focal int
}

func solve(scanner *bufio.Scanner) int {
	boxes := [256]box{}
	scanner.Scan()

	for _, step := range strings.Split(scanner.Text(), ",") {
		if parts := strings.Split(step, "="); len(parts) > 1 {
			// add or replace lens
			label := parts[0]
			boxNum := hash(label)
			focal, _ := strconv.Atoi(parts[1])
			boxes[boxNum] = boxes[boxNum].addUpdateLens(lens{label: label, focal: focal})
		} else if parts := strings.Split(step, "-"); len(parts) > 1 {
			// remove lens
			label := parts[0]
			boxNum := hash(label)

			boxes[boxNum] = boxes[boxNum].removeLens(lens{label: label})
		}
	}

	total := 0
	for i, box := range boxes {
		for j, lens := range box.lenses {
			total += (1 + i) * (1 + j) * lens.focal
		}
	}

	return total
}

func hash(s string) int {
	result := 0
	for _, char := range s {
		result += int(char)
		result *= 17
		result %= 256
	}
	return result
}

// addUpdateLens adds the lens with the specified label to the box, or updates its focal length.
func (b box) addUpdateLens(l lens) box {
	// search for lens
	for i, lens := range b.lenses {
		if lens.label != l.label {
			continue
		}
		b.lenses[i] = l
		return b
	}

	b.lenses = append(b.lenses, l)

	return b
}

// removeLens removes the lens with the specified label from the box, if present.
func (b box) removeLens(l lens) box {
	// search for lens
	for i, lens := range b.lenses {
		if lens.label != l.label {
			continue
		}
		b.lenses = append(b.lenses[:i], b.lenses[i+1:]...)
		return b
	}

	return b
}
