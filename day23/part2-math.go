// http://adventofcode.com/2017/day/23
// part 2 - MATH!
package main

import (
	"fmt"
	"math/big"
)

func main() {
	composite := 0
	start := big.NewInt(109300)
	end := big.NewInt(126300)
	increment := big.NewInt(17)

	for b := start; b.Cmp(end) <= 0; b.Add(b, increment) {
		if !b.ProbablyPrime(20) {
			composite++
		}
	}

	fmt.Println(composite)
}
