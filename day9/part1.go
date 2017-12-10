// http://adventofcode.com/2017/day/9
// part 1
package main

import (
	"bufio"
	"fmt"
	"os"
)

// Score returns the total score of group(s) g
func Score(g string) (score int) {
	level := 0
	garbageMode := false
	cancelMode := false

	for _, c := range g {
		if cancelMode {
			cancelMode = false
			continue
		}

		switch {
		case !garbageMode && c == '{':
			level++
		case !garbageMode && c == '}':
			score += level
			level--
		case !garbageMode && c == '<':
			garbageMode = true
		case garbageMode && c == '>':
			garbageMode = false
		case garbageMode && c == '!':
			cancelMode = true
		}
	}

	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		score := Score(scanner.Text())
		fmt.Println(score)
	}
}
