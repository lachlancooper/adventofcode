// http://adventofcode.com/2017/day/8
// part 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// memory contains any number of named registers and their values
var memory = make(map[string]int)

// process instruction i
func processInstruction(i []string) {
	// decode instruction
	register := i[0]
	operation := i[1]
	value, _ := strconv.Atoi(i[2])
	operand1 := memory[i[4]]
	condition := i[5]
	operand2, _ := strconv.Atoi(i[6])

	// test condition
	switch condition {
	case ">":
		if !(operand1 > operand2) {
			return
		}
	case ">=":
		if !(operand1 >= operand2) {
			return
		}
	case "<":
		if !(operand1 < operand2) {
			return
		}
	case "<=":
		if !(operand1 <= operand2) {
			return
		}
	case "==":
		if !(operand1 == operand2) {
			return
		}
	case "!=":
		if !(operand1 != operand2) {
			return
		}
	}

	// condition passed, perform operation
	switch operation {
	case "inc":
		memory[register] += value
	case "dec":
		memory[register] -= value
	}
}

// return the largest value in any register
func largestValue() (max int) {
	for _, val := range memory {
		if val > max {
			max = val
		}
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		processInstruction(strings.Fields(scanner.Text()))
	}

	fmt.Println(largestValue())
}
