package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(solve(scanner))
}

type machine struct {
	modules map[string]*module
	pulses  []pulse
}

type module struct {
	moduleType
	name string
	// state of the flipflop, false = off, true = on
	on bool
	// input modules, with most recently-seen pulse type
	inputs map[string]pulseType
	// destination modules
	destinations []string
}

type moduleType int

const (
	flipflop moduleType = iota
	conjunction
	broadcast
)

type pulse struct {
	pulseType
	src string
	dst string
}

type pulseType int

const (
	low pulseType = iota
	high
)

func (t pulseType) String() string {
	switch t {
	case low:
		return "low"
	default:
		return "high"
	}
}

func solve(scanner *bufio.Scanner) int {
	thisMachine := machine{}
	thisMachine.modules = make(map[string]*module)

	// read modules and connections
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " -> ")
		name, moduleType := detectMod(line[0])

		thisMachine.modules[name] = &module{
			moduleType:   moduleType,
			name:         name,
			inputs:       make(map[string]pulseType),
			destinations: strings.Split(line[1], ", "),
		}
	}

	// track inputs on each conjunction module
	for _, mod := range thisMachine.modules {
		for _, dst := range mod.destinations {
			target, ok := thisMachine.modules[dst]
			if !ok {
				continue
			}
			// if the destination is not a conjunction, skip it
			if target.moduleType != conjunction {
				continue
			}
			// the destination is a conjunction, add it to inputs, with seen set to low
			target.inputs[mod.name] = low
		}
	}

	highCount, lowCount := 0, 0
	presses := 1000
	for i := 0; i < presses; i++ {
		// add a single low pulse from button to broadcaster
		thisMachine.pulses = []pulse{{low, "button", "broadcaster"}}

		newHigh, newLow := thisMachine.run()
		highCount += newHigh
		lowCount += newLow

		fmt.Println()
	}

	return highCount * lowCount
}

// run simulates the flow of all pulses through a machine.
// Returns a count of all the high and low pulses sent.
func (m *machine) run() (int, int) {
	highCount, lowCount := 0, 0

	for len(m.pulses) != 0 {
		sentPulses := []pulse{}

		// send all current pulses, collecting output pulses
		for _, p := range m.pulses {
			fmt.Printf("%s -%s-> %s\n", p.src, p.pulseType, p.dst)

			switch p.pulseType {
			case low:
				lowCount++
			case high:
				highCount++
			}

			target, ok := m.modules[p.dst]
			if !ok {
				continue
			}
			sentPulses = append(sentPulses, target.receive(p)...)
		}

		m.pulses = sentPulses
	}

	return highCount, lowCount
}

// receive simulates a module receiving a single pulse.
func (mod *module) receive(in pulse) []pulse {
	var outType pulseType

	switch mod.moduleType {
	case flipflop:
		// if flipflop receives high pulse, it is ignored
		if in.pulseType == high {
			return nil
		}

		// if it receives low pulse, its state is flipped
		// if it was off, it sends a high pulse
		// if it was on, it sends a low pulse
		outType = high
		if mod.on {
			outType = low
		}
		mod.on = !mod.on
	case conjunction:
		// conjunction remembers the type of the most recent pulse from each input
		mod.inputs[in.src] = in.pulseType

		// if all inputs are high pulses, it sends a low pulse
		// otherwise, it sends a high pulse
		outType = low
		for _, seen := range mod.inputs {
			if seen == low {
				outType = high
				break
			}
		}
	case broadcast:
		// broadcast sends the same pulse to all destination modules
		outType = in.pulseType
	}

	// send out pulses to destinations
	output := []pulse{}
	for _, dst := range mod.destinations {
		output = append(output, pulse{outType, mod.name, dst})
	}

	return output
}

// detectMod detects the name and type of a module based on its name.
func detectMod(s string) (string, moduleType) {
	switch s[:1] {
	case "%":
		return s[1:], flipflop
	case "&":
		return s[1:], conjunction
	default:
		return s, broadcast
	}
}
