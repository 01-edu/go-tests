package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {

	args := [][]string{
		{"A", "B"},
		{"a", "b"},
		{"A", "b"},
		{"A", "1"},
		{" "},
		{"Hello", "World"},
		{""},
		{"1", "2"},
		{"a", "B"},
	}
	for _, arg := range args {
		challenge.Program("samecase", arg...)
	}
}
