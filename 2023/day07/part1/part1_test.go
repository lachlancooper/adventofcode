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
		"example": {want: 6440},
		"input":   {want: 250347426},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			file, err := os.Open("testdata/" + name)
			if err != nil {
				t.Errorf("unable to open file %s", name)
			}
			scanner := bufio.NewScanner(file)
			if got := winnings(scanner); got != tt.want {
				t.Errorf("countWins() = %v, want %v", got, tt.want)
			}
		})
	}
}
