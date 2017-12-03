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

// sqr returns i raised to the power of 2
func sqr(i int) int {
	return i * i
}

// manhattan returns the distance from the given square
// to the centre square of a spiral pattern:
//
// 37  36  35  34  33  32  31
// 38  17  16  15  14  13  30
// 39  18   5   4   3  12  29
// 40  19   6   1   2  11  28
// 41  20   7   8   9  10  27
// 42  21  22  23  24  25  26
// 43  44  45  46  47  48  49
func manhattan(s int) int {
	if s == 1 {
		return 0
	}

	// layer is absolute distance from grid centre
	//  2- 9: 1
	// 10-25: 2
	// 26-49: 3
	// 50-81: 4
	layer := int(math.Ceil(math.Sqrt(float64(s)))) / 2

	// layerstart is number of first square of current layer
	//  2- 9: 2
	// 10-25:10
	// 26-49:26
	// 50-81:50
	layerstart := sqr(layer*2-1+layer%2) + 1

	// edgelen is number of squares on each edge of current layer
	edgelen := layer * 2

	edgecentre := (layerstart - 1) + (s-layerstart)/edgelen*edgelen + edgelen/2

	// first, move along edge to centre square
	//dist := abs((s-layerstart)%(edgelen) + 1 - layer)
	dist := abs(s - edgecentre)
	fmt.Println("Square", s, "is in layer", layer, "which started at", layerstart, ". It takes", dist, "steps to centre of edge (", (s-layerstart)/edgelen, edgecentre, "), then", layer, "steps to grid centre.")

	// then move inward to grid centre, 1 step per layer
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
