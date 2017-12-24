// http://adventofcode.com/2017/day/23
// part 2 - loop
package main

import (
	"fmt"
)

func main() {
	var h int

	for b := 109300; b <= 126300; b += 17 {
	loop2:
		for d := 2; d <= b; d++ {
			for e := 2; e <= b/d; e++ {
				if d*e == b {
					h++
					break loop2
				}
			}
		}
	}

	fmt.Println(h)
}
