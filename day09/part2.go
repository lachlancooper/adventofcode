// http://adventofcode.com/2017/day/9
// part 2
package main

import (
	"bufio"
	"fmt"
	"os"
)

// Garbage returns the total number of garbage characters
// within group(s) g
func Garbage(g string) (count int) {
	garbageMode := false
	cancelMode := false

	for _, c := range g {
		if cancelMode {
			cancelMode = false
			continue
		}

		switch {
		case !garbageMode && c == '<':
			garbageMode = true
		case garbageMode && c == '>':
			garbageMode = false
		case garbageMode && c == '!':
			cancelMode = true
		case garbageMode:
			count++
		}
	}

	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		garbage := Garbage(scanner.Text())
		fmt.Println(garbage)
	}
}
