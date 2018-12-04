// http://adventofcode.com/2018/day/1
// part 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	freq := 0

	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		freq += val
	}
	fmt.Println(freq)
}
