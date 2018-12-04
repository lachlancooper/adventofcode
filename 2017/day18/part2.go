// http://adventofcode.com/2017/day/18
// part 2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// how many simultaneous programs?
const PROGS = 2

// instructions are the same for all programs
var ins = make([][]string, 0)

// one set of registers per program
var reg = make([]map[string]int, PROGS)

// each program receives a list values from another
var queue = make([][]int, PROGS)

// each program can be locked (waiting on input) independently
var locked = make([]bool, PROGS)

// each program has a separate pc
var pc = make([]int, 2)

// track sent messages per program
var sent = make([]int, 2)

// each program starts with its program ID in register p
func init() {
	for i := range reg {
		reg[i] = map[string]int{"p": i}
	}
}

// are all programs locked (i.e. deadlock)?
func deadlock() bool {
	for _, q := range locked {
		if !q {
			return false
		}
	}
	return true
}

// run program id until it locks
func runprog(id int) {
	for ; pc[id] < len(ins); pc[id]++ {
		// fetch instruction
		i := ins[pc[id]]
		op := i[0]
		dst := i[1]

		// decode operand values
		val := make([]int, 0)
		for _, c := range i[1:] {
			if v, err := strconv.Atoi(c); err == nil {
				val = append(val, v)
			} else {
				val = append(val, reg[id][c])
			}
		}

		// execute
		switch op {
		case "set":
			reg[id][dst] = val[1]
		case "add":
			reg[id][dst] += val[1]
		case "mul":
			reg[id][dst] *= val[1]
		case "mod":
			reg[id][dst] %= val[1]
		case "jgz":
			if val[0] > 0 {
				pc[id] += val[1] - 1
			}
		case "snd":
			// find id of next prog
			other := (id + 1) % PROGS

			// append to other queue
			queue[other] = append(queue[other], val[0])
			// unlock other prog
			locked[other] = false

			// stats
			sent[id]++
		case "rcv":
			if len(queue[id]) > 0 {
				// pop from queue
				reg[id][dst], queue[id] = queue[id][0], queue[id][1:]
			} else {
				// no queued value yet
				locked[id] = true
				return
			}
		}
	}

	// we ran out of instructions, lock
	locked[id] = true
}

// run until deadlocked
func run() {
	for !deadlock() {
		for i, q := range locked {
			if !q {
				runprog(i)
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

	// how many messages did program 1 send?
	fmt.Println(sent[1])
}
