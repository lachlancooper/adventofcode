// http://adventofcode.com/2017/day/15
// part 2
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
	for {
		a = (a * 16807) % 2147483647
		if (a % 4) == 0 {
			break
		}
	}
	for {
		b = (b * 48271) % 2147483647
		if (b % 8) == 0 {
			break
		}
	}

	return a, b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	gena, _ := strconv.Atoi(strings.Fields(scanner.Text())[4])
	scanner.Scan()
	genb, _ := strconv.Atoi(strings.Fields(scanner.Text())[4])

	count := 0
	for i := 0; i < 5*1000*1000; i++ {
		gena, genb = generate(gena, genb)
		if 0xFFFF&gena == 0xFFFF&genb {
			count++
		}
	}
	fmt.Println(count)
}
