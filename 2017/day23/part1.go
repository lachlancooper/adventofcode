// http://adventofcode.com/2017/day/23
// part 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// reg holds any number of named registers and their values
var reg = make(map[string]int)

// ins holds a list of instructions
var ins = make([][]string, 0)

// run executes instructions sequentially
func run() {
	mulcount := 0

	for pc := 0; pc < len(ins); pc++ {
		// fetch
		i := ins[pc]

		// collect operand
		var value int
		if v, err := strconv.Atoi(i[2]); err == nil {
			value = v
		} else {
			value = reg[i[2]]
		}

		// execute
		switch i[0] {
		case "set":
			reg[i[1]] = value
		case "sub":
			reg[i[1]] -= value
		case "mul":
			reg[i[1]] *= value
			mulcount++
		case "jnz":
			// collect condition
			var cond int
			if v, err := strconv.Atoi(i[1]); err == nil {
				cond = v
			} else {
				cond = reg[i[1]]
			}

			if cond != 0 {
				pc += value - 1
			}
		}
	}
	fmt.Println(mulcount)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ins = append(ins, strings.Fields(scanner.Text()))
	}

	run()
}
