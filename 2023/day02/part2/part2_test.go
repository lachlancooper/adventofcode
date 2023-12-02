package main

import (
	"bufio"
	"os"
	"testing"
)

func Test_sumPowers(t *testing.T) {
	tests := map[string]struct {
		scanner *bufio.Scanner
		want    int
	}{
		"example": {want: 2286},
		"input1":  {want: 67335},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			file, err := os.Open("testdata/" + name)
			if err != nil {
				t.Errorf("unable to open file %s", name)
			}
			scanner := bufio.NewScanner(file)
			if got := sumPowers(scanner); got != tt.want {
				t.Errorf("sumIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}
