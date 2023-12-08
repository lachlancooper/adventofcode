package main

import (
	"bufio"
	"os"
	"testing"
)

func Test_countWins(t *testing.T) {
	tests := map[string]struct {
		scanner *bufio.Scanner
		want    int
	}{
		"example": {want: 35},
		"input":   {want: 993500720},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			file, err := os.Open("testdata/" + name)
			if err != nil {
				t.Errorf("unable to open file %s", name)
			}
			scanner := bufio.NewScanner(file)
			if got := countWins(scanner); got != tt.want {
				t.Errorf("countWins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wins(t *testing.T) {
	tests := map[string]struct {
		time     int
		distance int
		want     int
	}{
		"example1":     {7, 9, 4},
		"example2":     {15, 40, 8},
		"example3":     {30, 200, 9},
		"input1":       {40, 219, 27},
		"input2":       {81, 1012, 50},
		"input3":       {77, 1365, 22},
		"input4":       {72, 1089, 29},
		"examplepart2": {71530, 940200, 71503},
		"inputpart2":   {40_817_772, 219_101_213_651_089, 28_101_347},
		"other":        {250, 15000, 51},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := wins(tt.time, tt.distance); got != tt.want {
				t.Errorf("wins() = %v, want %v", got, tt.want)
			}
		})
	}
}
