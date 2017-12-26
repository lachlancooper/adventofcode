# Advent of Code 2017

This is my repository of solutions to [AoC 2017](http://adventofcode.com/2017). I took it as a chance to get some practical (?) experience with Go, having completed only the language tour prior to this event.

My experience with the language was great. While many of the puzzles obviously suggested functional approaches, many were also very well suited to exploiting Go's zero values. A [trick](https://blog.golang.org/go-maps-in-action) I used a lot was storing "active" values in a map of `bool`s, deleting "inactive" values and relying on false being the zero value for the boolean type.

Generally my approach was very straightforward, with no goroutines and limited use of recursion. If I do this next year, or go back to earlier years, I might try to use goroutines where possible.

# Conventions

Every puzzle (except day 25) has two solutions, `part1.go` and `part2.go`.

(Nearly) all my solutions read puzzle input from stdin exactly as provided. I've included my input and any sample inputs as files named `input`, `input1`, `input2`, etc. The only exception is day 23 part 2 (which I could definitely modify to read starting parameters as input).
