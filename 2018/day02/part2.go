// http://adventofcode.com/2018/day/2
// part 2
package main

import (
	"bufio"
	"fmt"
	"os"
)

// compare reports whether two strings differ by exactly one rune
func compare(box1 string, box2 string) (oneDiff bool) {
	for i := 0; i < len(box1); i++ {
		if box1[i] != box2[i] {
			if oneDiff {
				// we've already seen one diff, so another eliminates this pair
				return false
			}
			oneDiff = true
		}
	}
	return oneDiff
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	boxes := make([]string, 0)

	for scanner.Scan() {
		boxid := scanner.Text()

		for _, b := range boxes {
			if compare(boxid, b) {
				fmt.Println(boxid)
				fmt.Println(b)
				return
			}
		}

		boxes = append(boxes, boxid)
	}
}
