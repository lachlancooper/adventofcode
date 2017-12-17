// http://adventofcode.com/2017/day/17
// part 2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var cycle int

func main() {
	// read step count
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cycle, _ = strconv.Atoi(scanner.Text())

	currentpos := 0
	var answer int
	// element 0 never moves from position 0, since
	// we always append *after* whatever position
	// we're considering
	//
	// thus, we only need to look out for the times that
	// insert position is 1; the rest of the buffer doesn't
	// need to be simulated at all
	//
	// the value it inserts at that time is cycle count + 1
	for i := 0; i < 50*1000*1000; i++ {
		insertpos := (currentpos+cycle)%(i+1) + 1
		currentpos = insertpos

		if insertpos == 1 {
			answer = i + 1
		}
	}
	fmt.Println(answer)
}
