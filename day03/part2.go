// http://adventofcode.com/2017/day/3
// part 2
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// squares holds the values of all squares calculated so far.
// neighboursum will write to it whenever it calculates
// a new value
// squares 1 and 2 are special and precalculated
var squares = []int{1, 1}

// neighboursum returns the sum of values of all squares
// neighbouring (including diagonals) the given square
// squares are numbered outward in a spiral pattern
//
// 37  36  35  34  33  32  31
// 38  17  16  15  14  13  30
// 39  18   5   4   3  12  29
// 40  19   6   1   2  11  28
// 41  20   7   8   9  10  27
// 42  21  22  23  24  25  26
// 43  44  45  46  47  48  49
//
//                    5022 2450
// 147  142  133  122   59 2391
// 304    5    4    2   57 2275
// 330   10    1    1   54 2105
// 351   11   23   25   26 1968
// 362  747  806  880  931  957
func neighboursum(s int) int {
	if s <= len(squares) {
		return squares[s-1]
	}

	sum := 0
	ceilsqrt := int(math.Ceil(math.Sqrt(float64(s))))
	nextoddpowroot := ceilsqrt + (1 - ceilsqrt%2)
	prevoddpowroot := nextoddpowroot - 2
	layer := nextoddpowroot / 2
	sizeofedge := layer * 2
	nextoddpow := int(math.Pow(float64(nextoddpowroot), 2))
	prevoddpow := int(math.Pow(float64(prevoddpowroot), 2))
	edge := (s - prevoddpow - 1) / sizeofedge
	prevcorner := prevoddpow + sizeofedge*edge
	nextcorner := prevcorner + sizeofedge
	startoflayer := prevoddpow + 1
	endoflayer := nextoddpow

	// we always neighbour the previous square
	sum += neighboursum(s - 1)

	// all corners neighbour their inner corner
	if s == nextcorner {
		innercorner := s - ((layer-1)*8 + (edge+1)*2)
		sum += neighboursum(innercorner)
	} else {
		// all squares that start layers only neighbour the previous layer start
		if s == startoflayer {
			sum += neighboursum(s - ((layer - 1) * 8))
		} else {
			// all non-corners neighbour their inner square
			innersquare := s - ((layer-1)*8 + 1 + edge*2)
			sum += neighboursum(innersquare)

			if s == prevcorner+1 {
				// all squares next to the previous corner also neighbour around the corner
				sum += neighboursum(prevcorner - 1)
			} else if s == startoflayer+1 {
				// all squares next to start of layer also neighbour around the corner
				sum += neighboursum(startoflayer - 1)
			} else {
				// all other squares neighbour the inner square - 1
				sum += neighboursum(innersquare - 1)
			}
			if s == nextcorner-1 {
			} else {
				// all squares not before a corner also neighbour the innner square + 1
				sum += neighboursum(innersquare + 1)
			}
		}
	}
	// all layer ends and just prior also neighbour layer start
	if s == endoflayer || s == endoflayer-1 {
		sum += neighboursum(startoflayer)
	}

	squares = append(squares, sum)
	return sum
}

// nextlarger returns the next neighboursum value
// larger than the given value
func nextlarger(val int) (next int) {
	for i := 2; next <= val; i++ {
		next = neighboursum(i)
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		fmt.Println(nextlarger(value))
	}
}
