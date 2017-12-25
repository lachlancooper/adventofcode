// http://adventofcode.com/2017/day/25
// part 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var tape = make(map[int]int)
var cursor int
var curstate string
var states = make(map[string]state)

type state struct {
	write [2]int
	move  [2]string
	next  [2]string
}

// atoi wraps strconv.Atoi, ignoring errors
func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

// run for the specified number of steps
func run(steps int) {
	for i := 0; i < steps; i++ {
		curval := tape[cursor]

		towrite := states[curstate].write[curval]
		tomove := states[curstate].move[curval]
		tonext := states[curstate].next[curval]

		// write
		switch towrite {
		case 0:
			delete(tape, cursor)
		case 1:
			tape[cursor] = towrite
		}

		// move
		switch tomove {
		case "left":
			cursor--
		case "right":
			cursor++
		}

		// next
		curstate = tonext
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// initial parameters
	scanner.Scan()
	curstate = strings.TrimRight(strings.Fields(scanner.Text())[3], ".")
	scanner.Scan()
	stepcount := atoi(strings.Fields(scanner.Text())[5])
	scanner.Scan()

	// states
	for scanner.Scan() {
		name := strings.TrimRight(strings.Fields(scanner.Text())[2], ":")
		scanner.Scan()

		scanner.Scan()
		write0 := atoi(strings.TrimRight(strings.Fields(scanner.Text())[4], "."))
		scanner.Scan()
		move0 := strings.TrimRight(strings.Fields(scanner.Text())[6], ".")
		scanner.Scan()
		next0 := strings.TrimRight(strings.Fields(scanner.Text())[4], ".")
		scanner.Scan()

		scanner.Scan()
		write1 := atoi(strings.TrimRight(strings.Fields(scanner.Text())[4], "."))
		scanner.Scan()
		move1 := strings.TrimRight(strings.Fields(scanner.Text())[6], ".")
		scanner.Scan()
		next1 := strings.TrimRight(strings.Fields(scanner.Text())[4], ".")
		scanner.Scan()

		states[name] = state{[2]int{write0, write1}, [2]string{move0, move1}, [2]string{next0, next1}}
	}

	run(stepcount)
	fmt.Println(len(tape))
}
