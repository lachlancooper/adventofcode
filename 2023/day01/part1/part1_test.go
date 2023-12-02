package main

import (
	"bufio"
	"os"
	"testing"
)

func Test_calibrate(t *testing.T) {
	tests := map[string]struct {
		scanner *bufio.Scanner
		want    int
	}{
		"example": {want: 142},
		"input1":  {want: 56506},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			file, err := os.Open("testdata/" + name)
			if err != nil {
				t.Errorf("unable to open file %s", name)
			}
			scanner := bufio.NewScanner(file)
			if got := calibrate(scanner); got != tt.want {
				t.Errorf("calibrate() = %v, want %v", got, tt.want)
			}
		})
	}
}
