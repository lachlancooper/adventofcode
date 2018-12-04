// http://adventofcode.com/2018/day/1
// part 2
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
	changes := make([]int, 0)
	seen := make(map[int]bool)

	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		changes = append(changes, val)
	}
	for true {
		for _, v := range changes {
			freq += v
			if seen[freq] {
				fmt.Println(freq)
				return
			}
			seen[freq] = true
		}
	}
}
