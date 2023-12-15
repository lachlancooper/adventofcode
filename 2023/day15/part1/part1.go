package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(solve(scanner))
}

func solve(scanner *bufio.Scanner) int {
	total := 0
	scanner.Scan()

	for _, step := range strings.Split(scanner.Text(), ",") {
		total += hash(step)
	}

	return total
}

func hash(s string) int {
	result := 0
	for _, char := range s {
		result += int(char)
		result *= 17
		result %= 256
	}
	return result
}
