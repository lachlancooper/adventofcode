// http://adventofcode.com/2017/day/3
// part 1
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// abs returns the absolute value of i
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// square returns i raised to the power of 2
func square(i int) int {
	return i * i
}

// manhattan returns the distance from s
// to square 1 of a spiral grid:
// 17  16  15  14  13
// 18   5   4   3  12
// 19   6   1   2  11
// 20   7   8   9  10
// 21  22  23---> ...
func manhattan(s int) int {
	if s == 1 {
		return 0
	}

	// layer is the chessboard distance of s from square 1
	// 2   2   2   2   2
	// 2   1   1   1   2
	// 2   1   0   1   2
	// 2   1   1   1   2
	// 2   2   2   2   2
	layer := int(math.Ceil(math.Sqrt(float64(s)))) / 2

	// prevlayerend is the last square of the previous layer
	prevlayerend := square(layer*2 - 1)

	// edgelen is number of squares on each edge of current layer
	edgelen := layer * 2

	// nearestaxis is the closest square in the current layer
	// collinear with square 1, according to the cardinal axes
	nearestaxis := prevlayerend + (s-prevlayerend)/edgelen*edgelen + edgelen/2

	// first move to nearest axis
	dist := abs(s - nearestaxis)
	// then move inward to square 1
	dist += layer

	return dist
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		square, _ := strconv.Atoi(scanner.Text())
		fmt.Println(manhattan(square))
	}
}
