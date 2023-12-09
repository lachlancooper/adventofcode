package main

import (
	"bufio"
	"os"
	"testing"
)

func Test_countTotalsReverse(t *testing.T) {
	tests := map[string]struct {
		scanner *bufio.Scanner
		want    int
	}{
		"example3": {want: 5},
		"example":  {want: 2},
		"input":    {want: 803},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			file, err := os.Open("testdata/" + name)
			if err != nil {
				t.Errorf("unable to open file %s", name)
			}
			scanner := bufio.NewScanner(file)
			if got := countTotalsReverse(scanner); got != tt.want {
				t.Errorf("countTotalsReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
