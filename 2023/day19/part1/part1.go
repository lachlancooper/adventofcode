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

type part struct {
	x, m, a, s int
}

func solve(scanner *bufio.Scanner) int {
	total := 0
	workflows := make(map[string]workflow)
	parts := []part{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		l := strings.Split(strings.TrimSuffix(line, "}"), "{")

		if l[0] == "" {
			// scan parts
			parts = append(parts, scanPart(l[1]))
		} else {
			// scan workflows
			workflows[l[0]] = workflow{scanRules(l[1])}
		}
	}

	// test parts
	for _, part := range parts {
		if testPart(workflows, part) {
			total += part.x + part.m + part.a + part.s
		}
	}

	return total
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

func scanPart(input string) part {
	result := part{}

	for _, rating := range strings.Split(input, ",") {
		bits := strings.Split(rating, "=")
		val, _ := strconv.Atoi(bits[1])

		switch bits[0] {
		case "x":
			result.x = val
		case "m":
			result.m = val
		case "a":
			result.a = val
		case "s":
			result.s = val
		}
	}

	return result
}

func testPart(workflows map[string]workflow, p part) bool {
	wf := "in"

	for {
		for _, rule := range workflows[wf].rules {
			var v int
			switch rule.field {
			case "x":
				v = p.x
			case "m":
				v = p.m
			case "a":
				v = p.a
			case "s":
				v = p.s
			}

			found := false
			switch rule.op {
			case ">":
				if v > rule.val {
					found = true
				}
			case "<":
				if v < rule.val {
					found = true
				}
			case "":
				found = true
			}

			if !found {
				continue
			}

			wf = rule.target
			break
		}

		switch wf {
		case "A":
			return true
		case "R":
			return false
		}
	}

	panic("Could not find result for part")
}
