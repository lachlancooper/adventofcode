package main

import (
	"bufio"
	"os"
	"testing"
)

func Test_lowestLocation(t *testing.T) {
	tests := map[string]struct {
		scanner *bufio.Scanner
		want    int
	}{
		"example": {want: 46},
		"input":   {want: 4917124},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			file, err := os.Open("testdata/" + name)
			if err != nil {
				t.Errorf("unable to open file %s", name)
			}
			scanner := bufio.NewScanner(file)
			if got := lowestLocation(scanner); got != tt.want {
				t.Errorf("lowestLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
