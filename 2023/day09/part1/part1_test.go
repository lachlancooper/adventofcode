package main

import (
	"bufio"
	"os"
	"testing"
)

func Test_countTotals(t *testing.T) {
	tests := map[string]struct {
		scanner *bufio.Scanner
		want    int
	}{
		"example1": {want: 18},
		"example":  {want: 114},
		"input":    {want: 1681758908},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			file, err := os.Open("testdata/" + name)
			if err != nil {
				t.Errorf("unable to open file %s", name)
			}
			scanner := bufio.NewScanner(file)
			if got := countTotals(scanner); got != tt.want {
				t.Errorf("countTotals() = %v, want %v", got, tt.want)
			}
		})
	}
}
