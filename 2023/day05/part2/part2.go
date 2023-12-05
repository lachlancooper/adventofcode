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

// seedMap stores a single type of map.
type seedMap struct {
	dstName string

	offsets []offset
}

// offset defines a range of IDs to convert from one type to another.
type offset struct {
	start int
	end   int
	value int
}

// seedRange represents a range of seeds.
type seedRange struct {
	start int
	end   int
}

// lowestLocation finds the lowest location number that corresponds to any of the initial seeds.
func lowestLocation(scanner *bufio.Scanner) int {
	lowest := math.MaxInt
	var srcName string
	startSeeds := []seedRange{}

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
				seeds := strings.Split(parts[1], " ")

				// range over each pair of seeds
				for i := 1; i < len(seeds); i += 2 {
					start, _ := strconv.Atoi(seeds[i])
					length, _ := strconv.Atoi(seeds[i+1])
					startSeeds = append(startSeeds, seedRange{
						start: start,
						end:   start + length - 1,
					})
					fmt.Printf("Identified start seed range: %d-%d.\n", start, start+length-1)
				}
			} else {
				// we have a map name, add it to the almanac
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
					start: srcRange,
					end:   srcRange + length - 1,
					value: dstRange - srcRange,
				},
			)
			almanac[srcName] = currentMap
		}
	}

	// check all startSeeds
	for _, seed := range startSeeds {
		if result := testSeedRange(almanac, seed, "seed"); result < lowest {
			fmt.Printf("\tFound new lowest for range %d-%d: %d\n", seed.start, seed.end, result)
			lowest = result
		}
	}

	return lowest
}

// testSeedRange finds the lowest value for a given seed range through the almanac, given a starting map name.
func testSeedRange(almanac map[string]seedMap, seed seedRange, mapName string) int {
	fmt.Printf("testing range %d-%d in %s, ", seed.start, seed.end, mapName)

	// if we don't find our map, we're at the last one
	// so lowest is just the lowest value in our range
	thisMap, ok := almanac[mapName]
	if !ok {
		fmt.Printf("did not find map %s, returning %d.\n", mapName, seed.start)
		return seed.start
	}

	nextMap := thisMap.dstName

	// check all custom mappings in this map
	for _, mapping := range thisMap.offsets {
		// our whole range is totally outside this mapping
		// move on to the next one
		if seed.end < mapping.start || seed.start > mapping.end {
			continue
		}

		// our range is totally inside this mapping
		// so the lowest value must be the lowest (mapped) value in our range
		// call the next iteration
		if seed.start >= mapping.start && seed.end <= mapping.end {
			fmt.Printf("found containing mapping %d-%d, ", mapping.start, mapping.end)
			return testSeedRange(almanac,
				seedRange{
					start: seed.start + mapping.value,
					end:   seed.end + mapping.value,
				},
				nextMap,
			)
		}

		// we have an overlap, so break our range down into sub-ranges within this map
		// seed      start |        | end
		// mapping        start |
		if seed.start < mapping.start && seed.end > mapping.start {
			fmt.Printf("breaking into sub-ranges: %d-%d and %d-%d.\n", seed.start, mapping.start-1, mapping.start, seed.end)
			return lowest(
				testSeedRange(almanac,
					seedRange{
						start: seed.start,
						end:   mapping.start - 1,
					},
					mapName,
				),
				testSeedRange(almanac,
					seedRange{
						start: mapping.start,
						end:   seed.end,
					},
					mapName,
				),
			)
		}

		// we have an overlap, so break our range down into sub-ranges within this map
		// seed           start |        | end
		// mapping   start |        | end
		if seed.start >= mapping.start && seed.end > mapping.end {
			fmt.Printf("breaking into sub-ranges: %d-%d and %d-%d.\n", seed.start, mapping.end, mapping.end+1, seed.end)
			return lowest(
				testSeedRange(almanac,
					seedRange{
						start: seed.start,
						end:   mapping.end - 1,
					},
					mapName,
				),
				testSeedRange(almanac,
					seedRange{
						start: mapping.end + 1,
						end:   seed.end,
					},
					mapName,
				),
			)
		}
	}

	// we didn't find any matching mapping, so just check our range on the next map
	return testSeedRange(almanac, seed, nextMap)
}

func lowest(x, y int) int {
	if x < y {
		return x
	}
	return y
}
