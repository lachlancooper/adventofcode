// http://adventofcode.com/2018/day/2
// part 1
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	twoLetterIds := 0
	threeLetterIds := 0

	for scanner.Scan() {
		count := make(map[rune]int)
		boxid := scanner.Text()
		hasTwo := 0
		hasThree := 0

		// tally occurences of each rune
		for _, r := range boxid {
			count[r] += 1
		}

		// check for exactly two or three of any rune
		for _, f := range count {
			if f == 2 {
				hasTwo = 1
			} else if f == 3 {
				hasThree = 1
			}
		}

		twoLetterIds += hasTwo
		threeLetterIds += hasThree
	}

	fmt.Println(twoLetterIds * threeLetterIds)
}
