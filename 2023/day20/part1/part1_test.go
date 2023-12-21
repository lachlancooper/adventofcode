package main

import (
	"bufio"
	"os"
	"testing"
)

func Test_solve(t *testing.T) {
	tests := map[string]struct {
		scanner *bufio.Scanner
		want    int
	}{
		"example1": {want: 32000000},
		"example2": {want: 11687500},
		"input":    {want: 711650489},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			file, err := os.Open("testdata/" + name)
			if err != nil {
				t.Errorf("unable to open file %s", name)
			}
			scanner := bufio.NewScanner(file)
			if got := solve(scanner); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
