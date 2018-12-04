// http://adventofcode.com/2017/day/17
// part 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var cycle int
var buffer = []int{0}

func main() {
	// read step count
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cycle, _ = strconv.Atoi(scanner.Text())

	currentpos := 0
	for i := 0; i < 2017; i++ {
		insertpos := (currentpos+cycle)%len(buffer) + 1

		// insert an entry at insertpos by:
		// - extending buffer by one
		// - copying all elements after insertpos forward by one
		// - setting the element at insertpos to the desired value
		buffer = append(buffer, 0)
		copy(buffer[insertpos+1:], buffer[insertpos:])
		buffer[insertpos] = i + 1

		currentpos = insertpos
	}
	fmt.Println(buffer[currentpos+1])
}
