package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(lowestLocation(scanner))
}

// store initial seeds in slice

// for each map type:
// e.g.
// seed-to-soil map:
// 50 98 2
// 52 50 48
//
// dst src len

// seedMap stores a single type of map.
type seedMap struct {
	dstName string

	offsets []offset
}

// offset defines a range of IDs to convert from one type to another.
type offset struct {
	rangeStart int
	rangeEnd   int
	value      int
}

// lowestLocation finds the lowest location number that corresponds to any of the initial seeds.
func lowestLocation(scanner *bufio.Scanner) int {
	lowest := math.MaxInt
	var srcName string
	startSeeds := []int{}

	// almanac contains all maps.
	almanac := make(map[string]seedMap)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Split(line, ":")

		if len(parts) == 2 {
			if parts[0] == "seeds" {
				// we have the initial seed list
				for _, seed := range strings.Split(parts[1], " ") {
					if seed == "" {
						continue
					}
					val, _ := strconv.Atoi(seed)
					startSeeds = append(startSeeds, val)
				}
			} else {
				// use source map name as current map
				mapParts := strings.Split(strings.Split(parts[0], " ")[0], "-")
				srcName = mapParts[0]
				almanac[srcName] = seedMap{
					dstName: mapParts[2],
				}
			}
		} else {
			// read values into current map
			parts = strings.Split(line, " ")
			dstRange, _ := strconv.Atoi(parts[0])
			srcRange, _ := strconv.Atoi(parts[1])
			length, _ := strconv.Atoi(parts[2])

			currentMap := almanac[srcName]
			currentMap.offsets = append(currentMap.offsets,
				offset{
					rangeStart: srcRange,
					rangeEnd:   srcRange + length,
					value:      dstRange - srcRange,
				},
			)
			almanac[srcName] = currentMap
		}
	}

	// check all startSeeds
	for _, seed := range startSeeds {
		if result := testSeed(almanac, seed); result < lowest {
			lowest = result
		}
	}

	return lowest
}

// testSeed tests the path of a given seed through the almanac.
func testSeed(almanac map[string]seedMap, seed int) int {
	nextMap := "seed"

	for {
		fmt.Printf("%s %d, ", nextMap, seed)

		thisMap, ok := almanac[nextMap]
		if !ok {
			break
		}

		nextMap = thisMap.dstName

		// check for differing seed mapping from src to dst
		for _, mapping := range thisMap.offsets {
			if seed >= mapping.rangeStart && seed <= mapping.rangeEnd {
				seed += mapping.value
				break
			}
		}
		// otherwise seed value doesn't change
	}

	fmt.Printf("\n")

	return seed
}
