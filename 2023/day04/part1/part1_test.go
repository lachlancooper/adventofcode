package main

import (
	"bufio"
	"os"
	"testing"
)

func Test_sumWins(t *testing.T) {
	tests := map[string]struct {
		scanner *bufio.Scanner
		want    int
	}{
		"example": {want: 13},
		"input":   {want: 32609},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			file, err := os.Open("testdata/" + name)
			if err != nil {
				t.Errorf("unable to open file %s", name)
			}
			scanner := bufio.NewScanner(file)
			if got := sumWins(scanner); got != tt.want {
				t.Errorf("sumWins() = %v, want %v", got, tt.want)
			}
		})
	}
}
