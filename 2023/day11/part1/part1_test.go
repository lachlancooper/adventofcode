package main

import (
	"bufio"
	"os"
	"testing"
)

func Test_sumPaths(t *testing.T) {
	tests := map[string]struct {
		scanner *bufio.Scanner
		want    int
	}{
		"example": {want: 374},
		"input":   {want: 6890},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			file, err := os.Open("testdata/" + name)
			if err != nil {
				t.Errorf("unable to open file %s", name)
			}
			scanner := bufio.NewScanner(file)
			if got := sumPaths(scanner); got != tt.want {
				t.Errorf("sumPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
