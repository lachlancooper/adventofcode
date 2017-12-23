// http://adventofcode.com/2017/day/18
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
	var sound int

outer:
	for pc := 0; pc < len(ins); pc++ {
		// fetch
		i := ins[pc]

		// collect operand
		var value int
		if len(i) == 3 {
			if v, err := strconv.Atoi(i[2]); err == nil {
				value = v
			} else {
				value = reg[i[2]]
			}
		}

		// execute
		switch i[0] {
		case "snd":
			sound = reg[i[1]]
		case "set":
			reg[i[1]] = value
		case "add":
			reg[i[1]] += value
		case "mul":
			reg[i[1]] *= value
		case "mod":
			reg[i[1]] %= value
		case "rcv":
			if reg[i[1]] != 0 {
				fmt.Println(sound)
				break outer
			}
		case "jgz":
			if reg[i[1]] > 0 {
				pc += value - 1
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ins = append(ins, strings.Fields(scanner.Text()))
	}

	run()
}
