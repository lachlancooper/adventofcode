package main

import (
	"bufio"
	"os"
	"testing"
)

func Test_sumParts(t *testing.T) {
	tests := map[string]struct {
		scanner *bufio.Scanner
		want    int
	}{
		"example": {want: 4361},
		"input":   {want: 532331},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			file, err := os.Open("testdata/" + name)
			if err != nil {
				t.Errorf("unable to open file %s", name)
			}
			scanner := bufio.NewScanner(file)
			if got := sumParts(scanner); got != tt.want {
				t.Errorf("sumIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}
