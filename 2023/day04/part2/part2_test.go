package main

import (
	"bufio"
	"os"
	"testing"
)

func Test_sumCards(t *testing.T) {
	tests := map[string]struct {
		scanner *bufio.Scanner
		want    int
	}{
		"example": {want: 30},
		"input":   {want: 14624680},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			file, err := os.Open("testdata/" + name)
			if err != nil {
				t.Errorf("unable to open file %s", name)
			}
			scanner := bufio.NewScanner(file)
			if got := sumCards(scanner); got != tt.want {
				t.Errorf("sumCards() = %v, want %v", got, tt.want)
			}
		})
	}
}
