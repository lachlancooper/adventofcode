// http://adventofcode.com/2017/day/7
// part 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// tower has a name and weight. it may support
// any number of subtowers
type tower struct {
	name      string
	weight    int
	subtowers []string
}

// city holds pointers to all towers
var city = make(map[string]*tower)

// surtowers maps tower name to surtower name
var surtowers = make(map[string]string)

// buildCity returns a city object built by scanning every
// input line and inserting all towers found
func buildCity() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		// add a new tower with the given name and weight
		name := line[0]
		weight, _ := strconv.Atoi(strings.Trim(line[1], "()"))
		city[name] = &tower{name, weight, []string{}}

		// tower has no sub-towers
		if len(line) == 2 {
			continue
		}

		// link sub-towers
		subtowers := line[3:]
		for _, s := range subtowers {
			s = strings.TrimRight(s, ",")
			city[name].subtowers = append(city[name].subtowers, s)
			surtowers[s] = name
		}
	}
}

// getBottom returns the bottom tower of city
func getBottom() (bottom string) {
	for _, t := range surtowers {
		for bottom = t; ; {
			s, ok := surtowers[bottom]
			if !ok {
				return
			}
			bottom = s
		}
	}
	return
}

func main() {
	buildCity()
	fmt.Println(getBottom())
}
