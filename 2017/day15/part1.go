// http://adventofcode.com/2017/day/15
// part 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// generate one cycle
func generate(a, b int) (int, int) {
	a = (a * 16807) % 2147483647
	b = (b * 48271) % 2147483647
	return a, b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	gena, _ := strconv.Atoi(strings.Fields(scanner.Text())[4])
	scanner.Scan()
	genb, _ := strconv.Atoi(strings.Fields(scanner.Text())[4])

	count := 0
	for i := 0; i < 40*1000*1000; i++ {
		gena, genb = generate(gena, genb)
		if 0xFFFF&gena == 0xFFFF&genb {
			count++
		}
	}
	fmt.Println(count)
}
