package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(solve(scanner))
}

type workflow struct {
	rules []rule
}

type rule struct {
	field  string
	op     string
	val    int
	target string
}

type partRange struct {
	minX, maxX int
	minM, maxM int
	minA, maxA int
	minS, maxS int
}

func solve(scanner *bufio.Scanner) int {
	workflows := make(map[string]workflow)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		l := strings.Split(strings.TrimSuffix(line, "}"), "{")
		if l[0] == "" {
			continue
		}

		// scan workflows
		workflows[l[0]] = workflow{scanRules(l[1])}
	}

	// test all possible parts
	min := 1
	max := 4000
	return countParts(workflows, partRange{
		minX: min, maxX: max,
		minM: min, maxM: max,
		minA: min, maxA: max,
		minS: min, maxS: max,
	}, "in")
}

func scanRules(input string) []rule {
	rules := []rule{}

	for _, r := range strings.Split(input, ",") {
		t := rule{}

		c := strings.Split(r, ":")
		if len(c) == 1 {
			t.target = c[0]
		} else {
			t.field = string(c[0][0])
			t.op = string(c[0][1])
			t.val, _ = strconv.Atoi(c[0][2:])
			t.target = c[1]
		}

		rules = append(rules, t)
	}

	return rules
}

func countParts(workflows map[string]workflow, p partRange, wf string) int {
	for {
		for _, rule := range workflows[wf].rules {
			var min, max int
			// preemptively split range around rule.val
			lowGt, highGt := p, p
			lowLt, highLt := p, p

			switch rule.field {
			case "x":
				min, max = p.minX, p.maxX
				lowGt.maxX, highGt.minX = rule.val, rule.val+1
				lowLt.maxX, highLt.minX = rule.val-1, rule.val
			case "m":
				min, max = p.minM, p.maxM
				lowGt.maxM, highGt.minM = rule.val, rule.val+1
				lowLt.maxM, highLt.minM = rule.val-1, rule.val
			case "a":
				min, max = p.minA, p.maxA
				lowGt.maxA, highGt.minA = rule.val, rule.val+1
				lowLt.maxA, highLt.minA = rule.val-1, rule.val
			case "s":
				min, max = p.minS, p.maxS
				lowGt.maxS, highGt.minS = rule.val, rule.val+1
				lowLt.maxS, highLt.minS = rule.val-1, rule.val
			}

			found := false
			switch rule.op {
			case ">":
				if min > rule.val {
					// whole range is sent to target workflow
					found = true
				} else if max > rule.val {
					// range is split:
					// - low continues evaluation
					// - high is sent to target workflow
					return countParts(workflows, lowGt, wf) + countParts(workflows, highGt, rule.target)
				}
			case "<":
				if max < rule.val {
					// whole range is sent to target workflow
					found = true
				} else if min < rule.val {
					// range is split:
					// - low is sent to target workflow
					// - high continues evaluation
					return countParts(workflows, lowLt, rule.target) + countParts(workflows, highLt, wf)
				}
			case "":
				// whole range is sent to default workflow
				found = true
			}

			if found {
				wf = rule.target
				break
			}
		}

		switch wf {
		case "A":
			return (p.maxX - p.minX + 1) * (p.maxM - p.minM + 1) * (p.maxA - p.minA + 1) * (p.maxS - p.minS + 1)
		case "R":
			return 0
		}
	}

	panic("Could not find result for part range")
}
