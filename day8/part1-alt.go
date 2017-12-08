// http://adventofcode.com/2017/day/8
// part 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Knetic/govaluate"
)

// memory contains any number of named registers and their values
var memory = make(map[string]int)

// process instruction i
func processInstruction(i []string) {
	// decode instruction
	register := i[0]
	operation := i[1]
	value, _ := strconv.Atoi(i[2])

	// test condition
	condition := fmt.Sprintf("%v %v %v", memory[i[4]], i[5], i[6])
	expression, _ := govaluate.NewEvaluableExpression(condition)
	result, _ := expression.Evaluate(nil)
	if !result.(bool) {
		return
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
