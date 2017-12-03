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

// manhattan returns the distance from the given square
// to the centre of a spiral pattern:
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

	//	fmt.Println("s =", s)

	ceilsqrt := int(math.Ceil(math.Sqrt(float64(s))))
	nextoddpowroot := ceilsqrt + (1 - ceilsqrt%2)
	//	fmt.Println("nextoddpowroot =", nextoddpowroot)

	prevoddpowroot := nextoddpowroot - 2
	layer := nextoddpowroot / 2
	//	fmt.Println("layer =", layer)

	sizeofedge := layer * 2
	//	fmt.Println("sizeofedge =", sizeofedge)

	prevoddpow := int(math.Pow(float64(prevoddpowroot), 2))
	edge := (s - prevoddpow - 1) / sizeofedge
	//	fmt.Println("edge =", edge)

	centreofedge := prevoddpow + edge*sizeofedge + sizeofedge/2
	//	fmt.Println("centreofedge =", centreofedge)

	//	fmt.Println()

	switch {
	case s > centreofedge:
		return 1 + manhattan(s-1)
	case s < centreofedge:
		return 1 + manhattan(s+1)
	// case s == centreofedge:
	default:
		return layer
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		square, _ := strconv.Atoi(scanner.Text())
		fmt.Println(manhattan(square))
	}
}

/*
layer 0: 1
layer 1: 2-9
	edge 0: 2-3
	edge 1: 4-5
	edge 2: 6-7
	edge 3: 8-9
layer 2: 10-25
	edge 0: 10-13
	edge 1: 14-17
	edge 2: 18-21
	edge 3: 22-25
layer 3: 26-49
	size 24
layer 4: 50-81
	size 32

layer number is
	sqrt(next odd power)-2

size of each layer is
	8 * (layer number)

cases for position in layer
bottom right:
right: sub
top right:
top:
top left:
left:
bottom left:
bottom:
*/
