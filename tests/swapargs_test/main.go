package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	table := [][]string{
		{"choumi", "is", "the", "best", "cat"},
		{" "},
		{""},
		{"hello", "world"},
		{"1", "2"},
		{"", " "},
	}
	for _, s := range table {
		challenge.Program("swapargs", s...)
	}
}
