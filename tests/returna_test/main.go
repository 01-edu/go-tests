package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	args := [][]string{
		{"Exam"},
		{""},
		{"Hello blah blah", "more blabh lah"},
		{"See you at ", "Monday"},
		{"a"},
		{"30", "20"},
		{"A B C E"},
		{"4 3 2 1 0"},
		{"A B C D E"},
		{"13 285 8528 -25a"},
		{"first go "},
		{"pool"},
	}

	for _, s := range args {
		challenge.Program("returna", s...)
	}
}
