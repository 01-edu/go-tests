package main

import (
	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	args := [][]string{
		{""},
		{"One", "ring!"},
		{"testing spaces and #!*"},
		{"more", "than", "three", "arguments"},
		{"Upper anD LoWer cAsE"},
		{lib.RandWords()},
	}

	for _, v := range args {
		challenge.Program("alphamirror", v...)
	}
}
