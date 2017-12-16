// http://adventofcode.com/2017/day/16
// part 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// dance := "abcde"
		dance := "abcdefghijklmnop"

		moves := strings.Split(scanner.Text(), ",")
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
		fmt.Println(dance)
	}
}
