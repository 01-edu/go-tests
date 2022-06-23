package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	args := [][]string{
		{"S"},
		{"One", "ring!"},
		{"A"},
		{"E"},
		{"z"},
		{"testing spaces and #!*"},
		{"more", "than", "three", "arguments"},
		{"W"},
		{"1"},
		{""},
	}
	for _, v := range args {
		challenge.Program("printascii", v...)
	}
}
