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

	// read nodes, and find starting nodes
	nodes := make(map[string]node)
	current := []string{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		name := line[0]
		nodes[name] = node{strings.Trim(line[2], "(,"), strings.Trim(line[3], ")")}
		if name[2] == 'A' {
			current = append(current, name)
		}
	}

	// follow instructions, starting at **A, counting steps to reach **Z simultaneously
	allSteps := []int{}

	// find step count for each node
	for _, node := range current {
		// follow instructions for this node
		steps := 0
		for ; node[2] != 'Z'; steps++ {
			if instructions[steps%len(instructions)] == 'L' {
				node = nodes[node].left
			} else {
				node = nodes[node].right
			}
		}
		allSteps = append(allSteps, steps)
	}

	return LCM(allSteps[0], allSteps[1], allSteps[2:]...)
}

// greatest common divisor (GCD) via Euclidean algorithm
// (stolen from go playground)
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
// (stolen from go playground)
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
