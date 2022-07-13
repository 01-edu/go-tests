package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	args := [][]string{
		{""},
		{"1", "One", "ring!"},
		{"2", "testing spaces and #!*"},
		{"3", "more", "than", "three", "arguments"},
		{"10", "Upper anD LoWer cAsE"},
		{"0", "Hello", "World", "everyOne"},
		{"-1", "Hello", "World", "everyOne"},
		{"4", "Hello", "World", "everyOne"},
	}

	for _, v := range args {
		challenge.Program("argrotn", v...)
	}
}
