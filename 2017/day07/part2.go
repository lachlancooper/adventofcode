// http://adventofcode.com/2017/day/7
// part 2
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

// totalWeight returns the total weight of n,
// including itself and all subtowers
func totalWeight(t *tower) (weight int) {
	weight = t.weight
	for _, s := range t.subtowers {
		weight += totalWeight(city[s])
	}
	return
}

// findUnbalanced returns the name of an unbalanced subtower
// and its ideal weight, searching from t
// both are zero valued if all subtowers are balanced
func findUnbalanced(t *tower) (string, int) {
	subtowerweights := make(map[int]string)
	weightcounts := make(map[int]int)

	for _, s := range t.subtowers {
		// is there an unbalanced tower higher up?
		badname, badweight := findUnbalanced(city[s])
		if badname != "" {
			return badname, badweight
		}

		// no, take note of total subtower weight
		weight := totalWeight(city[s])
		subtowerweights[weight] = city[s].name
		weightcounts[weight]++
	}

	// are *my* subtowers unbalanced?
	var commonWeight int
	var unbalancedTower string
	var unbalancedTowerWeight int
	for w, c := range weightcounts {
		if c == 1 {
			unbalancedTower = subtowerweights[w]
			unbalancedTowerWeight = w
		} else {
			commonWeight = w
		}
	}
	if unbalancedTower != "" {
		offset := commonWeight - unbalancedTowerWeight
		idealWeight := city[unbalancedTower].weight + offset
		return unbalancedTower, idealWeight
	}

	// no problems here
	return "", 0
}

func main() {
	buildCity()
	bottom := city[getBottom()]
	fmt.Println(findUnbalanced(bottom))
}
