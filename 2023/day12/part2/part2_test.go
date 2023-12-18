package main

import (
	"bufio"
	"os"
	"testing"
)

func Test_solve(t *testing.T) {
	tests := map[string]struct {
		scanner *bufio.Scanner
		want    int
	}{
		"example1": {want: 1},
		"example2": {want: 16384},
		"example3": {want: 1},
		"example4": {want: 16},
		"example5": {want: 2500},
		// "example6": {want: 506250},
		// "example": {want: 525152},
		// "input":   {want: 6935},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			file, err := os.Open("testdata/" + name)
			if err != nil {
				t.Errorf("unable to open file %s", name)
			}
			scanner := bufio.NewScanner(file)
			if got := solve(scanner); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_place(t *testing.T) {
	tests := []struct {
		input   string
		expect  int
		want    string
		success bool
	}{
		{"?", 1, "*", true},
		{"?#", 1, "", false},
		{".", 1, "", false},
		{"#", 2, "", false},
		{"##", 2, "**", true},
		{"###", 2, "", false},
		{"#?#", 3, "***", true},
		{"#???.", 2, "**.?.", true},
		{".#.", 1, ".*.", true},
		{".#", 1, ".*", true},
		{"#", 1, "*", true},
		{"#?", 2, "**", true},
		{".#?", 2, ".**", true},
		{".#??", 2, ".**.", true},
		{".#???", 2, ".**.?", true},
		{".#?.#?.#?.#?.#", 1, ".*..#?.#?.#?.#", true},
		{".?#.#?.#?.#?.#", 1, "", false},
		{"*.*.***.*.*.***.*.*.***.*.*.***.*.*.###", 3, "*.*.***.*.*.***.*.*.***.*.*.***.*.*.***", true},
		{"*....******..*****.*.....******..*****.*.....******..*****.*.....******..*****.*.....******..#####.", 5, "*....******..*****.*.....******..*****.*.....******..*****.*.....******..*****.*.....******..*****.", true},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, success := place(tt.input, tt.expect)
			if got != tt.want && tt.success {
				t.Errorf("place() got = %v, want %v", got, tt.want)
			}
			if success != tt.success {
				t.Errorf("place() success = %v, want %v", success, tt.success)
			}
		})
	}
}
