package main

import "github.com/01-edu/go-tests/lib/challenge"

func main() {
	args := [][]string{
		{"S", "A"},
		{"One", "ring!"},
		{"A", "B"},
		{"E", "Z"},
		{"z", "R"},
		{"testing spaces and #!*"},
		{"more", "than", "three", "arguments"},
		{"W"},
		{"1"},
		{" "},
		{""},
		{},
	}
	for _, arg := range args {
		challenge.Program("sumascii", arg...)
	}
}
