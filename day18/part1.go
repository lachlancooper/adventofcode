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

	for pc := 0; pc < len(ins); pc++ {
		// fetch
		i := ins[pc]
		op := i[0]
		dst := i[1]

		// decode operand values
		val := make([]int, 0)
		for _, c := range i[1:] {
			if v, err := strconv.Atoi(c); err == nil {
				val = append(val, v)
			} else {
				val = append(val, reg[c])
			}
		}

		// execute
		switch op {
		case "snd":
			sound = val[0]
		case "set":
			reg[dst] = val[1]
		case "add":
			reg[dst] += val[1]
		case "mul":
			reg[dst] *= val[1]
		case "mod":
			reg[dst] %= val[1]
		case "rcv":
			if reg[dst] != 0 {
				fmt.Println(sound)
				return
			}
		case "jgz":
			if val[0] > 0 {
				pc += val[1] - 1
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
