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

	stepstocentreofedge := int(math.Abs(float64(s - centreofedge)))

	// move to "centre" of current edge
	// then directly to centre of grid (== layer)
	return stepstocentreofedge + layer
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		square, _ := strconv.Atoi(scanner.Text())
		fmt.Println(manhattan(square))
	}
}
