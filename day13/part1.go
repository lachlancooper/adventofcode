// http://adventofcode.com/2017/day/13
// part 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	severity := 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ": ")
		layer, _ := strconv.Atoi(line[0])
		depth, _ := strconv.Atoi(line[1])

		if layer%(depth*2-2) == 0 {
			severity += layer * depth
		}
	}
	fmt.Println(severity)
}
