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
		"example": {want: 405},
		"input":   {want: 35360},
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
