package main

import (
	"bufio"
	"os"
	"testing"
)

func Test_countSteps(t *testing.T) {
	tests := map[string]struct {
		scanner *bufio.Scanner
		want    int
	}{
		"example1": {want: 2},
		"example2": {want: 6},
		"input":    {want: 19241},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			file, err := os.Open("testdata/" + name)
			if err != nil {
				t.Errorf("unable to open file %s", name)
			}
			scanner := bufio.NewScanner(file)
			if got := countSteps(scanner); got != tt.want {
				t.Errorf("countSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
