package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	left, right string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(countSteps(scanner))
}

func countSteps(scanner *bufio.Scanner) int {
	// read instructions
	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()

	// read nodes
	nodes := make(map[string]node)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		nodes[line[0]] = node{strings.Trim(line[2], "(,"), strings.Trim(line[3], ")")}
	}

	// follow instructions, starting at AAA, counting steps to reach ZZZ
	steps := 0
	for current := "AAA"; current != "ZZZ"; steps++ {
		if instructions[steps%len(instructions)] == 'L' {
			current = nodes[current].left
		} else {
			current = nodes[current].right
		}
	}

	return steps
}
