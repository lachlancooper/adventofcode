// http://adventofcode.com/2017/day/21
// part 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// fixed starting image
var image = []string{".#.", "..#", "###"}
var rules = make(map[string]string)

// print given string slice as image
func printimage(i []string) {
	for _, line := range i {
		fmt.Println(line)
	}
	fmt.Println()
}

// convert square to string representation
func flatten(square []string) string {
	return strings.Join(square, "/")
}

// convert string representation to square
func unflatten(flat string) []string {
	return strings.Split(flat, "/")
}

// rotate square once
// direction doesn't matter but must be consistent
// (we do it clockwise)
func rotate(square []string) []string {
	newsquare := make([]string, 0)
	for i := range square {
		newline := ""
		for j := range square {
			newline += string(square[len(square)-j-1][i])
		}
		newsquare = append(newsquare, newline)
	}
	return newsquare
}

// flip square once
// direction doesn't matter
// we do it vertically (by line)
func flip(square []string) []string {
	newsquare := make([]string, 0)
	for i := range square {
		newsquare = append(newsquare, square[len(square)-i-1])
	}
	return newsquare
}

// generate permutations of square to find matching rule
func match(square []string) []string {
	// flip x2
	for f := 0; f < 2; f++ {
		// rotate x4
		for r := 0; r < 4; r++ {
			// found a match?
			if val, ok := rules[flatten(square)]; ok {
				return unflatten(val)
			}
			square = rotate(square)
		}
		square = flip(square)
	}

	return []string{}
}

// apply rules once
func step() {
	var squaresize int

	// find dimensions
	if len(image)%2 == 0 {
		squaresize = 2
	} else {
		squaresize = 3
	}

	// create new (larger) image with placeholder (blank) lines
	newimagesize := len(image) / squaresize * (squaresize + 1)
	newimage := make([]string, newimagesize)

	// range over squares
	for sy := 0; sy < len(image)/squaresize; sy++ {
		for sx := 0; sx < len(image)/squaresize; sx++ {

			// cut out each square
			square := make([]string, 0)
			for y := 0; y < squaresize; y++ {
				line := image[sy*squaresize+y][sx*squaresize : (sx+1)*squaresize]
				square = append(square, line)
			}

			// match and insert into new image
			newsquare := match(square)
			for y := range newsquare {
				newimage[y+sy*(squaresize+1)] += newsquare[y]
			}
		}
	}
	image = newimage
}

// apply rules for r iterations
func run(r int) {
	for i := 0; i < r; i++ {
		step()
	}
}

// count the active pixels in i
func countactive(i []string) (sum int) {
	for _, s := range i {
		sum += strings.Count(s, "#")
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		l := strings.Fields(scanner.Text())
		rules[l[0]] = l[2]
	}

	run(18)

	fmt.Println(countactive(image))
}
