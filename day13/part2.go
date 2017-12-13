// http://adventofcode.com/2017/day/13
// part 2
//
// caught iff you meet a scanner on its cycle, i.e.
// your arrival (layer + delay) % depth*2-2 == 0
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var firewall = make(map[int]int)

// caught calculates whether or not a packet traversing
// firewall starting at picosecond s is caught
func caught(s int) bool {
	for layer, depth := range firewall {
		if (layer+s)%(depth*2-2) == 0 {
			return true
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ": ")
		layer, _ := strconv.Atoi(line[0])
		depth, _ := strconv.Atoi(line[1])
		firewall[layer] = depth
	}
	for i := 0; ; i++ {
		if !caught(i) {
			fmt.Println(i)
			break
		}
	}
}
