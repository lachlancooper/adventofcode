package main

import (
	"bufio"
	"os"
	"testing"
)

func Test_sumPaths(t *testing.T) {
	tests := map[string]struct {
		multiple int
		want     int
	}{
		"example": {multiple: 100, want: 8410},
		"input":   {multiple: 1_000_000, want: 752936133304},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			file, err := os.Open("testdata/" + name)
			if err != nil {
				t.Errorf("unable to open file %s", name)
			}
			scanner := bufio.NewScanner(file)
			if got := sumPaths(scanner, tt.multiple); got != tt.want {
				t.Errorf("sumPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
