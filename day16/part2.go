// http://adventofcode.com/2017/day/16
// part 2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// naive solution is way too slow
// 2.5s for 1000 dances, ~29 days for 1B dances
// all dances are identical, so all we really care
// about is the start and end

// build a mapping from start to end and apply that
// in a single operation

// still won't be efficient, but only in the region of
// 1B things to process rather than 10T+

func spin(dance string, d int) string {
	return dance[len(dance)-d:] + dance[:len(dance)-d]
}

func exchange(dance string, a, b int) string {
	bytea, byteb := dance[a], dance[b]
	dance = dance[:a] + string(byteb) + dance[a+1:]
	dance = dance[:b] + string(bytea) + dance[b+1:]
	return dance
}

func partner(dance string, bytea, byteb byte) string {
	var a, b int
	for i := 0; i < len(dance); i++ {
		c := dance[i]
		switch c {
		case bytea:
			a = i
		case byteb:
			b = i
		}
	}
	dance = dance[:a] + string(byteb) + dance[a+1:]
	dance = dance[:b] + string(bytea) + dance[b+1:]
	return dance
}

// round calculates the positions after a single round of the given dance
func round(dance string, moves []string) string {
	for _, move := range moves {
		switch move[0] {
		case 's':
			size, _ := strconv.Atoi(move[1:])
			dance = spin(dance, size)
		case 'x':
			ab := strings.Split(move[1:], "/")
			a, _ := strconv.Atoi(ab[0])
			b, _ := strconv.Atoi(ab[1])
			dance = exchange(dance, a, b)
		case 'p':
			a := move[1]
			b := move[3]
			dance = partner(dance, a, b)
		}
	}
	return dance
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		dance := "abcdefghijklmnop"

		for i := 0; i < 37; i++ {
			fmt.Println(i, dance)
			dance = round(dance, strings.Split(line, ","))
		}
		fmt.Println(37, dance)
		fmt.Println((1000 * 1000 * 1000) % 36)
	}
}
